package ics

import "time"

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

func (p *parser) readAttendeeProperty() (property, error) {
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

func (p *parser) readContactProperty() (property, error) {
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

func (p *parser) readOrganizerProperty() (property, error) {
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
	DateTime dateTime
}

func (p *parser) readRecurrenceIDProperty() (property, error) {
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

func (p *parser) readRelatedToProperty() (property, error) {
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

func (p *parser) readURLProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return url(v), nil
}

type uid string

func (p *parser) readUIDProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return uid(v), nil
}
