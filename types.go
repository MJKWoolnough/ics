package ics

import (
	"encoding/base64"
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

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
	*b = cb
	return err
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
	t, err := time.Parse("20060102", data)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *Date) Encode(w io.Writer) {
	b := make([]byte, 0, 8)
	w.Write([]byte(d.AppendFormat(b, "20060102")))
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
		t, err := time.ParseInLocation("20060102T150405", data, l)
		if err != nil {
			return err
		}
		d.Time = t
	} else if len(data) > 0 && data[len(data)-1] == 'Z' {
		t, err := time.ParseInLocation("20060102T150405Z", data, time.UTC)
		if err != nil {
			return err
		}
		d.Time = t
	} else {
		t, err := time.ParseInLocation("20060102T150405", data, time.Local)
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
		b = d.AppendFormat(b, "20060102T150405Z")
	case time.Local:
		b = d.AppendFormat(b, "20060102T150405")
	default:
		b = d.AppendFormat(b, "20060102T150405")
	}
	w.Write(b)
}

type Duration struct {
	Negative                             bool
	Weeks, Days, Hours, Minutes, Seconds uint
}

func (d *Duration) Decode(params map[string]string, data string) error {
	return nil
}

func itoa(i uint) []byte {

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
		return nil
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
	w.Write([]byte(strconv.FormatFloat(float64(*f), 'f', 64)))
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
	return nil
}

func (t *Text) Encode(w io.Writer) {

}

type Time struct {
	time.Time
}

func (t *Time) Decode(params map[string]string, data string) error {
	return nil
}

func (t *Time) Encode(w io.Writer) {

}

type URI struct {
	url.URL
}

func (u *URI) Decode(params map[string]string, data string) error {
	cu, err := url.Parse(data)
	if err != nil {
		return err
	}
	u.URL = cu
	return nil
}

func (u *URI) Encode(w io.Writer) {
	w.Write([]byte(u.URL.String()))
}

type UTCOffset int

func (u *UTCOffset) Decode(params map[string]string, data string) error {
	return nil
}

func (u *UTCOffset) Encode(w io.Writer) {

}

// Errors
var (
	ErrInvalidEncoding = errors.New("invalid binary encoding")
	ErrInvalidPeriod   = errors.New("invalid period")
)
