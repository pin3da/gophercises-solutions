package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseLinks(t *testing.T) {
	filenames := [...]string{"ex1.html", "ex2.html", "ex3.html", "ex4.html"}

	expected := [][]Link{
		[]Link{
			Link{"/other-page", "A link to another page"},
		},
		[]Link{
			Link{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
			Link{"https://github.com/gophercises", "Gophercises is on Github!"},
		},
		[]Link{
			Link{"#", "Login"},
			Link{"/lost", "Lost? Need help?"},
			Link{"https://twitter.com/marcusolsson", "@marcusolsson"},
		},
		[]Link{
			Link{"/dog-cat", "dog cat"},
		},
	}

	for i, name := range filenames {
		file, err := os.Open("test_data/" + name)
		if err != nil {
			t.Errorf("Provided test file does not exist: %v\n", err)
		}
		defer file.Close()
		actual, err := ParseLinks(file)
		if !reflect.DeepEqual(actual, expected[i]) {
			t.Errorf("failed to parse %v", name)
		}
	}
}
