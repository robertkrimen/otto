package main

import (
	"fmt"
	"unicode"
)

// ucfirst converts the first rune of val to uppercase an returns the result.
func ucfirst(val string) string {
	r := []rune(val)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}

// dict is a template helper returns a map created from alternating values.
// Values must be passed as key, value pairs with key being a string.
// It can be used to pass the combination of multiple values to a template.
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("map requires parameters which are multiple of 2 got %d", len(values))
	}

	m := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("map keys must be strings got %T", values[i])
		}
		m[key] = values[i+1]
	}

	return m, nil
}
