package ics

// File automatically generated with ./genParams.sh

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/MJKWoolnough/parser"
)

type AlternativeRepresentation URI

func (t *AlternativeRepresentation) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	if vs[0].Type != tokenParamQuotedValue {
		return ErrInvalidParam
	}
	var q URI
	if err := q.decode(nil, vs[0].Data); err != nil {
		return err
	}
	*t = AlternativeRepresentation(q)
	return nil
}

func (t AlternativeRepresentation) encode(w writer) {
	if len(t.String()) == 0 {
		return
	}
	w.WriteString(";ALTREP=")
	q := URI(t)
	q.encode(w)
}

func (t AlternativeRepresentation) valid() error {
	q := URI(t)
	return q.valid()
}

type CommonName string

func (t *CommonName) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	*t = CommonName(decode6868(vs[0].Data))
	return nil
}

func (t CommonName) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";CN=")
	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.WriteString(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.WriteString(encode6868(string(t)))
	}
}

func (t CommonName) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return ErrInvalidText
	}
	return nil
}

type CalendarUserType uint8

const (
	CalendarUserTypeUnknown CalendarUserType = iota
	CalendarUserTypeIndividual
	CalendarUserTypeGroup
	CalendarUserTypeResource
	CalendarUserTypeRoom
)

func (t *CalendarUserType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "INDIVIDUAL":
		*t = CalendarUserTypeIndividual
	case "GROUP":
		*t = CalendarUserTypeGroup
	case "RESOURCE":
		*t = CalendarUserTypeResource
	case "ROOM":
		*t = CalendarUserTypeRoom
	default:
		*t = CalendarUserTypeUnknown
	}
	return nil
}

func (t CalendarUserType) encode(w writer) {
	w.WriteString(";CUTYPE=")
	switch t {
	case CalendarUserTypeIndividual:
		w.WriteString("INDIVIDUAL")
	case CalendarUserTypeGroup:
		w.WriteString("GROUP")
	case CalendarUserTypeResource:
		w.WriteString("RESOURCE")
	case CalendarUserTypeRoom:
		w.WriteString("ROOM")
	default:
		w.WriteString("UNKNOWN")
	}
}

func (t CalendarUserType) valid() error {
	return nil
}

type Delegator []CalendarAddress

func (t *Delegator) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return ErrInvalidParam
		}
		var q CalendarAddress
		if err := q.decode(nil, v.Data); err != nil {
			return err
		}
		*t = append(*t, q)
	}
	return nil
}

func (t Delegator) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";DELEGATED-FROM=")
	for n, v := range t {
		if n > 0 {
			w.WriteString(",")
		}
		q := CalendarAddress(v)
		q.encode(w)
	}
}

func (t Delegator) valid() error {
	for _, v := range t {
		if err := v.valid(); err != nil {
			return err
		}
	}
	return nil
}

type Delagatee []CalendarAddress

func (t *Delagatee) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return ErrInvalidParam
		}
		var q CalendarAddress
		if err := q.decode(nil, v.Data); err != nil {
			return err
		}
		*t = append(*t, q)
	}
	return nil
}

func (t Delagatee) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";DELEGATED-TO=")
	for n, v := range t {
		if n > 0 {
			w.WriteString(",")
		}
		q := CalendarAddress(v)
		q.encode(w)
	}
}

func (t Delagatee) valid() error {
	for _, v := range t {
		if err := v.valid(); err != nil {
			return err
		}
	}
	return nil
}

type DirectoryEntry string

func (t *DirectoryEntry) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	if vs[0].Type != tokenParamQuotedValue {
		return ErrInvalidParam
	}
	*t = DirectoryEntry(decode6868(vs[0].Data))
	return nil
}

func (t DirectoryEntry) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";DIR=")
	w.WriteString("\"")
	w.WriteString(encode6868(string(t)))
	w.WriteString("\"")
}

func (t DirectoryEntry) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return ErrInvalidText
	}
	return nil
}

type Encoding uint8

const (
	Encoding8bit Encoding = iota
	EncodingBase64
)

func (t *Encoding) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "8BIT":
		*t = Encoding8bit
	case "BASE64":
		*t = EncodingBase64
	default:
		return ErrInvalidParam
	}
	return nil
}

func (t Encoding) encode(w writer) {
	w.WriteString(";ENCODING=")
	switch t {
	case Encoding8bit:
		w.WriteString("8BIT")
	case EncodingBase64:
		w.WriteString("BASE64")
	}
}

func (t Encoding) valid() error {
	switch t {
	case Encoding8bit, EncodingBase64:
	default:
		return ErrInvalidValue
	}
	return nil
}

type FormatType string

var regexFormatType *regexp.Regexp

func (t *FormatType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	if !regexFormatType.MatchString(vs[0].Data) {
		return ErrInvalidParam
	}
	*t = FormatType(vs[0].Data)
	return nil
}

func (t FormatType) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";FMTTYPE=")
	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.WriteString(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.WriteString(encode6868(string(t)))
	}
}

func (t FormatType) valid() error {
	if !regexFormatType.Match([]byte(t)) {
		return ErrInvalidValue
	}
	return nil
}

type FreeBusyType uint8

const (
	FreeBusyTypeUnknown FreeBusyType = iota
	FreeBusyTypeFree
	FreeBusyTypeBusy
	FreeBusyTypeBusyUnavailable
	FreeBusyTypeBusyTentative
)

func (t *FreeBusyType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "FREE":
		*t = FreeBusyTypeFree
	case "BUSY":
		*t = FreeBusyTypeBusy
	case "BUSY-UNAVAILABLE":
		*t = FreeBusyTypeBusyUnavailable
	case "BUSY-TENTATIVE":
		*t = FreeBusyTypeBusyTentative
	default:
		*t = FreeBusyTypeUnknown
	}
	return nil
}

func (t FreeBusyType) encode(w writer) {
	w.WriteString(";FBTYPE=")
	switch t {
	case FreeBusyTypeFree:
		w.WriteString("FREE")
	case FreeBusyTypeBusy:
		w.WriteString("BUSY")
	case FreeBusyTypeBusyUnavailable:
		w.WriteString("BUSY-UNAVAILABLE")
	case FreeBusyTypeBusyTentative:
		w.WriteString("BUSY-TENTATIVE")
	default:
		w.WriteString("UNKNOWN")
	}
}

func (t FreeBusyType) valid() error {
	return nil
}

type Language string

func (t *Language) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	*t = Language(decode6868(vs[0].Data))
	return nil
}

func (t Language) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";LANGUAGE=")
	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.WriteString(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.WriteString(encode6868(string(t)))
	}
}

func (t Language) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return ErrInvalidText
	}
	return nil
}

type Member []string

func (t *Member) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return ErrInvalidParam
		}
		*t = append(*t, decode6868(v.Data))
	}
	return nil
}

func (t Member) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";MEMBER=")
	for n, v := range t {
		if n > 0 {
			w.WriteString(",")
		}
		w.WriteString("\"")
		w.WriteString(encode6868(string(v)))
		w.WriteString("\"")
	}
}

func (t Member) valid() error {
	for _, v := range t {
		if strings.ContainsAny(string(v), nonsafeChars[:31]) {
			return ErrInvalidText
		}
	}
	return nil
}

type ParticipationStatus uint8

const (
	ParticipationStatusUnknown ParticipationStatus = iota
	ParticipationStatusNeedsAction
	ParticipationStatusAccepted
	ParticipationStatusDeclined
	ParticipationStatusTentative
	ParticipationStatusDelegated
	ParticipationStatusCompleted
	ParticipationStatusInProcess
)

func (t *ParticipationStatus) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "NEEDS-ACTION":
		*t = ParticipationStatusNeedsAction
	case "ACCEPTED":
		*t = ParticipationStatusAccepted
	case "DECLINED":
		*t = ParticipationStatusDeclined
	case "TENTATIVE":
		*t = ParticipationStatusTentative
	case "DELEGATED":
		*t = ParticipationStatusDelegated
	case "COMPLETED":
		*t = ParticipationStatusCompleted
	case "IN-PROCESS":
		*t = ParticipationStatusInProcess
	default:
		*t = ParticipationStatusUnknown
	}
	return nil
}

func (t ParticipationStatus) encode(w writer) {
	w.WriteString(";PARTSTAT=")
	switch t {
	case ParticipationStatusNeedsAction:
		w.WriteString("NEEDS-ACTION")
	case ParticipationStatusAccepted:
		w.WriteString("ACCEPTED")
	case ParticipationStatusDeclined:
		w.WriteString("DECLINED")
	case ParticipationStatusTentative:
		w.WriteString("TENTATIVE")
	case ParticipationStatusDelegated:
		w.WriteString("DELEGATED")
	case ParticipationStatusCompleted:
		w.WriteString("COMPLETED")
	case ParticipationStatusInProcess:
		w.WriteString("IN-PROCESS")
	default:
		w.WriteString("UNKNOWN")
	}
}

func (t ParticipationStatus) valid() error {
	return nil
}

type Range struct{}

func (t *Range) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	if strings.ToUpper(vs[0].Data) != "THISANDFUTURE" {
		return ErrInvalidParam
	}
	return nil
}

func (t Range) encode(w writer) {
	w.WriteString(";RANGE=")
	w.WriteString("THISANDFUTURE")
}

func (t Range) valid() error {
	return nil
}

type Related uint8

const (
	RelatedStart Related = iota
	RelatedEnd
)

func (t *Related) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "START":
		*t = RelatedStart
	case "END":
		*t = RelatedEnd
	default:
		return ErrInvalidParam
	}
	return nil
}

func (t Related) encode(w writer) {
	w.WriteString(";RELATED=")
	switch t {
	case RelatedStart:
		w.WriteString("START")
	case RelatedEnd:
		w.WriteString("END")
	}
}

func (t Related) valid() error {
	switch t {
	case RelatedStart, RelatedEnd:
	default:
		return ErrInvalidValue
	}
	return nil
}

type RelationshipType uint8

const (
	RelationshipTypeUnknown RelationshipType = iota
	RelationshipTypeParent
	RelationshipTypeChild
	RelationshipTypeSibling
)

func (t *RelationshipType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "PARENT":
		*t = RelationshipTypeParent
	case "CHILD":
		*t = RelationshipTypeChild
	case "SIBLING":
		*t = RelationshipTypeSibling
	default:
		*t = RelationshipTypeUnknown
	}
	return nil
}

func (t RelationshipType) encode(w writer) {
	w.WriteString(";RELTYPE=")
	switch t {
	case RelationshipTypeParent:
		w.WriteString("PARENT")
	case RelationshipTypeChild:
		w.WriteString("CHILD")
	case RelationshipTypeSibling:
		w.WriteString("SIBLING")
	default:
		w.WriteString("UNKNOWN")
	}
}

func (t RelationshipType) valid() error {
	return nil
}

type ParticipationRole uint8

const (
	ParticipationRoleUnknown ParticipationRole = iota
	ParticipationRoleRequiredParticipant
	ParticipationRoleChair
	ParticipationRoleOptParticipant
	ParticipationRoleNonParticipant
)

func (t *ParticipationRole) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "REQ-PARTICIPANT":
		*t = ParticipationRoleRequiredParticipant
	case "CHAIR":
		*t = ParticipationRoleChair
	case "OPT-PARTICIPANT":
		*t = ParticipationRoleOptParticipant
	case "NON-PARTICIPANT":
		*t = ParticipationRoleNonParticipant
	default:
		*t = ParticipationRoleUnknown
	}
	return nil
}

func (t ParticipationRole) encode(w writer) {
	w.WriteString(";ROLE=")
	switch t {
	case ParticipationRoleRequiredParticipant:
		w.WriteString("REQ-PARTICIPANT")
	case ParticipationRoleChair:
		w.WriteString("CHAIR")
	case ParticipationRoleOptParticipant:
		w.WriteString("OPT-PARTICIPANT")
	case ParticipationRoleNonParticipant:
		w.WriteString("NON-PARTICIPANT")
	default:
		w.WriteString("UNKNOWN")
	}
}

func (t ParticipationRole) valid() error {
	return nil
}

type RSVP Boolean

func (t *RSVP) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	var q Boolean
	if err := q.decode(nil, vs[0].Data); err != nil {
		return err
	}
	*t = RSVP(q)
	return nil
}

func (t RSVP) encode(w writer) {
	if !t {
		return
	}
	w.WriteString(";RSVP=")
	q := Boolean(t)
	q.encode(w)
}

func (t RSVP) valid() error {
	return nil
}

type SentBy string

func (t *SentBy) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	if vs[0].Type != tokenParamQuotedValue {
		return ErrInvalidParam
	}
	*t = SentBy(decode6868(vs[0].Data))
	return nil
}

func (t SentBy) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";SENT-BY=")
	w.WriteString("\"")
	w.WriteString(encode6868(string(t)))
	w.WriteString("\"")
}

func (t SentBy) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return ErrInvalidText
	}
	return nil
}

type TimezoneID string

func (t *TimezoneID) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	*t = TimezoneID(decode6868(vs[0].Data))
	return nil
}

func (t TimezoneID) encode(w writer) {
	if len(t) == 0 {
		return
	}
	w.WriteString(";TZID=")
	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.WriteString(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.WriteString(encode6868(string(t)))
	}
}

func (t TimezoneID) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return ErrInvalidText
	}
	return nil
}

type Value uint8

const (
	ValueUnknown Value = iota
	ValueBinary
	ValueBoolean
	ValueCalendarAddress
	ValueDate
	ValueDateTime
	ValueDuration
	ValueFloat
	ValueInteger
	ValuePeriod
	ValueRecur
	ValueText
	ValueTime
	ValueURI
	ValueUTCOffset
)

func (t *Value) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return ErrInvalidParam
	}
	switch strings.ToUpper(vs[0].Data) {
	case "BINARY":
		*t = ValueBinary
	case "BOOLEAN":
		*t = ValueBoolean
	case "CAL-ADDRESS":
		*t = ValueCalendarAddress
	case "DATE":
		*t = ValueDate
	case "DATE-TIME":
		*t = ValueDateTime
	case "DURATION":
		*t = ValueDuration
	case "FLOAT":
		*t = ValueFloat
	case "INTEGER":
		*t = ValueInteger
	case "PERIOD":
		*t = ValuePeriod
	case "RECUR":
		*t = ValueRecur
	case "TEXT":
		*t = ValueText
	case "TIME":
		*t = ValueTime
	case "URI":
		*t = ValueURI
	case "UTC-OFFSET":
		*t = ValueUTCOffset
	default:
		*t = ValueUnknown
	}
	return nil
}

func (t Value) encode(w writer) {
	w.WriteString(";VALUE=")
	switch t {
	case ValueBinary:
		w.WriteString("BINARY")
	case ValueBoolean:
		w.WriteString("BOOLEAN")
	case ValueCalendarAddress:
		w.WriteString("CAL-ADDRESS")
	case ValueDate:
		w.WriteString("DATE")
	case ValueDateTime:
		w.WriteString("DATE-TIME")
	case ValueDuration:
		w.WriteString("DURATION")
	case ValueFloat:
		w.WriteString("FLOAT")
	case ValueInteger:
		w.WriteString("INTEGER")
	case ValuePeriod:
		w.WriteString("PERIOD")
	case ValueRecur:
		w.WriteString("RECUR")
	case ValueText:
		w.WriteString("TEXT")
	case ValueTime:
		w.WriteString("TIME")
	case ValueURI:
		w.WriteString("URI")
	case ValueUTCOffset:
		w.WriteString("UTC-OFFSET")
	default:
		w.WriteString("UNKNOWN")
	}
}

func (t Value) valid() error {
	return nil
}

func decode6868(s string) string {
	t := parser.NewStringTokeniser(s)
	d := make([]byte, 0, len(s))
	var ru [4]byte
Loop:
	for {
		c := t.ExceptRun("^")
		d = append(d, t.Get()...)
		switch c {
		case -1:
			break Loop
		case '^':
			t.Accept("^")
			switch t.Peek() {
			case -1:
				d = append(d, '^')
				break Loop
			case 'n':
				d = append(d, '\n')
			case '\'':
				d = append(d, '"')
			case '^':
				d = append(d, '^')
			default:
				d = append(d, '^')
				l := utf8.EncodeRune(ru[:], c)
				d = append(d, ru[:l]...)
			}
			t.Except("")
		}
	}
	return string(d)
}

func encode6868(s string) string {
	t := parser.NewStringTokeniser(s)
	d := make([]byte, 0, len(s))
Loop:
	for {
		c := t.ExceptRun("\n^\"")
		d = append(d, t.Get()...)
		switch c {
		case -1:
			break Loop
		case '\n':
			d = append(d, '^', 'n')
		case '^':
			d = append(d, '^', '^')
		case '"':
			d = append(d, '^', '\'')
		}
	}
	return string(d)
}

func init() {
	regexFormatType = regexp.MustCompile("[A-Z0-9!#$&.+-^_]/[A-Z0-9!#$&.+-^_]")
}

// Errors
var (
	ErrInvalidParam = errors.New("invalid param value")
	ErrInvalidValue = errors.New("invalid value")
)
