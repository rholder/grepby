// Copyright 2015 Ray Holder
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
