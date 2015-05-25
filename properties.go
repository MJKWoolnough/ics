package ics

import (
	"errors"
	"strings"
)

const (
	beginp           = "BEGIN"
	endp             = "END"
	calscalep        = "CALSCALE"
	methodp          = "METHOD"
	prodidp          = "PRODID"
	versionp         = "VERSION"
	attachp          = "ATTACH"
	categoriesp      = "CATEGORIES"
	classp           = "CLASS"
	commentp         = "COMMENT"
	descriptionp     = "DESCRIPTION"
	geop             = "GEO"
	locationp        = "LOCATION"
	percentcompletep = "PERCENT-COMPLETE"
	priorityp        = "PRIORITY"
	resourcesp       = "RESOURCES"
	statusp          = "STATUS"
	summaryp         = "SUMMARY"
	completedp       = "COMPLETED"
	dtendp           = "DTEND"
	duep             = "DUE"
	dtstartp         = "DTSTART"
	durationp        = "DURATION"
	freebusyp        = "FREEBUSY"
	transpp          = "TRANSP"
	tzidp            = "TZID"
	tznamep          = "TZNAME"
	tzoffsetfromp    = "TZOFFSETFROM"
	tzoffsettop      = "TZOFFSETTO"
	tzurlp           = "TZURL"
	attendeep        = "ATTENDEE"
	contactp         = "CONTACT"
	organizerp       = "ORGANIZER"
	recuridp         = "RECURRENCE-ID"
	relatedp         = "RELATED-TO"
	urlp             = "URL"
	uidp             = "UID"
	exdatep          = "EXDATE"
	rdatep           = "RDATE"
	rrulep           = "RRULE"
	actionp          = "ACTION"
	repeatp          = "REPEAT"
	triggerp         = "TRIGGER"
	createdp         = "CREATED"
	dtstampp         = "DTSTAMP"
	lastmodp         = "LAST-MODIFIED"
	seqp             = "SEQUENCE"
	rstatusp         = "REQUEST-STATUS"
)

type property interface{}

type begin string

func (p *parser) readBeginProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return begin(unescape(v)), nil
}

type end string

func (p *parser) readEndProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return end(unescape(v)), nil
}

type requestStatus struct {
	Language          string
	StatusCode        int
	StatusDescription string
	Extra             string
}

func (p *parser) readRequestStatusProperty() (property, error) {
	as, err := p.readAttributes(languageparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	parts := strings.SplitN(v, ";", 3)
	if len(parts) < 2 {
		return nil, ErrUnsupportedValue
	}
	if len(parts[0]) != 4 {
		return nil, ErrUnsupportedValue
	}
	if parts[0][0] < '1' || parts[0][0] > '4' || parts[0][1] != '.' || parts[0][2] < '0' || parts[0][2] > '9' || parts[0][3] < '0' || parts[0][3] > '9' {
		return nil, ErrUnsupportedValue
	}

	r := requestStatus{
		StatusCode:        int(parts[0][0]-'0')*100 + int(parts[0][2]-'0')*10 + int(parts[0][3]-'0'),
		StatusDescription: string(unescape(parts[1])),
	}
	if len(parts) == 3 {
		r.Extra = string(unescape(parts[2]))
	}
	if l, ok := as[languageparam]; ok {
		r.Language = l.String()
	}
	return r, nil
}

type propertyData struct {
	Name   string
	Params map[string]attribute
	Value  string
}

func (p *parser) readUnknownProperty(name string) (property, error) {
	vs, err := p.readAttributes("*")
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return propertyData{
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
