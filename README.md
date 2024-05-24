# setgraph-formatter

Silly tool to convert set graph format to format expect by my coach

## Usage

### Install

Download the repo and `go install` or `go install github.com/romulets/setgraph-formatter`

### Running

```
setgraph-formatter $IN_FILE
```

`$IN_FILE` can have multiple sessions, separated by a blank line. See `./testdata/in.txt` for example.

Output: new file with today's date containing the parsed file