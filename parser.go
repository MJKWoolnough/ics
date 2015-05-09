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
		case "ALTREP":
			a, err = newAltRepParam(vs)
		case "CN":
			a, err = newCommonNameParam(vs)
		case "CUTYPE":
			a, err = newCalendarUserTypeParam(vs)
		case "DELEGATED-FROM":
			a, err = newDelegatorsParam(vs)
		case "DELEGATED-TO":
			a, err = newDelegateeParam(vs)
		case "DIR":
			a, err = newDirectoryEntryRefParam(vs)
		case "ENCODING":
			a, err = newEncodingParam(vs)
		case "FMTTYPE":
			a, err = newFmtTypeParam(vs)
		case "FBTIME":
			a, err = newFreeBusyTimeParam(vs)
		case "LANGUAGE":
			a, err = newLanguageParam(vs)
		case "MEMBER":
			a, err = newMemberParam(vs)
		case "PARTSTAT":
			a, err = newParticipationStatusParam(vs)
		case "RANGE":
			a, err = newRangeParam(vs)
		case "RELATED":
			a, err = newAlarmTriggerRelationshipParam(vs)
		case "RELTYPE":
			a, err = newRelationshipTypeParam(vs)
		case "ROLE":
			a, err = newParticipationRoleParam(vs)
		case "RSVP":
			a, err = newRSVPExpectationParam(vs)
		case "SENT-BY":
			a, err = newSentByParam(vs)
		case "TZID":
			a, err = newTimezoneIDParam(vs)
		case "VALUE":
			a, err = newValueParam(vs)
		default:
			a, err = newUnknownParam(vs)
		}
		if err != nil {
			return nil, err
		}
		for _, pn := range accepted {
			if pn == p.t.data {
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

var ErrInvalidToken = errors.New("received invalid token")
