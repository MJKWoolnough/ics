package ics

import (
	"encoding/base64"
	"strconv"
)

type attach struct {
	URI  bool
	Mime string
	Data []byte
}

func (p *parser) readAttachProperty() (property, error) {
	as, err := p.readAttributes(fmttypeparam, encodingparam, valuetypeparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	uri := true
	enc, encOK := as[encodingparam]
	val, valOK := as[valuetypeparam]
	var data []byte
	if encOK && valOK {
		uri = false
		if enc.(encoding) != encodingBase64 || val.(value) != valueBinary {
			return nil, ErrUnsupportedValue
		}
		data, err = base64.StdEncoding.DecodeString(v)
		if err != nil {
			return nil, err
		}

	} else if encOK == valOK {
		data = []byte(unescape(v))
	} else {
		return nil, ErrInvalidAttributeCombination
	}
	return attach{
		uri,
		string(as[fmttypeparam].(fmtType)),
		data,
	}, nil
}

type categories struct {
	Language   string
	Categories []string
}

func (p *parser) readCategoriesProperty() (property, error) {
	as, err := p.readAttributes(languageparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	language := ""
	if l, ok := as[languageparam]; ok {
		language = l.String()
	}
	return categories{
		language,
		textSplit(v, ','),
	}, nil
}

const (
	classPublic class = iota
	classPrivate
	classConfidential
)

type class int

func (p *parser) readClassProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "PUBLIC":
		return classPublic, nil
	case "PRIVATE":
		return classPrivate, nil
	case "CONFIDENTIAL":
		return classConfidential, nil
	default:
		return classPrivate, nil
	}
}

type comment altrepLanguageData

func (p *parser) readCommentProperty() (property, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return comment{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type description altrepLanguageData

func (p *parser) readDescriptionProperty() (property, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return description{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type geo struct {
	Latitude, Longitude float64
}

func (p *parser) readGeoProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	parts := textSplit(v, ';')
	if len(parts) != 2 {
		return nil, ErrUnsupportedValue
	}
	la, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return nil, err
	}
	lo, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return nil, err
	}
	return geo{la, lo}, nil
}

type location altrepLanguageData

func (p *parser) readLocationProperty() (property, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return location{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}

type percentComplete int

func (p *parser) readPercentCompleteProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	pc, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	if pc < 0 || pc > 100 {
		return nil, ErrUnsupportedValue
	}
	return percentComplete(pc), nil
}

type priority int

func (p *parser) readPriorityProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	pc, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	if pc < 0 || pc > 9 {
		return nil, ErrUnsupportedValue
	}
	return priority(pc), nil
}

type resources struct {
	Altrep, Language string
	Resources        []string
}

func (p *parser) readResourcesProperty() (property, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return resources{
		altRep,
		languageStr,
		textSplit(v, ','),
	}, nil
}

const (
	statusTentative status = iota
	statusConfirmed
	statusNeedsAction
	statusCompleted
	statusInProgress
	statusDraft
	statusFinal
	statusCancelled
)

type status int

func (p *parser) readStatusProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	switch v {
	case "TENTATIVE":
		return statusTentative, nil
	case "CONFIRMED":
		return statusConfirmed, nil
	case "NEED-ACTION":
		return statusNeedsAction, nil
	case "COMPLETED":
		return statusCompleted, nil
	case "IN-PROGRESS":
		return statusInProgress, nil
	case "DRAFT":
		return statusDraft, nil
	case "FINAL":
		return statusFinal, nil
	case "CANCELLED":
		return statusCancelled, nil
	default:
		return nil, ErrUnsupportedValue
	}
}

type summary altrepLanguageData

func (p *parser) readSummaryProperty() (property, error) {
	as, err := p.readAttributes(altrepparam, languageparam)
	if err != nil {
		return nil, err
	}
	var altRep, languageStr string
	if alt, ok := as[altrepparam]; ok {
		altRep = string(alt.(altrep))
	}
	if l, ok := as[languageparam]; ok {
		languageStr = string(l.(language))
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	return summary{
		altRep,
		languageStr,
		string(unescape(v)),
	}, nil
}
