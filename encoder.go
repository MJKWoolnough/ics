package ics

import (
	"errors"
	"io"
)

type Encoder struct {
	f   folder
	err error
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		f: newFolder(w),
	}
}

func (e *Encoder) Encode(c *Calendar) error {
	c.encode(e)
	return e.err
}

func (e *Encoder) writeProperty(p property) {
	if e.err != nil {
		return
	}
	if !p.Validate() {
		e.err = ErrInvalidProperty
		return
	}
	e.err = e.f.writeLine(p.Data().Bytes())
}

// Errors

var ErrInvalidProperty = errors.New("property failed to validate")
