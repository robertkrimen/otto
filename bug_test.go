package otto

import (
	. "./terst"
	"testing"
)

func Test_262(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise:
        eval("42 = 42;");
    `, "ReferenceError: Invalid left-hand side in assignment")
}

func Test_issue5(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`'abc' === 'def'`, "false")
	test(`'\t' === '\r'`, "false")
}

func Test_issue13(t *testing.T) {
	Terst(t)

	otto, test := runTestWithOtto()
	value, err := otto.ToValue(map[string]interface{}{
		"string": "Xyzzy",
		"number": 42,
		"array":  []string{"def", "ghi"},
	})
	if err != nil {
		FailNow(err)
	}
	fn, err := otto.Object(`
    (function(value){
        return ""+[value.string, value.number, value.array]
    })
    `)
	if err != nil {
		FailNow(err)
	}
	result, err := fn.Value().Call(fn.Value(), value)
	if err != nil {
		FailNow(err)
	}
	Is(result.toString(), "Xyzzy,42,def,ghi")

	anything := struct {
		Abc interface{}
	}{
		Abc: map[string]interface{}{
			"def": []interface{}{
				[]interface{}{
					"a", "b", "c", "", "d", "e",
				},
				map[string]interface{}{
					"jkl": "Nothing happens.",
				},
			},
			"ghi": -1,
		},
	}
	otto.Set("anything", anything)
	test(`
        [
            anything,
            "~",
            anything.Abc,
            "~",
            anything.Abc.def,
            "~",
            anything.Abc.def[1].jkl,
            "~",
            anything.Abc.ghi,
        ];
        `, "[object Object],~,[object Object],~,a,b,c,,d,e,[object Object],~,Nothing happens.,~,-1",
	)

}

func Test_issue16(t *testing.T) {
	Terst(t)

	otto, test := runTestWithOtto()
	test(`
        var def = {
            "abc": ["abc"],
            "xyz": ["xyz"]
        };
        def.abc.concat(def.xyz);
    `, "abc,xyz")

	otto.Set("ghi", []string{"jkl", "mno"})
	test(`
        def.abc.concat(def.xyz).concat(ghi);
    `, "abc,xyz,jkl,mno")

	test(`
        ghi.concat(def.abc.concat(def.xyz));
    `, "jkl,mno,abc,xyz")

	otto.Set("pqr", []interface{}{"jkl", 42, 3.14159, true})
	test(`
        pqr.concat(ghi, def.abc, def, def.xyz);
    `, "jkl,42,3.14159,true,jkl,mno,abc,[object Object],xyz")

	test(`
        pqr.concat(ghi, def.abc, def, def.xyz).length;
    `, "9")
}
