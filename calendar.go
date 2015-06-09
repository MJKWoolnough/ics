package ics

import (
	"errors"

	"github.com/MJKWoolnough/bitmask"
)

const vCalendar = "VCALENDAR"

type Calendar struct {
	ProductID, Method string
	Events            []Event
	Todo              []Todo
	Journals          []Journal
	FreeBusy          []FreeBusy
	Timezones         []Timezone
}

func (c *Calendar) decode(d Decoder) error {
	bm := bitmask.New(4)
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case productID:
			if bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			c.ProductID = string(p)
		case version:
			if bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			if p.Min != "2.0" && p.Max != "2.0" {
				return ErrUnsupportedVersion
			}
		case calscale:
			if bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			if p != "GREGORIAN" {
				return ErrUnsupportedCalendar
			}
		case method:
			if bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			c.Method = string(p)
		case begin:
			if !bm.Get(0) || !bm.Get(1) {
				return ErrRequiredMissing
			}
			switch p {
			case vEvent:
				if err = c.decodeEvent(d); err != nil {
					return err
				}
			case vTodo:
				if err = c.decodeTodo(d); err != nil {
					return err
				}
			case vJournal:
				if err = c.decodeJournal(d); err != nil {
					return err
				}
			case vFreeBusy:
				if err = c.decodeFreeBusy(d); err != nil {
					return err
				}
			case vTimezone:
				if err = c.decodeTimezone(d); err != nil {
					return err
				}
			default:
				if err = d.readUnknownComponent(string(p)); err != nil {
					return err
				}
			}
		case end:
			if !bm.Get(0) || !bm.Get(1) {
				return ErrRequiredMissing
			}
			if p != vCalendar {
				return ErrInvalidEnd
			}
			return nil
		}
	}
}

func (c *Calendar) encode(e *Encoder) {
	e.writeProperty(begin(vCalendar))
	e.writeProperty(productID(c.ProductID))
	e.writeProperty(version{"2.0", "2.0"})
	if c.Method != "" {
		e.writeProperty(method(c.Method))
	}
	c.writeTimezoneData(e)
	c.writeEventData(e)
	c.writeFreeBusyData(e)
	c.writeJournalData(e)
	c.writeTodoData(e)
	e.writeProperty(end(vCalendar))
}

// Errors

var (
	ErrUnsupportedCalendar = errors.New("unsupported calendar")
	ErrUnsupportedVersion  = errors.New("unsupported ics version")
)
