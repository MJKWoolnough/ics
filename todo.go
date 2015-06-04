package ics

import (
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
	Duration           time.Duration
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

func (c *Calendar) decodeTodo(d Decoder) error {
	bm := bitmask.New(20)
	var t Todo
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
			t.Duration = p.Duration
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
			for i := 0; i < 17; i++ {
				if !bm.Get(i) {
					return ErrRequiredMissing
				}
			}
			c.Todo = append(c.Todo, t)
			return nil
		}
	}
	return nil
}

func (c *Calendar) encodeTodos(e Encoder) error {
	return nil
}
