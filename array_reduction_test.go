package otto

import (
	"strings"
	"testing"
)

func TestArrayReductionSparseAccumulator(t *testing.T) {
	for _, method := range []string{"reduce", "reduceRight"} {
		t.Run(method, func(t *testing.T) {
			cases := []struct{ name, script, want string }{
				{"holes", `var a = new Array(3); try { a.METHOD(function() { throw 'called'; }); 'no error'; } catch (e) { e instanceof TypeError ? 'TypeError' : String(e); }`, "TypeError"},
				{"generic holes", `try { Array.prototype.METHOD.call({length: 3}, function() {}); 'no error'; } catch (e) { e instanceof TypeError ? 'TypeError' : String(e); }`, "TypeError"},
				{"deleted elements", `var a = [1, 2]; delete a[0]; delete a[1]; try { a.METHOD(function() {}); 'no error'; } catch (e) { e instanceof TypeError ? 'TypeError' : String(e); }`, "TypeError"},
				{"explicit undefined", `var calls = 0; var r = new Array(3).METHOD(function() { calls++; }, undefined); String(r) + ':' + calls;`, "undefined:0"},
				{"explicit seed", `var calls = 0; var r = new Array(3).METHOD(function() { calls++; }, 7); r + ':' + calls;`, "7:0"},
				{"undefined element", `var calls = 0; var r = [undefined].METHOD(function() { calls++; }); String(r) + ':' + calls;`, "undefined:0"},
				{"middle element", `var a = new Array(5); a[2] = 9; a.METHOD(function() { throw 'called'; });`, "9"},
				{"inherited element", `var p = {2: 11}; var a = Object.create(p); a.length = 5; Array.prototype.METHOD.call(a, function() { throw 'called'; });`, "11"},
				{"inherited undefined", `var a = Object.create({2: undefined}); a.length = 5; String(Array.prototype.METHOD.call(a, function() { throw 'called'; }));`, "undefined"},
				{"inherited getter", `var n = 0; var p = {}; Object.defineProperty(p, '2', {get: function() { n++; return 13; }}); var a = Object.create(p); a.length = 5; Array.prototype.METHOD.call(a, function() { throw 'called'; }) + ':' + n;`, "13:1"},
				{"throwing seed getter", `var a = {length: 3}; Object.defineProperty(a, '1', {get: function() { throw 'seed'; }}); try { Array.prototype.METHOD.call(a, function() {}); } catch (e) { String(e); }`, "seed"},
				{"undefined seed with another element", `var a = new Array(3); a[0] = undefined; a[2] = undefined; var n = 0; var r = a.METHOD(function(acc, v) { n++; return typeof acc + ':' + typeof v; }); r + ':' + n;`, "undefined:undefined:1"},
				{"sparse callback object", `var a = new Array(5); a[1] = 4; a[3] = 6; var n = 0; a.METHOD(function(acc, v, i, o) { if (o !== a || o[i] !== v || typeof i !== 'number') throw 'wrong callback'; n++; return acc + v; }, 10) + ':' + n;`, "20:2"},
				{"numeric callback index", `var result = ''; [10, 20, 30].METHOD(function(a, v, i) { result += typeof i; return a; }, 0); result;`, "numbernumbernumber"},
			}
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					script := strings.ReplaceAll(tc.script, "METHOD", method)
					v, err := New().Run(script)
					if err != nil {
						t.Fatal(err)
					}
					if got := v.String(); got != tc.want {
						t.Fatalf("got %q, want %q", got, tc.want)
					}
				})
			}
		})
	}
}
