package ics

import (
	"errors"
	"io"
)

type Decoder struct {
	p parser
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		p: newParser(r),
	}
}

func (d *Decoder) Decode() (*Calendar, error) {
	c, err := d.p.GetComponent()
	if err != nil {
		return nil, err
	}
	b, ok := c.(*beginComponent)
	if !ok || *b != beVCalendar {
		return nil, ErrInvalidComponent
	}
	return d.readCalendar()
}

func (d *Decoder) readUnknown(v string) error {
	for {
		c, err := d.p.GetComponent()
		if err != nil {
			return err
		}
		switch c := c.(type) {
		case *beginComponent:
			if err = d.readUnknown(string(*c)); err != nil {
				return err
			}
		case *endComponent:
			if string(*c) == v {
				return nil
			}
			return ErrInvalidComponent
		}
	}
}

func (d *Decoder) readCalendar() (*Calendar, error) {
	cal := new(Calendar)
	var (
		pID, cV, cS, m, o bool
		//methodStr = ""
		//calScaleStr = "GREGORIAN"
	)
	for {
		c, err := d.p.GetComponent()
		if err != nil {
			return nil, err
		}
		switch c := c.(type) {
		case *calScale:
			if o {
				return nil, ErrOutOfOrder
			} else if cS {
				return nil, ErrDuplicateComponent
			}
			cS = true
			//calScaleStr = string(*c)
		case *method:
			if o {
				return nil, ErrOutOfOrder
			} else if m {
				return nil, ErrDuplicateComponent
			}
			m = true
			//method = string(*c)
		case *productID:
			if pID {
				return nil, ErrDuplicateComponent
			}
			pID = true
			cal.SetProductID(string(*c))
		case *version:
			if v {
				return nil, ErrDuplicateComponent
			}
			v = true
			if *c[0] != "2.0" {
				return nil, ErrUnsupportedVersion
			}
		case *beginComponent:
			if !pID || !v {
				return nil, ErrMissingComponent
			}
			o = true
			switch *c {
			case beVEvent:
				e, err := d.readEvent()
				if err != nil {
					return nil, err
				}
				cal.AddEvent(e)
			case beVTodo:
				t, err := d.readTodo()
				if err != nil {
					return nil, err
				}
				cal.AddTodo(t)
			case beVJournal:
				j, err := d.readJournal()
				if err != nil {
					return nil, err
				}
				cal.AddJournal(j)
			case beVFreeBusy:
				fb, err := d.readFreeBusy()
				if err != nil {
					return nil, err
				}
				cal.AddFreeBusy(fb)
			case beVTimezone:
				tz, err := d.readTimezone()
				if err != nil {
					return nil, err
				}
				cal.AddTimezone(tz)
			default:
				err = d.readUnknown(string(*c))
				if err != nil {
					return nil, err
				}
			}
		case *endComponent:
			if *c == beVCalendar {
				if !pID || !v || !o {
					return nil, ErrMissingComponent
				}
				return cal, nil
			}
			return nil, ErrInvalidComponent
		}
	}
}

func (d *Decoder) readEvent() (*Event, error) {
	for {
		c, err := d.p.GetComponent()
		if err != nil {
			return nil, err
		}
		switch c := c.(type) {

		case *beginComponent:
			switch c {
			default:
				err = d.readUnknown(string(*c))
				if err != nil {
					return nil, err
				}
			}
		case *endComponent:
			if string(*c) == beVEvent {
				return nil, nil
			}
			return nil, ErrInvalidComponent
		}
	}
}

func (d *Decoder) readTodo() (*Todo, error) {
	for {
		c, err := d.p.GetComponent()
		if err != nil {
			return nil, err
		}
		switch c := c.(type) {
		case *beginComponent:
			switch c {
			default:
				err = d.readUnknown(string(*c))
				if err != nil {
					return nil, err
				}
			}
		case *endComponent:
			if string(*c) == beVTodo {
				return nil, nil
			}
			return nil, ErrInvalidComponent
		}
	}
}

func (d *Decoder) readJournal() (*Journal, error) {
	for {
		c, err := d.p.GetComponent()
		if err != nil {
			return nil, err
		}
		switch c := c.(type) {
		case *beginComponent:
			switch c {
			default:
				err = d.readUnknown(string(*c))
				if err != nil {
					return nil, err
				}
			}
		case *endComponent:
			if string(*c) == beVJournal {
				return nil, nil
			}
			return nil, ErrInvalidComponent
		}
	}
}

func (d *Decoder) readFreeBusy() (*FreeBusy, error) {
	for {
		c, err := d.p.GetComponent()
		if err != nil {
			return nil, err
		}
		switch c := c.(type) {
		case *beginComponent:
			switch c {
			default:
				err = d.readUnknown(string(*c))
				if err != nil {
					return nil, err
				}
			}
		case *endComponent:
			if string(*c) == beVFreeBusy {
				return nil, nil
			}
			return nil, ErrInvalidComponent
		}
	}
}

func (d *Decoder) readTimezone() (*Timezone, error) {
	for {
		c, err := d.p.GetComponent()
		if err != nil {
			return nil, err
		}
		switch c := c.(type) {
		case *beginComponent:
			switch c {
			default:
				err = d.readUnknown(string(*c))
				if err != nil {
					return nil, err
				}
			}
		case *endComponent:
			if string(*c) == beVTimezone {
				return nil, nil
			}
			return nil, ErrInvalidComponent
		}
	}
}

//Errors

var (
	ErrInvalidComponent   = errors.New("received invalid component")
	ErrMissingComponent   = errors.New("required component(s) missing")
	ErrOutOfOrder         = errors.New("component out of order")
	ErrDuplicateComponent = errors.New("received second instance of unique component")
	ErrUnsupportedVersion = errors.New("unsupported ics version")
)
