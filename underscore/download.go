//go:build generate

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	url    = flag.String("url", "", "url to read from")
	output = flag.String("output", "", "output file to write the result too")
)

func download(url, output string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("new request failed: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var f *os.File
	if output != "" {
		if f, err = os.Create(output); err != nil {
			return fmt.Errorf("create file %q failed: %w", output, err)
		}

		defer f.Close()
	} else {
		f = os.Stdout
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		return fmt.Errorf("body save: %w", err)
	}

	return nil
}

func main() {
	flag.Parse()

	switch {
	case len(*url) == 0:
		log.Fatal("missing required --url parameter")
	}

	if err := download(*url, *output); err != nil {
		log.Fatal(err)
	}
}
