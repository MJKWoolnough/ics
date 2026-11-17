package ics

import (
	"errors"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"vimagination.zapto.org/parser"
)

type params interface {
	decode([]parser.Token) error
	encode(writer)
}

type paramTests[T params] struct {
	Input  []parser.Token
	Param  T
	Error  error
	Output string
}

func testParams[T params](t *testing.T, tests []paramTests[T]) {
	for n, test := range tests {
		v := reflect.New(reflect.TypeFor[T]().Elem()).Interface().(params)

		if err := v.decode(test.Input); !errors.Is(err, test.Error) {
			if errv := reflect.New(reflect.TypeOf(test.Error)); !errors.As(err, errv.Interface()) || !reflect.DeepEqual(test.Error, errv.Elem().Interface().(error)) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Error, err)

				continue
			}
		}

		if test.Error == nil {
			if !reflect.DeepEqual(v, test.Param) {
				t.Errorf("test %d: expecting param %#v, got %#v", n+1, test.Param, v)
			} else if test.Output != "" {
				var buf strings.Builder

				v.encode(&buf)

				if out := buf.String(); out != test.Output {
					t.Errorf("test %d: expecting output %q, got %q", n+1, test.Output, out)
				}
			}
		}
	}
}

func testStrings[T ~string, PV interface {
	*T
	params
}](t *testing.T, prefix string, req parser.TokenType) {
	expect := prefix + "abc"

	if req == tokenParamQuotedValue {
		expect = prefix + "\"abc\""
	}

	tests := []paramTests[PV]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: req, Data: ""},
			},
			Param: new(T("")),
		},
		{ // 3
			Input: []parser.Token{
				{Type: req, Data: "abc"},
				{Type: req, Data: "def"},
			},
			Error: ErrInvalidParam,
		},
		{ // 4
			Input: []parser.Token{
				{Type: req, Data: "abc"},
			},
			Param:  new(T("abc")),
			Output: expect,
		},
		{ // 5
			Input: []parser.Token{
				{Type: req, Data: "abc;"},
			},
			Param:  new(T("abc;")),
			Output: prefix + "\"abc;\"",
		},
	}

	if req == tokenParamQuotedValue {
		tests = append(tests, paramTests[PV]{
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		})
	}

	testParams(t, tests)
}

func testOptions[T params](t *testing.T, outputPrefix string, options map[string]T) {
	tests := []paramTests[T]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{{}, {}},
			Error: ErrInvalidParam,
		},
	}

	for opt, v := range options {
		for i := range 1 << len(opt) {
			var sb strings.Builder

			for n, c := range opt {
				if c < 'A' || c > 'Z' {
					sb.WriteByte(byte(c))

					continue
				}

				if i&(1<<n) == 0 {
					sb.WriteByte(byte(c) | 0x20)
				} else {
					sb.WriteByte(byte(c) & 0xDF)
				}
			}

			tests = append(tests, paramTests[T]{
				Input:  []parser.Token{{Data: sb.String()}},
				Param:  v,
				Output: outputPrefix + string(opt),
			})
		}
	}

	if _, ok := options["UNKNOWN"]; !ok {
		tests = append(tests, paramTests[T]{
			Input: []parser.Token{{Data: "UnKnOwN"}},
			Error: ErrInvalidParam,
		})
	}

	testParams(t, tests)
}

func TestParamAlternativeRepresentation(t *testing.T) {
	u, _ := url.Parse("http://www.example.com")

	testParams(t, []paramTests[*ParamAlternativeRepresentation]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamName, Data: "http://www.example.com"},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
			},
			Param:  &ParamAlternativeRepresentation{*u},
			Output: ";ALTREP=http://www.example.com",
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamQuotedValue, Data: "http://www.example2.com"},
			},
			Error: ErrInvalidParam,
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://a%:b@example.com"},
			},
			Error: url.EscapeError("%"),
		},
	})
}

func TestParamCommonName(t *testing.T) {
	testStrings[ParamCommonName](t, ";CN=", tokenParamValue)
}

func TestParamCalendarUserType(t *testing.T) {
	testOptions(t, ";CUTYPE=", map[string]*ParamCalendarUserType{
		"INDIVIDUAL": new(CalendarUserTypeIndividual),
		"GROUP":      new(CalendarUserTypeGroup),
		"RESOURCE":   new(CalendarUserTypeResource),
		"ROOM":       new(CalendarUserTypeRoom),
		"UNKNOWN":    new(CalendarUserTypeUnknown),
	})
}

func TestParamDelegator(t *testing.T) {
	u1, _ := url.Parse("http://www.example.com")
	u2, _ := url.Parse("http://www.another.com")

	testParams(t, []paramTests[*ParamDelegator]{
		{ // 1
			Input: []parser.Token{},
			Param: new(ParamDelegator),
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
			},
			Param: &ParamDelegator{
				{*u1},
			},
			Output: ";DELEGATED-FROM=http://www.example.com",
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://a%:@www.example.com"},
			},
			Error: url.EscapeError("%"),
		},
		{ // 6
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamQuotedValue, Data: "http://www.another.com"},
			},
			Param: &ParamDelegator{
				{*u1},
				{*u2},
			},
			Output: ";DELEGATED-FROM=http://www.example.com,http://www.another.com",
		},
	})
}

func TestParamDelagatee(t *testing.T) {
	u1, _ := url.Parse("http://www.example.com")
	u2, _ := url.Parse("http://www.another.com")

	testParams(t, []paramTests[*ParamDelagatee]{
		{ // 1
			Input: []parser.Token{},
			Param: new(ParamDelagatee),
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
			},
			Param: &ParamDelagatee{
				{*u1},
			},
			Output: ";DELEGATED-TO=http://www.example.com",
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://a%:@www.example.com"},
			},
			Error: url.EscapeError("%"),
		},
		{ // 6
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "http://www.example.com"},
				{Type: tokenParamQuotedValue, Data: "http://www.another.com"},
			},
			Param: &ParamDelagatee{
				{*u1},
				{*u2},
			},
			Output: ";DELEGATED-TO=http://www.example.com,http://www.another.com",
		},
	})
}

func TestParamDirectoryEntry(t *testing.T) {
	testParams(t, []paramTests[*ParamDirectoryEntry]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: ""},
			},
			Param: NewDirectoryEntry(""),
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "abc"},
				{Type: tokenParamQuotedValue, Data: "def"},
			},
			Error: ErrInvalidParam,
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "abc"},
			},
			Param:  NewDirectoryEntry("abc"),
			Output: ";DIR=\"abc\"",
		},
	})
}

func TestParamEncoding(t *testing.T) {
	testOptions(t, ";ENCODING=", map[string]*ParamEncoding{
		"8BIT":   new(Encoding8bit),
		"BASE64": new(EncodingBase64),
	})
}

func TestParamFormatType(t *testing.T) {
	testParams(t, []paramTests[*ParamFormatType]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc/123"},
				{Type: tokenParamValue, Data: "def/456"},
			},
			Error: ErrInvalidParam,
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc/123"},
			},
			Param:  NewFormatType("abc/123"),
			Output: ";FMTTYPE=abc/123",
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "abc;/123"},
			},
			Param:  NewFormatType("abc;/123"),
			Output: ";FMTTYPE=\"abc;/123\"",
		},
	})
}

func TestParamFreeBusyType(t *testing.T) {
	testOptions(t, ";FBTYPE=", map[string]*ParamFreeBusyType{
		"FREE":             new(FreeBusyTypeFree),
		"BUSY":             new(FreeBusyTypeBusy),
		"BUSY-UNAVAILABLE": new(FreeBusyTypeBusyUnavailable),
		"BUSY-TENTATIVE":   new(FreeBusyTypeBusyTentative),
		"UNKNOWN":          new(FreeBusyTypeUnknown),
	})
}

func TestParamLanguage(t *testing.T) {
	testStrings[ParamLanguage](t, ";LANGUAGE=", tokenParamValue)
}

func TestParamMember(t *testing.T) {
	testParams(t, []paramTests[*ParamMember]{
		{ // 1
			Input: []parser.Token{},
			Param: new(ParamMember),
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "abc"},
				{Type: tokenParamValue, Data: ""},
			},
			Error: ErrInvalidParam,
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "abc"},
			},
			Param:  NewMember(ParamMember{"abc"}),
			Output: ";MEMBER=\"abc\"",
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "abc;"},
			},
			Param:  NewMember(ParamMember{"abc;"}),
			Output: ";MEMBER=\"abc;\"",
		},
		{ // 6
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "abc"},
				{Type: tokenParamQuotedValue, Data: "def"},
			},
			Param:  NewMember(ParamMember{"abc", "def"}),
			Output: ";MEMBER=\"abc\",\"def\"",
		},
	})
}

func TestParamParticipationStatus(t *testing.T) {
	testOptions(t, ";PARTSTAT=", map[string]*ParamParticipationStatus{
		"NEEDS-ACTION": new(ParticipationStatusNeedsAction),
		"ACCEPTED":     new(ParticipationStatusAccepted),
		"DECLINED":     new(ParticipationStatusDeclined),
		"TENTATIVE":    new(ParticipationStatusTentative),
		"DELEGATED":    new(ParticipationStatusDelegated),
		"COMPLETED":    new(ParticipationStatusCompleted),
		"IN-PROCESS":   new(ParticipationStatusInProcess),
		"UNKNOWN":      new(ParticipationStatusUnknown),
	})
}

func TestParamRange(t *testing.T) {
	testParams(t, []paramTests[*ParamRange]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "abc"},
			},
			Error: ErrInvalidParam,
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "THISANDFUTURE"},
			},
			Param:  new(ParamRange),
			Output: ";RANGE=THISANDFUTURE",
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamQuotedValue, Data: "THISANDFUTURE"},
				{Type: tokenParamQuotedValue, Data: "THISANDFUTURE"},
			},
			Error: ErrInvalidParam,
		},
		{ // 5
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "THISANDFUTURE"},
			},
			Param:  new(ParamRange),
			Output: ";RANGE=THISANDFUTURE",
		},
	})
}

func TestParamRelated(t *testing.T) {
	testOptions(t, ";RELATED=", map[string]*ParamRelated{
		"START": new(RelatedStart),
		"END":   new(RelatedEnd),
	})
}

func TestParamRelationshipType(t *testing.T) {
	testOptions(t, ";RELTYPE=", map[string]*ParamRelationshipType{
		"PARENT":  new(RelationshipTypeParent),
		"CHILD":   new(RelationshipTypeChild),
		"SIBLING": new(RelationshipTypeSibling),
		"UNKNOWN": new(RelationshipTypeUnknown),
	})
}

func TestParamParticipationRole(t *testing.T) {
	testOptions(t, ";ROLE=", map[string]*ParamParticipationRole{
		"REQ-PARTICIPANT": new(ParticipationRoleRequiredParticipant),
		"CHAIR":           new(ParticipationRoleChair),
		"OPT-PARTICIPANT": new(ParticipationRoleOptParticipant),
		"NON-PARTICIPANT": new(ParticipationRoleNonParticipant),
		"UNKNOWN":         new(ParticipationRoleUnknown),
	})
}

func TestParamRSVP(t *testing.T) {
	testParams(t, []paramTests[*ParamRSVP]{
		{ // 1
			Input: []parser.Token{},
			Error: ErrInvalidParam,
		},
		{ // 2
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "TRUE"},
			},
			Param:  NewRSVP(true),
			Output: ";RSVP=TRUE",
		},
		{ // 3
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "FALSE"},
			},
			Param: NewRSVP(false),
		},
		{ // 4
			Input: []parser.Token{
				{Type: tokenParamValue, Data: "TRUE"},
				{Type: tokenParamValue, Data: "FALSE"},
			},
			Error: ErrInvalidParam,
		},
	})
}

func TestParamSentBy(t *testing.T) {
	testStrings[ParamSentBy](t, ";SENT-BY=", tokenParamQuotedValue)
}

func TestParamTimezoneID(t *testing.T) {
	testStrings[ParamTimezoneID](t, ";TZID=", tokenParamValue)
}

func TestParamValue(t *testing.T) {
	testOptions(t, ";VALUE=", map[string]*ParamValue{
		"BINARY":      new(ValueBinary),
		"BOOLEAN":     new(ValueBoolean),
		"CAL-ADDRESS": new(ValueCalendarAddress),
		"DATE":        new(ValueDate),
		"DATE-TIME":   new(ValueDateTime),
		"DURATION":    new(ValueDuration),
		"FLOAT":       new(ValueFloat),
		"INTEGER":     new(ValueInteger),
		"PERIOD":      new(ValuePeriod),
		"RECUR":       new(ValueRecur),
		"TEXT":        new(ValueText),
		"TIME":        new(ValueTime),
		"URI":         new(ValueURI),
		"UTC-OFFSET":  new(ValueUTCOffset),
		"UNKNOWN":     new(ValueUnknown),
	})
}
