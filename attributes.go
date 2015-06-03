package ics

import "errors"

const (
	altrepparam    = "ALTREP"
	cnparam        = "CN"
	cutypeparam    = "CUTYPE"
	delfromparam   = "DELEGATED-FROM"
	deltoparam     = "DELEGATED-TO"
	dirparam       = "DIR"
	encodingparam  = "ENCODING"
	fmttypeparam   = "FMTTYPE"
	fbtypeparam    = "FBTYPE"
	languageparam  = "LANGUAGE"
	memberparam    = "MEMBER"
	partstatparam  = "PARTSTAT"
	rangeparam     = "RANGE"
	trigrelparam   = "RELATED"
	reltypeparam   = "RELTYPE"
	roleparam      = "ROLE"
	rsvpparam      = "RSVP"
	sentbyparam    = "SENT-BY"
	tzidparam      = "TZID"
	valuetypeparam = "VALUE"
)

type attribute interface {
	Bytes() []byte
	String() string
}

type altrep string

func newAltRepParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	if vs[0].typ != tokenParamQValue {
		return nil, ErrIncorrectParamValueType
	}
	return altrep(vs[0].data), nil
}

func (a altrep) Bytes() []byte {
	return dquote(escape6868(string(a)))
}

func (a altrep) String() string {
	return string(a)
}

type commonName string

func newCommonNameParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	return commonName(vs[0].data), nil
}

func (c commonName) Bytes() []byte {
	return escape6868(string(c))
}

func (c commonName) String() string {
	return string(c)
}

const (
	cuIndividual calendarUserType = iota
	cuGroup
	cuResource
	cuRoom
	cuUnknown
)

type calendarUserType int

func newCalendarUserTypeParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "INDIVIDUAL":
		return cuIndividual, nil
	case "GROUP":
		return cuGroup, nil
	case "RESOURCE":
		return cuResource, nil
	case "ROOM":
		return cuRoom, nil
	default:
		return cuUnknown, nil
	}
}

func (c calendarUserType) Bytes() []byte {
	return []byte(c.String())
}

func (c calendarUserType) String() string {
	switch c {
	case cuIndividual:
		return "INDIVIDUAL"
	case cuGroup:
		return "GROUP"
	case cuResource:
		return "RESOURCE"
	case cuRoom:
		return "ROOM"
	default:
		return "UNKNOWN"
	}
}

type delegators []string

func newDelegatorsParam(vs []token) (attribute, error) {
	if len(vs) == 0 {
		return nil, ErrIncorrectNumParamValues
	}
	d := make(delegators, 0, 1)
	for _, v := range vs {
		if v.typ != tokenParamQValue {
			return nil, ErrIncorrectParamValueType
		}
		d = append(d, v.data)
	}
	return d, nil
}

func (d delegators) Bytes() []byte {
	var ds []byte
	for n, dg := range d {
		if n > 0 {
			ds = append(ds, ',')
		}
		ds = append(ds, dquote(escape6868(dg))...)
	}
	return ds
}

func (d delegators) String() string {
	return string(d.Bytes())
}

type delegatee []string

func newDelegateeParam(vs []token) (attribute, error) {
	if len(vs) == 0 {
		return nil, ErrIncorrectNumParamValues
	}
	d := make(delegatee, 0, 1)
	for _, v := range vs {
		if v.typ != tokenParamQValue {
			return nil, ErrIncorrectParamValueType
		}
		d = append(d, v.data)
	}
	return d, nil
}

func (d delegatee) Bytes() []byte {
	var ds []byte
	for n, dg := range d {
		if n > 0 {
			ds = append(ds, ',')
		}
		ds = append(ds, dquote(escape6868(dg))...)
	}
	return ds
}

func (d delegatee) String() string {
	return string(d.Bytes())
}

type directoryEntryRef string

func newDirectoryEntryRefParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	if vs[0].typ != tokenParamQValue {
		return nil, ErrIncorrectParamValueType
	}
	return directoryEntryRef(vs[0].data), nil
}

func (d directoryEntryRef) Bytes() []byte {
	return dquote([]byte(d))
}

func (d directoryEntryRef) String() string {
	return string(d)
}

const (
	encoding8Bit encoding = iota
	encodingBase64
)

type encoding int

func newEncodingParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "8BIT":
		return encoding8Bit, nil
	case "BASE64":
		return encodingBase64, nil
	default:
		return nil, ErrUnsupportedParamValue
	}
}

func (e encoding) Bytes() []byte {
	return []byte(e.String())
}

func (e encoding) String() string {
	switch e {
	case encoding8Bit:
		return "8BIT"
	case encodingBase64:
		return "BASE64"
	default:
		return "unknownencoding"
	}
}

type fmtType string

func newFmtTypeParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	//validate value
	return fmtType(vs[0].data), nil
}

func (f fmtType) Bytes() []byte {
	return dquoteIfNeeded([]byte(f))
}

func (f fmtType) String() string {
	return string(f)
}

const (
	fbBusy freeBusy = iota
	fbFree
	fbBusyUnavailable
	fbBusyTentative
)

type freeBusy int

func newFreeBusyParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "BUSY":
		return fbBusy, nil
	case "FREE":
		return fbFree, nil
	case "BUSY-UNAVAILABLE":
		return fbBusyUnavailable, nil
	case "BUSY-TENTATIVE":
		return fbBusyTentative, nil
	default:
		return fbBusy, nil
	}
}

func (fb freeBusy) Bytes() []byte {
	return []byte(fb.String())
}

func (fb freeBusy) String() string {
	switch fb {
	case fbBusy:
		return "BUSY"
	case fbFree:
		return "FREE"
	case fbBusyUnavailable:
		return "BUSY-UNAVAILABLE"
	case fbBusyTentative:
		return "BUSY-TENTATIVE"
	default:
		return "BUSY"
	}
}

type language string

func newLanguageParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	if vs[0].typ != tokenParamQValue {
		return nil, ErrIncorrectParamValueType
	}
	return language(vs[0].data), nil
}

func (l language) Bytes() []byte {
	return dquote([]byte(l))
}

func (l language) String() string {
	return string(l)
}

type members []string

func newMemberParam(vs []token) (attribute, error) {
	if len(vs) == 0 {
		return nil, ErrIncorrectNumParamValues
	}
	m := make(members, 0, 1)
	for _, v := range vs {
		if v.typ != tokenParamQValue {
			return nil, ErrIncorrectParamValueType
		}
		m = append(m, v.data)
	}
	return m, nil
}

func (ms members) Bytes() []byte {
	var mstr []byte
	for n, m := range ms {
		if n > 0 {
			mstr = append(mstr, ',')
		}
		mstr = append(mstr, dquote(escape6868(m))...)
	}
	return mstr
}

func (ms members) String() string {
	return string(ms.Bytes())
}

const (
	psNeedsAction participationStatus = iota
	psAccepted
	psDeclined
	psTentative
	psDelegated
	psCompleted
	psInProgress
)

type participationStatus int

func newParticipationStatusParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "NEEDS-ACTION":
		return psNeedsAction, nil
	case "ACCEPTED":
		return psAccepted, nil
	case "DECLINED":
		return psDeclined, nil
	case "TENTATIVE":
		return psTentative, nil
	case "DELEGATED":
		return psDelegated, nil
	case "COMPLETED":
		return psCompleted, nil
	case "IN-PROGRESS":
		return psInProgress, nil
	default:
		return psNeedsAction, nil
	}
}

func (p participationStatus) Bytes() []byte {
	return []byte(p.String())
}

func (p participationStatus) String() string {
	switch p {
	case psNeedsAction:
		return "NEEDS-ACTION"
	case psAccepted:
		return "ACCEPTED"
	case psDeclined:
		return "DECLINED"
	case psTentative:
		return "TENTATIVE"
	case psDelegated:
		return "DELEGATED"
	case psCompleted:
		return "COMPLETED"
	case psInProgress:
		return "IN-PROGRESS"
	default:
		return "NEEDS-ACTION"
	}
}

func (p participationStatus) ValidForEvent() bool {
	switch p {
	case psNeedsAction, psAccepted, psDeclined, psTentative, psDelegated:
		return true
	default:
		return false
	}
}

func (p participationStatus) ValidForTodo() bool {
	switch p {
	case psNeedsAction, psAccepted, psDeclined, psTentative, psDelegated, psCompleted, psInProgress:
		return true
	default:
		return false
	}
}

func (p participationStatus) ValidForJournal() bool {
	switch p {
	case psNeedsAction, psAccepted, psDeclined:
		return true
	default:
		return false
	}
}

const (
	rngThisAndFuture rangeParam = iota
	rngThisAndPrior
)

type rangeParam int

func newRangeParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "THISANDFUTURE":
		return rngThisAndFuture, nil
	case "THISANDPRIOR":
		return rngThisAndPrior, nil
	default:
		return nil, ErrUnsupportedParamValue
	}
}

func (r rangeParam) Bytes() []byte {
	return []byte(r.String())
}

func (r rangeParam) String() string {
	switch r {
	case rngThisAndFuture:
		return "THISANDFUTURE"
	case rngThisAndPrior:
		return "THISANDPRIOR"
	default:
		return "unknown"
	}
}

const (
	atrStart alarmTriggerRelationship = iota
	atrEnd
)

type alarmTriggerRelationship int

func newAlarmTriggerRelationshipParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "START":
		return atrStart, nil
	case "END":
		return atrEnd, nil
	default:
		return atrStart, nil
	}
}

func (a alarmTriggerRelationship) Bytes() []byte {
	return []byte(a.String())
}

func (a alarmTriggerRelationship) String() string {
	switch a {
	case atrStart:
		return "START"
	case atrEnd:
		return "END"
	default:
		return "START"
	}
}

const (
	rtParent relationshipType = iota
	rtChild
	rtSibling
)

type relationshipType int

func newRelationshipTypeParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "PARENT":
		return rtParent, nil
	case "CHILD":
		return rtChild, nil
	case "SIBLING":
		return rtSibling, nil
	default:
		return rtParent, nil
	}
}

func (r relationshipType) Bytes() []byte {
	return []byte(r.String())
}

func (r relationshipType) String() string {
	switch r {
	case rtParent:
		return "PARENT"
	case rtChild:
		return "CHILD"
	case rtSibling:
		return "SIBLING"
	default:
		return "PARENT"
	}
}

const (
	prRequiredParticipant participationRole = iota
	prChair
	prOptionalParticipant
	prNonParticipant
)

type participationRole int

func newParticipationRoleParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "REQ-PARTICIPANT":
		return prRequiredParticipant, nil
	case "CHAIR":
		return prChair, nil
	case "OPT-PARTICIPANT":
		return prOptionalParticipant, nil
	case "NON-PARTICIPANT":
		return prNonParticipant, nil
	default:
		return prRequiredParticipant, nil
	}
}

func (p participationRole) Bytes() []byte {
	return []byte(p.String())
}

func (p participationRole) String() string {
	switch p {
	case prRequiredParticipant:
		return "REQ-PARTICIPANT"
	case prChair:
		return "CHAIR"
	case prOptionalParticipant:
		return "OPT-PARTICIPANT"
	case prNonParticipant:
		return "NON-PARTICIPANT"
	default:
		return "REQ-PARTICIPANT"
	}
}

type rsvp bool

func newRSVPExpectationParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "TRUE":
		return rsvp(true), nil
	case "FALSE":
		return rsvp(false), nil
	default:
		return nil, ErrUnsupportedParamValue
	}
}

func (r rsvp) Bytes() []byte {
	return []byte(r.String())
}

func (r rsvp) String() string {
	if r {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

type sentBy string

func newSentByParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	if vs[0].typ != tokenParamQValue {
		return nil, ErrIncorrectParamValueType
	}
	return sentBy(vs[0].data), nil
}

func (s sentBy) Bytes() []byte {
	return dquote(escape6868(string(s)))
}

func (s sentBy) String() string {
	return string(s)
}

type timezoneID string

func newTimezoneIDParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	if vs[0].typ != tokenParamQValue {
		return nil, ErrIncorrectParamValueType
	}
	return timezoneID(vs[0].data), nil
}

func (t timezoneID) Bytes() []byte {
	return dquote(escape6868(string(t)))
}

func (t timezoneID) String() string {
	return string(t)
}

const (
	valueBinary value = iota
	valueBoolean
	valueCalAddress
	valueDate
	valueDateTime
	valueDuration
	valueFloat
	valueInteger
	valuePeriod
	valueRecurrence
	valueText
	valueTime
	valueURI
	valueUTCOffset
)

type value int

func newValueParam(vs []token) (attribute, error) {
	if len(vs) != 1 {
		return nil, ErrIncorrectNumParamValues
	}
	switch vs[0].data {
	case "BINARY":
		return valueBinary, nil
	case "BOOLEAN":
		return valueBoolean, nil
	case "CAL-ADDRESS":
		return valueCalAddress, nil
	case "DATE":
		return valueDate, nil
	case "DATE-TIME":
		return valueDateTime, nil
	case "DURATION":
		return valueDuration, nil
	case "FLOAT":
		return valueFloat, nil
	case "INTEGER":
		return valueInteger, nil
	case "PERIOD":
		return valuePeriod, nil
	case "RECUR":
		return valueRecurrence, nil
	case "TEXT":
		return valueText, nil
	case "TIME":
		return valueTime, nil
	case "URI":
		return valueURI, nil
	case "UTC-OFFSET":
		return valueUTCOffset, nil
	default:
		return nil, ErrUnsupportedParamValue
	}
}

func (v value) Bytes() []byte {
	return []byte(v.String())
}

func (v value) String() string {
	switch v {
	case valueBinary:
		return "BINARY"
	case valueBoolean:
		return "BOOLEAN"
	case valueCalAddress:
		return "CAL-ADDRESS"
	case valueDate:
		return "DATE"
	case valueDateTime:
		return "DATE-TIME"
	case valueDuration:
		return "DURATION"
	case valueFloat:
		return "FLOAT"
	case valueInteger:
		return "INTEGER"
	case valuePeriod:
		return "PERIOD"
	case valueRecurrence:
		return "RECUR"
	case valueText:
		return "TEXT"
	case valueTime:
		return "TIME"
	case valueURI:
		return "URI"
	case valueUTCOffset:
		return "UTC-OFFSET"
	default:
		return "unknownvalue"
	}
}

type unknownParam []token

func newUnknownParam(vs []token) (attribute, error) {
	return unknownParam(vs), nil
}

func (u unknownParam) Bytes() []byte {
	var toRet []byte
	for n, uv := range u {
		if n > 0 {
			toRet = append(toRet, ',')
		}
		if uv.typ == tokenParamQValue {
			toRet = append(toRet, dquote(escape6868(uv.data))...)
		} else {
			toRet = append(toRet, escape6868(uv.data)...)
		}
	}
	return toRet
}

func (u unknownParam) String() string {
	return string(u.Bytes())
}

// Errors

var (
	ErrIncorrectNumParamValues = errors.New("incorrect number of param values")
	ErrUnsupportedParamValue   = errors.New("unsupported param value")
	ErrIncorrectParamValueType = errors.New("incorrect param value type")
)
