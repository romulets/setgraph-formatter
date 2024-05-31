# setgraph-formatter

Silly tool to convert set graph format to format expect by my coach

## Usage

### Install

Download the repo and `go install` or `go install github.com/romulets/setgraph-formatter`

### Running

```shell
setgraph-formatter $IN_FILE
```

`$IN_FILE` can have multiple sessions, separated by a blank line. See `./testdata/in.txt` for example.

Output: converted sessions

common usage to get only the values:

```shell
setgraph-formatter setgraph.in | awk -F '\t' '{print $2}'
```

If you rather get it saved to a file, pass the `-f` param, like

```shell
setgraph-formatter $IN_FILE -f
```

If you rather read from clipboard instead of file:

```shell
setgraph-formatter -c
```

There is also an option to sort the output based on a list of values provided in file, the option is `-s $FILE`. 
You can see an example of the file in `./testdata/format.txt`

Personal current favorite flavor:

```shell
setgraph-formatter -c -s ~/sort.idx | awk -F '\t' '{print $2}'
```
