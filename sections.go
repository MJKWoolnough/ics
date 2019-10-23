package ics

// File automatically generated with ./genSections.sh

import (
	"fmt"
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
			return fmt.Errorf(errDecodingType, cCalendar, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cCalendar, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VEVENT":
				var e Event
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cCalendar, cEvent, err)
				}
				s.Event = append(s.Event, e)
			case "VTODO":
				var e Todo
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cCalendar, cTodo, err)
				}
				s.Todo = append(s.Todo, e)
			case "VJOURNAL":
				var e Journal
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cCalendar, cJournal, err)
				}
				s.Journal = append(s.Journal, e)
			case "VFREEBUSY":
				var e FreeBusy
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cCalendar, cFreeBusy, err)
				}
				s.FreeBusy = append(s.FreeBusy, e)
			case "VTIMEZONE":
				var e Timezone
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cCalendar, cTimezone, err)
				}
				s.Timezone = append(s.Timezone, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cCalendar, err)
				}
			}
		case "VERSION":
			if requiredVersion {
				return fmt.Errorf(errMultiple, cCalendar, cVersion)
			}
			requiredVersion = true
			if err := s.Version.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cCalendar, cVersion, err)
			}
		case "PRODID":
			if requiredProductID {
				return fmt.Errorf(errMultiple, cCalendar, cProductID)
			}
			requiredProductID = true
			if err := s.ProductID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cCalendar, cProductID, err)
			}
		case "END":
			if value != "VCALENDAR" {
				return fmt.Errorf(errDecodingType, cCalendar, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredVersion || !requiredProductID {
		return fmt.Errorf(errDecodingType, cCalendar, ErrMissingRequired)
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
		return fmt.Errorf(errValidatingProp, cCalendar, cVersion, err)
	}
	if err := s.ProductID.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cCalendar, cProductID, err)
	}
	for n := range s.Event {
		if err := s.Event[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cCalendar, cEvent, err)
		}
	}
	for n := range s.Todo {
		if err := s.Todo[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cCalendar, cTodo, err)
		}
	}
	for n := range s.Journal {
		if err := s.Journal[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cCalendar, cJournal, err)
		}
	}
	for n := range s.FreeBusy {
		if err := s.FreeBusy[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cCalendar, cFreeBusy, err)
		}
	}
	for n := range s.Timezone {
		if err := s.Timezone[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cCalendar, cTimezone, err)
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
			return fmt.Errorf(errDecodingType, cEvent, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cEvent, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VALARM":
				var e Alarm
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cEvent, cAlarm, err)
				}
				s.Alarm = append(s.Alarm, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cEvent, err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return fmt.Errorf(errMultiple, cEvent, cDateTimeStamp)
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cDateTimeStamp, err)
			}
		case "UID":
			if requiredUID {
				return fmt.Errorf(errMultiple, cEvent, cUID)
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cUID, err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return fmt.Errorf(errMultiple, cEvent, cDateTimeStart)
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cDateTimeStart, err)
			}
		case "CLASS":
			if s.Class != nil {
				return fmt.Errorf(errMultiple, cEvent, cClass)
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cClass, err)
			}
		case "CREATED":
			if s.Created != nil {
				return fmt.Errorf(errMultiple, cEvent, cCreated)
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cCreated, err)
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return fmt.Errorf(errMultiple, cEvent, cDescription)
			}
			s.Description = new(PropDescription)
			if err := s.Description.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cDescription, err)
			}
		case "GEO":
			if s.Geo != nil {
				return fmt.Errorf(errMultiple, cEvent, cGeo)
			}
			s.Geo = new(PropGeo)
			if err := s.Geo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cGeo, err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return fmt.Errorf(errMultiple, cEvent, cLastModified)
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cLastModified, err)
			}
		case "LOCATION":
			if s.Location != nil {
				return fmt.Errorf(errMultiple, cEvent, cLocation)
			}
			s.Location = new(PropLocation)
			if err := s.Location.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cLocation, err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return fmt.Errorf(errMultiple, cEvent, cOrganizer)
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cOrganizer, err)
			}
		case "PRIORITY":
			if s.Priority != nil {
				return fmt.Errorf(errMultiple, cEvent, cPriority)
			}
			s.Priority = new(PropPriority)
			if err := s.Priority.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cPriority, err)
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return fmt.Errorf(errMultiple, cEvent, cSequence)
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cSequence, err)
			}
		case "STATUS":
			if s.Status != nil {
				return fmt.Errorf(errMultiple, cEvent, cStatus)
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cStatus, err)
			}
		case "SUMMARY":
			if s.Summary != nil {
				return fmt.Errorf(errMultiple, cEvent, cSummary)
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cSummary, err)
			}
		case "TRANSP":
			if s.TimeTransparency != nil {
				return fmt.Errorf(errMultiple, cEvent, cTimeTransparency)
			}
			s.TimeTransparency = new(PropTimeTransparency)
			if err := s.TimeTransparency.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cTimeTransparency, err)
			}
		case "URL":
			if s.URL != nil {
				return fmt.Errorf(errMultiple, cEvent, cURL)
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cURL, err)
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return fmt.Errorf(errMultiple, cEvent, cRecurrenceID)
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cRecurrenceID, err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return fmt.Errorf(errMultiple, cEvent, cRecurrenceRule)
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cRecurrenceRule, err)
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return fmt.Errorf(errMultiple, cEvent, cDateTimeEnd)
			}
			s.DateTimeEnd = new(PropDateTimeEnd)
			if err := s.DateTimeEnd.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cDateTimeEnd, err)
			}
		case "DURATION":
			if s.Duration != nil {
				return fmt.Errorf(errMultiple, cEvent, cDuration)
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cDuration, err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cAttachment, err)
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cAttendee, err)
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cCategories, err)
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cComment, err)
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cContact, err)
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cExceptionDateTime, err)
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cRequestStatus, err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cRelatedTo, err)
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cResources, err)
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cEvent, cRecurrenceDateTimes, err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VEVENT" {
				return fmt.Errorf(errDecodingType, cEvent, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return fmt.Errorf(errDecodingType, cEvent, ErrMissingRequired)
	}
	if s.DateTimeEnd != nil && s.Duration != nil {
		return fmt.Errorf(errDecodingType, cEvent, ErrRequirementNotMet)
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
		return fmt.Errorf(errValidatingProp, cEvent, cDateTimeStamp, err)
	}
	if err := s.UID.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cEvent, cUID, err)
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cDateTimeStart, err)
		}
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cClass, err)
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cCreated, err)
		}
	}
	if s.Description != nil {
		if err := s.Description.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cDescription, err)
		}
	}
	if s.Geo != nil {
		if err := s.Geo.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cGeo, err)
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cLastModified, err)
		}
	}
	if s.Location != nil {
		if err := s.Location.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cLocation, err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cOrganizer, err)
		}
	}
	if s.Priority != nil {
		if err := s.Priority.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cPriority, err)
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cSequence, err)
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cStatus, err)
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cSummary, err)
		}
	}
	if s.TimeTransparency != nil {
		if err := s.TimeTransparency.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cTimeTransparency, err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cURL, err)
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cRecurrenceID, err)
		}
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cRecurrenceRule, err)
		}
	}
	if s.DateTimeEnd != nil {
		if err := s.DateTimeEnd.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cDateTimeEnd, err)
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cDuration, err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cAttachment, err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cAttendee, err)
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cCategories, err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cComment, err)
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cContact, err)
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cExceptionDateTime, err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cRequestStatus, err)
		}
	}
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cRelatedTo, err)
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cResources, err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cRecurrenceDateTimes, err)
		}
	}
	for n := range s.Alarm {
		if err := s.Alarm[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cEvent, cAlarm, err)
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
			return fmt.Errorf(errDecodingType, cTodo, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cTodo, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "VALARM":
				var e Alarm
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cTodo, cAlarm, err)
				}
				s.Alarm = append(s.Alarm, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cTodo, err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return fmt.Errorf(errMultiple, cTodo, cDateTimeStamp)
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cDateTimeStamp, err)
			}
		case "UID":
			if requiredUID {
				return fmt.Errorf(errMultiple, cTodo, cUID)
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cUID, err)
			}
		case "CLASS":
			if s.Class != nil {
				return fmt.Errorf(errMultiple, cTodo, cClass)
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cClass, err)
			}
		case "COMPLETED":
			if s.Completed != nil {
				return fmt.Errorf(errMultiple, cTodo, cCompleted)
			}
			s.Completed = new(PropCompleted)
			if err := s.Completed.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cCompleted, err)
			}
		case "CREATED":
			if s.Created != nil {
				return fmt.Errorf(errMultiple, cTodo, cCreated)
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cCreated, err)
			}
		case "DESCRIPTION":
			if s.Description != nil {
				return fmt.Errorf(errMultiple, cTodo, cDescription)
			}
			s.Description = new(PropDescription)
			if err := s.Description.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cDescription, err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return fmt.Errorf(errMultiple, cTodo, cDateTimeStart)
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cDateTimeStart, err)
			}
		case "GEO":
			if s.Geo != nil {
				return fmt.Errorf(errMultiple, cTodo, cGeo)
			}
			s.Geo = new(PropGeo)
			if err := s.Geo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cGeo, err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return fmt.Errorf(errMultiple, cTodo, cLastModified)
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cLastModified, err)
			}
		case "LOCATION":
			if s.Location != nil {
				return fmt.Errorf(errMultiple, cTodo, cLocation)
			}
			s.Location = new(PropLocation)
			if err := s.Location.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cLocation, err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return fmt.Errorf(errMultiple, cTodo, cOrganizer)
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cOrganizer, err)
			}
		case "PERCENT-COMPLETE":
			if s.PercentComplete != nil {
				return fmt.Errorf(errMultiple, cTodo, cPercentComplete)
			}
			s.PercentComplete = new(PropPercentComplete)
			if err := s.PercentComplete.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cPercentComplete, err)
			}
		case "PRIORITY":
			if s.Priority != nil {
				return fmt.Errorf(errMultiple, cTodo, cPriority)
			}
			s.Priority = new(PropPriority)
			if err := s.Priority.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cPriority, err)
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return fmt.Errorf(errMultiple, cTodo, cRecurrenceID)
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cRecurrenceID, err)
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return fmt.Errorf(errMultiple, cTodo, cSequence)
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cSequence, err)
			}
		case "STATUS":
			if s.Status != nil {
				return fmt.Errorf(errMultiple, cTodo, cStatus)
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cStatus, err)
			}
		case "SUMMARY":
			if s.Summary != nil {
				return fmt.Errorf(errMultiple, cTodo, cSummary)
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cSummary, err)
			}
		case "URL":
			if s.URL != nil {
				return fmt.Errorf(errMultiple, cTodo, cURL)
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cURL, err)
			}
		case "DUE":
			if s.Due != nil {
				return fmt.Errorf(errMultiple, cTodo, cDue)
			}
			s.Due = new(PropDue)
			if err := s.Due.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cDue, err)
			}
		case "DURATION":
			if s.Duration != nil {
				return fmt.Errorf(errMultiple, cTodo, cDuration)
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cDuration, err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cAttachment, err)
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cAttendee, err)
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cCategories, err)
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cComment, err)
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cContact, err)
			}
			s.Contact = append(s.Contact, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cExceptionDateTime, err)
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cRequestStatus, err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cRelatedTo, err)
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cResources, err)
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTodo, cRecurrenceDateTimes, err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VTODO" {
				return fmt.Errorf(errDecodingType, cTodo, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return fmt.Errorf(errDecodingType, cTodo, ErrMissingRequired)
	}
	if s.Duration != nil && (s.DateTimeStart == nil) {
		return fmt.Errorf(errDecodingType, cTodo, ErrRequirementNotMet)
	}
	if s.Due != nil && s.Duration != nil {
		return fmt.Errorf(errDecodingType, cTodo, ErrRequirementNotMet)
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
		return fmt.Errorf(errValidatingProp, cTodo, cDateTimeStamp, err)
	}
	if err := s.UID.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cTodo, cUID, err)
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cClass, err)
		}
	}
	if s.Completed != nil {
		if err := s.Completed.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cCompleted, err)
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cCreated, err)
		}
	}
	if s.Description != nil {
		if err := s.Description.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cDescription, err)
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cDateTimeStart, err)
		}
	}
	if s.Geo != nil {
		if err := s.Geo.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cGeo, err)
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cLastModified, err)
		}
	}
	if s.Location != nil {
		if err := s.Location.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cLocation, err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cOrganizer, err)
		}
	}
	if s.PercentComplete != nil {
		if err := s.PercentComplete.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cPercentComplete, err)
		}
	}
	if s.Priority != nil {
		if err := s.Priority.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cPriority, err)
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cRecurrenceID, err)
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cSequence, err)
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cStatus, err)
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cSummary, err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cURL, err)
		}
	}
	if s.Due != nil {
		if err := s.Due.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cDue, err)
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cDuration, err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cAttachment, err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cAttendee, err)
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cCategories, err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cComment, err)
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cContact, err)
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cExceptionDateTime, err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cRequestStatus, err)
		}
	}
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cRelatedTo, err)
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cResources, err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cRecurrenceDateTimes, err)
		}
	}
	for n := range s.Alarm {
		if err := s.Alarm[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTodo, cAlarm, err)
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
			return fmt.Errorf(errDecodingType, cJournal, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cJournal, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cJournal, err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return fmt.Errorf(errMultiple, cJournal, cDateTimeStamp)
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cDateTimeStamp, err)
			}
		case "UID":
			if requiredUID {
				return fmt.Errorf(errMultiple, cJournal, cUID)
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cUID, err)
			}
		case "CLASS":
			if s.Class != nil {
				return fmt.Errorf(errMultiple, cJournal, cClass)
			}
			s.Class = new(PropClass)
			if err := s.Class.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cClass, err)
			}
		case "CREATED":
			if s.Created != nil {
				return fmt.Errorf(errMultiple, cJournal, cCreated)
			}
			s.Created = new(PropCreated)
			if err := s.Created.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cCreated, err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return fmt.Errorf(errMultiple, cJournal, cDateTimeStart)
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cDateTimeStart, err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return fmt.Errorf(errMultiple, cJournal, cLastModified)
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cLastModified, err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return fmt.Errorf(errMultiple, cJournal, cOrganizer)
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cOrganizer, err)
			}
		case "RECURRENCE-ID":
			if s.RecurrenceID != nil {
				return fmt.Errorf(errMultiple, cJournal, cRecurrenceID)
			}
			s.RecurrenceID = new(PropRecurrenceID)
			if err := s.RecurrenceID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cRecurrenceID, err)
			}
		case "SEQUENCE":
			if s.Sequence != nil {
				return fmt.Errorf(errMultiple, cJournal, cSequence)
			}
			s.Sequence = new(PropSequence)
			if err := s.Sequence.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cSequence, err)
			}
		case "STATUS":
			if s.Status != nil {
				return fmt.Errorf(errMultiple, cJournal, cStatus)
			}
			s.Status = new(PropStatus)
			if err := s.Status.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cStatus, err)
			}
		case "SUMMARY":
			if s.Summary != nil {
				return fmt.Errorf(errMultiple, cJournal, cSummary)
			}
			s.Summary = new(PropSummary)
			if err := s.Summary.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cSummary, err)
			}
		case "URL":
			if s.URL != nil {
				return fmt.Errorf(errMultiple, cJournal, cURL)
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cURL, err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return fmt.Errorf(errMultiple, cJournal, cRecurrenceRule)
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cRecurrenceRule, err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cAttachment, err)
			}
			s.Attachment = append(s.Attachment, e)
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cAttendee, err)
			}
			s.Attendee = append(s.Attendee, e)
		case "CATEGORIES":
			var e PropCategories
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cCategories, err)
			}
			s.Categories = append(s.Categories, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cComment, err)
			}
			s.Comment = append(s.Comment, e)
		case "CONTACT":
			var e PropContact
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cContact, err)
			}
			s.Contact = append(s.Contact, e)
		case "DESCRIPTION":
			var e PropDescription
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cDescription, err)
			}
			s.Description = append(s.Description, e)
		case "EXDATE":
			var e PropExceptionDateTime
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cExceptionDateTime, err)
			}
			s.ExceptionDateTime = append(s.ExceptionDateTime, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cRequestStatus, err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "RELATED-TO":
			var e PropRelatedTo
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cRelatedTo, err)
			}
			s.RelatedTo = append(s.RelatedTo, e)
		case "RESOURCES":
			var e PropResources
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cResources, err)
			}
			s.Resources = append(s.Resources, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cJournal, cRecurrenceDateTimes, err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "END":
			if value != "VJOURNAL" {
				return fmt.Errorf(errDecodingType, cJournal, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return fmt.Errorf(errDecodingType, cJournal, ErrMissingRequired)
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
		return fmt.Errorf(errValidatingProp, cJournal, cDateTimeStamp, err)
	}
	if err := s.UID.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cJournal, cUID, err)
	}
	if s.Class != nil {
		if err := s.Class.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cClass, err)
		}
	}
	if s.Created != nil {
		if err := s.Created.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cCreated, err)
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cDateTimeStart, err)
		}
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cLastModified, err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cOrganizer, err)
		}
	}
	if s.RecurrenceID != nil {
		if err := s.RecurrenceID.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cRecurrenceID, err)
		}
	}
	if s.Sequence != nil {
		if err := s.Sequence.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cSequence, err)
		}
	}
	if s.Status != nil {
		if err := s.Status.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cStatus, err)
		}
	}
	if s.Summary != nil {
		if err := s.Summary.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cSummary, err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cURL, err)
		}
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cRecurrenceRule, err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cAttachment, err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cAttendee, err)
		}
	}
	for n := range s.Categories {
		if err := s.Categories[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cCategories, err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cComment, err)
		}
	}
	for n := range s.Contact {
		if err := s.Contact[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cContact, err)
		}
	}
	for n := range s.Description {
		if err := s.Description[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cDescription, err)
		}
	}
	for n := range s.ExceptionDateTime {
		if err := s.ExceptionDateTime[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cExceptionDateTime, err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cRequestStatus, err)
		}
	}
	for n := range s.RelatedTo {
		if err := s.RelatedTo[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cRelatedTo, err)
		}
	}
	for n := range s.Resources {
		if err := s.Resources[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cResources, err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cJournal, cRecurrenceDateTimes, err)
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
			return fmt.Errorf(errDecodingType, cFreeBusy, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cFreeBusy, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cFreeBusy, err)
				}
			}
		case "DTSTAMP":
			if requiredDateTimeStamp {
				return fmt.Errorf(errMultiple, cFreeBusy, cDateTimeStamp)
			}
			requiredDateTimeStamp = true
			if err := s.DateTimeStamp.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cDateTimeStamp, err)
			}
		case "UID":
			if requiredUID {
				return fmt.Errorf(errMultiple, cFreeBusy, cUID)
			}
			requiredUID = true
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cUID, err)
			}
		case "CONTACT":
			if s.Contact != nil {
				return fmt.Errorf(errMultiple, cFreeBusy, cContact)
			}
			s.Contact = new(PropContact)
			if err := s.Contact.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cContact, err)
			}
		case "DTSTART":
			if s.DateTimeStart != nil {
				return fmt.Errorf(errMultiple, cFreeBusy, cDateTimeStart)
			}
			s.DateTimeStart = new(PropDateTimeStart)
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cDateTimeStart, err)
			}
		case "DTEND":
			if s.DateTimeEnd != nil {
				return fmt.Errorf(errMultiple, cFreeBusy, cDateTimeEnd)
			}
			s.DateTimeEnd = new(PropDateTimeEnd)
			if err := s.DateTimeEnd.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cDateTimeEnd, err)
			}
		case "ORGANIZER":
			if s.Organizer != nil {
				return fmt.Errorf(errMultiple, cFreeBusy, cOrganizer)
			}
			s.Organizer = new(PropOrganizer)
			if err := s.Organizer.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cOrganizer, err)
			}
		case "URL":
			if s.URL != nil {
				return fmt.Errorf(errMultiple, cFreeBusy, cURL)
			}
			s.URL = new(PropURL)
			if err := s.URL.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cURL, err)
			}
		case "ATTENDEE":
			var e PropAttendee
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cAttendee, err)
			}
			s.Attendee = append(s.Attendee, e)
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cComment, err)
			}
			s.Comment = append(s.Comment, e)
		case "FREEBUSY":
			var e PropFreeBusy
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cFreeBusy, err)
			}
			s.FreeBusy = append(s.FreeBusy, e)
		case "REQUEST-STATUS":
			var e PropRequestStatus
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cFreeBusy, cRequestStatus, err)
			}
			s.RequestStatus = append(s.RequestStatus, e)
		case "END":
			if value != "VFREEBUSY" {
				return fmt.Errorf(errDecodingType, cFreeBusy, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStamp || !requiredUID {
		return fmt.Errorf(errDecodingType, cFreeBusy, ErrMissingRequired)
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
		return fmt.Errorf(errValidatingProp, cFreeBusy, cDateTimeStamp, err)
	}
	if err := s.UID.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cFreeBusy, cUID, err)
	}
	if s.Contact != nil {
		if err := s.Contact.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cContact, err)
		}
	}
	if s.DateTimeStart != nil {
		if err := s.DateTimeStart.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cDateTimeStart, err)
		}
	}
	if s.DateTimeEnd != nil {
		if err := s.DateTimeEnd.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cDateTimeEnd, err)
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cOrganizer, err)
		}
	}
	if s.URL != nil {
		if err := s.URL.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cURL, err)
		}
	}
	for n := range s.Attendee {
		if err := s.Attendee[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cAttendee, err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cComment, err)
		}
	}
	for n := range s.FreeBusy {
		if err := s.FreeBusy[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cFreeBusy, err)
		}
	}
	for n := range s.RequestStatus {
		if err := s.RequestStatus[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cFreeBusy, cRequestStatus, err)
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
			return fmt.Errorf(errDecodingType, cTimezone, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cTimezone, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			case "STANDARD":
				var e Standard
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cTimezone, cStandard, err)
				}
				s.Standard = append(s.Standard, e)
			case "DAYLIGHT":
				var e Daylight
				if err := e.decode(t); err != nil {
					return fmt.Errorf(errDecodingProp, cTimezone, cDaylight, err)
				}
				s.Daylight = append(s.Daylight, e)
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cTimezone, err)
				}
			}
		case "TZID":
			if requiredTimezoneID {
				return fmt.Errorf(errMultiple, cTimezone, cTimezoneID)
			}
			requiredTimezoneID = true
			if err := s.TimezoneID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTimezone, cTimezoneID, err)
			}
		case "LAST-MOD":
			if s.LastModified != nil {
				return fmt.Errorf(errMultiple, cTimezone, cLastModified)
			}
			s.LastModified = new(PropLastModified)
			if err := s.LastModified.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTimezone, cLastModified, err)
			}
		case "TZURL":
			if s.TimezoneURL != nil {
				return fmt.Errorf(errMultiple, cTimezone, cTimezoneURL)
			}
			s.TimezoneURL = new(PropTimezoneURL)
			if err := s.TimezoneURL.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cTimezone, cTimezoneURL, err)
			}
		case "END":
			if value != "VTIMEZONE" {
				return fmt.Errorf(errDecodingType, cTimezone, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredTimezoneID {
		return fmt.Errorf(errDecodingType, cTimezone, ErrMissingRequired)
	}
	if s.Standard == nil && s.Daylight == nil {
		return fmt.Errorf(errDecodingType, cTimezone, ErrRequirementNotMet)
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
		return fmt.Errorf(errValidatingProp, cTimezone, cTimezoneID, err)
	}
	if s.LastModified != nil {
		if err := s.LastModified.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTimezone, cLastModified, err)
		}
	}
	if s.TimezoneURL != nil {
		if err := s.TimezoneURL.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTimezone, cTimezoneURL, err)
		}
	}
	for n := range s.Standard {
		if err := s.Standard[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTimezone, cStandard, err)
		}
	}
	for n := range s.Daylight {
		if err := s.Daylight[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cTimezone, cDaylight, err)
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
			return fmt.Errorf(errDecodingType, cStandard, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cStandard, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cStandard, err)
				}
			}
		case "DTSTART":
			if requiredDateTimeStart {
				return fmt.Errorf(errMultiple, cStandard, cDateTimeStart)
			}
			requiredDateTimeStart = true
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cStandard, cDateTimeStart, err)
			}
		case "TZOFFSETTO":
			if requiredTimezoneOffsetTo {
				return fmt.Errorf(errMultiple, cStandard, cTimezoneOffsetTo)
			}
			requiredTimezoneOffsetTo = true
			if err := s.TimezoneOffsetTo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cStandard, cTimezoneOffsetTo, err)
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return fmt.Errorf(errMultiple, cStandard, cTimezoneOffsetFrom)
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cStandard, cTimezoneOffsetFrom, err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return fmt.Errorf(errMultiple, cStandard, cRecurrenceRule)
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cStandard, cRecurrenceRule, err)
			}
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cStandard, cComment, err)
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cStandard, cRecurrenceDateTimes, err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e PropTimezoneName
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cStandard, cTimezoneName, err)
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if value != "STANDARD" {
				return fmt.Errorf(errDecodingType, cStandard, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffsetTo || !requiredTimezoneOffsetFrom {
		return fmt.Errorf(errDecodingType, cStandard, ErrMissingRequired)
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
		return fmt.Errorf(errValidatingProp, cStandard, cDateTimeStart, err)
	}
	if err := s.TimezoneOffsetTo.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cStandard, cTimezoneOffsetTo, err)
	}
	if err := s.TimezoneOffsetFrom.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cStandard, cTimezoneOffsetFrom, err)
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cStandard, cRecurrenceRule, err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cStandard, cComment, err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cStandard, cRecurrenceDateTimes, err)
		}
	}
	for n := range s.TimezoneName {
		if err := s.TimezoneName[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cStandard, cTimezoneName, err)
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
			return fmt.Errorf(errDecodingType, cDaylight, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cDaylight, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cDaylight, err)
				}
			}
		case "DTSTART":
			if requiredDateTimeStart {
				return fmt.Errorf(errMultiple, cDaylight, cDateTimeStart)
			}
			requiredDateTimeStart = true
			if err := s.DateTimeStart.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cDaylight, cDateTimeStart, err)
			}
		case "TZOFFSETTO":
			if requiredTimezoneOffsetTo {
				return fmt.Errorf(errMultiple, cDaylight, cTimezoneOffsetTo)
			}
			requiredTimezoneOffsetTo = true
			if err := s.TimezoneOffsetTo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cDaylight, cTimezoneOffsetTo, err)
			}
		case "TZOFFSETFROM":
			if requiredTimezoneOffsetFrom {
				return fmt.Errorf(errMultiple, cDaylight, cTimezoneOffsetFrom)
			}
			requiredTimezoneOffsetFrom = true
			if err := s.TimezoneOffsetFrom.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cDaylight, cTimezoneOffsetFrom, err)
			}
		case "RRULE":
			if s.RecurrenceRule != nil {
				return fmt.Errorf(errMultiple, cDaylight, cRecurrenceRule)
			}
			s.RecurrenceRule = new(PropRecurrenceRule)
			if err := s.RecurrenceRule.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cDaylight, cRecurrenceRule, err)
			}
		case "COMMENT":
			var e PropComment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cDaylight, cComment, err)
			}
			s.Comment = append(s.Comment, e)
		case "RDATE":
			var e PropRecurrenceDateTimes
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cDaylight, cRecurrenceDateTimes, err)
			}
			s.RecurrenceDateTimes = append(s.RecurrenceDateTimes, e)
		case "TZNAME":
			var e PropTimezoneName
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cDaylight, cTimezoneName, err)
			}
			s.TimezoneName = append(s.TimezoneName, e)
		case "END":
			if value != "DAYLIGHT" {
				return fmt.Errorf(errDecodingType, cDaylight, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDateTimeStart || !requiredTimezoneOffsetTo || !requiredTimezoneOffsetFrom {
		return fmt.Errorf(errDecodingType, cDaylight, ErrMissingRequired)
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
		return fmt.Errorf(errValidatingProp, cDaylight, cDateTimeStart, err)
	}
	if err := s.TimezoneOffsetTo.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cDaylight, cTimezoneOffsetTo, err)
	}
	if err := s.TimezoneOffsetFrom.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cDaylight, cTimezoneOffsetFrom, err)
	}
	if s.RecurrenceRule != nil {
		if err := s.RecurrenceRule.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cDaylight, cRecurrenceRule, err)
		}
	}
	for n := range s.Comment {
		if err := s.Comment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cDaylight, cComment, err)
		}
	}
	for n := range s.RecurrenceDateTimes {
		if err := s.RecurrenceDateTimes[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cDaylight, cRecurrenceDateTimes, err)
		}
	}
	for n := range s.TimezoneName {
		if err := s.TimezoneName[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cDaylight, cTimezoneName, err)
		}
	}
	return nil
}

// AlarmAudio provides a group of components that define an Audio Alarm
type AlarmAudio struct {
	Trigger       PropTrigger
	Duration      *PropDuration
	Repeat        *PropRepeat
	Attachment    []PropAttachment
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}

func (s *AlarmAudio) decode(t tokeniser) error {
	var requiredTrigger bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return fmt.Errorf(errDecodingType, cAlarmAudio, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cAlarmAudio, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cAlarmAudio, err)
				}
			}
		case "TRIGGER":
			if requiredTrigger {
				return fmt.Errorf(errMultiple, cAlarmAudio, cTrigger)
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cTrigger, err)
			}
		case "DURATION":
			if s.Duration != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cDuration)
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cDuration, err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cRepeat)
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cRepeat, err)
			}
		case "ATTACH":
			var e PropAttachment
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cAttachment, err)
			}
			s.Attachment = append(s.Attachment, e)
		case "UID":
			if s.UID != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cUID)
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cUID, err)
			}
		case "ALARM-AGENT":
			var e PropAlarmAgent
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cAlarmAgent, err)
			}
			s.AlarmAgent = append(s.AlarmAgent, e)
		case "STATUS":
			if s.AlarmStatus != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cAlarmStatus)
			}
			s.AlarmStatus = new(PropAlarmStatus)
			if err := s.AlarmStatus.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cAlarmStatus, err)
			}
		case "LAST-TRIGGERED":
			var e PropLastTriggered
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cLastTriggered, err)
			}
			s.LastTriggered = append(s.LastTriggered, e)
		case "ACKNOWLEDGED":
			if s.Acknowledged != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cAcknowledged)
			}
			s.Acknowledged = new(PropAcknowledged)
			if err := s.Acknowledged.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cAcknowledged, err)
			}
		case "PROXIMITY":
			if s.Proximity != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cProximity)
			}
			s.Proximity = new(PropProximity)
			if err := s.Proximity.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cProximity, err)
			}
		case "GEO-LOCATION":
			var e PropGeoLocation
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cGeoLocation, err)
			}
			s.GeoLocation = append(s.GeoLocation, e)
		case "RELATED-TO":
			if s.RelatedTo != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cRelatedTo)
			}
			s.RelatedTo = new(PropRelatedTo)
			if err := s.RelatedTo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cRelatedTo, err)
			}
		case "DEFAULT-ALARM":
			if s.DefaultAlarm != nil {
				return fmt.Errorf(errMultiple, cAlarmAudio, cDefaultAlarm)
			}
			s.DefaultAlarm = new(PropDefaultAlarm)
			if err := s.DefaultAlarm.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmAudio, cDefaultAlarm, err)
			}
		case "END":
			if value != "VALARM" {
				return fmt.Errorf(errDecodingType, cAlarmAudio, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredTrigger {
		return fmt.Errorf(errDecodingType, cAlarmAudio, ErrMissingRequired)
	}
	if s.GeoLocation != nil && (s.Proximity == nil) {
		return fmt.Errorf(errDecodingType, cAlarmAudio, ErrRequirementNotMet)
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
	for n := range s.AlarmAgent {
		s.AlarmAgent[n].encode(w)
	}
	if s.AlarmStatus != nil {
		s.AlarmStatus.encode(w)
	}
	for n := range s.LastTriggered {
		s.LastTriggered[n].encode(w)
	}
	if s.Acknowledged != nil {
		s.Acknowledged.encode(w)
	}
	if s.Proximity != nil {
		s.Proximity.encode(w)
	}
	for n := range s.GeoLocation {
		s.GeoLocation[n].encode(w)
	}
	if s.RelatedTo != nil {
		s.RelatedTo.encode(w)
	}
	if s.DefaultAlarm != nil {
		s.DefaultAlarm.encode(w)
	}
}

func (s *AlarmAudio) valid() error {
	if err := s.Trigger.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cAlarmAudio, cTrigger, err)
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cDuration, err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cRepeat, err)
		}
	}
	for n := range s.Attachment {
		if err := s.Attachment[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cAttachment, err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cUID, err)
		}
	}
	for n := range s.AlarmAgent {
		if err := s.AlarmAgent[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cAlarmAgent, err)
		}
	}
	if s.AlarmStatus != nil {
		if err := s.AlarmStatus.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cAlarmStatus, err)
		}
	}
	for n := range s.LastTriggered {
		if err := s.LastTriggered[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cLastTriggered, err)
		}
	}
	if s.Acknowledged != nil {
		if err := s.Acknowledged.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cAcknowledged, err)
		}
	}
	if s.Proximity != nil {
		if err := s.Proximity.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cProximity, err)
		}
	}
	for n := range s.GeoLocation {
		if err := s.GeoLocation[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cGeoLocation, err)
		}
	}
	if s.RelatedTo != nil {
		if err := s.RelatedTo.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cRelatedTo, err)
		}
	}
	if s.DefaultAlarm != nil {
		if err := s.DefaultAlarm.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmAudio, cDefaultAlarm, err)
		}
	}
	return nil
}

// AlarmDisplay provides a group of components that define a Display Alarm
type AlarmDisplay struct {
	Description   PropDescription
	Trigger       PropTrigger
	Duration      *PropDuration
	Repeat        *PropRepeat
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}

func (s *AlarmDisplay) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return fmt.Errorf(errDecodingType, cAlarmDisplay, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cAlarmDisplay, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cAlarmDisplay, err)
				}
			}
		case "DESCRIPTION":
			if requiredDescription {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cDescription)
			}
			requiredDescription = true
			if err := s.Description.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cDescription, err)
			}
		case "TRIGGER":
			if requiredTrigger {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cTrigger)
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cTrigger, err)
			}
		case "DURATION":
			if s.Duration != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cDuration)
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cDuration, err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cRepeat)
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cRepeat, err)
			}
		case "UID":
			if s.UID != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cUID)
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cUID, err)
			}
		case "ALARM-AGENT":
			var e PropAlarmAgent
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cAlarmAgent, err)
			}
			s.AlarmAgent = append(s.AlarmAgent, e)
		case "STATUS":
			if s.AlarmStatus != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cAlarmStatus)
			}
			s.AlarmStatus = new(PropAlarmStatus)
			if err := s.AlarmStatus.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cAlarmStatus, err)
			}
		case "LAST-TRIGGERED":
			var e PropLastTriggered
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cLastTriggered, err)
			}
			s.LastTriggered = append(s.LastTriggered, e)
		case "ACKNOWLEDGED":
			if s.Acknowledged != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cAcknowledged)
			}
			s.Acknowledged = new(PropAcknowledged)
			if err := s.Acknowledged.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cAcknowledged, err)
			}
		case "PROXIMITY":
			if s.Proximity != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cProximity)
			}
			s.Proximity = new(PropProximity)
			if err := s.Proximity.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cProximity, err)
			}
		case "GEO-LOCATION":
			var e PropGeoLocation
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cGeoLocation, err)
			}
			s.GeoLocation = append(s.GeoLocation, e)
		case "RELATED-TO":
			if s.RelatedTo != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cRelatedTo)
			}
			s.RelatedTo = new(PropRelatedTo)
			if err := s.RelatedTo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cRelatedTo, err)
			}
		case "DEFAULT-ALARM":
			if s.DefaultAlarm != nil {
				return fmt.Errorf(errMultiple, cAlarmDisplay, cDefaultAlarm)
			}
			s.DefaultAlarm = new(PropDefaultAlarm)
			if err := s.DefaultAlarm.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmDisplay, cDefaultAlarm, err)
			}
		case "END":
			if value != "VALARM" {
				return fmt.Errorf(errDecodingType, cAlarmDisplay, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDescription || !requiredTrigger {
		return fmt.Errorf(errDecodingType, cAlarmDisplay, ErrMissingRequired)
	}
	if s.GeoLocation != nil && (s.Proximity == nil) {
		return fmt.Errorf(errDecodingType, cAlarmDisplay, ErrRequirementNotMet)
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
	for n := range s.AlarmAgent {
		s.AlarmAgent[n].encode(w)
	}
	if s.AlarmStatus != nil {
		s.AlarmStatus.encode(w)
	}
	for n := range s.LastTriggered {
		s.LastTriggered[n].encode(w)
	}
	if s.Acknowledged != nil {
		s.Acknowledged.encode(w)
	}
	if s.Proximity != nil {
		s.Proximity.encode(w)
	}
	for n := range s.GeoLocation {
		s.GeoLocation[n].encode(w)
	}
	if s.RelatedTo != nil {
		s.RelatedTo.encode(w)
	}
	if s.DefaultAlarm != nil {
		s.DefaultAlarm.encode(w)
	}
}

func (s *AlarmDisplay) valid() error {
	if err := s.Description.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cAlarmDisplay, cDescription, err)
	}
	if err := s.Trigger.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cAlarmDisplay, cTrigger, err)
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cDuration, err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cRepeat, err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cUID, err)
		}
	}
	for n := range s.AlarmAgent {
		if err := s.AlarmAgent[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cAlarmAgent, err)
		}
	}
	if s.AlarmStatus != nil {
		if err := s.AlarmStatus.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cAlarmStatus, err)
		}
	}
	for n := range s.LastTriggered {
		if err := s.LastTriggered[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cLastTriggered, err)
		}
	}
	if s.Acknowledged != nil {
		if err := s.Acknowledged.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cAcknowledged, err)
		}
	}
	if s.Proximity != nil {
		if err := s.Proximity.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cProximity, err)
		}
	}
	for n := range s.GeoLocation {
		if err := s.GeoLocation[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cGeoLocation, err)
		}
	}
	if s.RelatedTo != nil {
		if err := s.RelatedTo.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cRelatedTo, err)
		}
	}
	if s.DefaultAlarm != nil {
		if err := s.DefaultAlarm.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmDisplay, cDefaultAlarm, err)
		}
	}
	return nil
}

// AlarmEmail provides a group of components that define an Email Alarm
type AlarmEmail struct {
	Description   PropDescription
	Trigger       PropTrigger
	Summary       PropSummary
	Attendee      *PropAttendee
	Duration      *PropDuration
	Repeat        *PropRepeat
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}

func (s *AlarmEmail) decode(t tokeniser) error {
	var requiredDescription, requiredTrigger, requiredSummary bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return fmt.Errorf(errDecodingType, cAlarmEmail, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cAlarmEmail, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cAlarmEmail, err)
				}
			}
		case "DESCRIPTION":
			if requiredDescription {
				return fmt.Errorf(errMultiple, cAlarmEmail, cDescription)
			}
			requiredDescription = true
			if err := s.Description.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cDescription, err)
			}
		case "TRIGGER":
			if requiredTrigger {
				return fmt.Errorf(errMultiple, cAlarmEmail, cTrigger)
			}
			requiredTrigger = true
			if err := s.Trigger.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cTrigger, err)
			}
		case "SUMMARY":
			if requiredSummary {
				return fmt.Errorf(errMultiple, cAlarmEmail, cSummary)
			}
			requiredSummary = true
			if err := s.Summary.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cSummary, err)
			}
		case "ATTENDEE":
			if s.Attendee != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cAttendee)
			}
			s.Attendee = new(PropAttendee)
			if err := s.Attendee.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cAttendee, err)
			}
		case "DURATION":
			if s.Duration != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cDuration)
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cDuration, err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cRepeat)
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cRepeat, err)
			}
		case "UID":
			if s.UID != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cUID)
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cUID, err)
			}
		case "ALARM-AGENT":
			var e PropAlarmAgent
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cAlarmAgent, err)
			}
			s.AlarmAgent = append(s.AlarmAgent, e)
		case "STATUS":
			if s.AlarmStatus != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cAlarmStatus)
			}
			s.AlarmStatus = new(PropAlarmStatus)
			if err := s.AlarmStatus.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cAlarmStatus, err)
			}
		case "LAST-TRIGGERED":
			var e PropLastTriggered
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cLastTriggered, err)
			}
			s.LastTriggered = append(s.LastTriggered, e)
		case "ACKNOWLEDGED":
			if s.Acknowledged != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cAcknowledged)
			}
			s.Acknowledged = new(PropAcknowledged)
			if err := s.Acknowledged.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cAcknowledged, err)
			}
		case "PROXIMITY":
			if s.Proximity != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cProximity)
			}
			s.Proximity = new(PropProximity)
			if err := s.Proximity.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cProximity, err)
			}
		case "GEO-LOCATION":
			var e PropGeoLocation
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cGeoLocation, err)
			}
			s.GeoLocation = append(s.GeoLocation, e)
		case "RELATED-TO":
			if s.RelatedTo != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cRelatedTo)
			}
			s.RelatedTo = new(PropRelatedTo)
			if err := s.RelatedTo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cRelatedTo, err)
			}
		case "DEFAULT-ALARM":
			if s.DefaultAlarm != nil {
				return fmt.Errorf(errMultiple, cAlarmEmail, cDefaultAlarm)
			}
			s.DefaultAlarm = new(PropDefaultAlarm)
			if err := s.DefaultAlarm.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmEmail, cDefaultAlarm, err)
			}
		case "END":
			if value != "VALARM" {
				return fmt.Errorf(errDecodingType, cAlarmEmail, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredDescription || !requiredTrigger || !requiredSummary {
		return fmt.Errorf(errDecodingType, cAlarmEmail, ErrMissingRequired)
	}
	if s.GeoLocation != nil && (s.Proximity == nil) {
		return fmt.Errorf(errDecodingType, cAlarmEmail, ErrRequirementNotMet)
	}
	if t := s.Duration == nil; t == (s.Repeat == nil) {
		return fmt.Errorf(errDecodingType, cAlarmEmail, ErrRequirementNotMet)
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
	for n := range s.AlarmAgent {
		s.AlarmAgent[n].encode(w)
	}
	if s.AlarmStatus != nil {
		s.AlarmStatus.encode(w)
	}
	for n := range s.LastTriggered {
		s.LastTriggered[n].encode(w)
	}
	if s.Acknowledged != nil {
		s.Acknowledged.encode(w)
	}
	if s.Proximity != nil {
		s.Proximity.encode(w)
	}
	for n := range s.GeoLocation {
		s.GeoLocation[n].encode(w)
	}
	if s.RelatedTo != nil {
		s.RelatedTo.encode(w)
	}
	if s.DefaultAlarm != nil {
		s.DefaultAlarm.encode(w)
	}
}

func (s *AlarmEmail) valid() error {
	if err := s.Description.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cAlarmEmail, cDescription, err)
	}
	if err := s.Trigger.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cAlarmEmail, cTrigger, err)
	}
	if err := s.Summary.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cAlarmEmail, cSummary, err)
	}
	if s.Attendee != nil {
		if err := s.Attendee.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cAttendee, err)
		}
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cDuration, err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cRepeat, err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cUID, err)
		}
	}
	for n := range s.AlarmAgent {
		if err := s.AlarmAgent[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cAlarmAgent, err)
		}
	}
	if s.AlarmStatus != nil {
		if err := s.AlarmStatus.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cAlarmStatus, err)
		}
	}
	for n := range s.LastTriggered {
		if err := s.LastTriggered[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cLastTriggered, err)
		}
	}
	if s.Acknowledged != nil {
		if err := s.Acknowledged.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cAcknowledged, err)
		}
	}
	if s.Proximity != nil {
		if err := s.Proximity.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cProximity, err)
		}
	}
	for n := range s.GeoLocation {
		if err := s.GeoLocation[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cGeoLocation, err)
		}
	}
	if s.RelatedTo != nil {
		if err := s.RelatedTo.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cRelatedTo, err)
		}
	}
	if s.DefaultAlarm != nil {
		if err := s.DefaultAlarm.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmEmail, cDefaultAlarm, err)
		}
	}
	return nil
}

// AlarmURI provies a group of components that define a URI Alarm
type AlarmURI struct {
	URI           PropURI
	Duration      *PropDuration
	Repeat        *PropRepeat
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}

func (s *AlarmURI) decode(t tokeniser) error {
	var requiredURI bool
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return fmt.Errorf(errDecodingType, cAlarmURI, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cAlarmURI, io.ErrUnexpectedEOF)
		}
		params := p.Data[1 : len(p.Data)-1]
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cAlarmURI, err)
				}
			}
		case "URI":
			if requiredURI {
				return fmt.Errorf(errMultiple, cAlarmURI, cURI)
			}
			requiredURI = true
			if err := s.URI.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cURI, err)
			}
		case "DURATION":
			if s.Duration != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cDuration)
			}
			s.Duration = new(PropDuration)
			if err := s.Duration.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cDuration, err)
			}
		case "REPEAT":
			if s.Repeat != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cRepeat)
			}
			s.Repeat = new(PropRepeat)
			if err := s.Repeat.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cRepeat, err)
			}
		case "UID":
			if s.UID != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cUID)
			}
			s.UID = new(PropUID)
			if err := s.UID.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cUID, err)
			}
		case "ALARM-AGENT":
			var e PropAlarmAgent
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cAlarmAgent, err)
			}
			s.AlarmAgent = append(s.AlarmAgent, e)
		case "STATUS":
			if s.AlarmStatus != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cAlarmStatus)
			}
			s.AlarmStatus = new(PropAlarmStatus)
			if err := s.AlarmStatus.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cAlarmStatus, err)
			}
		case "LAST-TRIGGERED":
			var e PropLastTriggered
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cLastTriggered, err)
			}
			s.LastTriggered = append(s.LastTriggered, e)
		case "ACKNOWLEDGED":
			if s.Acknowledged != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cAcknowledged)
			}
			s.Acknowledged = new(PropAcknowledged)
			if err := s.Acknowledged.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cAcknowledged, err)
			}
		case "PROXIMITY":
			if s.Proximity != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cProximity)
			}
			s.Proximity = new(PropProximity)
			if err := s.Proximity.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cProximity, err)
			}
		case "GEO-LOCATION":
			var e PropGeoLocation
			if err := e.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cGeoLocation, err)
			}
			s.GeoLocation = append(s.GeoLocation, e)
		case "RELATED-TO":
			if s.RelatedTo != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cRelatedTo)
			}
			s.RelatedTo = new(PropRelatedTo)
			if err := s.RelatedTo.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cRelatedTo, err)
			}
		case "DEFAULT-ALARM":
			if s.DefaultAlarm != nil {
				return fmt.Errorf(errMultiple, cAlarmURI, cDefaultAlarm)
			}
			s.DefaultAlarm = new(PropDefaultAlarm)
			if err := s.DefaultAlarm.decode(params, value); err != nil {
				return fmt.Errorf(errDecodingProp, cAlarmURI, cDefaultAlarm, err)
			}
		case "END":
			if value != "VALARM" {
				return fmt.Errorf(errDecodingType, cAlarmURI, ErrInvalidEnd)
			}
			break Loop
		}
	}
	if !requiredURI {
		return fmt.Errorf(errDecodingType, cAlarmURI, ErrMissingRequired)
	}
	if s.GeoLocation != nil && (s.Proximity == nil) {
		return fmt.Errorf(errDecodingType, cAlarmURI, ErrRequirementNotMet)
	}
	if t := s.Duration == nil; t == (s.Repeat == nil) {
		return fmt.Errorf(errDecodingType, cAlarmURI, ErrRequirementNotMet)
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
	for n := range s.AlarmAgent {
		s.AlarmAgent[n].encode(w)
	}
	if s.AlarmStatus != nil {
		s.AlarmStatus.encode(w)
	}
	for n := range s.LastTriggered {
		s.LastTriggered[n].encode(w)
	}
	if s.Acknowledged != nil {
		s.Acknowledged.encode(w)
	}
	if s.Proximity != nil {
		s.Proximity.encode(w)
	}
	for n := range s.GeoLocation {
		s.GeoLocation[n].encode(w)
	}
	if s.RelatedTo != nil {
		s.RelatedTo.encode(w)
	}
	if s.DefaultAlarm != nil {
		s.DefaultAlarm.encode(w)
	}
}

func (s *AlarmURI) valid() error {
	if err := s.URI.valid(); err != nil {
		return fmt.Errorf(errValidatingProp, cAlarmURI, cURI, err)
	}
	if s.Duration != nil {
		if err := s.Duration.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cDuration, err)
		}
	}
	if s.Repeat != nil {
		if err := s.Repeat.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cRepeat, err)
		}
	}
	if s.UID != nil {
		if err := s.UID.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cUID, err)
		}
	}
	for n := range s.AlarmAgent {
		if err := s.AlarmAgent[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cAlarmAgent, err)
		}
	}
	if s.AlarmStatus != nil {
		if err := s.AlarmStatus.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cAlarmStatus, err)
		}
	}
	for n := range s.LastTriggered {
		if err := s.LastTriggered[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cLastTriggered, err)
		}
	}
	if s.Acknowledged != nil {
		if err := s.Acknowledged.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cAcknowledged, err)
		}
	}
	if s.Proximity != nil {
		if err := s.Proximity.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cProximity, err)
		}
	}
	for n := range s.GeoLocation {
		if err := s.GeoLocation[n].valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cGeoLocation, err)
		}
	}
	if s.RelatedTo != nil {
		if err := s.RelatedTo.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cRelatedTo, err)
		}
	}
	if s.DefaultAlarm != nil {
		if err := s.DefaultAlarm.valid(); err != nil {
			return fmt.Errorf(errValidatingProp, cAlarmURI, cDefaultAlarm, err)
		}
	}
	return nil
}

// AlarmNone
type AlarmNone struct{}

func (s *AlarmNone) decode(t tokeniser) error {
Loop:
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return fmt.Errorf(errDecodingType, cAlarmNone, err)
		} else if p.Type == parser.PhraseDone {
			return fmt.Errorf(errDecodingType, cAlarmNone, io.ErrUnexpectedEOF)
		}
		value := p.Data[len(p.Data)-1].Data
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			switch n := strings.ToUpper(value); n {
			default:
				if err := decodeDummy(t, n); err != nil {
					return fmt.Errorf(errDecodingType, cAlarmNone, err)
				}
			}
		case "END":
			if value != "VALARM" {
				return fmt.Errorf(errDecodingType, cAlarmNone, ErrInvalidEnd)
			}
			break Loop
		}
	}
	return nil
}

func (s *AlarmNone) encode(w writer) {
}

func (s *AlarmNone) valid() error {
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
	errMultiple                       = "error decoding %s: multiple %s"
	errMissing                        = "error validating %s: missing %s"
	cCalendar                         = "Calendar"
	cEvent                            = "Event"
	cTodo                             = "Todo"
	cJournal                          = "Journal"
	cTimezone                         = "Timezone"
	cStandard                         = "Standard"
	cDaylight                         = "Daylight"
	cAlarmAudio                       = "AlarmAudio"
	cAlarmDisplay                     = "AlarmDisplay"
	cAlarmEmail                       = "AlarmEmail"
	cAlarmURI                         = "AlarmURI"
	cAlarmNone                        = "AlarmNone"
)
