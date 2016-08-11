package ics

// File automatically generated with ./genSections.sh

import "strings"

type Calendar struct {
	Version  Version
	ProdID   ProdID
	Event    []Event
	Todo     []Todo
	Journal  []Journal
	FreeBusy []FreeBusy
	Timezone []Timezone
}

func (s *Calendar) decode(t tokeniser) error {
	var requiredVersion, requiredProdID bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			case "VEVENT":
				var e Event
				if err := e.Event.decode(p); err != nil {
					return err
				}
				s.Event = append(s.Event, e)
			case "VTODO":
				var e Todo
				if err := e.Todo.decode(p); err != nil {
					return err
				}
				s.Todo = append(s.Todo, e)
			case "VJOURNAL":
				var e Journal
				if err := e.Journal.decode(p); err != nil {
					return err
				}
				s.Journal = append(s.Journal, e)
			case "VFREEBUSY":
				var e FreeBusy
				if err := e.FreeBusy.decode(p); err != nil {
					return err
				}
				s.FreeBusy = append(s.FreeBusy, e)
			case "VTIMEZONE":
				var e Timezone
				if err := e.Timezone.decode(p); err != nil {
					return err
				}
				s.Timezone = append(s.Timezone, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "VERSION":
			if requiredVersion {
				return ErrMultipleSingle
			}
			requiredVersion = true
			if err := s.Version.decode(p); err != nil {
				return err
			}
		case "PRODID":
			if requiredProdID {
				return ErrMultipleSingle
			}
			requiredProdID = true
			if err := s.ProdID.decode(p); err != nil {
				return err
			}
		case "END":
			if p[len(p)-1].Data != "VCALENDAR" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredVersion || !requiredProdID {
		return ErrMissingRequired
	}
	return nil
}

func (s *Calendar) encode(w writer) {
	w.WriteString("BEGIN:VCALENDAR")
	s.Version.encode(w)
	s.ProdID.encode(w)
	for n := range s.Event {
		s.Event[n].encode(w)
	}
	for n := range s.Todo {
		s.Todo[n].encode(w)
	}
	for n := range s.Journal {
		s.Journal[n].encode(w)
	}
	for n := range s.FreeBusy {
		s.FreeBusy[n].encode(w)
	}
	for n := range s.Timezone {
		s.Timezone[n].encode(w)
	}
	w.WriteString("END:VCALENDAR")
}

func (s *Calendar) valid() error {
	if err := s.Version.valid(); err != nil {
		return err
	}
	if err := s.ProdID.valid(); err != nil {
		return err
	}
	for n := range s.Event {
		if err := s.Event[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Todo {
		if err := s.Todo[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Journal {
		if err := s.Journal[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.FreeBusy {
		if err := s.FreeBusy[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Timezone {
		if err := s.Timezone[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type Event struct {
	DateTimeStamp       DateTimeStamp
	UID                 UID
	DateTimeStart       *DateTimeStart
	Class               *Class
	Created             *Created
	Description         *Description
	Geo                 *Geo
	LastModified        *LastModified
	Location            *Location
	Organizer           *Organizer
	Priority            *Priority
	Seq                 *Seq
	Status              *Status
	Summary             *Summary
	TimeTransparency    *TimeTransparency
	Url                 *Url
	RecurID             *RecurID
	RecurrenceRule      *RecurrenceRule
	DateTimeEnd         *DateTimeEnd
	Duration            *Duration
	Attachment          []Attachment
	Attendee            []Attendee
	Categories          []Categories
	Comment             []Comment
	Contact             []Contact
	ExceptionDateTime   []ExceptionDateTime
	RequestStatus       []RequestStatus
	Related             []Related
	Resources           []Resources
	RecurrenceDateTimes []RecurrenceDateTimes
	Alarm               []Alarm
}

func (s *Event) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			case "VALARM":
				var e Alarm
				if err := e.Alarm.decode(p); err != nil {
					return err
				}
				s.Alarm = append(s.Alarm, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return ErrMultipleSingle
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(p); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(p); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(DateTimeStart)
			if err := s.DateTimeStart.decode(p); err != nil {
				return err
			}
		case "CLASS":
			if s.Class != nil {
				return ErrMultipleSingle
			}
			s.Class = new(Class)
			if err := s.Class.decode(p); err != nil {
				return err
			}
		case "CREATED":
			if s.Created != nil {
				return ErrMultipleSingle
			}
			s.Created = new(Created)
			if err := s.Created.decode(p); err != nil {
				return err
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return ErrMultipleSingle
			}
			s.Description = new(Description)
			if err := s.Description.decode(p); err != nil {
				return err
			}
		case "GEO":
			if s.Geo != nil {
				return ErrMultipleSingle
			}
			s.Geo = new(Geo)
			if err := s.Geo.decode(p); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(LastModified)
			if err := s.LastModified.decode(p); err != nil {
				return err
			}
		case "LOCATION":
			if s.Location != nil {
				return ErrMultipleSingle
			}
			s.Location = new(Location)
			if err := s.Location.decode(p); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(Organizer)
			if err := s.Organizer.decode(p); err != nil {
				return err
			}
		case "PRIORITY":
			if s.Priority != nil {
				return ErrMultipleSingle
			}
			s.Priority = new(Priority)
			if err := s.Priority.decode(p); err != nil {
				return err
			}
		case "SEQ":
			if s.Seq != nil {
				return ErrMultipleSingle
			}
			s.Seq = new(Seq)
			if err := s.Seq.decode(p); err != nil {
				return err
			}
		case "STATUS":
			if s.Status != nil {
				return ErrMultipleSingle
			}
			s.Status = new(Status)
			if err := s.Status.decode(p); err != nil {
				return err
			}
		case "SUMMARY":
			if s.Summary != nil {
				return ErrMultipleSingle
			}
			s.Summary = new(Summary)
			if err := s.Summary.decode(p); err != nil {
				return err
			}
		case "TRANSP":
			if s.TimeTransparency != nil {
				return ErrMultipleSingle
			}
			s.TimeTransparency = new(TimeTransparency)
			if err := s.TimeTransparency.decode(p); err != nil {
				return err
			}
		case "URL":
			if s.Url != nil {
				return ErrMultipleSingle
			}
			s.Url = new(Url)
			if err := s.Url.decode(p); err != nil {
				return err
			}
		case "RECURID":
			if s.RecurID != nil {
				return ErrMultipleSingle
			}
			s.RecurID = new(RecurID)
			if err := s.RecurID.decode(p); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(RecurrenceRule)
			if err := s.RecurrenceRule.decode(p); err != nil {
				return err
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return ErrMultipleSingle
			}
			s.DateTimeEnd = new(DateTimeEnd)
			if err := s.DateTimeEnd.decode(p); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(Duration)
			if err := s.Duration.decode(p); err != nil {
				return err
			}
		case "ATTACH":
			var e Attachment
			if err := e.Attachment.decode(p); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e Attendee
			if err := e.Attendee.decode(p); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e Categories
			if err := e.Categories.decode(p); err != nil {
				return err
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e Comment
			if err := e.Comment.decode(p); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e Contact
			if err := e.Contact.decode(p); err != nil {
				return err
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e ExceptionDateTime
			if err := e.ExceptionDateTime.decode(p); err != nil {
				return err
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e RequestStatus
			if err := e.RequestStatus.decode(p); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED":
			var e Related
			if err := e.Related.decode(p); err != nil {
				return err
			}
			s.Related = append(s.Related, e)
		case "RESOURCES":
			var e Resources
			if err := e.Resources.decode(p); err != nil {
				return err
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e RecurrenceDateTimes
			if err := e.RecurrenceDateTimes.decode(p); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if p[len(p)-1].Data != "VEVENT" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return ErrMissingRequired
	}
	if s.DateTimeEnd != nil && s.Duration != nil {
		return ErrRequirementNotMet
	}
	return nil
}

func (s *Event) encode(w writer) {
	w.WriteString("BEGIN:VEVENT")
	s.DateTimeStamp.encode(w)
	s.UID.encode(w)
	if s.DateTimeStart != nil {
		s.DateTimeStart.encode(w)
	}
	if s.Class != nil {
		s.Class.encode(w)
	}
	if s.Created != nil {
		s.Created.encode(w)
	}
	if s.Description != nil {
		s.Description.encode(w)
	}
	if s.Geo != nil {
		s.Geo.encode(w)
	}
	if s.LastModified != nil {
		s.LastModified.encode(w)
	}
	if s.Location != nil {
		s.Location.encode(w)
	}
	if s.Organizer != nil {
		s.Organizer.encode(w)
	}
	if s.Priority != nil {
		s.Priority.encode(w)
	}
	if s.Seq != nil {
		s.Seq.encode(w)
	}
	if s.Status != nil {
		s.Status.encode(w)
	}
	if s.Summary != nil {
		s.Summary.encode(w)
	}
	if s.TimeTransparency != nil {
		s.TimeTransparency.encode(w)
	}
	if s.Url != nil {
		s.Url.encode(w)
	}
	if s.RecurID != nil {
		s.RecurID.encode(w)
	}
	if s.RecurrenceRule != nil {
		s.RecurrenceRule.encode(w)
	}
	if s.DateTimeEnd != nil {
		s.DateTimeEnd.encode(w)
	}
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	for n := range s.Attachment {
		s.Attachment[n].encode(w)
	}
	for n := range s.Attendee {
		s.Attendee[n].encode(w)
	}
	for n := range s.Categories {
		s.Categories[n].encode(w)
	}
	for n := range s.Comment {
		s.Comment[n].encode(w)
	}
	for n := range s.Contact {
		s.Contact[n].encode(w)
	}
	for n := range s.ExceptionDateTime {
		s.ExceptionDateTime[n].encode(w)
	}
	for n := range s.RequestStatus {
		s.RequestStatus[n].encode(w)
	}
	for n := range s.Related {
		s.Related[n].encode(w)
	}
	for n := range s.Resources {
		s.Resources[n].encode(w)
	}
	for n := range s.RecurrenceDateTimes {
		s.RecurrenceDateTimes[n].encode(w)
	}
	for n := range s.Alarm {
		s.Alarm[n].encode(w)
	}
	w.WriteString("END:VEVENT")
}

func (s *Event) valid() error {
	if err := s.DateTimeStamp.valid(); err != nil {
		return err
	}
	if err := s.UID.valid(); err != nil {
		return err
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return err
		}
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return err
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return err
		}
	}
	if s.Description != nil {
		if err := s.Description.valid(); err != nil {
			return err
		}
	}
	if s.Geo != nil {
		if err := s.Geo.valid(); err != nil {
			return err
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return err
		}
	}
	if s.Location != nil {
		if err := s.Location.valid(); err != nil {
			return err
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return err
		}
	}
	if s.Priority != nil {
		if err := s.Priority.valid(); err != nil {
			return err
		}
	}
	if s.Seq != nil {
		if err := s.Seq.valid(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return err
		}
	}
	if s.TimeTransparency != nil {
		if err := s.TimeTransparency.valid(); err != nil {
			return err
		}
	}
	if s.Url != nil {
		if err := s.Url.valid(); err != nil {
			return err
		}
	}
	if s.RecurID != nil {
		if err := s.RecurID.valid(); err != nil {
			return err
		}
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return err
		}
	}
	if s.DateTimeEnd != nil {
		if err := s.DateTimeEnd.valid(); err != nil {
			return err
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return err
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Related {
		if err := s.Related[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Alarm {
		if err := s.Alarm[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type Todo struct {
	DateTimeStamp       DateTimeStamp
	UID                 UID
	Class               *Class
	Completed           *Completed
	Created             *Created
	Description         *Description
	DateTimeStart       *DateTimeStart
	Duration            *Duration
	Geo                 *Geo
	LastModified        *LastModified
	Location            *Location
	Organizer           *Organizer
	Percent             *Percent
	Priority            *Priority
	RecurID             *RecurID
	Seq                 *Seq
	Status              *Status
	Summary             *Summary
	Url                 *Url
	Due                 *Due
	Duration            *Duration
	Attachment          []Attachment
	Attendee            []Attendee
	Categories          []Categories
	Comment             []Comment
	Contact             []Contact
	ExceptionDateTime   []ExceptionDateTime
	RequestStatus       []RequestStatus
	Related             []Related
	Resources           []Resources
	RecurrenceDateTimes []RecurrenceDateTimes
}

func (s *Todo) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return ErrMultipleSingle
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(p); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(p); err != nil {
				return err
			}
		case "CLASS":
			if s.Class != nil {
				return ErrMultipleSingle
			}
			s.Class = new(Class)
			if err := s.Class.decode(p); err != nil {
				return err
			}
		case "COMPLETED":
			if s.Completed != nil {
				return ErrMultipleSingle
			}
			s.Completed = new(Completed)
			if err := s.Completed.decode(p); err != nil {
				return err
			}
		case "CREATED":
			if s.Created != nil {
				return ErrMultipleSingle
			}
			s.Created = new(Created)
			if err := s.Created.decode(p); err != nil {
				return err
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return ErrMultipleSingle
			}
			s.Description = new(Description)
			if err := s.Description.decode(p); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(DateTimeStart)
			if err := s.DateTimeStart.decode(p); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(Duration)
			if err := s.Duration.decode(p); err != nil {
				return err
			}
		case "GEO":
			if s.Geo != nil {
				return ErrMultipleSingle
			}
			s.Geo = new(Geo)
			if err := s.Geo.decode(p); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(LastModified)
			if err := s.LastModified.decode(p); err != nil {
				return err
			}
		case "LOCATION":
			if s.Location != nil {
				return ErrMultipleSingle
			}
			s.Location = new(Location)
			if err := s.Location.decode(p); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(Organizer)
			if err := s.Organizer.decode(p); err != nil {
				return err
			}
		case "PERCENT":
			if s.Percent != nil {
				return ErrMultipleSingle
			}
			s.Percent = new(Percent)
			if err := s.Percent.decode(p); err != nil {
				return err
			}
		case "PRIORITY":
			if s.Priority != nil {
				return ErrMultipleSingle
			}
			s.Priority = new(Priority)
			if err := s.Priority.decode(p); err != nil {
				return err
			}
		case "RECURID":
			if s.RecurID != nil {
				return ErrMultipleSingle
			}
			s.RecurID = new(RecurID)
			if err := s.RecurID.decode(p); err != nil {
				return err
			}
		case "SEQ":
			if s.Seq != nil {
				return ErrMultipleSingle
			}
			s.Seq = new(Seq)
			if err := s.Seq.decode(p); err != nil {
				return err
			}
		case "STATUS":
			if s.Status != nil {
				return ErrMultipleSingle
			}
			s.Status = new(Status)
			if err := s.Status.decode(p); err != nil {
				return err
			}
		case "SUMMARY":
			if s.Summary != nil {
				return ErrMultipleSingle
			}
			s.Summary = new(Summary)
			if err := s.Summary.decode(p); err != nil {
				return err
			}
		case "URL":
			if s.Url != nil {
				return ErrMultipleSingle
			}
			s.Url = new(Url)
			if err := s.Url.decode(p); err != nil {
				return err
			}
		case "DUE":
			if s.Due != nil {
				return ErrMultipleSingle
			}
			s.Due = new(Due)
			if err := s.Due.decode(p); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(Duration)
			if err := s.Duration.decode(p); err != nil {
				return err
			}
		case "ATTACH":
			var e Attachment
			if err := e.Attachment.decode(p); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e Attendee
			if err := e.Attendee.decode(p); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e Categories
			if err := e.Categories.decode(p); err != nil {
				return err
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e Comment
			if err := e.Comment.decode(p); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e Contact
			if err := e.Contact.decode(p); err != nil {
				return err
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e ExceptionDateTime
			if err := e.ExceptionDateTime.decode(p); err != nil {
				return err
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e RequestStatus
			if err := e.RequestStatus.decode(p); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED":
			var e Related
			if err := e.Related.decode(p); err != nil {
				return err
			}
			s.Related = append(s.Related, e)
		case "RESOURCES":
			var e Resources
			if err := e.Resources.decode(p); err != nil {
				return err
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e RecurrenceDateTimes
			if err := e.RecurrenceDateTimes.decode(p); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if p[len(p)-1].Data != "VTODO" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return ErrMissingRequired
	}
	if s.Duration != nil && (s.DateTimeStart == nil) {
		return ErrRequirementNotMet
	}
	if s.Due != nil && s.Duration != nil {
		return ErrRequirementNotMet
	}
	return nil
}

func (s *Todo) encode(w writer) {
	w.WriteString("BEGIN:VTODO")
	s.DateTimeStamp.encode(w)
	s.UID.encode(w)
	if s.Class != nil {
		s.Class.encode(w)
	}
	if s.Completed != nil {
		s.Completed.encode(w)
	}
	if s.Created != nil {
		s.Created.encode(w)
	}
	if s.Description != nil {
		s.Description.encode(w)
	}
	if s.DateTimeStart != nil {
		s.DateTimeStart.encode(w)
	}
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	if s.Geo != nil {
		s.Geo.encode(w)
	}
	if s.LastModified != nil {
		s.LastModified.encode(w)
	}
	if s.Location != nil {
		s.Location.encode(w)
	}
	if s.Organizer != nil {
		s.Organizer.encode(w)
	}
	if s.Percent != nil {
		s.Percent.encode(w)
	}
	if s.Priority != nil {
		s.Priority.encode(w)
	}
	if s.RecurID != nil {
		s.RecurID.encode(w)
	}
	if s.Seq != nil {
		s.Seq.encode(w)
	}
	if s.Status != nil {
		s.Status.encode(w)
	}
	if s.Summary != nil {
		s.Summary.encode(w)
	}
	if s.Url != nil {
		s.Url.encode(w)
	}
	if s.Due != nil {
		s.Due.encode(w)
	}
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	for n := range s.Attachment {
		s.Attachment[n].encode(w)
	}
	for n := range s.Attendee {
		s.Attendee[n].encode(w)
	}
	for n := range s.Categories {
		s.Categories[n].encode(w)
	}
	for n := range s.Comment {
		s.Comment[n].encode(w)
	}
	for n := range s.Contact {
		s.Contact[n].encode(w)
	}
	for n := range s.ExceptionDateTime {
		s.ExceptionDateTime[n].encode(w)
	}
	for n := range s.RequestStatus {
		s.RequestStatus[n].encode(w)
	}
	for n := range s.Related {
		s.Related[n].encode(w)
	}
	for n := range s.Resources {
		s.Resources[n].encode(w)
	}
	for n := range s.RecurrenceDateTimes {
		s.RecurrenceDateTimes[n].encode(w)
	}
	w.WriteString("END:VTODO")
}

func (s *Todo) valid() error {
	if err := s.DateTimeStamp.valid(); err != nil {
		return err
	}
	if err := s.UID.valid(); err != nil {
		return err
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return err
		}
	}
	if s.Completed != nil {
		if err := s.Completed.valid(); err != nil {
			return err
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return err
		}
	}
	if s.Description != nil {
		if err := s.Description.valid(); err != nil {
			return err
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return err
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return err
		}
	}
	if s.Geo != nil {
		if err := s.Geo.valid(); err != nil {
			return err
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return err
		}
	}
	if s.Location != nil {
		if err := s.Location.valid(); err != nil {
			return err
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return err
		}
	}
	if s.Percent != nil {
		if err := s.Percent.valid(); err != nil {
			return err
		}
	}
	if s.Priority != nil {
		if err := s.Priority.valid(); err != nil {
			return err
		}
	}
	if s.RecurID != nil {
		if err := s.RecurID.valid(); err != nil {
			return err
		}
	}
	if s.Seq != nil {
		if err := s.Seq.valid(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return err
		}
	}
	if s.Url != nil {
		if err := s.Url.valid(); err != nil {
			return err
		}
	}
	if s.Due != nil {
		if err := s.Due.valid(); err != nil {
			return err
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return err
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Related {
		if err := s.Related[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type Journal struct {
	DateTimeStamp       DateTimeStamp
	UID                 UID
	Class               *Class
	Created             *Created
	DateTimeStart       *DateTimeStart
	LastModified        *LastModified
	Organizer           *Organizer
	RecurID             *RecurID
	Seq                 *Seq
	Status              *Status
	Summary             *Summary
	Url                 *Url
	RecurrenceRule      *RecurrenceRule
	Attachment          []Attachment
	Attendee            []Attendee
	Categories          []Categories
	Comment             []Comment
	Contact             []Contact
	ExceptionDateTime   []ExceptionDateTime
	RequestStatus       []RequestStatus
	Related             []Related
	Resources           []Resources
	RecurrenceDateTimes []RecurrenceDateTimes
}

func (s *Journal) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return ErrMultipleSingle
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(p); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(p); err != nil {
				return err
			}
		case "CLASS":
			if s.Class != nil {
				return ErrMultipleSingle
			}
			s.Class = new(Class)
			if err := s.Class.decode(p); err != nil {
				return err
			}
		case "CREATED":
			if s.Created != nil {
				return ErrMultipleSingle
			}
			s.Created = new(Created)
			if err := s.Created.decode(p); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(DateTimeStart)
			if err := s.DateTimeStart.decode(p); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(LastModified)
			if err := s.LastModified.decode(p); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(Organizer)
			if err := s.Organizer.decode(p); err != nil {
				return err
			}
		case "RECURID":
			if s.RecurID != nil {
				return ErrMultipleSingle
			}
			s.RecurID = new(RecurID)
			if err := s.RecurID.decode(p); err != nil {
				return err
			}
		case "SEQ":
			if s.Seq != nil {
				return ErrMultipleSingle
			}
			s.Seq = new(Seq)
			if err := s.Seq.decode(p); err != nil {
				return err
			}
		case "STATUS":
			if s.Status != nil {
				return ErrMultipleSingle
			}
			s.Status = new(Status)
			if err := s.Status.decode(p); err != nil {
				return err
			}
		case "SUMMARY":
			if s.Summary != nil {
				return ErrMultipleSingle
			}
			s.Summary = new(Summary)
			if err := s.Summary.decode(p); err != nil {
				return err
			}
		case "URL":
			if s.Url != nil {
				return ErrMultipleSingle
			}
			s.Url = new(Url)
			if err := s.Url.decode(p); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(RecurrenceRule)
			if err := s.RecurrenceRule.decode(p); err != nil {
				return err
			}
		case "ATTACH":
			var e Attachment
			if err := e.Attachment.decode(p); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e Attendee
			if err := e.Attendee.decode(p); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e Categories
			if err := e.Categories.decode(p); err != nil {
				return err
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e Comment
			if err := e.Comment.decode(p); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e Contact
			if err := e.Contact.decode(p); err != nil {
				return err
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e ExceptionDateTime
			if err := e.ExceptionDateTime.decode(p); err != nil {
				return err
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e RequestStatus
			if err := e.RequestStatus.decode(p); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED":
			var e Related
			if err := e.Related.decode(p); err != nil {
				return err
			}
			s.Related = append(s.Related, e)
		case "RESOURCES":
			var e Resources
			if err := e.Resources.decode(p); err != nil {
				return err
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e RecurrenceDateTimes
			if err := e.RecurrenceDateTimes.decode(p); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if p[len(p)-1].Data != "VJOURNAL" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return ErrMissingRequired
	}
	return nil
}

func (s *Journal) encode(w writer) {
	w.WriteString("BEGIN:VJOURNAL")
	s.DateTimeStamp.encode(w)
	s.UID.encode(w)
	if s.Class != nil {
		s.Class.encode(w)
	}
	if s.Created != nil {
		s.Created.encode(w)
	}
	if s.DateTimeStart != nil {
		s.DateTimeStart.encode(w)
	}
	if s.LastModified != nil {
		s.LastModified.encode(w)
	}
	if s.Organizer != nil {
		s.Organizer.encode(w)
	}
	if s.RecurID != nil {
		s.RecurID.encode(w)
	}
	if s.Seq != nil {
		s.Seq.encode(w)
	}
	if s.Status != nil {
		s.Status.encode(w)
	}
	if s.Summary != nil {
		s.Summary.encode(w)
	}
	if s.Url != nil {
		s.Url.encode(w)
	}
	if s.RecurrenceRule != nil {
		s.RecurrenceRule.encode(w)
	}
	for n := range s.Attachment {
		s.Attachment[n].encode(w)
	}
	for n := range s.Attendee {
		s.Attendee[n].encode(w)
	}
	for n := range s.Categories {
		s.Categories[n].encode(w)
	}
	for n := range s.Comment {
		s.Comment[n].encode(w)
	}
	for n := range s.Contact {
		s.Contact[n].encode(w)
	}
	for n := range s.ExceptionDateTime {
		s.ExceptionDateTime[n].encode(w)
	}
	for n := range s.RequestStatus {
		s.RequestStatus[n].encode(w)
	}
	for n := range s.Related {
		s.Related[n].encode(w)
	}
	for n := range s.Resources {
		s.Resources[n].encode(w)
	}
	for n := range s.RecurrenceDateTimes {
		s.RecurrenceDateTimes[n].encode(w)
	}
	w.WriteString("END:VJOURNAL")
}

func (s *Journal) valid() error {
	if err := s.DateTimeStamp.valid(); err != nil {
		return err
	}
	if err := s.UID.valid(); err != nil {
		return err
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return err
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return err
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return err
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return err
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return err
		}
	}
	if s.RecurID != nil {
		if err := s.RecurID.valid(); err != nil {
			return err
		}
	}
	if s.Seq != nil {
		if err := s.Seq.valid(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return err
		}
	}
	if s.Url != nil {
		if err := s.Url.valid(); err != nil {
			return err
		}
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return err
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Related {
		if err := s.Related[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type FreeBusy struct {
	DateTimeStamp DateTimeStamp
	UID           UID
	Contact       *Contact
	DateTimeStart *DateTimeStart
	DateTimeEnd   *DateTimeEnd
	Organizer     *Organizer
	Url           *Url
	Attendee      []Attendee
	Comment       []Comment
	FreeBusy      []FreeBusy
	RequestStatus []RequestStatus
}

func (s *FreeBusy) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return ErrMultipleSingle
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(p); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(p); err != nil {
				return err
			}
		case "CONTACT":
			if s.Contact != nil {
				return ErrMultipleSingle
			}
			s.Contact = new(Contact)
			if err := s.Contact.decode(p); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(DateTimeStart)
			if err := s.DateTimeStart.decode(p); err != nil {
				return err
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return ErrMultipleSingle
			}
			s.DateTimeEnd = new(DateTimeEnd)
			if err := s.DateTimeEnd.decode(p); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(Organizer)
			if err := s.Organizer.decode(p); err != nil {
				return err
			}
		case "URL":
			if s.Url != nil {
				return ErrMultipleSingle
			}
			s.Url = new(Url)
			if err := s.Url.decode(p); err != nil {
				return err
			}
		case "ATTENDEE":
			var e Attendee
			if err := e.Attendee.decode(p); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "COMMENT":
			var e Comment
			if err := e.Comment.decode(p); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "FREEBUSY":
			var e FreeBusy
			if err := e.FreeBusy.decode(p); err != nil {
				return err
			}
			s.FreeBusy = append(s.FreeBusy, e)
		case "REQUEST-STATUS":
			var e RequestStatus
			if err := e.RequestStatus.decode(p); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "END":
			if p[len(p)-1].Data != "VFREEBUSY" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return ErrMissingRequired
	}
	return nil
}

func (s *FreeBusy) encode(w writer) {
	w.WriteString("BEGIN:VFREEBUSY")
	s.DateTimeStamp.encode(w)
	s.UID.encode(w)
	if s.Contact != nil {
		s.Contact.encode(w)
	}
	if s.DateTimeStart != nil {
		s.DateTimeStart.encode(w)
	}
	if s.DateTimeEnd != nil {
		s.DateTimeEnd.encode(w)
	}
	if s.Organizer != nil {
		s.Organizer.encode(w)
	}
	if s.Url != nil {
		s.Url.encode(w)
	}
	for n := range s.Attendee {
		s.Attendee[n].encode(w)
	}
	for n := range s.Comment {
		s.Comment[n].encode(w)
	}
	for n := range s.FreeBusy {
		s.FreeBusy[n].encode(w)
	}
	for n := range s.RequestStatus {
		s.RequestStatus[n].encode(w)
	}
	w.WriteString("END:VFREEBUSY")
}

func (s *FreeBusy) valid() error {
	if err := s.DateTimeStamp.valid(); err != nil {
		return err
	}
	if err := s.UID.valid(); err != nil {
		return err
	}
	if s.Contact != nil {
		if err := s.Contact.valid(); err != nil {
			return err
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return err
		}
	}
	if s.DateTimeEnd != nil {
		if err := s.DateTimeEnd.valid(); err != nil {
			return err
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return err
		}
	}
	if s.Url != nil {
		if err := s.Url.valid(); err != nil {
			return err
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.FreeBusy {
		if err := s.FreeBusy[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type Timezone struct {
	TimezoneID   TimezoneID
	LastModified *LastModified
	TimezoneURL  *TimezoneURL
	Standard     []Standard
	Daylight     []Daylight
}

func (s *Timezone) decode(t tokeniser) error {
	var requiredTimezoneID bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			case "STANDARD":
				var e Standard
				if err := e.Standard.decode(p); err != nil {
					return err
				}
				s.Standard = append(s.Standard, e)
			case "DAYLIGHT":
				var e Daylight
				if err := e.Daylight.decode(p); err != nil {
					return err
				}
				s.Daylight = append(s.Daylight, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "TZID":
			if requiredTimezoneID {
				return ErrMultipleSingle
			}
			requiredTimezoneID = true
			if err := s.TimezoneID.decode(p); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(LastModified)
			if err := s.LastModified.decode(p); err != nil {
				return err
			}
		case "TZURL":
			if s.TimezoneURL != nil {
				return ErrMultipleSingle
			}
			s.TimezoneURL = new(TimezoneURL)
			if err := s.TimezoneURL.decode(p); err != nil {
				return err
			}
		case "END":
			if p[len(p)-1].Data != "VTIMEZONE" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredTimezoneID {
		return ErrMissingRequired
	}
	if s.Standard == nil && s.Daylight == nil {
		return ErrRequirementNotMet
	}
	return nil
}

func (s *Timezone) encode(w writer) {
	w.WriteString("BEGIN:VTIMEZONE")
	s.TimezoneID.encode(w)
	if s.LastModified != nil {
		s.LastModified.encode(w)
	}
	if s.TimezoneURL != nil {
		s.TimezoneURL.encode(w)
	}
	for n := range s.Standard {
		s.Standard[n].encode(w)
	}
	for n := range s.Daylight {
		s.Daylight[n].encode(w)
	}
	w.WriteString("END:VTIMEZONE")
}

func (s *Timezone) valid() error {
	if err := s.TimezoneID.valid(); err != nil {
		return err
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return err
		}
	}
	if s.TimezoneURL != nil {
		if err := s.TimezoneURL.valid(); err != nil {
			return err
		}
	}
	for n := range s.Standard {
		if err := s.Standard[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.Daylight {
		if err := s.Daylight[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type Standard struct {
	DateTimeStart       DateTimeStart
	TimezoneOffset      TimezoneOffset
	TimezoneOffsetFrom  TimezoneOffsetFrom
	RecurrenceRule      *RecurrenceRule
	Comment             []Comment
	RecurrenceDateTimes []RecurrenceDateTimes
	TimezoneName        []TimezoneName
}

func (s *Standard) decode(t tokeniser) error {
	var requiredDateTimeStart, requiredTimezoneOffset, requiredTimezoneOffsetFrom bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DTSTART":
			if requiredDateTimeStart {
				return ErrMultipleSingle
			}
			requiredDateTimeStart = true
			if err := s.DateTimeStart.decode(p); err != nil {
				return err
			}
		case "TZOFFSET":
			if requiredTimezoneOffset {
				return ErrMultipleSingle
			}
			requiredTimezoneOffset = true
			if err := s.TimezoneOffset.decode(p); err != nil {
				return err
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return ErrMultipleSingle
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(p); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(RecurrenceRule)
			if err := s.RecurrenceRule.decode(p); err != nil {
				return err
			}
		case "COMMENT":
			var e Comment
			if err := e.Comment.decode(p); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e RecurrenceDateTimes
			if err := e.RecurrenceDateTimes.decode(p); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e TimezoneName
			if err := e.TimezoneName.decode(p); err != nil {
				return err
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if p[len(p)-1].Data != "STANDARD" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffset || !requiredTimezoneOffsetFrom {
		return ErrMissingRequired
	}
	return nil
}

func (s *Standard) encode(w writer) {
	w.WriteString("BEGIN:STANDARD")
	s.DateTimeStart.encode(w)
	s.TimezoneOffset.encode(w)
	s.TimezoneOffsetFrom.encode(w)
	if s.RecurrenceRule != nil {
		s.RecurrenceRule.encode(w)
	}
	for n := range s.Comment {
		s.Comment[n].encode(w)
	}
	for n := range s.RecurrenceDateTimes {
		s.RecurrenceDateTimes[n].encode(w)
	}
	for n := range s.TimezoneName {
		s.TimezoneName[n].encode(w)
	}
	w.WriteString("END:STANDARD")
}

func (s *Standard) valid() error {
	if err := s.DateTimeStart.valid(); err != nil {
		return err
	}
	if err := s.TimezoneOffset.valid(); err != nil {
		return err
	}
	if err := s.TimezoneOffsetFrom.valid(); err != nil {
		return err
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return err
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.TimezoneName {
		if err := s.TimezoneName[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type Daylight struct {
	DateTimeStart       DateTimeStart
	TimezoneOffset      TimezoneOffset
	TimezoneOffsetFrom  TimezoneOffsetFrom
	RecurrenceRule      *RecurrenceRule
	Comment             []Comment
	RecurrenceDateTimes []RecurrenceDateTimes
	TimezoneName        []TimezoneName
}

func (s *Daylight) decode(t tokeniser) error {
	var requiredDateTimeStart, requiredTimezoneOffset, requiredTimezoneOffsetFrom bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DTSTART":
			if requiredDateTimeStart {
				return ErrMultipleSingle
			}
			requiredDateTimeStart = true
			if err := s.DateTimeStart.decode(p); err != nil {
				return err
			}
		case "TZOFFSET":
			if requiredTimezoneOffset {
				return ErrMultipleSingle
			}
			requiredTimezoneOffset = true
			if err := s.TimezoneOffset.decode(p); err != nil {
				return err
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return ErrMultipleSingle
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(p); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(RecurrenceRule)
			if err := s.RecurrenceRule.decode(p); err != nil {
				return err
			}
		case "COMMENT":
			var e Comment
			if err := e.Comment.decode(p); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e RecurrenceDateTimes
			if err := e.RecurrenceDateTimes.decode(p); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e TimezoneName
			if err := e.TimezoneName.decode(p); err != nil {
				return err
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if p[len(p)-1].Data != "DAYLIGHT" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffset || !requiredTimezoneOffsetFrom {
		return ErrMissingRequired
	}
	return nil
}

func (s *Daylight) encode(w writer) {
	w.WriteString("BEGIN:DAYLIGHT")
	s.DateTimeStart.encode(w)
	s.TimezoneOffset.encode(w)
	s.TimezoneOffsetFrom.encode(w)
	if s.RecurrenceRule != nil {
		s.RecurrenceRule.encode(w)
	}
	for n := range s.Comment {
		s.Comment[n].encode(w)
	}
	for n := range s.RecurrenceDateTimes {
		s.RecurrenceDateTimes[n].encode(w)
	}
	for n := range s.TimezoneName {
		s.TimezoneName[n].encode(w)
	}
	w.WriteString("END:DAYLIGHT")
}

func (s *Daylight) valid() error {
	if err := s.DateTimeStart.valid(); err != nil {
		return err
	}
	if err := s.TimezoneOffset.valid(); err != nil {
		return err
	}
	if err := s.TimezoneOffsetFrom.valid(); err != nil {
		return err
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return err
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return err
		}
	}
	for n := range s.TimezoneName {
		if err := s.TimezoneName[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type AlarmAudio struct {
	Trigger    Trigger
	Duration   *Duration
	Repeat     *Repeat
	Attachment []Attachment
}

func (s *AlarmAudio) decode(t tokeniser) error {
	var requiredTrigger bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "TRIGGER":
			if requiredTrigger {
				return ErrMultipleSingle
			}
			requiredTrigger = true
			if err := s.Trigger.decode(p); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(Duration)
			if err := s.Duration.decode(p); err != nil {
				return err
			}
		case "REPEAT":
			if s.Repeat != nil {
				return ErrMultipleSingle
			}
			s.Repeat = new(Repeat)
			if err := s.Repeat.decode(p); err != nil {
				return err
			}
		case "ATTACH":
			var e Attachment
			if err := e.Attachment.decode(p); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "END":
			if p[len(p)-1].Data != "VALARMAUDIO" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredTrigger {
		return ErrMissingRequired
	}
	return nil
}

func (s *AlarmAudio) encode(w writer) {
	w.WriteString("BEGIN:VALARM")
	s.Trigger.encode(w)
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	if s.Repeat != nil {
		s.Repeat.encode(w)
	}
	for n := range s.Attachment {
		s.Attachment[n].encode(w)
	}
	w.WriteString("END:VALARM")
}

func (s *AlarmAudio) valid() error {
	if err := s.Trigger.valid(); err != nil {
		return err
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return err
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return err
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return err
		}
	}
	return nil
}

type AlarmDisplay struct {
	Description Description
	Trigger     Trigger
	Duration    *Duration
	Repeat      *Repeat
}

func (s *AlarmDisplay) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DESCRIPTION":
			if requiredDescription {
				return ErrMultipleSingle
			}
			requiredDescription = true
			if err := s.Description.decode(p); err != nil {
				return err
			}
		case "TRIGGER":
			if requiredTrigger {
				return ErrMultipleSingle
			}
			requiredTrigger = true
			if err := s.Trigger.decode(p); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(Duration)
			if err := s.Duration.decode(p); err != nil {
				return err
			}
		case "REPEAT":
			if s.Repeat != nil {
				return ErrMultipleSingle
			}
			s.Repeat = new(Repeat)
			if err := s.Repeat.decode(p); err != nil {
				return err
			}
		case "END":
			if p[len(p)-1].Data != "VALARMDISPLAY" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDescription || !requiredTrigger {
		return ErrMissingRequired
	}
	return nil
}

func (s *AlarmDisplay) encode(w writer) {
	w.WriteString("BEGIN:VALARM")
	s.Description.encode(w)
	s.Trigger.encode(w)
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	if s.Repeat != nil {
		s.Repeat.encode(w)
	}
	w.WriteString("END:VALARM")
}

func (s *AlarmDisplay) valid() error {
	if err := s.Description.valid(); err != nil {
		return err
	}
	if err := s.Trigger.valid(); err != nil {
		return err
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return err
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return err
		}
	}
	return nil
}

type AlarmEmail struct {
	Description Description
	Trigger     Trigger
	Summary     Summary
	Attendee    *Attendee
	Duration    *Duration
	Repeat      *Repeat
}

func (s *AlarmEmail) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger, requiredSummary bool
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(p[len(p)-1].Data); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return err
				}
			}
		case "DESCRIPTION":
			if requiredDescription {
				return ErrMultipleSingle
			}
			requiredDescription = true
			if err := s.Description.decode(p); err != nil {
				return err
			}
		case "TRIGGER":
			if requiredTrigger {
				return ErrMultipleSingle
			}
			requiredTrigger = true
			if err := s.Trigger.decode(p); err != nil {
				return err
			}
		case "SUMMARY":
			if requiredSummary {
				return ErrMultipleSingle
			}
			requiredSummary = true
			if err := s.Summary.decode(p); err != nil {
				return err
			}
		case "ATTENDEE":
			if s.Attendee != nil {
				return ErrMultipleSingle
			}
			s.Attendee = new(Attendee)
			if err := s.Attendee.decode(p); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(Duration)
			if err := s.Duration.decode(p); err != nil {
				return err
			}
		case "REPEAT":
			if s.Repeat != nil {
				return ErrMultipleSingle
			}
			s.Repeat = new(Repeat)
			if err := s.Repeat.decode(p); err != nil {
				return err
			}
		case "END":
			if p[len(p)-1].Data != "VALARMEMAIL" {
				return ErrInvalidEnd
			}
			break
		}
	}
	if !requiredDescription || !requiredTrigger || !requiredSummary {
		return ErrMissingRequired
	}
	if t := s.Duration == nil; t == (s.Repeat == nil) {
		return ErrRequirementNotMet
	}
	return nil
}

func (s *AlarmEmail) encode(w writer) {
	w.WriteString("BEGIN:VALARM")
	s.Description.encode(w)
	s.Trigger.encode(w)
	s.Summary.encode(w)
	if s.Attendee != nil {
		s.Attendee.encode(w)
	}
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	if s.Repeat != nil {
		s.Repeat.encode(w)
	}
	w.WriteString("END:VALARM")
}

func (s *AlarmEmail) valid() error {
	if err := s.Description.valid(); err != nil {
		return err
	}
	if err := s.Trigger.valid(); err != nil {
		return err
	}
	if err := s.Summary.valid(); err != nil {
		return err
	}
	if s.Attendee != nil {
		if err := s.Attendee.valid(); err != nil {
			return err
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return err
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return err
		}
	}
	return nil
}

// decodeDummy reads unknown sections, discarding the data
func decodeDummy(t tokeniser, n string) error {
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		}
		switch strings.ToUpper(p[0].Data) {
		case "BEGIN":
			if err := decodeDummy(t, p[len(p)-1].Data); err != nil {
				return err
			}
		case "END":
			if strings.ToUpper(p[len(p)-1].Data) == n {
				return nil
			}
			return ErrInvalidEnd
		}
	}
}

// Errors
var (
	ErrMultipleSingle    = errors.New("unique property found multiple times")
	ErrInvalidEnd        = errors.New("invalid end of section")
	ErrMissingRequired   = errors.New("required property missing")
	ErrRequirementNotMet = errors.New("requirement not met")
)
