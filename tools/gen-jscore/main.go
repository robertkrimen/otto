// Command gen-jscore generates go representations of JavaScript core types file.
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

//go:embed .gen-jscore.yaml
var configData []byte

//go:embed templates/*
var templates embed.FS

// jsType represents JavaScript type to generate.
type jsType struct {
	Name            string     `yaml:"name"`
	Core            bool       `yaml:"core"`
	ObjectClass     string     `yaml:"objectClass"`
	ObjectPrototype string     `yaml:"objectPrototype"`
	Class           string     `yaml:"class"`
	Value           string     `yaml:"value"`
	Properties      []property `yaml:"properties"`
	Prototype       *prototype `yaml:"prototype"`
}

// BlankConstructor is a default fallback returning false for templates.
func (t jsType) BlankConstructor() bool {
	return false
}

// prototype represents a JavaScript prototype to generate.
type prototype struct {
	Value       string     `yaml:"value"`
	ObjectClass string     `yaml:"objectClass"`
	Prototype   string     `yaml:"prototype"`
	Properties  []property `yaml:"properties"`
}

// Property returns the property with the given name.
func (p prototype) Property(name string) (*property, error) {
	for _, prop := range p.Properties {
		if prop.Name == name {
			return &prop, nil
		}
	}

	return nil, fmt.Errorf("missing property %q", name)
}

// property represents a JavaScript property to generate.
type property struct {
	Name     string `yaml:"name"`
	Call     string `yaml:"call"`
	Function int    `yaml:"function"`
	Mode     string `yaml:"mode"`
	Value    string `yaml:"value"`
	Kind     string `yaml:"kind"`
}

// value represents a JavaScript value to generate a Value creator for.
type value struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// config represents our configuration.
type config struct {
	Types  []jsType `yaml:"types"`
	Log    jsType   `yaml:"log"`
	Values []value  `yaml:"values"`
}

// Type returns the type for name.
func (c config) Type(name string) (*jsType, error) {
	for _, t := range c.Types {
		if t.Name == name {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("missing type %q", name)
}

// generate generates the context file writing the output to filename.
func generate(filename string) (err error) {
	var cfg config
	if err := yaml.Unmarshal(configData, &cfg); err != nil {
		return fmt.Errorf("decode config: %w", err)
	}

	tmpl := template.New("base").Funcs(template.FuncMap{
		"ucfirst":  ucfirst,
		"dict":     dict,
		"contains": strings.Contains,
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
	flag.StringVar(&filename, "output", "inline.go", "the filename to write the generated code to")
	if err := generate(filename); err != nil {
		log.Fatal(err)
	}
}
