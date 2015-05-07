package ics

import (
	"errors"
	"io"
)

type parser struct {
	l *lexer
}

func newParser(r io.Reader) parser {
	return parser{newLexer(r)}
}

func (p parser) GetComponent() (component, error) {
	t, err := p.l.GetToken()
	if err != nil {
		return nil, err
	} else if t.typ != tokenName {
		return nil, ErrInvalidToken
	}
	c := componentFromToken(t.data)
	t, err = p.l.GetToken()
	if err != nil {
		return nil, err
	}
	for {
		switch t.typ {
		case tokenParamName:
			pn := t
			values := make([]token, 0, 1)
		Loop:
			for {
				t, err = p.l.GetToken()
				if err != nil {
					return nil, err
				}
				switch t.typ {
				case tokenParamValue, tokenParamQValue:
					values = append(values, t)
				default:
					break Loop
				}
			}
			attr, err := attributeFromTokens(pn, values)
			if err != nil {
				return nil, err
			}
			if err = c.setAttribute(attr); err != nil {
				return nil, err
			}
		case tokenValue:
			if err = c.setValue(t.data); err != nil {
				return nil, err
			}
			return c, nil
		default:
			return nil, ErrInvalidToken //should never happen, an error should already have been caught
		}
	}
}

// Errors

var ErrInvalidToken = errors.New("received invalid token")
