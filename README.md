# grepby [![Build Status](https://travis-ci.org/rholder/grepby.svg?branch=master)](https://travis-ci.org/rholder/grepby)

Use `grepby` to count lines that match regular expressions. It's a bit like
having group by for grep.

## Installation
```bash
go get github.com/rholder/grepby
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

## Usage
```
Usage: grepby [regex1] [regex2] [regex3]...

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
```

## License
Debinate is released under version 2.0 of the
[Apache License](http://www.apache.org/licenses/LICENSE-2.0).