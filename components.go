package ics

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"
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

type calscale string

func (p *parser) readCalScaleComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return calscale(unescape(v)), nil
}

type method string

func (p *parser) readMethodComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return method(unescape(v)), nil
}

type productID string

func (p *parser) readProductIDComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return productID(unescape(v)), nil
}

type version struct {
	Min, Max string
}

func (p *parser) readVersionComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	parts := textSplit(v, ';')
	if len(parts) > 2 {
		return nil, ErrUnsupportedValue
	} else if len(parts) == 2 {
		return version{parts[0], parts[1]}, nil
	} else {
		return version{parts[0], parts[0]}, nil
	}
}

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
	v, err := p.readValue()
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
		data, err = base64.StdEncoding.DecodeString(v)
		if err != nil {
			return nil, err
		}

	} else if encOK == valOK {
		data = []byte(unescape(v))
	} else {
		return nil, ErrInvalidAttributeCombination
	}
	return attach{
		uri,
		string(as[fmttypeparam].(fmtType)),
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
		textSplit(v, ','),
	}, nil
}

const (
	classPublic class = iota
	classPrivate
	classConfidential
)

type class int

func (p *parser) readClassComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "PUBLIC":
		return classPublic, nil
	case "PRIVATE":
		return classPrivate, nil
	case "CONFIDENTIAL":
		return classConfidential, nil
	default:
		return classPrivate, nil
	}
}

type comment struct {
	Altrep, Language, Comment string
}

func (p *parser) readCommentComponent() (component, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return comment{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type description struct {
	Altrep, Language, Description string
}

func (p *parser) readDescriptionComponent() (component, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return description{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type geo struct {
	Latitude, Longitude float64
}

func (p *parser) readGeoComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	parts := textSplit(v, ';')
	if len(parts) != 2 {
		return nil, ErrUnsupportedValue
	}
	la, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return nil, err
	}
	lo, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return nil, err
	}
	return geo{la, lo}, nil
}

type location struct {
	Altrep, Language, Location string
}

func (p *parser) readLocationComponent() (component, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return location{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type percentComplete int

func (p *parser) readPercentCompleteComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	pc, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	if pc < 0 || pc > 100 {
		return nil, ErrUnsupportedValue
	}
	return percentComplete(pc), nil
}

type priority int

func (p *parser) readPriorityComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	pc, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	if pc < 0 || pc > 9 {
		return nil, ErrUnsupportedValue
	}
	return priority(pc), nil
}

type resources struct {
	Altrep, Language string
	Resources        []string
}

func (p *parser) readResourcesComponent() (component, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return resources{
		altRep,
		languageStr,
		textSplit(v, ','),
	}, nil
}

const (
	statusTentative status = iota
	statusConfirmed
	statusNeedsAction
	statusCompleted
	statusInProgress
	statusDraft
	statusFinal
	statusCancelled
)

type status int

func (p *parser) readStatusComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "TENTATIVE":
		return statusTentative, nil
	case "CONFIRMED":
		return statusConfirmed, nil
	case "NEED-ACTION":
		return statusNeedsAction, nil
	case "COMPLETED":
		return statusCompleted, nil
	case "IN-PROGRESS":
		return statusInProgress, nil
	case "DRAFT":
		return statusDraft, nil
	case "FINAL":
		return statusFinal, nil
	case "CANCELLED":
		return statusCancelled, nil
	default:
		return nil, ErrUnsupportedValue
	}
}

type summary struct {
	Altrep, Language, Summary string
}

func (p *parser) readSummaryComponent() (component, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return summary{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type completed time.Time

func (p *parser) readCompletedComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	t, err := time.ParseInLocation("20060102T150405Z", v, time.UTC)
	if err != nil {
		return nil, err
	}
	return completed(t), nil
}

type dateTimeEnd struct {
	justDate bool
	Time     time.Time
}

func (p *parser) readDateTimeOrTime() (t time.Time, justDate bool, err error) {
	as, err := p.readAttributes(tzidparam, valuetypeparam)
	if err != nil {
		return t, justDate, err
	}
	var (
		l *time.Location
	)
	if tzid, ok := as[tzidparam]; ok {
		l, err = time.LoadLocation(string(tzid.(timezoneID)))
		if err != nil {
			return t, justDate, err
		}
	}
	if v, ok := as[valuetypeparam]; ok {
		val := v.(value)
		switch val {
		case valueDate:
			justDate = true
		case valueDateTime:
			justDate = false
		default:
			return t, justDate, ErrUnsupportedValue
		}
	}
	v, err := p.readValue()
	if err != nil {
		return t, justDate, err
	}
	if justDate {
		t, err = parseDate(v)
	} else {
		t, err = parseDateTime(v, l)
	}
	return t, justDate, err
}

func (p *parser) readDateTimeEndComponent() (component, error) {
	t, j, err := p.readDateTimeOrTime()
	if err != nil {
		return nil, err
	}
	return dateTimeEnd{j, t}, nil
}

type dateTimeDue struct {
	justDate bool
	Time     time.Time
}

func (p *parser) readDateTimeDueComponent() (component, error) {
	t, j, err := p.readDateTimeOrTime()
	if err != nil {
		return nil, err
	}
	return dateTimeDue{j, t}, nil
}

type dateTimeStart struct {
	justDate bool
	Time     time.Time
}

func (p *parser) readDateTimeStartComponent() (component, error) {
	t, j, err := p.readDateTimeOrTime()
	if err != nil {
		return nil, err
	}
	return dateTimeStart{j, t}, nil
}

type duration time.Duration

func (p *parser) readDurationComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	d, err := parseDuration(v)
	if err != nil {
		return nil, err
	}
	return duration(d), nil
}

type freeBusyTime struct {
	Typ     freeBusy
	Periods []period
}

type period struct {
	FixedDuration bool
	Start, End    time.Time
}

func (p *parser) readFreeBusyTimeComponent() (component, error) {
	as, err := p.readAttributes(fbtypeparam)
	if err != nil {
		return nil, err
	}
	var fb freeBusy
	if f, ok := as[fbtypeparam]; ok {
		fb = f.(freeBusy)
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	periods := make([]period, 0, 1)

	for _, pd := range textSplit(v, ',') {
		parts := strings.Split(pd, "/")
		if len(parts) != 2 {
			return nil, ErrUnsupportedValue
		}
		if parts[0][len(parts[0])-1] != 'Z' {
			return nil, ErrUnsupportedValue
		}
		start, err := parseDateTime(parts[0], nil)
		if err != nil {
			return nil, err
		}
		var (
			end           time.Time
			fixedDuration bool
		)
		if parts[1][len(parts[1])-1] == 'Z' {
			end, err = parseDateTime(parts[1], nil)
			if err != nil {
				return nil, err
			}
		} else {
			d, err := parseDuration(parts[1])
			if err != nil {
				return nil, err
			}
			if d < 0 {
				return nil, ErrUnsupportedValue
			}
			end = start.Add(d)
			fixedDuration = true
		}
		periods = append(periods, period{fixedDuration, start, end})
	}
	return freeBusyTime{
		Typ:     fb,
		Periods: periods,
	}, nil
}

const (
	TTOpaque timeTransparency = iota
	TTTransparent
)

type timeTransparency int

func (p *parser) readTimeTransparencyComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "OPAQUE":
		return TTOpaque, nil
	case "TRANSPARENT":
		return TTTransparent, nil
	default:
		return nil, ErrUnsupportedValue
	}
}

//type timeZoneID string //(in attributes)

func (p *parser) readTimezoneID() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneID(v), nil
}

type timezoneName struct {
	Language, Name string
}

func (p *parser) readTimezoneName() (component, error) {
	as, err := p.readAttributes(languageparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var languageStr string
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	return timezoneName{languageStr, v}, nil
}

type timezoneOffsetFrom nil

func (p *parser) readTimezoneOffsetFrom() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	tzo, err := parseOffset(v)
	if err != nil {
		return nil, err
	}
	return timezoneOffsetFrom(tzo), nil
}

type timezoneOffsetTo time.Duration

func (p *parser) readTimezoneOffsetTo() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	tzo, err := parseOffset(v)
	if err != nil {
		return nil, err
	}
	return timezoneOffsetTo(tzo), nil
}

type timezoneURL string

func (p *parser) readTimezoneURL() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneURL(v), nil
}

type unknown struct {
	Name   string
	Params map[string]attribute
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
