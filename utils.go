package ics

import (
	"errors"
	"strconv"
	"time"

	strparse "github.com/MJKWoolnough/parser"
)

func escape(s string) []byte {
	p := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\\':
			p = append(p, '\\', '\\')
		case ';':
			p = append(p, '\\', ';')
		case ',':
			p = append(p, '\\', ',')
		case '\n':
			p = append(p, '\\', 'n')
		default:
			p = append(p, s[i])
		}
	}
	return p
}

func unescape(p string) []byte {
	u := make([]byte, 0, len(p))
	for i := 0; i < len(p); i++ {
		if p[i] == '\\' && i+1 < len(p) {
			i++
			switch p[i] {
			case '\\', ';', ',':
				u = append(u, p[i])
			case 'N', 'n':
				u = append(u, '\n')
			default:
				u = append(u, '\\', p[i])
			}
		} else {
			u = append(u, p[i])
		}
	}
	return u
}

func escape6868(s string) []byte {
	p := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\n':
			p = append(p, '^', 'n')
		case '^':
			p = append(p, '^', '^')
		case '"':
			p = append(p, '^', '\'')
		default:
			p = append(p, s[i])
		}
	}
	return p
}

func unescape6868(p string) []byte {
	u := make([]byte, 0, len(p))
	for i := 0; i < len(p); i++ {
		if p[i] == '^' && i+1 < len(p) {
			i++
			switch p[i] {
			case 'n':
				u = append(u, '\n') //crlf on windows?
			case '^':
				u = append(u, '^')
			case '\'':
				u = append(u, '"')
			default:
				u = append(u, '^', p[i])
			}
		} else {
			u = append(u, p[i])
		}
	}
	return u
}

func textSplit(s string, delim byte) []string {
	toRet := make([]string, 0, 1)
	lastPos := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\\':
			i++
		case delim:
			toRet = append(toRet, string(unescape(s[lastPos:i])))
			lastPos = i + 1
		}
	}
	if lastPos <= len(s) {
		toRet = append(toRet, string(unescape(s[lastPos:])))
	}
	return toRet
}

type dateTime struct {
	justDate bool
	time.Time
}

func (dt dateTime) Add(d time.Duration) dateTime {
	if d%24*time.Hour != 0 {
		dt.justDate = false
	}
	dt.Time = dt.Time.Add(d)
	return dt
}

func (dt dateTime) AddDate(years, months, days int) dateTime {
	dt.Time = dt.Time.AddDate(years, months, days)
	return dt
}

func (dt dateTime) In(loc *time.Location) dateTime {
	dt.Time = dt.Time.In(loc)
	return dt
}

func (dt dateTime) String() string {
	if dt.justDate {
		return dt.Format("20060102")
	}
	switch dt.Location() {
	case nil, time.UTC:
		return dt.Format("20060102T150405Z")
	default:
		return dt.Format("20060102T150405")
	}
}

func parseDate(s string) (dateTime, error) {
	t, err := time.Parse("20060102", s)
	return dateTime{true, t}, err
}

func parseDateTime(s string, l *time.Location) (dateTime, error) {
	var (
		t   time.Time
		err error
	)
	if l == nil {
		if s[len(s)-1] == 'Z' {
			t, err = time.Parse("20060102T150405Z", s)
		} else {
			t, err = time.ParseInLocation("20060102T150405", s, time.Local)
		}
	} else {
		t, err = time.ParseInLocation("20060102T150405", s, l)
	}
	return dateTime{Time: t}, err
}

func parseTime(s string, l *time.Location) (time.Time, error) {
	if l == nil {
		if s[len(s)-1] == 'Z' {
			return time.Parse("150405Z", s)
		} else {
			return time.ParseInLocation("150405Z", s, time.Local)
		}
	}
	return time.ParseInLocation("150405", s, l)
}

const nums = "0123456789"

func parseDuration(s string) (time.Duration, error) {
	p := strparse.NewStringParser(s)
	var (
		dur time.Duration
		neg bool
	)
	if p.Accept("-") {
		neg = true
	} else {
		p.Accept("+")
	}
	if !p.Accept("P") {
		return 0, ErrInvalidDuration
	}
	p.Get()
	if !p.Accept("T") {
		p.AcceptRun(nums)
		num := p.Get()
		if len(num) == 0 {
			return 0, ErrInvalidDuration
		}
		n, _ := strconv.Atoi(num)
		p.Accept("DW")
		switch p.Get() {
		case "D":
			dur = time.Duration(n) * time.Hour * 24
		case "W":
			return time.Duration(n) * time.Hour * 24 * 7, nil
		default:
			return 0, ErrInvalidDuration
		}
		p.Except("")
		switch p.Get() {
		case "":
			if neg {
				return -dur, nil
			}
			return dur, nil
		case "T":
		default:
			return 0, ErrInvalidDuration
		}
	} else {
		p.Get()
	}
	toRead := "HMS"
	var readTime bool
	for len(toRead) > 0 {
		p.AcceptRun(nums)
		num := p.Get()
		if len(num) == 0 {
			if !readTime {
				return 0, ErrInvalidDuration
			}
			break
		}
		n, _ := strconv.Atoi(num)
		p.Accept(toRead)
		switch p.Get() {
		case "H":
			dur += time.Duration(n) * time.Hour
			toRead = "MS"
		case "M":
			dur += time.Duration(n) * time.Minute
			toRead = "S"
		case "S":
			dur += time.Duration(n) * time.Second
			toRead = ""
		default:
			return 0, ErrInvalidDuration
		}
		readTime = true
	}
	if neg {
		return -dur, nil
	}
	return dur, nil
}

func durationString(d time.Duration) string {
	toRet := make([]byte, 0, 16)
	if d < 0 {
		toRet = append(toRet, '-')
		d = -d
	}
	toRet = append(toRet, 'P')
	if d%(time.Hour*24*7) == 0 {
		toRet = append(toRet, strconv.FormatInt(int64(d/(time.Hour*24*7)), 10)...)
		toRet = append(toRet, 'W')
	} else {
		if d >= time.Hour*24 {
			toRet = append(toRet, strconv.FormatInt(int64(d/(time.Hour*24)), 10)...)
			toRet = append(toRet, 'D')
			d = d % (time.Hour * 24)
		}
		if d > 0 {
			toRet = append(toRet, 'T')
			if d >= time.Hour {
				toRet = append(toRet, strconv.FormatInt(int64(d/time.Hour), 10)...)
				toRet = append(toRet, 'H')
				d = d % time.Hour
			}
			if d >= time.Minute {
				toRet = append(toRet, strconv.FormatInt(int64(d/time.Minute), 10)...)
				toRet = append(toRet, 'M')
				d = d % time.Minute
			}
			if d >= time.Second {
				toRet = append(toRet, strconv.FormatInt(int64(d/time.Second), 10)...)
				toRet = append(toRet, 'S')
			}
		}
	}
	return string(toRet)
}

func parseOffset(s string) (int, error) {
	if len(s) != 5 && len(s) != 7 {
		return 0, ErrInvalidOffset
	}
	var neg bool
	if s[0] == '-' {
		neg = true
	} else if s[0] != '+' {
		return 0, ErrInvalidOffset
	}
	hours, err := strconv.Atoi(s[1:3])
	if err != nil {
		return 0, ErrInvalidOffset
	}
	minutes, err := strconv.Atoi(s[3:5])
	if err != nil {
		return 0, ErrInvalidOffset
	}
	var seconds int
	if len(s) == 7 {
		seconds, err = strconv.Atoi(s[5:7])
		if err != nil {
			return 0, ErrInvalidOffset
		}
	}
	val := hours*3600 + minutes*60 + seconds
	if neg {
		return -val, nil
	}
	return val, nil
}

func offsetString(o int) string {
	toRet := make([]byte, 1, 7)
	if o < 0 {
		toRet[0] = '-'
		o = -o
	} else {
		toRet[0] = '+'
	}
	toRet = append(toRet, strconv.Itoa(o/3600)...)
	toRet = append(toRet, strconv.Itoa((o%3600)/60)...)
	seconds := o % 60
	if seconds > 0 {
		toRet = append(toRet, strconv.Itoa(seconds)...)
	}
	return string(toRet)
}

func dquote(p []byte) []byte {
	q := make([]byte, 0, len(p)+2)
	q = append(q, '"')
	q = append(q, p...)
	q = append(q, '"')
	return q
}

func dquoteIfNeeded(p []byte) []byte {
	for _, c := range p {
		switch c {
		case ';', ':', ',':
			return dquote(p)
		}
	}
	return p
}

// Errors

var (
	ErrInvalidDuration = errors.New("invalid duration string")
	ErrInvalidOffset   = errors.New("invalid offset string")
)
