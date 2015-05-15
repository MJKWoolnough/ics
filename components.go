package ics

import (
	"errors"
	"strings"
)

const (
	beginc           = "BEGIN"
	endc             = "END"
	calscalec        = "CALSCALE"
	methodc          = "METHOD"
	prodidc          = "PRODID"
	versionc         = "VERSION"
	attachc          = "ATTACH"
	categoriesc      = "CATEGORIES"
	classc           = "CLASS"
	commentc         = "COMMENT"
	descriptionc     = "DESCRIPTION"
	geoc             = "GEO"
	locationc        = "LOCATION"
	percentcompletec = "PERCENT-COMPLETE"
	priorityc        = "PRIORITY"
	resourcesc       = "RESOURCES"
	statusc          = "STATUS"
	summaryc         = "SUMMARY"
	completedc       = "COMPLETED"
	dtendc           = "DTEND"
	duec             = "DUE"
	dtstartc         = "DTSTART"
	durationc        = "DURATION"
	freebusyc        = "FREEBUSY"
	transpc          = "TRANSP"
	tzidc            = "TZID"
	tznamec          = "TZNAME"
	tzoffsetfromc    = "TZOFFSETFROM"
	tzoffsettoc      = "TZOFFSETTO"
	tzurlc           = "TZURL"
	attendeec        = "ATTENDEE"
	contactc         = "CONTACT"
	organizerc       = "ORGANIZER"
	recuridc         = "RECURRENCE-ID"
	relatedc         = "RELATED-TO"
	urlc             = "URL"
	uidc             = "UID"
	exdatec          = "EXDATE"
	rdatec           = "RDATE"
	rrulec           = "RRULE"
	actionc          = "ACTION"
	repeatc          = "REPEAT"
	triggerc         = "TRIGGER"
	createdc         = "CREATED"
	dtstampc         = "DTSTAMP"
	lastmodc         = "LAST-MODIFIED"
	seqc             = "SEQUENCE"
	rstatusc         = "REQUEST-STATUS"
)

type component interface{}

type begin string

func (p *parser) readBeginComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return begin(unescape(v)), nil
}

type end string

func (p *parser) readEndComponent() (component, error) {
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

func (p *parser) readRequestStatusComponent() (component, error) {
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

type unknown struct {
	Name   string
	Params map[string]attribute
	Value  string
}

func (p *parser) readUnknownComponent(name string) (component, error) {
	vs, err := p.readAttributes("*")
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
