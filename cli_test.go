package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCliNoArgs(t *testing.T) {
	args := []string{}
	in := strings.NewReader("")
	err := cli(args, in, ioutil.Discard, ioutil.Discard)
	if err == nil {
		t.Fatal("Expected an error")
	} else {
		if err.Error() != "Invalid number of arguments." {
			t.Fatal("Expected argument error")
		}
	}
}

func TestCliArgs(t *testing.T) {
	args := []string{"aaa"}
	in := strings.NewReader("potato\naaa\nmeep")
	err := cli(args, in, ioutil.Discard, ioutil.Discard)
	if err != nil {
		t.Fatal("Unexpected error")
	}

	// TODO add more tests to verify counts
}
