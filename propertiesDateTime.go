package ics

import (
	"strings"
	"time"
)

type completed struct {
	time.Time
}

func (p *parser) readCompletedProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var c completed
	c.Time, err = time.ParseInLocation("20060102T150405Z", v, time.UTC)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type dateTimeEnd struct {
	dateTime
}

func (p *parser) readDateTimeOrTime() (t dateTime, err error) {
	as, err := p.readAttributes(tzidparam, valuetypeparam)
	if err != nil {
		return t, err
	}
	var (
		l        *time.Location
		justDate bool
	)
	if tzid, ok := as[tzidparam]; ok {
		l, err = time.LoadLocation(tzid.String())
		if err != nil {
			return t, err
		}
	}
	if v, ok := as[valuetypeparam]; ok {
		val := v.(value)
		switch val {
		case valueDate:
			justDate = true
		case valueDateTime:
			justDate = false
		default:
			return t, ErrUnsupportedValue
		}
	}
	v, err := p.readValue()
	if err != nil {
		return t, err
	}
	if justDate {
		t, err = parseDate(v)
	} else {
		t, err = parseDateTime(v, l)
	}
	return t, err
}

func (p *parser) readDateTimeEndProperty() (property, error) {
	t, err := p.readDateTimeOrTime()
	if err != nil {
		return nil, err
	}
	return dateTimeEnd{t}, nil
}

type dateTimeDue struct {
	dateTime dateTime
}

func (p *parser) readDateTimeDueProperty() (property, error) {
	t, err := p.readDateTimeOrTime()
	if err != nil {
		return nil, err
	}
	return dateTimeDue{t}, nil
}

type dateTimeStart struct {
	dateTime dateTime
}

func (p *parser) readDateTimeStartProperty() (property, error) {
	t, err := p.readDateTimeOrTime()
	if err != nil {
		return nil, err
	}
	return dateTimeStart{t}, nil
}

type duration struct {
	time.Duration
}

func (p *parser) readDurationProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var d duration
	d.Duration, err = parseDuration(v)
	if err != nil {
		return nil, err
	}
	return d, nil
}

type freeBusyTime struct {
	Typ     freeBusy
	Periods []period
}

type period struct {
	FixedDuration bool
	Start, End    dateTime
}

func parsePeriods(v string, l *time.Location) ([]period, error) {
	periods := make([]period, 0, 1)

	for _, pd := range textSplit(v, ',') {
		parts := strings.Split(pd, "/")
		if len(parts) != 2 {
			return nil, ErrUnsupportedValue
		}
		if parts[0][len(parts[0])-1] != 'Z' {
			return nil, ErrUnsupportedValue
		}
		start, err := parseDateTime(parts[0], l)
		if err != nil {
			return nil, err
		}
		var (
			end           dateTime
			fixedDuration bool
		)
		if parts[1][len(parts[1])-1] == 'Z' {
			end, err = parseDateTime(parts[1], l)
			if err != nil {
				return nil, err
			}
		} else {
			d, err := parseDuration(parts[1])
			if err != nil {
				return nil, err
			}
			if d < 0 {
				return nil, ErrUnsupportedValue
			}
			end = start.Add(d)
			fixedDuration = true
		}
		periods = append(periods, period{fixedDuration, start, end})
	}
	return periods, nil
}

func (p *parser) readFreeBusyTimeProperty() (property, error) {
	as, err := p.readAttributes(fbtypeparam)
	if err != nil {
		return nil, err
	}
	var fb freeBusy
	if f, ok := as[fbtypeparam]; ok {
		fb = f.(freeBusy)
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	periods, err := parsePeriods(v, nil)
	if err != nil {
		return nil, err
	}
	return freeBusyTime{
		Typ:     fb,
		Periods: periods,
	}, nil
}

const (
	TTOpaque timeTransparency = iota
	TTTransparent
)

type timeTransparency int

func (p *parser) readTimeTransparencyProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "OPAQUE":
		return TTOpaque, nil
	case "TRANSPARENT":
		return TTTransparent, nil
	default:
		return nil, ErrUnsupportedValue
	}
}
