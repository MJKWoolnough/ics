package ics

import "strings"

type component interface {
	setAttribute(token, []token) error
	setValue(token) error
}

func componentFromToken(t token) component {
	switch strings.ToUpper(t.data) {
	default:
		return &UnknownComponent{
			Name: t.data,
		}
	}
}

type UnknownComponent struct {
	Name, Value string
	Attributes  []Attribute
}

func (u *UnknownComponent) setAttribute(pn token, pvs []token) error {
	attr, err := attributeFromTokens(pn, pvs)
	if err != nil {
		return err
	}
	u.Attributes = append(u.Attributes, attr)
	return nil
}

func (u *UnknownComponent) setValue(v token) error {
	u.Value = v.data
	return nil
}
