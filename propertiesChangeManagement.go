package ics

import "strconv"

type created struct {
	dateTime
}

func (p *parser) readCreateProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var c created
	c.dateTime, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type dateStamp struct {
	dateTime
}

func (p *parser) readDateStampProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var d dateStamp
	d.dateTime, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return d, nil
}

type lastModified struct {
	dateTime
}

func (p *parser) readLastModifiedProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var l lastModified
	l.dateTime, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return l, nil
}

type sequence int

func (p *parser) readSequenceProperty() (property, error) {
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
