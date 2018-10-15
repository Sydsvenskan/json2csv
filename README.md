# JSON to CSV conversion

[![GoDoc](https://godoc.org/github.com/Sydsvenskan/json2csv?status.svg)](https://godoc.org/github.com/Sydsvenskan/json2csv) [![Goreport](https://goreportcard.com/badge/github.com/Sydsvenskan/json2csv)](https://goreportcard.com/report/github.com/Sydsvenskan/json2csv) [![Build Status](https://travis-ci.org/Sydsvenskan/json2csv.svg?branch=master)](https://travis-ci.org/Sydsvenskan/json2csv) [![codecov](https://codecov.io/gh/Sydsvenskan/json2csv/branch/master/graph/badge.svg)](https://codecov.io/gh/Sydsvenskan/json2csv)

Converts arbitrary JSON data to CSV. Does it make sense? Not always, but it can be useful for the right use-case :)

Usage:

```go
func main() {
	json2csv.Convert(strings.NewReader(`[
{
  "foo": "bar",
  "baz": [{"a":"b"}, 1, 1.3]
},
{
  "foo": "fi fum",
  "baz": [{"a":"c"}, 3],
  "extra": "fine"
},
"it's JSON, what do you expect?"
]`), os.Stdout)
	// Output:
	// baz.0.a,baz.1,baz.2,extra,foo,text
	// b,1,1.3,,bar,
	// c,3,,fine,fi fum,
	// ,,,,,"it's JSON, what do you expect?"
}
```
