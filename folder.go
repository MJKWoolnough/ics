package ics

import (
	"io"
	"unicode/utf8"
)

const lineLength = 75

type folder struct {
	w io.Writer
}

func newFolder(w io.Writer) folder {
	return folder{w}
}

func (f folder) WriteLine(p []byte) (err error) {
	pos := 0
	var bufArr [1024]byte
	buf := bufArr[:0]
	for len(p) > 0 {
		_, s := utf8.DecodeRune(p)
		if pos+s > lineLength {
			if len(buf)+3 > 1024 {
				_, err = f.w.Write(buf)
				if err != nil {
					return err
				}
				buf = buf[:0]
			}
			buf = append(buf, '\r', '\n', ' ')
			pos = 0
		}
		if len(buf)+s > 1021 {
			_, err = f.w.Write(buf)
			if err != nil {
				return err
			}
			buf = buf[:0]
		}
		buf = append(buf, p[:s])
		pos += s
		p = p[s:]
	}
	_, err = f.w.Write(buf)
	return err
}
