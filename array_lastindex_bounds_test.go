package otto

import "testing"

func TestArrayLastIndexOfBoundsAndCoercion(t *testing.T) {
	cases := []struct{ name, source, want string }{
		{"at length", `Array.prototype.lastIndexOf.call({0:'x',1:'x',length:1},'x',1)`, "0"},
		{"at length absent", `Array.prototype.lastIndexOf.call({1:'x',length:1},'x',1)`, "-1"},
		{"inherited past length", `var p={1:'x'};var o=Object.create(p);o.length=1;Array.prototype.lastIndexOf.call(o,'x',1)`, "-1"},
		{"getter past length", `var n=0,o={0:'x',length:1};Object.defineProperty(o,'1',{get:function(){n++;return 'x'}});[Array.prototype.lastIndexOf.call(o,'x',1),n].join(',')`, "0,0"},
		{"captured length", `var o={0:'x',length:1};var f={valueOf:function(){o[1]='x';o.length=2;return 1}};Array.prototype.lastIndexOf.call(o,'x',f)`, "0"},
		{"zero ignores coercion", `var n=0;var f={valueOf:function(){n++;throw 'bad'}};var v=Array.prototype.lastIndexOf.call({length:0},'x',f);[v,n].join(',')`, "-1,0"},
		{"zero ignores properties", `Array.prototype.lastIndexOf.call({0:'x',1:'x',length:0},'x',1)`, "-1"},
		{"zero ignores inherited", `var o=Object.create({0:'x'});o.length=0;Array.prototype.lastIndexOf.call(o,'x',0)`, "-1"},
		{"zero getter", `var n=0,o={length:0};Object.defineProperty(o,'0',{get:function(){n++;return 'x'}});[Array.prototype.lastIndexOf.call(o,'x',0),n].join(',')`, "-1,0"},
		{"zero uint32 length", `var n=0;var f={valueOf:function(){n++;return 0}};[Array.prototype.lastIndexOf.call({0:'x',length:4294967296},'x',f),n].join(',')`, "-1,0"},
		{"empty array", `[].lastIndexOf('x',1)`, "-1"},
		{"ordinary backwards", `['a','b','a'].lastIndexOf('a',1)`, "0"},
		{"default last", `['a','b','a'].lastIndexOf('a')`, "2"},
		{"explicit undefined", `['a','b','a'].lastIndexOf('a',undefined)`, "0"},
		{"positive infinity", `['a','b','a'].lastIndexOf('a',Infinity)`, "2"},
		{"negative infinity", `['a','b','a'].lastIndexOf('a',-Infinity)`, "-1"},
		{"negative offset", `['a','b','a'].lastIndexOf('a',-2)`, "0"},
		{"negative too far", `['a','b','a'].lastIndexOf('a',-4)`, "-1"},
		{"fraction", `['a','b','a'].lastIndexOf('a',1.9)`, "0"},
		{"nan", `['a','b','a'].lastIndexOf('a',NaN)`, "0"},
		{"inherited within length", `var o=Object.create({1:'x'});o.length=3;Array.prototype.lastIndexOf.call(o,'x')`, "1"},
		{"sparse skips undefined", `new Array(3).lastIndexOf(undefined)`, "-1"},
		{"present undefined", `[0,undefined].lastIndexOf(undefined)`, "1"},
		{"coerces once", `var n=0,f={valueOf:function(){n++;return 2}};[['x','y','x'].lastIndexOf('x',f),n].join(',')`, "2,1"},
		{"length before from", `var log=[],o={0:'x'};Object.defineProperty(o,'length',{get:function(){log.push('length');return 1}});var f={valueOf:function(){log.push('from');return 0}};[Array.prototype.lastIndexOf.call(o,'x',f),log.join(':')].join(',')`, "0,length:from"},
		{"nonempty throws", `var f={valueOf:function(){throw 'sentinel'}};try {Array.prototype.lastIndexOf.call({length:1},'x',f)}catch(e){e}`, "sentinel"},
		{"null receiver first", `var n=0,f={valueOf:function(){n++;return 0}};try{Array.prototype.lastIndexOf.call(null,'x',f)}catch(e){[e instanceof TypeError,n].join(',')}`, "true,0"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			v, e := New().Run(c.source)
			if e != nil {
				t.Fatalf("unexpected error: %v", e)
			}
			if got := v.String(); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
