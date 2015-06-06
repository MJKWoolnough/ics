package ics

import (
	"io"
	"unicode/utf8"
)

const lineLength = 75

var (
	delims       = [...]byte{'\r', '\n', ' '}
	continuation = delims[:]
	endLine      = delims[:2]
)

type folder struct {
	w io.Writer
}

func newFolder(w io.Writer) folder {
	return folder{w}
}

func (f folder) writeLine(p []byte) (err error) {
	pos := 0
	var utf [5]byte
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		if pos+s > lineLength {
			f.w.Write(continuation)
			pos = 1
		}
		l := utf8.EncodeRune(utf[:], r)
		_, err = f.w.Write(utf[:l])
		if err != nil {
			return err
		}
		pos += s
		p = p[s:]
	}
	_, err = f.w.Write(endLine)
	if err != nil {
		return err
	}
	return nil
}
