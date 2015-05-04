package ics

import (
	"bytes"
	"io"
)

var (
	lineWrap = []byte{'\r', '\n', ' '}
	lineEnd  = []byte{'\r', '\n'}
)

type unfolder struct {
	buf bytes.Buffer
}

func NewUnfolder(r io.Reader) (*unfolder, error) {
	var u unfolder
	_, err := u.buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u *unfolder) ReadLine() ([]byte, error) {
	var toRet []byte
	for {
		p, err := u.buf.ReadBytes('\r')
		if err != nil {
			return nil, err
		}
		c, err := u.buf.ReadByte()
		if err != nil {
			return nil, err
		}
		if c != '\n' {
			u.buf.UnreadByte()
			toRet = append(toRet, p...)
			continue
		}
		toRet = append(toRet, p[:len(p)-1]...)
		c, _ := u.buf.ReadByte()
		if c != ' ' {
			u.buf.UnreadByte()
			break
		}
	}
	return toRet, nil
}
