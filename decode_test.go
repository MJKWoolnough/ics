package ics

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func (c Calendar) Error() string {
	var buf bytes.Buffer
	buf.WriteString("Calendar: ")
	indentWrite(reflect.ValueOf(&c), 1, &buf)
	return buf.String()
}

func indentWrite(v reflect.Value, indentLevel int, w writer) {
	w.WriteString(v.Type().String())
	w.WriteString(" ")
	if v.Type().Kind() == reflect.Ptr {
		if v.IsNil() {
			w.WriteString("<nil>\n")
			return
		}
		v = v.Elem()
	}
	if s, ok := v.Interface().(interface {
		String() string
	}); ok {
		w.WriteString(s.String())
	} else {
		switch v.Type().Kind() {
		case reflect.Struct:
			w.WriteString("{\n")
			for i := 0; i < v.Type().NumField(); i++ {
				indent(w, indentLevel)
				w.WriteString(v.Type().Field(i).Name)
				w.WriteString(": ")
				indentWrite(v.Field(i), indentLevel+1, w)
			}
			indent(w, indentLevel-1)
			w.WriteString("}")
		case reflect.String:
			fmt.Fprintf(w, "%q", v.String())
		case reflect.Slice:
			w.WriteString("{")
			if v.Len() > 0 {
				w.WriteString("\n")
				for i := 0; i < v.Len(); i++ {
					indent(w, indentLevel)
					fmt.Fprintf(w, "%d", i)
					w.WriteString(": ")
					indentWrite(v.Index(i), indentLevel+1, w)
				}
				indent(w, indentLevel-1)
			}
			w.WriteString("}")
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			w.WriteString(v.Kind().String())
			w.WriteString(" <")
			fmt.Fprintf(w, "%d", v.Int())
			w.WriteString(">")
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			w.WriteString(v.Kind().String())
			w.WriteString(" <")
			fmt.Fprintf(w, "%d", v.Uint())
			w.WriteString(">")
		case reflect.Bool:
			if v.Bool() {
				w.WriteString("True")
			} else {
				w.WriteString("False")
			}
		default:
			w.WriteString(v.Kind().String())
		}
	}
	w.WriteString("\n")
}

func indent(w io.Writer, indentLevel int) {
	in := [1]byte{'	'}
	for i := 0; i < indentLevel; i++ {
		w.Write(in[:])
	}
}

func TestDecode(t *testing.T) {
	tzny, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	tests := []struct {
		Input  string
		Output *Calendar
		Error  error
	}{
		{
			Error: io.ErrUnexpectedEOF,
		},
		{
			Input: "BEGIN:VCALENDAR\r\nPRODID:TestDecode\r\nVERSION:2.0\r\nEND:VCALENDAR\r\n",
			Output: &Calendar{
				ProdID:  "TestDecode",
				Version: "2.0",
			},
		},
		{
			Input: "BEGIN:VCALENDAR\r\n" +
				"PRODID:-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN\r\n" +
				"VERSION:2.0\r\n" +
				"BEGIN:VEVENT\r\n" +
				"DTSTAMP:19960704T120000Z\r\n" +
				"UID:uid1@example.com\r\n" +
				"ORGANIZER:mailto:jsmith@example.com\r\n" +
				"DTSTART:19960918T143000Z\r\n" +
				"DTEND:19960920T220000Z\r\n" +
				"STATUS:CONFIRMED\r\n" +
				"CATEGORIES:CONFERENCE\r\n" +
				"SUMMARY:Networld+Interop Conference\r\n" +
				"DESCRIPTION:Networld+Interop Conference\r\n  and Exhibit\\nAtlanta World Congress Center\\n\r\n Atlanta\\, Georgia\r\n" +
				"END:VEVENT\r\n" +
				"END:VCALENDAR\r\n",
			Output: &Calendar{
				ProdID:  "-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN",
				Version: "2.0",
				Event: []Event{
					{
						DateTimeStamp: PropDateTimeStamp{
							Time: time.Date(1996, 7, 4, 12, 0, 0, 0, time.UTC),
						},
						UID: "uid1@example.com",
						Organizer: &PropOrganizer{
							CalendarAddress: CalendarAddress{
								URL: url.URL{
									Scheme: "mailto",
									Opaque: "jsmith@example.com",
								},
							},
						},
						DateTimeStart: &PropDateTimeStart{
							DateTime: &DateTime{
								Time: time.Date(1996, 9, 18, 14, 30, 0, 0, time.UTC),
							},
						},
						DateTimeEnd: &PropDateTimeEnd{
							DateTime: &DateTime{
								Time: time.Date(1996, 9, 20, 22, 0, 0, 0, time.UTC),
							},
						},
						Status: StatusConfirmed.New(),
						Categories: []PropCategories{
							{
								MText: MText{"CONFERENCE"},
							},
						},
						Summary: &PropSummary{
							Text: "Networld+Interop Conference",
						},
						Description: &PropDescription{
							Text: "Networld+Interop Conference and Exhibit\nAtlanta World Congress Center\nAtlanta, Georgia",
						},
					},
				},
			},
		},
		{
			Input: "BEGIN:VCALENDAR\r\n" +
				"PRODID:-//RDU Software//NONSGML HandCal//EN\r\n" +
				"VERSION:2.0\r\n" +
				"BEGIN:VTIMEZONE\r\n" +
				"TZID:America/New_York\r\n" +
				"BEGIN:STANDARD\r\n" +
				"DTSTART:19981025T020000\r\n" +
				"TZOFFSETFROM:-0400\r\n" +
				"TZOFFSETTO:-0500\r\n" +
				"TZNAME:EST\r\n" +
				"END:STANDARD\r\n" +
				"BEGIN:DAYLIGHT\r\n" +
				"DTSTART:19990404T020000\r\n" +
				"TZOFFSETFROM:-0500\r\n" +
				"TZOFFSETTO:-0400\r\n" +
				"TZNAME:EDT\r\n" +
				"END:DAYLIGHT\r\n" +
				"END:VTIMEZONE\r\n" +
				"BEGIN:VEVENT\r\n" +
				"DTSTAMP:19980309T231000Z\r\n" +
				"UID:guid-1.example.com\r\n" +
				"ORGANIZER:mailto:mrbig@example.com\r\n" +
				"ATTENDEE;RSVP=TRUE;ROLE=REQ-PARTICIPANT;CUTYPE=GROUP:\r\n" +
				" mailto:employee-A@example.com\r\n" +
				"DESCRIPTION:Project XYZ Review Meeting\r\n" +
				"CATEGORIES:MEETING\r\n" +
				"CLASS:PUBLIC\r\n" +
				"CREATED:19980309T130000Z\r\n" +
				"SUMMARY:XYZ Project Review\r\n" +
				"DTSTART;TZID=America/New_York:19980312T083000\r\n" +
				"DTEND;TZID=America/New_York:19980312T093000\r\n" +
				"LOCATION:1CP Conference Room 4350\r\n" +
				"END:VEVENT\r\n" +
				"END:VCALENDAR\r\n",
			Output: &Calendar{
				ProdID:  "-//RDU Software//NONSGML HandCal//EN",
				Version: "2.0",
				Event: []Event{
					{
						DateTimeStamp: PropDateTimeStamp{
							Time: time.Date(1998, 3, 9, 23, 10, 0, 0, time.UTC),
						},
						UID: "guid-1.example.com",
						Organizer: &PropOrganizer{
							CalendarAddress: CalendarAddress{
								URL: url.URL{
									Scheme: "mailto",
									Opaque: "mrbig@example.com",
								},
							},
						},
						Attendee: []PropAttendee{
							{
								CalendarUserType:  CalendarUserTypeGroup.New(),
								RSVP:              NewRSVP(true),
								ParticipationRole: ParticipationRoleRequiredParticipant.New(),
								CalendarAddress: CalendarAddress{
									URL: url.URL{
										Scheme: "mailto",
										Opaque: "employee-A@example.com",
									},
								},
							},
						},
						Description: &PropDescription{
							Text: "Project XYZ Review Meeting",
						},
						Categories: []PropCategories{
							{
								MText: MText{"MEETING"},
							},
						},
						Class: ClassPublic.New(),
						Created: &PropCreated{
							Time: time.Date(1998, 3, 9, 13, 0, 0, 0, time.UTC),
						},
						DateTimeStart: &PropDateTimeStart{
							DateTime: &DateTime{
								Time: time.Date(1998, 3, 12, 8, 30, 0, 0, tzny),
							},
						},
						DateTimeEnd: &PropDateTimeEnd{
							DateTime: &DateTime{
								Time: time.Date(1998, 3, 12, 9, 30, 0, 0, tzny),
							},
						},
						Location: &PropLocation{
							Text: "1CP Conference Room 4350",
						},
						Summary: &PropSummary{
							Text: "XYZ Project Review",
						},
					},
				},
				Timezone: []Timezone{
					{
						TimezoneID: "America/New_York",
						Standard: []Standard{
							{
								DateTimeStart: PropDateTimeStart{
									DateTime: &DateTime{
										Time: time.Date(1998, 10, 25, 2, 0, 0, 0, time.Local),
									},
								},
								TimezoneOffsetFrom: -4 * 3600,
								TimezoneOffsetTo:   -5 * 3600,
								TimezoneName: []PropTimezoneName{
									{
										Text: "EST",
									},
								},
							},
						},
						Daylight: []Daylight{
							{
								DateTimeStart: PropDateTimeStart{
									DateTime: &DateTime{
										Time: time.Date(1999, 4, 4, 2, 0, 0, 0, time.Local),
									},
								},
								TimezoneOffsetFrom: -5 * 3600,
								TimezoneOffsetTo:   -4 * 3600,
								TimezoneName: []PropTimezoneName{
									{
										Text: "EDT",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Input: "BEGIN:VCALENDAR\r\n" +
				"VERSION:2.0\r\n" +
				"PRODID:-//ABC Corporation//NONSGML My Product//EN\r\n" +
				"BEGIN:VTODO\r\n" +
				"DTSTAMP:19980130T134500Z\r\n" +
				"SEQUENCE:2\r\n" +
				"UID:uid4@example.com\r\n" +
				"ORGANIZER:mailto:unclesam@example.com\r\n" +
				"ATTENDEE;PARTSTAT=ACCEPTED:mailto:jqpublic@example.com\r\n" +
				"DUE:19980415T000000\r\n" +
				"STATUS:NEEDS-ACTION\r\n" +
				"SUMMARY:Submit Income Taxes\r\n" +
				"BEGIN:VALARM\r\n" +
				"ACTION:AUDIO\r\n" +
				"TRIGGER;VALUE=DATE-TIME:19980403T120000Z\r\n" +
				"ATTACH;FMTTYPE=audio/basic:http://example.com/pub/audio-\r\n" +
				" files/ssbanner.aud\r\n" +
				"REPEAT:4\r\n" +
				"DURATION:PT1H\r\n" +
				"END:VALARM\r\n" +
				"END:VTODO\r\n" +
				"END:VCALENDAR\r\n",
			Output: &Calendar{
				Version: "2.0",
				ProdID:  "-//ABC Corporation//NONSGML My Product//EN",
				Todo: []Todo{
					{
						DateTimeStamp: PropDateTimeStamp{
							Time: time.Date(1998, 1, 30, 13, 45, 0, 0, time.UTC),
						},
						Sequence: NewSequence(2),
						UID:      "uid4@example.com",
						Organizer: &PropOrganizer{
							CalendarAddress: CalendarAddress{
								URL: url.URL{
									Scheme: "mailto",
									Opaque: "unclesam@example.com",
								},
							},
						},
						Attendee: []PropAttendee{
							{
								ParticipationStatus: ParticipationStatusAccepted.New(),
								CalendarAddress: CalendarAddress{
									URL: url.URL{
										Scheme: "mailto",
										Opaque: "jqpublic@example.com",
									},
								},
							},
						},
						Due: &PropDue{
							DateTime: &DateTime{
								Time: time.Date(1998, 4, 15, 0, 0, 0, 0, time.Local),
							},
						},
						Status: StatusNeedsAction.New(),
						Summary: &PropSummary{
							Text: "Submit Income Taxes",
						},
						Alarm: []Alarm{
							{
								AlarmType: &AlarmAudio{
									Trigger: PropTrigger{
										DateTime: &DateTime{
											Time: time.Date(1998, 4, 3, 12, 0, 0, 0, time.UTC),
										},
									},
									Attachment: []PropAttachment{
										{
											FormatType: NewFormatType("audio/basic"),
											URI: &URI{
												URL: url.URL{
													Scheme: "http",
													Host:   "example.com",
													Path:   "/pub/audio-files/ssbanner.aud",
												},
											},
										},
									},
									Repeat: NewRepeat(4),
									Duration: &PropDuration{
										Hours: 1,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Input: "BEGIN:VCALENDAR\r\n" +
				"VERSION:2.0\r\n" +
				"PRODID:-//ABC Corporation//NONSGML My Product//EN\r\n" +
				"BEGIN:VJOURNAL\r\n" +
				"DTSTAMP:19970324T120000Z\r\n" +
				"UID:uid5@example.com\r\n" +
				"ORGANIZER:mailto:jsmith@example.com\r\n" +
				"STATUS:DRAFT\r\n" +
				"CLASS:PUBLIC\r\n" +
				"CATEGORIES:Project Report,XYZ,Weekly Meeting\r\n" +
				"DESCRIPTION:Project xyz Review Meeting Minutes\\n\r\n" +
				" Agenda\\n1. Review of project version 1.0 requirements.\\n2.\r\n" +
				"  Definition of project processes.\\n3. Review of project schedule.\\n\r\n" +
				" Participants: John Smith\\, Jane Doe\\, Jim Dandy\\n-It was\r\n" +
				"  decided that the requirements need to be signed off by\r\n" +
				"  product marketing.\\n-Project processes were accepted.\\n\r\n" +
				" -Project schedule needs to account for scheduled holidays\r\n" +
				"  and employee vacation time. Check with HR for specific\r\n" +
				"  dates.\\n-New schedule will be distributed by Friday.\\n-\r\n" +
				" Next weeks meeting is cancelled. No meeting until 3/23.\r\n" +
				"END:VJOURNAL\r\n" +
				"END:VCALENDAR\r\n",
			Output: &Calendar{
				Version: "2.0",
				ProdID:  "-//ABC Corporation//NONSGML My Product//EN",
				Journal: []Journal{
					{
						DateTimeStamp: PropDateTimeStamp{
							Time: time.Date(1997, 3, 24, 12, 0, 0, 0, time.UTC),
						},
						UID: "uid5@example.com",
						Organizer: &PropOrganizer{
							CalendarAddress: CalendarAddress{
								URL: url.URL{
									Scheme: "mailto",
									Opaque: "jsmith@example.com",
								},
							},
						},
						Status: StatusDraft.New(),
						Class:  ClassPublic.New(),
						Categories: []PropCategories{
							{
								MText: MText{
									"Project Report",
									"XYZ",
									"Weekly Meeting",
								},
							},
						},
						Description: []PropDescription{
							{
								Text: "Project xyz Review Meeting Minutes\n" +
									"Agenda\n" +
									"1. Review of project version 1.0 requirements.\n" +
									"2. Definition of project processes.\n" +
									"3. Review of project schedule.\n" +
									"Participants: John Smith, Jane Doe, Jim Dandy\n" +
									"-It was decided that the requirements need to be signed off by product marketing.\n" +
									"-Project processes were accepted.\n" +
									"-Project schedule needs to account for scheduled holidays and employee vacation time. Check with HR for specific dates.\n" +
									"-New schedule will be distributed by Friday.\n" +
									"-Next weeks meeting is cancelled. No meeting until 3/23.",
							},
						},
					},
				},
			},
		},
	}

	for n, test := range tests {
		c, err := Decode(strings.NewReader(test.Input))
		if err != test.Error {
			t.Errorf("test %d: expecting error %q, got %q", n+1, test.Error, err)
		} else if !reflect.DeepEqual(c, test.Output) {
			t.Errorf("test %d: expecting calendar %s, got %s", n+1, test.Output, c)
		}
	}
}
