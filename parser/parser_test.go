package parser

import (
	"reflect"
	"testing"

	ts "github.com/anonx/espresso/parser/testdata"
)

func TestParseFile(t *testing.T) {
	res, err := ParseFile("./testdata/basic.espresso")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(ts.Basic, res) {
		t.Errorf("Incorrect result of ParseFile. Expected:\n%v\nGot:\n%v", ts.Basic, res)
	}
}
