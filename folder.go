package ics

import "unicode/utf8"

type folder struct {
	w    writer
	line uint8
}

const maxLineLength = 75

var eol = [...]byte{'\r', '\n', ' '}

func (f *folder) Write(q []byte) (int, error) {
	var (
		r    rune
		s, n int
	)
	for pos := 0; pos < len(q); pos += s {
		r, s = utf8.DecodeRune(q[pos:])
		f.line += uint8(s)
		if r == '\n' {
			f.line = 0
		} else if r == '\r' {
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
			f.line = uint8(s)
		}
	}
	if len(q) > 0 {
		m, err := f.w.Write(q)
		n += m
		return n, err
	}
	return n, nil
}

func (f *folder) WriteString(q string) (int, error) {
	var (
		r    rune
		s, n int
	)
	for pos := 0; pos < len(q); pos += s {
		r, s = utf8.DecodeRuneInString(q[pos:])
		f.line += uint8(s)
		if r == '\n' {
			f.line = 0
		} else if r == '\r' {
		} else if f.line > maxLineLength {
			if pos > 0 {
				m, err := f.w.WriteString(q[:pos])
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
			f.line = uint8(s)
		}
	}
	if len(q) > 0 {
		m, err := f.w.WriteString(q)
		n += m
		return n, err
	}
	return n, nil
}
