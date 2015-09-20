package parser

import (
	"reflect"
	"testing"

	ts "github.com/anonx/espresso/parser/testdata"
)

var correctFiles = map[string]interface{}{
	"./testdata/basic.espresso": ts.Basic,
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
		}
	}
}

func TestParseFile(t *testing.T) {
	for path, exp := range correctFiles {
		res, err := ParseFile(path)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(exp, res) {
			t.Errorf("Incorrect result of ParseFile. Expected:\n%v\nGot:\n%v", exp, res)
		}
	}
}

func TestParseFile_Debug(t *testing.T) {
	for path, exp := range correctFiles {
		res, err := ParseFile(path, Debug(true), Memoize(true), Recover(true))
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(exp, res) {
			t.Errorf("Incorrect result of ParseFile. Expected:\n%v\nGot:\n%v", exp, res)
		}
	}
}
