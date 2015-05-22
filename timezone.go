package ics

import (
	"time"

	"github.com/MJKWoolnough/bitmask"
)

const (
	vTimezone = "VTIMEZONE"
	standard  = "STANDARD"
	daylight  = "DAYLIGHT"
)

type Timezone struct {
	ID           string
	LastModified time.Time
	URL          timezoneURL
	Timezones    []TimezoneData
}

func (c *Calendar) decodeTimezone(d Decoder) error {
	bm := bitmask.New(4)
	var t Timezone
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case timezoneID:
			if !bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			t.ID = string(p)
		case lastModified:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			t.LastModified = p.Time
		case timezoneURL:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			t.URL = p
		case begin:
			var td TimezoneData
			switch p {
			case daylight:
				td.IsDaylight = true
				fallthrough
			case standard:
				bm.Set(3, true)
				err = td.decode(d)
				if err != nil {
					return err
				}
				t.Timezones = append(t.Timezones, td)
			default:
				if err = d.readUnknownComponent(string(p)); err != nil {
					return err
				}
			}
		case end:
			if p != vTimezone {
				return ErrInvalidEnd
			}
			if !bm.Get(0) || !bm.Get(3) {
				return ErrRequiredMissing
			}
			return nil
		}
	}
	return nil
}

type TimezoneData struct {
	IsDaylight      bool
	Start           dateTime
	OffsetFrom      timezoneOffsetFrom
	OffsetTo        timezoneOffsetTo
	RecurrenceRule  recurrenceRule
	Comments        []comment
	RecurrenceDates []recurrenceDate
	Names           map[string][]string
}

func (t *TimezoneData) decode(d Decoder) error {
	bm := bitmask.New(3)
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateTimeStart:
			if !bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			t.Start = p.dateTime
		case timezoneOffsetTo:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			t.OffsetTo = p
		case timezoneOffsetFrom:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			t.OffsetFrom = p
		case recurrenceRule:
			t.RecurrenceRule = p
		case comment:
			t.Comments = append(t.Comments, p)
		case recurrenceDate:
			t.RecurrenceDates = append(t.RecurrenceDates, p)
		case timezoneName:
			var names []string
			if n, ok := t.Names[p.Language]; ok {
				names = n
			}
			names = append(names, p.Name)
			t.Names[p.Language] = names
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if t.IsDaylight {
				if p != daylight {
					return ErrInvalidEnd
				}
			} else {
				if p != standard {
					return ErrInvalidEnd
				}
			}
			if !bm.Get(0) || !bm.Get(1) || !bm.Get(2) {
				return ErrRequiredMissing
			}
			return nil
		}
	}
}
