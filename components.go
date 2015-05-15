package ics

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

	strparse "github.com/MJKWoolnough/parser"
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

type completed struct {
	time.Time
}

func (p *parser) readCompletedComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var c completed
	c.Time, err = time.ParseInLocation("20060102T150405Z", v, time.UTC)
	if err != nil {
		return nil, err
	}
	return c, nil
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

type duration struct {
	time.Duration
}

func (p *parser) readDurationComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var d duration
	d.Duration, err = parseDuration(v)
	if err != nil {
		return nil, err
	}
	return d, nil
}

type freeBusyTime struct {
	Typ     freeBusy
	Periods []period
}

type period struct {
	FixedDuration bool
	Start, End    time.Time
}

func parsePeriods(v string, l *time.Location) ([]period, error) {
	periods := make([]period, 0, 1)

	for _, pd := range textSplit(v, ',') {
		parts := strings.Split(pd, "/")
		if len(parts) != 2 {
			return nil, ErrUnsupportedValue
		}
		if parts[0][len(parts[0])-1] != 'Z' {
			return nil, ErrUnsupportedValue
		}
		start, err := parseDateTime(parts[0], l)
		if err != nil {
			return nil, err
		}
		var (
			end           time.Time
			fixedDuration bool
		)
		if parts[1][len(parts[1])-1] == 'Z' {
			end, err = parseDateTime(parts[1], l)
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
	return periods, nil
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
	periods, err := parsePeriods(v, nil)
	if err != nil {
		return nil, err
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

func (p *parser) readTimezoneIDComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneID(v), nil
}

type timezoneName struct {
	Language, Name string
}

func (p *parser) readTimezoneNameComponent() (component, error) {
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

type timezoneOffsetFrom int

func (p *parser) readTimezoneOffsetFromComponent() (component, error) {
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

type timezoneOffsetTo struct {
	time.Duration
}

func (p *parser) readTimezoneOffsetToComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var tzo timezoneOffsetTo
	tzo.Duration, err = parseOffset(v)
	if err != nil {
		return nil, err
	}
	return timezoneOffsetTo(tzo), nil
}

type timezoneURL string

func (p *parser) readTimezoneURLComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneURL(v), nil
}

type attendee struct {
	CalendarUserType    calendarUserType
	Members             members
	Role                participationRole
	ParticipationStatus participationStatus
	RSVP                rsvp
	Delegatee           delegatee
	Delegator           delegators
	SentBy              sentBy
	CommonName          commonName
	DirectoryEntryRef   directoryEntryRef
	Language            language
	Address             string
}

func (p *parser) readAttendeeComponent() (component, error) {
	as, err := p.readAttributes(cutypeparam, memberparam, roleparam, partstatparam, rsvpparam, deltoparam, delfromparam, sentbyparam, cnparam, dirparam, languageparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var a attendee
	a.Address = string(unescape(v))
	if pm, ok := as[cutypeparam]; ok {
		a.CalendarUserType = pm.(calendarUserType)
	}
	if pm, ok := as[memberparam]; ok {
		a.Members = pm.(members)
	}
	if pm, ok := as[roleparam]; ok {
		a.Role = pm.(participationRole)
	}
	if pm, ok := as[partstatparam]; ok {
		a.ParticipationStatus = pm.(participationStatus)
	}
	if pm, ok := as[rsvpparam]; ok {
		a.RSVP = pm.(rsvp)
	}
	if pm, ok := as[deltoparam]; ok {
		a.Delegatee = pm.(delegatee)
	}
	if pm, ok := as[delfromparam]; ok {
		a.Delegator = pm.(delegators)
	}
	if pm, ok := as[sentbyparam]; ok {
		a.SentBy = pm.(sentBy)
	}
	if pm, ok := as[cnparam]; ok {
		a.CommonName = pm.(commonName)
	}
	if pm, ok := as[dirparam]; ok {
		a.DirectoryEntryRef = pm.(directoryEntryRef)
	}
	if pm, ok := as[languageparam]; ok {
		a.Language = pm.(language)
	}
	return a, nil
}

type contact struct {
	Altrep, Language, Value string
}

func (p *parser) readContactComponent() (component, error) {
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
	return contact{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type organizer struct {
	CommonName        commonName
	DirectoryEntryRef directoryEntryRef
	SentBy            sentBy
	Language          language
	Name              string
}

func (p *parser) readOrganizerComponent() (component, error) {
	as, err := p.readAttributes(cnparam, dirparam, sentbyparam, languageparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var o organizer
	o.Name = string(unescape(v))
	if pm, ok := as[cnparam]; ok {
		o.CommonName = pm.(commonName)
	}
	if pm, ok := as[dirparam]; ok {
		o.DirectoryEntryRef = pm.(directoryEntryRef)
	}
	if pm, ok := as[sentbyparam]; ok {
		o.SentBy = pm.(sentBy)
	}
	if pm, ok := as[languageparam]; ok {
		o.Language = pm.(language)
	}
	return o, nil
}

type recurrenceID struct {
	Range    rangeParam
	JustDate bool
	DateTime time.Time
}

func (p *parser) readRecurrenceIDComponent() (component, error) {
	as, err := p.readAttributes(valuetypeparam, tzidparam, rangeparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var (
		r recurrenceID
		l *time.Location
	)
	if tzid, ok := as[tzidparam]; ok {
		l, err = time.LoadLocation(tzid.String())
		if err != nil {
			return nil, err
		}
	}
	if val, ok := as[valuetypeparam]; ok && val.(value) == valueDate {
		r.JustDate = true
		r.DateTime, err = parseDate(v)
	} else {
		r.DateTime, err = parseDateTime(v, l)
	}
	if err != nil {
		return nil, err
	}
	if rng, ok := as[rangeparam]; ok {
		r.Range = rng.(rangeParam)
	}
	return r, nil
}

type relatedTo struct {
	RelationshipType relationshipType
	Value            string
}

func (p *parser) readRelatedToComponent() (component, error) {
	as, err := p.readAttributes(reltypeparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	r := relatedTo{Value: v}
	if rel, ok := as[reltypeparam]; ok {
		r.RelationshipType = rel.(relationshipType)
	}
	return r, nil
}

type url string

func (p *parser) readURLComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return url(v), nil
}

type uid string

func (p *parser) readUIDComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return uid(v), nil
}

type exceptionDate struct {
	JustDate bool
	DateTime time.Time
}

func (p *parser) readExceptionDateComponent() (component, error) {
	as, err := p.readAttributes(tzidparam, valuetypeparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var (
		e exceptionDate
		l *time.Location
	)
	if tzid, ok := as[tzidparam]; ok {
		l, err = time.LoadLocation(tzid.String())
		if err != nil {
			return nil, err
		}
	}
	if val, ok := as[valuetypeparam]; ok && val.(value) == valueDate {
		e.JustDate = true
		e.DateTime, err = parseDate(v)
	} else {
		e.DateTime, err = parseDateTime(v, l)
	}
	if err != nil {
		return nil, err
	}
	return e, nil
}

type recurrenceDate struct {
	JustDate bool
	Periods  []period
}

func (p *parser) readRecurrenceDateComponent() (component, error) {
	as, err := p.readAttributes(tzidparam, valuetypeparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var (
		r recurrenceDate
		l *time.Location
	)
	if tzid, ok := as[tzidparam]; ok {
		l, err = time.LoadLocation(tzid.String())
		if err != nil {
			return nil, err
		}
	}
	if val, ok := as[valuetypeparam]; ok && val.(value) == valuePeriod {
		r.Periods, err = parsePeriods(v, l)
		if err != nil {
			return nil, err
		}
	} else if ok && val.(value) == valueDate {
		r.JustDate = true
		parts := textSplit(v, ',')
		r.Periods = make([]period, 0, len(parts))
		for _, tm := range parts {
			t, err := parseDate(tm)
			if err != nil {
				return nil, err
			}
			r.Periods = append(r.Periods, period{Start: t})
		}
	} else {
		parts := textSplit(v, ',')
		r.Periods = make([]period, 0, len(parts))
		for _, tm := range parts {
			t, err := parseDateTime(tm, l)
			if err != nil {
				return nil, err
			}
			r.Periods = append(r.Periods, period{Start: t})
		}
	}
	return r, nil
}

const (
	freqSecondly frequency = iota + 1
	freqMinutely
	freqHourly
	freqDaily
	freqWeekly
	freqMonthly
	freqYearly
)

type frequency int

type recurrenceRule struct {
	Frequency                                                                      frequency
	Until                                                                          time.Time
	Count, Interval                                                                int
	BySecond, ByMinute, ByHour, ByMonthDay, ByYearDay, ByWeekNo, ByMonth, BySetPos []int
	ByDay                                                                          [][2]int
	WeekStart                                                                      int
}

func (p *parser) readRecurrenceRuleComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var (
		r                                                                                                             recurrenceRule
		freqSet, untilCountSet, intervalSet, bsSet, bmSet, bhSet, bdSet, bmdSet, bydSet, bwnSet, bmoSet, bstSet, wSet bool
	)
	for _, rule := range strings.Split(v, ";") {
		parts := strings.SplitN(rule, "=", 2)
		if len(parts) != 2 {
			return nil, ErrUnsupportedValue
		}
		if !freqSet && parts[0] != "FREQ" {
			return nil, ErrUnsupportedValue
		}
		switch parts[0] {
		case "FREQ":
			if freqSet {
				return nil, ErrInvalidAttributeCombination
			}
			freqSet = true
			switch parts[1] {
			case "SECONDLY":
				r.Frequency = freqSecondly
			case "MINUTELY":
				r.Frequency = freqMinutely
			case "HOURLY":
				r.Frequency = freqHourly
			case "DAILY":
				r.Frequency = freqDaily
			case "WEEKLY":
				r.Frequency = freqWeekly
			case "MONTHLY":
				r.Frequency = freqMonthly
			case "YEARLY":
				r.Frequency = freqYearly
			default:
				return nil, ErrUnsupportedValue
			}
		case "UNTIL":
			if untilCountSet {
				return nil, ErrInvalidAttributeCombination
			}
			untilCountSet = true
			if strings.IndexByte(parts[1], 'T') >= 0 {
				r.Until, err = parseDateTime(parts[1], nil)
			} else {
				r.Until, err = parseDate(parts[1])
			}
			if err != nil {
				return nil, err
			}
		case "COUNT":
			if untilCountSet {
				return nil, ErrInvalidAttributeCombination
			}
			untilCountSet = true
			r.Count, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
		case "INTERVAL":
			if intervalSet {
				return nil, ErrInvalidAttributeCombination
			}
			intervalSet = true
			r.Interval, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
			if r.Interval <= 0 {
				return nil, ErrUnsupportedValue
			}
		case "BYSECOND":
			if bsSet {
				return nil, ErrInvalidAttributeCombination
			}
			bsSet = true
			for _, sec := range strings.Split(parts[1], ",") {
				s, err := strconv.Atoi(sec)
				if err != nil {
					return nil, err
				}
				if s < 0 || s > 60 {
					return nil, ErrUnsupportedValue
				}
				r.BySecond = append(r.BySecond, s)
			}
		case "BYMINUTE":
			if bmSet {
				return nil, ErrInvalidAttributeCombination
			}
			bmSet = true
			for _, min := range strings.Split(parts[1], ",") {
				m, err := strconv.Atoi(min)
				if err != nil {
					return nil, err
				}
				if m < 0 || m > 59 {
					return nil, ErrUnsupportedValue
				}
				r.ByMinute = append(r.ByMinute, m)
			}
		case "BYHOUR":
			if bhSet {
				return nil, ErrInvalidAttributeCombination
			}
			bhSet = true
			for _, hour := range strings.Split(parts[1], ",") {
				h, err := strconv.Atoi(hour)
				if err != nil {
					return nil, err
				}
				if h < 0 || h > 23 {
					return nil, ErrUnsupportedValue
				}
				r.ByHour = append(r.ByHour, h)
			}
		case "BYDAY":
			if bdSet {
				return nil, ErrInvalidAttributeCombination
			}
			bdSet = true
			for _, day := range strings.Split(parts[1], ",") {
				p := strparse.NewStringParser(day)
				var (
					neg, num bool
					n, w     int
				)
				if p.Accept("-") {
					num = true
					neg = true
				} else if p.Accept("+") {
					num = true
				}
				pos := len(p.Get())
				p.AcceptRun("0123456789")
				if p.Len() == 0 {
					if num {
						return nil, ErrUnsupportedValue
					}
				} else {
					numStr := p.Get()
					pos += len(numStr)
					n, _ = strconv.Atoi(numStr)
					if neg {
						n = -n
					}
					if n < -53 || n > 53 || n == 0 {
						return nil, ErrUnsupportedValue
					}
				}
				switch parts[1][pos:] {
				case "SU":
				case "MO":
				case "TU":
				case "WE":
				case "TH":
				case "FR":
				case "SA":
				default:
					return nil, ErrUnsupportedValue
				}
				r.ByDay = append(r.ByDay, [2]int{n, w})
			}
		case "BYMONTHDAY":
			if bmdSet {
				return nil, ErrInvalidAttributeCombination
			}
			bmdSet = true
			for _, monthday := range strings.Split(parts[1], ",") {
				md, err := strconv.Atoi(monthday)
				if err != nil {
					return nil, err
				}
				if md < -31 || md > 31 || md == 0 {
					return nil, ErrUnsupportedValue
				}
				r.ByMonthDay = append(r.ByMonthDay, md)
			}
		case "BYYEARDAY":
			if bydSet {
				return nil, ErrInvalidAttributeCombination
			}
			bydSet = true
			for _, yearday := range strings.Split(parts[1], ",") {
				yd, err := strconv.Atoi(yearday)
				if err != nil {
					return nil, err
				}
				if yd < -366 || yd > 366 || yd == 0 {
					return nil, ErrUnsupportedValue
				}
				r.ByYearDay = append(r.ByYearDay, yd)
			}
		case "BYWEEKNO":
			if bwnSet {
				return nil, ErrInvalidAttributeCombination
			}
			bwnSet = true
			for _, week := range strings.Split(parts[1], ",") {
				w, err := strconv.Atoi(week)
				if err != nil {
					return nil, err
				}
				if w < -53 || w > 53 || w == 0 {
					return nil, ErrUnsupportedValue
				}
				r.ByWeekNo = append(r.ByWeekNo, w)
			}
		case "BYMONTH":
			if bmoSet {
				return nil, ErrInvalidAttributeCombination
			}
			bmoSet = true
			for _, month := range strings.Split(parts[1], ",") {
				m, err := strconv.Atoi(month)
				if err != nil {
					return nil, err
				}
				if m < 1 || m > 12 {
					return nil, ErrUnsupportedValue
				}
				r.ByMonth = append(r.ByMonth, m)
			}
		case "BYSETPOS":
			if bstSet {
				return nil, ErrInvalidAttributeCombination
			}
			bstSet = true
			for _, setpos := range strings.Split(parts[1], ",") {
				sp, err := strconv.Atoi(setpos)
				if err != nil {
					return nil, err
				}
				if sp < -366 || sp > 366 || sp == 0 {
					return nil, ErrUnsupportedValue
				}
				r.BySetPos = append(r.BySetPos, sp)
			}
		case "WKST":
			if wSet {
				return nil, ErrInvalidAttributeCombination
			}
			wSet = true
			r.WeekStart, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
		default:
		}
	}
	if !intervalSet {
		r.Interval = 1
	}
	if !freqSet {
		return nil, ErrUnsupportedValue
	}
	return r, nil
}

const (
	actionUnknown action = iota
	actionAudio
	actionDisplay
	actionEmail
)

type action int

func (p *parser) readActionComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "AUDIO":
		return actionAudio, nil
	case "DISPLAY":
		return actionDisplay, nil
	case "EMAIL":
		return actionEmail, nil
	default:
		return actionUnknown, nil
	}
}

type repeat int

func (p *parser) readRepeatComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	return repeat(n), nil
}

type related int

type trigger struct {
	DateTime time.Time
	Related  alarmTriggerRelationship
	Duration time.Duration
}

func (p *parser) readTriggerComponent() (component, error) {
	as, err := p.readAttributes(valuetypeparam, trigrelparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var t trigger
	if val, ok := as[valuetypeparam]; ok && val.(value) == valueDateTime {
		if v[len(v)-1] != 'Z' {
			return nil, ErrUnsupportedValue
		}
		t.DateTime, err = parseDateTime(v, nil)
		if err != nil {
			return nil, err
		}
		t.Related = -1
	} else {
		if rel, ok := as[trigrelparam]; ok {
			t.Related = rel.(alarmTriggerRelationship)
		} else {
			t.Related = atrStart
		}
		t.Duration, err = parseDuration(v)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

type created struct {
	time.Time
}

func (p *parser) readCreateComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var c created
	c.Time, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type dateStamp struct {
	time.Time
}

func (p *parser) readDateStampComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var d dateStamp
	d.Time, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return d, nil
}

type lastModified struct {
	time.Time
}

func (p *parser) readLastModifiedComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	if v[len(v)-1] != 'Z' {
		return nil, ErrUnsupportedValue
	}
	var l lastModified
	l.Time, err = parseDateTime(v, nil)
	if err != nil {
		return nil, err
	}
	return l, nil
}

type sequence int

func (p *parser) readSequenceComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	s, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	if s < 0 {
		return nil, ErrUnsupportedValue
	}
	return sequence(s), nil
}

type requestStatus struct {
	Language, Status string
}

func (p *parser) readRequestStatusComponent() (component, error) {
	as, err := p.readAttributes(languageparam)
	if err != nil {
		return nil, err
	}
	var r requestStatus
	r.Status, err = p.readValue()
	if err != nil {
		return nil, err
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
