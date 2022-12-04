package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/parser"
)

var (
	flagTest   *bool = flag.Bool("test", false, "")
	flagTeport *bool = flag.Bool("report", false, "")
)

var (
	matchReferenceErrorNotDefined = regexp.MustCompile(`^ReferenceError: \S+ is not defined$`)
	matchLookahead                = regexp.MustCompile(`Invalid regular expression: re2: Invalid \(\?[=!]\) <lookahead>`)
	matchBackreference            = regexp.MustCompile(`Invalid regular expression: re2: Invalid \\\d <backreference>`)
	matchTypeErrorUndefined       = regexp.MustCompile(`^TypeError: Cannot access member '[^']+' of undefined$`)
)

var target = map[string]string{
	"test-angular-bindonce.js": "fail",  // (anonymous): Line 1:944 Unexpected token ( (and 40 more errors)
	"test-jsforce.js":          "fail",  // (anonymous): Line 9:28329 RuneError (and 5 more errors)
	"test-chaplin.js":          "parse", // Error: Chaplin requires Common.js or AMD modules
	"test-dropbox.js.js":       "parse", // Error: dropbox.js loaded in an unsupported JavaScript environment.
	"test-epitome.js":          "parse", // TypeError: undefined is not a function
	"test-portal.js":           "parse", // TypeError
	"test-reactive-coffee.js":  "parse", // Dependencies are not met for reactive: _ and $ not found
	"test-scriptaculous.js":    "parse", // script.aculo.us requires the Prototype JavaScript framework >= 1.6.0.3
	"test-waypoints.js":        "parse", // TypeError: undefined is not a function
	"test-webuploader.js":      "parse", // Error: `jQuery` is undefined
	"test-xuijs.js":            "parse", // TypeError: undefined is not a function
}

// http://cdnjs.com/
// http://api.cdnjs.com/libraries

type libraries struct {
	Results []library `json:"results"`
}

type library struct {
	Name   string `json:"name"`
	Latest string `json:"latest"`
}

func (l library) fetch() error {
	if !strings.HasSuffix(l.Latest, ".js") {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
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

	f, err := os.Create("test-" + l.Name + ".js")
	if err != nil {
		return fmt.Errorf("create library %q: %w", l.Name, err)
	}

	if _, err = io.Copy(f, resp.Body); err != nil {
		return fmt.Errorf("write library %q: %w", l.Name, err)
	}

	return nil
}

func test(filename string) error {
	script, err := os.ReadFile(filename) //nolint: gosec
	if err != nil {
		return err
	}

	if !*flagTeport {
		fmt.Fprintln(os.Stdout, filename, len(script))
	}

	parse := false
	if target[filename] != "parse" {
		vm := otto.New()
		if _, err = vm.Run(string(script)); err != nil {
			value := err.Error()
			switch {
			case matchReferenceErrorNotDefined.MatchString(value),
				matchTypeErrorUndefined.MatchString(value),
				matchLookahead.MatchString(value),
				matchBackreference.MatchString(value):
			default:
				return err
			}
			parse = true
		}
	}

	if parse {
		_, err = parser.ParseFile(nil, filename, string(script), parser.IgnoreRegExpErrors)
		if err != nil {
			return err
		}
		target[filename] = "parse"
	}

	return nil
}

func fetchAll() error {
	resp, err := http.Get("http://api.cdnjs.com/libraries") //nolint: noctx
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint: errcheck

	var libs libraries
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&libs); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	var wg sync.WaitGroup
	errs := make(chan error, 5)
	for _, lib := range libs.Results {
		wg.Add(1)
		go func(lib library) {
			defer wg.Done()
			errs <- lib.fetch()
		}(lib)
	}

	defer func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func report() error {
	files, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("read dir: %w", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(writer, "", "\t| Status")
	fmt.Fprintln(writer, "---", "\t| ---")
	for _, file := range files {
		filename := file.Name()
		if !strings.HasPrefix(filename, "test-") {
			continue
		}
		err := test(filename)
		option := target[filename]
		name := strings.TrimPrefix(strings.TrimSuffix(filename, ".js"), "test-")
		if err != nil {
			fmt.Fprintln(writer, name, "\t| fail")
			continue
		}

		switch option {
		case "":
			fmt.Fprintln(writer, name, "\t| pass")
		case "parse":
			fmt.Fprintln(writer, name, "\t| pass (parse)")
		case "re2":
			fmt.Fprintln(writer, name, "\t| unknown (re2)")
		}
	}
	return writer.Flush()
}

func main() {
	flag.Parse()

	var filename string
	err := func() error {
		if flag.Arg(0) == "fetch" {
			return fetchAll()
		}

		if *flagTeport {
			return report()
		}

		filename = flag.Arg(0)
		return test(filename)
	}()
	if err != nil {
		if filename != "" {
			if *flagTest && target[filename] == "fail" {
				goto exit
			}
			fmt.Fprintf(os.Stderr, "%s: %s\n", filename, err.Error())
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(64)
	}
exit:
}
