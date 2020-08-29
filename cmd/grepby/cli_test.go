// Copyright 2015-2020 Ray Holder
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io"
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

func TestCliArgsTail(t *testing.T) {
	args := []string{"aaa", "--tail"}
	in := strings.NewReader("potato\naaa\nmeep")
	err := cli(args, in, ioutil.Discard, ioutil.Discard)
	if err != nil {
		t.Fatal("Unexpected error")
	}

	// TODO add more tests to verify counts
}

func TestCliArgsErr(t *testing.T) {
	args := []string{".**"}
	in := strings.NewReader("potato\naaa\nmeep")
	err := cli(args, in, ioutil.Discard, ioutil.Discard)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "error parsing regexp") {
			t.Fatal("Expected a regexp error")
		}
	} else {
		t.Fatal("Expected an error")
	}
}

func TestCliHelp(t *testing.T) {
	args := []string{"a", "b", "c", "--help"}

	// --help should short circuit before reading
	var in io.Reader
	err := cli(args, in, ioutil.Discard, ioutil.Discard)
	if err != nil {
		t.Fatal("Unexpected error")
	}
}

func TestCliVersion(t *testing.T) {
	args := []string{"a", "b", "c", "--version"}

	// --version should short circuit before reading
	var in io.Reader
	err := cli(args, in, ioutil.Discard, ioutil.Discard)
	if err != nil {
		t.Fatal("Unexpected error")
	}
}

func TestCliErr(t *testing.T) {
	args := []string{"a", "b", "c", "--tail=potato"}

	// err should short circuit before reading
	var in io.Reader
	err := cli(args, in, ioutil.Discard, ioutil.Discard)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "strconv.ParseUint") {
			t.Fatal("Expected a parse error")
		}
	} else {
		t.Fatal("Expected an error")
	}
}

func checkPrint(t *testing.T, expected bool, invert bool, matchedLine bool, outputMatches bool) {
	var print = shouldPrintMatch(invert, matchedLine, outputMatches)
	if expected != print {
		t.Fatal("Unexpected value for print,", expected, invert, matchedLine, outputMatches)
	}
}

func TestCliShouldPrintMatch(t *testing.T) {
	checkPrint(t, false, false, false, false)
	checkPrint(t, false, false, false, true)
	checkPrint(t, false, false, true, false)
	checkPrint(t, true, false, true, true)

	// there shouldn't be a case where we invert but don't output, so meh
	checkPrint(t, true, true, false, false)

	checkPrint(t, true, true, false, true)
	checkPrint(t, false, true, true, false)
	checkPrint(t, false, true, true, true)
}
