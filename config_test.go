package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestConfigDefault(t *testing.T) {
	args := []string{"a", "b", "c"}
	config, _ := newConfig(args, os.Stdout, os.Stderr)

	if config.tail {
		t.Fatal("Expected no --tail")
	}

	if !reflect.DeepEqual(args, config.patterns) {
		t.Fatal("Expected all arguments to pass through")
	}

	if config.countOutput != os.Stdout {
		t.Fatal("Expected default count output to be stdout")
	}

	if config.matchOutput != nil {
		t.Fatal("Expected default match output to be nil")
	}
}

func TestConfigTail(t *testing.T) {
	args := []string{"a", "b", "c", "--tail"}
	expectedParameters := []string{"a", "b", "c"}
	config, _ := newConfig(args, os.Stdout, os.Stderr)

	if !config.tail {
		t.Fatal("Expected --tail")
	}

	if config.tailDelay != 2 {
		t.Fatal("Expected 2 tail delay")
	}

	if !reflect.DeepEqual(expectedParameters, config.patterns) {
		t.Fatal("Expected only parameters to pass through")
	}

	if config.countOutput != os.Stderr {
		t.Fatal("Expected count output to be stderr")
	}

	if config.matchOutput != nil {
		t.Fatal("Expected match output to be nil")
	}
}

func TestConfigTailNumbers(t *testing.T) {
	args := []string{"a", "b", "c", "--tail=99"}
	expectedParameters := []string{"a", "b", "c"}
	config, _ := newConfig(args, os.Stdout, os.Stderr)

	if !config.tail {
		t.Fatal("Expected --tail")
	}

	if config.tailDelay != 99 {
		t.Fatal("Expected 99 tail delay")
	}

	if !reflect.DeepEqual(expectedParameters, config.patterns) {
		t.Fatal("Expected only parameters to pass through")
	}

	if config.countOutput != os.Stderr {
		t.Fatal("Expected count output to be stderr")
	}

	if config.matchOutput != nil {
		t.Fatal("Expected match output to be nil")
	}
}

func TestConfigTailBogus(t *testing.T) {
	args := []string{"a", "b", "c", "--tail=99potato"}
	config, err := newConfig(args, os.Stdout, os.Stderr)

	if config != nil {
		t.Fatal("Unxpected config created")
	}

	if err == nil {
		t.Fatal("Expected an error")
	} else {
		if !strings.HasPrefix(err.Error(), "strconv.ParseInt") {
			t.Fatal("Expected strconv.ParseInt")
		}
	}
}

func TestConfigOutput(t *testing.T) {
	args := []string{"a", "b", "c", "--output"}
	expectedParameters := []string{"a", "b", "c"}
	config, _ := newConfig(args, os.Stdout, os.Stderr)

	if config.tail {
		t.Fatal("Did not expect --tail")
	}

	if !config.output {
		t.Fatal("Expected --output")
	}

	if !reflect.DeepEqual(expectedParameters, config.patterns) {
		t.Fatal("Expected only parameters to pass through")
	}

	if config.countOutput != os.Stderr {
		t.Fatal("Expected count output to be stderr")
	}

	if config.matchOutput != os.Stdout {
		t.Fatal("Expected match output to be stdout")
	}
}

func TestConfigTailOutput(t *testing.T) {
	args := []string{"a", "b", "c", "--output", "--tail"}
	expectedParameters := []string{"a", "b", "c"}
	config, _ := newConfig(args, os.Stdout, os.Stderr)

	if !config.tail {
		t.Fatal("Expected --tail")
	}

	if config.tailDelay != 2 {
		t.Fatal("Expected 2 tail delay")
	}

	if !config.output {
		t.Fatal("Expected --output")
	}

	if !reflect.DeepEqual(expectedParameters, config.patterns) {
		t.Fatal("Expected only parameters to pass through")
	}

	if config.countOutput != os.Stderr {
		t.Fatal("Expected count output to be stderr")
	}

	if config.matchOutput != os.Stdout {
		t.Fatal("Expected match output to be stdout")
	}
}
