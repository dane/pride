# pride

Run any terminal command with pride!

## Usage

Have `pride` execute your command and decorate the output:

```
$ pride -- go test -v ./...
```

Pipe any `STDOUT` into `pride`:

```
$ go test -v ./... | pride -
```

## Installation

```
$ go install github.com/dane/pride
```

## Credit

Thank you @fatih for [github.com/fatih/color](https://github.com/fatih/color)
and @mattn for [github.com/mattn/go-colorable](https://github.com/mattn/go-colorable).
