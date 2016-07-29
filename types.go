package ics

import (
	"encoding/base64"
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/MJKWoolnough/parser"
)

const dateTimeFormat = "20060102T150405Z"

type Binary []byte

func (b *Binary) Decode(params map[string]string, data string) error {
	if params["ENCODING"] != "BASE64" {
		return ErrInvalidEncoding
	}
	cb, err := base64.StdEncoding.DecodeString(data)
	*b = cb
	return err
}

func (b *Binary) Encode(w io.Writer) {
	e := base64.NewEncoder(base64.StdEncoding, w)
	e.Write([]byte(*b))
	e.Close()
}

type Boolean bool

func (b *Boolean) Decode(params map[string]string, data string) error {
	cb, err := strconv.ParseBool(data)
	*b = Boolean(cb)
	if err != nil {
		return ErrInvalidBoolean
	}
	return nil
}

var (
	booleanTrue  = [...]byte{'T', 'R', 'U', 'E'}
	booleanFalse = [...]byte{'F', 'A', 'L', 'S', 'E'}
)

func (b *Boolean) Encode(w io.Writer) {
	if *b {
		w.Write(booleanTrue[:])
	} else {
		w.Write(booleanFalse[:])
	}
}

type CalAddress struct {
	URI
}

type Date struct {
	time.Time
}

func (d *Date) Decode(params map[string]string, data string) error {
	t, err := time.Parse(dateTimeFormat[:8], data)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *Date) Encode(w io.Writer) {
	b := make([]byte, 0, 8)
	w.Write([]byte(d.AppendFormat(b, dateTimeFormat[:8])))
}

type DateTime struct {
	time.Time
}

func (d *DateTime) Decode(params map[string]string, data string) error {
	if tz, ok := params["TZID"]; ok {
		l, err := time.LoadLocation(tz)
		if err != nil {
			return err
		}
		t, err := time.ParseInLocation(dateTimeFormat[:15], data, l)
		if err != nil {
			return err
		}
		d.Time = t
	} else if len(data) > 0 && data[len(data)-1] == 'Z' {
		t, err := time.ParseInLocation(dateTimeFormat, data, time.UTC)
		if err != nil {
			return err
		}
		d.Time = t
	} else {
		t, err := time.ParseInLocation(dateTimeFormat[:15], data, time.Local)
		if err != nil {
			return err
		}
		d.Time = t
	}
	return nil
}

func (d *DateTime) Encode(w io.Writer) {
	b := make([]byte, 0, 16)
	switch d.Location() {
	case time.UTC:
		b = d.AppendFormat(b, dateTimeFormat)
	case time.Local:
		b = d.AppendFormat(b, dateTimeFormat[:15])
	default:
		b = d.AppendFormat(b, dateTimeFormat[:15])
	}
	w.Write(b)
}

type Duration struct {
	Negative                             bool
	Weeks, Days, Hours, Minutes, Seconds uint
}

func (d *Duration) Decode(params map[string]string, data string) error {
	t := parser.NewStringTokeniser(data)
	if t.Accept("-") {
		d.Negative = true
	} else {
		t.Accept("+")
	}
	if !t.Accept("P") {
		return ErrInvalidDuration
	}
	var level uint8
	for t.Peek() != -1 {
		if t.Accept("T") {
			level = 1
		}
		t.Get()
		mode := t.AcceptRun("0123456789")
		n, err := strconv.ParseUint(t.Get(), 10, 0)
		num := uint(n)
		if err != nil {
			return err
		}
		switch mode {
		case 'W':
			if level > 0 {
				return ErrInvalidDuration
			}
			t.Accept("W")
			if t.Peek() != -1 {
				return ErrInvalidDuration
			}
			d.Weeks = num
			return nil
		case 'D':
			if level > 0 {
				return ErrInvalidDuration
			}
			t.Accept("D")
			d.Days = num
			level = 1
		case 'H':
			if level > 1 {
				return ErrInvalidDuration
			}
			t.Accept("H")
			d.Hours = num
			level = 2
		case 'M':
			if level > 2 {
				return ErrInvalidDuration
			}
			t.Accept("M")
			d.Minutes = num
			level = 3
		case 'S':
			if level > 3 {
				return ErrInvalidDuration
			}
			t.Accept("S")
			if t.Peek() != -1 {
				return ErrInvalidDuration
			}
			d.Seconds = num
		default:
			return ErrInvalidDuration
		}
	}
	if level == 0 {
		return ErrInvalidDuration
	}
	return nil
}

func itoa(n uint) []byte {
	if n == 0 {
		return []byte{'0'}
	}
	var digits [20]byte
	pos := 20
	for ; n > 0; n /= 10 {
		pos--
		digits[pos] = '0' + byte(n%10)
	}
	return digits[pos:]
}

func (d *Duration) Encode(w io.Writer) {
	data := make([]byte, 0, 64)
	if d.Negative {
		data = append(data, '-')
	}
	data = append(data, 'P')
	if d.Weeks != 0 {
		data = append(data, itoa(d.Weeks)...)
		data = append(data, 'W')
	} else {
		if d.Days != 0 {
			data = append(data, itoa(d.Days)...)
			data = append(data, 'D')
		}
		if d.Days == 0 || (d.Hours != 0 || d.Minutes != 0 || d.Seconds != 0) {
			data = append(data, 'T')
			if d.Hours != 0 {
				data = append(data, itoa(d.Hours)...)
				data = append(data, 'H')
				if d.Minutes != 0 || d.Seconds != 0 {
					data = append(data, itoa(d.Minutes)...)
					data = append(data, 'M')
					if d.Seconds != 0 {
						data = append(data, itoa(d.Seconds)...)
						data = append(data, 'S')
					}
				}
			} else if d.Minutes != 0 {
				data = append(data, itoa(d.Minutes)...)
				data = append(data, 'M')
				if d.Seconds != 0 {
					data = append(data, itoa(d.Seconds)...)
					data = append(data, 'S')
				}
			} else {
				data = append(data, itoa(d.Seconds)...)
				data = append(data, 'S')
			}
		}
	}
	w.Write(data)
}

type Float float64

func (f *Float) Decode(params map[string]string, data string) error {
	cf, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return err
	}
	*f = Float(cf)
	return nil
}

func (f *Float) Encode(w io.Writer) {
	w.Write([]byte(strconv.FormatFloat(float64(*f), 'f', -1, 64)))
}

type Integer int32

func (i *Integer) Decode(params map[string]string, data string) error {
	ci, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		return err
	}
	*i = Integer(ci)
	return nil
}

func (i *Integer) Encode(w io.Writer) {
	w.Write([]byte(strconv.FormatInt(int64(*i), 10)))
}

type Period struct {
	Start, End DateTime
	Duration   Duration
}

func (p *Period) Decode(params map[string]string, data string) error {
	i := strings.IndexByte(data, '/')
	if i == -1 || len(data) == i+1 {
		return ErrInvalidPeriod
	}
	err := p.Start.Decode(params, data[:i])
	if err != nil {
		return err
	}
	if data[i+1] == 'P' || data[i+1] == '+' {
		return p.Duration.Decode(params, data[i+1:])
	}
	return p.End.Decode(params, data[i+1:])
}

func (p *Period) Encode(w io.Writer) {
	p.Start.Encode(w)
	w.Write([]byte{'/'})
	if p.End.IsZero() {
		p.Duration.Encode(w)
	} else {
		p.End.Encode(w)
	}
}

type Recur struct {
}

func (r *Recur) Decode(params map[string]string, data string) error {
	return nil
}

func (r *Recur) Encode(w io.Writer) {

}

type Text string

func (t *Text) Decode(params map[string]string, data string) error {
	st := parser.NewStringTokeniser(data)
	d := make([]byte, 0, len(data))
	ru := make([]byte, 4)
Loop:
	for {
		c := st.ExceptRun("\";:\\,^")
		d = append(d, st.Get()...)
		switch c {
		case -1:
			break Loop
		case '\\':
			st.Accept("\\")
			switch st.Peek() {
			case '\\':
				d = append(d, '\\')
			case ';':
				d = append(d, ';')
			case ',':
				d = append(d, ',')
			case 'N', 'n':
				d = append(d, '\n')
			default:
				return ErrInvalidText
			}
			st.Except("")
		case '^':
			st.Accept("^")
			switch c := st.Peek(); c {
			case 'n':
				d = append(d, '\n')
			case -1, '^':
				d = append(d, '^')
			case '\'':
				d = append(d, '"')
			default:
				d = append(d, '^')
				l := utf8.EncodeRune(ru, c)
				d = append(d, ru[:l]...)
			}
			st.Except("")
		default:
			return ErrInvalidText
		}
		st.Get()
	}
	*t = Text(d)
	return nil
}

func (t *Text) Encode(w io.Writer) {
	d := make([]byte, 0, len(*t)+256)
	ru := make([]byte, 4)
	for _, c := range *t {
		switch c {
		case '\\':
			d = append(d, '\\', '\\')
		case '\n':
			d = append(d, '\\', 'n')
		case ';':
			d = append(d, '\\', ';')
		case ',':
			d = append(d, '\\', ',')
		case '^':
			d = append(d, '^', '^')
		case '"':
			d = append(d, '^', '\'')
		default:
			l := utf8.EncodeRune(ru, c)
			d = append(d, ru[:l]...)
		}
	}
	w.Write(d)
}

type Time struct {
	time.Time
}

func (t *Time) Decode(params map[string]string, data string) error {
	if tz, ok := params["TZID"]; ok {
		l, err := time.LoadLocation(tz)
		if err != nil {
			return err
		}
		ct, err := time.ParseInLocation(dateTimeFormat[9:15], data, l)
		if err != nil {
			return err
		}
		t.Time = ct
	} else if len(data) > 0 && data[len(data)-1] == 'Z' {
		ct, err := time.ParseInLocation(dateTimeFormat[9:], data, time.UTC)
		if err != nil {
			return err
		}
		t.Time = ct
	} else {
		ct, err := time.ParseInLocation(dateTimeFormat[9:15], data, time.Local)
		if err != nil {
			return err
		}
		t.Time = ct
	}
	return nil
}

func (t *Time) Encode(w io.Writer) {
	b := make([]byte, 0, 7)
	switch t.Location() {
	case time.UTC:
		b = t.AppendFormat(b, dateTimeFormat[9:])
	case time.Local:
		b = t.AppendFormat(b, dateTimeFormat[9:15])
	default:
		b = t.AppendFormat(b, dateTimeFormat[9:15])
	}
	w.Write(b)
}

type URI struct {
	url.URL
}

func (u *URI) Decode(params map[string]string, data string) error {
	cu, err := url.Parse(data)
	if err != nil {
		return err
	}
	u.URL = *cu
	return nil
}

func (u *URI) Encode(w io.Writer) {
	w.Write([]byte(u.URL.String()))
}

type UTCOffset int

func (u *UTCOffset) Decode(params map[string]string, data string) error {
	t := parser.NewStringTokeniser(data)
	neg := false
	if t.Accept("-") {
		neg = true
	} else {
		t.Accept("+")
	}
	t.Get()
	if !t.Accept("0123456789") || !t.Accept("0123456789") {
		return ErrInvalidOffset
	}
	h, _ := strconv.ParseInt(t.Get(), 10, 32)
	if !t.Accept("0123456789") || !t.Accept("0123456789") {
		return ErrInvalidOffset
	}
	m, _ := strconv.ParseInt(t.Get(), 10, 32)
	if m >= 60 {
		return ErrInvalidOffset
	}
	var s int64
	if t.Accept("0123456789") {
		if !t.Accept("0123456789") || t.Peek() != -1 {
			return ErrInvalidOffset
		}
		s, _ = strconv.ParseInt(t.Get(), 10, 32)
		if s >= 60 {
			return ErrInvalidOffset
		}
	} else if t.Peek() != -1 {
		return ErrInvalidOffset
	}
	*u = UTCOffset(3600*h + 60*m + s)
	if neg {
		if *u == 0 {
			return ErrInvalidOffset
		}
		*u = -(*u)
	}
	return nil
}

func (u *UTCOffset) Encode(w io.Writer) {
	o := int64(*u)
	b := make([]byte, 0, 7)
	if o < 0 {
		b = append(b, '-')
		o = -o
	}
	s := byte(o % 60)
	o /= 60
	m := byte(o % 60)
	h := byte(o / 60)
	if h > 99 {
		h = 0
	}
	b = append(b, '0'+h/10)
	b = append(b, '0'+h%10)
	b = append(b, '0'+m/10)
	b = append(b, '0'+m%10)
	if s > 0 {
		b = append(b, '0'+s/10)
		b = append(b, '0'+s%10)
	}
	w.Write(b)
}

// Errors
var (
	ErrInvalidEncoding = errors.New("invalid binary encoding")
	ErrInvalidPeriod   = errors.New("invalid period")
	ErrInvalidDuration = errors.New("invalid duration")
	ErrInvalidText     = errors.New("invalid encoded text")
	ErrInvalidBoolean  = errors.New("invalid boolean")
	ErrInvalidOffset   = errors.New("invalid offset")
)
