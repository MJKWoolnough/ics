package ics

// File automatically generated with ./genParams.sh

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"vimagination.zapto.org/errors"
	"vimagination.zapto.org/parser"
)

// AlternativeRepresentation is an alternate text representation for the
// property value
type ParamAlternativeRepresentation URI

func (t *ParamAlternativeRepresentation) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding AlternativeRepresentation: ", ErrInvalidParam)
	}
	if vs[0].Type != tokenParamQuotedValue {
		return errors.WithContext("error decoding AlternativeRepresentation: ", ErrInvalidParam)
	}
	var q URI
	if err := q.decode(nil, vs[0].Data); err != nil {
		return errors.WithContext("error decoding AlternativeRepresentation: ", err)
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
		return errors.WithContext("error validating AlternativeRepresentation: ", err)
	}
	return nil
}

// CommonName is the common name to be associated with the calendar user
// specified by the property
type ParamCommonName string

// NewCommonName returns a *ParamCommonName for ease of use with optional values
func NewCommonName(v ParamCommonName) *ParamCommonName {
	return &v
}

func (t *ParamCommonName) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding CommonName: ", ErrInvalidParam)
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
		return errors.WithContext("error validating CommonName: ", ErrInvalidText)
	}
	return nil
}

// CalendarUserType is identify the type of calendar user specified by the
// property
type ParamCalendarUserType uint8

// CalendarUserType constant values
const (
	CalendarUserTypeUnknown ParamCalendarUserType = iota
	CalendarUserTypeIndividual
	CalendarUserTypeGroup
	CalendarUserTypeResource
	CalendarUserTypeRoom
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (t ParamCalendarUserType) New() *ParamCalendarUserType {
	return &t
}

func (t *ParamCalendarUserType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding CalendarUserType: ", ErrInvalidParam)
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

// Delegator is used to specify the calendar users that have delegated their
// participation to the calendar user specified by the property
type ParamDelegator []CalendarAddress

func (t *ParamDelegator) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return errors.WithContext("error decoding Delegator: ", ErrInvalidParam)
		}
		var q CalendarAddress
		if err := q.decode(nil, v.Data); err != nil {
			return errors.WithContext("error decoding Delegator: ", err)
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
			return errors.WithContext("error validation Delegator: ", err)
		}
	}
	return nil
}

// Delagatee is used to specify the calendar users to whom the calendar user
// specified by the property has delegated participation
type ParamDelagatee []CalendarAddress

func (t *ParamDelagatee) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return errors.WithContext("error decoding Delagatee: ", ErrInvalidParam)
		}
		var q CalendarAddress
		if err := q.decode(nil, v.Data); err != nil {
			return errors.WithContext("error decoding Delagatee: ", err)
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
			return errors.WithContext("error validation Delagatee: ", err)
		}
	}
	return nil
}

// DirectoryEntry is a reference to a directory entry associated with the
// calendar
type ParamDirectoryEntry string

// NewDirectoryEntry returns a *ParamDirectoryEntry for ease of use with optional values
func NewDirectoryEntry(v ParamDirectoryEntry) *ParamDirectoryEntry {
	return &v
}

func (t *ParamDirectoryEntry) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding DirectoryEntry: ", ErrInvalidParam)
	}
	if vs[0].Type != tokenParamQuotedValue {
		return errors.WithContext("error decoding DirectoryEntry: ", ErrInvalidParam)
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
		return errors.WithContext("error validating DirectoryEntry: ", ErrInvalidText)
	}
	return nil
}

// Encoding is the inline encoding for the property value
type ParamEncoding uint8

// Encoding constant values
const (
	Encoding8bit ParamEncoding = iota
	EncodingBase64
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (t ParamEncoding) New() *ParamEncoding {
	return &t
}

func (t *ParamEncoding) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding Encoding: ", ErrInvalidParam)
	}
	switch strings.ToUpper(vs[0].Data) {
	case "8BIT":
		*t = Encoding8bit
	case "BASE64":
		*t = EncodingBase64
	default:
		return errors.WithContext("error decoding Encoding: ", ErrInvalidParam)
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
		return errors.WithContext("error validating Encoding: ", ErrInvalidValue)
	}
	return nil
}

// FormatType is the content type of a referenced object
type ParamFormatType string

// NewFormatType returns a *ParamFormatType for ease of use with optional values
func NewFormatType(v ParamFormatType) *ParamFormatType {
	return &v
}

var regexFormatType *regexp.Regexp

func (t *ParamFormatType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding FormatType: ", ErrInvalidParam)
	}
	if !regexFormatType.MatchString(vs[0].Data) {
		return errors.WithContext("error decoding FormatType: ", ErrInvalidParam)
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
		return errors.WithContext("error validating FormatType: ", ErrInvalidValue)
	}
	return nil
}

// FreeBusyType is used to specify the free or busy time type
type ParamFreeBusyType uint8

// FreeBusyType constant values
const (
	FreeBusyTypeUnknown ParamFreeBusyType = iota
	FreeBusyTypeFree
	FreeBusyTypeBusy
	FreeBusyTypeBusyUnavailable
	FreeBusyTypeBusyTentative
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (t ParamFreeBusyType) New() *ParamFreeBusyType {
	return &t
}

func (t *ParamFreeBusyType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding FreeBusyType: ", ErrInvalidParam)
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

// Language is the language for text values
type ParamLanguage string

// NewLanguage returns a *ParamLanguage for ease of use with optional values
func NewLanguage(v ParamLanguage) *ParamLanguage {
	return &v
}

func (t *ParamLanguage) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding Language: ", ErrInvalidParam)
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
		return errors.WithContext("error validating Language: ", ErrInvalidText)
	}
	return nil
}

// Member is used to specify the group or list membership of the calendar
type ParamMember []string

// NewMember returns a *ParamMember for ease of use with optional values
func NewMember(v ParamMember) *ParamMember {
	return &v
}

func (t *ParamMember) decode(vs []parser.Token) error {
	for _, v := range vs {
		if v.Type != tokenParamQuotedValue {
			return errors.WithContext("error decoding Member: ", ErrInvalidParam)
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
			return errors.WithContext("error validating Member: ", ErrInvalidText)
		}
	}
	return nil
}

// ParticipationStatus is used to specify the participation status for the
// calendar
type ParamParticipationStatus uint8

// ParticipationStatus constant values
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
// optional values)
func (t ParamParticipationStatus) New() *ParamParticipationStatus {
	return &t
}

func (t *ParamParticipationStatus) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding ParticipationStatus: ", ErrInvalidParam)
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

// Range is used to specify the effective range of recurrence instances from the
// instance specified by the recurrence identifier specified by the property
type ParamRange struct{}

func (t *ParamRange) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding Range: ", ErrInvalidParam)
	}
	if strings.ToUpper(vs[0].Data) != "THISANDFUTURE" {
		return errors.WithContext("error decoding Range", ErrInvalidParam)
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

// Related is the relationship of the alarm trigger with respect to the start or
// end of the calendar component
type ParamRelated uint8

// Related constant values
const (
	RelatedStart ParamRelated = iota
	RelatedEnd
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (t ParamRelated) New() *ParamRelated {
	return &t
}

func (t *ParamRelated) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding Related: ", ErrInvalidParam)
	}
	switch strings.ToUpper(vs[0].Data) {
	case "START":
		*t = RelatedStart
	case "END":
		*t = RelatedEnd
	default:
		return errors.WithContext("error decoding Related: ", ErrInvalidParam)
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
		return errors.WithContext("error validating Related: ", ErrInvalidValue)
	}
	return nil
}

// RelationshipType is the type of hierarchical relationship associated with the
// calendar component specified by the property
type ParamRelationshipType uint8

// RelationshipType constant values
const (
	RelationshipTypeUnknown ParamRelationshipType = iota
	RelationshipTypeParent
	RelationshipTypeChild
	RelationshipTypeSibling
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (t ParamRelationshipType) New() *ParamRelationshipType {
	return &t
}

func (t *ParamRelationshipType) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding RelationshipType: ", ErrInvalidParam)
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

// ParticipationRole is used to specify the participation role for the calendar
// user specified by the property
type ParamParticipationRole uint8

// ParticipationRole constant values
const (
	ParticipationRoleUnknown ParamParticipationRole = iota
	ParticipationRoleRequiredParticipant
	ParticipationRoleChair
	ParticipationRoleOptParticipant
	ParticipationRoleNonParticipant
)

// New returns a pointer to the type (used with constants for ease of use with
// optional values)
func (t ParamParticipationRole) New() *ParamParticipationRole {
	return &t
}

func (t *ParamParticipationRole) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding ParticipationRole: ", ErrInvalidParam)
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

// RSVP is used to specify whether there is an expectation of a favor of a reply
// from the calendar user specified by the property value
type ParamRSVP Boolean

// NewRSVP returns a *ParamRSVP for ease of use with optional values
func NewRSVP(v ParamRSVP) *ParamRSVP {
	return &v
}

func (t *ParamRSVP) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding RSVP: ", ErrInvalidParam)
	}
	var q Boolean
	if err := q.decode(nil, vs[0].Data); err != nil {
		return errors.WithContext("error decoding RSVP: ", err)
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

// SentBy is used to specify the calendar user that is acting on behalf of the
// calendar user specified by the property
type ParamSentBy string

// NewSentBy returns a *ParamSentBy for ease of use with optional values
func NewSentBy(v ParamSentBy) *ParamSentBy {
	return &v
}

func (t *ParamSentBy) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding SentBy: ", ErrInvalidParam)
	}
	if vs[0].Type != tokenParamQuotedValue {
		return errors.WithContext("error decoding SentBy: ", ErrInvalidParam)
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
		return errors.WithContext("error validating SentBy: ", ErrInvalidText)
	}
	return nil
}

// TimezoneID is used to specify the identifier for the time zone definition for
// a time component in the property value
type ParamTimezoneID string

// NewTimezoneID returns a *ParamTimezoneID for ease of use with optional values
func NewTimezoneID(v ParamTimezoneID) *ParamTimezoneID {
	return &v
}

func (t *ParamTimezoneID) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding TimezoneID: ", ErrInvalidParam)
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
		return errors.WithContext("error validating TimezoneID: ", ErrInvalidText)
	}
	return nil
}

// Value is used to explicitly specify the value type format for a property
// value
type ParamValue uint8

// Value constant values
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
// optional values)
func (t ParamValue) New() *ParamValue {
	return &t
}

func (t *ParamValue) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding Value: ", ErrInvalidParam)
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

// URI
type ParamURI URI

func (t *ParamURI) decode(vs []parser.Token) error {
	if len(vs) != 1 {
		return errors.WithContext("error decoding URI: ", ErrInvalidParam)
	}
	if vs[0].Type != tokenParamQuotedValue {
		return errors.WithContext("error decoding URI: ", ErrInvalidParam)
	}
	var q URI
	if err := q.decode(nil, vs[0].Data); err != nil {
		return errors.WithContext("error decoding URI: ", err)
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
		return errors.WithContext("error validating URI: ", err)
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

// Errors
const (
	ErrInvalidParam errors.Error = "invalid param value"
	ErrInvalidValue errors.Error = "invalid value"
)
