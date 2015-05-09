package ics

import (
	"errors"
	"io"
)

type parser struct {
	l *lexer
	t token
}

func newParser(r io.Reader) *parser {
	return &parser{newLexer(r)}
}

func (p *parser) GetComponent() (c component, err error) {
	p.t, err = p.l.GetToken()
	if err != nil {
		return nil, err
	} else if t.typ != tokenName {
		return nil, ErrInvalidToken
	}
	switch p.t.data {
	case "ATTACH":
		return p.readAttachComponent()
	default:
		return p.readUnknownComponent(p.t.data)
	}
}

func (p *parser) readAttributes(accepted ...string) (as map[string]attribute, err error) {
	as = make(map[string]attribute)
	var a attribute
	for {
		if p.t.typ != tokenParamName {
			p.t, err = p.l.GetToken()
			if err != nil {
				return nil, err
			}
			if p.t.typ == tokenParamValue {
				return as, nil
			}
		}
		vs := make([]token, 0, 1)
		for {
			t, err := p.l.GetToken()
			if err != nil {
				return nil, err
			}
			if t.typ != tokenParamValue {
				p.t = t
				return vs, nil
			}
			vs = append(vs, t)
		}
		switch p.t.data {
		case altrepparam:
			a, err = newAltRepParam(vs)
		case cnparam:
			a, err = newCommonNameParam(vs)
		case cutypeparam:
			a, err = newCalendarUserTypeParam(vs)
		case delfromparam:
			a, err = newDelegatorsParam(vs)
		case deltoparam:
			a, err = newDelegateeParam(vs)
		case dirparam:
			a, err = newDirectoryEntryRefParam(vs)
		case encodingparam:
			a, err = newEncodingParam(vs)
		case fmttypeparam:
			a, err = newFmtTypeParam(vs)
		case fbtypeparam:
			a, err = newFreeBusyTimeParam(vs)
		case languageparam:
			a, err = newLanguageParam(vs)
		case memberparam:
			a, err = newMemberParam(vs)
		case partstatparam:
			a, err = newParticipationStatusParam(vs)
		case rangeparam:
			a, err = newRangeParam(vs)
		case trigrelparam:
			a, err = newAlarmTriggerRelationshipParam(vs)
		case reltypeparam:
			a, err = newRelationshipTypeParam(vs)
		case roleparam:
			a, err = newParticipationRoleParam(vs)
		case rsvpparam:
			a, err = newRSVPExpectationParam(vs)
		case sentbyparam:
			a, err = newSentByParam(vs)
		case tzidparam:
			a, err = newTimezoneIDParam(vs)
		case valuetypeparam:
			a, err = newValueParam(vs)
		default:
			a, err = newUnknownParam(vs)
		}
		if err != nil {
			return nil, err
		}
		for _, pn := range accepted {
			if pn == p.t.data {
				if _, ok := as[pn]; ok {
					return nil, ErrDuplicateParam
				}
				as[pn] = a
			}
		}
		if p.t.typ == tokenParamValue {
			return as, nil
		}
	}
}

func (p *parser) readValue() (v string, err error) {
	if p.t.typ != tokenValue {
		_, err = p.readAttributes()
		if err != nil || a == nil {
			break
		}
		if err == nil && p.t.typ != tokenValue {
			err = ErrInvalidToken
		}
	}
	return p.t.data, err
}

// Errors

var (
	ErrInvalidToken   = errors.New("received invalid token")
	ErrDuplicateParam = errors.New("duplicate parameter")
)
