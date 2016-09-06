package ics

import (
	"io"
	"unicode/utf8"
)

type folder struct {
	w    io.Writer
	err  error
	line uint8
}

const maxLineLength = 75

var eol = [...]byte{'\r', '\n', ' '}

func (f *folder) Write(q []byte) (int, error) {
	if f.err != nil {
		return 0, f.err
	}
	var (
		r       rune
		s, n, m int
	)
	for pos := 0; pos < len(q); pos += s {
		r, s = utf8.DecodeRune(q[pos:])
		f.line += uint8(s)
		if r == '\n' {
			f.line = 0
		} else if r == '\r' {
		} else if f.line > maxLineLength {
			if pos > 0 {
				m, f.err = f.w.Write(q[:pos])
				n += m
				if f.err != nil {
					return n, f.err
				}
				q = q[pos:]
			}
			_, f.err = f.w.Write(eol[:])
			if f.err != nil {
				return n, f.err
			}

			pos = 0
			f.line = uint8(s) + 1
		}
	}
	if len(q) > 0 {
		m, f.err = f.w.Write(q)
		n += m
	}
	return n, f.err
}

func (f *folder) WriteString(q string) (int, error) {
	return f.Write([]byte(q))
}
