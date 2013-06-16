package otto

import (
	. "./terst"
	"fmt"
	"testing"
	Time "time"
)

func TestDate(t *testing.T) {
	Terst(t)

	// Passing or failing should not be dependent on what time zone we're in
	defer mockTimeLocal(Time.UTC)()

	test := runTest()

	time := Time.Unix(1348616313, 47*1000*1000).Local()
	check := func(run string, value int) {
		test(run, fmt.Sprintf("%d", value))
	}

	test(`Date`, "[function]")
	test(`new Date(0).toUTCString()`, "Thu, 01 Jan 1970 00:00:00 UTC")
	test(`new Date(0).toGMTString()`, "Thu, 01 Jan 1970 00:00:00 GMT")
	if false {
		// TODO toLocale{Date,Time}String
		test(`new Date(0).toLocaleString()`, "")
		test(`new Date(0).toLocaleDateString()`, "")
		test(`new Date(0).toLocaleTimeString()`, "")
	}
	test(`new Date(1348616313).getTime()`, "1348616313")
	test(`new Date(1348616313).toUTCString()`, "Fri, 16 Jan 1970 14:36:56 UTC")
	test(`abc = new Date(1348616313047); abc.toUTCString()`, "Tue, 25 Sep 2012 23:38:33 UTC")
	check(`abc.getYear()`, time.Year()-1900)
	check(`abc.getFullYear()`, time.Year())
	check(`abc.getUTCFullYear()`, 2012)
	check(`abc.getMonth()`, int(time.Month())-1) // Remember, the JavaScript month is 0-based
	check(`abc.getUTCMonth()`, 8)
	check(`abc.getDate()`, time.Day())
	check(`abc.getUTCDate()`, 25)
	check(`abc.getDay()`, int(time.Weekday()))
	check(`abc.getUTCDay()`, 2)
	check(`abc.getHours()`, time.Hour())
	check(`abc.getUTCHours()`, 23)
	check(`abc.getMinutes()`, time.Minute())
	check(`abc.getUTCMinutes()`, 38)
	check(`abc.getSeconds()`, time.Second())
	check(`abc.getUTCSeconds()`, 33)
	check(`abc.getMilliseconds()`, time.Nanosecond()/(1000*1000)) // In honor of the 47%
	check(`abc.getUTCMilliseconds()`, 47)
	_, offset := time.Zone()
	check(`abc.getTimezoneOffset()`, offset/-60)

	test(`new Date("Xyzzy").getTime()`, "NaN")

	test(`abc.setFullYear(2011); abc.toUTCString()`, "Sun, 25 Sep 2011 23:38:33 UTC")
	test(`new Date(12564504e5).toUTCString()`, "Sun, 25 Oct 2009 06:00:00 UTC")
	test(`new Date(2009, 9, 25).toUTCString()`, "Sun, 25 Oct 2009 00:00:00 UTC")
	test(`+(new Date(2009, 9, 25))`, "1256428800000")

	format := "Mon, 2 Jan 2006 15:04:05 MST"

	tme := Time.Unix(1256450400, 0)
	time = Time.Date(tme.Year(), tme.Month(), tme.Day(), tme.Hour(), tme.Minute(), tme.Second(), tme.Nanosecond(), tme.Location()).UTC()

	time = Time.Date(tme.Year(), tme.Month(), tme.Day(), tme.Hour(), tme.Minute(), tme.Second(), 2001*1000*1000, tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setMilliseconds(2001); abc.toUTCString()`, time.Format(format))

	time = Time.Date(tme.Year(), tme.Month(), tme.Day(), tme.Hour(), tme.Minute(), 61, tme.Nanosecond(), tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setSeconds("61"); abc.toUTCString()`, time.Format(format))

	time = Time.Date(tme.Year(), tme.Month(), tme.Day(), tme.Hour(), 61, tme.Second(), tme.Nanosecond(), tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setMinutes("61"); abc.toUTCString()`, time.Format(format))

	time = Time.Date(tme.Year(), tme.Month(), tme.Day(), 5, tme.Minute(), tme.Second(), tme.Nanosecond(), tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setHours("5"); abc.toUTCString()`, time.Format(format))

	time = Time.Date(tme.Year(), tme.Month(), 26, tme.Hour(), tme.Minute(), tme.Second(), tme.Nanosecond(), tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setDate("26"); abc.toUTCString()`, time.Format(format))

	time = Time.Date(tme.Year(), 10, tme.Day(), tme.Hour(), tme.Minute(), tme.Second(), tme.Nanosecond(), tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setMonth(9); abc.toUTCString()`, time.Format(format))
	test(`abc = new Date(12564504e5); abc.setMonth("09"); abc.toUTCString()`, time.Format(format))

	time = Time.Date(tme.Year(), 11, tme.Day(), tme.Hour(), tme.Minute(), tme.Second(), tme.Nanosecond(), tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setMonth("10"); abc.toUTCString()`, time.Format(format))

	time = Time.Date(2010, tme.Month(), tme.Day(), tme.Hour(), tme.Minute(), tme.Second(), tme.Nanosecond(), tme.Location()).UTC()
	test(`abc = new Date(12564504e5); abc.setFullYear(2010); abc.toUTCString()`, time.Format(format))

	test(`new Date("2001-01-01T10:01:02.000").getTime()`, "978343262000")

	// Date()
	test(`typeof Date()`, "string")
	test(`typeof Date(2006, 1, 2)`, "string")

	test(`
        abc = Object.getOwnPropertyDescriptor(Date, "parse");
        [ abc.value === Date.parse, abc.writable, abc.enumerable, abc.configurable ];
    `, "true,true,false,true")

	test(`
        abc = Object.getOwnPropertyDescriptor(Date.prototype, "toTimeString");
        [ abc.value === Date.prototype.toTimeString, abc.writable, abc.enumerable, abc.configurable ];
    `, "true,true,false,true")

	test(`
        var abc = Object.getOwnPropertyDescriptor(Date, "prototype");
        [   [ typeof Date.prototype ],
            [ abc.writable, abc.enumerable, abc.configurable ] ];
    `, "object,false,false,false")
}

func TestDate_parse(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`Date.parse("2001-01-01T10:01:02.000")`, "978343262000")
	test(`Date.parse("2006-01-02T15:04:05.000")`, "1136214245000")
	test(`Date.parse("2006")`, "1136073600000")
	test(`Date.parse("1970-01-16T14:36:56+00:00")`, "1348616000")
	test(`Date.parse("1970-01-16T14:36:56.313+00:00")`, "1348616313")
	test(`Date.parse("1970-01-16T14:36:56.000")`, "1348616000")
}

func TestDate_UTC(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`Date.UTC(2009, 9, 25)`, "1256428800000")
}

func TestDate_toISOString(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`new Date(0).toISOString()`, "1970-01-01T00:00:00.000Z")
}

func TestDate_toJSON(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`new Date(0).toJSON()`, "1970-01-01T00:00:00.000Z")
}

func TestDate_setYear(t *testing.T) {
	Terst(t)

	// Passing or failing should not be dependent on what time zone we're in
	defer mockTimeLocal(Time.UTC)()

	test := runTest()
	test(`new Date(12564504e5).setYear(96)`, "846223200000")
	test(`new Date(12564504e5).setYear(1996)`, "846223200000")
	test(`new Date(12564504e5).setYear(2000)`, "972453600000")
}

func mockTimeLocal(location *Time.Location) func() {
	local := Time.Local
	Time.Local = location
	return func() {
		Time.Local = local
	}
}

func TestDateDefaultValue(t *testing.T) {
	Terst(t)

	test := runTest()
	test(`
        var date = new Date();
        date + 0 === date.toString() + "0";
    `, "true")
}

func TestDate_April1978(t *testing.T) {
	Terst(t)

	// Passing or failing should not be dependent on what time zone we're in
	defer mockTimeLocal(Time.UTC)()

	test := runTest()
	test(`
        var abc = new Date(1978,3);
        [ abc.getYear(), abc.getMonth(), abc.valueOf() ];
    `, "78,3,260236800000")
}
