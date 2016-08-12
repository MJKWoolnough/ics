package ics

// File automatically generated with ./genParams.sh

import (
	"strings"
)

type Action uint8

const (
	ActionAudio Action = iota
	ActionDisplay
	ActionEmail
)

func (p *Action) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "AUDIO":
		*p = ActionAudio
	case "DISPLAY":
		*p = ActionDisplay
	case "EMAIL":
		*p = ActionEmail
	default:
		return ErrInvalidValue
	}
	return nil
}

func (p *Action) encode(w writer) {
	w.WriteString("ACTION:")
	switch *p {
	case Audio:
		w.WriteString("AUDIO")
	case Display:
		w.WriteString("DISPLAY")
	case Email:
		w.WriteString("EMAIL")
	}
	w.WriteString("\r\n")
}

func (p *Action) valid() error {
	switch *p {
	case Audio, Display, Email:
	default:
		return ErrInvalidValue
	}
	return nil
}

type Attachment struct {
	FormatType *FormatType
	Uri        *Uri
	Binary     Binary
}

func (p *Attachment) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "FMTTYPE":
			if p.FormatType != nil {
				return ErrDuplicateParam
			}
			p.FormatType = new(FormatType)
			if err := p.FormatType.decode(pValues); err != nil {
				return err
			}
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "URI":
				vType = 0
			case "Binary":
				vType = 1
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.Uri = new(Uri)
		if err := p.Uri.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		if err := p.Binary.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *Attachment) encode(w writer) {
	w.WriteString("ATTACH")
	if p.FormatType != nil {
		p.FormatType.encode(w)
	}
	if p.Uri != nil {
		p.Uri.aencode(w)
	}
	if p.Binary != nil {
		p.Binary.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *Attachment) valid() error {
	if p.FormatType != nil {
		if err := p.FormatType.valid(); err != nil {
			return err
		}
	}
	c := 0
	if p.Uri != nil {
		if err := p.Uri.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Binary != nil {
		if err := p.Binary.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type Attendee struct {
	CalendarUserType    *CalendarUserType
	Member              Member
	ParticipationRole   *ParticipationRole
	ParticipationStatus *ParticipationStatus
	RSVP                *RSVP
	Delagatee           Delagatee
	Delegator           Delegator
	SentBy              *SentBy
	CommonName          *CommonName
	DirectoryEntry      *DirectoryEntry
	Language            *Language
	CalAddress
}

func (p *Attendee) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "CUTYPE":
			if p.CalendarUserType != nil {
				return ErrDuplicateParam
			}
			p.CalendarUserType = new(CalendarUserType)
			if err := p.CalendarUserType.decode(pValues); err != nil {
				return err
			}
		case "MEMBER":
			if p.Member != nil {
				return ErrDuplicateParam
			}
			if err := p.Member.decode(pValues); err != nil {
				return err
			}
		case "ROLE":
			if p.ParticipationRole != nil {
				return ErrDuplicateParam
			}
			p.ParticipationRole = new(ParticipationRole)
			if err := p.ParticipationRole.decode(pValues); err != nil {
				return err
			}
		case "PARTSTAT":
			if p.ParticipationStatus != nil {
				return ErrDuplicateParam
			}
			p.ParticipationStatus = new(ParticipationStatus)
			if err := p.ParticipationStatus.decode(pValues); err != nil {
				return err
			}
		case "RSVP":
			if p.RSVP != nil {
				return ErrDuplicateParam
			}
			p.RSVP = new(RSVP)
			if err := p.RSVP.decode(pValues); err != nil {
				return err
			}
		case "DELEGATED-TO":
			if p.Delagatee != nil {
				return ErrDuplicateParam
			}
			if err := p.Delagatee.decode(pValues); err != nil {
				return err
			}
		case "DELEGATED-FROM":
			if p.Delegator != nil {
				return ErrDuplicateParam
			}
			if err := p.Delegator.decode(pValues); err != nil {
				return err
			}
		case "SENT-BY":
			if p.SentBy != nil {
				return ErrDuplicateParam
			}
			p.SentBy = new(SentBy)
			if err := p.SentBy.decode(pValues); err != nil {
				return err
			}
		case "CN":
			if p.CommonName != nil {
				return ErrDuplicateParam
			}
			p.CommonName = new(CommonName)
			if err := p.CommonName.decode(pValues); err != nil {
				return err
			}
		case "DIR":
			if p.DirectoryEntry != nil {
				return ErrDuplicateParam
			}
			p.DirectoryEntry = new(DirectoryEntry)
			if err := p.DirectoryEntry.decode(pValues); err != nil {
				return err
			}
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Caladdress.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Attendee) encode(w writer) {
	w.WriteString("ATTENDEE")
	if p.CalendarUserType != nil {
		p.CalendarUserType.encode(w)
	}
	if p.Member != nil {
		p.Member.encode(w)
	}
	if p.ParticipationRole != nil {
		p.ParticipationRole.encode(w)
	}
	if p.ParticipationStatus != nil {
		p.ParticipationStatus.encode(w)
	}
	if p.RSVP != nil {
		p.RSVP.encode(w)
	}
	if p.Delagatee != nil {
		p.Delagatee.encode(w)
	}
	if p.Delegator != nil {
		p.Delegator.encode(w)
	}
	if p.SentBy != nil {
		p.SentBy.encode(w)
	}
	if p.CommonName != nil {
		p.CommonName.encode(w)
	}
	if p.DirectoryEntry != nil {
		p.DirectoryEntry.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Caladdress.aencode(w)
	w.WriteString("\r\n")
}

func (p *Attendee) valid() error {
	if p.CalendarUserType != nil {
		if err := p.CalendarUserType.valid(); err != nil {
			return err
		}
	}
	if p.Member != nil {
		if err := p.Member.valid(); err != nil {
			return err
		}
	}
	if p.ParticipationRole != nil {
		if err := p.ParticipationRole.valid(); err != nil {
			return err
		}
	}
	if p.ParticipationStatus != nil {
		if err := p.ParticipationStatus.valid(); err != nil {
			return err
		}
	}
	if p.RSVP != nil {
		if err := p.RSVP.valid(); err != nil {
			return err
		}
	}
	if p.Delagatee != nil {
		if err := p.Delagatee.valid(); err != nil {
			return err
		}
	}
	if p.Delegator != nil {
		if err := p.Delegator.valid(); err != nil {
			return err
		}
	}
	if p.SentBy != nil {
		if err := p.SentBy.valid(); err != nil {
			return err
		}
	}
	if p.CommonName != nil {
		if err := p.CommonName.valid(); err != nil {
			return err
		}
	}
	if p.DirectoryEntry != nil {
		if err := p.DirectoryEntry.valid(); err != nil {
			return err
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Caladdress.valid(); err != nil {
		return err
	}
	return nil
}

type CalendarScale uint8

const (
	CalendarScaleGregorian CalendarScale = iota
)

func (p *CalendarScale) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "GREGORIAN":
		*p = CalendarScaleGregorian
	default:
		return ErrInvalidValue
	}
	return nil
}

func (p *CalendarScale) encode(w writer) {
	w.WriteString("CALSCALE:")
	switch *p {
	case Gregorian:
		w.WriteString("GREGORIAN")
	}
	w.WriteString("\r\n")
}

func (p *CalendarScale) valid() error {
	switch *p {
	case Gregorian:
	default:
		return ErrInvalidValue
	}
	return nil
}

type Catergories struct {
	Language *Language
	MText
}

func (p *Catergories) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Mtext.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Catergories) encode(w writer) {
	w.WriteString("CATERGORIES")
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Mtext.aencode(w)
	w.WriteString("\r\n")
}

func (p *Catergories) valid() error {
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Mtext.valid(); err != nil {
		return err
	}
	return nil
}

type Class uint8

const (
	ClassPublic Class = iota
	ClassPrivate
	ClassConfidential
)

func (p *Class) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "PUBLIC":
		*p = ClassPublic
	case "PRIVATE":
		*p = ClassPrivate
	case "CONFIDENTIAL":
		*p = ClassConfidential
	default:
		return ErrInvalidValue
	}
	return nil
}

func (p *Class) encode(w writer) {
	w.WriteString("CLASS:")
	switch *p {
	case Public:
		w.WriteString("PUBLIC")
	case Private:
		w.WriteString("PRIVATE")
	case Confidential:
		w.WriteString("CONFIDENTIAL")
	}
	w.WriteString("\r\n")
}

func (p *Class) valid() error {
	switch *p {
	case Public, Private, Confidential:
	default:
		return ErrInvalidValue
	}
	return nil
}

type Comment struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *Comment) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "ALTREP":
			if p.AlternativeRepresentation != nil {
				return ErrDuplicateParam
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return err
			}
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Text.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Comment) encode(w writer) {
	w.WriteString("COMMENT")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *Comment) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return err
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Text.valid(); err != nil {
		return err
	}
	return nil
}

type Completed DateTime

func (p *Completed) decode(params []parser.Token, value string) error {
	var t DateTime
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Completed(t)
	return nil
}

func (p *Completed) encode(w writer) {
	w.WriteString("COMPLETED")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Completed) valid() error {
	t := DateTime(*p)
	return t.valid()
}

type Contact struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *Contact) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "ALTREP":
			if p.AlternativeRepresentation != nil {
				return ErrDuplicateParam
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return err
			}
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Text.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Contact) encode(w writer) {
	w.WriteString("CONTACT")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *Contact) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return err
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Text.valid(); err != nil {
		return err
	}
	return nil
}

type Created DateTime

func (p *Created) decode(params []parser.Token, value string) error {
	var t DateTime
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Created(t)
	return nil
}

func (p *Created) encode(w writer) {
	w.WriteString("CREATED")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Created) valid() error {
	t := DateTime(*p)
	return t.valid()
}

type Description struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *Description) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "ALTREP":
			if p.AlternativeRepresentation != nil {
				return ErrDuplicateParam
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return err
			}
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Text.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Description) encode(w writer) {
	w.WriteString("DESCRIPTION")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *Description) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return err
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Text.valid(); err != nil {
		return err
	}
	return nil
}

type DateTimeEnd struct {
	DateTime *DateTime
	Date     *Date
}

func (p *DateTimeEnd) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *DateTimeEnd) encode(w writer) {
	w.WriteString("DTEND")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *DateTimeEnd) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type DateTimeStamp DateTime

func (p *DateTimeStamp) decode(params []parser.Token, value string) error {
	var t DateTime
	if err := t.decode(value); err != nil {
		return err
	}
	*p = DateTimeStamp(t)
	return nil
}

func (p *DateTimeStamp) encode(w writer) {
	w.WriteString("DTSTAMP")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *DateTimeStamp) valid() error {
	t := DateTime(*p)
	return t.valid()
}

type DateTimeStart struct {
	DateTime *DateTime
	Date     *Date
}

func (p *DateTimeStart) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *DateTimeStart) encode(w writer) {
	w.WriteString("DTSTART")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *DateTimeStart) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type Due struct {
	DateTime *DateTime
	Date     *Date
}

func (p *Due) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *Due) encode(w writer) {
	w.WriteString("DUE")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *Due) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type Duration Duration

func (p *Duration) decode(params []parser.Token, value string) error {
	var t Duration
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Duration(t)
	return nil
}

func (p *Duration) encode(w writer) {
	w.WriteString("DURATION")
	t := Duration(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Duration) valid() error {
	t := Duration(*p)
	return t.valid()
}

type ExceptionDateTime struct {
	DateTime *DateTime
	Date     *Date
}

func (p *ExceptionDateTime) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *ExceptionDateTime) encode(w writer) {
	w.WriteString("EXDATE")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *ExceptionDateTime) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type FreeBusy struct {
	FreeBusyType *FreeBusyType
	Period
}

func (p *FreeBusy) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "FBTYPE":
			if p.FreeBusyType != nil {
				return ErrDuplicateParam
			}
			p.FreeBusyType = new(FreeBusyType)
			if err := p.FreeBusyType.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Period.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *FreeBusy) encode(w writer) {
	w.WriteString("FREEBUSY")
	if p.FreeBusyType != nil {
		p.FreeBusyType.encode(w)
	}
	p.Period.aencode(w)
	w.WriteString("\r\n")
}

func (p *FreeBusy) valid() error {
	if p.FreeBusyType != nil {
		if err := p.FreeBusyType.valid(); err != nil {
			return err
		}
	}
	if err := p.Period.valid(); err != nil {
		return err
	}
	return nil
}

type Geo TFloat

func (p *Geo) decode(params []parser.Token, value string) error {
	var t TFloat
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Geo(t)
	return nil
}

func (p *Geo) encode(w writer) {
	w.WriteString("GEO")
	t := TFloat(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Geo) valid() error {
	t := TFloat(*p)
	return t.valid()
}

type LastModified DateTime

func (p *LastModified) decode(params []parser.Token, value string) error {
	var t DateTime
	if err := t.decode(value); err != nil {
		return err
	}
	*p = LastModified(t)
	return nil
}

func (p *LastModified) encode(w writer) {
	w.WriteString("LAST-MODIFIED")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *LastModified) valid() error {
	t := DateTime(*p)
	return t.valid()
}

type Location struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *Location) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "ALTREP":
			if p.AlternativeRepresentation != nil {
				return ErrDuplicateParam
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return err
			}
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Text.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Location) encode(w writer) {
	w.WriteString("LOCATION")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *Location) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return err
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Text.valid(); err != nil {
		return err
	}
	return nil
}

type Method Text

func (p *Method) decode(params []parser.Token, value string) error {
	var t Text
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Method(t)
	return nil
}

func (p *Method) encode(w writer) {
	w.WriteString("METHOD")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Method) valid() error {
	t := Text(*p)
	return t.valid()
}

type Organizer struct {
	CommonName     *CommonName
	DirectoryEntry *DirectoryEntry
	SentBy         *SentBy
	Language       *Language
	CalAddress
}

func (p *Organizer) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "CN":
			if p.CommonName != nil {
				return ErrDuplicateParam
			}
			p.CommonName = new(CommonName)
			if err := p.CommonName.decode(pValues); err != nil {
				return err
			}
		case "DIR":
			if p.DirectoryEntry != nil {
				return ErrDuplicateParam
			}
			p.DirectoryEntry = new(DirectoryEntry)
			if err := p.DirectoryEntry.decode(pValues); err != nil {
				return err
			}
		case "SENT-BY":
			if p.SentBy != nil {
				return ErrDuplicateParam
			}
			p.SentBy = new(SentBy)
			if err := p.SentBy.decode(pValues); err != nil {
				return err
			}
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Caladdress.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Organizer) encode(w writer) {
	w.WriteString("ORGANIZER")
	if p.CommonName != nil {
		p.CommonName.encode(w)
	}
	if p.DirectoryEntry != nil {
		p.DirectoryEntry.encode(w)
	}
	if p.SentBy != nil {
		p.SentBy.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Caladdress.aencode(w)
	w.WriteString("\r\n")
}

func (p *Organizer) valid() error {
	if p.CommonName != nil {
		if err := p.CommonName.valid(); err != nil {
			return err
		}
	}
	if p.DirectoryEntry != nil {
		if err := p.DirectoryEntry.valid(); err != nil {
			return err
		}
	}
	if p.SentBy != nil {
		if err := p.SentBy.valid(); err != nil {
			return err
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Caladdress.valid(); err != nil {
		return err
	}
	return nil
}

type PercentComplete Integer

func (p *PercentComplete) decode(params []parser.Token, value string) error {
	var t Integer
	if err := t.decode(value); err != nil {
		return err
	}
	*p = PercentComplete(t)
	return nil
}

func (p *PercentComplete) encode(w writer) {
	w.WriteString("PERCENT-COMPLETE")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PercentComplete) valid() error {
	t := Integer(*p)
	return t.valid()
}

type Priority Integer

func (p *Priority) decode(params []parser.Token, value string) error {
	var t Integer
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Priority(t)
	return nil
}

func (p *Priority) encode(w writer) {
	w.WriteString("PRIORITY")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Priority) valid() error {
	t := Integer(*p)
	return t.valid()
}

type ProdID Text

func (p *ProdID) decode(params []parser.Token, value string) error {
	var t Text
	if err := t.decode(value); err != nil {
		return err
	}
	*p = ProdID(t)
	return nil
}

func (p *ProdID) encode(w writer) {
	w.WriteString("PRODID")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *ProdID) valid() error {
	t := Text(*p)
	return t.valid()
}

type RecurrenceDateTimes struct {
	DateTime *DateTime
	Date     *Date
	Period   *Period
}

func (p *RecurrenceDateTimes) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			case "PERIOD":
				vType = 2
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return err
		}
	case 2:
		p.Period = new(Period)
		if err := p.Period.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *RecurrenceDateTimes) encode(w writer) {
	w.WriteString("RDATE")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		p.Date.aencode(w)
	}
	if p.Period != nil {
		p.Period.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *RecurrenceDateTimes) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Period != nil {
		if err := p.Period.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type RecurrenceId struct {
	Range    *Range
	DateTime *DateTime
	Date     *Date
}

func (p *RecurrenceId) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "RANGE":
			if p.Range != nil {
				return ErrDuplicateParam
			}
			p.Range = new(Range)
			if err := p.Range.decode(pValues); err != nil {
				return err
			}
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *RecurrenceId) encode(w writer) {
	w.WriteString("RECURRENCE-ID")
	if p.Range != nil {
		p.Range.encode(w)
	}
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *RecurrenceId) valid() error {
	if p.Range != nil {
		if err := p.Range.valid(); err != nil {
			return err
		}
	}
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return err
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type RelatedTo struct {
	RelationshipType *RelationshipType
	Text
}

func (p *RelatedTo) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "RELTYPE":
			if p.RelationshipType != nil {
				return ErrDuplicateParam
			}
			p.RelationshipType = new(RelationshipType)
			if err := p.RelationshipType.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Text.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *RelatedTo) encode(w writer) {
	w.WriteString("RELATED-TO")
	if p.RelationshipType != nil {
		p.RelationshipType.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *RelatedTo) valid() error {
	if p.RelationshipType != nil {
		if err := p.RelationshipType.valid(); err != nil {
			return err
		}
	}
	if err := p.Text.valid(); err != nil {
		return err
	}
	return nil
}

type Repeat Integer

func (p *Repeat) decode(params []parser.Token, value string) error {
	var t Integer
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Repeat(t)
	return nil
}

func (p *Repeat) encode(w writer) {
	w.WriteString("REPEAT")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Repeat) valid() error {
	t := Integer(*p)
	return t.valid()
}

type RequestStatus Text

func (p *RequestStatus) decode(params []parser.Token, value string) error {
	var t Text
	if err := t.decode(value); err != nil {
		return err
	}
	*p = RequestStatus(t)
	return nil
}

func (p *RequestStatus) encode(w writer) {
	w.WriteString("REQUEST-STATUS")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *RequestStatus) valid() error {
	t := Text(*p)
	return t.valid()
}

type Resources struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	MText
}

func (p *Resources) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "ALTREP":
			if p.AlternativeRepresentation != nil {
				return ErrDuplicateParam
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return err
			}
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Mtext.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Resources) encode(w writer) {
	w.WriteString("RESOURCES")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Mtext.aencode(w)
	w.WriteString("\r\n")
}

func (p *Resources) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return err
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Mtext.valid(); err != nil {
		return err
	}
	return nil
}

type RecurrenceRule Recur

func (p *RecurrenceRule) decode(params []parser.Token, value string) error {
	var t Recur
	if err := t.decode(value); err != nil {
		return err
	}
	*p = RecurrenceRule(t)
	return nil
}

func (p *RecurrenceRule) encode(w writer) {
	w.WriteString("RRULE")
	t := Recur(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *RecurrenceRule) valid() error {
	t := Recur(*p)
	return t.valid()
}

type Sequence Integer

func (p *Sequence) decode(params []parser.Token, value string) error {
	var t Integer
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Sequence(t)
	return nil
}

func (p *Sequence) encode(w writer) {
	w.WriteString("SEQUENCE")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Sequence) valid() error {
	t := Integer(*p)
	return t.valid()
}

type Status uint8

const (
	StatusTentative Status = iota
	StatusConfirmed
	StatusCancelled
	StatusNeedsAction
	StatusCompleted
	StatusInProcess
	StatusDraft
	StatusFinal
)

func (p *Status) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "TENTATIVE":
		*p = StatusTentative
	case "CONFIRMED":
		*p = StatusConfirmed
	case "CANCELLED":
		*p = StatusCancelled
	case "NEEDS-ACTION":
		*p = StatusNeedsAction
	case "COMPLETED":
		*p = StatusCompleted
	case "IN-PROCESS":
		*p = StatusInProcess
	case "DRAFT":
		*p = StatusDraft
	case "FINAL":
		*p = StatusFinal
	default:
		return ErrInvalidValue
	}
	return nil
}

func (p *Status) encode(w writer) {
	w.WriteString("STATUS:")
	switch *p {
	case Tentative:
		w.WriteString("TENTATIVE")
	case Confirmed:
		w.WriteString("CONFIRMED")
	case Cancelled:
		w.WriteString("CANCELLED")
	case NeedsAction:
		w.WriteString("NEEDS-ACTION")
	case Completed:
		w.WriteString("COMPLETED")
	case InProcess:
		w.WriteString("IN-PROCESS")
	case Draft:
		w.WriteString("DRAFT")
	case Final:
		w.WriteString("FINAL")
	}
	w.WriteString("\r\n")
}

func (p *Status) valid() error {
	switch *p {
	case Tentative, Confirmed, Cancelled, NeedsAction, Completed, InProcess, Draft, Final:
	default:
		return ErrInvalidValue
	}
	return nil
}

type Summary struct {
	AlternativeRepresentation *AlternativeRepresentation
	Langauge                  *Langauge
	Text
}

func (p *Summary) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "ALTREP":
			if p.AlternativeRepresentation != nil {
				return ErrDuplicateParam
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return err
			}
		case "LANGAUGE":
			if p.Langauge != nil {
				return ErrDuplicateParam
			}
			p.Langauge = new(Langauge)
			if err := p.Langauge.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Text.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *Summary) encode(w writer) {
	w.WriteString("SUMMARY")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Langauge != nil {
		p.Langauge.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *Summary) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return err
		}
	}
	if p.Langauge != nil {
		if err := p.Langauge.valid(); err != nil {
			return err
		}
	}
	if err := p.Text.valid(); err != nil {
		return err
	}
	return nil
}

type TimeTransparency uint8

const (
	TimeTransparencyOpaque TimeTransparency = iota
	TimeTransparencyTransparent
)

func (p *TimeTransparency) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "OPAQUE":
		*p = TimeTransparencyOpaque
	case "TRANSPARENT":
		*p = TimeTransparencyTransparent
	default:
		return ErrInvalidValue
	}
	return nil
}

func (p *TimeTransparency) encode(w writer) {
	w.WriteString("TRANSP:")
	switch *p {
	case Opaque:
		w.WriteString("OPAQUE")
	case Transparent:
		w.WriteString("TRANSPARENT")
	}
	w.WriteString("\r\n")
}

func (p *TimeTransparency) valid() error {
	switch *p {
	case Opaque, Transparent:
	default:
		return ErrInvalidValue
	}
	return nil
}

type Trigger struct {
	Duration *Duration
	DateTime *DateTime
}

func (p *Trigger) decode(params []parser.Token, value string) error {
	vType := -1
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "VALUE":
			if len(pValues) != 1 {
				return ErrInvalidValue
			}
			if vType != -1 {
				return ErrDuplicateParam
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DURATION":
				vType = 0
			case "DATE-TIME":
				vType = 1
			default:
				return ErrInvalidValue
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if vType == -1 {
		vType = 0
	}
	switch vType {
	case 0:
		p.Duration = new(Duration)
		if err := p.Duration.decode(oParams, value); err != nil {
			return err
		}
	case 1:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return err
		}
	}
	return nil
}

func (p *Trigger) encode(w writer) {
	w.WriteString("TRIGGER")
	if p.Duration != nil {
		p.Duration.aencode(w)
	}
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *Trigger) valid() error {
	c := 0
	if p.Duration != nil {
		if err := p.Duration.valid(); err != nil {
			return err
		}
		c++
	}
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return err
		}
		c++
	}
	if c != 1 {
		return ErrInvalidValue
	}
	return nil
}

type TimezoneID Text

func (p *TimezoneID) decode(params []parser.Token, value string) error {
	var t Text
	if err := t.decode(value); err != nil {
		return err
	}
	*p = TimezoneID(t)
	return nil
}

func (p *TimezoneID) encode(w writer) {
	w.WriteString("TZID")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *TimezoneID) valid() error {
	t := Text(*p)
	return t.valid()
}

type TimezoneName struct {
	Language *Language
	Text
}

func (p *TimezoneName) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		pName := strings.ToUpper(params[0].Data)
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		params = params[i:]
		switch pName {
		case "LANGUAGE":
			if p.Language != nil {
				return ErrDuplicateParam
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return err
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.Text.decode(oParams, value); err != nil {
		return err
	}
	return nil
}

func (p *TimezoneName) encode(w writer) {
	w.WriteString("TZNAME")
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *TimezoneName) valid() error {
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return err
		}
	}
	if err := p.Text.valid(); err != nil {
		return err
	}
	return nil
}

type TimezoneOffsetFrom UTCOffset

func (p *TimezoneOffsetFrom) decode(params []parser.Token, value string) error {
	var t UTCOffset
	if err := t.decode(value); err != nil {
		return err
	}
	*p = TimezoneOffsetFrom(t)
	return nil
}

func (p *TimezoneOffsetFrom) encode(w writer) {
	w.WriteString("TZOFFSETFROM")
	t := UTCOffset(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *TimezoneOffsetFrom) valid() error {
	t := UTCOffset(*p)
	return t.valid()
}

type TimezoneOffsetTo UTCOffset

func (p *TimezoneOffsetTo) decode(params []parser.Token, value string) error {
	var t UTCOffset
	if err := t.decode(value); err != nil {
		return err
	}
	*p = TimezoneOffsetTo(t)
	return nil
}

func (p *TimezoneOffsetTo) encode(w writer) {
	w.WriteString("TZOFFSETTO")
	t := UTCOffset(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *TimezoneOffsetTo) valid() error {
	t := UTCOffset(*p)
	return t.valid()
}

type TimezoneURL URI

func (p *TimezoneURL) decode(params []parser.Token, value string) error {
	var t URI
	if err := t.decode(value); err != nil {
		return err
	}
	*p = TimezoneURL(t)
	return nil
}

func (p *TimezoneURL) encode(w writer) {
	w.WriteString("TZURL")
	t := URI(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *TimezoneURL) valid() error {
	t := URI(*p)
	return t.valid()
}

type UID Text

func (p *UID) decode(params []parser.Token, value string) error {
	var t Text
	if err := t.decode(value); err != nil {
		return err
	}
	*p = UID(t)
	return nil
}

func (p *UID) encode(w writer) {
	w.WriteString("UID")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *UID) valid() error {
	t := Text(*p)
	return t.valid()
}

type URL URI

func (p *URL) decode(params []parser.Token, value string) error {
	var t URI
	if err := t.decode(value); err != nil {
		return err
	}
	*p = URL(t)
	return nil
}

func (p *URL) encode(w writer) {
	w.WriteString("URL")
	t := URI(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *URL) valid() error {
	t := URI(*p)
	return t.valid()
}

type Version Text

func (p *Version) decode(params []parser.Token, value string) error {
	var t Text
	if err := t.decode(value); err != nil {
		return err
	}
	*p = Version(t)
	return nil
}

func (p *Version) encode(w writer) {
	w.WriteString("VERSION")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *Version) valid() error {
	t := Text(*p)
	return t.valid()
}

// Errors
var (
	ErrInvalidValue = errors.New("invalid value")
)
