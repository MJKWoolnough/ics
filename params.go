package ics

// File automatically generated with ./genParams.sh

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"vimagination.zapto.org/parser"
)

// ParamAlternativeRepresentation.
type ParamAlternativeRepresentation URI

func (t *ParamAlternativeRepresentation) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cAlternativeRepresentation, ErrInvalidParam)
	}

	if vs[0].Type != tokenParamQuotedValue {
		return fmt.Errorf(errDecodingType, cAlternativeRepresentation, ErrInvalidParam)
	}

	var q URI

	if err := q.decode(nil, vs[0].Data); err != nil {
		return fmt.Errorf(errDecodingType, cAlternativeRepresentation, err)
	}

	*t = ParamAlternativeRepresentation(q)

	return nil
}

func (t ParamAlternativeRepresentation) encode(w writer) {
	if len(t.String()) == 0 {
		return
	}

	w.WriteString(";ALTREP=")

	q := URI(t)

	q.encode(w)
}

func (t ParamAlternativeRepresentation) valid() error {
	q := URI(t)

	if err := q.valid(); err != nil {
		return fmt.Errorf(errValidatingType, cAlternativeRepresentation, err)
	}

	return nil
}

// ParamCommonName.
type ParamCommonName string

// NewCommonName returns a *ParamCommonName for ease of use with optional values.
func NewCommonName(v ParamCommonName) *ParamCommonName {
	return &v
}

func (t *ParamCommonName) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cCommonName, ErrInvalidParam)
	}

	*t = ParamCommonName(decode6868(vs[0].Data))

	return nil
}

func (t ParamCommonName) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";CN=")

	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.Write(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.Write(encode6868(string(t)))
	}
}

func (t ParamCommonName) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return fmt.Errorf(errValidatingType, cCommonName, ErrInvalidText)
	}

	return nil
}

// ParamCalendarUserType.
type ParamCalendarUserType uint8

// CalendarUserType constant values.
const (
	CalendarUserTypeUnknown ParamCalendarUserType = iota
	CalendarUserTypeIndividual
	CalendarUserTypeGroup
	CalendarUserTypeResource
	CalendarUserTypeRoom
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamCalendarUserType) New() *ParamCalendarUserType {
	return &t
}

func (t *ParamCalendarUserType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cCalendarUserType, ErrInvalidParam)
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

func (t ParamCalendarUserType) encode(w writer) {
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

func (t ParamCalendarUserType) valid() error {
	return nil
}

// ParamDelegator.
type ParamDelegator []CalendarAddress

func (t *ParamDelegator) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return fmt.Errorf(errDecodingType, cDelegator, ErrInvalidParam)
		}

		var q CalendarAddress

		if err := q.decode(nil, v.Data); err != nil {
			return fmt.Errorf(errDecodingType, cDelegator, err)
		}

		*t = append(*t, q)
	}

	return nil
}

func (t ParamDelegator) encode(w writer) {
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

func (t ParamDelegator) valid() error {
	for _, v := range t {
		if err := v.valid(); err != nil {
			return fmt.Errorf(errValidatingType, cDelegator, err)
		}
	}

	return nil
}

// ParamDelagatee.
type ParamDelagatee []CalendarAddress

func (t *ParamDelagatee) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return fmt.Errorf(errDecodingType, cDelagatee, ErrInvalidParam)
		}

		var q CalendarAddress

		if err := q.decode(nil, v.Data); err != nil {
			return fmt.Errorf(errDecodingType, cDelagatee, err)
		}

		*t = append(*t, q)
	}

	return nil
}

func (t ParamDelagatee) encode(w writer) {
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

func (t ParamDelagatee) valid() error {
	for _, v := range t {
		if err := v.valid(); err != nil {
			return fmt.Errorf(errValidatingType, cDelagatee, err)
		}
	}

	return nil
}

// ParamDirectoryEntry.
type ParamDirectoryEntry string

// NewDirectoryEntry returns a *ParamDirectoryEntry for ease of use with optional values.
func NewDirectoryEntry(v ParamDirectoryEntry) *ParamDirectoryEntry {
	return &v
}

func (t *ParamDirectoryEntry) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cDirectoryEntry, ErrInvalidParam)
	}

	if vs[0].Type != tokenParamQuotedValue {
		return fmt.Errorf(errDecodingType, cDirectoryEntry, ErrInvalidParam)
	}

	*t = ParamDirectoryEntry(decode6868(vs[0].Data))

	return nil
}

func (t ParamDirectoryEntry) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";DIR=")
	w.WriteString("\"")
	w.Write(encode6868(string(t)))
	w.WriteString("\"")
}

func (t ParamDirectoryEntry) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return fmt.Errorf(errValidatingType, cDirectoryEntry, ErrInvalidText)
	}

	return nil
}

// ParamEncoding.
type ParamEncoding uint8

// Encoding constant values.
const (
	Encoding8bit ParamEncoding = iota
	EncodingBase64
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamEncoding) New() *ParamEncoding {
	return &t
}

func (t *ParamEncoding) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cEncoding, ErrInvalidParam)
	}

	switch strings.ToUpper(vs[0].Data) {
	case "8BIT":
		*t = Encoding8bit
	case "BASE64":
		*t = EncodingBase64
	default:
		return fmt.Errorf(errDecodingType, cEncoding, ErrInvalidParam)
	}

	return nil
}

func (t ParamEncoding) encode(w writer) {
	w.WriteString(";ENCODING=")

	switch t {
	case Encoding8bit:
		w.WriteString("8BIT")
	case EncodingBase64:
		w.WriteString("BASE64")
	}
}

func (t ParamEncoding) valid() error {
	switch t {
	case Encoding8bit, EncodingBase64:
	default:
		return fmt.Errorf(errValidatingType, cEncoding, ErrInvalidValue)
	}

	return nil
}

// ParamFormatType.
type ParamFormatType string

// NewFormatType returns a *ParamFormatType for ease of use with optional values.
func NewFormatType(v ParamFormatType) *ParamFormatType {
	return &v
}

var regexFormatType *regexp.Regexp

func (t *ParamFormatType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cFormatType, ErrInvalidParam)
	}

	if !regexFormatType.MatchString(vs[0].Data) {
		return fmt.Errorf(errDecodingType, cFormatType, ErrInvalidParam)
	}

	*t = ParamFormatType(vs[0].Data)

	return nil
}

func (t ParamFormatType) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";FMTTYPE=")

	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.Write(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.Write(encode6868(string(t)))
	}
}

func (t ParamFormatType) valid() error {
	if !regexFormatType.Match([]byte(t)) {
		return fmt.Errorf(errValidatingType, cFormatType, ErrInvalidValue)
	}

	return nil
}

// ParamFreeBusyType.
type ParamFreeBusyType uint8

// FreeBusyType constant values.
const (
	FreeBusyTypeUnknown ParamFreeBusyType = iota
	FreeBusyTypeFree
	FreeBusyTypeBusy
	FreeBusyTypeBusyUnavailable
	FreeBusyTypeBusyTentative
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamFreeBusyType) New() *ParamFreeBusyType {
	return &t
}

func (t *ParamFreeBusyType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cFreeBusyType, ErrInvalidParam)
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

func (t ParamFreeBusyType) encode(w writer) {
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

func (t ParamFreeBusyType) valid() error {
	return nil
}

// ParamLanguage.
type ParamLanguage string

// NewLanguage returns a *ParamLanguage for ease of use with optional values.
func NewLanguage(v ParamLanguage) *ParamLanguage {
	return &v
}

func (t *ParamLanguage) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cLanguage, ErrInvalidParam)
	}

	*t = ParamLanguage(decode6868(vs[0].Data))

	return nil
}

func (t ParamLanguage) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";LANGUAGE=")

	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.Write(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.Write(encode6868(string(t)))
	}
}

func (t ParamLanguage) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return fmt.Errorf(errValidatingType, cLanguage, ErrInvalidText)
	}

	return nil
}

// ParamMember.
type ParamMember []string

// NewMember returns a *ParamMember for ease of use with optional values.
func NewMember(v ParamMember) *ParamMember {
	return &v
}

func (t *ParamMember) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return fmt.Errorf(errDecodingType, cMember, ErrInvalidParam)
		}

		*t = append(*t, decode6868(v.Data))
	}

	return nil
}

func (t ParamMember) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";MEMBER=")

	for n, v := range t {
		if n > 0 {
			w.WriteString(",")
		}

		w.WriteString("\"")
		w.Write(encode6868(string(v)))
		w.WriteString("\"")
	}
}

func (t ParamMember) valid() error {
	for _, v := range t {
		if strings.ContainsAny(string(v), nonsafeChars[:31]) {
			return fmt.Errorf(errValidatingType, cMember, ErrInvalidText)
		}
	}

	return nil
}

// ParamParticipationStatus.
type ParamParticipationStatus uint8

// ParticipationStatus constant values.
const (
	ParticipationStatusUnknown ParamParticipationStatus = iota
	ParticipationStatusNeedsAction
	ParticipationStatusAccepted
	ParticipationStatusDeclined
	ParticipationStatusTentative
	ParticipationStatusDelegated
	ParticipationStatusCompleted
	ParticipationStatusInProcess
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamParticipationStatus) New() *ParamParticipationStatus {
	return &t
}

func (t *ParamParticipationStatus) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cParticipationStatus, ErrInvalidParam)
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

func (t ParamParticipationStatus) encode(w writer) {
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

func (t ParamParticipationStatus) valid() error {
	return nil
}

// ParamRange.
type ParamRange struct{}

func (t *ParamRange) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cRange, ErrInvalidParam)
	}

	if strings.ToUpper(vs[0].Data) != "THISANDFUTURE" {
		return fmt.Errorf(errDecodingType, cRange, ErrInvalidParam)
	}

	return nil
}

func (t ParamRange) encode(w writer) {
	w.WriteString(";RANGE=")
	w.WriteString("THISANDFUTURE")
}

func (t ParamRange) valid() error {
	return nil
}

// ParamRelated.
type ParamRelated uint8

// Related constant values.
const (
	RelatedStart ParamRelated = iota
	RelatedEnd
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamRelated) New() *ParamRelated {
	return &t
}

func (t *ParamRelated) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cRelated, ErrInvalidParam)
	}

	switch strings.ToUpper(vs[0].Data) {
	case "START":
		*t = RelatedStart
	case "END":
		*t = RelatedEnd
	default:
		return fmt.Errorf(errDecodingType, cRelated, ErrInvalidParam)
	}

	return nil
}

func (t ParamRelated) encode(w writer) {
	w.WriteString(";RELATED=")

	switch t {
	case RelatedStart:
		w.WriteString("START")
	case RelatedEnd:
		w.WriteString("END")
	}
}

func (t ParamRelated) valid() error {
	switch t {
	case RelatedStart, RelatedEnd:
	default:
		return fmt.Errorf(errValidatingType, cRelated, ErrInvalidValue)
	}

	return nil
}

// ParamRelationshipType.
type ParamRelationshipType uint8

// RelationshipType constant values.
const (
	RelationshipTypeUnknown ParamRelationshipType = iota
	RelationshipTypeParent
	RelationshipTypeChild
	RelationshipTypeSibling
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamRelationshipType) New() *ParamRelationshipType {
	return &t
}

func (t *ParamRelationshipType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cRelationshipType, ErrInvalidParam)
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

func (t ParamRelationshipType) encode(w writer) {
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

func (t ParamRelationshipType) valid() error {
	return nil
}

// ParamParticipationRole.
type ParamParticipationRole uint8

// ParticipationRole constant values.
const (
	ParticipationRoleUnknown ParamParticipationRole = iota
	ParticipationRoleRequiredParticipant
	ParticipationRoleChair
	ParticipationRoleOptParticipant
	ParticipationRoleNonParticipant
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamParticipationRole) New() *ParamParticipationRole {
	return &t
}

func (t *ParamParticipationRole) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cParticipationRole, ErrInvalidParam)
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

func (t ParamParticipationRole) encode(w writer) {
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

func (t ParamParticipationRole) valid() error {
	return nil
}

// ParamRSVP.
type ParamRSVP Boolean

// NewRSVP returns a *ParamRSVP for ease of use with optional values.
func NewRSVP(v ParamRSVP) *ParamRSVP {
	return &v
}

func (t *ParamRSVP) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cRSVP, ErrInvalidParam)
	}

	var q Boolean

	if err := q.decode(nil, vs[0].Data); err != nil {
		return fmt.Errorf(errDecodingType, cRSVP, err)
	}

	*t = ParamRSVP(q)

	return nil
}

func (t ParamRSVP) encode(w writer) {
	if !t {
		return
	}

	w.WriteString(";RSVP=")

	q := Boolean(t)

	q.encode(w)
}

func (t ParamRSVP) valid() error {
	return nil
}

// ParamSentBy.
type ParamSentBy string

// NewSentBy returns a *ParamSentBy for ease of use with optional values.
func NewSentBy(v ParamSentBy) *ParamSentBy {
	return &v
}

func (t *ParamSentBy) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cSentBy, ErrInvalidParam)
	}

	if vs[0].Type != tokenParamQuotedValue {
		return fmt.Errorf(errDecodingType, cSentBy, ErrInvalidParam)
	}

	*t = ParamSentBy(decode6868(vs[0].Data))

	return nil
}

func (t ParamSentBy) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";SENT-BY=")
	w.WriteString("\"")
	w.Write(encode6868(string(t)))
	w.WriteString("\"")
}

func (t ParamSentBy) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return fmt.Errorf(errValidatingType, cSentBy, ErrInvalidText)
	}

	return nil
}

// ParamTimezoneID.
type ParamTimezoneID string

// NewTimezoneID returns a *ParamTimezoneID for ease of use with optional values.
func NewTimezoneID(v ParamTimezoneID) *ParamTimezoneID {
	return &v
}

func (t *ParamTimezoneID) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cTimezoneID, ErrInvalidParam)
	}

	*t = ParamTimezoneID(decode6868(vs[0].Data))

	return nil
}

func (t ParamTimezoneID) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";TZID=")

	if strings.ContainsAny(string(t), nonsafeChars[32:]) {
		w.WriteString("\"")
		w.Write(encode6868(string(t)))
		w.WriteString("\"")
	} else {
		w.Write(encode6868(string(t)))
	}
}

func (t ParamTimezoneID) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return fmt.Errorf(errValidatingType, cTimezoneID, ErrInvalidText)
	}

	return nil
}

// ParamValue.
type ParamValue uint8

// Value constant values.
const (
	ValueUnknown ParamValue = iota
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

// New returns a pointer to the type (used with constants for ease of use with
// optional values).
func (t ParamValue) New() *ParamValue {
	return &t
}

func (t *ParamValue) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cValue, ErrInvalidParam)
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

func (t ParamValue) encode(w writer) {
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

func (t ParamValue) valid() error {
	return nil
}

// ParamURI.
type ParamURI URI

func (t *ParamURI) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cURI, ErrInvalidParam)
	}

	if vs[0].Type != tokenParamQuotedValue {
		return fmt.Errorf(errDecodingType, cURI, ErrInvalidParam)
	}

	var q URI

	if err := q.decode(nil, vs[0].Data); err != nil {
		return fmt.Errorf(errDecodingType, cURI, err)
	}

	*t = ParamURI(q)

	return nil
}

func (t ParamURI) encode(w writer) {
	if len(t.String()) == 0 {
		return
	}

	w.WriteString(";URI=")

	q := URI(t)

	q.encode(w)
}

func (t ParamURI) valid() error {
	q := URI(t)

	if err := q.valid(); err != nil {
		return fmt.Errorf(errValidatingType, cURI, err)
	}

	return nil
}

// ParamID.
type ParamID string

// NewID returns a *ParamID for ease of use with optional values.
func NewID(v ParamID) *ParamID {
	return &v
}

func (t *ParamID) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cID, ErrInvalidParam)
	}

	if vs[0].Type != tokenParamQuotedValue {
		return fmt.Errorf(errDecodingType, cID, ErrInvalidParam)
	}

	*t = ParamID(decode6868(vs[0].Data))

	return nil
}

func (t ParamID) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";ID=")
	w.WriteString("\"")
	w.Write(encode6868(string(t)))
	w.WriteString("\"")
}

func (t ParamID) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return fmt.Errorf(errValidatingType, cID, ErrInvalidText)
	}

	return nil
}

// ParamAgentID.
type ParamAgentID string

// NewAgentID returns a *ParamAgentID for ease of use with optional values.
func NewAgentID(v ParamAgentID) *ParamAgentID {
	return &v
}

func (t *ParamAgentID) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return fmt.Errorf(errDecodingType, cAgentID, ErrInvalidParam)
	}

	if vs[0].Type != tokenParamQuotedValue {
		return fmt.Errorf(errDecodingType, cAgentID, ErrInvalidParam)
	}

	*t = ParamAgentID(decode6868(vs[0].Data))

	return nil
}

func (t ParamAgentID) encode(w writer) {
	if len(t) == 0 {
		return
	}

	w.WriteString(";AGENT-ID=")
	w.WriteString("\"")
	w.Write(encode6868(string(t)))
	w.WriteString("\"")
}

func (t ParamAgentID) valid() error {
	if strings.ContainsAny(string(t), nonsafeChars[:31]) {
		return fmt.Errorf(errValidatingType, cAgentID, ErrInvalidText)
	}

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
			switch t.Next() {
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
		}
	}
	return string(d)
}

func encode6868(s string) []byte {
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
	return d
}

func init() {
	regexFormatType = regexp.MustCompile("[A-Za-z0-9!#$&.+-^_]/[A-Za-z0-9!#$&.+-^_]")
}

// Errors.
var (
	ErrInvalidParam = errors.New("invalid param value")
	ErrInvalidValue = errors.New("invalid value")
)

const (
	errDecodingType            = "error decoding %s: %w"
	errValidatingType          = "error decoding %s: %w"
	cAlternativeRepresentation = "AlternativeRepresentation"
	cCommonName                = "CommonName"
	cCalendarUserType          = "CalendarUserType"
	cDelegator                 = "Delegator"
	cDelagatee                 = "Delagatee"
	cDirectoryEntry            = "DirectoryEntry"
	cEncoding                  = "Encoding"
	cFormatType                = "FormatType"
	cFreeBusyType              = "FreeBusyType"
	cLanguage                  = "Language"
	cMember                    = "Member"
	cParticipationStatus       = "ParticipationStatus"
	cRange                     = "Range"
	cRelated                   = "Related"
	cRelationshipType          = "RelationshipType"
	cParticipationRole         = "ParticipationRole"
	cRSVP                      = "RSVP"
	cSentBy                    = "SentBy"
	cTimezoneID                = "TimezoneID"
	cValue                     = "Value"
	cURI                       = "URI"
	cID                        = "ID"
	cAgentID                   = "AgentID"
)
