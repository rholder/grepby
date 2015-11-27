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
	"testing"
)

// Return true if any pattern in the rollup matches the string.
func matches(rollup *Rollup, value string) bool {
	for _, pc := range rollup.patterns {
		if pc.regex.MatchString(value) {
			return true
		}
	}
	return false
}

func TestRollup(t *testing.T) {
	acceptable := []string{
		"abc", "def", "ghi", "abcdef", "ghiblahblah", "defabc",
	}
	unacceptable := []string{
		"foo", "bar", "foobar",
	}

	args := []string{"abc", "def", "ghi"}
	config, err := newConfig(args, ioutil.Discard, ioutil.Discard)
	rollup, err := newRollup(config)
	if err != nil {
		t.Fatal("Unxpected rollup error:", err)
	}

	if rollup.total != 0 {
		t.Fatal("Expected rollup total to start at 0:", rollup.total)
	}

	for _, expected := range acceptable {
		if !matches(rollup, expected) {
			t.Fatal("Expected to match:", expected)
		}
	}

	for _, unexpected := range unacceptable {
		if matches(rollup, unexpected) {
			t.Fatal("Not expected to match:", unexpected)
		}
	}
}
