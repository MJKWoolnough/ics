package ics

//type timezoneID string //(in attributes)

func (p *parser) readTimezoneIDProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneID(v), nil
}

func (t timezoneID) Validate() bool {
	return true
}

func (t timezoneID) Data() propertyData {
	return propertyData{
		Name:  tzidp,
		Value: string(t),
	}
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

func (t timezoneName) Validate() bool {
	return true
}

func (t timezoneName) Data() propertyData {
	params := make(map[string]attribute)
	if t.Language != "" {
		params[languageparam] = language(t.Language)
	}
	return propertyData{
		Name:   tznamep,
		Params: params,
		Value:  t.Name,
	}
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

func (t timezoneOffsetFrom) Validate() bool {
	return true
}

func (t timezoneOffsetFrom) Data() propertyData {
	return propertyData{
		Name:  tzoffsetfromp,
		Value: offsetString(int(t)),
	}
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

func (t timezoneOffsetTo) Validate() bool {
	return true
}

func (t timezoneOffsetTo) Data() propertyData {
	return propertyData{
		Name:  tzoffsettop,
		Value: offsetString(int(t)),
	}
}

type timezoneURL string

func (p *parser) readTimezoneURLProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return timezoneURL(v), nil
}

func (t timezoneURL) Validate() bool {
	return true
}

func (t timezoneURL) Data() propertyData {
	return propertyData{
		Name:  tzurlp,
		Value: string(t),
	}
}
