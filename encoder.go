package ics

import (
	"bufio"
	"errors"
	"io"
)

type writer interface {
	Write([]byte) (int, error)
	WriteString(string) (int, error)
}

type writeError struct {
	w   io.Writer
	err error
}

func (w *writeError) Write(p []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	var n int
	n, w.err = w.w.Write(p)
	return n, w.err
}

func Encode(w io.Writer, cal *Calender) error {
	if !cal.valid() {
		return ErrInvalidCalendar
	}
	we := writeError{
		w: &folder{w: w},
	}
	b := bufio.NewWriter(&we)
	cal.encode(b)
	b.Flush()
	return we.err
}

// Errors
var (
	ErrInvalidCalendar = errors.New("invalid calendar")
)