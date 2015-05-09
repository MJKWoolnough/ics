package ics

import (
	"encoding/base64"
	"errors"
)

type component interface{}

type attach struct {
	URI  bool
	Mime string
	Data []byte
}

func (p *parser) readAttachComponent() (component, error) {
	as, err := p.readAttributes("FMTTYPE", "ENCODING", "VALUE")
	if err != nil {
		return nil, err
	}
	value, err := p.readValue()
	if err != nil {
		return nil, err
	}
	uri := true
	enc, encOK := as["ENCODING"]
	val, valOK := as["VALUE"]
	var data []byte
	if encOK && valOK {
		uri = false
		if enc.(encoding) != encodingBase64 || val.(value) != valueBinary {
			return nil, ErrUnsupportedValue
		}
		data, err = base64.StdEncoding.DecodeString(value)
		if err != nil {
			return nil, err
		}

	} else if encOK == valOK {
		data = []byte(value)
	} else {
		return nil, ErrInvalidAttributeCombination
	}
	return attach{
		uri,
		as["FMTTYPE"],
		data,
	}, nil
}

type unknown struct {
	Name   string
	Params []token
	Value  string
}

func (p *parser) readUnknownComponent(name string) (component, error) {
	vs, err := p.readAttributes()
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return unknown{
		name,
		vs,
		v,
	}, err
}

// Errors

var (
	ErrUnsupportedValue            = errors.New("attribute contained unsupported value")
	ErrInvalidAttributeCombination = errors.New("invalid combination of attributes")
)
