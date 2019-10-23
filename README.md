# ics
--
    import "vimagination.zapto.org/ics"

Package ics implements an encoder and decoder for iCalendar files

## Usage

```go
const (
	ErrInvalidStructure   errors.Error = "invalid structure"
	ErrMissingAlarmAction errors.Error = "missing alarm action"
	ErrInvalidAlarm       errors.Error = "invalid alarm type"
)
```
Errors

```go
var (
	ErrInvalidParam = errors.New("invalid param value")
	ErrInvalidValue = errors.New("invalid value")
)
```
Errors

```go
var (
	ErrInvalidEnd        = errors.New("invalid end of section")
	ErrMissingRequired   = errors.New("required property missing")
	ErrRequirementNotMet = errors.New("requirement not met")
)
```
Errors

```go
var (
	ErrInvalidContentLine                 = errors.New("invalid content line")
	ErrInvalidContentLineName             = errors.New("invalid content line name")
	ErrInvalidContentLineParamName        = errors.New("invalid content line param name")
	ErrInvalidContentLineQuotedParamValue = errors.New("invalid content line quoted param value")
	ErrInvalidContentLineParamValue       = errors.New("invalid content line param value")
	ErrInvalidContentLineValue            = errors.New("invalid content line value")
)
```
Errors

```go
var (
	ErrInvalidEncoding        = errors.New("invalid Binary encoding")
	ErrInvalidPeriod          = errors.New("invalid Period")
	ErrInvalidDuration        = errors.New("invalid Duration")
	ErrInvalidText            = errors.New("invalid encoded text")
	ErrInvalidBoolean         = errors.New("invalid Boolean")
	ErrInvalidOffset          = errors.New("invalid UTC Offset")
	ErrInvalidRecur           = errors.New("invalid Recur")
	ErrInvalidTime            = errors.New("invalid Time")
	ErrInvalidFloat           = errors.New("invalid float")
	ErrInvalidTFloat          = errors.New("invalid number of floats")
	ErrInvalidPeriodStart     = errors.New("invalid start of Period")
	ErrInvalidPeriodDuration  = errors.New("invalid Period duration")
	ErrInvalidPeriodEnd       = errors.New("invalid end of Period")
	ErrInvalidRecurFrequency  = errors.New("invalid Recur frequency")
	ErrInvalidRecurBySecond   = errors.New("invalid Recur BySecond")
	ErrInvalidRecurByMinute   = errors.New("invalid Recur ByMinute")
	ErrInvalidRecurByHour     = errors.New("invalid Recur ByHour")
	ErrInvalidRecurByDay      = errors.New("invalid Recur ByDay")
	ErrInvalidRecurByMonthDay = errors.New("invalid Recur ByMonthDay")
	ErrInvalidRecurByYearDay  = errors.New("invalid Recur ByYearDay")
	ErrInvalidRecurByWeekNum  = errors.New("invalid Recur ByWeekNum")
	ErrInvalidRecurByMonth    = errors.New("invalid Recur ByMonth")
	ErrInvalidRecurBySetPos   = errors.New("invalid Recur BySetPos")
	ErrInvalidRecurWeekStart  = errors.New("invalid Recur WeekStart")
)
```
Errors

```go
var (
	ErrDuplicateParam = errors.New("duplicate param")
)
```
Errors

```go
var (
	ErrInvalidCalendar = errors.New("invalid calendar")
)
```
Errors

#### func  Encode

```go
func Encode(w io.Writer, cal *Calendar) error
```
Encode encodes the given iCalendar object into the writer. It first validates
the iCalendar object so as not to write invalid data to the writer

#### type Alarm

```go
type Alarm struct {
	AlarmType
}
```

Alarm is the encompassing type for the three alarm types

#### type AlarmAudio

```go
type AlarmAudio struct {
	Trigger       PropTrigger
	Duration      *PropDuration
	Repeat        *PropRepeat
	Attachment    []PropAttachment
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}
```

AlarmAudio provides a group of components that define an Audio Alarm

#### func (AlarmAudio) Type

```go
func (AlarmAudio) Type() string
```
Type returns the type of the alarm "AUDIO"

#### type AlarmDisplay

```go
type AlarmDisplay struct {
	Description   PropDescription
	Trigger       PropTrigger
	Duration      *PropDuration
	Repeat        *PropRepeat
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}
```

AlarmDisplay provides a group of components that define a Display Alarm

#### func (AlarmDisplay) Type

```go
func (AlarmDisplay) Type() string
```
Type returns the type of the alarm "DISPLAY"

#### type AlarmEmail

```go
type AlarmEmail struct {
	Description   PropDescription
	Trigger       PropTrigger
	Summary       PropSummary
	Attendee      *PropAttendee
	Duration      *PropDuration
	Repeat        *PropRepeat
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}
```

AlarmEmail provides a group of components that define an Email Alarm

#### func (AlarmEmail) Type

```go
func (AlarmEmail) Type() string
```
Type returns the type of the alarm "EMAIL"

#### type AlarmNone

```go
type AlarmNone struct{}
```

AlarmNone

#### func (AlarmNone) Type

```go
func (AlarmNone) Type() string
```
Type returns the type of the alarm "NONE"

#### type AlarmType

```go
type AlarmType interface {
	Type() string
	// contains filtered or unexported methods
}
```

AlarmType is an interface this is fulfilled by AlarmAudio, AlarmDisplay and
AlarmEmail

#### type AlarmURI

```go
type AlarmURI struct {
	URI           PropURI
	Duration      *PropDuration
	Repeat        *PropRepeat
	UID           *PropUID
	AlarmAgent    []PropAlarmAgent
	AlarmStatus   *PropAlarmStatus
	LastTriggered []PropLastTriggered
	Acknowledged  *PropAcknowledged
	Proximity     *PropProximity
	GeoLocation   []PropGeoLocation
	RelatedTo     *PropRelatedTo
	DefaultAlarm  *PropDefaultAlarm
}
```

AlarmURI provies a group of components that define a URI Alarm

#### func (AlarmURI) Type

```go
func (AlarmURI) Type() string
```
Type returns the type of the alarm "URI"

#### type Binary

```go
type Binary []byte
```

Binary is inline binary data

#### type Boolean

```go
type Boolean bool
```

Boolean is true or false

#### type Calendar

```go
type Calendar struct {
	Version   PropVersion
	ProductID PropProductID
	Event     []Event
	Todo      []Todo
	Journal   []Journal
	FreeBusy  []FreeBusy
	Timezone  []Timezone
}
```

Calendar represents a iCalendar object

#### func  Decode

```go
func Decode(r io.Reader) (*Calendar, error)
```
Decode decodes an iCalendar object from the given reader

#### type CalendarAddress

```go
type CalendarAddress struct {
	url.URL
}
```

CalendarAddress contains a calendar user address

#### type Date

```go
type Date struct {
	time.Time
}
```

Date is a Calendar Data

#### type DateTime

```go
type DateTime struct {
	time.Time
}
```

DateTime is a Calendar Date and Time

#### type DayRecur

```go
type DayRecur struct {
	Day        WeekDay
	Occurrence int8
}
```

DayRecur is used to reprent the nth day in a time period, be it 2nd Monday in a
Month, or 31st Friday in a year, etc.

#### type Daylight

```go
type Daylight struct {
	DateTimeStart       PropDateTimeStart
	TimezoneOffsetTo    PropTimezoneOffsetTo
	TimezoneOffsetFrom  PropTimezoneOffsetFrom
	RecurrenceRule      *PropRecurrenceRule
	Comment             []PropComment
	RecurrenceDateTimes []PropRecurrenceDateTimes
	TimezoneName        []PropTimezoneName
}
```

Daylight represents daylight savings timezone rules

#### type Duration

```go
type Duration struct {
	Negative                             bool
	Weeks, Days, Hours, Minutes, Seconds uint
}
```

Duration is a duration of time

#### type Event

```go
type Event struct {
	DateTimeStamp       PropDateTimeStamp
	UID                 PropUID
	DateTimeStart       *PropDateTimeStart
	Class               *PropClass
	Created             *PropCreated
	Description         *PropDescription
	Geo                 *PropGeo
	LastModified        *PropLastModified
	Location            *PropLocation
	Organizer           *PropOrganizer
	Priority            *PropPriority
	Sequence            *PropSequence
	Status              *PropStatus
	Summary             *PropSummary
	TimeTransparency    *PropTimeTransparency
	URL                 *PropURL
	RecurrenceID        *PropRecurrenceID
	RecurrenceRule      *PropRecurrenceRule
	DateTimeEnd         *PropDateTimeEnd
	Duration            *PropDuration
	Attachment          []PropAttachment
	Attendee            []PropAttendee
	Categories          []PropCategories
	Comment             []PropComment
	Contact             []PropContact
	ExceptionDateTime   []PropExceptionDateTime
	RequestStatus       []PropRequestStatus
	RelatedTo           []PropRelatedTo
	Resources           []PropResources
	RecurrenceDateTimes []PropRecurrenceDateTimes
	Alarm               []Alarm
}
```

Event provides a group of components that describe an event

#### type Float

```go
type Float float64
```

Float contains a real-number value

#### type FreeBusy

```go
type FreeBusy struct {
	DateTimeStamp PropDateTimeStamp
	UID           PropUID
	Contact       *PropContact
	DateTimeStart *PropDateTimeStart
	DateTimeEnd   *PropDateTimeEnd
	Organizer     *PropOrganizer
	URL           *PropURL
	Attendee      []PropAttendee
	Comment       []PropComment
	FreeBusy      []PropFreeBusy
	RequestStatus []PropRequestStatus
}
```

FreeBusy provides a group of components that describe either a request for
free/busy time, describe a response to a request for free/busy time, or describe
a published set of busy time

#### type Frequency

```go
type Frequency uint8
```

Frequency represents the Recurrence frequency

```go
const (
	Secondly Frequency = iota
	Minutely
	Hourly
	Daily
	Weekly
	Monthly
	Yearly
)
```
Frequency constant values

#### type Integer

```go
type Integer int32
```

Integer is a signed integer value

#### type Journal

```go
type Journal struct {
	DateTimeStamp       PropDateTimeStamp
	UID                 PropUID
	Class               *PropClass
	Created             *PropCreated
	DateTimeStart       *PropDateTimeStart
	LastModified        *PropLastModified
	Organizer           *PropOrganizer
	RecurrenceID        *PropRecurrenceID
	Sequence            *PropSequence
	Status              *PropStatus
	Summary             *PropSummary
	URL                 *PropURL
	RecurrenceRule      *PropRecurrenceRule
	Attachment          []PropAttachment
	Attendee            []PropAttendee
	Categories          []PropCategories
	Comment             []PropComment
	Contact             []PropContact
	Description         []PropDescription
	ExceptionDateTime   []PropExceptionDateTime
	RequestStatus       []PropRequestStatus
	RelatedTo           []PropRelatedTo
	Resources           []PropResources
	RecurrenceDateTimes []PropRecurrenceDateTimes
}
```

Journal provides a group of components that describe a journal entry

#### type MText

```go
type MText []Text
```

MText contains multiple text values

#### type Month

```go
type Month uint8
```

Month is a numeric representation of a Month of the Year

```go
const (
	UnknownMonth Month = iota
	January
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```
Month Constant Values

#### type ParamAgentID

```go
type ParamAgentID string
```

AgentID

#### func  NewAgentID

```go
func NewAgentID(v ParamAgentID) *ParamAgentID
```
NewAgentID returns a *ParamAgentID for ease of use with optional values

#### type ParamAlternativeRepresentation

```go
type ParamAlternativeRepresentation URI
```

AlternativeRepresentation is an alternate text representation for the property
value

#### type ParamCalendarUserType

```go
type ParamCalendarUserType uint8
```

CalendarUserType is identify the type of calendar user specified by the property

```go
const (
	CalendarUserTypeUnknown ParamCalendarUserType = iota
	CalendarUserTypeIndividual
	CalendarUserTypeGroup
	CalendarUserTypeResource
	CalendarUserTypeRoom
)
```
CalendarUserType constant values

#### func (ParamCalendarUserType) New

```go
func (t ParamCalendarUserType) New() *ParamCalendarUserType
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type ParamCommonName

```go
type ParamCommonName string
```

CommonName is the common name to be associated with the calendar user specified
by the property

#### func  NewCommonName

```go
func NewCommonName(v ParamCommonName) *ParamCommonName
```
NewCommonName returns a *ParamCommonName for ease of use with optional values

#### type ParamDelagatee

```go
type ParamDelagatee []CalendarAddress
```

Delagatee is used to specify the calendar users to whom the calendar user
specified by the property has delegated participation

#### type ParamDelegator

```go
type ParamDelegator []CalendarAddress
```

Delegator is used to specify the calendar users that have delegated their
participation to the calendar user specified by the property

#### type ParamDirectoryEntry

```go
type ParamDirectoryEntry string
```

DirectoryEntry is a reference to a directory entry associated with the calendar

#### func  NewDirectoryEntry

```go
func NewDirectoryEntry(v ParamDirectoryEntry) *ParamDirectoryEntry
```
NewDirectoryEntry returns a *ParamDirectoryEntry for ease of use with optional
values

#### type ParamEncoding

```go
type ParamEncoding uint8
```

Encoding is the inline encoding for the property value

```go
const (
	Encoding8bit ParamEncoding = iota
	EncodingBase64
)
```
Encoding constant values

#### func (ParamEncoding) New

```go
func (t ParamEncoding) New() *ParamEncoding
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type ParamFormatType

```go
type ParamFormatType string
```

FormatType is the content type of a referenced object

#### func  NewFormatType

```go
func NewFormatType(v ParamFormatType) *ParamFormatType
```
NewFormatType returns a *ParamFormatType for ease of use with optional values

#### type ParamFreeBusyType

```go
type ParamFreeBusyType uint8
```

FreeBusyType is used to specify the free or busy time type

```go
const (
	FreeBusyTypeUnknown ParamFreeBusyType = iota
	FreeBusyTypeFree
	FreeBusyTypeBusy
	FreeBusyTypeBusyUnavailable
	FreeBusyTypeBusyTentative
)
```
FreeBusyType constant values

#### func (ParamFreeBusyType) New

```go
func (t ParamFreeBusyType) New() *ParamFreeBusyType
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type ParamID

```go
type ParamID string
```

ID

#### func  NewID

```go
func NewID(v ParamID) *ParamID
```
NewID returns a *ParamID for ease of use with optional values

#### type ParamLanguage

```go
type ParamLanguage string
```

Language is the language for text values

#### func  NewLanguage

```go
func NewLanguage(v ParamLanguage) *ParamLanguage
```
NewLanguage returns a *ParamLanguage for ease of use with optional values

#### type ParamMember

```go
type ParamMember []string
```

Member is used to specify the group or list membership of the calendar

#### func  NewMember

```go
func NewMember(v ParamMember) *ParamMember
```
NewMember returns a *ParamMember for ease of use with optional values

#### type ParamParticipationRole

```go
type ParamParticipationRole uint8
```

ParticipationRole is used to specify the participation role for the calendar
user specified by the property

```go
const (
	ParticipationRoleUnknown ParamParticipationRole = iota
	ParticipationRoleRequiredParticipant
	ParticipationRoleChair
	ParticipationRoleOptParticipant
	ParticipationRoleNonParticipant
)
```
ParticipationRole constant values

#### func (ParamParticipationRole) New

```go
func (t ParamParticipationRole) New() *ParamParticipationRole
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type ParamParticipationStatus

```go
type ParamParticipationStatus uint8
```

ParticipationStatus is used to specify the participation status for the calendar

```go
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
```
ParticipationStatus constant values

#### func (ParamParticipationStatus) New

```go
func (t ParamParticipationStatus) New() *ParamParticipationStatus
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type ParamRSVP

```go
type ParamRSVP Boolean
```

RSVP is used to specify whether there is an expectation of a favor of a reply
from the calendar user specified by the property value

#### func  NewRSVP

```go
func NewRSVP(v ParamRSVP) *ParamRSVP
```
NewRSVP returns a *ParamRSVP for ease of use with optional values

#### type ParamRange

```go
type ParamRange struct{}
```

Range is used to specify the effective range of recurrence instances from the
instance specified by the recurrence identifier specified by the property

#### type ParamRelated

```go
type ParamRelated uint8
```

Related is the relationship of the alarm trigger with respect to the start or
end of the calendar component

```go
const (
	RelatedStart ParamRelated = iota
	RelatedEnd
)
```
Related constant values

#### func (ParamRelated) New

```go
func (t ParamRelated) New() *ParamRelated
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type ParamRelationshipType

```go
type ParamRelationshipType uint8
```

RelationshipType is the type of hierarchical relationship associated with the
calendar component specified by the property

```go
const (
	RelationshipTypeUnknown ParamRelationshipType = iota
	RelationshipTypeParent
	RelationshipTypeChild
	RelationshipTypeSibling
)
```
RelationshipType constant values

#### func (ParamRelationshipType) New

```go
func (t ParamRelationshipType) New() *ParamRelationshipType
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type ParamSentBy

```go
type ParamSentBy string
```

SentBy is used to specify the calendar user that is acting on behalf of the
calendar user specified by the property

#### func  NewSentBy

```go
func NewSentBy(v ParamSentBy) *ParamSentBy
```
NewSentBy returns a *ParamSentBy for ease of use with optional values

#### type ParamTimezoneID

```go
type ParamTimezoneID string
```

TimezoneID is used to specify the identifier for the time zone definition for a
time component in the property value

#### func  NewTimezoneID

```go
func NewTimezoneID(v ParamTimezoneID) *ParamTimezoneID
```
NewTimezoneID returns a *ParamTimezoneID for ease of use with optional values

#### type ParamURI

```go
type ParamURI URI
```

URI

#### type ParamValue

```go
type ParamValue uint8
```

Value is used to explicitly specify the value type format for a property value

```go
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
```
Value constant values

#### func (ParamValue) New

```go
func (t ParamValue) New() *ParamValue
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type Period

```go
type Period struct {
	Start, End DateTime
	Duration   Duration
}
```

Period represents a precise period of time/

Only one of End or Duration will be used. If Period.End.IsZero() is true, then
it uses Period.Duration

#### type PropAcknowledged

```go
type PropAcknowledged DateTime
```

PropAcknowledged

#### type PropAction

```go
type PropAction uint8
```

PropAction defines the action to be invoked when an alarm is triggered

```go
const (
	ActionAudio PropAction = iota
	ActionDisplay
	ActionEmail
)
```
PropAction constant values

#### func (PropAction) New

```go
func (p PropAction) New() *PropAction
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type PropAlarmAgent

```go
type PropAlarmAgent struct {
	URI     *ParamURI
	ID      *ParamID
	AgentID *ParamAgentID
	Text
}
```

PropAlarmAgent

#### type PropAlarmStatus

```go
type PropAlarmStatus uint8
```

PropAlarmStatus

```go
const (
	AlarmStatusActive PropAlarmStatus = iota
	AlarmStatusCancelled
	AlarmStatusCompleted
)
```
PropAlarmStatus constant values

#### func (PropAlarmStatus) New

```go
func (p PropAlarmStatus) New() *PropAlarmStatus
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type PropAttachment

```go
type PropAttachment struct {
	FormatType *ParamFormatType
	URI        *URI
	Binary     *Binary
}
```

PropAttachment provides the capability to associate a document object with a
calendar component

#### type PropAttendee

```go
type PropAttendee struct {
	CalendarUserType    *ParamCalendarUserType
	Member              ParamMember
	ParticipationRole   *ParamParticipationRole
	ParticipationStatus *ParamParticipationStatus
	RSVP                *ParamRSVP
	Delagatee           ParamDelagatee
	Delegator           ParamDelegator
	SentBy              *ParamSentBy
	CommonName          *ParamCommonName
	DirectoryEntry      *ParamDirectoryEntry
	Language            *ParamLanguage
	CalendarAddress
}
```

PropAttendee defines an "Attendee" within a calendar component

#### type PropCalendarScale

```go
type PropCalendarScale uint8
```

PropCalendarScale defines the calendar scale

```go
const (
	CalendarScaleGregorian PropCalendarScale = iota
)
```
PropCalendarScale constant values

#### func (PropCalendarScale) New

```go
func (p PropCalendarScale) New() *PropCalendarScale
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type PropCategories

```go
type PropCategories struct {
	Language *ParamLanguage
	MText
}
```

PropCategories defines the categories for a calendar component

#### type PropClass

```go
type PropClass uint8
```

PropClass defines the access classification for a calendar component

```go
const (
	ClassPublic PropClass = iota
	ClassPrivate
	ClassConfidential
)
```
PropClass constant values

#### func (PropClass) New

```go
func (p PropClass) New() *PropClass
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type PropComment

```go
type PropComment struct {
	AlternativeRepresentation *ParamAlternativeRepresentation
	Language                  *ParamLanguage
	Text
}
```

PropComment specifies non-processing information intended to provide a comment
to the calendar user

#### type PropCompleted

```go
type PropCompleted DateTime
```

PropCompleted defines the date and time that a to-do was actually completed

#### type PropContact

```go
type PropContact struct {
	AlternativeRepresentation *ParamAlternativeRepresentation
	Language                  *ParamLanguage
	Text
}
```

PropContact is used to represent contact information or alternately a reference
to contact information associated with the calendar component

#### type PropCreated

```go
type PropCreated DateTime
```

PropCreated specifies the date and time that the calendar information was
created by the calendar user agent in the calendar store

#### type PropDateTimeEnd

```go
type PropDateTimeEnd struct {
	DateTime *DateTime
	Date     *Date
}
```

PropDateTimeEnd specifies the date and time that a calendar component ends

#### type PropDateTimeStamp

```go
type PropDateTimeStamp DateTime
```

PropDateTimeStamp specifies the date and time that the calendar object was
created unless the calendar object has no METHOD property, in which case it
specifies the date and time that the information with the calendar was last
revised

#### type PropDateTimeStart

```go
type PropDateTimeStart struct {
	DateTime *DateTime
	Date     *Date
}
```

PropDateTimeStart specifies when the calendar component begins

#### type PropDefaultAlarm

```go
type PropDefaultAlarm Boolean
```

PropDefaultAlarm

#### type PropDescription

```go
type PropDescription struct {
	AlternativeRepresentation *ParamAlternativeRepresentation
	Language                  *ParamLanguage
	Text
}
```

PropDescription provides a more complete description of the calendar component
than that provided by the "SUMMARY" property

#### type PropDue

```go
type PropDue struct {
	DateTime *DateTime
	Date     *Date
}
```

PropDue defines the date and time that a to-do is expected to be completed

#### type PropDuration

```go
type PropDuration Duration
```

PropDuration specifies a positive duration of time

#### type PropExceptionDateTime

```go
type PropExceptionDateTime struct {
	DateTime *DateTime
	Date     *Date
}
```

PropExceptionDateTime defines the list of DATE-TIME exceptions for recurring
events, to-dos, journal entries, or time zone definitions

#### type PropFreeBusy

```go
type PropFreeBusy struct {
	FreeBusyType *ParamFreeBusyType
	Period
}
```

PropFreeBusy defines one or more free or busy time intervals

#### type PropGeo

```go
type PropGeo TFloat
```

PropGeo specifies information related to the global position for the activity
specified by a calendar component

#### type PropGeoLocation

```go
type PropGeoLocation URI
```

PropGeoLocation

#### type PropLastModified

```go
type PropLastModified DateTime
```

PropLastModified specifies the date and time that the information associated
with the calendar component was last revised in the calendar store

#### type PropLastTriggered

```go
type PropLastTriggered DateTime
```

PropLastTriggered

#### type PropLocation

```go
type PropLocation struct {
	AlternativeRepresentation *ParamAlternativeRepresentation
	Language                  *ParamLanguage
	Text
}
```

PropLocation defines the intended venue for the activity defined by a calendar
component

#### type PropMethod

```go
type PropMethod Text
```

PropMethod defines the iCalendar object method associated with the calendar
object

#### type PropOrganizer

```go
type PropOrganizer struct {
	CommonName     *ParamCommonName
	DirectoryEntry *ParamDirectoryEntry
	SentBy         *ParamSentBy
	Language       *ParamLanguage
	CalendarAddress
}
```

PropOrganizer defines the organizer for a calendar component

#### type PropPercentComplete

```go
type PropPercentComplete Integer
```

PropPercentComplete is used by an assignee or delegatee of a to-do to convey the
percent completion of a to-do to the "Organizer"

#### func  NewPercentComplete

```go
func NewPercentComplete(v PropPercentComplete) *PropPercentComplete
```
NewPercentComplete generates a pointer to a constant value. Used when manually
creating Calendar values

#### type PropPriority

```go
type PropPriority Integer
```

PropPriority defines the relative priority for a calendar component

#### func  NewPriority

```go
func NewPriority(v PropPriority) *PropPriority
```
NewPriority generates a pointer to a constant value. Used when manually creating
Calendar values

#### type PropProductID

```go
type PropProductID Text
```

PropProductID specifies the identifier for the product that created the
iCalendar object

#### type PropProximity

```go
type PropProximity Text
```

PropProximity

#### type PropRecurrenceDateTimes

```go
type PropRecurrenceDateTimes struct {
	DateTime *DateTime
	Date     *Date
	Period   *Period
}
```

PropRecurrenceDateTimes defines the list of DATE-TIME values for recurring
events, to-dos, journal entries, or time zone definitions

#### type PropRecurrenceID

```go
type PropRecurrenceID struct {
	Range    *ParamRange
	DateTime *DateTime
	Date     *Date
}
```

PropRecurrenceID is used to identify a specific instance of a recurring Event,
Todo or Journal

#### type PropRecurrenceRule

```go
type PropRecurrenceRule Recur
```

PropRecurrenceRule defines a rule or repeating pattern for recurring events,
to-dos, journal entries, or time zone definitions

#### type PropRelatedTo

```go
type PropRelatedTo struct {
	RelationshipType *ParamRelationshipType
	Text
}
```

PropRelatedTo is used to represent a relationship or reference between one
calendar component and another

#### type PropRepeat

```go
type PropRepeat Integer
```

PropRepeat defines the number of times the alarm should be repeated, after the
initial trigger

#### func  NewRepeat

```go
func NewRepeat(v PropRepeat) *PropRepeat
```
NewRepeat generates a pointer to a constant value. Used when manually creating
Calendar values

#### type PropRequestStatus

```go
type PropRequestStatus Text
```

PropRequestStatus defines the status code returned for a scheduling request

#### type PropResources

```go
type PropResources struct {
	AlternativeRepresentation *ParamAlternativeRepresentation
	Language                  *ParamLanguage
	MText
}
```

PropResources defines the equipment or resources anticipated for an activity
specified by a calendar component

#### type PropSequence

```go
type PropSequence Integer
```

PropSequence defines the revision sequence number of the calendar component
within a sequence of revisions

#### func  NewSequence

```go
func NewSequence(v PropSequence) *PropSequence
```
NewSequence generates a pointer to a constant value. Used when manually creating
Calendar values

#### type PropStatus

```go
type PropStatus uint8
```

PropStatus defines the overall status or confirmation for the calendar component

```go
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
```
PropStatus constant values

#### func (PropStatus) New

```go
func (p PropStatus) New() *PropStatus
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type PropSummary

```go
type PropSummary struct {
	AlternativeRepresentation *ParamAlternativeRepresentation
	Language                  *ParamLanguage
	Text
}
```

PropSummary defines a short summary or subject for the calendar component

#### type PropTimeTransparency

```go
type PropTimeTransparency uint8
```

PropTimeTransparency defines whether or not an event is transparent to busy time
searches

```go
const (
	TimeTransparencyOpaque PropTimeTransparency = iota
	TimeTransparencyTransparent
)
```
PropTimeTransparency constant values

#### func (PropTimeTransparency) New

```go
func (p PropTimeTransparency) New() *PropTimeTransparency
```
New returns a pointer to the type (used with constants for ease of use with
optional values)

#### type PropTimezoneID

```go
type PropTimezoneID Text
```

PropTimezoneID specifies the text value that uniquely identifies the "VTIMEZONE"
calendar component in the scope of an iCalendar object

#### type PropTimezoneName

```go
type PropTimezoneName struct {
	Language *ParamLanguage
	Text
}
```

PropTimezoneName specifies the customary designation for a time zone description

#### type PropTimezoneOffsetFrom

```go
type PropTimezoneOffsetFrom UTCOffset
```

PropTimezoneOffsetFrom specifies the offset that is in use prior to this time
zone observance

#### type PropTimezoneOffsetTo

```go
type PropTimezoneOffsetTo UTCOffset
```

PropTimezoneOffsetTo specifies the offset that is in use in this time zone
observance

#### type PropTimezoneURL

```go
type PropTimezoneURL URI
```

PropTimezoneURL provides a means for a "VTIMEZONE" component to point to a
network location that can be used to retrieve an up- to-date version of itself

#### type PropTrigger

```go
type PropTrigger struct {
	Duration *Duration
	DateTime *DateTime
}
```

PropTrigger specifies when an alarm will trigger

#### type PropUID

```go
type PropUID Text
```

PropUID defines the persistent, globally unique identifier for the calendar
component

#### type PropURI

```go
type PropURI URI
```

PropURI

#### type PropURL

```go
type PropURL URI
```

PropURL defines a Uniform Resource Locator associated with the iCalendar object

#### type PropVersion

```go
type PropVersion Text
```

PropVersion specifies the identifier corresponding to the highest version number
or the minimum and maximum range of the iCalendar specification that is required
in order to interpret the iCalendar object

#### type Recur

```go
type Recur struct {
	Frequency  Frequency
	WeekStart  WeekDay
	UntilTime  bool
	Until      time.Time
	Count      uint64
	Interval   uint64
	BySecond   []uint8
	ByMinute   []uint8
	ByHour     []uint8
	ByDay      []DayRecur
	ByMonthDay []int8
	ByYearDay  []int16
	ByWeekNum  []int8
	ByMonth    []Month
	BySetPos   []int16
}
```

Recur contains a recurrence rule specification

#### type Standard

```go
type Standard struct {
	DateTimeStart       PropDateTimeStart
	TimezoneOffsetTo    PropTimezoneOffsetTo
	TimezoneOffsetFrom  PropTimezoneOffsetFrom
	RecurrenceRule      *PropRecurrenceRule
	Comment             []PropComment
	RecurrenceDateTimes []PropRecurrenceDateTimes
	TimezoneName        []PropTimezoneName
}
```

Standard represents standard timezone rules

#### type TFloat

```go
type TFloat [2]float64
```

TFloat is a pair of float values used for coords

#### type Text

```go
type Text string
```

Text contains human-readable text

#### type Time

```go
type Time struct {
	time.Time
}
```

Time contains a precise time

#### type Timezone

```go
type Timezone struct {
	TimezoneID   PropTimezoneID
	LastModified *PropLastModified
	TimezoneURL  *PropTimezoneURL
	Standard     []Standard
	Daylight     []Daylight
}
```

Timezone provide a group of components that defines a time zone

#### type Todo

```go
type Todo struct {
	DateTimeStamp       PropDateTimeStamp
	UID                 PropUID
	Class               *PropClass
	Completed           *PropCompleted
	Created             *PropCreated
	Description         *PropDescription
	DateTimeStart       *PropDateTimeStart
	Geo                 *PropGeo
	LastModified        *PropLastModified
	Location            *PropLocation
	Organizer           *PropOrganizer
	PercentComplete     *PropPercentComplete
	Priority            *PropPriority
	RecurrenceID        *PropRecurrenceID
	Sequence            *PropSequence
	Status              *PropStatus
	Summary             *PropSummary
	URL                 *PropURL
	Due                 *PropDue
	Duration            *PropDuration
	Attachment          []PropAttachment
	Attendee            []PropAttendee
	Categories          []PropCategories
	Comment             []PropComment
	Contact             []PropContact
	ExceptionDateTime   []PropExceptionDateTime
	RequestStatus       []PropRequestStatus
	RelatedTo           []PropRelatedTo
	Resources           []PropResources
	RecurrenceDateTimes []PropRecurrenceDateTimes
	Alarm               []Alarm
}
```

Todo provides a group of components that describe a to-do

#### type URI

```go
type URI struct {
	url.URL
}
```

URI contains a reference to another piece of data

#### type UTCOffset

```go
type UTCOffset int
```

UTCOffset contains the offset from UTC to local time

#### type WeekDay

```go
type WeekDay uint8
```

WeekDay is a numeric representation of a Day of the Week

```go
const (
	UnknownDay WeekDay = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```
Weekday constant values
