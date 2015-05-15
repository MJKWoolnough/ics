package ics

import (
	"strconv"
	"time"
)

type created struct {
	time.Time
}

func (p *parser) readCreateComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var c created
	c.Time, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type dateStamp struct {
	time.Time
}

func (p *parser) readDateStampComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var d dateStamp
	d.Time, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return d, nil
}

type lastModified struct {
	time.Time
}

func (p *parser) readLastModifiedComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var l lastModified
	l.Time, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return l, nil
}

type sequence int

func (p *parser) readSequenceComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	s, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	if s < 0 {
		return nil, ErrUnsupportedValue
	}
	return sequence(s), nil
}
