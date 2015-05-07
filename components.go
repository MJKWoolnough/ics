package ics

import (
	"errors"
	"strings"
)

type componentType int

type component interface {
	setAttribute(attribute) error
	setValue(string) error
}

func componentFromToken(s string) component {
	switch s {
	case "BEGIN":
		return new(beginComponent)
	case "END":
		return new(endComponent)
	case "CALSCALE":
		return new(calScale)
	case "METHOD":
		return new(method)
	case "PRODID":
		return new(productID)
	case "VERSION":
		return new(version)
	default:
		return &unknownComponent{
			Name: s,
		}
	}
}

const (
	beVCalendar = "VCALENDAR"
	beVEvent    = "VEVENT"
	beVTodo     = "VTODO"
	beVJournal  = "VJOURNAL"
	beVFreeBusy = "VFREEBUSY"
	beVTimezone = "VTIMEZONE"
	beStandard  = "STANDARD"
	beDaylight  = "DAYLIGHT"
	beVAlarm    = "VALARM"
)

type beginComponent string

func (beginComponent) setAttribute(_ attribute) error {
	return ErrNoAttributes
}

func (b *beginComponent) setValue(v string) error {
	*b = beginComponent(v)
	return nil
}

type endComponent string

func (endComponent) setAttribute(_ attribute) error {
	return ErrNoAttributes
}

func (e *endComponent) setValue(v string) error {
	*e = endComponent(v)
	return nil
}

type calScale string

func (calScale) setAttribute(_ attribute) error {
	return nil
}

func (c *calScale) setValue(v string) error {
	*c = calScale(v)
	return nil
}

type method string

func (method) setAttribute(_ attribute) error {
	return nil
}

func (m *method) setValue(v string) error {
	*m = method(v)
	return nil
}

type productID string

func (productID) setAttribute(_ attribute) error {
	return nil
}

func (p *productID) setValue(v string) error {
	*p = productID(v)
	return nil
}

type version [2]string

func (version) setAttribute(_ attribute) error {
	return nil
}

func (ver *version) setValue(v string) error {
	parts := strings.SplitN(v, ";", 2)
	if len(parts) == 1 {
		ver[0] = v
		ver[1] = v
	} else {
		ver[0] = parts[0]
		ver[1] = parts[1]
	}
	return nil
}

type unknownComponent struct {
	Name, Value string
	Attributes  []attribute
}

func (u *unknownComponent) setAttribute(attr attribute) error {
	u.Attributes = append(u.Attributes, attr)
	return nil
}

func (u *unknownComponent) setValue(v string) error {
	u.Value = v
	return nil
}

// Errors

var (
	ErrNoAttributes = errors.New("component doesn't allow attributes")
)
