package ics

//type timeZoneID string //(in attributes)

func (p *parser) readTimezoneIDProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneID(v), nil
}

type timezoneName struct {
	Language, Name string
}

func (p *parser) readTimezoneNameProperty() (property, error) {
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

func (p *parser) readTimezoneOffsetFromProperty() (property, error) {
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

type timezoneOffsetTo int

func (p *parser) readTimezoneOffsetToProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	tzo, err := parseOffset(v)
	if err != nil {
		return nil, err
	}
	return timezoneOffsetTo(tzo), nil
}

type timezoneURL string

func (p *parser) readTimezoneURLProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneURL(v), nil
}
