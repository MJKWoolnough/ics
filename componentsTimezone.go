package ics

import "time"

//type timeZoneID string //(in attributes)

func (p *parser) readTimezoneIDComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneID(v), nil
}

type timezoneName struct {
	Language, Name string
}

func (p *parser) readTimezoneNameComponent() (component, error) {
	as, err := p.readAttributes(languageparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var languageStr string
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	return timezoneName{languageStr, v}, nil
}

type timezoneOffsetFrom int

func (p *parser) readTimezoneOffsetFromComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	tzo, err := parseOffset(v)
	if err != nil {
		return nil, err
	}
	return timezoneOffsetFrom(tzo), nil
}

type timezoneOffsetTo struct {
	time.Duration
}

func (p *parser) readTimezoneOffsetToComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var tzo timezoneOffsetTo
	tzo.Duration, err = parseOffset(v)
	if err != nil {
		return nil, err
	}
	return timezoneOffsetTo(tzo), nil
}

type timezoneURL string

func (p *parser) readTimezoneURLComponent() (component, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneURL(v), nil
}
