package ics

// File automatically generated with ./genSections.sh

import (
	"errors"
	"io"
	"strings"

	"github.com/MJKWoolnough/parser"
)

// Calendar represents a iCalendar object
type Calendar struct {
	Version  PropVersion
	ProdID   PropProdID
	Event    []Event
	Todo     []Todo
	Journal  []Journal
	FreeBusy []FreeBusy
	Timezone []Timezone
}

func (s *Calendar) decode(t tokeniser) error {
	var requiredVersion, requiredProdID bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VEVENT":
				var e Event
				if err := e.decode(t); err != nil {
					return err
				}
				s.Event = append(s.Event, e)
			case "VTODO":
				var e Todo
				if err := e.decode(t); err != nil {
					return err
				}
				s.Todo = append(s.Todo, e)
			case "VJOURNAL":
				var e Journal
				if err := e.decode(t); err != nil {
					return err
				}
				s.Journal = append(s.Journal, e)
			case "VFREEBUSY":
				var e FreeBusy
				if err := e.decode(t); err != nil {
					return err
				}
				s.FreeBusy = append(s.FreeBusy, e)
			case "VTIMEZONE":
				var e Timezone
				if err := e.decode(t); err != nil {
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
			if err := s.Version.decode(params, value); err != nil {
				return err
			}
		case "PRODID":
			if requiredProdID {
				return ErrMultipleSingle
			}
			requiredProdID = true
			if err := s.ProdID.decode(params, value); err != nil {
				return err
			}
		case "END":
			if value != "VCALENDAR" {
				return ErrInvalidEnd
			}
			break Loop
		}
	}
	if !requiredVersion || !requiredProdID {
		return ErrMissingRequired
	}
	return nil
}

func (s *Calendar) encode(w writer) {
	w.WriteString("BEGIN:VCALENDAR\r\n")
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
	w.WriteString("END:VCALENDAR\r\n")
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

// Event provides a group of components that describe an event
type Event struct {
	DateTimeStamp       PropDateTimeStamp
	UID                 PropUID
	DateTimeStart       *PropDateTimeStart
	Class               *PropClass
	Created             *PropCreated
	Description         *PropDescription
	Geo                 *PropGeo
	LastModified        *PropLastModified
	Location            *PropLocation
	Organizer           *PropOrganizer
	Priority            *PropPriority
	Sequence            *PropSequence
	Status              *PropStatus
	Summary             *PropSummary
	TimeTransparency    *PropTimeTransparency
	URL                 *PropURL
	RecurrenceID        *PropRecurrenceID
	RecurrenceRule      *PropRecurrenceRule
	DateTimeEnd         *PropDateTimeEnd
	Duration            *PropDuration
	Attachment          []PropAttachment
	Attendee            []PropAttendee
	Categories          []PropCategories
	Comment             []PropComment
	Contact             []PropContact
	ExceptionDateTime   []PropExceptionDateTime
	RequestStatus       []PropRequestStatus
	RelatedTo           []PropRelatedTo
	Resources           []PropResources
	RecurrenceDateTimes []PropRecurrenceDateTimes
	Alarm               []Alarm
}

func (s *Event) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VALARM":
				var e Alarm
				if err := e.decode(t); err != nil {
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
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return err
			}
		case "CLASS":
			if s.Class != nil {
				return ErrMultipleSingle
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return err
			}
		case "CREATED":
			if s.Created != nil {
				return ErrMultipleSingle
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return err
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return ErrMultipleSingle
			}
			s.Description = new(PropDescription)
			if err := s.Description.decode(params, value); err != nil {
				return err
			}
		case "GEO":
			if s.Geo != nil {
				return ErrMultipleSingle
			}
			s.Geo = new(PropGeo)
			if err := s.Geo.decode(params, value); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return err
			}
		case "LOCATION":
			if s.Location != nil {
				return ErrMultipleSingle
			}
			s.Location = new(PropLocation)
			if err := s.Location.decode(params, value); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return err
			}
		case "PRIORITY":
			if s.Priority != nil {
				return ErrMultipleSingle
			}
			s.Priority = new(PropPriority)
			if err := s.Priority.decode(params, value); err != nil {
				return err
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return ErrMultipleSingle
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return err
			}
		case "STATUS":
			if s.Status != nil {
				return ErrMultipleSingle
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return err
			}
		case "SUMMARY":
			if s.Summary != nil {
				return ErrMultipleSingle
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return err
			}
		case "TRANSP":
			if s.TimeTransparency != nil {
				return ErrMultipleSingle
			}
			s.TimeTransparency = new(PropTimeTransparency)
			if err := s.TimeTransparency.decode(params, value); err != nil {
				return err
			}
		case "URL":
			if s.URL != nil {
				return ErrMultipleSingle
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return err
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return err
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return ErrMultipleSingle
			}
			s.DateTimeEnd = new(PropDateTimeEnd)
			if err := s.DateTimeEnd.decode(params, value); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return err
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VEVENT" {
				return ErrInvalidEnd
			}
			break Loop
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
	w.WriteString("BEGIN:VEVENT\r\n")
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
	if s.Sequence != nil {
		s.Sequence.encode(w)
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
	if s.URL != nil {
		s.URL.encode(w)
	}
	if s.RecurrenceID != nil {
		s.RecurrenceID.encode(w)
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
	for n := range s.RelatedTo {
		s.RelatedTo[n].encode(w)
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
	w.WriteString("END:VEVENT\r\n")
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
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
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
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return err
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
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
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
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

// Todo provides a group of components that describe a to-do
type Todo struct {
	DateTimeStamp       PropDateTimeStamp
	UID                 PropUID
	Class               *PropClass
	Completed           *PropCompleted
	Created             *PropCreated
	Description         *PropDescription
	DateTimeStart       *PropDateTimeStart
	Geo                 *PropGeo
	LastModified        *PropLastModified
	Location            *PropLocation
	Organizer           *PropOrganizer
	PercentComplete     *PropPercentComplete
	Priority            *PropPriority
	RecurrenceID        *PropRecurrenceID
	Sequence            *PropSequence
	Status              *PropStatus
	Summary             *PropSummary
	URL                 *PropURL
	Due                 *PropDue
	Duration            *PropDuration
	Attachment          []PropAttachment
	Attendee            []PropAttendee
	Categories          []PropCategories
	Comment             []PropComment
	Contact             []PropContact
	ExceptionDateTime   []PropExceptionDateTime
	RequestStatus       []PropRequestStatus
	RelatedTo           []PropRelatedTo
	Resources           []PropResources
	RecurrenceDateTimes []PropRecurrenceDateTimes
	Alarm               []Alarm
}

func (s *Todo) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VALARM":
				var e Alarm
				if err := e.decode(t); err != nil {
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
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return err
			}
		case "CLASS":
			if s.Class != nil {
				return ErrMultipleSingle
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return err
			}
		case "COMPLETED":
			if s.Completed != nil {
				return ErrMultipleSingle
			}
			s.Completed = new(PropCompleted)
			if err := s.Completed.decode(params, value); err != nil {
				return err
			}
		case "CREATED":
			if s.Created != nil {
				return ErrMultipleSingle
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return err
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return ErrMultipleSingle
			}
			s.Description = new(PropDescription)
			if err := s.Description.decode(params, value); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return err
			}
		case "GEO":
			if s.Geo != nil {
				return ErrMultipleSingle
			}
			s.Geo = new(PropGeo)
			if err := s.Geo.decode(params, value); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return err
			}
		case "LOCATION":
			if s.Location != nil {
				return ErrMultipleSingle
			}
			s.Location = new(PropLocation)
			if err := s.Location.decode(params, value); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return err
			}
		case "PERCENT-COMPLETE":
			if s.PercentComplete != nil {
				return ErrMultipleSingle
			}
			s.PercentComplete = new(PropPercentComplete)
			if err := s.PercentComplete.decode(params, value); err != nil {
				return err
			}
		case "PRIORITY":
			if s.Priority != nil {
				return ErrMultipleSingle
			}
			s.Priority = new(PropPriority)
			if err := s.Priority.decode(params, value); err != nil {
				return err
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return err
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return ErrMultipleSingle
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return err
			}
		case "STATUS":
			if s.Status != nil {
				return ErrMultipleSingle
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return err
			}
		case "SUMMARY":
			if s.Summary != nil {
				return ErrMultipleSingle
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return err
			}
		case "URL":
			if s.URL != nil {
				return ErrMultipleSingle
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return err
			}
		case "DUE":
			if s.Due != nil {
				return ErrMultipleSingle
			}
			s.Due = new(PropDue)
			if err := s.Due.decode(params, value); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return err
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VTODO" {
				return ErrInvalidEnd
			}
			break Loop
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
	w.WriteString("BEGIN:VTODO\r\n")
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
	if s.PercentComplete != nil {
		s.PercentComplete.encode(w)
	}
	if s.Priority != nil {
		s.Priority.encode(w)
	}
	if s.RecurrenceID != nil {
		s.RecurrenceID.encode(w)
	}
	if s.Sequence != nil {
		s.Sequence.encode(w)
	}
	if s.Status != nil {
		s.Status.encode(w)
	}
	if s.Summary != nil {
		s.Summary.encode(w)
	}
	if s.URL != nil {
		s.URL.encode(w)
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
	for n := range s.RelatedTo {
		s.RelatedTo[n].encode(w)
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
	w.WriteString("END:VTODO\r\n")
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
	if s.PercentComplete != nil {
		if err := s.PercentComplete.valid(); err != nil {
			return err
		}
	}
	if s.Priority != nil {
		if err := s.Priority.valid(); err != nil {
			return err
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return err
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
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
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
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
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
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

// Journal provides a group of components that describe a journal entry
type Journal struct {
	DateTimeStamp       PropDateTimeStamp
	UID                 PropUID
	Class               *PropClass
	Created             *PropCreated
	DateTimeStart       *PropDateTimeStart
	LastModified        *PropLastModified
	Organizer           *PropOrganizer
	RecurrenceID        *PropRecurrenceID
	Sequence            *PropSequence
	Status              *PropStatus
	Summary             *PropSummary
	URL                 *PropURL
	RecurrenceRule      *PropRecurrenceRule
	Attachment          []PropAttachment
	Attendee            []PropAttendee
	Categories          []PropCategories
	Comment             []PropComment
	Contact             []PropContact
	ExceptionDateTime   []PropExceptionDateTime
	RequestStatus       []PropRequestStatus
	RelatedTo           []PropRelatedTo
	Resources           []PropResources
	RecurrenceDateTimes []PropRecurrenceDateTimes
}

func (s *Journal) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
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
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return err
			}
		case "CLASS":
			if s.Class != nil {
				return ErrMultipleSingle
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return err
			}
		case "CREATED":
			if s.Created != nil {
				return ErrMultipleSingle
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return err
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return err
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return ErrMultipleSingle
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return err
			}
		case "STATUS":
			if s.Status != nil {
				return ErrMultipleSingle
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return err
			}
		case "SUMMARY":
			if s.Summary != nil {
				return ErrMultipleSingle
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return err
			}
		case "URL":
			if s.URL != nil {
				return ErrMultipleSingle
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return err
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VJOURNAL" {
				return ErrInvalidEnd
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return ErrMissingRequired
	}
	return nil
}

func (s *Journal) encode(w writer) {
	w.WriteString("BEGIN:VJOURNAL\r\n")
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
	if s.RecurrenceID != nil {
		s.RecurrenceID.encode(w)
	}
	if s.Sequence != nil {
		s.Sequence.encode(w)
	}
	if s.Status != nil {
		s.Status.encode(w)
	}
	if s.Summary != nil {
		s.Summary.encode(w)
	}
	if s.URL != nil {
		s.URL.encode(w)
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
	for n := range s.RelatedTo {
		s.RelatedTo[n].encode(w)
	}
	for n := range s.Resources {
		s.Resources[n].encode(w)
	}
	for n := range s.RecurrenceDateTimes {
		s.RecurrenceDateTimes[n].encode(w)
	}
	w.WriteString("END:VJOURNAL\r\n")
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
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return err
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
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
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
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
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
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

// FreeBusy provides a group of components that describe either a request for
// free/busy time, describe a response to a request for free/busy time, or
// describe a published set of busy time
type FreeBusy struct {
	DateTimeStamp PropDateTimeStamp
	UID           PropUID
	Contact       *PropContact
	DateTimeStart *PropDateTimeStart
	DateTimeEnd   *PropDateTimeEnd
	Organizer     *PropOrganizer
	URL           *PropURL
	Attendee      []PropAttendee
	Comment       []PropComment
	FreeBusy      []PropFreeBusy
	RequestStatus []PropRequestStatus
}

func (s *FreeBusy) decode(t tokeniser) error {
	var requiredDateTimeStamp, requiredUID bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
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
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return err
			}
		case "UID":
			if requiredUID {
				return ErrMultipleSingle
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return err
			}
		case "CONTACT":
			if s.Contact != nil {
				return ErrMultipleSingle
			}
			s.Contact = new(PropContact)
			if err := s.Contact.decode(params, value); err != nil {
				return err
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return ErrMultipleSingle
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return err
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return ErrMultipleSingle
			}
			s.DateTimeEnd = new(PropDateTimeEnd)
			if err := s.DateTimeEnd.decode(params, value); err != nil {
				return err
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return ErrMultipleSingle
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return err
			}
		case "URL":
			if s.URL != nil {
				return ErrMultipleSingle
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return err
			}
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attendee = append(s.Attendee, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "FREEBUSY":
			var e PropFreeBusy
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.FreeBusy = append(s.FreeBusy, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "END":
			if value != "VFREEBUSY" {
				return ErrInvalidEnd
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return ErrMissingRequired
	}
	return nil
}

func (s *FreeBusy) encode(w writer) {
	w.WriteString("BEGIN:VFREEBUSY\r\n")
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
	if s.URL != nil {
		s.URL.encode(w)
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
	w.WriteString("END:VFREEBUSY\r\n")
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
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
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

// Timezone provide a group of components that defines a time zone
type Timezone struct {
	TimezoneID   PropTimezoneID
	LastModified *PropLastModified
	TimezoneURL  *PropTimezoneURL
	Standard     []Standard
	Daylight     []Daylight
}

func (s *Timezone) decode(t tokeniser) error {
	var requiredTimezoneID bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "STANDARD":
				var e Standard
				if err := e.decode(t); err != nil {
					return err
				}
				s.Standard = append(s.Standard, e)
			case "DAYLIGHT":
				var e Daylight
				if err := e.decode(t); err != nil {
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
			if err := s.TimezoneID.decode(params, value); err != nil {
				return err
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return ErrMultipleSingle
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return err
			}
		case "TZURL":
			if s.TimezoneURL != nil {
				return ErrMultipleSingle
			}
			s.TimezoneURL = new(PropTimezoneURL)
			if err := s.TimezoneURL.decode(params, value); err != nil {
				return err
			}
		case "END":
			if value != "VTIMEZONE" {
				return ErrInvalidEnd
			}
			break Loop
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
	w.WriteString("BEGIN:VTIMEZONE\r\n")
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
	w.WriteString("END:VTIMEZONE\r\n")
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

// Standard represents standard timezone rules
type Standard struct {
	DateTimeStart       PropDateTimeStart
	TimezoneOffsetTo    PropTimezoneOffsetTo
	TimezoneOffsetFrom  PropTimezoneOffsetFrom
	RecurrenceRule      *PropRecurrenceRule
	Comment             []PropComment
	RecurrenceDateTimes []PropRecurrenceDateTimes
	TimezoneName        []PropTimezoneName
}

func (s *Standard) decode(t tokeniser) error {
	var requiredDateTimeStart, requiredTimezoneOffsetTo, requiredTimezoneOffsetFrom bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
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
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return err
			}
		case "TZOFFSETTO":
			if requiredTimezoneOffsetTo {
				return ErrMultipleSingle
			}
			requiredTimezoneOffsetTo = true
			if err := s.TimezoneOffsetTo.decode(params, value); err != nil {
				return err
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return ErrMultipleSingle
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(params, value); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return err
			}
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e PropTimezoneName
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if value != "STANDARD" {
				return ErrInvalidEnd
			}
			break Loop
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffsetTo || !requiredTimezoneOffsetFrom {
		return ErrMissingRequired
	}
	return nil
}

func (s *Standard) encode(w writer) {
	w.WriteString("BEGIN:STANDARD\r\n")
	s.DateTimeStart.encode(w)
	s.TimezoneOffsetTo.encode(w)
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
	w.WriteString("END:STANDARD\r\n")
}

func (s *Standard) valid() error {
	if err := s.DateTimeStart.valid(); err != nil {
		return err
	}
	if err := s.TimezoneOffsetTo.valid(); err != nil {
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

// Daylight represents daylight savings timezone rules
type Daylight struct {
	DateTimeStart       PropDateTimeStart
	TimezoneOffsetTo    PropTimezoneOffsetTo
	TimezoneOffsetFrom  PropTimezoneOffsetFrom
	RecurrenceRule      *PropRecurrenceRule
	Comment             []PropComment
	RecurrenceDateTimes []PropRecurrenceDateTimes
	TimezoneName        []PropTimezoneName
}

func (s *Daylight) decode(t tokeniser) error {
	var requiredDateTimeStart, requiredTimezoneOffsetTo, requiredTimezoneOffsetFrom bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
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
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return err
			}
		case "TZOFFSETTO":
			if requiredTimezoneOffsetTo {
				return ErrMultipleSingle
			}
			requiredTimezoneOffsetTo = true
			if err := s.TimezoneOffsetTo.decode(params, value); err != nil {
				return err
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return ErrMultipleSingle
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(params, value); err != nil {
				return err
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return ErrMultipleSingle
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return err
			}
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e PropTimezoneName
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if value != "DAYLIGHT" {
				return ErrInvalidEnd
			}
			break Loop
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffsetTo || !requiredTimezoneOffsetFrom {
		return ErrMissingRequired
	}
	return nil
}

func (s *Daylight) encode(w writer) {
	w.WriteString("BEGIN:DAYLIGHT\r\n")
	s.DateTimeStart.encode(w)
	s.TimezoneOffsetTo.encode(w)
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
	w.WriteString("END:DAYLIGHT\r\n")
}

func (s *Daylight) valid() error {
	if err := s.DateTimeStart.valid(); err != nil {
		return err
	}
	if err := s.TimezoneOffsetTo.valid(); err != nil {
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

// AlarmAudio provides a group of components that define an Audio Alarm
type AlarmAudio struct {
	Trigger    PropTrigger
	Duration   *PropDuration
	Repeat     *PropRepeat
	Attachment []PropAttachment
}

func (s *AlarmAudio) decode(t tokeniser) error {
	var requiredTrigger bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
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
			if err := s.Trigger.decode(params, value); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return err
			}
		case "REPEAT":
			if s.Repeat != nil {
				return ErrMultipleSingle
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return err
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return err
			}
			s.Attachment = append(s.Attachment, e)
		case "END":
			if value != "VALARM" {
				return ErrInvalidEnd
			}
			break Loop
		}
	}
	if !requiredTrigger {
		return ErrMissingRequired
	}
	return nil
}

func (s *AlarmAudio) encode(w writer) {
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

// AlarmDisplay provides a group of components that define a Display Alarm
type AlarmDisplay struct {
	Description PropDescription
	Trigger     PropTrigger
	Duration    *PropDuration
	Repeat      *PropRepeat
}

func (s *AlarmDisplay) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
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
			if err := s.Description.decode(params, value); err != nil {
				return err
			}
		case "TRIGGER":
			if requiredTrigger {
				return ErrMultipleSingle
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return err
			}
		case "REPEAT":
			if s.Repeat != nil {
				return ErrMultipleSingle
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return err
			}
		case "END":
			if value != "VALARM" {
				return ErrInvalidEnd
			}
			break Loop
		}
	}
	if !requiredDescription || !requiredTrigger {
		return ErrMissingRequired
	}
	return nil
}

func (s *AlarmDisplay) encode(w writer) {
	s.Description.encode(w)
	s.Trigger.encode(w)
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	if s.Repeat != nil {
		s.Repeat.encode(w)
	}
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

// AlarmEmail provides a group of components that define an Email Alarm
type AlarmEmail struct {
	Description PropDescription
	Trigger     PropTrigger
	Summary     PropSummary
	Attendee    *PropAttendee
	Duration    *PropDuration
	Repeat      *PropRepeat
}

func (s *AlarmEmail) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger, requiredSummary bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
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
			if err := s.Description.decode(params, value); err != nil {
				return err
			}
		case "TRIGGER":
			if requiredTrigger {
				return ErrMultipleSingle
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return err
			}
		case "SUMMARY":
			if requiredSummary {
				return ErrMultipleSingle
			}
			requiredSummary = true
			if err := s.Summary.decode(params, value); err != nil {
				return err
			}
		case "ATTENDEE":
			if s.Attendee != nil {
				return ErrMultipleSingle
			}
			s.Attendee = new(PropAttendee)
			if err := s.Attendee.decode(params, value); err != nil {
				return err
			}
		case "DURATION":
			if s.Duration != nil {
				return ErrMultipleSingle
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return err
			}
		case "REPEAT":
			if s.Repeat != nil {
				return ErrMultipleSingle
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return err
			}
		case "END":
			if value != "VALARM" {
				return ErrInvalidEnd
			}
			break Loop
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
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			if err := decodeDummy(t, p.Data[len(p.Data)-1].Data); err != nil {
				return err
			}
		case "END":
			if strings.ToUpper(p.Data[len(p.Data)-1].Data) == n {
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
