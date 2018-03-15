package ics

// File automatically generated with ./genProperties.sh

import (
	"strings"

	"github.com/MJKWoolnough/errors"
	"github.com/MJKWoolnough/parser"
)

// PropAction defines the action to be invoked when an alarm is triggered
type PropAction uint8

// PropAction constant values
const (
	ActionAudio PropAction = iota
	ActionDisplay
	ActionEmail
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (p PropAction) New() *PropAction {
	return &p
}

func (p *PropAction) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "AUDIO":
		*p = ActionAudio
	case "DISPLAY":
		*p = ActionDisplay
	case "EMAIL":
		*p = ActionEmail
	default:
		return errors.WithContext("error decoding Action: ", ErrInvalidValue)
	}
	return nil
}

func (p *PropAction) encode(w writer) {
	w.WriteString("ACTION:")
	switch *p {
	case ActionAudio:
		w.WriteString("AUDIO")
	case ActionDisplay:
		w.WriteString("DISPLAY")
	case ActionEmail:
		w.WriteString("EMAIL")
	}
	w.WriteString("\r\n")
}

func (p *PropAction) valid() error {
	switch *p {
	case ActionAudio, ActionDisplay, ActionEmail:
	default:
		return errors.WithContext("error validating Action: ", ErrInvalidValue)
	}
	return nil
}

// PropAttachment provides the capability to associate a document object with a
// calendar component
type PropAttachment struct {
	FormatType *FormatType
	URI        *URI
	Binary     *Binary
}

func (p *PropAttachment) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Attachment->FormatType: ", ErrDuplicateParam)
			}
			p.FormatType = new(FormatType)
			if err := p.FormatType.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attachment->FormatType: ", err)
			}
		case "VALUE":
			if len(pValues) != 1 {
				return errors.WithContext("error decoding Attachment->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding Attachment->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "URI":
				vType = 0
			case "BINARY":
				vType = 1
			default:
				return errors.WithContext("error decoding Attachment: ", ErrInvalidValue)
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
		p.URI = new(URI)
		if err := p.URI.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding Attachment->URI: ", err)
		}
	case 1:
		p.Binary = new(Binary)
		if err := p.Binary.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding Attachment->Binary: ", err)
		}
	}
	return nil
}

func (p *PropAttachment) encode(w writer) {
	w.WriteString("ATTACH")
	if p.FormatType != nil {
		p.FormatType.encode(w)
	}
	if p.URI != nil {
		p.URI.aencode(w)
	}
	if p.Binary != nil {
		w.WriteString(";VALUE=BINARY")
		p.Binary.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropAttachment) valid() error {
	if p.FormatType != nil {
		if err := p.FormatType.valid(); err != nil {
			return errors.WithContext("error validating Attachment->FormatType: ", err)
		}
	}
	c := 0
	if p.URI != nil {
		if err := p.URI.valid(); err != nil {
			return errors.WithContext("error validating Attachment->URI: ", err)
		}
		c++
	}
	if p.Binary != nil {
		if err := p.Binary.valid(); err != nil {
			return errors.WithContext("error validating Attachment->Binary: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating Attachment: ", ErrInvalidValue)
	}
	return nil
}

// PropAttendee defines an "Attendee" within a calendar component
type PropAttendee struct {
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
	CalendarAddress
}

func (p *PropAttendee) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Attendee->CalendarUserType: ", ErrDuplicateParam)
			}
			p.CalendarUserType = new(CalendarUserType)
			if err := p.CalendarUserType.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->CalendarUserType: ", err)
			}
		case "MEMBER":
			if p.Member != nil {
				return errors.WithContext("error decoding Attendee->Member: ", ErrDuplicateParam)
			}
			if err := p.Member.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->Member: ", err)
			}
		case "ROLE":
			if p.ParticipationRole != nil {
				return errors.WithContext("error decoding Attendee->ParticipationRole: ", ErrDuplicateParam)
			}
			p.ParticipationRole = new(ParticipationRole)
			if err := p.ParticipationRole.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->ParticipationRole: ", err)
			}
		case "PARTSTAT":
			if p.ParticipationStatus != nil {
				return errors.WithContext("error decoding Attendee->ParticipationStatus: ", ErrDuplicateParam)
			}
			p.ParticipationStatus = new(ParticipationStatus)
			if err := p.ParticipationStatus.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->ParticipationStatus: ", err)
			}
		case "RSVP":
			if p.RSVP != nil {
				return errors.WithContext("error decoding Attendee->RSVP: ", ErrDuplicateParam)
			}
			p.RSVP = new(RSVP)
			if err := p.RSVP.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->RSVP: ", err)
			}
		case "DELEGATED-TO":
			if p.Delagatee != nil {
				return errors.WithContext("error decoding Attendee->Delagatee: ", ErrDuplicateParam)
			}
			if err := p.Delagatee.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->Delagatee: ", err)
			}
		case "DELEGATED-FROM":
			if p.Delegator != nil {
				return errors.WithContext("error decoding Attendee->Delegator: ", ErrDuplicateParam)
			}
			if err := p.Delegator.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->Delegator: ", err)
			}
		case "SENT-BY":
			if p.SentBy != nil {
				return errors.WithContext("error decoding Attendee->SentBy: ", ErrDuplicateParam)
			}
			p.SentBy = new(SentBy)
			if err := p.SentBy.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->SentBy: ", err)
			}
		case "CN":
			if p.CommonName != nil {
				return errors.WithContext("error decoding Attendee->CommonName: ", ErrDuplicateParam)
			}
			p.CommonName = new(CommonName)
			if err := p.CommonName.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->CommonName: ", err)
			}
		case "DIR":
			if p.DirectoryEntry != nil {
				return errors.WithContext("error decoding Attendee->DirectoryEntry: ", ErrDuplicateParam)
			}
			p.DirectoryEntry = new(DirectoryEntry)
			if err := p.DirectoryEntry.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->DirectoryEntry: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Attendee->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Attendee->Language: ", err)
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.CalendarAddress.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Attendee->CalendarAddress: ", err)
	}
	return nil
}

func (p *PropAttendee) encode(w writer) {
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
	p.CalendarAddress.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropAttendee) valid() error {
	if p.CalendarUserType != nil {
		if err := p.CalendarUserType.valid(); err != nil {
			return errors.WithContext("error validating Attendee->CalendarUserType: ", err)
		}
	}
	if p.Member != nil {
		if err := p.Member.valid(); err != nil {
			return errors.WithContext("error validating Attendee->Member: ", err)
		}
	}
	if p.ParticipationRole != nil {
		if err := p.ParticipationRole.valid(); err != nil {
			return errors.WithContext("error validating Attendee->ParticipationRole: ", err)
		}
	}
	if p.ParticipationStatus != nil {
		if err := p.ParticipationStatus.valid(); err != nil {
			return errors.WithContext("error validating Attendee->ParticipationStatus: ", err)
		}
	}
	if p.RSVP != nil {
		if err := p.RSVP.valid(); err != nil {
			return errors.WithContext("error validating Attendee->RSVP: ", err)
		}
	}
	if p.Delagatee != nil {
		if err := p.Delagatee.valid(); err != nil {
			return errors.WithContext("error validating Attendee->Delagatee: ", err)
		}
	}
	if p.Delegator != nil {
		if err := p.Delegator.valid(); err != nil {
			return errors.WithContext("error validating Attendee->Delegator: ", err)
		}
	}
	if p.SentBy != nil {
		if err := p.SentBy.valid(); err != nil {
			return errors.WithContext("error validating Attendee->SentBy: ", err)
		}
	}
	if p.CommonName != nil {
		if err := p.CommonName.valid(); err != nil {
			return errors.WithContext("error validating Attendee->CommonName: ", err)
		}
	}
	if p.DirectoryEntry != nil {
		if err := p.DirectoryEntry.valid(); err != nil {
			return errors.WithContext("error validating Attendee->DirectoryEntry: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Attendee->Language: ", err)
		}
	}
	if err := p.CalendarAddress.valid(); err != nil {
		return errors.WithContext("error validating Attendee->CalendarAddress: ", err)
	}
	return nil
}

// PropCalendarScale defines the calendar scale
type PropCalendarScale uint8

// PropCalendarScale constant values
const (
	CalendarScaleGregorian PropCalendarScale = iota
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (p PropCalendarScale) New() *PropCalendarScale {
	return &p
}

func (p *PropCalendarScale) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "GREGORIAN":
		*p = CalendarScaleGregorian
	default:
		return errors.WithContext("error decoding CalendarScale: ", ErrInvalidValue)
	}
	return nil
}

func (p *PropCalendarScale) encode(w writer) {
	w.WriteString("CALSCALE:")
	switch *p {
	case CalendarScaleGregorian:
		w.WriteString("GREGORIAN")
	}
	w.WriteString("\r\n")
}

func (p *PropCalendarScale) valid() error {
	switch *p {
	case CalendarScaleGregorian:
	default:
		return errors.WithContext("error validating CalendarScale: ", ErrInvalidValue)
	}
	return nil
}

// PropCategories defines the categories for a calendar component
type PropCategories struct {
	Language *Language
	MText
}

func (p *PropCategories) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Categories->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Categories->Language: ", err)
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.MText.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Categories->MText: ", err)
	}
	return nil
}

func (p *PropCategories) encode(w writer) {
	w.WriteString("CATEGORIES")
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.MText.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropCategories) valid() error {
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Categories->Language: ", err)
		}
	}
	if err := p.MText.valid(); err != nil {
		return errors.WithContext("error validating Categories->MText: ", err)
	}
	return nil
}

// PropClass defines the access classification for a calendar component
type PropClass uint8

// PropClass constant values
const (
	ClassPublic PropClass = iota
	ClassPrivate
	ClassConfidential
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (p PropClass) New() *PropClass {
	return &p
}

func (p *PropClass) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "PUBLIC":
		*p = ClassPublic
	case "PRIVATE":
		*p = ClassPrivate
	case "CONFIDENTIAL":
		*p = ClassConfidential
	default:
		return errors.WithContext("error decoding Class: ", ErrInvalidValue)
	}
	return nil
}

func (p *PropClass) encode(w writer) {
	w.WriteString("CLASS:")
	switch *p {
	case ClassPublic:
		w.WriteString("PUBLIC")
	case ClassPrivate:
		w.WriteString("PRIVATE")
	case ClassConfidential:
		w.WriteString("CONFIDENTIAL")
	}
	w.WriteString("\r\n")
}

func (p *PropClass) valid() error {
	switch *p {
	case ClassPublic, ClassPrivate, ClassConfidential:
	default:
		return errors.WithContext("error validating Class: ", ErrInvalidValue)
	}
	return nil
}

// PropComment specifies non-processing information intended to provide a
// comment to the calendar user
type PropComment struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *PropComment) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Comment->AlternativeRepresentation: ", ErrDuplicateParam)
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return errors.WithContext("error decoding Comment->AlternativeRepresentation: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Comment->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Comment->Language: ", err)
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
		return errors.WithContext("error decoding Comment->Text: ", err)
	}
	return nil
}

func (p *PropComment) encode(w writer) {
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

func (p *PropComment) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return errors.WithContext("error validating Comment->AlternativeRepresentation: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Comment->Language: ", err)
		}
	}
	if err := p.Text.valid(); err != nil {
		return errors.WithContext("error validating Comment->Text: ", err)
	}
	return nil
}

// PropCompleted defines the date and time that a to-do was actually completed
type PropCompleted DateTime

func (p *PropCompleted) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t DateTime
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Completed: ", err)
	}
	*p = PropCompleted(t)
	return nil
}

func (p *PropCompleted) encode(w writer) {
	w.WriteString("COMPLETED")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropCompleted) valid() error {
	t := DateTime(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Completed: ", err)
	}
	return nil
}

// PropContact is used to represent contact information or alternately a
// reference to contact information associated with the calendar component
type PropContact struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *PropContact) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Contact->AlternativeRepresentation: ", ErrDuplicateParam)
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return errors.WithContext("error decoding Contact->AlternativeRepresentation: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Contact->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Contact->Language: ", err)
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
		return errors.WithContext("error decoding Contact->Text: ", err)
	}
	return nil
}

func (p *PropContact) encode(w writer) {
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

func (p *PropContact) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return errors.WithContext("error validating Contact->AlternativeRepresentation: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Contact->Language: ", err)
		}
	}
	if err := p.Text.valid(); err != nil {
		return errors.WithContext("error validating Contact->Text: ", err)
	}
	return nil
}

// PropCreated specifies the date and time that the calendar information was
// created by the calendar user agent in the calendar store
type PropCreated DateTime

func (p *PropCreated) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t DateTime
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Created: ", err)
	}
	*p = PropCreated(t)
	return nil
}

func (p *PropCreated) encode(w writer) {
	w.WriteString("CREATED")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropCreated) valid() error {
	t := DateTime(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Created: ", err)
	}
	return nil
}

// PropDescription provides a more complete description of the calendar
// component than that provided by the "SUMMARY" property
type PropDescription struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *PropDescription) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Description->AlternativeRepresentation: ", ErrDuplicateParam)
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return errors.WithContext("error decoding Description->AlternativeRepresentation: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Description->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Description->Language: ", err)
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
		return errors.WithContext("error decoding Description->Text: ", err)
	}
	return nil
}

func (p *PropDescription) encode(w writer) {
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

func (p *PropDescription) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return errors.WithContext("error validating Description->AlternativeRepresentation: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Description->Language: ", err)
		}
	}
	if err := p.Text.valid(); err != nil {
		return errors.WithContext("error validating Description->Text: ", err)
	}
	return nil
}

// PropDateTimeEnd specifies the date and time that a calendar component ends
type PropDateTimeEnd struct {
	DateTime *DateTime
	Date     *Date
}

func (p *PropDateTimeEnd) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding DateTimeEnd->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding DateTimeEnd->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return errors.WithContext("error decoding DateTimeEnd: ", ErrInvalidValue)
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
			return errors.WithContext("error decoding DateTimeEnd->DateTime: ", err)
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding DateTimeEnd->Date: ", err)
		}
	}
	return nil
}

func (p *PropDateTimeEnd) encode(w writer) {
	w.WriteString("DTEND")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		w.WriteString(";VALUE=DATE")
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropDateTimeEnd) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return errors.WithContext("error validating DateTimeEnd->DateTime: ", err)
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return errors.WithContext("error validating DateTimeEnd->Date: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating DateTimeEnd: ", ErrInvalidValue)
	}
	return nil
}

// PropDateTimeStamp specifies the date and time that the calendar object was
// created unless the calendar object has no METHOD property, in which case it
// specifies the date and time that the information with the calendar was last
// revised
type PropDateTimeStamp DateTime

func (p *PropDateTimeStamp) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t DateTime
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding DateTimeStamp: ", err)
	}
	*p = PropDateTimeStamp(t)
	return nil
}

func (p *PropDateTimeStamp) encode(w writer) {
	w.WriteString("DTSTAMP")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropDateTimeStamp) valid() error {
	t := DateTime(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating DateTimeStamp: ", err)
	}
	return nil
}

// PropDateTimeStart specifies when the calendar component begins
type PropDateTimeStart struct {
	DateTime *DateTime
	Date     *Date
}

func (p *PropDateTimeStart) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding DateTimeStart->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding DateTimeStart->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return errors.WithContext("error decoding DateTimeStart: ", ErrInvalidValue)
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
			return errors.WithContext("error decoding DateTimeStart->DateTime: ", err)
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding DateTimeStart->Date: ", err)
		}
	}
	return nil
}

func (p *PropDateTimeStart) encode(w writer) {
	w.WriteString("DTSTART")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		w.WriteString(";VALUE=DATE")
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropDateTimeStart) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return errors.WithContext("error validating DateTimeStart->DateTime: ", err)
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return errors.WithContext("error validating DateTimeStart->Date: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating DateTimeStart: ", ErrInvalidValue)
	}
	return nil
}

// PropDue defines the date and time that a to-do is expected to be completed
type PropDue struct {
	DateTime *DateTime
	Date     *Date
}

func (p *PropDue) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Due->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding Due->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return errors.WithContext("error decoding Due: ", ErrInvalidValue)
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
			return errors.WithContext("error decoding Due->DateTime: ", err)
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding Due->Date: ", err)
		}
	}
	return nil
}

func (p *PropDue) encode(w writer) {
	w.WriteString("DUE")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		w.WriteString(";VALUE=DATE")
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropDue) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return errors.WithContext("error validating Due->DateTime: ", err)
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return errors.WithContext("error validating Due->Date: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating Due: ", ErrInvalidValue)
	}
	return nil
}

// PropDuration specifies a positive duration of time
type PropDuration Duration

func (p *PropDuration) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Duration
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Duration: ", err)
	}
	*p = PropDuration(t)
	return nil
}

func (p *PropDuration) encode(w writer) {
	w.WriteString("DURATION")
	t := Duration(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropDuration) valid() error {
	t := Duration(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Duration: ", err)
	}
	return nil
}

// PropExceptionDateTime defines the list of DATE-TIME exceptions for recurring
// events, to-dos, journal entries, or time zone definitions
type PropExceptionDateTime struct {
	DateTime *DateTime
	Date     *Date
}

func (p *PropExceptionDateTime) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding ExceptionDateTime->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding ExceptionDateTime->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return errors.WithContext("error decoding ExceptionDateTime: ", ErrInvalidValue)
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
			return errors.WithContext("error decoding ExceptionDateTime->DateTime: ", err)
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding ExceptionDateTime->Date: ", err)
		}
	}
	return nil
}

func (p *PropExceptionDateTime) encode(w writer) {
	w.WriteString("EXDATE")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		w.WriteString(";VALUE=DATE")
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropExceptionDateTime) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return errors.WithContext("error validating ExceptionDateTime->DateTime: ", err)
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return errors.WithContext("error validating ExceptionDateTime->Date: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating ExceptionDateTime: ", ErrInvalidValue)
	}
	return nil
}

// PropFreeBusy defines one or more free or busy time intervals
type PropFreeBusy struct {
	FreeBusyType *FreeBusyType
	Period
}

func (p *PropFreeBusy) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding FreeBusy->FreeBusyType: ", ErrDuplicateParam)
			}
			p.FreeBusyType = new(FreeBusyType)
			if err := p.FreeBusyType.decode(pValues); err != nil {
				return errors.WithContext("error decoding FreeBusy->FreeBusyType: ", err)
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
		return errors.WithContext("error decoding FreeBusy->Period: ", err)
	}
	return nil
}

func (p *PropFreeBusy) encode(w writer) {
	w.WriteString("FREEBUSY")
	if p.FreeBusyType != nil {
		p.FreeBusyType.encode(w)
	}
	p.Period.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropFreeBusy) valid() error {
	if p.FreeBusyType != nil {
		if err := p.FreeBusyType.valid(); err != nil {
			return errors.WithContext("error validating FreeBusy->FreeBusyType: ", err)
		}
	}
	if err := p.Period.valid(); err != nil {
		return errors.WithContext("error validating FreeBusy->Period: ", err)
	}
	return nil
}

// PropGeo specifies information related to the global position for the activity
// specified by a calendar component
type PropGeo TFloat

func (p *PropGeo) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t TFloat
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Geo: ", err)
	}
	*p = PropGeo(t)
	return nil
}

func (p *PropGeo) encode(w writer) {
	w.WriteString("GEO")
	t := TFloat(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropGeo) valid() error {
	t := TFloat(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Geo: ", err)
	}
	return nil
}

// PropLastModified specifies the date and time that the information associated
// with the calendar component was last revised in the calendar store
type PropLastModified DateTime

func (p *PropLastModified) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t DateTime
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding LastModified: ", err)
	}
	*p = PropLastModified(t)
	return nil
}

func (p *PropLastModified) encode(w writer) {
	w.WriteString("LAST-MODIFIED")
	t := DateTime(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropLastModified) valid() error {
	t := DateTime(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating LastModified: ", err)
	}
	return nil
}

// PropLocation defines the intended venue for the activity defined by a
// calendar component
type PropLocation struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *PropLocation) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Location->AlternativeRepresentation: ", ErrDuplicateParam)
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return errors.WithContext("error decoding Location->AlternativeRepresentation: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Location->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Location->Language: ", err)
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
		return errors.WithContext("error decoding Location->Text: ", err)
	}
	return nil
}

func (p *PropLocation) encode(w writer) {
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

func (p *PropLocation) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return errors.WithContext("error validating Location->AlternativeRepresentation: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Location->Language: ", err)
		}
	}
	if err := p.Text.valid(); err != nil {
		return errors.WithContext("error validating Location->Text: ", err)
	}
	return nil
}

// PropMethod defines the iCalendar object method associated with the calendar
// object
type PropMethod Text

func (p *PropMethod) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Text
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Method: ", err)
	}
	*p = PropMethod(t)
	return nil
}

func (p *PropMethod) encode(w writer) {
	w.WriteString("METHOD")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropMethod) valid() error {
	t := Text(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Method: ", err)
	}
	return nil
}

// PropOrganizer defines the organizer for a calendar component
type PropOrganizer struct {
	CommonName     *CommonName
	DirectoryEntry *DirectoryEntry
	SentBy         *SentBy
	Language       *Language
	CalendarAddress
}

func (p *PropOrganizer) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Organizer->CommonName: ", ErrDuplicateParam)
			}
			p.CommonName = new(CommonName)
			if err := p.CommonName.decode(pValues); err != nil {
				return errors.WithContext("error decoding Organizer->CommonName: ", err)
			}
		case "DIR":
			if p.DirectoryEntry != nil {
				return errors.WithContext("error decoding Organizer->DirectoryEntry: ", ErrDuplicateParam)
			}
			p.DirectoryEntry = new(DirectoryEntry)
			if err := p.DirectoryEntry.decode(pValues); err != nil {
				return errors.WithContext("error decoding Organizer->DirectoryEntry: ", err)
			}
		case "SENT-BY":
			if p.SentBy != nil {
				return errors.WithContext("error decoding Organizer->SentBy: ", ErrDuplicateParam)
			}
			p.SentBy = new(SentBy)
			if err := p.SentBy.decode(pValues); err != nil {
				return errors.WithContext("error decoding Organizer->SentBy: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Organizer->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Organizer->Language: ", err)
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.CalendarAddress.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Organizer->CalendarAddress: ", err)
	}
	return nil
}

func (p *PropOrganizer) encode(w writer) {
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
	p.CalendarAddress.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropOrganizer) valid() error {
	if p.CommonName != nil {
		if err := p.CommonName.valid(); err != nil {
			return errors.WithContext("error validating Organizer->CommonName: ", err)
		}
	}
	if p.DirectoryEntry != nil {
		if err := p.DirectoryEntry.valid(); err != nil {
			return errors.WithContext("error validating Organizer->DirectoryEntry: ", err)
		}
	}
	if p.SentBy != nil {
		if err := p.SentBy.valid(); err != nil {
			return errors.WithContext("error validating Organizer->SentBy: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Organizer->Language: ", err)
		}
	}
	if err := p.CalendarAddress.valid(); err != nil {
		return errors.WithContext("error validating Organizer->CalendarAddress: ", err)
	}
	return nil
}

// PropPercentComplete is used by an assignee or delegatee of a to-do to convey
// the percent completion of a to-do to the "Organizer"
type PropPercentComplete Integer

// NewPercentComplete generates a pointer to a constant value.
// Used when manually creating Calendar values
func NewPercentComplete(v PropPercentComplete) *PropPercentComplete {
	return &v
}

func (p *PropPercentComplete) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Integer
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding PercentComplete: ", err)
	}
	*p = PropPercentComplete(t)
	return nil
}

func (p *PropPercentComplete) encode(w writer) {
	w.WriteString("PERCENT-COMPLETE")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropPercentComplete) valid() error {
	t := Integer(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating PercentComplete: ", err)
	}
	return nil
}

// PropPriority defines the relative priority for a calendar component
type PropPriority Integer

// NewPriority generates a pointer to a constant value.
// Used when manually creating Calendar values
func NewPriority(v PropPriority) *PropPriority {
	return &v
}

func (p *PropPriority) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Integer
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Priority: ", err)
	}
	*p = PropPriority(t)
	return nil
}

func (p *PropPriority) encode(w writer) {
	w.WriteString("PRIORITY")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropPriority) valid() error {
	t := Integer(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Priority: ", err)
	}
	return nil
}

// PropProductID specifies the identifier for the product that created the
// iCalendar object
type PropProductID Text

func (p *PropProductID) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Text
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding ProductID: ", err)
	}
	*p = PropProductID(t)
	return nil
}

func (p *PropProductID) encode(w writer) {
	w.WriteString("PRODID")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropProductID) valid() error {
	t := Text(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating ProductID: ", err)
	}
	return nil
}

// PropRecurrenceDateTimes defines the list of DATE-TIME values for recurring
// events, to-dos, journal entries, or time zone definitions
type PropRecurrenceDateTimes struct {
	DateTime *DateTime
	Date     *Date
	Period   *Period
}

func (p *PropRecurrenceDateTimes) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding RecurrenceDateTimes->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding RecurrenceDateTimes->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			case "PERIOD":
				vType = 2
			default:
				return errors.WithContext("error decoding RecurrenceDateTimes: ", ErrInvalidValue)
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
			return errors.WithContext("error decoding RecurrenceDateTimes->DateTime: ", err)
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding RecurrenceDateTimes->Date: ", err)
		}
	case 2:
		p.Period = new(Period)
		if err := p.Period.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding RecurrenceDateTimes->Period: ", err)
		}
	}
	return nil
}

func (p *PropRecurrenceDateTimes) encode(w writer) {
	w.WriteString("RDATE")
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		w.WriteString(";VALUE=DATE")
		p.Date.aencode(w)
	}
	if p.Period != nil {
		w.WriteString(";VALUE=PERIOD")
		p.Period.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropRecurrenceDateTimes) valid() error {
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return errors.WithContext("error validating RecurrenceDateTimes->DateTime: ", err)
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return errors.WithContext("error validating RecurrenceDateTimes->Date: ", err)
		}
		c++
	}
	if p.Period != nil {
		if err := p.Period.valid(); err != nil {
			return errors.WithContext("error validating RecurrenceDateTimes->Period: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating RecurrenceDateTimes: ", ErrInvalidValue)
	}
	return nil
}

// PropRecurrenceID is used to identify a specific instance of a recurring
// Event, Todo or Journal
type PropRecurrenceID struct {
	Range    *Range
	DateTime *DateTime
	Date     *Date
}

func (p *PropRecurrenceID) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding RecurrenceID->Range: ", ErrDuplicateParam)
			}
			p.Range = new(Range)
			if err := p.Range.decode(pValues); err != nil {
				return errors.WithContext("error decoding RecurrenceID->Range: ", err)
			}
		case "VALUE":
			if len(pValues) != 1 {
				return errors.WithContext("error decoding RecurrenceID->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding RecurrenceID->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DATE-TIME":
				vType = 0
			case "DATE":
				vType = 1
			default:
				return errors.WithContext("error decoding RecurrenceID: ", ErrInvalidValue)
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
			return errors.WithContext("error decoding RecurrenceID->DateTime: ", err)
		}
	case 1:
		p.Date = new(Date)
		if err := p.Date.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding RecurrenceID->Date: ", err)
		}
	}
	return nil
}

func (p *PropRecurrenceID) encode(w writer) {
	w.WriteString("RECURRENCE-ID")
	if p.Range != nil {
		p.Range.encode(w)
	}
	if p.DateTime != nil {
		p.DateTime.aencode(w)
	}
	if p.Date != nil {
		w.WriteString(";VALUE=DATE")
		p.Date.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropRecurrenceID) valid() error {
	if p.Range != nil {
		if err := p.Range.valid(); err != nil {
			return errors.WithContext("error validating RecurrenceID->Range: ", err)
		}
	}
	c := 0
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return errors.WithContext("error validating RecurrenceID->DateTime: ", err)
		}
		c++
	}
	if p.Date != nil {
		if err := p.Date.valid(); err != nil {
			return errors.WithContext("error validating RecurrenceID->Date: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating RecurrenceID: ", ErrInvalidValue)
	}
	return nil
}

// PropRelatedTo is used to represent a relationship or reference between one
// calendar component and another
type PropRelatedTo struct {
	RelationshipType *RelationshipType
	Text
}

func (p *PropRelatedTo) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding RelatedTo->RelationshipType: ", ErrDuplicateParam)
			}
			p.RelationshipType = new(RelationshipType)
			if err := p.RelationshipType.decode(pValues); err != nil {
				return errors.WithContext("error decoding RelatedTo->RelationshipType: ", err)
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
		return errors.WithContext("error decoding RelatedTo->Text: ", err)
	}
	return nil
}

func (p *PropRelatedTo) encode(w writer) {
	w.WriteString("RELATED-TO")
	if p.RelationshipType != nil {
		p.RelationshipType.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropRelatedTo) valid() error {
	if p.RelationshipType != nil {
		if err := p.RelationshipType.valid(); err != nil {
			return errors.WithContext("error validating RelatedTo->RelationshipType: ", err)
		}
	}
	if err := p.Text.valid(); err != nil {
		return errors.WithContext("error validating RelatedTo->Text: ", err)
	}
	return nil
}

// PropRepeat defines the number of times the alarm should be repeated, after
// the initial trigger
type PropRepeat Integer

// NewRepeat generates a pointer to a constant value.
// Used when manually creating Calendar values
func NewRepeat(v PropRepeat) *PropRepeat {
	return &v
}

func (p *PropRepeat) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Integer
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Repeat: ", err)
	}
	*p = PropRepeat(t)
	return nil
}

func (p *PropRepeat) encode(w writer) {
	w.WriteString("REPEAT")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropRepeat) valid() error {
	t := Integer(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Repeat: ", err)
	}
	return nil
}

// PropRequestStatus defines the status code returned for a scheduling request
type PropRequestStatus Text

func (p *PropRequestStatus) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Text
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding RequestStatus: ", err)
	}
	*p = PropRequestStatus(t)
	return nil
}

func (p *PropRequestStatus) encode(w writer) {
	w.WriteString("REQUEST-STATUS")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropRequestStatus) valid() error {
	t := Text(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating RequestStatus: ", err)
	}
	return nil
}

// PropResources defines the equipment or resources anticipated for an activity
// specified by a calendar component
type PropResources struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	MText
}

func (p *PropResources) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Resources->AlternativeRepresentation: ", ErrDuplicateParam)
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return errors.WithContext("error decoding Resources->AlternativeRepresentation: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Resources->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Resources->Language: ", err)
			}
		default:
			for _, v := range pValues {
				ts = append(ts, v.Data)
			}
			oParams[pName] = strings.Join(ts, ",")
			ts = ts[:0]
		}
	}
	if err := p.MText.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Resources->MText: ", err)
	}
	return nil
}

func (p *PropResources) encode(w writer) {
	w.WriteString("RESOURCES")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.MText.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropResources) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return errors.WithContext("error validating Resources->AlternativeRepresentation: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Resources->Language: ", err)
		}
	}
	if err := p.MText.valid(); err != nil {
		return errors.WithContext("error validating Resources->MText: ", err)
	}
	return nil
}

// PropRecurrenceRule defines a rule or repeating pattern for recurring events,
// to-dos, journal entries, or time zone definitions
type PropRecurrenceRule Recur

func (p *PropRecurrenceRule) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Recur
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding RecurrenceRule: ", err)
	}
	*p = PropRecurrenceRule(t)
	return nil
}

func (p *PropRecurrenceRule) encode(w writer) {
	w.WriteString("RRULE")
	t := Recur(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropRecurrenceRule) valid() error {
	t := Recur(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating RecurrenceRule: ", err)
	}
	return nil
}

// PropSequence defines the revision sequence number of the calendar component
// within a sequence of revisions
type PropSequence Integer

// NewSequence generates a pointer to a constant value.
// Used when manually creating Calendar values
func NewSequence(v PropSequence) *PropSequence {
	return &v
}

func (p *PropSequence) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Integer
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Sequence: ", err)
	}
	*p = PropSequence(t)
	return nil
}

func (p *PropSequence) encode(w writer) {
	w.WriteString("SEQUENCE")
	t := Integer(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropSequence) valid() error {
	t := Integer(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Sequence: ", err)
	}
	return nil
}

// PropStatus defines the overall status or confirmation for the calendar
// component
type PropStatus uint8

// PropStatus constant values
const (
	StatusTentative PropStatus = iota
	StatusConfirmed
	StatusCancelled
	StatusNeedsAction
	StatusCompleted
	StatusInProcess
	StatusDraft
	StatusFinal
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (p PropStatus) New() *PropStatus {
	return &p
}

func (p *PropStatus) decode(params []parser.Token, value string) error {
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
		return errors.WithContext("error decoding Status: ", ErrInvalidValue)
	}
	return nil
}

func (p *PropStatus) encode(w writer) {
	w.WriteString("STATUS:")
	switch *p {
	case StatusTentative:
		w.WriteString("TENTATIVE")
	case StatusConfirmed:
		w.WriteString("CONFIRMED")
	case StatusCancelled:
		w.WriteString("CANCELLED")
	case StatusNeedsAction:
		w.WriteString("NEEDS-ACTION")
	case StatusCompleted:
		w.WriteString("COMPLETED")
	case StatusInProcess:
		w.WriteString("IN-PROCESS")
	case StatusDraft:
		w.WriteString("DRAFT")
	case StatusFinal:
		w.WriteString("FINAL")
	}
	w.WriteString("\r\n")
}

func (p *PropStatus) valid() error {
	switch *p {
	case StatusTentative, StatusConfirmed, StatusCancelled, StatusNeedsAction, StatusCompleted, StatusInProcess, StatusDraft, StatusFinal:
	default:
		return errors.WithContext("error validating Status: ", ErrInvalidValue)
	}
	return nil
}

// PropSummary defines a short summary or subject for the calendar component
type PropSummary struct {
	AlternativeRepresentation *AlternativeRepresentation
	Language                  *Language
	Text
}

func (p *PropSummary) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Summary->AlternativeRepresentation: ", ErrDuplicateParam)
			}
			p.AlternativeRepresentation = new(AlternativeRepresentation)
			if err := p.AlternativeRepresentation.decode(pValues); err != nil {
				return errors.WithContext("error decoding Summary->AlternativeRepresentation: ", err)
			}
		case "LANGUAGE":
			if p.Language != nil {
				return errors.WithContext("error decoding Summary->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding Summary->Language: ", err)
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
		return errors.WithContext("error decoding Summary->Text: ", err)
	}
	return nil
}

func (p *PropSummary) encode(w writer) {
	w.WriteString("SUMMARY")
	if p.AlternativeRepresentation != nil {
		p.AlternativeRepresentation.encode(w)
	}
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropSummary) valid() error {
	if p.AlternativeRepresentation != nil {
		if err := p.AlternativeRepresentation.valid(); err != nil {
			return errors.WithContext("error validating Summary->AlternativeRepresentation: ", err)
		}
	}
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating Summary->Language: ", err)
		}
	}
	if err := p.Text.valid(); err != nil {
		return errors.WithContext("error validating Summary->Text: ", err)
	}
	return nil
}

// PropTimeTransparency defines whether or not an event is transparent to busy
// time searches
type PropTimeTransparency uint8

// PropTimeTransparency constant values
const (
	TimeTransparencyOpaque PropTimeTransparency = iota
	TimeTransparencyTransparent
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (p PropTimeTransparency) New() *PropTimeTransparency {
	return &p
}

func (p *PropTimeTransparency) decode(params []parser.Token, value string) error {
	switch strings.ToUpper(value) {
	case "OPAQUE":
		*p = TimeTransparencyOpaque
	case "TRANSPARENT":
		*p = TimeTransparencyTransparent
	default:
		return errors.WithContext("error decoding TimeTransparency: ", ErrInvalidValue)
	}
	return nil
}

func (p *PropTimeTransparency) encode(w writer) {
	w.WriteString("TRANSP:")
	switch *p {
	case TimeTransparencyOpaque:
		w.WriteString("OPAQUE")
	case TimeTransparencyTransparent:
		w.WriteString("TRANSPARENT")
	}
	w.WriteString("\r\n")
}

func (p *PropTimeTransparency) valid() error {
	switch *p {
	case TimeTransparencyOpaque, TimeTransparencyTransparent:
	default:
		return errors.WithContext("error validating TimeTransparency: ", ErrInvalidValue)
	}
	return nil
}

// PropTrigger specifies when an alarm will trigger
type PropTrigger struct {
	Duration *Duration
	DateTime *DateTime
}

func (p *PropTrigger) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding Trigger->Value: ", ErrInvalidValue)
			}
			if vType != -1 {
				return errors.WithContext("error decoding Trigger->Value: ", ErrDuplicateParam)
			}
			switch strings.ToUpper(pValues[0].Data) {
			case "DURATION":
				vType = 0
			case "DATE-TIME":
				vType = 1
			default:
				return errors.WithContext("error decoding Trigger: ", ErrInvalidValue)
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
			return errors.WithContext("error decoding Trigger->Duration: ", err)
		}
	case 1:
		p.DateTime = new(DateTime)
		if err := p.DateTime.decode(oParams, value); err != nil {
			return errors.WithContext("error decoding Trigger->DateTime: ", err)
		}
	}
	return nil
}

func (p *PropTrigger) encode(w writer) {
	w.WriteString("TRIGGER")
	if p.Duration != nil {
		p.Duration.aencode(w)
	}
	if p.DateTime != nil {
		w.WriteString(";VALUE=DATE-TIME")
		p.DateTime.aencode(w)
	}
	w.WriteString("\r\n")
}

func (p *PropTrigger) valid() error {
	c := 0
	if p.Duration != nil {
		if err := p.Duration.valid(); err != nil {
			return errors.WithContext("error validating Trigger->Duration: ", err)
		}
		c++
	}
	if p.DateTime != nil {
		if err := p.DateTime.valid(); err != nil {
			return errors.WithContext("error validating Trigger->DateTime: ", err)
		}
		c++
	}
	if c != 1 {
		return errors.WithContext("error validating Trigger: ", ErrInvalidValue)
	}
	return nil
}

// PropTimezoneID specifies the text value that uniquely identifies the
// "VTIMEZONE" calendar component in the scope of an iCalendar object
type PropTimezoneID Text

func (p *PropTimezoneID) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Text
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding TimezoneID: ", err)
	}
	*p = PropTimezoneID(t)
	return nil
}

func (p *PropTimezoneID) encode(w writer) {
	w.WriteString("TZID")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropTimezoneID) valid() error {
	t := Text(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating TimezoneID: ", err)
	}
	return nil
}

// PropTimezoneName specifies the customary designation for a time zone
// description
type PropTimezoneName struct {
	Language *Language
	Text
}

func (p *PropTimezoneName) decode(params []parser.Token, value string) error {
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
				return errors.WithContext("error decoding TimezoneName->Language: ", ErrDuplicateParam)
			}
			p.Language = new(Language)
			if err := p.Language.decode(pValues); err != nil {
				return errors.WithContext("error decoding TimezoneName->Language: ", err)
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
		return errors.WithContext("error decoding TimezoneName->Text: ", err)
	}
	return nil
}

func (p *PropTimezoneName) encode(w writer) {
	w.WriteString("TZNAME")
	if p.Language != nil {
		p.Language.encode(w)
	}
	p.Text.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropTimezoneName) valid() error {
	if p.Language != nil {
		if err := p.Language.valid(); err != nil {
			return errors.WithContext("error validating TimezoneName->Language: ", err)
		}
	}
	if err := p.Text.valid(); err != nil {
		return errors.WithContext("error validating TimezoneName->Text: ", err)
	}
	return nil
}

// PropTimezoneOffsetFrom specifies the offset that is in use prior to this time
// zone observance
type PropTimezoneOffsetFrom UTCOffset

func (p *PropTimezoneOffsetFrom) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t UTCOffset
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding TimezoneOffsetFrom: ", err)
	}
	*p = PropTimezoneOffsetFrom(t)
	return nil
}

func (p *PropTimezoneOffsetFrom) encode(w writer) {
	w.WriteString("TZOFFSETFROM")
	t := UTCOffset(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropTimezoneOffsetFrom) valid() error {
	t := UTCOffset(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating TimezoneOffsetFrom: ", err)
	}
	return nil
}

// PropTimezoneOffsetTo specifies the offset that is in use in this time zone
// observance
type PropTimezoneOffsetTo UTCOffset

func (p *PropTimezoneOffsetTo) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t UTCOffset
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding TimezoneOffsetTo: ", err)
	}
	*p = PropTimezoneOffsetTo(t)
	return nil
}

func (p *PropTimezoneOffsetTo) encode(w writer) {
	w.WriteString("TZOFFSETTO")
	t := UTCOffset(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropTimezoneOffsetTo) valid() error {
	t := UTCOffset(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating TimezoneOffsetTo: ", err)
	}
	return nil
}

// PropTimezoneURL provides a means for a "VTIMEZONE" component to point to a
// network location that can be used to retrieve an up- to-date version of
// itself
type PropTimezoneURL URI

func (p *PropTimezoneURL) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t URI
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding TimezoneURL: ", err)
	}
	*p = PropTimezoneURL(t)
	return nil
}

func (p *PropTimezoneURL) encode(w writer) {
	w.WriteString("TZURL")
	t := URI(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropTimezoneURL) valid() error {
	t := URI(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating TimezoneURL: ", err)
	}
	return nil
}

// PropUID defines the persistent, globally unique identifier for the calendar
// component
type PropUID Text

func (p *PropUID) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Text
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding UID: ", err)
	}
	*p = PropUID(t)
	return nil
}

func (p *PropUID) encode(w writer) {
	w.WriteString("UID")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropUID) valid() error {
	t := Text(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating UID: ", err)
	}
	return nil
}

// PropURL defines a Uniform Resource Locator associated with the iCalendar
// object
type PropURL URI

func (p *PropURL) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t URI
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding URL: ", err)
	}
	*p = PropURL(t)
	return nil
}

func (p *PropURL) encode(w writer) {
	w.WriteString("URL")
	t := URI(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropURL) valid() error {
	t := URI(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating URL: ", err)
	}
	return nil
}

// PropVersion specifies the identifier corresponding to the highest version
// number or the minimum and maximum range of the iCalendar specification that
// is required in order to interpret the iCalendar object
type PropVersion Text

func (p *PropVersion) decode(params []parser.Token, value string) error {
	oParams := make(map[string]string)
	var ts []string
	for len(params) > 0 {
		i := 1
		for i < len(params) && params[i].Type != tokenParamName {
			i++
		}
		pValues := params[1:i]
		for _, v := range pValues {
			ts = append(ts, v.Data)
		}
		oParams[strings.ToUpper(params[0].Data)] = strings.Join(ts, ",")
		params = params[i:]
		ts = ts[:0]
	}
	var t Text
	if err := t.decode(oParams, value); err != nil {
		return errors.WithContext("error decoding Version: ", err)
	}
	*p = PropVersion(t)
	return nil
}

func (p *PropVersion) encode(w writer) {
	w.WriteString("VERSION")
	t := Text(*p)
	t.aencode(w)
	w.WriteString("\r\n")
}

func (p *PropVersion) valid() error {
	t := Text(*p)
	if err := t.valid(); err != nil {
		return errors.WithContext("error validating Version: ", err)
	}
	return nil
}

// Errors
const (
	ErrDuplicateParam errors.Error = "duplicate param"
)
