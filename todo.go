package ics

import (
	"math"
	"time"

	"github.com/MJKWoolnough/bitmask"
)

const vTodo = "VTODO"

type Todo struct {
	LastModified       time.Time
	UID                string
	Class              class
	Completed, Created time.Time
	Description        description
	Start, End         dateTime
	Duration           duration
	Geo                geo
	Location           location
	Organizer          organizer
	PercentComplete    percentComplete
	Priority           priority
	RecurrenceID       recurrenceID
	Sequence           sequence
	Status             status
	Summary            summary
	URL                url
	RecurrenceRule     recurrenceRule
	Attachments        []attach
	Attendees          []attendee
	Categories         map[string][]string
	Comments           []comment
	Contacts           []contact
	ExceptionDates     []exceptionDate
	RequestStatuses    []requestStatus
	RelatedTo          []relatedTo
	Resources          []resources
	RecurrenceDates    []recurrenceDate
}

func NewTodo() Todo {
	var t Todo
	t.Geo.Latitude = math.NaN()
	t.Geo.Longitude = math.NaN()
	t.Class = -1
	t.Sequence = -1
	t.Status = -1
	t.RecurrenceRule.Frequency = -1
	return t
}

func (c *Calendar) decodeTodo(d Decoder) error {
	bm := bitmask.New(20)
	t := NewTodo()
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateStamp:
			if bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			t.LastModified = p.Time
		case uid:
			if bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			t.UID = string(p)
		case class:
			if bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			t.Class = p
		case completed:
			if bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			t.Completed = p.Time
		case created:
			if bm.SetIfNot(4, true) {
				return ErrMultipleUnique
			}
			t.Created = p.Time
		case description:
			if bm.SetIfNot(5, true) {
				return ErrMultipleUnique
			}
			t.Description = p
		case dateTimeStart:
			if bm.SetIfNot(6, true) {
				return ErrMultipleUnique
			}
			t.Start = p.dateTime
		case geo:
			if bm.SetIfNot(7, true) {
				return ErrMultipleUnique
			}
			t.Geo = p
		case lastModified:
			if bm.SetIfNot(8, true) {
				return ErrMultipleUnique
			}
			t.LastModified = p.Time
		case location:
			if bm.SetIfNot(9, true) {
				return ErrMultipleUnique
			}
			t.Location = p
		case organizer:
			if bm.SetIfNot(10, true) {
				return ErrMultipleUnique
			}
			t.Organizer = p
		case percentComplete:
			if bm.SetIfNot(11, true) {
				return ErrMultipleUnique
			}
			t.PercentComplete = p
		case priority:
			if bm.SetIfNot(12, true) {
				return ErrMultipleUnique
			}
			t.Priority = p
		case recurrenceID:
			if bm.SetIfNot(13, true) {
				return ErrMultipleUnique
			}
			t.RecurrenceID = p
		case sequence:
			if bm.SetIfNot(14, true) {
				return ErrMultipleUnique
			}
			t.Sequence = p
		case status:
			if bm.SetIfNot(15, true) {
				return ErrMultipleUnique
			}
			t.Status = p
		case summary:
			if bm.SetIfNot(16, true) {
				return ErrMultipleUnique
			}
			t.Summary = p
		case url:
			if bm.SetIfNot(17, true) {
				return ErrMultipleUnique
			}
			t.URL = p
		case recurrenceRule:
			t.RecurrenceRule = p
		case dateTimeDue:
			if bm.Get(19) {
				return ErrInvalidComponentCombination
			}
			if bm.SetIfNot(18, true) {
				return ErrMultipleUnique
			}
			t.End = p.dateTime
		case duration:
			if bm.Get(18) {
				return ErrInvalidComponentCombination
			}
			if bm.SetIfNot(19, true) {
				return ErrMultipleUnique
			}
			t.Duration = p
		case attach:
			t.Attachments = append(t.Attachments, p)
		case attendee:
			t.Attendees = append(t.Attendees, p)
		case categories:
			var cats []string
			if cts, ok := t.Categories[p.Language]; ok {
				cats = cts
			}
			cats = append(cats, p.Categories...)
			t.Categories[p.Language] = cats
		case comment:
			t.Comments = append(t.Comments, p)
		case contact:
			t.Contacts = append(t.Contacts, p)
		case exceptionDate:
			t.ExceptionDates = append(t.ExceptionDates, p)
		case requestStatus:
			t.RequestStatuses = append(t.RequestStatuses, p)
		case relatedTo:
			t.RelatedTo = append(t.RelatedTo, p)
		case resources:
			t.Resources = append(t.Resources, p)
		case recurrenceDate:
			t.RecurrenceDates = append(t.RecurrenceDates, p)
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vTodo {
				return ErrInvalidEnd
			}
			if !bm.Get(0) || !bm.Get(1) {
				return ErrRequiredMissing
			}
			if bm.Get(6) != bm.Get(19) {
				return ErrInvalidComponentCombination
			}
			c.Todo = append(c.Todo, t)
			return nil
		}
	}
	return nil
}

func (c *Calendar) writeTodoData(e *Encoder) {
	for _, t := range c.Todo {
		e.writeProperty(begin(vTodo))
		e.writeProperty(dateStamp{dateTime{Time: t.LastModified}})
		e.writeProperty(uid(t.UID))
		if t.Class >= 0 {
			e.writeProperty(t.Class)
		}
		if !t.Completed.IsZero() {
			e.writeProperty(completed{dateTime{Time: t.Completed}})
		}
		if !t.Created.IsZero() {
			e.writeProperty(created{dateTime{Time: t.Completed}})
		}
		if t.Description.String != "" {
			e.writeProperty(t.Description)
		}
		if !t.Start.IsZero() {
			e.writeProperty(dateTimeStart{t.Start})
		}
		if t.Geo.Latitude == t.Geo.Latitude && t.Geo.Longitude == t.Geo.Longitude {
			e.writeProperty(t.Geo)
		}
		if !t.LastModified.IsZero() {
			e.writeProperty(lastModified{dateTime{Time: t.LastModified}})
		}
		if t.Location.String != "" {
			e.writeProperty(t.Location)
		}
		if t.Organizer.Name != "" {
			e.writeProperty(t.Organizer)
		}
		if t.PercentComplete > 0 {
			e.writeProperty(t.PercentComplete)
		}
		if t.Priority > 0 {
			e.writeProperty(t.Priority)
		}
		if !t.RecurrenceID.DateTime.IsZero() {
			e.writeProperty(t.RecurrenceID)
		}
		if t.Sequence >= 0 {
			e.writeProperty(t.Sequence)
		}
		if t.Status >= 0 {
			e.writeProperty(t.Status)
		}
		if t.Summary.String != "" {
			e.writeProperty(t.Summary)
		}
		if t.URL != "" {
			e.writeProperty(t.URL)
		}
		if t.RecurrenceRule.Frequency >= 0 {
			e.writeProperty(t.RecurrenceRule)
		}
		if !t.End.IsZero() {
			e.writeProperty(dateTimeDue{t.End})
		} else if !t.Start.IsZero() && t.Duration.Duration != 0 {
			e.writeProperty(t.Duration)
		}
		for _, p := range t.Attachments {
			e.writeProperty(p)
		}
		for _, p := range t.Attendees {
			e.writeProperty(p)
		}
		for l, cs := range t.Categories {
			e.writeProperty(categories{
				Language:   l,
				Categories: cs,
			})
		}
		for _, p := range t.Comments {
			e.writeProperty(p)
		}
		for _, p := range t.Contacts {
			e.writeProperty(p)
		}
		for _, p := range t.ExceptionDates {
			e.writeProperty(p)
		}
		for _, p := range t.RequestStatuses {
			e.writeProperty(p)
		}
		for _, p := range t.RelatedTo {
			e.writeProperty(p)
		}
		for _, p := range t.Resources {
			e.writeProperty(p)
		}
		for _, p := range t.RecurrenceDates {
			e.writeProperty(p)
		}
		e.writeProperty(end(vTodo))
	}
}
