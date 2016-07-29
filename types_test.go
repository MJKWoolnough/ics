package ics

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"
)

type typeTest struct {
	Params map[string]string
	Data   string
	Input  interface {
		Decode(map[string]string, string) error
	}
	Match interface {
		Encode(io.Writer)
	}
	Output string
	Error  error
}

var emptyMap = make(map[string]string)

func testType(t *testing.T, tests []typeTest) {
	var buf bytes.Buffer
	for n, test := range tests {
		if test.Params == nil {
			test.Params = emptyMap
		}
		err := test.Input.Decode(test.Params, test.Data)
		if err != test.Error {
			if !reflect.DeepEqual(err, test.Error) {
				t.Errorf("test %d: expecting error %s, got %s", n+1, test.Error, err)
				continue
			}
		}
		if test.Error == nil {
			if !reflect.DeepEqual(test.Input, test.Match) {
				t.Errorf("test %d: input does not match expected", n+1)
				continue
			}
			test.Match.Encode(&buf)
			if str := buf.String(); str != test.Output {
				t.Errorf("test %d: expecting output string %q, got %q", n+1, test.Output, str)
			}
			buf.Reset()
		}
	}
}

func TestBinary(t *testing.T) {
	testType(t, []typeTest{
		{
			Input: &Binary{},
			Match: &Binary{},
			Error: ErrInvalidEncoding,
		},
		{
			Params: map[string]string{"ENCODING": "BASE64"},
			Data:   "MTIzNDU=",
			Input:  &Binary{},
			Match:  &Binary{'1', '2', '3', '4', '5'},
			Output: "MTIzNDU=",
		},
	})
}

func TestBoolean(t *testing.T) {
	tr := new(Boolean)
	fa := new(Boolean)
	*tr = true
	testType(t, []typeTest{
		{
			Data:   "False",
			Input:  fa,
			Match:  fa,
			Output: "FALSE",
		},
		{
			Data:   "true",
			Input:  tr,
			Match:  tr,
			Output: "TRUE",
		},
		{
			Data:  "HotDog",
			Input: fa,
			Error: ErrInvalidBoolean,
		},
	})
}

func TestDate(t *testing.T) {
	testType(t, []typeTest{
		{
			Data:   "20011225",
			Input:  &Date{},
			Match:  &Date{time.Date(2001, 12, 25, 0, 0, 0, 0, time.UTC)},
			Output: "20011225",
		},
		{
			Data:   "20081111",
			Input:  &Date{},
			Match:  &Date{time.Date(2008, 11, 11, 0, 0, 0, 0, time.UTC)},
			Output: "20081111",
		},
	})
}

func TestDateTime(t *testing.T) {
	l, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	testType(t, []typeTest{
		{
			Data:   "20011225T131415",
			Input:  &DateTime{},
			Match:  &DateTime{time.Date(2001, 12, 25, 13, 14, 15, 0, time.Local)},
			Output: "20011225T131415",
		},
		{
			Data:   "20011225T131415Z",
			Input:  &DateTime{},
			Match:  &DateTime{time.Date(2001, 12, 25, 13, 14, 15, 0, time.UTC)},
			Output: "20011225T131415Z",
		},
		{
			Params: map[string]string{"TZID": "America/New_York"},
			Data:   "20011225T131415",
			Input:  &DateTime{},
			Match:  &DateTime{time.Date(2001, 12, 25, 13, 14, 15, 0, l)},
			Output: "20011225T131415",
		},
	})
}

func TestDuration(t *testing.T) {
	testType(t, []typeTest{
		{
			Data:   "P1W",
			Input:  &Duration{},
			Match:  &Duration{Weeks: 1},
			Output: "P1W",
		},
		{
			Data:   "P1D",
			Input:  &Duration{},
			Match:  &Duration{Days: 1},
			Output: "P1D",
		},
		{
			Data:   "PT1H",
			Input:  &Duration{},
			Match:  &Duration{Hours: 1},
			Output: "PT1H",
		},
		{
			Data:   "PT1M",
			Input:  &Duration{},
			Match:  &Duration{Minutes: 1},
			Output: "PT1M",
		},
		{
			Data:   "PT1S",
			Input:  &Duration{},
			Match:  &Duration{Seconds: 1},
			Output: "PT1S",
		},
		{
			Data:   "P25W",
			Input:  &Duration{},
			Match:  &Duration{Weeks: 25},
			Output: "P25W",
		},
		{
			Data:  "P25W1D",
			Input: &Duration{},
			Error: ErrInvalidDuration,
		},
		{
			Data:  "P25G",
			Input: &Duration{},
			Error: ErrInvalidDuration,
		},
		{
			Data:   "P1DT2H3M4S",
			Input:  &Duration{},
			Match:  &Duration{Days: 1, Hours: 2, Minutes: 3, Seconds: 4},
			Output: "P1DT2H3M4S",
		},
		{
			Data:   "PT0S",
			Input:  &Duration{},
			Match:  &Duration{},
			Output: "PT0S",
		},
		{
			Data:   "+P1DT2H3M4S",
			Input:  &Duration{},
			Match:  &Duration{Days: 1, Hours: 2, Minutes: 3, Seconds: 4},
			Output: "P1DT2H3M4S",
		},
		{
			Data:   "-P1DT2H3M4S",
			Input:  &Duration{},
			Match:  &Duration{Negative: true, Days: 1, Hours: 2, Minutes: 3, Seconds: 4},
			Output: "-P1DT2H3M4S",
		},
	})
}

func TestPeriod(t *testing.T) {
	testType(t, []typeTest{
		{
			Data:  "19970101T180000Z/19970102T070000Z",
			Input: &Period{},
			Match: &Period{
				Start: DateTime{Time: time.Date(1997, 1, 1, 18, 0, 0, 0, time.UTC)},
				End:   DateTime{Time: time.Date(1997, 1, 2, 7, 0, 0, 0, time.UTC)},
			},
			Output: "19970101T180000Z/19970102T070000Z",
		},
		{
			Data:  "19970101T180000Z/PT5H30M",
			Input: &Period{},
			Match: &Period{
				Start:    DateTime{Time: time.Date(1997, 1, 1, 18, 0, 0, 0, time.UTC)},
				Duration: Duration{Hours: 5, Minutes: 30},
			},
			Output: "19970101T180000Z/PT5H30M",
		},
	})
}

func TestRecur(t *testing.T) {
	testType(t, []typeTest{
		{
			Data:   "FREQ=SECONDLY",
			Input:  &Recur{},
			Match:  &Recur{},
			Output: "FREQ=SECONDLY",
		},
		{
			Data:  "FREQ=MINUTELY",
			Input: &Recur{},
			Match: &Recur{
				Frequency: Minutely,
			},
			Output: "FREQ=MINUTELY",
		},
		{
			Data:  "FREQ=HOURLY;COUNT=10",
			Input: &Recur{},
			Match: &Recur{
				Frequency: Hourly,
				Count:     10,
			},
			Output: "FREQ=HOURLY;COUNT=10",
		},
		{
			Data:  "FREQ=DAILY;UNTIL=20011125",
			Input: &Recur{},
			Match: &Recur{
				Frequency: Daily,
				Until:     time.Date(2001, 11, 25, 0, 0, 0, 0, time.UTC),
			},
			Output: "FREQ=DAILY;UNTIL=20011125",
		},
		{
			Data:  "FREQ=WEEKLY;UNTIL=20011125T131415",
			Input: &Recur{},
			Match: &Recur{
				Frequency: Weekly,
				Until:     time.Date(2001, 11, 25, 13, 14, 15, 0, time.Local),
				UntilTime: true,
			},
			Output: "FREQ=WEEKLY;UNTIL=20011125T131415",
		},
		{
			Data:  "UNTIL=20011125;FREQ=MONTHLY",
			Input: &Recur{},
			Match: &Recur{
				Frequency: Monthly,
				Until:     time.Date(2001, 11, 25, 0, 0, 0, 0, time.UTC),
			},
			Output: "FREQ=MONTHLY;UNTIL=20011125",
		},
		{
			Data:  "FREQ=YEARLY;UNTIL=20011125T131415Z",
			Input: &Recur{},
			Match: &Recur{
				Frequency: Yearly,
				Until:     time.Date(2001, 11, 25, 13, 14, 15, 0, time.UTC),
				UntilTime: true,
			},
			Output: "FREQ=YEARLY;UNTIL=20011125T131415Z",
		},
		{
			Data:  "FREQ=UNKNOWN",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;COUNT=-1",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;UNTIL=ABC",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;INTERVAL=10",
			Input: &Recur{},
			Match: &Recur{
				Interval: 10,
			},
			Output: "FREQ=SECONDLY;INTERVAL=10",
		},
		{
			Data:  "FREQ=SECONDLY;BYSECOND=1,2,60",
			Input: &Recur{},
			Match: &Recur{
				BySecond: []uint8{1, 2, 60},
			},
			Output: "FREQ=SECONDLY;BYSECOND=1,2,60",
		},
		{
			Data:  "FREQ=SECONDLY;BYSECOND=1,2,61",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYSECOND=1,2,-1",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYSECOND=1,2,A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMINUTE=1,2,59",
			Input: &Recur{},
			Match: &Recur{
				ByMinute: []uint8{1, 2, 59},
			},
			Output: "FREQ=SECONDLY;BYMINUTE=1,2,59",
		},
		{
			Data:  "FREQ=SECONDLY;BYMINUTE=1,2,60",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMINUTE=1,2,-1",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMINUTE=1,2,A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYHOUR=1,2,23",
			Input: &Recur{},
			Match: &Recur{
				ByHour: []uint8{1, 2, 23},
			},
			Output: "FREQ=SECONDLY;BYHOUR=1,2,23",
		},
		{
			Data:  "FREQ=SECONDLY;BYHOUR=1,2,24",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYHOUR=1,2,-1",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYHOUR=1,2,A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYDAY=WE,+2FR,53MO,-53SU,-9TU",
			Input: &Recur{},
			Match: &Recur{
				ByDay: []DayRecur{
					{Day: Wednesday},
					{Day: Friday, Occurence: 2},
					{Day: Monday, Occurence: 53},
					{Day: Sunday, Occurence: -53},
					{Day: Tuesday, Occurence: -9},
				},
			},
			Output: "FREQ=SECONDLY;BYDAY=WE,2FR,53MO,-53SU,-9TU",
		},
		{
			Data:  "FREQ=SECONDLY;BYDAY=UN",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYDAY=-54MO",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYDAY=54MO",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYDAY=MON",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYDAY=MO,",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTHDAY=1,31,-1,-31",
			Input: &Recur{},
			Match: &Recur{
				ByMonthDay: []int8{1, 31, -1, -31},
			},
			Output: "FREQ=SECONDLY;BYMONTHDAY=1,31,-1,-31",
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTHDAY=32",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTHDAY=-32",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTHDAY=0",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTHDAY=A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTHDAY=1,",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYEARDAY=1,366,-1,-366",
			Input: &Recur{},
			Match: &Recur{
				ByYearDay: []int16{1, 366, -1, -366},
			},
			Output: "FREQ=SECONDLY;BYYEARDAY=1,366,-1,-366",
		},
		{
			Data:  "FREQ=SECONDLY;BYYEARDAY=0",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYEARDAY=367",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYEARDAY=-367",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYEARDAY=1,",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYEARDAY=A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYWEEKNO=1,52,-1,-52",
			Input: &Recur{},
			Match: &Recur{
				ByWeekNum: []int8{1, 52, -1, -52},
			},
			Output: "FREQ=SECONDLY;BYWEEKNO=1,52,-1,-52",
		},
		{
			Data:  "FREQ=SECONDLY;BYYWEEKNO=1,0",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYWEEKNO=54",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYWEEKNO=-54",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYYWEEKNO=A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTH=1,2,12",
			Input: &Recur{},
			Match: &Recur{
				ByMonth: []Month{January, February, December},
			},
			Output: "FREQ=SECONDLY;BYMONTH=1,2,12",
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTH=1,2,13",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTH=1,2,-1",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYMONTH=1,2,A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYSETPOS=1,366,-1,-366",
			Input: &Recur{},
			Match: &Recur{
				BySetPos: []int16{1, 366, -1, -366},
			},
			Output: "FREQ=SECONDLY;BYSETPOS=1,366,-1,-366",
		},
		{
			Data:  "FREQ=SECONDLY;BYSETPOS=0",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYSETPOS=367",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYSETPOS=-367",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYSETPOS=1,",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;BYSETPOS=A",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;WKST=SU",
			Input: &Recur{},
			Match: &Recur{
				WeekStart: Sunday,
			},
			Output: "FREQ=SECONDLY;WKST=SU",
		},
		{
			Data:  "FREQ=SECONDLY;WKST=MO",
			Input: &Recur{},
			Match: &Recur{
				WeekStart: Monday,
			},
			Output: "FREQ=SECONDLY;WKST=MO",
		},
		{
			Data:  "FREQ=SECONDLY;WKST=TU",
			Input: &Recur{},
			Match: &Recur{
				WeekStart: Tuesday,
			},
			Output: "FREQ=SECONDLY;WKST=TU",
		},
		{
			Data:  "FREQ=SECONDLY;WKST=WE",
			Input: &Recur{},
			Match: &Recur{
				WeekStart: Wednesday,
			},
			Output: "FREQ=SECONDLY;WKST=WE",
		},
		{
			Data:  "FREQ=SECONDLY;WKST=TH",
			Input: &Recur{},
			Match: &Recur{
				WeekStart: Thursday,
			},
			Output: "FREQ=SECONDLY;WKST=TH",
		},
		{
			Data:  "FREQ=SECONDLY;WKST=FR",
			Input: &Recur{},
			Match: &Recur{
				WeekStart: Friday,
			},
			Output: "FREQ=SECONDLY;WKST=FR",
		},
		{
			Data:  "FREQ=SECONDLY;WKST=SA",
			Input: &Recur{},
			Match: &Recur{
				WeekStart: Saturday,
			},
			Output: "FREQ=SECONDLY;WKST=SA",
		},
		{
			Data:  "FREQ=SECONDLY;WKST=UN",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;UNKNOWN=1",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
		{
			Data:  "FREQ=SECONDLY;UNKOWN",
			Input: &Recur{},
			Error: ErrInvalidRecur,
		},
	})
}

func TestText(t *testing.T) {
	newText := func(s string) *Text {
		nt := Text(s)
		return &nt
	}
	testType(t, []typeTest{
		{
			Data:   "",
			Input:  new(Text),
			Match:  newText(""),
			Output: "",
		},
		{
			Data:   "Jackdaws love my big Sphinx of quartz",
			Input:  new(Text),
			Match:  newText("Jackdaws love my big Sphinx of quartz"),
			Output: "Jackdaws love my big Sphinx of quartz",
		},
		{
			Data:   "Project XYZ Final Review\\nConference Room - 3B\\nCome Prepared.",
			Input:  new(Text),
			Match:  newText("Project XYZ Final Review\nConference Room - 3B\nCome Prepared."),
			Output: "Project XYZ Final Review\\nConference Room - 3B\\nCome Prepared.",
		},
		{
			Data:   "\\\\\\;\\,\\N\\n",
			Input:  new(Text),
			Match:  newText("\\;,\n\n"),
			Output: "\\\\\\;\\,\\n\\n",
		},
		{
			Data:  ";",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  ",",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  "\\",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  "\"",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:  "\\g",
			Input: new(Text),
			Error: ErrInvalidText,
		},
		{
			Data:   "^n",
			Input:  new(Text),
			Match:  newText("\n"),
			Output: "\\n",
		},
		{
			Data:   "^^",
			Input:  new(Text),
			Match:  newText("^"),
			Output: "^^",
		},
		{
			Data:   "^'",
			Input:  new(Text),
			Match:  newText("\""),
			Output: "^'",
		},
	})
}

func TestTime(t *testing.T) {
	l, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	testType(t, []typeTest{
		{
			Data:   "010203",
			Input:  &Time{},
			Match:  &Time{Time: time.Date(0, 1, 1, 1, 2, 3, 0, time.Local)},
			Output: "010203",
		},
		{
			Data:   "010203Z",
			Input:  &Time{},
			Match:  &Time{Time: time.Date(0, 1, 1, 1, 2, 3, 0, time.UTC)},
			Output: "010203Z",
		},
		{
			Params: map[string]string{"TZID": "America/New_York"},
			Data:   "010203",
			Input:  &Time{},
			Match:  &Time{Time: time.Date(0, 1, 1, 1, 2, 3, 0, l)},
			Output: "010203",
		},
	})
}

func TestUTCOffset(t *testing.T) {
	newUTC := func(i int) *UTCOffset {
		u := UTCOffset(i)
		return &u
	}
	testType(t, []typeTest{
		{
			Data:   "0100",
			Input:  new(UTCOffset),
			Match:  newUTC(3600),
			Output: "0100",
		},
		{
			Data:   "0230",
			Input:  new(UTCOffset),
			Match:  newUTC(9000),
			Output: "0230",
		},
		{
			Data:   "023045",
			Input:  new(UTCOffset),
			Match:  newUTC(9045),
			Output: "023045",
		},
		{
			Data:   "+0230",
			Input:  new(UTCOffset),
			Match:  newUTC(9000),
			Output: "0230",
		},
		{
			Data:   "-0230",
			Input:  new(UTCOffset),
			Match:  newUTC(-9000),
			Output: "-0230",
		},
		{
			Data:  "0260",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "020060",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "0261",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "020061",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "0",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "00",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "000",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "00000",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "T",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "0000T",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "000000T",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "-0000",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
		{
			Data:  "-000000",
			Input: new(UTCOffset),
			Error: ErrInvalidOffset,
		},
	})
}
