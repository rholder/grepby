package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const version = "1.0.0"
const usageText = `Usage: grepby [regex1] [regex2] [regex3]...

  Use grepby to count lines that match regular expressions. It's a bit like
  having group by for grep.

  By default, all of stdin in read and the aggregate counts are output to
  stdout. When --tail or --output are used or combined, counts are output to
  stderr and matching lines are output to stdout.

Options:
  --help          Print this help
  --tail          Print aggregate output every 2 seconds to stderr
  --tail=10       Print aggregate output every 10 seconds to stderr
  --output        Print all lines that match at least one regex to stdout
  --version       Print the version number

Examples:
  grepby 'potato' 'banana' '[Tt]omato' < groceries.txt")
  20% -  600 - potato")
  13% -  400 - banana")
  17% -  500 - [Tt]omato")
  50% - 1500 - (unmatched)")

Report bugs and find the latest updates at https://github.com/rholder/grepby.
`

type Config struct {
	help          bool
	tail          bool
	tailDelay     float64
	output        bool
	countOutput   io.Writer
	matchOutput   io.Writer
	patterns      []string
	countTemplate string
	version       bool
}

type PatternCount struct {
	pattern string
	count   uint64
	regex   *regexp.Regexp
}

type Rollup struct {
	config   *Config
	patterns []*PatternCount
	total    uint64
}

func newRollup(config *Config) (*Rollup, error) {
	rollup := Rollup{}
	rollup.total = 0
	rollup.config = config
	for _, pattern := range config.patterns {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		pc := PatternCount{pattern, 0, regex}
		rollup.patterns = append(rollup.patterns, &pc)
	}
	return &rollup, nil
}

func newConfig(args []string, stdout io.Writer, stderr io.Writer) (*Config, error) {
	config := Config{}
	config.countOutput = stdout
	config.tailDelay = 2.0

	enableTail := false
	enableOutput := false

	// default is to output a count to stdout when complete
	var patterns []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--tail") {
			enableTail = true
			if strings.HasPrefix(arg, "--tail=") {
				td, err := strconv.Atoi(arg[7:])
				if err != nil {
					return nil, err
				}
				config.tailDelay = float64(td)
			}
		} else if "--output" == arg {
			enableOutput = true
		} else if "--version" == arg {
			config.version = true
		} else if "--help" == arg {
			config.help = true
		} else {
			patterns = append(patterns, arg)
		}
	}
	config.patterns = patterns

	// --tail always outputs counts to stderr
	if enableTail {
		config.tail = true
		config.countOutput = stderr
	}

	// --output outputs matches to stdout and forces counts to stderr
	if enableOutput {
		config.output = true
		config.countOutput = stderr
		config.matchOutput = stdout
	}

	// TODO make configurable via argument
	config.countTemplate = "%4.0f%% - %6v - %v" + "\n"

	return &config, nil
}

// Output the rollup counts.
func outputCounts(rollup *Rollup) {
	var totalMatched uint64 = 0
	output := rollup.config.countOutput
	template := rollup.config.countTemplate

	for _, pc := range rollup.patterns {
		totalMatched += pc.count
	}

	if rollup.config.tail {
		fmt.Fprintf(output, "(last %v lines)\n", rollup.total)
	}

	totalUnmatched := rollup.total - totalMatched
	for _, pc := range rollup.patterns {
		var percentMatched float64 = 0
		if rollup.total != 0 {
			percentMatched = 100 * float64(pc.count) / float64(rollup.total)
		}
		fmt.Fprintf(output, template, percentMatched, pc.count, pc.pattern)
	}
	var percentUnmatched float64 = 0
	if rollup.total != 0 {
		percentUnmatched = 100 * float64(totalUnmatched) / float64(rollup.total)
	}
	fmt.Fprintf(output, template, percentUnmatched, totalUnmatched, "(unmatched)")
}

// Update counts from the given input line. Return true if there was a match.
func updateCounts(rollup *Rollup, line string) bool {
	rollup.total += 1
	for _, pc := range rollup.patterns {
		// only first matching pattern counts
		if pc.regex.MatchString(line) {
			pc.count += 1
			return true
		}
	}
	return false
}

func cli(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	if len(args) == 0 {
		fmt.Fprintln(stdout, usageText)
		return errors.New("Invalid number of arguments.")
	}

	config, err := newConfig(args, stdout, stderr)
	if err != nil {
		return err
	}

	// short circuit on --version
	if config.version {
		fmt.Fprintln(stdout, version)
		return nil
	}

	// short circuit on --help
	if config.help {
		fmt.Fprintln(stdout, usageText)
		return nil
	}

	rollup, err := newRollup(config)
	if err != nil {
		return err
	}

	// read from input
	last := time.Now()
	scanner := bufio.NewScanner(stdin)
	shouldOutputMatch := rollup.config.output
	matchOutput := rollup.config.matchOutput
	for scanner.Scan() {
		line := scanner.Text()
		matched := updateCounts(rollup, line)
		if shouldOutputMatch && matched {
			fmt.Fprintln(matchOutput, line)
		}
		if config.tail {
			now := time.Now()
			if now.Sub(last).Seconds() > config.tailDelay {
				outputCounts(rollup)
				last = now
			}
		}
	}
	outputCounts(rollup)
	return nil
}

func main() {
	args := os.Args[1:]
	err := cli(args, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
