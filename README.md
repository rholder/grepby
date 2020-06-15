# grepby
[![Build Status](http://img.shields.io/travis/rholder/grepby.svg)](https://travis-ci.org/rholder/grepby)
[![Latest Version](https://img.shields.io/github/v/release/rholder/grepby?color=bright-green&sort=semver)](https://github.com/rholder/grepby/releases/latest)
[![License](http://img.shields.io/badge/license-apache%202-brightgreen.svg)](https://github.com/rholder/grepby/blob/master/LICENSE)
[![Coverage](https://gocover.io/_badge/github.com/rholder/grepby/cmd/grepby)](https://gocover.io/github.com/rholder/grepby/cmd/grepby)
[![Go Report Card](https://goreportcard.com/badge/github.com/rholder/grepby)](https://goreportcard.com/report/github.com/rholder/grepby)

Use `grepby` to count lines that match regular expressions. It's a bit like
having group by for grep. If you've ever wanted to sample a fast scrolling log
file and get an idea about what the percentage of several types of errors look
like as they stream by, then this might be the tool you've been searching for.

## Features
* Process an entire stream and give full aggregate match counts
* Monitor a live stream of lines with --tail sending counts to stderr
* Composable API with --output of lines matching 1 or more regexes
* Output inverted, non-matching lines with --invert

## Installation
Release binaries are available for several platforms.

#### Linux
Drop the binary into your path, such as `/usr/local/bin`:
```
sudo curl -o /usr/local/bin/grepby -L "https://github.com/rholder/grepby/releases/download/v1.2.0/grepby_linux_amd64" && \
sudo chmod +x /usr/local/bin/grepby
```

#### OSX
Drop the binary into your path, such as `/usr/local/bin`:
```
sudo curl -o /usr/local/bin/grepby -L "https://github.com/rholder/grepby/releases/download/v1.2.0/grepby_darwin_amd64" && \
sudo chmod +x /usr/local/bin/grepby
```

#### Windows
Download the `.exe` from the latest [releases](https://github.com/rholder/grepby/releases/latest).

#### Source
Install it from source with `go get`:
```bash
go get github.com/rholder/grepby
```

## Usage
```
Usage: grepby [regex1] [regex2] [regex3]...

  Use grepby to count lines that match regular expressions. It's a bit like
  having group by for grep.

  By default, all of stdin is read and the aggregate counts are output to
  stdout. When --tail or --output are used or combined, counts are output to
  stderr and matching lines are output to stdout. When --invert is used,
  non-matching lines are output to stdout and counts are output to stderr.

Options:

  --help          Print this help
  --tail          Print aggregate output every 2 seconds to stderr
  --tail=10       Print aggregate output every 10 seconds to stderr
  --output        Print all lines that match at least one regex to stdout
  --invert        Invert matching and output non-matching lines
  --version       Print the version number
```

## Examples
Read an entire file and output a count for each regex to stdout:
```bash
grepby potato banana '[Tt]omato' < groceries.txt
```
```
 20% -   600 - potato
 13% -   400 - banana
 17% -   500 - [Tt]omato
 50% -  1500 - (unmatched)
```

Tail a log file, sending a count for each regex to stderr every 2 seconds:
```bash
tail -f app.log | grepby --tail ERROR SUCCESS
```
```
(last 200 lines)
 20% -    40 - ERROR
 80% -   160 - SUCCESS
  0% -     0 - (unmatched)
```

Tail a log file, printing all matching lines to stdout:
```bash
tail -f app.log | grepby --output FATAL ERROR WARNING
```
```
WARNING a weird thing happened
WARNING another weird thing happened
ERROR an error occurred
FATAL unrecoverable error
WARNING a bad thing happened
```

## License
Grepby is released under version 2.0 of the
[Apache License](http://www.apache.org/licenses/LICENSE-2.0).
