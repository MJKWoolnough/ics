package ics

import (
	"strings"
	"time"
)

type completed struct {
	dateTime
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

func (c completed) Validate() bool {
	return c.Location() == time.UTC
}

func (c completed) Data() propertyData {
	return propertyData{
		Name:  completedp,
		Value: c.String(),
	}
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

func dateTimeOrTimeData(name string, d dateTime) propertyData {
	params := make(map[string]attribute)
	if d.justDate {
		params[valuetypeparam] = valueDate
	}
	if d.Location() != time.UTC {
		params[tzidparam] = timezoneID(d.Location().String())
	}
	return propertyData{
		Name:   name,
		Params: params,
		Value:  d.String(),
	}
}

func (p *parser) readDateTimeEndProperty() (property, error) {
	t, err := p.readDateTimeOrTime()
	if err != nil {
		return nil, err
	}
	return dateTimeEnd{t}, nil
}

func (d dateTimeEnd) Validate() bool {
	return true
}

func (d dateTimeEnd) Data() propertyData {
	return dateTimeOrTimeData(dtendp, d.dateTime)
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

func (d dateTimeDue) Validate() bool {
	return true
}

func (d dateTimeDue) Data() propertyData {
	return dateTimeOrTimeData(duep, d.dateTime)
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

func (d dateTimeStart) Validate() bool {
	return true
}

func (d dateTimeStart) Data() propertyData {
	return dateTimeOrTimeData(dtstartp, d.dateTime)
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

func (d duration) Validate() bool {
	return true
}

func (d duration) Data() propertyData {
	return propertyData{
		Name:  durationp,
		Value: durationString(d.Duration),
	}
}

type freeBusyTime struct {
	Typ     freeBusy
	Periods []period
}

type period struct {
	FixedDuration bool
	Start, End    dateTime
}

func (p period) Bytes() []byte {
	val := make([]byte, 0, 64)
	val = append(val, p.Start.String()...)
	val = append(val, '/')
	if p.FixedDuration {
		val = append(val, durationString(p.End.Sub(p.Start.Time))...)
	} else {
		val = append(val, p.End.String()...)
	}
	return val
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

func (f freeBusyTime) Validate() bool {
	return f.Typ >= fbBusy && f.Typ <= fbBusyTentative
}

func (f freeBusyTime) Data() propertyData {
	params := make(map[string]attribute)
	params[fbtypeparam] = f.Typ
	val := make([]byte, 0, len(f.Periods)*64)
	for n, period := range f.Periods {
		if n > 0 {
			val = append(val, ',')
		}
		val = append(val, period.Bytes()...)
	}
	return propertyData{
		Name:   freebusyp,
		Params: params,
		Value:  string(val),
	}
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

func (t timeTransparency) String() string {
	switch t {
	case TTOpaque:
		return "OPAQUE"
	case TTTransparent:
		return "TRANSPARENT"
	default:
		return "UNKNOWN"
	}
}

func (t timeTransparency) Validate() bool {
	switch t {
	case TTOpaque, TTTransparent:
		return true
	default:
		return false
	}
}

func (t timeTransparency) Data() propertyData {
	return propertyData{
		Name:  transpp,
		Value: t.String(),
	}
}
