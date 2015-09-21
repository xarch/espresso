package parser

import (
	"reflect"
	"testing"

	ts "github.com/anonx/espresso/parser/testdata"
)

var correctFiles = map[string]interface{}{
	"./testdata/boolean.espresso":       ts.Boolean,
	"./testdata/env.espresso":           ts.Env,
	"./testdata/integer.espresso":       ts.Integer,
	"./testdata/key_quotless.espresso":  ts.KeyQuotless,
	"./testdata/list.espresso":          ts.List,
	"./testdata/object.espresso":        ts.Object,
	"./testdata/object_inline.espresso": ts.ObjectInline,
	"./testdata/string.espresso":        ts.String,
	"./testdata/string_inline.espresso": ts.StringInline,

	"./testdata/complex.espresso": ts.Complex,
}

var incorrectFiles = []string{
	"./testdata/unclosedBraces.espresso",
}

func TestParseFile_NonExistent(t *testing.T) {
	res, err := ParseFile("file_that_does_not_exist")
	if res != nil || err == nil {
		t.Error("File does not exist, error expected.")
	}
}

func TestParseFile_Incorrect(t *testing.T) {
	for _, path := range incorrectFiles {
		res, err := ParseFile(path, Debug(true))
		if res != nil || err == nil {
			t.Error(err)
			t.FailNow()
		}
	}
}

func TestParseFile(t *testing.T) {
	for path, exp := range correctFiles {
		res, err := ParseFile(path)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if !reflect.DeepEqual(exp, res) {
			t.Errorf("Incorrect result of ParseFile. Expected:\n%v\nGot:\n%v", exp, res)
			t.FailNow()
		}
	}
}

func TestParseFile_Debug(t *testing.T) {
	for path, exp := range correctFiles {
		res, err := ParseFile(path, Debug(true), Memoize(true), Recover(true))
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if !reflect.DeepEqual(exp, res) {
			t.Errorf("Incorrect result of ParseFile. Expected:\n%v\nGot:\n%v", exp, res)
			t.FailNow()
		}
	}
}
