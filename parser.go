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

func (p parser) GetComponent(r io.Reader) (component, error) {
	t, err := p.l.GetToken()
	if err != nil {
		return nil, err
	} else if t.typ != TokenName {
		return nil, ErrInvalidToken
	}
	c := componentFromToken(t)
	t, err = p.l.GetToken()
	if err != nil {
		return nil, err
	}
	for {
		switch t.typ {
		case TokenParamName:
			pn := t
			values := make([]token, 0, 1)
		Loop:
			for {
				t, err = p.l.GetToken()
				if err != nil {
					return nil, err
				}
				switch t.typ {
				case TokenParamValue, TokenParamQValue:
					values = append(values, t)
				default:
					break Loop
				}
			}
			if err := c.setAttribute(pn, values); err != nil {
				return nil, err
			}
		case TokenValue:
			if err = c.setValue(t); err != nil {
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
