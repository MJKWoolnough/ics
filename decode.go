package ics

import "io"

type section interface {
	decode(tokeniser) error
	encode(w writer)
	valid() bool
}

func Decode(r io.Reader) (*Calendar, error) {
	cal := new(Calendar)
	err := cal.decode(newTokeniser(&unfolder{r: r}))
	return cal, err
}
