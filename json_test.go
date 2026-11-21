package otto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkJSON_parse(b *testing.B) {
	vm := New()
	for i := 0; i < b.N; i++ {
		_, err := vm.Run(`JSON.parse("1")`)
		require.NoError(b, err)
		_, err = vm.Run(`JSON.parse("[1,2,3]")`)
		require.NoError(b, err)
		_, err = vm.Run(`JSON.parse('{"a":{"x":100,"y":110},"b":[10,20,30],"c":"zazazaza"}')`)
		require.NoError(b, err)
		_, err = vm.Run(`JSON.parse("[1,2,3]", function(k, v) { return undefined })`)
		require.NoError(b, err)
	}
}

func TestJSON_parse(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		test(`
            JSON.parse("1");
        `, 1)

		test(`
            JSON.parse("null");
        `, "null") // TODO Can we make this nil?

		test(`
            var abc = JSON.parse('"a\uFFFFbc"');
            [ abc[0], abc[2], abc[3], abc.length ];
        `, "a,b,c,4")

		test(`
            JSON.parse("[1, 2, 3]");
        `, "1,2,3")

		test(`
            JSON.parse('{ "abc": 1, "def":2 }').abc;
        `, 1)

		test(`
            JSON.parse('{ "abc": { "x": 100, "y": 110 }, "def": [ 10, 20 ,30 ], "ghi": "zazazaza" }').def;
        `, "10,20,30")

		test(`raise:
            JSON.parse("12\t\r\n 34");
        `, "SyntaxError: invalid character '3' after top-level value")

		test(`
            JSON.parse("[1, 2, 3]", function() { return undefined });
        `, "undefined")

		test(`raise:
            JSON.parse("");
        `, "SyntaxError: unexpected end of JSON input")

		test(`raise:
            JSON.parse("[1, 2, 3");
        `, "SyntaxError: unexpected end of JSON input")

		test(`raise:
            JSON.parse("[1, 2, ; abc=10");
        `, "SyntaxError: invalid character ';' looking for beginning of value")

		test(`raise:
            JSON.parse("[1, 2, function(){}]");
        `, "SyntaxError: invalid character 'u' in literal false (expecting 'a')")
	})
}

func TestJSON_stringify(t *testing.T) {
	tt(t, func() {
		test, _ := test()

		defer mockUTC()()

		test(`
            JSON.stringify(function(){});
        `, "undefined")

		test(`
            JSON.stringify(new Boolean(false));
        `, "false")

		test(`
            JSON.stringify({a1: {b1: [1,2,3,4], b2: {c1: 1, c2: 2}}, a2: 'a2'}, null, -5);
        `, `{"a1":{"b1":[1,2,3,4],"b2":{"c1":1,"c2":2}},"a2":"a2"}`)

		test(`
            JSON.stringify(undefined);
        `, "undefined")

		test(`
            JSON.stringify(1);
        `, "1")

		test(`
            JSON.stringify("abc def");
        `, "\"abc def\"")

		test(`
            JSON.stringify(3.14159);
        `, "3.14159")

		test(`
            JSON.stringify([]);
        `, "[]")

		test(`
            JSON.stringify([1, 2, 3]);
        `, "[1,2,3]")

		test(`
            JSON.stringify([true, false, null]);
        `, "[true,false,null]")

		test(`
            JSON.stringify({
                abc: { x: 100, y: 110 },
                def: [ 10, 20, 30 ],
                ghi: "zazazaza"
            });
        `, `{"abc":{"x":100,"y":110},"def":[10,20,30],"ghi":"zazazaza"}`)

		test(`
            JSON.stringify([
                'e',
                {pluribus: 'unum'}
            ], null, '\t');
        `, "[\n\t\"e\",\n\t{\n\t\t\"pluribus\": \"unum\"\n\t}\n]")

		test(`
            JSON.stringify(new Date(0));
        `, `"1970-01-01T00:00:00.000Z"`)

		test(`
            JSON.stringify([ new Date(0) ], function(key, value){
                return this[key] instanceof Date ? 'Date(' + this[key] + ')' : value
            });
        `, `["Date(Thu, 01 Jan 1970 00:00:00 UTC)"]`)

		test(`
            JSON.stringify({
                abc: 1,
                def: 2,
                ghi: 3
            }, ['abc','def']);
        `, `{"abc":1,"def":2}`)

		test(`raise:
            var abc = {
                def: null
            };
            abc.def = abc;
            JSON.stringify(abc)
        `, "TypeError: Converting circular structure to JSON")

		test(`raise:
            var abc= [ null ];
            abc[0] = abc;
            JSON.stringify(abc);
        `, "TypeError: Converting circular structure to JSON")

		test(`raise:
            var abc = {
                def: {}
            };
            abc.def.ghi = abc;
            JSON.stringify(abc)
        `, "TypeError: Converting circular structure to JSON")

		test(`
            var ghi = { "pi": 3.14159 };
            var abc = {
                def: {}
            };
            abc.ghi = ghi;
            abc.def.ghi = ghi;
            JSON.stringify(abc);
        `, `{"def":{"ghi":{"pi":3.14159}},"ghi":{"pi":3.14159}}`)
	})
}

func TestJSON_stringifyReplacerArrayCompaction(t *testing.T) {
	cases := []struct {
		name, script, want string
	}{
		{"leading_ignored", `JSON.stringify({keep: 1}, [null, "keep"])`, `{"keep":1}`},
		{"middle_ignored", `JSON.stringify({a: 1, b: 2}, ["a", false, "b"])`, `{"a":1,"b":2}`},
		{"duplicate", `JSON.stringify({a: 1, b: 2}, ["a", "a", "b"])`, `{"a":1,"b":2}`},
		{"sparse", `JSON.stringify({keep: 1}, [,,"keep"])`, `{"keep":1}`},
		{"boxed_duplicate", `JSON.stringify({1: 1, b: 2}, [new Number(1), "1", new String("b")])`, `{"1":1,"b":2}`},
		{"ignored_object", `JSON.stringify({keep: 1}, [{toString: function() { throw "unused"; }}, "keep"])`, `{"keep":1}`},
		{"empty_key_not_selected", `JSON.stringify({"": "private", keep: 1}, [null, "keep"])`, `{"keep":1}`},
		{"nested", `JSON.stringify({child: {keep: 1, drop: 2}}, [false, "child", "keep"])`, `{"child":{"keep":1}}`},
		{"getter_visited", `(function() { var calls = 0, obj = {}; Object.defineProperty(obj, "keep", {get: function() { calls++; return 7; }}); return JSON.stringify(obj, [null, "keep"]) + ":" + calls; })()`, `{"keep":7}:1`},
		{"all_valid_control", `JSON.stringify({a: 1, b: 2, drop: 3}, ["a", "b"])`, `{"a":1,"b":2}`},
		{"all_ignored_control", `JSON.stringify({a: 1}, [null, false, {}])`, `{}`},
		{"empty_control", `JSON.stringify({a: 1}, [])`, `{}`},
		{"empty_key_selected_control", `JSON.stringify({"": 1, a: 2}, [""])`, `{"":1}`},
		{"array_control", `JSON.stringify([1, {keep: 2}], ["keep"])`, `[1,{"keep":2}]`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			value, err := New().Run(tc.script)
			if err != nil {
				t.Fatal(err)
			}
			if got := value.String(); got != tc.want {
				t.Fatalf("got %s, want %s", got, tc.want)
			}
		})
	}
}
