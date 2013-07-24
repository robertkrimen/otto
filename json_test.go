package otto

import (
	. "./terst"
	"testing"
	"time"
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
	test(`JSON.parse('{"a":1,"b":2}').a`, "1")
	test(`JSON.parse('{"a":{"x":100,"y":110},"b":[10,20,30],"c":"zazazaza"}').b`, "10,20,30")
	test(`JSON.parse("[1,2,3]", function(k, v) { return undefined })`, "undefined")

	test(`raise: JSON.parse("")`, "SyntaxError: Unexpected end of input")
	test(`raise: JSON.parse("[1,2,3")`, "SyntaxError: Unexpected end of input")
	test(`raise: JSON.parse("[1,2,;x=10")`, "SyntaxError: Unexpected token ;")
	test(`raise: JSON.parse("[1,2,function(){}]")`, "SyntaxError: Unexpected token function")
}

func TestJSON_stringify(t *testing.T) {
	Terst(t)
	test := runTest()

	defer mockTimeLocal(time.UTC)()

	test(`JSON.stringify(undefined)`, "undefined")
	test(`JSON.stringify(1)`, "1")
	test(`JSON.stringify([])`, "[]")
	test(`JSON.stringify([1,2,3])`, "[1,2,3]")
	test(`JSON.stringify([true, false, null])`, "[true,false,null]")
	test(`JSON.stringify({a:{x:100,y:110},b:[10,20,30],c:"zazazaza"})`, `{"a":{"x":100,"y":110},"b":[10,20,30],"c":"zazazaza"}`)

	test(`JSON.stringify(['e', {pluribus: 'unum'}], null, '\t')`, "[\n\t\"e\",\n\t{\n\t\t\"pluribus\": \"unum\"\n\t}\n]")
	test(`JSON.stringify(new Date(0))`, `"1970-01-01T00:00:00.000Z"`)
	test(`JSON.stringify([new Date(0)], function(k,v){
		return this[k] instanceof Date ? 'Date(' + this[k] + ')' : v
	})`, `["Date(Thu, 01 Jan 1970 00:00:00 UTC)"]`)
	test(`JSON.stringify({a:1,b:2,c:3}, ['a','b'])`, `{"a":1,"b":2}`)
}
