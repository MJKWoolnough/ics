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

func (c *Calendar) decodeEvent(d Decoder) error {
	bm := bitmask.New(19)
	var e Event
	e.Geo.Latitude = math.NaN()
	e.Geo.Longitude = math.NaN()
	e.Priority = -1
	e.Class = -1
	e.TimeTransparency = -1
	e.Sequence = -1
	e.Status = -1
	e.RecurrenceRule.Frequency = -1
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

func (c *Calendar) eventData() []property {
	data := make([]property, 0, 1024)
	for _, e := range c.Events {
		data = append(data, begin(vEvent))
		data = append(data, dateStamp{dateTime{Time: e.LastModified}}, uid(e.UID))
		if c.Method == "" || !e.Start.dateTime.IsZero() {
			data = append(data, e.Start)
		}
		if e.Class >= 0 {
			data = append(data, e.Class)
		}
		if !e.Created.IsZero() {
			data = append(data, created{dateTime{Time: e.Created}})
		}
		if e.Description.String != "" {
			data = append(data, e.Description)
		}
		if e.Geo.Latitude == e.Geo.Latitude && e.Geo.Longitude == e.Geo.Longitude {
			data = append(data, e.Geo)
		}
		if e.Location.String != "" {
			data = append(data, e.Location)
		}
		if e.Organizer.Name != "" {
			data = append(data, e.Organizer)
		}
		if e.Priority >= 0 {
			data = append(data, e.Priority)
		}
		if e.Sequence >= 0 {
			data = append(data, e.Sequence)
		}
		if e.Status >= 0 {
			data = append(data, e.Status)
		}
		if e.Summary.String != "" {
			data = append(data, e.Summary)
		}
		if e.TimeTransparency >= 0 {
			data = append(data, e.TimeTransparency)
		}
		if e.URL != "" {
			data = append(data, e.URL)
		}
		if !e.RecurrenceID.DateTime.IsZero() {
			data = append(data, e.RecurrenceID)
		}
		if e.RecurrenceRule.Frequency >= 0 {
			data = append(data, e.RecurrenceRule)
		}
		if !e.End.IsZero() {
			data = append(data, e.End)
		} else if e.Duration.Duration > 0 {
			data = append(data, e.Duration)
		}
		for _, p := range e.Attachments {
			data = append(data, p)
		}
		for _, p := range e.Attendees {
			data = append(data, p)
		}
		for l, cs := range e.Categories {
			data = append(data, categories{
				Language:   l,
				Categories: cs,
			})
		}
		for _, p := range e.Comments {
			data = append(data, p)
		}
		for _, p := range e.Contacts {
			data = append(data, p)
		}
		for _, p := range e.ExceptionDates {
			data = append(data, p)
		}
		for _, p := range e.RequestStatus {
			data = append(data, p)
		}
		for _, p := range e.RelatedTo {
			data = append(data, p)
		}
		for _, p := range e.Resources {
			data = append(data, p)
		}
		for _, p := range e.RecurrenceDate {
			data = append(data, p)
		}
		for _, a := range e.Alarms {
			data = append(data, begin(vAlarm))
			data = append(data, a.alarmData()...)
			data = append(data, end(vAlarm))
		}
		data = append(data, end(vEvent))
	}
	return data
}
