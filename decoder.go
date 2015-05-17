package ics

import (
	"errors"
	"io"
)

type Decoder struct {
	p *parser
}

func NewDecoder(r io.Reader) Decoder {
	return Decoder{
		p: newParser(r),
	}
}

func (d Decoder) Decode() (*Calendar, error) {
	p, err := d.p.GetProperty()
	if err != nil {
		return nil, err
	}
	if c, ok := p.(begin); ok && c == vCalendar {
		cal := new(Calendar)
		err = cal.decode(d)
		if err != nil {
			return nil, err
		}
		return cal, nil
	}
	return nil, ErrInvalidStart
}

func (d Decoder) readUnknownComponent(name string) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case begin:
			err = d.readUnknownComponent(string(p))
			if err != nil {
				return nil
			}
		case end:
			if string(p) == name {
				return nil
			}
			return ErrInvalidEnd
		}
	}
}

// Errors

var (
	ErrInvalidStart    = errors.New("invalid start component")
	ErrInvalidEnd      = errors.New("invalid end component")
	ErrMultipleUnique  = errors.New("multiple of a unique property")
	ErrRequiredMissing = errors.New("required property missing")
)
