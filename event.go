package ics

import (
	"time"

	"github.com/MJKWoolnough/bitmask"
)

const vEvent = "VEVENT"

type Event struct {
	LastModified, Created time.Time
	UID                   string
	Start                 dateTime
	class                 class
	Description           description
	Geo                   geo
	Location              location
	Organizer             organizer
	Priority              priority
	Sequence              sequence
	Status                status
	Summary               summary
	TimeTransparency      timeTransparency
	URL                   url
	RecurrenceID          recurrenceID
	RecurrenceRule        recurrenceRule
	End                   dateTime
	Duration              time.Duration
	Attachments           []attach
	Attendees             []attendee
	Categories            map[string][]string
	Comments              []comment
	Contacts              []contact
	ExceptionDates        []exceptionDate
	RequestStatus         []requestStatus
	RelatedTo             []relatedTo
	Resources             []resources
	RecurrenceDate        []recurrenceDate
	Alarms                []Alarm
}

func (c *Calendar) decodeEvent(d Decoder) error {
	bm := bitmask.New(19)
	var e Event
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateStamp:
			if !bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			if c.Method == "" {
				e.LastModified = p.Time
			}
		case uid:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			e.UID = string(p)
		case dateTimeStart:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			e.Start = p.dateTime
		case class:
			if !bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			e.class = p
		case created:
			if !bm.SetIfNot(4, true) {
				return ErrMultipleUnique
			}
			e.Created = p.Time
		case description:
			if !bm.SetIfNot(5, true) {
				return ErrMultipleUnique
			}
			e.Description = p
		case geo:
			if !bm.SetIfNot(6, true) {
				return ErrMultipleUnique
			}
			e.Geo = p
		case lastModified:
			if !bm.SetIfNot(7, true) {
				return ErrMultipleUnique
			}
			e.LastModified = p.Time
		case location:
			if !bm.SetIfNot(8, true) {
				return ErrMultipleUnique
			}
			e.Location = p
		case organizer:
			if !bm.SetIfNot(9, true) {
				return ErrMultipleUnique
			}
			e.Organizer = p
		case priority:
			if !bm.SetIfNot(10, true) {
				return ErrMultipleUnique
			}
			e.Priority = p
		case sequence:
			if !bm.SetIfNot(11, true) {
				return ErrMultipleUnique
			}
			e.Sequence = p
		case status:
			if !bm.SetIfNot(12, true) {
				return ErrMultipleUnique
			}
			e.Status = p
		case summary:
			if !bm.SetIfNot(13, true) {
				return ErrMultipleUnique
			}
			e.Summary = p
		case timeTransparency:
			if !bm.SetIfNot(14, true) {
				return ErrMultipleUnique
			}
			e.TimeTransparency = p
		case url:
			if !bm.SetIfNot(15, true) {
				return ErrMultipleUnique
			}
			e.URL = p
		case recurrenceID:
			if !bm.SetIfNot(16, true) {
				return ErrMultipleUnique
			}
			e.RecurrenceID = p
		case recurrenceRule:
			e.RecurrenceRule = p
		case dateTimeEnd:
			if bm.Get(18) {
				return ErrInvalidComponentCombination
			}
			if !bm.SetIfNot(17, true) {
				return ErrMultipleUnique
			}
			e.End = p.dateTime
		case duration:
			if bm.Get(17) {
				return ErrInvalidComponentCombination
			}
			if !bm.SetIfNot(18, true) {
				return ErrMultipleUnique
			}
			e.Duration = p.Duration
		case attach:
			e.Attachments = append(e.Attachments, p)
		case attendee:
			e.Attendees = append(e.Attendees, p)
		case categories:
			var cats []string
			if cts, ok := e.Categories[p.Language]; ok {
				cats = cts
			}
			cats = append(cats, p.Categories...)
			e.Categories[p.Language] = cats
		case comment:
			e.Comments = append(e.Comments, p)
		case contact:
			e.Contacts = append(e.Contacts, p)
		case exceptionDate:
			e.ExceptionDates = append(e.ExceptionDates, p)
		case requestStatus:
			e.RequestStatus = append(e.RequestStatus, p)
		case relatedTo:
			e.RelatedTo = append(e.RelatedTo, p)
		case resources:
			e.Resources = append(e.Resources, p)
		case recurrenceDate:
			e.RecurrenceDate = append(e.RecurrenceDate, p)
		case begin:
			switch p {
			case vAlarm:
				a, err := c.decodeAlarm(d)
				if err != nil {
					return nil
				}
				e.Alarms = append(e.Alarms, a)
			default:
				if err = d.readUnknownComponent(string(p)); err != nil {
					return err
				}
			}
		case end:
			if p != vEvent {
				return ErrInvalidEnd
			}
			c.Events = append(c.Events, e)
			return nil
		}
	}
	return nil
}
