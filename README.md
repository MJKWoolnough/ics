# ics

[![CI](https://github.com/MJKWoolnough/ics/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/ics/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/ics.svg)](https://pkg.go.dev/vimagination.zapto.org/ics)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/ics)](https://goreportcard.com/report/vimagination.zapto.org/ics)

--
    import "vimagination.zapto.org/ics"

Package ics implements an encoder and decoder for iCalendar files.

## Highlights

 - Parse iCalendar (ics) files/streams into object.
 - Encode objects into iCalendar files/streams.

## Usage

```go
package main

import (
	"fmt"
	"strings"

	"vimagination.zapto.org/ics"
)

func main() {
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
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/ics
