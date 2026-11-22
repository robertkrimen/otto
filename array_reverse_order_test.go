package otto

import (
	"fmt"
	"testing"
)

func TestArrayReverseSparseWriteBeforeDelete(t *testing.T) {
	for _, receiver := range []string{"array", "object"} {
		for _, testCase := range []struct {
			name  string
			setup string
			want  string
		}{
			{"write_rejected", `a[1] = "upper"; Object.preventExtensions(a);`, `TypeError|false:undefined|true:upper`},
			{"delete_rejected", `Object.defineProperty(a, "1", {value: "upper", configurable: false});`, `TypeError|true:upper|true:upper`},
			{"getter_delete_rejected", `Object.defineProperty(a, "1", {get: function () { return "upper"; }, configurable: false});`, `TypeError|true:upper|true:upper`},
			{"undefined_delete_rejected", `Object.defineProperty(a, "1", {value: undefined, configurable: false});`, `TypeError|true:undefined|true:undefined`},
			{"upper_only", `a[1] = "upper";`, `none|true:upper|false:undefined`},
			{"lower_only", `a[0] = "lower";`, `none|false:undefined|true:lower`},
			{"lower_write_rejected", `a[0] = "lower"; Object.preventExtensions(a);`, `TypeError|false:undefined|false:undefined`},
			{"both_present", `a[0] = "lower"; a[1] = "upper"; Object.preventExtensions(a);`, `none|true:upper|true:lower`},
			{"both_holes", `Object.preventExtensions(a);`, `none|false:undefined|false:undefined`},
			{"sealed", `a[1] = "upper"; Object.seal(a);`, `TypeError|false:undefined|true:upper`},
			{"frozen", `a[1] = "upper"; Object.freeze(a);`, `TypeError|false:undefined|true:upper`},
		} {
			t.Run(receiver+"/"+testCase.name, func(t *testing.T) {
				initial := `({length: 2})`
				if receiver == "array" {
					initial = `new Array(2)`
				}
				source := fmt.Sprintf(`(function () {
					var a = %s;
					%s
					var thrown = "none";
					try { Array.prototype.reverse.call(a); }
					catch (e) { thrown = e.name; }
					return [thrown,
						Object.prototype.hasOwnProperty.call(a, "0") + ":" + a[0],
						Object.prototype.hasOwnProperty.call(a, "1") + ":" + a[1]].join("|");
				}())`, initial, testCase.setup)
				value, err := New().Run(source)
				if err != nil {
					t.Fatal(err)
				}
				if got := value.String(); got != testCase.want {
					t.Fatalf("got %q, want %q", got, testCase.want)
				}
			})
		}
	}
}

func TestArrayReversePreservesEarlierPairsWhenDeletionFails(t *testing.T) {
	value, err := New().Run(`(function () {
		var a = ["first", , "inner", "last"];
		Object.defineProperty(a, "2", {configurable: false});
		var thrown = "none";
		try { a.reverse(); } catch (e) { thrown = e.name; }
		return [thrown, a.join(","), Object.prototype.hasOwnProperty.call(a, "1")].join("|");
	}())`)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := value.String(), "TypeError|last,inner,inner,first|true"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
