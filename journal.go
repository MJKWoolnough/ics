package ics

import (
	"time"

	"github.com/MJKWoolnough/bitmask"
)

const vJournal = "VJournal"

type Journal struct {
	LastModified, Created time.Time
	UID                   string
	Class                 class
	Start                 dateTime
	Organizer             organizer
	RecurrenceID          recurrenceID
	Sequence              sequence
	Status                status
	Summary               summary
	URL                   url
	RecurrenceRule        recurrenceRule
	Attachments           []attach
	Attendees             []attendee
	Categories            map[string][]string
	Comments              []comment
	Contacts              []contact
	Descriptions          []description
	ExceptionDates        []exceptionDate
	RelatedTo             []relatedTo
	RecurrenceDates       []recurrenceDate
	RequestStatus         []requestStatus
}

func (c *Calendar) decodeJournal(d Decoder) error {
	bm := bitmask.New(12)
	var j Journal
	j.Class = -1
	j.Sequence = -1
	j.Status = -1
	j.RecurrenceRule.Frequency = -1
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
			j.LastModified = p.Time
		case uid:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			j.UID = string(p)
		case class:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			j.Class = p
		case created:
			if !bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			j.Created = p.Time
		case dateTimeStart:
			if !bm.SetIfNot(4, true) {
				return ErrMultipleUnique
			}
			j.Start = p.dateTime
		case lastModified:
			if !bm.SetIfNot(5, true) {
				return ErrMultipleUnique
			}
			j.LastModified = p.Time
		case organizer:
			if !bm.SetIfNot(6, true) {
				return ErrMultipleUnique
			}
			j.Organizer = p
		case recurrenceID:
			if !bm.SetIfNot(7, true) {
				return ErrMultipleUnique
			}
			j.RecurrenceID = p
		case sequence:
			if !bm.SetIfNot(8, true) {
				return ErrMultipleUnique
			}
			j.Sequence = p
		case status:
			if !bm.SetIfNot(9, true) {
				return ErrMultipleUnique
			}
			j.Status = p
		case summary:
			if !bm.SetIfNot(10, true) {
				return ErrMultipleUnique
			}
			j.Summary = p
		case url:
			if !bm.SetIfNot(11, true) {
				return ErrMultipleUnique
			}
			j.URL = p
		case recurrenceRule:
			j.RecurrenceRule = p
		case attach:
			j.Attachments = append(j.Attachments, p)
		case attendee:
			j.Attendees = append(j.Attendees, p)
		case categories:
			var cats []string
			if cts, ok := j.Categories[p.Language]; ok {
				cats = cts
			}
			cats = append(cats, p.Categories...)
			j.Categories[p.Language] = cats
		case comment:
			j.Comments = append(j.Comments, p)
		case contact:
			j.Contacts = append(j.Contacts, p)
		case description:
			j.Descriptions = append(j.Descriptions, p)
		case exceptionDate:
			j.ExceptionDates = append(j.ExceptionDates, p)
		case relatedTo:
			j.RelatedTo = append(j.RelatedTo, p)
		case recurrenceDate:
			j.RecurrenceDates = append(j.RecurrenceDates, p)
		case requestStatus:
			j.RequestStatus = append(j.RequestStatus, p)
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vJournal {
				return ErrInvalidEnd
			}
			if !bm.Get(0) || !bm.Get(1) {
				return ErrRequiredMissing
			}
			c.Journals = append(c.Journals, j)
			return nil
		}
	}
	return nil
}

func (c *Calendar) writeJournalData(e *Encoder) {
	for _, j := range c.Journals {
		e.writeProperty(begin(vJournal))
		e.writeProperty(dateStamp{dateTime{Time: j.LastModified}})
		e.writeProperty(uid(j.UID))
		if j.Class >= 0 {
			e.writeProperty(j.Class)
		}
		if !j.Created.IsZero() {
			e.writeProperty(created{dateTime{Time: j.Created}})
		}
		if !j.Start.IsZero() {
			e.writeProperty(dateTimeStart{j.Start})
		}
		if j.Organizer.Name != "" {
			e.writeProperty(j.Organizer)
		}
		if !j.RecurrenceID.DateTime.IsZero() {
			e.writeProperty(j.RecurrenceID)
		}
		if j.Sequence >= 0 {
			e.writeProperty(j.Sequence)
		}
		if j.Status >= 0 {
			e.writeProperty(j.Status)
		}
		if j.Summary.String != "" {
			e.writeProperty(j.Summary)
		}
		if j.URL != "" {
			e.writeProperty(j.URL)
		}
		if j.RecurrenceRule.Frequency != -1 {
			e.writeProperty(j.RecurrenceRule)
		}
		for _, p := range j.Attachments {
			e.writeProperty(p)
		}
		for _, p := range j.Attendees {
			e.writeProperty(p)
		}
		for l, cs := range j.Categories {
			e.writeProperty(categories{
				Language:   l,
				Categories: cs,
			})
		}
		for _, p := range j.Comments {
			e.writeProperty(p)
		}
		for _, p := range j.Contacts {
			e.writeProperty(p)
		}
		for _, p := range j.Descriptions {
			e.writeProperty(p)
		}
		for _, p := range j.ExceptionDates {
			e.writeProperty(p)
		}
		for _, p := range j.RelatedTo {
			e.writeProperty(p)
		}
		for _, p := range j.RecurrenceDates {
			e.writeProperty(p)
		}
		for _, p := range j.RequestStatus {
			e.writeProperty(p)
		}
		e.writeProperty(end(vJournal))
	}
}
