package ics

import (
	"bufio"
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
	bw *bufio.Writer
}

func newFolder(w io.Writer) folder {
	return folder{bufio.NewWriterSize(w, 1024)}
}

func (f folder) WriteLine(p []byte) (err error) {
	pos := 0
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		if pos+s > lineLength {
			f.bw.Write(continuation)
			pos = 0
		}
		_, err = f.bw.WriteRune(r)
		if err != nil {
			return err
		}
		pos += s
		p = p[s:]
	}
	_, err = f.bw.Write(endLine)
	if err != nil {
		return err
	}
	return f.bw.Flush()
}
