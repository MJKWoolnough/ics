package ics

import (
	"strconv"
	"time"
)

type created struct {
	dateTime
}

func (p *parser) readCreatedProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	dateTime, err := parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return created{dateTime}, nil
}

func (c created) Validate() bool {
	return c.Location == time.UTC
}

func (c created) Data() propertyData {
	return propertyData{
		Name:  createdp,
		Value: c.String(),
	}
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

func (d dateStamp) Validate() bool {
	return d.Location == time.UTC
}

func (d dateStamp) Data() propertyData {
	return propertyData{
		Name:  dtstampp,
		Value: d.String(),
	}
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

func (l lastModified) Validate() bool {
	return l.Location == time.UTC
}

func (l lastModified) Data() propertyData {
	return propertyData{
		Name:  lastmodp,
		Value: l.String(),
	}
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

func (s sequence) Validate() bool {
	return s >= 0
}

func (s sequence) Data() propertyData {
	return propertData{
		Name:  seqp,
		Value: strconv.Itoa(int(s)),
	}
}
