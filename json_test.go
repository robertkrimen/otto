package otto

import (
	. "./terst"
	"testing"
)

func BenchmarkJSON_parse(b *testing.B) {
	o := New(); t := func(s string) { o.Run(s) }
	for i := 0; i < b.N; i++ {
		t(`JSON.parse("1")`)
		t(`JSON.parse("[1,2,3]")`)
		t(`JSON.parse('{"a":{"x":100,"y":110},"b":[10,20,30],"c":"zazazaza"}')`)
		t(`JSON.parse("[1,2,3]", function(k, v) { return undefined })`)
	}
}

func TestJSON_parse(t *testing.T) {
	Terst(t)
	test := runTest()

	test(`JSON.parse("1")`, "1")
	test(`JSON.parse("[1,2,3]")`, "1,2,3")
	test(`JSON.parse('{"a":{"x":100,"y":110},"b":[10,20,30],"c":"zazazaza"}').b`, "10,20,30")
	test(`JSON.parse("[1,2,3]", function(k, v) { return undefined })`, "undefined")

	test(`var a;try{JSON.parse("")}catch(e){a=e instanceof SyntaxError};a`, "true")
	test(`var a;try{JSON.parse("[1,2,3")}catch(e){a=e instanceof SyntaxError};a`, "true")
	test(`var a;try{JSON.parse("[1,2];x=10")}catch(e){a=e instanceof SyntaxError};a`, "true")
	test(`var a;try{JSON.parse("[1,2,function(){}]")}catch(e){a=e instanceof SyntaxError};a`, "true")
}

func TestJSON_stringify(t *testing.T) {
	Terst(t)
	test := runTest()

	test(`JSON.stringify(undefined)`, "undefined")
	test(`JSON.stringify(1)`, "1")
	test(`JSON.stringify([])`, "[]")
	test(`JSON.stringify([1,2,3])`, "[1,2,3]")
	test(`JSON.stringify([true, false, null])`, "[true,false,null]")
	test(`JSON.stringify({a:{x:100,y:110},b:[10,20,30],c:"zazazaza"})`, `{"a":{"x":100,"y":110},"b":[10,20,30],"c":"zazazaza"}`)

	test(`JSON.stringify(['e', {pluribus: 'unum'}], null, '\t')`, "[\n\t\"e\",\n\t{\n\t\t\"pluribus\": \"unum\"\n\t}\n]")
	test(`JSON.stringify([new Boolean(true), new Date(0), new Number(1).toJSON(), new String('abc').toJSON()])`,
		`[true,"1970-01-01T00:00:00.000Z",1,"abc"]`)
	test(`JSON.stringify([new Date(0)], function(k,v){
		return this[k] instanceof Date ? 'Date(' + this[k] + ')' : v
	})`, `["Date(Thu, 01 Jan 1970 01:00:00 CET)"]`)
	test(`JSON.stringify({a:1,b:2,c:3}, ['a','b'])`, `{"a":1,"b":2}`)
}

func TestJSON_toJSON(t *testing.T) {
	Terst(t)
	test := runTest()

	test(`typeof Boolean.prototype.toJSON === 'function'`, "true")
	test(`typeof Date.prototype.toJSON === 'function'`, "true")
	test(`typeof Number.prototype.toJSON === 'function'`, "true")
	test(`typeof String.prototype.toJSON === 'function'`, "true")

	test(`new Boolean(true).toJSON()`, "true")
	test(`new Date(0).toJSON()`, "1970-01-01T00:00:00.000Z")
	test(`new Number(1).toJSON()`, "1")
	test(`new String('abc').toJSON()`, "abc")

	test(`
		var object = {}
		object.toISOString = function() { return 1 }
		object.toJSON = Date.prototype.toJSON
		object.toJSON()
	`, "1")
}
