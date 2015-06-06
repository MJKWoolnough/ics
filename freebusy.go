package ics

import (
	"time"

	"github.com/MJKWoolnough/bitmask"
)

const vFreeBusy = "VFREEBUSY"

type FreeBusy struct {
	LastModified  time.Time
	UID           string
	Contact       contact
	Start, End    dateTime
	Organizer     organizer
	URL           url
	Attendees     []attendee
	Comments      []comment
	FreeBusy      []freeBusyTime
	RequestStatus []requestStatus
}

func (c *Calendar) decodeFreeBusy(d Decoder) error {
	bm := bitmask.New(7)
	var f FreeBusy
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
			f.LastModified = p.Time
		case uid:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			f.UID = string(p)
		case contact:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			f.Contact = p
		case dateTimeStart:
			if !bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			f.Start = p.dateTime
		case dateTimeEnd:
			if !bm.SetIfNot(4, true) {
				return ErrMultipleUnique
			}
			f.End = p.dateTime
		case organizer:
			if !bm.SetIfNot(5, true) {
				return ErrMultipleUnique
			}
			f.Organizer = p
		case url:
			if !bm.SetIfNot(6, true) {
				return ErrMultipleUnique
			}
			f.URL = p
		case attendee:
			f.Attendees = append(f.Attendees, p)
		case comment:
			f.Comments = append(f.Comments, p)
		case freeBusyTime:
			f.FreeBusy = append(f.FreeBusy, p)
		case requestStatus:
			f.RequestStatus = append(f.RequestStatus, p)
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vFreeBusy {
				return ErrInvalidEnd
			}
			if !bm.Get(0) || !bm.Get(1) {
				return ErrRequiredMissing
			}
			c.FreeBusy = append(c.FreeBusy, f)
			return nil
		}
	}
	return nil
}

func (c *Calendar) writeFreeBusyData(e *Encoder) {
	for _, f := range c.FreeBusy {
		e.writeProperty(begin(vFreeBusy))
		e.writeProperty(dateStamp{dateTime{Time: f.LastModified}})
		e.writeProperty(uid(f.UID))
		if f.Contact.String != "" {
			e.writeProperty(f.Contact)
		}
		if !f.Start.IsZero() {
			e.writeProperty(dateTimeStart{f.Start})
		}
		if !f.End.IsZero() {
			e.writeProperty(dateTimeEnd{f.End})
		}
		if f.Organizer.Name != "" {
			e.writeProperty(f.Organizer)
		}
		if f.URL != "" {
			e.writeProperty(f.URL)
		}
		for _, p := range f.Attendees {
			e.writeProperty(p)
		}
		for _, p := range f.Comments {
			e.writeProperty(p)
		}
		for _, p := range f.FreeBusy {
			e.writeProperty(p)
		}
		for _, p := range f.RequestStatus {
			e.writeProperty(p)
		}
		e.writeProperty(end(vFreeBusy))
	}
}
