package ics

import (
	"errors"
	"io"

	readerParser "github.com/MJKWoolnough/parser"
)

type parser struct {
	l *readerParser.Parser
	t readerParser.Token
}

func newParser(r io.Reader) *parser {
	return &parser{l: newLexer(r)}
}

func (p *parser) GetProperty() (c property, err error) {
	n, err := p.readName()
	if err != nil {
		return nil, err
	}
	switch n {
	case beginp:
		return p.readBeginProperty()
	case endp:
		return p.readEndProperty()
	case calscalep:
		return p.readCalScaleProperty()
	case methodp:
		return p.readMethodProperty()
	case prodidp:
		return p.readProductIDProperty()
	case versionp:
		return p.readVersionProperty()
	case attachp:
		return p.readAttachProperty()
	case categoriesp:
		return p.readCategoriesProperty()
	case classp:
		return p.readClassProperty()
	case commentp:
		return p.readCommentProperty()
	case descriptionp:
		return p.readDescriptionProperty()
	case geop:
		return p.readGeoProperty()
	case locationp:
		return p.readLocationProperty()
	case percentcompletep:
		return p.readPercentCompleteProperty()
	case priorityp:
		return p.readPriorityProperty()
	case resourcesp:
		return p.readResourcesProperty()
	case statusp:
		return p.readStatusProperty()
	case summaryp:
		return p.readSummaryProperty()
	case completedp:
		return p.readCompletedProperty()
	case dtendp:
		return p.readDateTimeEndProperty()
	case duep:
		return p.readDateTimeDueProperty()
	case dtstartp:
		return p.readDateTimeStartProperty()
	case durationp:
		return p.readDurationProperty()
	case freebusyp:
		return p.readFreeBusyTimeProperty()
	case transpp:
		return p.readTimeTransparencyProperty()
	case tzidp:
		return p.readTimezoneIDProperty()
	case tznamep:
		return p.readTimezoneNameProperty()
	case tzoffsetfromp:
		return p.readTimezoneOffsetFromProperty()
	case tzoffsettop:
		return p.readTimezoneOffsetToProperty()
	case tzurlp:
		return p.readTimezoneURLProperty()
	case attendeep:
		return p.readAttendeeProperty()
	case contactp:
		return p.readContactProperty()
	case organizerp:
		return p.readOrganizerProperty()
	case recuridp:
		return p.readRecurrenceIDProperty()
	case relatedp:
		return p.readRelatedToProperty()
	case urlp:
		return p.readURLProperty()
	case uidp:
		return p.readUIDProperty()
	case exdatep:
		return p.readExceptionDateProperty()
	case rdatep:
		return p.readRecurrenceDateProperty()
	case rrulep:
		return p.readRecurrenceRuleProperty()
	case actionp:
		return p.readActionProperty()
	case repeatp:
		return p.readRepeatProperty()
	case triggerp:
		return p.readTriggerProperty()
	case createdp:
		return p.readCreatedProperty()
	case dtstampp:
		return p.readDateStampProperty()
	case lastmodp:
		return p.readLastModifiedProperty()
	case seqp:
		return p.readSequenceProperty()
	case rstatusp:
		return p.readRequestStatusProperty()
	default:
		return p.readUnknownProperty(p.t.Data)
	}
}

func (p *parser) readName() (string, error) {
	if p.t.Type != tokenName {
		var err error
		p.t, err = p.l.GetToken()
		if err != nil {
			return "", err
		} else if p.t.Type != tokenName {
			return "", ErrInvalidToken
		}
	}
	return p.t.Data, nil

}

func (p *parser) readAttributes(accepted ...string) (as map[string]attribute, err error) {
	as = make(map[string]attribute)
	var (
		a   attribute
		all bool
	)
	if len(accepted) == 1 && accepted[0] == "*" {
		all = true
	}
	for {
		if p.t.Type != tokenParamName {
			p.t, err = p.l.GetToken()
			if err != nil {
				return nil, err
			}
			if p.t.Type == tokenValue {
				return as, nil
			}
		}
		pt := p.t
		vs := make([]readerParser.Token, 0, 1)
		for {
			t, err := p.l.GetToken()
			if err != nil {
				return nil, err
			}
			if t.Type != tokenParamValue && t.Type != tokenParamQValue {
				p.t = t
				break
			}
			vs = append(vs, t)
		}
		switch pt.Data {
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
			if all || pn == pt.Data {
				if _, ok := as[pt.Data]; ok {
					return nil, ErrDuplicateParam
				}
				as[pt.Data] = a
			}
		}
		if p.t.Type != tokenParamName {
			return as, nil
		}
	}
}

func (p *parser) readValue() (v string, err error) {
	if p.t.Type != tokenValue {
		_, err = p.readAttributes()
		if err == nil && p.t.Type != tokenValue {
			err = ErrInvalidToken
		}
	}
	return p.t.Data, err
}

// Errors

var (
	ErrInvalidToken   = errors.New("received invalid token")
	ErrDuplicateParam = errors.New("duplicate parameter")
)
