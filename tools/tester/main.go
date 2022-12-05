// Command tester automates the ability to download a suite of JavaScript libraries from a CDN and check if otto can handle them.
//
// It provides two commands via flags:
// * -fetch = Fetch all libraries from the CDN and store them in local testdata directory.
// * -report [file1 file2 ... fileN] = Report the results of trying to run the given or if empty all libraries in the testdata directory.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/parser"
)

const (
	// dataDir is where the libraries are downloaded to.
	dataDir = "testdata"

	// downloadWorkers is the number of workers that process downloads.
	downloadWorkers = 40

	// librariesURL is the source for JavaScript libraries for testing.
	librariesURL = "http://api.cdnjs.com/libraries"

	// requestTimeout is the maximum time we wait for a request to complete.
	requestTimeout = time.Second * 20
)

var (
	// testWorkers is the number of workers that process report.
	testWorkers = min(10, runtime.GOMAXPROCS(0))

	// noopConsole is a noopConsole which ignore log requests.
	noopConsole = map[string]interface{}{
		"log": func(call otto.FunctionCall) otto.Value {
			return otto.UndefinedValue()
		},
	}
)

var (
	matchReferenceErrorNotDefined = regexp.MustCompile(`^ReferenceError: \S+ is not defined$`)
	matchLookahead                = regexp.MustCompile(`Invalid regular expression: re2: Invalid \(\?[=!]\) <lookahead>`)
	matchBackReference            = regexp.MustCompile(`Invalid regular expression: re2: Invalid \\\d <backreference>`)
	matchTypeErrorUndefined       = regexp.MustCompile(`^TypeError: Cannot access member '[^']+' of undefined$`)
)

// broken identifies libraries which fail with a fatal error, so must be skipped.
var broken = map[string]string{
	"lets-plot.js":      "stack overflow",
	"knockout-es5.js":   "stack overflow",
	"sendbird-calls.js": "runtime: out of memory",
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// libraries represents fetch all libraries response.
type libraries struct {
	Results []library `json:"results"`
}

// library represents a single library in a libraries response.
type library struct {
	Name   string `json:"name"`
	Latest string `json:"latest"`
}

// fetch fetches itself and stores it in the dataDir.
func (l library) fetch() error {
	if !strings.HasSuffix(l.Latest, ".js") {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.Latest, nil)
	if err != nil {
		return fmt.Errorf("request library %q: %w", l.Name, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("get library %q: %w", l.Name, err)
	}
	defer resp.Body.Close() //nolint: errcheck

	name := l.Name
	if !strings.HasSuffix(name, ".js") {
		name += ".js"
	}

	f, err := os.Create(filepath.Join(dataDir, name)) //nolint: gosec
	if err != nil {
		return fmt.Errorf("create library %q: %w", l.Name, err)
	}

	if _, err = io.Copy(f, resp.Body); err != nil {
		return fmt.Errorf("write library %q: %w", l.Name, err)
	}

	return nil
}

// test runs the code from filename returning the time it took and any error
// encountered when running a full parse without IgnoreRegExpErrors in parseError.
func test(filename string) (took time.Duration, parseError, err error) { //nolint: nonamedreturns
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic on %q: %v", filename, r)
		}
	}()
	now := time.Now()
	defer func() {
		// Always set took.
		took = time.Since(now)
	}()

	if val := broken[filepath.Base(filename)]; val != "" {
		return 0, nil, fmt.Errorf("fatal %q", val)
	}

	script, err := os.ReadFile(filename) //nolint: gosec
	if err != nil {
		return 0, nil, err
	}

	vm := otto.New()
	if err := vm.Set("console", noopConsole); err != nil {
		return 0, nil, fmt.Errorf("set console: %w", err)
	}

	prog, err := parser.ParseFile(nil, filename, string(script), 0)
	if err != nil {
		val := err.Error()
		switch {
		case matchReferenceErrorNotDefined.MatchString(val),
			matchTypeErrorUndefined.MatchString(val),
			matchLookahead.MatchString(val),
			matchBackReference.MatchString(val):
			// RegExp issues, retry with IgnoreRegExpErrors.
			parseError = err
			if _, err = parser.ParseFile(nil, filename, string(script), parser.IgnoreRegExpErrors); err != nil {
				return 0, nil, err
			}
			return 0, parseError, nil
		default:
			return 0, nil, err
		}
	}

	_, err = vm.Run(prog)
	return 0, nil, err
}

// fetchAll fetches all files from src.
func fetchAll(src string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, librariesURL, nil)
	if err != nil {
		return fmt.Errorf("request libraries %q: %w", src, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("get libraries %q: %w", src, err)
	}
	defer resp.Body.Close() //nolint: errcheck

	var libs libraries
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&libs); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	if err := os.Mkdir(dataDir, 0o750); err != nil && !errors.Is(err, fs.ErrExist) {
		return fmt.Errorf("mkdir: %w", err)
	}

	var wg sync.WaitGroup
	work := make(chan library, downloadWorkers)
	errs := make(chan error, len(libs.Results))
	wg.Add(downloadWorkers)
	for i := 0; i < downloadWorkers; i++ {
		go func() {
			defer wg.Done()
			for lib := range work {
				fmt.Fprint(os.Stdout, ".")
				errs <- lib.fetch()
			}
		}()
	}

	fmt.Fprintf(os.Stdout, "Downloading %d libraries with %d workers ...", len(libs.Results), downloadWorkers)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, lib := range libs.Results {
			work <- lib
		}
		close(work)
	}()

	wg.Wait()
	close(errs)
	fmt.Fprintln(os.Stdout, " done")

	for e := range errs {
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			err = e
		}
	}

	return err
}

// result represents the result from a test.
type result struct {
	filename   string
	err        error
	parseError error
	took       time.Duration
}

// report runs test for all specified files, if none a specified all
// JavaScript files in our dataDir, outputting the results.
func report(files []string) error {
	if len(files) == 0 {
		var err error
		files, err = filepath.Glob(filepath.Join(dataDir, "*.js"))
		if err != nil {
			return fmt.Errorf("read dir: %w", err)
		}
	}

	var wg sync.WaitGroup
	workers := min(testWorkers, len(files))
	work := make(chan string, workers)
	results := make(chan result, len(files))
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for f := range work {
				fmt.Fprint(os.Stdout, ".")
				took, parseError, err := test(f)
				results <- result{
					filename:   f,
					err:        err,
					parseError: parseError,
					took:       took,
				}
			}
		}()
	}

	fmt.Fprintf(os.Stdout, "Reporting on %d libs with %d workers...", len(files), workers)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, f := range files {
			work <- f
		}
		close(work)
	}()

	wg.Wait()
	close(results)
	fmt.Fprintln(os.Stdout, " done")

	var fail, pass, parse int
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(writer, "Library", "\t| Took", "\t| Status")
	fmt.Fprintln(writer, "-------", "\t| ----", "\t| ------")

	for res := range results {
		switch {
		case res.err != nil:
			fmt.Fprintf(writer, "%s\t| %v\t| fail: %v\n", res.filename, res.took, res.err)
			fail++
		case res.parseError != nil:
			fmt.Fprintf(writer, "%s\t| %v\t| pass parse: %v\n", res.filename, res.took, res.parseError)
			parse++
		default:
			fmt.Fprintf(writer, "%s\t| %v\t| pass\n", res.filename, res.took)
			pass++
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush: %w", err)
	}

	fmt.Fprintf(os.Stdout, "\nSummary:\n - %d passes\n - %d parse passes\n - %d fails\n", pass, parse, fail)

	return nil
}

func main() {
	flagFetch := flag.Bool("fetch", false, "fetch all libraries for testing")
	flagReport := flag.Bool("report", false, "test and report the named files or all libraries if non specified")
	flag.Parse()

	var err error
	switch {
	case *flagFetch:
		err = fetchAll(librariesURL)
	case *flagReport:
		err = report(flag.Args())
	default:
		flag.PrintDefaults()
		err = fmt.Errorf("missing flag")
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(64)
	}
}
