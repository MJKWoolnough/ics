package ics

import (
	"io"
	"unicode/utf8"
)

type folder struct {
	w    io.Writer
	line uint8
}

const maxLineLength = 75

var eol = [...]byte{'\r', '\n', ' '}

func (f *folder) Write(p []byte) (int, error) {
	q := p
	var (
		r    rune
		s, n int
	)
	for pos := 0; pos < len(q); pos += s {
		r, s = utf8.DecodeRune(q[pos:])
		f.line += s
		if r == '\n' {
			f.line = 0
		} else if f.line > maxLineLength {
			if pos > 0 {
				m, err := f.w.Write(q[:pos])
				n += m
				if err != nil {
					return n, err
				}
				q = q[pos:]
			}
			_, err := f.w.Write(eol[:])
			if err != nil {
				return n, err
			}

			pos = 0
			f.line = s
		}
	}
	if len(q) > 0 {
		m, err := f.w.Write(q)
		n += m
		return n, err
	}
	return n, nil
}
