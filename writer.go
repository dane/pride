package main

import (
	"bytes"
	"io"

	"github.com/fatih/color"
)

const NL = '\n'

var colors = [...]*color.Color{
	color.New(color.FgRed),
	NewOrange(),
	color.New(color.FgYellow),
	color.New(color.FgGreen),
	color.New(color.FgBlue),
	NewViolet(),
}

func NewOrange() *color.Color {
	return color.New(
		color.Attribute(38),
		color.Attribute(2),
		color.Attribute(255),
		color.Attribute(165),
		color.Reset,
	)
}

func NewViolet() *color.Color {
	return color.New(
		color.Attribute(38),
		color.Attribute(5),
		color.Attribute(55),
	)
}

func NewWriter(dst io.Writer) io.Writer {
	return &writer{dst: dst}
}

type writer struct {
	dst io.Writer
	idx int
}

func (w *writer) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
	count := len(b)
	for {
		line, err := buf.ReadBytes(NL)
		if err != nil {
			if err == io.EOF {
				break
			}

			return count, err
		}

		c := colors[w.idx]
		w.idx++
		if w.idx >= len(colors) {
			w.idx = 0
		}

		if _, err = c.Fprintf(w.dst, "%s", line); err != nil {
			return count, err
		}
	}

	return count, nil
}
