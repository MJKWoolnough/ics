package ics_test

import (
	"fmt"
	"strings"

	"vimagination.zapto.org/ics"
)

func Example() {
	const icsData = "BEGIN:VCALENDAR\r\n" +
		"VERSION:2.0\r\n" +
		"PRODID:-//hacksw/handcal//NONSGML v1.0//EN\r\n" +
		"BEGIN:VEVENT\r\n" +
		"UID:uid1@example.com\r\n" +
		"ORGANIZER;CN=John Doe:MAILTO:john.doe@example.com\r\n" +
		"DTSTAMP:19970701T100000Z\r\n" +
		"DTSTART:19970714T170000Z\r\n" +
		"DTEND:19970715T040000Z\r\n" +
		"SUMMARY:Bastille Day Party\r\n" +
		"GEO:48.85299;2.36885\r\n" +
		"END:VEVENT\r\n" +
		"END:VCALENDAR"

	cal, err := ics.Decode(strings.NewReader(icsData))
	if err != nil {
		fmt.Println("Error: ", err)

		return
	}

	fmt.Println(cal.Event[0].DateTimeStart.DateTime.Time, "-", cal.Event[0].Summary.Text)

	// Output:
	// 1997-07-14 17:00:00 +0000 UTC - Bastille Day Party
}
