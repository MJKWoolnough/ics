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

func componentFromToken(t token) component {
	switch strings.ToUpper(t.data) {
	case "BEGIN":
		return new(beginComponent)
	case "END":
		return new(endComponent)
	default:
		return &unknownComponent{
			Name: t.data,
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
