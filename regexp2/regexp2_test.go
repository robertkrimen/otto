package regexp2_test

import (
	"regexp"
	"testing"

	re2 "github.com/dlclark/regexp2"
	"github.com/robertkrimen/otto/regexp2"
	"github.com/stretchr/testify/require"
)

func setRE2() func() {
	old := regexp2.DefaultOption
	regexp2.DefaultOption = re2.RE2
	return func() {
		regexp2.DefaultOption = old
	}
}

func TestRegexp_FindStringIndex(t *testing.T) {
	cleanup := setRE2()
	defer cleanup()

	tests := map[string]struct {
		expr string
		data string
		want []int
	}{
		"match": {
			expr: `ab?`,
			data: `tablett ablet`,
			want: []int{1, 3},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			re, err := regexp.Compile(tt.expr)
			require.NoError(t, err)

			res := re.FindStringIndex(tt.data)
			require.Equal(t, tt.want, res)

			re2, err := regexp2.Compile(tt.expr)
			require.NoError(t, err)

			res = re2.FindStringIndex(tt.data)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestRegexp_FindAllStringIndex(t *testing.T) {
	cleanup := setRE2()
	defer cleanup()

	tests := map[string]struct {
		expr string
		n    int
		data string
		want [][]int
	}{
		"all": {
			expr: `ab?`,
			n:    -1,
			data: `tablett ablet`,
			want: [][]int{
				{1, 3},
				{8, 10},
			},
		},
		"one": {
			expr: `ab?`,
			n:    1,
			data: `tablett ablet`,
			want: [][]int{
				{1, 3},
			},
		},
		"zero": {
			expr: `ab?`,
			n:    0,
			data: `tablett ablet`,
		},
		"no-match": {
			expr: `(\#production|\.min\.js)`,
			n:    -1,
			data: "",
			want: nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			re, err := regexp.Compile(tt.expr)
			require.NoError(t, err)

			res := re.FindAllStringIndex(tt.data, tt.n)
			require.Equal(t, tt.want, res)

			re2, err := regexp2.Compile(tt.expr)
			require.NoError(t, err)

			res = re2.FindAllStringIndex(tt.data, tt.n)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestRegexp_FindAllSubmatchIndex(t *testing.T) {
	cleanup := setRE2()
	defer cleanup()

	tests := map[string]struct {
		expr string
		n    int
		data string
		want [][]int
	}{
		"all": {
			expr: `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`,
			n:    -1,
			data: `
				# comment line
				option1: value1
				option2: value2
			`,
			want: [][]int{
				{24, 39, 24, 31, 33, 39},
				{44, 59, 44, 51, 53, 59},
			},
		},
		"one": {
			expr: `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`,
			n:    1,
			data: `
				# comment line
				option1: value1
				option2: value2
			`,
			want: [][]int{
				{24, 39, 24, 31, 33, 39},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			re, err := regexp.Compile(tt.expr)
			require.NoError(t, err)

			res := re.FindAllSubmatchIndex([]byte(tt.data), tt.n)
			require.Equal(t, tt.want, res)

			re2, err := regexp2.Compile(tt.expr)
			require.NoError(t, err)

			res = re2.FindAllSubmatchIndex([]byte(tt.data), tt.n)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestRegexp_FindStringSubmatchIndex(t *testing.T) {
	cleanup := setRE2()
	defer cleanup()

	data := `
		# comment line
		option1: value1
		option2: value2
	`
	tests := map[string]struct {
		expr string
		data string
		want []int
	}{
		"match": {
			expr: `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`,
			data: data,
			want: []int{20, 35, 20, 27, 29, 35},
		},
		"no-match": {
			expr: `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`,
			data: `# comment line`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			re, err := regexp.Compile(tt.expr)
			require.NoError(t, err)

			res := re.FindStringSubmatchIndex(tt.data)
			require.Equal(t, tt.want, res)

			re2, err := regexp2.Compile(tt.expr)
			require.NoError(t, err)

			res = re2.FindStringSubmatchIndex(tt.data)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestRegexp_FindAllStringSubmatchIndex(t *testing.T) {
	cleanup := setRE2()
	defer cleanup()

	data := `
		# comment line
		option1: value1
		option2: value2
	`
	tests := map[string]struct {
		expr string
		n    int
		want [][]int
	}{
		"all": {
			expr: `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`,
			n:    -1,
			want: [][]int{
				{20, 35, 20, 27, 29, 35},
				{38, 53, 38, 45, 47, 53},
			},
		},
		"one": {
			expr: `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`,
			n:    1,
			want: [][]int{
				{20, 35, 20, 27, 29, 35},
			},
		},
		"zero": {
			expr: `(?m)(?P<key>\w+):\s+(?P<value>\w+)$`,
			n:    0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			re, err := regexp.Compile(tt.expr)
			require.NoError(t, err)

			res := re.FindAllStringSubmatchIndex(data, tt.n)
			require.Equal(t, tt.want, res)

			re2, err := regexp2.Compile(tt.expr)
			require.NoError(t, err)

			res = re2.FindAllStringSubmatchIndex(data, tt.n)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestRegExp_Compile(t *testing.T) {
	t.Skip("regexp2 doesn't handle this case yet see: https://github.com/dlclark/regexp2/issues/54")
	_, err := regexp2.Compile(`[a-\s]`)
	require.NoError(t, err)
}
