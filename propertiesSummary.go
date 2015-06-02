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

func (a attach) Validate() bool {
	return true
}

func (a attach) Data() propertyData {
	params := make(map[string]attribute)
	if a.Mime != "" {
		params[fmttypeparam] = fmtType(a.Mime)
	}
	if !a.URI {
		params[encodingparam] = encodingBase64
		params[valuetypeparam] = valueBinary
	}
	return propertyData{
		Name:   attachp,
		Params: params,
		Value:  string(a.Data),
	}
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

func (c categories) Validate() bool {
	return true
}

func (c categories) Data() propertyData {
	params := make(map[string]params)
	if c.Language != "" {
		params[languageparam] = language(c.Language)
	}
	val := make([]byte, 0, 1024)
	for n, cat := range c.Categories {
		if n > 0 {
			val = append(val, ',')
		}
		val = append(val, escape(cat)...)
	}
	return propertyData{
		Name:   categoriesp,
		Params: params,
		Value:  string(val),
	}
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

func (c class) String() string {
	switch c {
	case classPublic:
		return "PUBLIC"
	case classPrivate:
		return "PRIVATE"
	case classConfidential:
		return "CONFIDENTIAL"
	default:
		return "PRIVATE"
	}
}

func (c class) Validate() bool {
	return c == classPublic || c == classPrivate || c == classConfidential
}

func (c class) Data() propertyData {
	return propertyData{
		Name:  classp,
		Value: c.String(),
	}
}

type comment struct {
	altrepLanguageData
}

func (p *parser) readCommentProperty() (property, error) {
	a, err := p.readAltrepLanguageData()
	if err != nil {
		return nil, err
	}
	return comment{a}, nil
}

func (c comment) Data() propertyData {
	return c.data(commentp)
}

type description struct {
	altrepLanguageData
}

func (p *parser) readDescriptionProperty() (property, error) {
	a, err := p.readAltrepLanguageData()
	if err != nil {
		return nil, err
	}
	return description{a}, nil
}

func (d description) Data() propertyData {
	return c.data(descriptionp)
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

func (g geo) Validate() bool {
	return true
}

func (g geo) Data() propertyData {
	return propertyData{
		Name:  geop,
		Value: strconv.FormatFloat(g.Latitude, 'f', -1, 64) + ";" + strconv.FormatFloat(g.Longitude, 'f', -1, 64),
	}
}

type location struct {
	altrepLanguageData
}

func (p *parser) readLocationProperty() (property, error) {
	a, err := p.readAltrepLanguageData()
	if err != nil {
		return nil, err
	}
	return location{a}, nil
}

func (l location) Data() propertyData {
	return c.data(locationp)
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

func (p percentComplete) Validate() bool {
	return p >= 0 && p <= 100
}

func (p percentComplete) Data() propertyData {
	return propertyData{
		Name:  percentcompletep,
		Value: strconv.Itoa(int(p)),
	}
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

func (p priority) Validate() bool {
	return p >= 0 && p <= 9
}

func (p percentComplete) Data() propertyData {
	return propertyData{
		Name:  priorityp,
		Value: strconv.Itoa(int(p)),
	}
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

func (r resources) Validate() bool {
	return true
}

func (r resources) Data() propertyData {
	params := make(map[string]attributes)
	if a.AltRep != "" {
		params[altrepparam] = altrep(a.AltRep)
	}
	if a.Language != "" {
		params[languageparam] = language(a.Language)
	}
	val := make([]byte, 0, 1024)
	for n, res := range r.Resources {
		if n > 0 {
			val = append(val, ',')
		}
		val = append(val, encode(res)...)
	}
	return propertyData{
		Name:   resourcesp,
		Params: params,
		Value:  string(val),
	}
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

func (s status) String() string {
	switch s {
	case statusTentative:
		return "TENTATIVE"
	case statusConfirmed:
		return "CONFIRMED"
	case statusNeedsAction:
		return "NEED-ACTION"
	case statusCompleted:
		return "COMPLETED"
	case statusInProgress:
		return "IN-PROGRESS"
	case statusDraft:
		return "DRAFT"
	case statusFinal:
		return "FINAL"
	case statusCancelled:
		return "CANCELLED"
	}
}

func (s status) Validate() bool {
	return s >= statusTentative && s <= statusCancelled
}

func (s status) Data() propertyData {
	return propertyData{
		Name:  statusp,
		Value: s.String(),
	}
}

type summary altrepLanguageData

func (p *parser) readSummaryProperty() (property, error) {
	a, err := p.readAltrepLanguageData()
	if err != nil {
		return nil, err
	}
	return summary{a}, nil
}

func (s summary) Data() propertyData {
	return c.data(summaryp)
}
