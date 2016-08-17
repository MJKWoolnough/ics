package ics

import (
	"errors"
	"io"
)

type writer interface {
	Write([]byte) (int, error)
	WriteString(string) (int, error)
}

// Encode encodes the given iCalendar object into the writer. It first
// validates the iCalendar object so as not to write invalid data to the writer
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
