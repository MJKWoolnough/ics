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
	return &parser{l: newLexer(r)}
}

func (p *parser) GetComponent() (c component, err error) {
	n, err := p.readName()
	if err != nil {
		return nil, err
	}
	switch n {
	case beginc:
		return p.readBeginComponent()
	case endc:
		return p.readEndComponent()
	case calscalec:
		return p.readCalScaleComponent()
	case methodc:
		return p.readMethodComponent()
	case prodidc:
		return p.readProductIDComponent()
	case versionc:
		return p.readVersionComponent()
	case attachc:
		return p.readAttachComponent()
	case categoriesc:
		return p.readCategoriesComponent()
	case classc:
		return p.readClassComponent()
	case commentc:
		return p.readCommentComponent()
	case descriptionc:
		return p.readDescriptionComponent()
	case geoc:
		return p.readGeoComponent()
	case locationc:
		return p.readLocationComponent()
	case percentcompletec:
		return p.readPercentCompleteComponent()
	case priorityc:
		return p.readPriorityComponent()
	case resourcesc:
		return p.readResourcesComponent()
	case statusc:
		return p.readStatusComponent()
	case summaryc:
		return p.readSummaryComponent()
	case completedc:
		return p.readCompletedComponent()
	case dtendc:
		return p.readDateTimeEndComponent()
	case duec:
		return p.readDateTimeDueComponent()
	case dtstartc:
		return p.readDateTimeStartComponent()
	case durationc:
		return p.readDurationComponent()
	case freebusyc:
		return p.readFreeBusyTimeComponent()
	case transpc:
		return p.readTimeTransparencyComponent()
	case tzidc:
		return p.readTimezoneIDComponent()
	case tznamec:
		return p.readTimezoneNameComponent()
	case tzoffsetfromc:
		return p.readTimezoneOffsetFromComponent()
	case tzoffsettoc:
		return p.readTimezoneOffsetToComponent()
	case tzurlc:
		return p.readTimezoneURLComponent()
	case attendeec:
		return p.readAttendeeComponent()
	case contactc:
		return p.readContactComponent()
	case organizerc:
		return p.readOrganizerComponent()
	case recuridc:
		return p.readRecurrenceIDComponent()
	case relatedc:
		return p.readRelatedToComponent()
	case urlc:
		return p.readURLComponent()
	case uidc:
		return p.readUIDComponent()
	case exdatec:
		return p.readExceptionDateComponent()
	case rdatec:
		return p.readRecurrenceDateComponent()
	case rrulec:
		return p.readRecurrenceRuleComponent()
	case actionc:
		return p.readActionComponent()
	case repeatc:
		return p.readRepeatComponent()
	case triggerc:
		return p.readTriggerComponent()
	case createdc:
		return p.readCreateComponent()
	case dtstampc:
		return p.readDateStampComponent()
	case lastmodc:
		return p.readLastModifiedComponent()
	case seqc:
		return p.readSequenceComponent()
	case rstatusc:
		return p.readRequestStatusComponent()
	default:
		return p.readUnknownComponent(p.t.data)
	}
}

func (p *parser) readName() (string, error) {
	if p.t.typ != tokenName {
		var err error
		p.t, err = p.l.GetToken()
		if err != nil {
			return "", err
		} else if p.t.typ != tokenName {
			return "", ErrInvalidToken
		}
	}
	return p.t.data, nil

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
		pt := p.t
		vs := make([]token, 0, 1)
		for {
			t, err := p.l.GetToken()
			if err != nil {
				return nil, err
			}
			if t.typ != tokenParamValue && t.typ != tokenParamQValue {
				p.t = t
				break
			}
			vs = append(vs, t)
		}
		switch pt.data {
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
			a, err = newFreeBusyParam(vs)
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
			if pn == pt.data {
				if _, ok := as[pn]; ok {
					return nil, ErrDuplicateParam
				}
				as[pn] = a
			}
		}
		if p.t.typ != tokenParamName {
			return as, nil
		}
	}
}

func (p *parser) readValue() (v string, err error) {
	if p.t.typ != tokenValue {
		_, err = p.readAttributes()
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
