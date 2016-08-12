package ics

import "io"

type section interface {
	decode(tokeniser) error
	encode(w writer)
	valid() error
}

func Decode(r io.Reader) (*SectionCalendar, error) {
	cal := new(SectionCalendar)
	err := cal.decode(newTokeniser(&unfolder{r: r}))
	return cal, err
}
