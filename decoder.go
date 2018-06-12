// Package ics implements an encoder and decoder for iCalendar files
package ics // import "vimagination.zapto.org/ics"

import (
	"io"
	"strings"
)

type section interface {
	decode(tokeniser) error
	encode(w writer)
	valid() error
}

// Decode decodes an iCalendar object from the given reader
func Decode(r io.Reader) (*Calendar, error) {
	t := newTokeniser(&unfolder{r: r})
	if p, err := t.GetPhrase(); err != nil {
		return nil, err
	} else if p.Type != phraseContentLine {
		if t.Err == io.EOF {
			return nil, io.ErrUnexpectedEOF
		}
		return nil, t.Err
	} else if strings.ToUpper(p.Data[0].Data) != "BEGIN" || strings.ToUpper(p.Data[len(p.Data)-1].Data) != "VCALENDAR" {
		return nil, ErrInvalidCalendar
	}
	cal := new(Calendar)
	if err := cal.decode(t); err != nil {
		return nil, err
	}
	return cal, nil
}
