package ics

type calscale string

func (p *parser) readCalScaleProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return calscale(unescape(v)), nil
}

func (c calscale) Validate() bool {
	return true
}

func (c calscale) Data() propertyData {
	return propertyData{
		Name:  calscalep,
		Value: string(escape(string(c))),
	}
}

type method string

func (p *parser) readMethodProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return method(unescape(v)), nil
}

func (m method) Validate() bool {
	return true
}

func (m method) Data() propertyData {
	return propertyData{
		Name:  methodp,
		Value: string(escape(string(m))),
	}
}

type productID string

func (p *parser) readProductIDProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return productID(unescape(v)), nil
}

func (p productID) Validate() bool {
	return true
}

func (p productID) Data() propertyData {
	return propertyData{
		Name:  prodidp,
		Value: string(escape(string(p))),
	}
}

type version struct {
	Min, Max string
}

func (p *parser) readVersionProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	parts := textSplit(v, ';')
	if len(parts) > 2 {
		return nil, ErrUnsupportedValue
	} else if len(parts) == 2 {
		return version{parts[0], parts[1]}, nil
	} else {
		return version{parts[0], parts[0]}, nil
	}
}

func (v version) Validate() bool {
	return true
}

func (v version) Data() propertyData {
	var val string
	if v.Min != v.Max {
		val = v.Min
	} else {
		val = v.Min + ";" + v.Max
	}
	return propertyData{
		Name:  versionp,
		Value: escape(val),
	}
}
