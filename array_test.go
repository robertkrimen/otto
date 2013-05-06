package otto

import (
	. "./terst"
	"testing"
)

func TestArray(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = ""+[0, 1, 2, 3];
		def = [].toString();
		ghi = [null, 4, "null"].toString();
	`)
	test("abc", "0,1,2,3")
	test("def", "")
	test("ghi", ",4,null")
	test(`new Array(0).length`, "0")
	test(`new Array(11).length`, "11")
	test(`new Array(11, 1).length`, "2")

	test(`
        abc = [0, 1, 2, 3];
        abc.xyzzy = "Nothing happens.";
        delete abc[1];
        var xyzzy = delete abc.xyzzy;
        [ abc, xyzzy, abc.xyzzy ];
    `, "0,,2,3,true,")
}

func TestArray_toString(t *testing.T) {
	Terst(t)

	{
		test := runTest()

		test(`
			Array.prototype.toString = function() {
				return "Nothing happens."
			}
			abc = Array.prototype.toString()
			def = [].toString()
			ghi = [null, 4, "null"].toString()
		`)
		test("abc", "Nothing happens.")
		test("def", "Nothing happens.")
		test("ghi", "Nothing happens.")

	}

	{
		test := runTest()

		test(`
			Array.prototype.join = undefined
			abc = Array.prototype.toString()
			def = [].toString()
			ghi = [null, 4, "null"].toString()
		`)
		test("abc", "[object Array]")
		test("def", "[object Array]")
		test("ghi", "[object Array]")
	}
}

func TestArray_concat(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0, 1, 2]
		def = [-1, -2, -3]
		ghi = abc.concat(def)
		jkl = abc.concat(def, 3, 4, 5)
		mno = def.concat(-4, -5, abc)
	`)
	test("ghi", "0,1,2,-1,-2,-3")
	test("jkl", "0,1,2,-1,-2,-3,3,4,5")
	test("mno", "-1,-2,-3,-4,-5,0,1,2")
}

func TestArray_splice(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0, 1, 2]
		def = abc.splice(1, 2, 3, 4, 5)
		ghi = [].concat(abc)
		jkl = ghi.splice(17, 21, 7, 8, 9)
	`)
	test("abc", "0,3,4,5")
	test("def", "1,2")
	test("ghi", "0,3,4,5,7,8,9")
	test("jkl", "")
}

func TestArray_shift(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0, 1, 2]
		def = abc.shift()
		ghi = [].concat(abc)
		jkl = abc.shift()
		mno = [].concat(abc)
		pqr = abc.shift()
		stu = [].concat(abc)
		vwx = abc.shift()
	`)
	test("abc", "")
	test("def", "0")
	test("ghi", "1,2")
	test("jkl", "1")
	test("mno", "2")
	test("pqr", "2")
	test("stu", "")
	test("vwx", "undefined")
}

func TestArray_push(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0]
		def = abc.push(1)
		ghi = [].concat(abc)
		jkl = abc.push(2,3,4)
	`)
	test("abc", "0,1,2,3,4")
	test("def", "2")
	test("ghi", "0,1")
	test("jkl", "5")
}

func TestArray_pop(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0,1]
		def = abc.pop()
		ghi = [].concat(abc)
		jkl = abc.pop()
		mno = [].concat(abc)
		pqr = abc.pop()
	`)
	test("abc", "")
	test("def", "1")
	test("ghi", "0")
	test("jkl", "0")
	test("mno", "")
	test("pqr", "undefined")
}

func TestArray_slice(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0,1,2,3]
		def = abc.slice()
		ghi = abc.slice(1)
		jkl = abc.slice(3,-1)
		mno = abc.slice(2,-1)
		pqr = abc.slice(-1, -10)
	`)
	test("def", "0,1,2,3")
	test("ghi", "1,2,3")
	test("jkl", "")
	test("mno", "2")
	test("pqr", "")
	test(`abc.slice(2, 2)`, "")
	test(`abc.slice(2, 3)`, "2")

	test(`
		abc = { 0: 0, 1: 1, 2: 2, 3: 3 }
		abc.length = 4
		def = Array.prototype.slice.call(abc)
		ghi = Array.prototype.slice.call(abc,1)
		jkl = Array.prototype.slice.call(abc,3,-1)
		mno = Array.prototype.slice.call(abc,2,-1)
		pqr = Array.prototype.slice.call(abc,-1,-10)
	`)
	// Array.prototype.slice is generic
	test("def", "0,1,2,3")
	test("ghi", "1,2,3")
	test("jkl", "")
	test("mno", "2")
	test("pqr", "")
}

func TestArray_sliceArguments(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		(function(){
			return Array.prototype.slice.call(arguments, 1)
		})({}, 1, 2, 3)
	`, "1,2,3")

}

func TestArray_unshift(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = []
		def = abc.unshift(0)
		ghi = [].concat(abc)
		jkl = abc.unshift(1,2,3,4)
	`)
	test("abc", "1,2,3,4,0")
	test("def", "1")
	test("ghi", "0")
	test("jkl", "5")
}

func TestArray_reverse(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0,1,2,3].reverse()
		def = [0,1,2].reverse()
	`)
	test("abc", "3,2,1,0")
	test("def", "2,1,0")
}

func TestArray_sort(t *testing.T) {
	Terst(t)

	test := runTest()

	test(`
		abc = [0,1,2,3].sort()
		def = [3,2,1,0].sort()
		ghi = [].sort()
		jkl = [0].sort()
		mno = [1,0].sort()
		pqr = [1,5,-10, 100, 8, 72, 401, 0.05].sort()
		stu = [1,5,-10, 100, 8, 72, 401, 0.05].sort(function(x, y){
			return x == y ? 0 : x < y ? -1 : 1
		})
	`)
	test("abc", "0,1,2,3")
	test("def", "0,1,2,3")
	test("ghi", "")
	test("jkl", "0")
	test("mno", "0,1")
	test("pqr", "-10,0.05,1,100,401,5,72,8")
	test("stu", "-10,0.05,1,5,8,72,100,401")
}

func TestArray_isArray(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`
        [ Array.isArray(), Array.isArray([]), Array.isArray({}) ];
    `, "false,true,false")
}

func TestArray_indexOf(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`['a', 'b', 'c', 'b'].indexOf('b')`, "1")
	test(`['a', 'b', 'c', 'b'].indexOf('b', 2)`, "3")
	test(`['a', 'b', 'c', 'b'].indexOf('b', -2)`, "3")
	test(`
		Object.prototype.indexOf = Array.prototype.indexOf
		var obj = {0: 'a', 1: 'b', 2: 'c', length: 3}
		// [typeof obj.indexOf === 'function', obj.indexOf('c')]
		obj.indexOf('c')
	`, "2")
}

func TestArray_lastIndexOf(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`['a', 'b', 'c', 'b'].lastIndexOf('b')`, "3")
	test(`['a', 'b', 'c', 'b'].lastIndexOf('b', 2)`, "1")
	test(`['a', 'b', 'c', 'b'].lastIndexOf('b', -2)`, "1")
	test(`
		Object.prototype.lastIndexOf = Array.prototype.lastIndexOf
		var obj = {0: 'a', 1: 'b', 2: 'c', 3: 'b', length: 4}
		// [typeof obj.lastIndexOf === 'function', obj.lastIndexOf('b')]
		obj.lastIndexOf('b')
	`, "3")
}

func TestArray_every(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: [].every("abc")`, "TypeError")
	test(`[].every(function(v,k,o) { return false })`, "true")
	test(`[1,2,3].every(function(v,k,o) { return false })`, "false")
	test(`[1,2,3].every(function(v,k,o) { return true })`, "true")
}

func TestArray_some(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: [].some("abc")`, "TypeError")
	test(`[].some(function(v,k,o) { return true })`, "false")
	test(`[1,2,3].some(function(v,k,o) { return false })`, "false")
	test(`[1,2,3].some(function(v,k,o) { return true })`, "true")
}

func TestArray_forEach(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: [].forEach("abc")`, "TypeError")
	test(`var a = 0;[].forEach(function(v,k,o) { a += v });a`, "0")
	test(`[1,2,3].forEach(function(v,k,o) { a += v });a`, "6")
}

func TestArray_map(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: [].map("abc")`, "TypeError")
	test(`[].map(function(v,k,o) { return 1 }).length`, "0")
	test(`[1,2,3].map(function(v,k,o) { return v * v })`, "1,4,9")
	test(`[1,2,3].map(function(v,k,o) { return 1 })`, "1,1,1")
}


func TestArray_filter(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: [].filter("abc")`, "TypeError")
	test(`[].filter(function(v,k,o) { return 1 }).length`, "0")
	test(`[1,2,3].filter(function(v,k,o) { return false }).length`, "0")
	test(`[1,2,3].filter(function(v,k,o) { return true })`, "1,2,3")
}

func TestArray_reduce(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`raise: [].reduce("abc")`, "TypeError")
	test(`raise: [].reduce(function(a,v,k,o) {})`, "TypeError")
	test(`[].reduce(function(a,v,k,o) {}, 0)`, "0")
	test(`['a','b','c'].reduce(function(a,v,k,o) { return a+', '+v })`, "a, b, c")
	test(`[1,2,3].reduce(function(a,v,k,o) { return a + v }, 4)`, "10")
	test(`[1,2,3].reduce(function(a,v,k,o) { return a + v })`, "6")
}
