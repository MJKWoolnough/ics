package ics

import "errors"

const vCalendar = "VCALENDAR"

type Calendar struct {
	ProductID, Method string
}

func (c *Calendar) decode(d Decoder) error {
	var pID, ver, cs, m bool
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case productID:
			if pID {
				return ErrMultipleUnique
			}
			pID = true
			c.ProductID = string(p)
		case version:
			if ver {
				return ErrMultipleUnique
			}
			ver = true
			if p.Min != "2.0" && p.Max != "2.0" {
				return ErrUnsupportedVersion
			}
		case calscale:
			if cs {
				return ErrMultipleUnique
			}
			cs = true
			if p != "GREGORIAN" {
				return ErrUnsupportedCalendar
			}
		case method:
			if m {
				return ErrMultipleUnique
			}
			m = true
			// do something with value?
		case begin:
			if !pID || !ver {
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
			if !pID || !ver {
				return ErrRequiredMissing
			}
			if p != vCalendar {
				return ErrInvalidEnd
			}
			return nil
		}
	}
}

// Errors

var (
	ErrUnsupportedCalendar = errors.New("unsupported calendar")
	ErrUnsupportedVersion  = errors.New("unsupported ics version")
)
