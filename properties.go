package ics

import (
	"bytes"
	"errors"
	"strconv"
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

type property interface {
	Validate() bool
	Data() propertyData
}

type begin string

func (p *parser) readBeginProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return begin(unescape(v)), nil
}

func (b begin) Validate() bool {
	switch b {
	case vCalendar, vEvent, vAlarm, vTodo, vJournal, vTimezone:
		return true
	}
	return false
}

func (b begin) Data() propertyData {
	return propertyData{
		Name:  beginp,
		Value: string(b),
	}
}

type end string

func (p *parser) readEndProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return end(unescape(v)), nil
}

func (e end) Validate() bool {
	switch e {
	case vCalendar, vEvent, vAlarm, vTodo, vJournal, vTimezone:
		return true
	}
	return false
}

func (e end) Data() propertyData {
	return propertyData{
		Name:  endp,
		Value: string(b),
	}
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

func (r requestStatus) Validate() bool {
	if r.StatusCode < 100 || r.StatusCode > 499 {
		return false
	}
	return true
}

func (r requestStatus) Data() propertyData {
	params := make(map[string]attribute)
	if r.Language != "" {
		params[languageparam] = language(r.Language)
	}
	parts := make([]string, 0, 3)
	parts = append(parts, strconv.Itoa(r.StatusCode), r.StatusDescription)
	if r.Extra != "" {
		parts = append(parts, r.Extra)
	}
	val := strings.Join(parts, ";")
	return propertyData{
		Name:   rstatusp,
		Params: params,
		Value:  val,
	}
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

func (p propertyData) Bytes() []byte {
	var buf bytes.Buffer
	buf.WriteString(p.Name)
	for k, v := range p.Params {
		buf.WriteByte(';')
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.Write(v.Bytes())
	}
	buf.WriteByte(':')
	buf.WriteString(p.Value)
	return buf.Bytes()
}

type altrepLanguageData struct {
	AltRep, Language, Data string
}

func (p *parser) readAltrepLanguageData() (altrepLanguageData, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return altrepLanguageData{}, err
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
		return altrepLanguageData{}, err
	}
	return altrepLanguageData{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

func (a altrepLanguageData) Validate() bool {
	return true
}

func (a altrepLanguageData) data(name string) propertyData {
	params := make(map[string]attributes)
	if a.AltRep != "" {
		params[altrepparam] = altrep(a.AltRep)
	}
	if a.Language != "" {
		params[languageparam] = language(a.Language)
	}
	return propertyData{
		Name:   name,
		Params: params,
		Value:  a.Data,
	}
}

// Errors

var (
	ErrUnsupportedValue            = errors.New("attribute contained unsupported value")
	ErrInvalidAttributeCombination = errors.New("invalid combination of attributes")
)
