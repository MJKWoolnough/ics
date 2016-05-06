package ics

import (
	"math"
	"time"

	"github.com/MJKWoolnough/bitmask"
)

const vEvent = "VEVENT"

type Event struct {
	LastModified, Created time.Time
	UID                   string
	Start                 dateTimeStart
	Class                 class
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
	End                   dateTimeEnd
	Duration              duration
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

func NewEvent() Event {
	var e Event
	e.Geo.Latitude = math.NaN()
	e.Geo.Longitude = math.NaN()
	e.Priority = -1
	e.Class = -1
	e.TimeTransparency = -1
	e.Sequence = -1
	e.Status = -1
	e.RecurrenceRule.Frequency = -1
	return e
}

func (c *Calendar) decodeEvent(d Decoder) error {
	bm := bitmask.New(19)
	e := NewEvent()
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
			e.Start = p
		case class:
			if !bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			e.Class = p
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
			e.End = p
		case duration:
			if bm.Get(17) {
				return ErrInvalidComponentCombination
			}
			if !bm.SetIfNot(18, true) {
				return ErrMultipleUnique
			}
			e.Duration = p
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

func (c *Calendar) writeEventData(e *Encoder) {
	for _, ev := range c.Events {
		e.writeProperty(begin(vEvent))
		e.writeProperty(dateStamp{dateTime{Time: ev.LastModified}})
		e.writeProperty(uid(ev.UID))
		if c.Method == "" || !ev.Start.dateTime.IsZero() {
			e.writeProperty(ev.Start)
		}
		if ev.Class >= 0 {
			e.writeProperty(ev.Class)
		}
		if !ev.Created.IsZero() {
			e.writeProperty(created{dateTime{Time: ev.Created}})
		}
		if ev.Description.String != "" {
			e.writeProperty(ev.Description)
		}
		if !math.IsNaN(ev.Geo.Latitude) && !math.IsNaN(ev.Geo.Longitude) {
			e.writeProperty(ev.Geo)
		}
		if ev.Location.String != "" {
			e.writeProperty(ev.Location)
		}
		if ev.Organizer.Name != "" {
			e.writeProperty(ev.Organizer)
		}
		if ev.Priority >= 0 {
			e.writeProperty(ev.Priority)
		}
		if ev.Sequence >= 0 {
			e.writeProperty(ev.Sequence)
		}
		if ev.Status >= 0 {
			e.writeProperty(ev.Status)
		}
		if ev.Summary.String != "" {
			e.writeProperty(ev.Summary)
		}
		if ev.TimeTransparency >= 0 {
			e.writeProperty(ev.TimeTransparency)
		}
		if ev.URL != "" {
			e.writeProperty(ev.URL)
		}
		if !ev.RecurrenceID.DateTime.IsZero() {
			e.writeProperty(ev.RecurrenceID)
		}
		if ev.RecurrenceRule.Frequency >= 0 {
			e.writeProperty(ev.RecurrenceRule)
		}
		if !ev.End.IsZero() {
			e.writeProperty(ev.End)
		} else if ev.Duration.Duration > 0 {
			e.writeProperty(ev.Duration)
		}
		for _, p := range ev.Attachments {
			e.writeProperty(p)
		}
		for _, p := range ev.Attendees {
			e.writeProperty(p)
		}
		for l, cs := range ev.Categories {
			e.writeProperty(categories{
				Language:   l,
				Categories: cs,
			})
		}
		for _, p := range ev.Comments {
			e.writeProperty(p)
		}
		for _, p := range ev.Contacts {
			e.writeProperty(p)
		}
		for _, p := range ev.ExceptionDates {
			e.writeProperty(p)
		}
		for _, p := range ev.RequestStatus {
			e.writeProperty(p)
		}
		for _, p := range ev.RelatedTo {
			e.writeProperty(p)
		}
		for _, p := range ev.Resources {
			e.writeProperty(p)
		}
		for _, p := range ev.RecurrenceDate {
			e.writeProperty(p)
		}
		for _, a := range ev.Alarms {
			e.writeProperty(begin(vAlarm))
			a.writeAlarmData(e)
			e.writeProperty(end(vAlarm))
		}
		e.writeProperty(end(vEvent))
	}
}
