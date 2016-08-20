package ics

import (
	"bytes"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	tzny, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	tests := []struct {
		Input  *Calendar
		Output string
		Error  error
	}{
		{
			Input: &Calendar{
				ProdID:  "TestDecode",
				Version: "2.0",
			},
			Output: "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID:TestDecode\r\nEND:VCALENDAR\r\n",
		},
		{
			Input: &Calendar{
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
			Output: "BEGIN:VCALENDAR\r\n" +
				"VERSION:2.0\r\n" +
				"PRODID:-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN\r\n" +
				"BEGIN:VEVENT\r\n" +
				"DTSTAMP:19960704T120000Z\r\n" +
				"UID:uid1@example.com\r\n" +
				"DTSTART:19960918T143000Z\r\n" +
				"DESCRIPTION:Networld+Interop Conference and Exhibit\\nAtlanta World Congress\r\n" +
				"  Center\\nAtlanta\\, Georgia\r\n" +
				"ORGANIZER:mailto:jsmith@example.com\r\n" +
				"STATUS:CONFIRMED\r\n" +
				"SUMMARY:Networld+Interop Conference\r\n" +
				"DTEND:19960920T220000Z\r\n" +
				"CATEGORIES:CONFERENCE\r\n" +
				"END:VEVENT\r\n" +
				"END:VCALENDAR\r\n",
		},
		{
			Input: &Calendar{
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
			Output: "BEGIN:VCALENDAR\r\n" +
				"VERSION:2.0\r\n" +
				"PRODID:-//RDU Software//NONSGML HandCal//EN\r\n" +
				"BEGIN:VEVENT\r\n" +
				"DTSTAMP:19980309T231000Z\r\n" +
				"UID:guid-1.example.com\r\n" +
				"DTSTART;TZID=America/New_York:19980312T083000\r\n" +
				"CLASS:PUBLIC\r\n" +
				"CREATED:19980309T130000Z\r\n" +
				"DESCRIPTION:Project XYZ Review Meeting\r\n" +
				"LOCATION:1CP Conference Room 4350\r\n" +
				"ORGANIZER:mailto:mrbig@example.com\r\n" +
				"SUMMARY:XYZ Project Review\r\n" +
				"DTEND;TZID=America/New_York:19980312T093000\r\n" +
				"ATTENDEE;CUTYPE=GROUP;ROLE=REQ-PARTICIPANT;RSVP=TRUE:mailto:employee-A@exam\r\n" +
				" ple.com\r\n" +
				"CATEGORIES:MEETING\r\n" +
				"END:VEVENT\r\n" +
				"BEGIN:VTIMEZONE\r\n" +
				"TZID:America/New_York\r\n" +
				"BEGIN:STANDARD\r\n" +
				"DTSTART:19981025T020000\r\n" +
				"TZOFFSETTO:-0500\r\n" +
				"TZOFFSETFROM:-0400\r\n" +
				"TZNAME:EST\r\n" +
				"END:STANDARD\r\n" +
				"BEGIN:DAYLIGHT\r\n" +
				"DTSTART:19990404T020000\r\n" +
				"TZOFFSETTO:-0400\r\n" +
				"TZOFFSETFROM:-0500\r\n" +
				"TZNAME:EDT\r\n" +
				"END:DAYLIGHT\r\n" +
				"END:VTIMEZONE\r\n" +
				"END:VCALENDAR\r\n",
		},
		{
			Input: &Calendar{
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
			Output: "BEGIN:VCALENDAR\r\n" +
				"VERSION:2.0\r\n" +
				"PRODID:-//ABC Corporation//NONSGML My Product//EN\r\n" +
				"BEGIN:VTODO\r\n" +
				"DTSTAMP:19980130T134500Z\r\n" +
				"UID:uid4@example.com\r\n" +
				"ORGANIZER:mailto:unclesam@example.com\r\n" +
				"SEQUENCE:2\r\n" +
				"STATUS:NEEDS-ACTION\r\n" +
				"SUMMARY:Submit Income Taxes\r\n" +
				"DUE:19980415T000000\r\n" +
				"ATTENDEE;PARTSTAT=ACCEPTED:mailto:jqpublic@example.com\r\n" +
				"BEGIN:VALARM\r\n" +
				"ACTION:AUDIO\r\n" +
				"TRIGGER;VALUE=DATE-TIME:19980403T120000Z\r\n" +
				"DURATION:PT1H\r\n" +
				"REPEAT:4\r\n" +
				"ATTACH;FMTTYPE=audio/basic:http://example.com/pub/audio-files/ssbanner.aud\r\n" +
				"END:VALARM\r\n" +
				"END:VTODO\r\n" +
				"END:VCALENDAR\r\n",
		},
		{
			Input: &Calendar{
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
			Output: "BEGIN:VCALENDAR\r\n" +
				"VERSION:2.0\r\n" +
				"PRODID:-//ABC Corporation//NONSGML My Product//EN\r\n" +
				"BEGIN:VJOURNAL\r\n" +
				"DTSTAMP:19970324T120000Z\r\n" +
				"UID:uid5@example.com\r\n" +
				"CLASS:PUBLIC\r\n" +
				"ORGANIZER:mailto:jsmith@example.com\r\n" +
				"STATUS:DRAFT\r\n" +
				"CATEGORIES:Project Report,XYZ,Weekly Meeting\r\n" +
				"DESCRIPTION:Project xyz Review Meeting Minutes\\nAgenda\\n1. Review of projec\r\n" +
				" t version 1.0 requirements.\\n2. Definition of project processes.\\n3. Review\r\n" +
				"  of project schedule.\\nParticipants: John Smith\\, Jane Doe\\, Jim Dandy\\n-It\r\n" +
				"  was decided that the requirements need to be signed off by product marketi\r\n" +
				" ng.\\n-Project processes were accepted.\\n-Project schedule needs to account \r\n" +
				" for scheduled holidays and employee vacation time. Check with HR for specif\r\n" +
				" ic dates.\\n-New schedule will be distributed by Friday.\\n-Next weeks meetin\r\n" +
				" g is cancelled. No meeting until 3/23.\r\n" +
				"END:VJOURNAL\r\n" +
				"END:VCALENDAR\r\n",
		},
		{
			Input: &Calendar{
				Version: "2.0",
				ProdID:  "-//RDU Software//NONSGML HandCal//EN",
				FreeBusy: []FreeBusy{
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
						DateTimeStart: &PropDateTimeStart{
							DateTime: &DateTime{
								Time: time.Date(1998, 3, 13, 14, 17, 11, 0, time.UTC),
							},
						},
						DateTimeEnd: &PropDateTimeEnd{
							DateTime: &DateTime{
								Time: time.Date(1998, 4, 10, 14, 17, 11, 0, time.UTC),
							},
						},
						FreeBusy: []PropFreeBusy{
							{
								Period: Period{
									Start: DateTime{
										Time: time.Date(1998, 3, 14, 23, 30, 0, 0, time.UTC),
									},
									End: DateTime{
										Time: time.Date(1998, 3, 15, 0, 30, 0, 0, time.UTC),
									},
								},
							},
							{
								Period: Period{
									Start: DateTime{
										Time: time.Date(1998, 3, 16, 15, 30, 0, 0, time.UTC),
									},
									End: DateTime{
										Time: time.Date(1998, 3, 16, 16, 30, 0, 0, time.UTC),
									},
								},
							},
							{
								Period: Period{
									Start: DateTime{
										Time: time.Date(1998, 3, 18, 3, 0, 0, 0, time.UTC),
									},
									End: DateTime{
										Time: time.Date(1998, 3, 18, 4, 0, 0, 0, time.UTC),
									},
								},
							},
						},
						URL: &PropURL{
							URL: url.URL{
								Scheme: "http",
								Host:   "www.example.com",
								Path:   "/calendar/busytime/jsmith.ifb",
							},
						},
					},
				},
			},
			Output: "BEGIN:VCALENDAR\r\n" +
				"VERSION:2.0\r\n" +
				"PRODID:-//RDU Software//NONSGML HandCal//EN\r\n" +
				"BEGIN:VFREEBUSY\r\n" +
				"DTSTAMP:19970324T120000Z\r\n" +
				"UID:uid5@example.com\r\n" +
				"DTSTART:19980313T141711Z\r\n" +
				"DTEND:19980410T141711Z\r\n" +
				"ORGANIZER:mailto:jsmith@example.com\r\n" +
				"URL:http://www.example.com/calendar/busytime/jsmith.ifb\r\n" +
				"FREEBUSY:19980314T233000Z/19980315T003000Z\r\n" +
				"FREEBUSY:19980316T153000Z/19980316T163000Z\r\n" +
				"FREEBUSY:19980318T030000Z/19980318T040000Z\r\n" +
				"END:VFREEBUSY\r\n" +
				"END:VCALENDAR\r\n",
		},
	}

	var buf bytes.Buffer
	for n, test := range tests {
		err := Encode(&buf, test.Input)
		if err != test.Error {
			t.Errorf("test %d: expecting error %q, got %q", n+1, test.Error, err)
		} else if str := buf.String(); str != test.Output {
			t.Errorf("test %d: expecting calendar %s, got %s", n+1, test.Output, str)
		} else {
			c, err := Decode(strings.NewReader(str))
			if err != nil {
				t.Errorf("test %d: unexpected error decoding encoded string: %s", n+1, err)
			} else if !reflect.DeepEqual(c, test.Input) {
				t.Errorf("test %d: expecting calendar %s, got %s", n+1, test.Input, c)
			}
		}
		buf.Reset()
	}
}
