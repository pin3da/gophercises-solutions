package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseLinks(t *testing.T) {
	filenames := [...]string{"ex2.html"}

	expected := []Link{
		Link{"https://www.twitter.com/joncalhoun", "Check me out on twitter <i class=\"fa fa-twitter\" aria-hidden=\"true\"></i>"},
		Link{"https://github.com/gophercises", "Gophercises is on <strong>Github</strong>!"},
	}

	for _, name := range filenames {
		file, err := os.Open("test_data/" + name)
		if err != nil {
			t.Errorf("Provided test file does not exist: %v\n", err)
		}
		defer file.Close()
		actual, err := ParseLinks(file)
		if !reflect.DeepEqual(actual, expected) {
			t.Fail()
		}
	}
}
