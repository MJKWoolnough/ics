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

func parseAddress(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return *u
}

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
	confirmed := StatusConfirmed
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
								URL: parseAddress("mailto:jsmith@example.com"),
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
						Status: &confirmed,
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
