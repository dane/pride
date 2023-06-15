package main

import (
	"bytes"
	"io"
	"testing"
)

func TestStdin(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
	}{
		{
			name: "write 6 lines",
			lines: []string{
				"foo\n",
				"bar\n",
				"baz\n",
				"cux\n",
				"qux\n",
				"quux\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build input and expected output
			var (
				in  bytes.Buffer
				exp bytes.Buffer
			)

			for i, line := range tt.lines {
				if _, err := in.WriteString(line); err != nil {
					t.Fatal(err)
				}

				_, err := colors[i].Fprintf(&exp, "%s", line)
				if err != nil {
					t.Fatal(err)
				}
			}

			// Build writer
			var out bytes.Buffer
			w := NewWriter(&out)
			if _, err := io.Copy(w, &in); err != nil {
				t.Fatal(err)
			}

			// Compare output to expected output
			want := exp.String()
			got := out.String()
			if got != want {
				t.Fatalf("got:\n%s\nwant:\n%s\n", got, want)
			}
		})
	}
}
