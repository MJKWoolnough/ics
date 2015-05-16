package ics

import (
	"strconv"
	"time"
)

const (
	actionUnknown action = iota
	actionAudio
	actionDisplay
	actionEmail
)

type action int

func (p *parser) readActionProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "AUDIO":
		return actionAudio, nil
	case "DISPLAY":
		return actionDisplay, nil
	case "EMAIL":
		return actionEmail, nil
	default:
		return actionUnknown, nil
	}
}

type repeat int

func (p *parser) readRepeatProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	return repeat(n), nil
}

type related int

type trigger struct {
	DateTime time.Time
	Related  alarmTriggerRelationship
	Duration time.Duration
}

func (p *parser) readTriggerProperty() (property, error) {
	as, err := p.readAttributes(valuetypeparam, trigrelparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var t trigger
	if val, ok := as[valuetypeparam]; ok && val.(value) == valueDateTime {
		if v[len(v)-1] != 'Z' {
			return nil, ErrUnsupportedValue
		}
		t.DateTime, err = parseDateTime(v, nil)
		if err != nil {
			return nil, err
		}
		t.Related = -1
	} else {
		if rel, ok := as[trigrelparam]; ok {
			t.Related = rel.(alarmTriggerRelationship)
		} else {
			t.Related = atrStart
		}
		t.Duration, err = parseDuration(v)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
