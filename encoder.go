package ics

import "io"

type Encoder struct {
	folder
}

func NewEncoder(w io.Writer) Encoder {
	return Encoder{
		folder: newFolder(w),
	}
}

func (e Encoder) Encode(c *Calendar) error {
	err := c.encode(e)
	if err != nil {
		return err
	}
	return e.folder.flush()
}
