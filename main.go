package main

import (
	"fmt"
	"io"
	"os"
)

const (
	FlagSTDIN = "-"
	FlagCmd   = "--"
)

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	switch os.Args[1] {
	case FlagSTDIN:
		w := NewWriter(os.Stdout)
		_, err := io.Copy(w, os.Stdin)
		exitIf(err, 1)
	case FlagCmd:
		if len(os.Args) <= 2 {
			usage()
		}

		code, err := command(os.Args[2:])
		exitIf(err, code)
		return
	default:
		usage()
	}

}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: pride -- [command]")
	os.Exit(1)
}

func exitIf(err error, code int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "pride: %q\n", err)
		os.Exit(code)
	}
}
