package ics

import (
	"errors"
	"io"
)

type writer interface {
	Write([]byte) (int, error)
	WriteString(string) (int, error)
}

func Encode(w io.Writer, cal *Calendar) error {
	if err := cal.valid(); err != nil {
		return err
	}
	f := folder{w: w}
	cal.encode(&f)
	return f.err
}

// Errors
var (
	ErrInvalidCalendar = errors.New("invalid calendar")
)
