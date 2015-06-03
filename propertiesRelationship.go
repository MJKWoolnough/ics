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

func (a attendee) Validate() bool {
	return a.CalendarUserType >= cuIndividual && a.CalendarUserType <= cuUnknown && a.ParticipationStatus >= psNeedsAction && a.ParticipationStatus <= psInProgress
}

func (a attendee) Data() propertyData {
	params := make(map[string]attribute)
	if a.CalendarUserType != cuUnknown {
		params[cutypeparam] = a.CalendarUserType
	}
	if len(a.Members) > 0 {
		params[memberparam] = a.Members
	}
	if a.Role != prRequiredParticipant {
		params[roleparam] = a.Role
	}
	if a.ParticipationStatus != psNeedsAction {
		params[partstatparam] = a.ParticipationStatus
	}
	if a.RSVP {
		params[rsvpparam] = a.RSVP
	}
	if len(a.Delegatee) > 0 {
		params[deltoparam] = a.Delegatee
	}
	if len(a.Delegator) > 0 {
		params[delfromparam] = a.Delegator
	}
	if a.SentBy != "" {
		params[sentbyparam] = a.SentBy
	}
	if a.DirectoryEntryRef != "" {
		params[dirparam] = a.DirectoryEntryRef
	}
	if a.Language != "" {
		params[languageparam] = a.Language
	}
	return propertyData{
		Name:   attendeep,
		Params: params,
		Value:  string(escape(a.Address)),
	}
}

type contact struct {
	altrepLanguageData
}

func (p *parser) readContactProperty() (property, error) {
	a, err := p.readAltrepLanguageData()
	if err != nil {
		return nil, err
	}
	return contact{a}, nil
}

func (c contact) Data() propertyData {
	return c.data(contactp)
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

func (o organizer) Validate() bool {
	return true
}

func (o organizer) Data() propertyData {
	params := make(map[string]attribute)
	if o.CommonName != "" {
		params[cnparam] = o.CommonName
	}
	if o.DirectoryEntryRef != "" {
		params[dirparam] = o.DirectoryEntryRef
	}
	if o.SentBy != "" {
		params[sentbyparam] = o.SentBy
	}
	if o.Language != "" {
		params[languageparam] = o.Language
	}
	return propertyData{
		Name:   organizerp,
		Params: params,
		Value:  string(escape(o.Name)),
	}
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

func (r recurrenceID) Validate() bool {
	return r.Range == rngThisAndFuture || r.Range == rngThisAndPrior
}

func (r recurrenceID) Data() propertyData {
	params := make(map[string]attribute)
	if r.JustDate {
		params[valuetypeparam] = valueDate
	} else if r.DateTime.Location() != time.UTC {
		params[tzidparam] = timezoneID(r.DateTime.Location().String())
	}
	params[rangeparam] = r.Range
	return propertyData{
		Name:   recuridp,
		Params: params,
		Value:  r.DateTime.String(),
	}
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

func (r relatedTo) Validate() bool {
	return r.RelationshipType >= rtParent && r.RelationshipType <= rtSibling
}

func (r relatedTo) Data() propertyData {
	params := make(map[string]attribute)
	if r.RelationshipType != rtParent {
		params[reltypeparam] = r.RelationshipType
	}
	return propertyData{
		Name:   relatedp,
		Params: params,
		Value:  r.Value,
	}
}

type url string

func (p *parser) readURLProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return url(v), nil
}

func (u url) Validate() bool {
	return true
}

func (u url) Data() propertyData {
	return propertyData{
		Name:  uidp,
		Value: string(u),
	}
}

type uid string

func (p *parser) readUIDProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return uid(v), nil
}

func (u uid) Validate() bool {
	return true
}

func (u uid) Data() propertyData {
	return propertyData{
		Name:  uidp,
		Value: string(u),
	}
}
