package ics

import (
	"encoding/base64"
	"errors"
)

const (
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
)

type component interface{}

type attach struct {
	URI  bool
	Mime string
	Data []byte
}

func (p *parser) readAttachComponent() (component, error) {
	as, err := p.readAttributes(fmttypeparam, encodingparam, valuetypeparam)
	if err != nil {
		return nil, err
	}
	value, err := p.readValue()
	if err != nil {
		return nil, err
	}
	uri := true
	enc, encOK := as[encodingparam]
	val, valOK := as[valuetypeparam]
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
		data = []byte(unescape(value))
	} else {
		return nil, ErrInvalidAttributeCombination
	}
	return attach{
		uri,
		as[fmttypeparam],
		data,
	}, nil
}

type categories struct {
	Language   string
	Categories []string
}

func (p *parser) readCategoriesComponent() (component, error) {
	as, err := p.readAttributes(languageparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	language := ""
	if l, ok := as[languageparam]; ok {
		language = l.String()
	}
	return categories{
		language,
		textSplit(v),
	}, nil
}

const (
	classPublic = iota
	classPrivate
	classConfidential
)

type class struct {
	Value int
}

func (p *parser) readClassComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var cv int
	switch v {
	case "PUBLIC":
		cv = classPublic
	case "PRIVATE":
		cv = classPrivate
	case "CONFIDENTIAL":
		cv = classConfidential
	default:
		cv = classPrivate
	}
	return class{cv}, nil
}

type comment struct {
	Altrep, Language string
}

func (p *parser) readCommentComponent() (component, error) {

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
