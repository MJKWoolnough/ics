package ics

// File automatically generated with ./genSections.sh

import (
	"io"
	"strings"

	"vimagination.zapto.org/errors"
	"vimagination.zapto.org/parser"
)

// Calendar represents a iCalendar object
type Calendar struct {
	Version   PropVersion
	ProductID PropProductID
	Event     []Event
	Todo      []Todo
	Journal   []Journal
	FreeBusy  []FreeBusy
	Timezone  []Timezone
}

func (s *Calendar) decode(t tokeniser) error {
	var requiredVersion, requiredProductID bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return errors.WithContext("error decoding Calendar: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding Calendar: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VEVENT":
				var e Event
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Calendar->Event: ", err)
				}
				s.Event = append(s.Event, e)
			case "VTODO":
				var e Todo
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Calendar->Todo: ", err)
				}
				s.Todo = append(s.Todo, e)
			case "VJOURNAL":
				var e Journal
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Calendar->Journal: ", err)
				}
				s.Journal = append(s.Journal, e)
			case "VFREEBUSY":
				var e FreeBusy
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Calendar->FreeBusy: ", err)
				}
				s.FreeBusy = append(s.FreeBusy, e)
			case "VTIMEZONE":
				var e Timezone
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Calendar->Timezone: ", err)
				}
				s.Timezone = append(s.Timezone, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding Calendar: ", err)
				}
			}
		case "VERSION":
			if requiredVersion {
				return errors.Error("error decoding Calendar: multiple Version")
			}
			requiredVersion = true
			if err := s.Version.decode(params, value); err != nil {
				return errors.WithContext("error decoding Calendar->Version: ", err)
			}
		case "PRODID":
			if requiredProductID {
				return errors.Error("error decoding Calendar: multiple ProductID")
			}
			requiredProductID = true
			if err := s.ProductID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Calendar->ProductID: ", err)
			}
		case "END":
			if value != "VCALENDAR" {
				return errors.WithContext("error decoding Calendar: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredVersion || !requiredProductID {
		return errors.WithContext("error decoding Calendar: ", ErrMissingRequired)
	}
	return nil
}

func (s *Calendar) encode(w writer) {
	w.WriteString("BEGIN:VCALENDAR\r\n")
	s.Version.encode(w)
	s.ProductID.encode(w)
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
		return errors.WithContext("error validating Calendar->Version: ", err)
	}
	if err := s.ProductID.valid(); err != nil {
		return errors.WithContext("error validating Calendar->ProductID: ", err)
	}
	for n := range s.Event {
		if err := s.Event[n].valid(); err != nil {
			return errors.WithContext("error validating Calendar->Event: ", err)
		}
	}
	for n := range s.Todo {
		if err := s.Todo[n].valid(); err != nil {
			return errors.WithContext("error validating Calendar->Todo: ", err)
		}
	}
	for n := range s.Journal {
		if err := s.Journal[n].valid(); err != nil {
			return errors.WithContext("error validating Calendar->Journal: ", err)
		}
	}
	for n := range s.FreeBusy {
		if err := s.FreeBusy[n].valid(); err != nil {
			return errors.WithContext("error validating Calendar->FreeBusy: ", err)
		}
	}
	for n := range s.Timezone {
		if err := s.Timezone[n].valid(); err != nil {
			return errors.WithContext("error validating Calendar->Timezone: ", err)
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
			return errors.WithContext("error decoding Event: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding Event: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VALARM":
				var e Alarm
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Event->Alarm: ", err)
				}
				s.Alarm = append(s.Alarm, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding Event: ", err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return errors.Error("error decoding Event: multiple DateTimeStamp")
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->DateTimeStamp: ", err)
			}
		case "UID":
			if requiredUID {
				return errors.Error("error decoding Event: multiple UID")
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->UID: ", err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return errors.Error("error decoding Event: multiple DateTimeStart")
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->DateTimeStart: ", err)
			}
		case "CLASS":
			if s.Class != nil {
				return errors.Error("error decoding Event: multiple Class")
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Class: ", err)
			}
		case "CREATED":
			if s.Created != nil {
				return errors.Error("error decoding Event: multiple Created")
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Created: ", err)
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return errors.Error("error decoding Event: multiple Description")
			}
			s.Description = new(PropDescription)
			if err := s.Description.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Description: ", err)
			}
		case "GEO":
			if s.Geo != nil {
				return errors.Error("error decoding Event: multiple Geo")
			}
			s.Geo = new(PropGeo)
			if err := s.Geo.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Geo: ", err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return errors.Error("error decoding Event: multiple LastModified")
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->LastModified: ", err)
			}
		case "LOCATION":
			if s.Location != nil {
				return errors.Error("error decoding Event: multiple Location")
			}
			s.Location = new(PropLocation)
			if err := s.Location.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Location: ", err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return errors.Error("error decoding Event: multiple Organizer")
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Organizer: ", err)
			}
		case "PRIORITY":
			if s.Priority != nil {
				return errors.Error("error decoding Event: multiple Priority")
			}
			s.Priority = new(PropPriority)
			if err := s.Priority.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Priority: ", err)
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return errors.Error("error decoding Event: multiple Sequence")
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Sequence: ", err)
			}
		case "STATUS":
			if s.Status != nil {
				return errors.Error("error decoding Event: multiple Status")
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Status: ", err)
			}
		case "SUMMARY":
			if s.Summary != nil {
				return errors.Error("error decoding Event: multiple Summary")
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Summary: ", err)
			}
		case "TRANSP":
			if s.TimeTransparency != nil {
				return errors.Error("error decoding Event: multiple TimeTransparency")
			}
			s.TimeTransparency = new(PropTimeTransparency)
			if err := s.TimeTransparency.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->TimeTransparency: ", err)
			}
		case "URL":
			if s.URL != nil {
				return errors.Error("error decoding Event: multiple URL")
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->URL: ", err)
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return errors.Error("error decoding Event: multiple RecurrenceID")
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->RecurrenceID: ", err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return errors.Error("error decoding Event: multiple RecurrenceRule")
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->RecurrenceRule: ", err)
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return errors.Error("error decoding Event: multiple DateTimeEnd")
			}
			s.DateTimeEnd = new(PropDateTimeEnd)
			if err := s.DateTimeEnd.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->DateTimeEnd: ", err)
			}
		case "DURATION":
			if s.Duration != nil {
				return errors.Error("error decoding Event: multiple Duration")
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Duration: ", err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Attachment: ", err)
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Attendee: ", err)
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Categories: ", err)
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Comment: ", err)
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Contact: ", err)
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->ExceptionDateTime: ", err)
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->RequestStatus: ", err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->RelatedTo: ", err)
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->Resources: ", err)
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Event->RecurrenceDateTimes: ", err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VEVENT" {
				return errors.WithContext("error decoding Event: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return errors.WithContext("error decoding Event: ", ErrMissingRequired)
	}
	if s.DateTimeEnd != nil && s.Duration != nil {
		return errors.WithContext("error decoding Event: ", ErrRequirementNotMet)
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
		return errors.WithContext("error validating Event->DateTimeStamp: ", err)
	}
	if err := s.UID.valid(); err != nil {
		return errors.WithContext("error validating Event->UID: ", err)
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return errors.WithContext("error validating Event->DateTimeStart: ", err)
		}
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return errors.WithContext("error validating Event->Class: ", err)
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return errors.WithContext("error validating Event->Created: ", err)
		}
	}
	if s.Description != nil {
		if err := s.Description.valid(); err != nil {
			return errors.WithContext("error validating Event->Description: ", err)
		}
	}
	if s.Geo != nil {
		if err := s.Geo.valid(); err != nil {
			return errors.WithContext("error validating Event->Geo: ", err)
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return errors.WithContext("error validating Event->LastModified: ", err)
		}
	}
	if s.Location != nil {
		if err := s.Location.valid(); err != nil {
			return errors.WithContext("error validating Event->Location: ", err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return errors.WithContext("error validating Event->Organizer: ", err)
		}
	}
	if s.Priority != nil {
		if err := s.Priority.valid(); err != nil {
			return errors.WithContext("error validating Event->Priority: ", err)
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
			return errors.WithContext("error validating Event->Sequence: ", err)
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return errors.WithContext("error validating Event->Status: ", err)
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return errors.WithContext("error validating Event->Summary: ", err)
		}
	}
	if s.TimeTransparency != nil {
		if err := s.TimeTransparency.valid(); err != nil {
			return errors.WithContext("error validating Event->TimeTransparency: ", err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return errors.WithContext("error validating Event->URL: ", err)
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return errors.WithContext("error validating Event->RecurrenceID: ", err)
		}
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return errors.WithContext("error validating Event->RecurrenceRule: ", err)
		}
	}
	if s.DateTimeEnd != nil {
		if err := s.DateTimeEnd.valid(); err != nil {
			return errors.WithContext("error validating Event->DateTimeEnd: ", err)
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return errors.WithContext("error validating Event->Duration: ", err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return errors.WithContext("error validating Event->Attachment: ", err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return errors.WithContext("error validating Event->Attendee: ", err)
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return errors.WithContext("error validating Event->Categories: ", err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return errors.WithContext("error validating Event->Comment: ", err)
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return errors.WithContext("error validating Event->Contact: ", err)
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return errors.WithContext("error validating Event->ExceptionDateTime: ", err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return errors.WithContext("error validating Event->RequestStatus: ", err)
		}
	}
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
			return errors.WithContext("error validating Event->RelatedTo: ", err)
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return errors.WithContext("error validating Event->Resources: ", err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return errors.WithContext("error validating Event->RecurrenceDateTimes: ", err)
		}
	}
	for n := range s.Alarm {
		if err := s.Alarm[n].valid(); err != nil {
			return errors.WithContext("error validating Event->Alarm: ", err)
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
			return errors.WithContext("error decoding Todo: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding Todo: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VALARM":
				var e Alarm
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Todo->Alarm: ", err)
				}
				s.Alarm = append(s.Alarm, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding Todo: ", err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return errors.Error("error decoding Todo: multiple DateTimeStamp")
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->DateTimeStamp: ", err)
			}
		case "UID":
			if requiredUID {
				return errors.Error("error decoding Todo: multiple UID")
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->UID: ", err)
			}
		case "CLASS":
			if s.Class != nil {
				return errors.Error("error decoding Todo: multiple Class")
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Class: ", err)
			}
		case "COMPLETED":
			if s.Completed != nil {
				return errors.Error("error decoding Todo: multiple Completed")
			}
			s.Completed = new(PropCompleted)
			if err := s.Completed.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Completed: ", err)
			}
		case "CREATED":
			if s.Created != nil {
				return errors.Error("error decoding Todo: multiple Created")
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Created: ", err)
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return errors.Error("error decoding Todo: multiple Description")
			}
			s.Description = new(PropDescription)
			if err := s.Description.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Description: ", err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return errors.Error("error decoding Todo: multiple DateTimeStart")
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->DateTimeStart: ", err)
			}
		case "GEO":
			if s.Geo != nil {
				return errors.Error("error decoding Todo: multiple Geo")
			}
			s.Geo = new(PropGeo)
			if err := s.Geo.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Geo: ", err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return errors.Error("error decoding Todo: multiple LastModified")
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->LastModified: ", err)
			}
		case "LOCATION":
			if s.Location != nil {
				return errors.Error("error decoding Todo: multiple Location")
			}
			s.Location = new(PropLocation)
			if err := s.Location.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Location: ", err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return errors.Error("error decoding Todo: multiple Organizer")
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Organizer: ", err)
			}
		case "PERCENT-COMPLETE":
			if s.PercentComplete != nil {
				return errors.Error("error decoding Todo: multiple PercentComplete")
			}
			s.PercentComplete = new(PropPercentComplete)
			if err := s.PercentComplete.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->PercentComplete: ", err)
			}
		case "PRIORITY":
			if s.Priority != nil {
				return errors.Error("error decoding Todo: multiple Priority")
			}
			s.Priority = new(PropPriority)
			if err := s.Priority.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Priority: ", err)
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return errors.Error("error decoding Todo: multiple RecurrenceID")
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->RecurrenceID: ", err)
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return errors.Error("error decoding Todo: multiple Sequence")
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Sequence: ", err)
			}
		case "STATUS":
			if s.Status != nil {
				return errors.Error("error decoding Todo: multiple Status")
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Status: ", err)
			}
		case "SUMMARY":
			if s.Summary != nil {
				return errors.Error("error decoding Todo: multiple Summary")
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Summary: ", err)
			}
		case "URL":
			if s.URL != nil {
				return errors.Error("error decoding Todo: multiple URL")
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->URL: ", err)
			}
		case "DUE":
			if s.Due != nil {
				return errors.Error("error decoding Todo: multiple Due")
			}
			s.Due = new(PropDue)
			if err := s.Due.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Due: ", err)
			}
		case "DURATION":
			if s.Duration != nil {
				return errors.Error("error decoding Todo: multiple Duration")
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Duration: ", err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Attachment: ", err)
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Attendee: ", err)
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Categories: ", err)
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Comment: ", err)
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Contact: ", err)
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->ExceptionDateTime: ", err)
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->RequestStatus: ", err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->RelatedTo: ", err)
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->Resources: ", err)
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Todo->RecurrenceDateTimes: ", err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VTODO" {
				return errors.WithContext("error decoding Todo: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return errors.WithContext("error decoding Todo: ", ErrMissingRequired)
	}
	if s.Duration != nil && (s.DateTimeStart == nil) {
		return errors.WithContext("error decoding Todo: ", ErrRequirementNotMet)
	}
	if s.Due != nil && s.Duration != nil {
		return errors.WithContext("error decoding Todo: ", ErrRequirementNotMet)
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
		return errors.WithContext("error validating Todo->DateTimeStamp: ", err)
	}
	if err := s.UID.valid(); err != nil {
		return errors.WithContext("error validating Todo->UID: ", err)
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return errors.WithContext("error validating Todo->Class: ", err)
		}
	}
	if s.Completed != nil {
		if err := s.Completed.valid(); err != nil {
			return errors.WithContext("error validating Todo->Completed: ", err)
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return errors.WithContext("error validating Todo->Created: ", err)
		}
	}
	if s.Description != nil {
		if err := s.Description.valid(); err != nil {
			return errors.WithContext("error validating Todo->Description: ", err)
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return errors.WithContext("error validating Todo->DateTimeStart: ", err)
		}
	}
	if s.Geo != nil {
		if err := s.Geo.valid(); err != nil {
			return errors.WithContext("error validating Todo->Geo: ", err)
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return errors.WithContext("error validating Todo->LastModified: ", err)
		}
	}
	if s.Location != nil {
		if err := s.Location.valid(); err != nil {
			return errors.WithContext("error validating Todo->Location: ", err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return errors.WithContext("error validating Todo->Organizer: ", err)
		}
	}
	if s.PercentComplete != nil {
		if err := s.PercentComplete.valid(); err != nil {
			return errors.WithContext("error validating Todo->PercentComplete: ", err)
		}
	}
	if s.Priority != nil {
		if err := s.Priority.valid(); err != nil {
			return errors.WithContext("error validating Todo->Priority: ", err)
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return errors.WithContext("error validating Todo->RecurrenceID: ", err)
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
			return errors.WithContext("error validating Todo->Sequence: ", err)
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return errors.WithContext("error validating Todo->Status: ", err)
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return errors.WithContext("error validating Todo->Summary: ", err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return errors.WithContext("error validating Todo->URL: ", err)
		}
	}
	if s.Due != nil {
		if err := s.Due.valid(); err != nil {
			return errors.WithContext("error validating Todo->Due: ", err)
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return errors.WithContext("error validating Todo->Duration: ", err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->Attachment: ", err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->Attendee: ", err)
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->Categories: ", err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->Comment: ", err)
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->Contact: ", err)
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->ExceptionDateTime: ", err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->RequestStatus: ", err)
		}
	}
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->RelatedTo: ", err)
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->Resources: ", err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->RecurrenceDateTimes: ", err)
		}
	}
	for n := range s.Alarm {
		if err := s.Alarm[n].valid(); err != nil {
			return errors.WithContext("error validating Todo->Alarm: ", err)
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
	Description         []PropDescription
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
			return errors.WithContext("error decoding Journal: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding Journal: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding Journal: ", err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return errors.Error("error decoding Journal: multiple DateTimeStamp")
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->DateTimeStamp: ", err)
			}
		case "UID":
			if requiredUID {
				return errors.Error("error decoding Journal: multiple UID")
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->UID: ", err)
			}
		case "CLASS":
			if s.Class != nil {
				return errors.Error("error decoding Journal: multiple Class")
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Class: ", err)
			}
		case "CREATED":
			if s.Created != nil {
				return errors.Error("error decoding Journal: multiple Created")
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Created: ", err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return errors.Error("error decoding Journal: multiple DateTimeStart")
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->DateTimeStart: ", err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return errors.Error("error decoding Journal: multiple LastModified")
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->LastModified: ", err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return errors.Error("error decoding Journal: multiple Organizer")
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Organizer: ", err)
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return errors.Error("error decoding Journal: multiple RecurrenceID")
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->RecurrenceID: ", err)
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return errors.Error("error decoding Journal: multiple Sequence")
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Sequence: ", err)
			}
		case "STATUS":
			if s.Status != nil {
				return errors.Error("error decoding Journal: multiple Status")
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Status: ", err)
			}
		case "SUMMARY":
			if s.Summary != nil {
				return errors.Error("error decoding Journal: multiple Summary")
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Summary: ", err)
			}
		case "URL":
			if s.URL != nil {
				return errors.Error("error decoding Journal: multiple URL")
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->URL: ", err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return errors.Error("error decoding Journal: multiple RecurrenceRule")
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->RecurrenceRule: ", err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Attachment: ", err)
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Attendee: ", err)
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Categories: ", err)
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Comment: ", err)
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Contact: ", err)
			}
			s.Contact = append(s.Contact, e)
		case "DESCRIPTION":
			var e PropDescription
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Description: ", err)
			}
			s.Description = append(s.Description, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->ExceptionDateTime: ", err)
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->RequestStatus: ", err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->RelatedTo: ", err)
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->Resources: ", err)
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Journal->RecurrenceDateTimes: ", err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VJOURNAL" {
				return errors.WithContext("error decoding Journal: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return errors.WithContext("error decoding Journal: ", ErrMissingRequired)
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
	for n := range s.Description {
		s.Description[n].encode(w)
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
		return errors.WithContext("error validating Journal->DateTimeStamp: ", err)
	}
	if err := s.UID.valid(); err != nil {
		return errors.WithContext("error validating Journal->UID: ", err)
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return errors.WithContext("error validating Journal->Class: ", err)
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return errors.WithContext("error validating Journal->Created: ", err)
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return errors.WithContext("error validating Journal->DateTimeStart: ", err)
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return errors.WithContext("error validating Journal->LastModified: ", err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return errors.WithContext("error validating Journal->Organizer: ", err)
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return errors.WithContext("error validating Journal->RecurrenceID: ", err)
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
			return errors.WithContext("error validating Journal->Sequence: ", err)
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return errors.WithContext("error validating Journal->Status: ", err)
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return errors.WithContext("error validating Journal->Summary: ", err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return errors.WithContext("error validating Journal->URL: ", err)
		}
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return errors.WithContext("error validating Journal->RecurrenceRule: ", err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->Attachment: ", err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->Attendee: ", err)
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->Categories: ", err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->Comment: ", err)
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->Contact: ", err)
		}
	}
	for n := range s.Description {
		if err := s.Description[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->Description: ", err)
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->ExceptionDateTime: ", err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->RequestStatus: ", err)
		}
	}
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->RelatedTo: ", err)
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->Resources: ", err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return errors.WithContext("error validating Journal->RecurrenceDateTimes: ", err)
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
			return errors.WithContext("error decoding FreeBusy: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding FreeBusy: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding FreeBusy: ", err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return errors.Error("error decoding FreeBusy: multiple DateTimeStamp")
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->DateTimeStamp: ", err)
			}
		case "UID":
			if requiredUID {
				return errors.Error("error decoding FreeBusy: multiple UID")
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->UID: ", err)
			}
		case "CONTACT":
			if s.Contact != nil {
				return errors.Error("error decoding FreeBusy: multiple Contact")
			}
			s.Contact = new(PropContact)
			if err := s.Contact.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->Contact: ", err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return errors.Error("error decoding FreeBusy: multiple DateTimeStart")
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->DateTimeStart: ", err)
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return errors.Error("error decoding FreeBusy: multiple DateTimeEnd")
			}
			s.DateTimeEnd = new(PropDateTimeEnd)
			if err := s.DateTimeEnd.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->DateTimeEnd: ", err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return errors.Error("error decoding FreeBusy: multiple Organizer")
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->Organizer: ", err)
			}
		case "URL":
			if s.URL != nil {
				return errors.Error("error decoding FreeBusy: multiple URL")
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->URL: ", err)
			}
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->Attendee: ", err)
			}
			s.Attendee = append(s.Attendee, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->Comment: ", err)
			}
			s.Comment = append(s.Comment, e)
		case "FREEBUSY":
			var e PropFreeBusy
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->FreeBusy: ", err)
			}
			s.FreeBusy = append(s.FreeBusy, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding FreeBusy->RequestStatus: ", err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "END":
			if value != "VFREEBUSY" {
				return errors.WithContext("error decoding FreeBusy: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return errors.WithContext("error decoding FreeBusy: ", ErrMissingRequired)
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
		return errors.WithContext("error validating FreeBusy->DateTimeStamp: ", err)
	}
	if err := s.UID.valid(); err != nil {
		return errors.WithContext("error validating FreeBusy->UID: ", err)
	}
	if s.Contact != nil {
		if err := s.Contact.valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->Contact: ", err)
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->DateTimeStart: ", err)
		}
	}
	if s.DateTimeEnd != nil {
		if err := s.DateTimeEnd.valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->DateTimeEnd: ", err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->Organizer: ", err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->URL: ", err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->Attendee: ", err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->Comment: ", err)
		}
	}
	for n := range s.FreeBusy {
		if err := s.FreeBusy[n].valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->FreeBusy: ", err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->RequestStatus: ", err)
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
			return errors.WithContext("error decoding Timezone: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding Timezone: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "STANDARD":
				var e Standard
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Timezone->Standard: ", err)
				}
				s.Standard = append(s.Standard, e)
			case "DAYLIGHT":
				var e Daylight
				if err := e.decode(t); err != nil {
					return errors.WithContext("error decoding Timezone->Daylight: ", err)
				}
				s.Daylight = append(s.Daylight, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding Timezone: ", err)
				}
			}
		case "TZID":
			if requiredTimezoneID {
				return errors.Error("error decoding Timezone: multiple TimezoneID")
			}
			requiredTimezoneID = true
			if err := s.TimezoneID.decode(params, value); err != nil {
				return errors.WithContext("error decoding Timezone->TimezoneID: ", err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return errors.Error("error decoding Timezone: multiple LastModified")
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return errors.WithContext("error decoding Timezone->LastModified: ", err)
			}
		case "TZURL":
			if s.TimezoneURL != nil {
				return errors.Error("error decoding Timezone: multiple TimezoneURL")
			}
			s.TimezoneURL = new(PropTimezoneURL)
			if err := s.TimezoneURL.decode(params, value); err != nil {
				return errors.WithContext("error decoding Timezone->TimezoneURL: ", err)
			}
		case "END":
			if value != "VTIMEZONE" {
				return errors.WithContext("error decoding Timezone: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredTimezoneID {
		return errors.WithContext("error decoding Timezone: ", ErrMissingRequired)
	}
	if s.Standard == nil && s.Daylight == nil {
		return errors.WithContext("error decoding Timezone: ", ErrRequirementNotMet)
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
		return errors.WithContext("error validating Timezone->TimezoneID: ", err)
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return errors.WithContext("error validating Timezone->LastModified: ", err)
		}
	}
	if s.TimezoneURL != nil {
		if err := s.TimezoneURL.valid(); err != nil {
			return errors.WithContext("error validating Timezone->TimezoneURL: ", err)
		}
	}
	for n := range s.Standard {
		if err := s.Standard[n].valid(); err != nil {
			return errors.WithContext("error validating Timezone->Standard: ", err)
		}
	}
	for n := range s.Daylight {
		if err := s.Daylight[n].valid(); err != nil {
			return errors.WithContext("error validating Timezone->Daylight: ", err)
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
			return errors.WithContext("error decoding Standard: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding Standard: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding Standard: ", err)
				}
			}
		case "DTSTART":
			if requiredDateTimeStart {
				return errors.Error("error decoding Standard: multiple DateTimeStart")
			}
			requiredDateTimeStart = true
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return errors.WithContext("error decoding Standard->DateTimeStart: ", err)
			}
		case "TZOFFSETTO":
			if requiredTimezoneOffsetTo {
				return errors.Error("error decoding Standard: multiple TimezoneOffsetTo")
			}
			requiredTimezoneOffsetTo = true
			if err := s.TimezoneOffsetTo.decode(params, value); err != nil {
				return errors.WithContext("error decoding Standard->TimezoneOffsetTo: ", err)
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return errors.Error("error decoding Standard: multiple TimezoneOffsetFrom")
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(params, value); err != nil {
				return errors.WithContext("error decoding Standard->TimezoneOffsetFrom: ", err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return errors.Error("error decoding Standard: multiple RecurrenceRule")
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return errors.WithContext("error decoding Standard->RecurrenceRule: ", err)
			}
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Standard->Comment: ", err)
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Standard->RecurrenceDateTimes: ", err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e PropTimezoneName
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Standard->TimezoneName: ", err)
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if value != "STANDARD" {
				return errors.WithContext("error decoding Standard: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffsetTo || !requiredTimezoneOffsetFrom {
		return errors.WithContext("error decoding Standard: ", ErrMissingRequired)
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
		return errors.WithContext("error validating Standard->DateTimeStart: ", err)
	}
	if err := s.TimezoneOffsetTo.valid(); err != nil {
		return errors.WithContext("error validating Standard->TimezoneOffsetTo: ", err)
	}
	if err := s.TimezoneOffsetFrom.valid(); err != nil {
		return errors.WithContext("error validating Standard->TimezoneOffsetFrom: ", err)
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return errors.WithContext("error validating Standard->RecurrenceRule: ", err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return errors.WithContext("error validating Standard->Comment: ", err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return errors.WithContext("error validating Standard->RecurrenceDateTimes: ", err)
		}
	}
	for n := range s.TimezoneName {
		if err := s.TimezoneName[n].valid(); err != nil {
			return errors.WithContext("error validating Standard->TimezoneName: ", err)
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
			return errors.WithContext("error decoding Daylight: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding Daylight: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding Daylight: ", err)
				}
			}
		case "DTSTART":
			if requiredDateTimeStart {
				return errors.Error("error decoding Daylight: multiple DateTimeStart")
			}
			requiredDateTimeStart = true
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return errors.WithContext("error decoding Daylight->DateTimeStart: ", err)
			}
		case "TZOFFSETTO":
			if requiredTimezoneOffsetTo {
				return errors.Error("error decoding Daylight: multiple TimezoneOffsetTo")
			}
			requiredTimezoneOffsetTo = true
			if err := s.TimezoneOffsetTo.decode(params, value); err != nil {
				return errors.WithContext("error decoding Daylight->TimezoneOffsetTo: ", err)
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return errors.Error("error decoding Daylight: multiple TimezoneOffsetFrom")
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(params, value); err != nil {
				return errors.WithContext("error decoding Daylight->TimezoneOffsetFrom: ", err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return errors.Error("error decoding Daylight: multiple RecurrenceRule")
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return errors.WithContext("error decoding Daylight->RecurrenceRule: ", err)
			}
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Daylight->Comment: ", err)
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Daylight->RecurrenceDateTimes: ", err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e PropTimezoneName
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding Daylight->TimezoneName: ", err)
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if value != "DAYLIGHT" {
				return errors.WithContext("error decoding Daylight: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffsetTo || !requiredTimezoneOffsetFrom {
		return errors.WithContext("error decoding Daylight: ", ErrMissingRequired)
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
		return errors.WithContext("error validating Daylight->DateTimeStart: ", err)
	}
	if err := s.TimezoneOffsetTo.valid(); err != nil {
		return errors.WithContext("error validating Daylight->TimezoneOffsetTo: ", err)
	}
	if err := s.TimezoneOffsetFrom.valid(); err != nil {
		return errors.WithContext("error validating Daylight->TimezoneOffsetFrom: ", err)
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return errors.WithContext("error validating Daylight->RecurrenceRule: ", err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return errors.WithContext("error validating Daylight->Comment: ", err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return errors.WithContext("error validating Daylight->RecurrenceDateTimes: ", err)
		}
	}
	for n := range s.TimezoneName {
		if err := s.TimezoneName[n].valid(); err != nil {
			return errors.WithContext("error validating Daylight->TimezoneName: ", err)
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
	UID        *PropUID
}

func (s *AlarmAudio) decode(t tokeniser) error {
	var requiredTrigger bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return errors.WithContext("error decoding AlarmAudio: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding AlarmAudio: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding AlarmAudio: ", err)
				}
			}
		case "TRIGGER":
			if requiredTrigger {
				return errors.Error("error decoding AlarmAudio: multiple Trigger")
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmAudio->Trigger: ", err)
			}
		case "DURATION":
			if s.Duration != nil {
				return errors.Error("error decoding AlarmAudio: multiple Duration")
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmAudio->Duration: ", err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return errors.Error("error decoding AlarmAudio: multiple Repeat")
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmAudio->Repeat: ", err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmAudio->Attachment: ", err)
			}
			s.Attachment = append(s.Attachment, e)
		case "UID":
			if s.UID != nil {
				return errors.Error("error decoding AlarmAudio: multiple UID")
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmAudio->UID: ", err)
			}
		case "END":
			if value != "VALARM" {
				return errors.WithContext("error decoding AlarmAudio: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredTrigger {
		return errors.WithContext("error decoding AlarmAudio: ", ErrMissingRequired)
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
	if s.UID != nil {
		s.UID.encode(w)
	}
}

func (s *AlarmAudio) valid() error {
	if err := s.Trigger.valid(); err != nil {
		return errors.WithContext("error validating AlarmAudio->Trigger: ", err)
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return errors.WithContext("error validating AlarmAudio->Duration: ", err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return errors.WithContext("error validating AlarmAudio->Repeat: ", err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return errors.WithContext("error validating AlarmAudio->Attachment: ", err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return errors.WithContext("error validating AlarmAudio->UID: ", err)
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
	UID         *PropUID
}

func (s *AlarmDisplay) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return errors.WithContext("error decoding AlarmDisplay: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding AlarmDisplay: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding AlarmDisplay: ", err)
				}
			}
		case "DESCRIPTION":
			if requiredDescription {
				return errors.Error("error decoding AlarmDisplay: multiple Description")
			}
			requiredDescription = true
			if err := s.Description.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmDisplay->Description: ", err)
			}
		case "TRIGGER":
			if requiredTrigger {
				return errors.Error("error decoding AlarmDisplay: multiple Trigger")
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmDisplay->Trigger: ", err)
			}
		case "DURATION":
			if s.Duration != nil {
				return errors.Error("error decoding AlarmDisplay: multiple Duration")
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmDisplay->Duration: ", err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return errors.Error("error decoding AlarmDisplay: multiple Repeat")
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmDisplay->Repeat: ", err)
			}
		case "UID":
			if s.UID != nil {
				return errors.Error("error decoding AlarmDisplay: multiple UID")
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmDisplay->UID: ", err)
			}
		case "END":
			if value != "VALARM" {
				return errors.WithContext("error decoding AlarmDisplay: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDescription || !requiredTrigger {
		return errors.WithContext("error decoding AlarmDisplay: ", ErrMissingRequired)
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
	if s.UID != nil {
		s.UID.encode(w)
	}
}

func (s *AlarmDisplay) valid() error {
	if err := s.Description.valid(); err != nil {
		return errors.WithContext("error validating AlarmDisplay->Description: ", err)
	}
	if err := s.Trigger.valid(); err != nil {
		return errors.WithContext("error validating AlarmDisplay->Trigger: ", err)
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return errors.WithContext("error validating AlarmDisplay->Duration: ", err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return errors.WithContext("error validating AlarmDisplay->Repeat: ", err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return errors.WithContext("error validating AlarmDisplay->UID: ", err)
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
	UID         *PropUID
}

func (s *AlarmEmail) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger, requiredSummary bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return errors.WithContext("error decoding AlarmEmail: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding AlarmEmail: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding AlarmEmail: ", err)
				}
			}
		case "DESCRIPTION":
			if requiredDescription {
				return errors.Error("error decoding AlarmEmail: multiple Description")
			}
			requiredDescription = true
			if err := s.Description.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmEmail->Description: ", err)
			}
		case "TRIGGER":
			if requiredTrigger {
				return errors.Error("error decoding AlarmEmail: multiple Trigger")
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmEmail->Trigger: ", err)
			}
		case "SUMMARY":
			if requiredSummary {
				return errors.Error("error decoding AlarmEmail: multiple Summary")
			}
			requiredSummary = true
			if err := s.Summary.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmEmail->Summary: ", err)
			}
		case "ATTENDEE":
			if s.Attendee != nil {
				return errors.Error("error decoding AlarmEmail: multiple Attendee")
			}
			s.Attendee = new(PropAttendee)
			if err := s.Attendee.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmEmail->Attendee: ", err)
			}
		case "DURATION":
			if s.Duration != nil {
				return errors.Error("error decoding AlarmEmail: multiple Duration")
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmEmail->Duration: ", err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return errors.Error("error decoding AlarmEmail: multiple Repeat")
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmEmail->Repeat: ", err)
			}
		case "UID":
			if s.UID != nil {
				return errors.Error("error decoding AlarmEmail: multiple UID")
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmEmail->UID: ", err)
			}
		case "END":
			if value != "VALARM" {
				return errors.WithContext("error decoding AlarmEmail: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDescription || !requiredTrigger || !requiredSummary {
		return errors.WithContext("error decoding AlarmEmail: ", ErrMissingRequired)
	}
	if t := s.Duration == nil; t == (s.Repeat == nil) {
		return errors.WithContext("error decoding AlarmEmail: ", ErrRequirementNotMet)
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
	if s.UID != nil {
		s.UID.encode(w)
	}
}

func (s *AlarmEmail) valid() error {
	if err := s.Description.valid(); err != nil {
		return errors.WithContext("error validating AlarmEmail->Description: ", err)
	}
	if err := s.Trigger.valid(); err != nil {
		return errors.WithContext("error validating AlarmEmail->Trigger: ", err)
	}
	if err := s.Summary.valid(); err != nil {
		return errors.WithContext("error validating AlarmEmail->Summary: ", err)
	}
	if s.Attendee != nil {
		if err := s.Attendee.valid(); err != nil {
			return errors.WithContext("error validating AlarmEmail->Attendee: ", err)
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return errors.WithContext("error validating AlarmEmail->Duration: ", err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return errors.WithContext("error validating AlarmEmail->Repeat: ", err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return errors.WithContext("error validating AlarmEmail->UID: ", err)
		}
	}
	return nil
}

// AlarmURI provies a group of components that define a URI Alarm
type AlarmURI struct {
	URI      PropURI
	Duration *PropDuration
	Repeat   *PropRepeat
	UID      *PropUID
}

func (s *AlarmURI) decode(t tokeniser) error {
	var requiredURI bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return errors.WithContext("error decoding AlarmURI: ", err)
		} else if p.Type == parser.PhraseDone {
			return errors.WithContext("error decoding AlarmURI: ", io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return errors.WithContext("error decoding AlarmURI: ", err)
				}
			}
		case "URI":
			if requiredURI {
				return errors.Error("error decoding AlarmURI: multiple URI")
			}
			requiredURI = true
			if err := s.URI.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmURI->URI: ", err)
			}
		case "DURATION":
			if s.Duration != nil {
				return errors.Error("error decoding AlarmURI: multiple Duration")
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmURI->Duration: ", err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return errors.Error("error decoding AlarmURI: multiple Repeat")
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmURI->Repeat: ", err)
			}
		case "UID":
			if s.UID != nil {
				return errors.Error("error decoding AlarmURI: multiple UID")
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return errors.WithContext("error decoding AlarmURI->UID: ", err)
			}
		case "END":
			if value != "VALARM" {
				return errors.WithContext("error decoding AlarmURI: ", ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredURI {
		return errors.WithContext("error decoding AlarmURI: ", ErrMissingRequired)
	}
	if t := s.Duration == nil; t == (s.Repeat == nil) {
		return errors.WithContext("error decoding AlarmURI: ", ErrRequirementNotMet)
	}
	return nil
}

func (s *AlarmURI) encode(w writer) {
	s.URI.encode(w)
	if s.Duration != nil {
		s.Duration.encode(w)
	}
	if s.Repeat != nil {
		s.Repeat.encode(w)
	}
	if s.UID != nil {
		s.UID.encode(w)
	}
}

func (s *AlarmURI) valid() error {
	if err := s.URI.valid(); err != nil {
		return errors.WithContext("error validating AlarmURI->URI: ", err)
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return errors.WithContext("error validating AlarmURI->Duration: ", err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return errors.WithContext("error validating AlarmURI->Repeat: ", err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return errors.WithContext("error validating AlarmURI->UID: ", err)
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
const (
	ErrInvalidEnd        errors.Error = "invalid end of section"
	ErrMissingRequired   errors.Error = "required property missing"
	ErrRequirementNotMet errors.Error = "requirement not met"
)
