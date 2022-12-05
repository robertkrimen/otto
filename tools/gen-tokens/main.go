// Command gen-tokens generates go representations of JavaScript tokens.
package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

//go:embed .gen-tokens.yaml
var configData []byte

//go:embed templates/*
var templates embed.FS

// token represents a JavaScript token.
type token struct {
	Group  string `yaml:"group"`
	Name   string `yaml:"name"`
	Symbol string `yaml:"symbol"`
	Future bool   `yaml:"future"`
	Strict bool   `yaml:"strict"`
}

// config represents our configuration.
type config struct {
	Tokens []token `yaml:"tokens"`
}

// generate generates the context file writing the output to filename.
func generate(filename string) (err error) {
	var cfg config
	if err := yaml.Unmarshal(configData, &cfg); err != nil {
		return fmt.Errorf("decode config: %w", err)
	}

	tmpl := template.New("base").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
	})

	tmpl, err = tmpl.ParseFS(templates, "templates/*.tmpl")
	if err != nil {
		return fmt.Errorf("parse templates: %w", err)
	}

	output, err := os.Create(filename) //nolint: gosec
	if err != nil {
		return fmt.Errorf("open output: %w", err)
	}

	defer func() {
		if errc := output.Close(); err == nil && errc != nil {
			err = errc
		}
	}()

	if err := tmpl.ExecuteTemplate(output, "root.tmpl", cfg); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	cmd := exec.Command("gofmt", "-w", filename)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("format output %q: %w", string(buf), err)
	}

	return nil
}

func main() {
	var filename string
	flag.StringVar(&filename, "output", "token_const.go", "the filename to write the generated code to")
	if err := generate(filename); err != nil {
		log.Fatal(err)
	}
}
