package json2csv_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Sydsvenskan/json2csv"
)

func ExampleObject() {
	json2csv.Convert(strings.NewReader(`{
  "foo": "bar",
  "baz": [{"a":"b"}, 1, 1.3]
}`), os.Stdout)
	// Output:
	// baz.0.a,baz.1,baz.2,foo
	// b,1,1.3,bar
}

func ExampleArray() {
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

func ExampleNumber() {
	json2csv.Convert(strings.NewReader("42"), os.Stdout)
	// Output:
	// number
	// 42
}

func ExampleArrayInArray() {
	json2csv.Convert(strings.NewReader(`[
["a", "b", "c"],
["one", "two", "three"]
]`), os.Stdout)
	// Output:
	// 0,1,2
	// a,b,c
	// one,two,three
}

func TestBadJSON(t *testing.T) {
	err := json2csv.Convert(
		bytes.NewReader([]byte("NO GOOD")), ioutil.Discard,
	)
	if err == nil {
		t.Error("expected JSON parsing to fail")
	}
}
