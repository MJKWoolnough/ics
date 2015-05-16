package ics

type calscale string

func (p *parser) readCalScaleProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return calscale(unescape(v)), nil
}

type method string

func (p *parser) readMethodProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return method(unescape(v)), nil
}

type productID string

func (p *parser) readProductIDProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return productID(unescape(v)), nil
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
