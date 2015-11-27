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
