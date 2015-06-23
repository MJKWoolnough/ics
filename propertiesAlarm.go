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

func (a action) Validate() bool {
	switch a {
	case actionAudio, actionDisplay, actionEmail:
		return true
	}
	return false
}

func (a action) Data() propertyData {
	return propertyData{
		Name:  actionp,
		Value: a.String(),
	}
}

func (a action) String() string {
	switch a {
	case actionAudio:
		return "AUDIO"
	case actionDisplay:
		return "DISPLAY"
	case actionEmail:
		return "EMAIL"
	default:
		return "UNKNOWN"
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

func (r repeat) Validate() bool {
	return true
}

func (r repeat) Data() propertyData {
	return propertyData{
		Name:  repeatp,
		Value: strconv.Itoa(int(r)),
	}
}

//type related int

type trigger struct {
	DateTime dateTime
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

func (t trigger) Validate() bool {
	return true
}

func (t trigger) Data() propertyData {
	params := make(map[string]attribute)
	if t.Related != atrStart {
		params[trigrelparam] = t.Related
	}
	var val string
	if t.DateTime.IsZero() {
		val = durationString(t.Duration)
	} else {
		params[valuetypeparam] = valueDateTime
		val = t.DateTime.String()
	}
	return propertyData{
		Name:   triggerp,
		Params: params,
		Value:  val,
	}
}
