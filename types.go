package ics

import (
	"encoding/base64"
	"errors"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/MJKWoolnough/parser"
)

const dateTimeFormat = "20060102T150405Z"

type Binary []byte

func (b *Binary) decode(params map[string]string, data string) error {
	if params["ENCODING"] != "BASE64" {
		return ErrInvalidEncoding
	}
	cb, err := base64.StdEncoding.DecodeString(data)
	*b = cb
	return err
}

func (b *Binary) aencode(w writer) {
	w.WriteString(";ENCODING=BASE64:")
	b.encode(w)
}

func (b *Binary) encode(w writer) {
	e := base64.NewEncoder(base64.StdEncoding, w)
	e.Write(*b)
	e.Close()
}

func (b *Binary) valid() error {
	return nil
}

type Boolean bool

func (b *Boolean) decode(_ map[string]string, data string) error {
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

func (b *Boolean) aencode(w writer) {
	w.WriteString(":")
	b.encode(w)
}

func (b *Boolean) encode(w writer) {
	if *b {
		w.Write(booleanTrue[:])
	} else {
		w.Write(booleanFalse[:])
	}
}

func (b *Boolean) valid() error {
	return nil
}

type CalendarAddress struct {
	URI
}

type Date struct {
	time.Time
}

func (d *Date) decode(_ map[string]string, data string) error {
	t, err := time.Parse(dateTimeFormat[:8], data)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *Date) aencode(w writer) {
	w.WriteString(":")
	d.encode(w)
}

func (d *Date) encode(w writer) {
	b := make([]byte, 0, 8)
	w.Write(d.AppendFormat(b, dateTimeFormat[:8]))
}

func (d *Date) valid() error {
	if d.IsZero() {
		return ErrInvalidTime
	}
	return nil
}

type DateTime struct {
	time.Time
}

func (d *DateTime) decode(params map[string]string, data string) error {
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

func (d *DateTime) aencode(w writer) {
	writeTimezone(w, d.Time)
	w.WriteString(":")
	d.encode(w)
}

func (d *DateTime) encode(w writer) {
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

func (d *DateTime) valid() error {
	if d.IsZero() {
		return ErrInvalidTime
	}
	return nil
}

type Duration struct {
	Negative                             bool
	Weeks, Days, Hours, Minutes, Seconds uint
}

func (d *Duration) decode(_ map[string]string, data string) error {
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

func (d *Duration) aencode(w writer) {
	w.WriteString(":")
	d.encode(w)
}

func (d *Duration) encode(w writer) {
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

func (d *Duration) valid() error {
	return nil
}

type Float float64

func (f *Float) decode(_ map[string]string, data string) error {
	cf, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return err
	}
	*f = Float(cf)
	return nil
}

func (f *Float) aencode(w writer) {
	w.WriteString(":")
	f.encode(w)
}

func (f *Float) encode(w writer) {
	w.WriteString(strconv.FormatFloat(float64(*f), 'f', -1, 64))
}

func (f *Float) valid() error {
	d := float64(*f)
	if !math.IsNaN(d) && !math.IsInf(d, 0) {
		return ErrInvalidFloat
	}
	return nil
}

type TFloat [2]float64

func (t *TFloat) decode(_ map[string]string, data string) error {
	fs := strings.SplitN(data, ";", 2)
	if len(fs) != 2 {
		return ErrInvalidTFloat
	}
	var err error
	(*t)[0], err = strconv.ParseFloat(fs[0], 64)
	if err != nil {
		return err
	}
	(*t)[1], err = strconv.ParseFloat(fs[1], 64)
	if err != nil {
		return err
	}
	return nil
}

func (t *TFloat) aencode(w writer) {
	w.WriteString(":")
	t.encode(w)
}

func (t *TFloat) encode(w writer) {
	w.WriteString(strconv.FormatFloat((*t)[0], 'f', -1, 64))
	w.WriteString(";")
	w.WriteString(strconv.FormatFloat((*t)[1], 'f', -1, 64))
}

func (t *TFloat) valid() error {
	d := float64((*t)[0])
	if !math.IsNaN(d) && !math.IsInf(d, 0) {
		return ErrInvalidFloat
	}
	d = float64((*t)[1])
	if !math.IsNaN(d) && !math.IsInf(d, 0) {
		return ErrInvalidFloat
	}
	return nil
}

type Integer int32

func (i *Integer) decode(_ map[string]string, data string) error {
	ci, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		return err
	}
	*i = Integer(ci)
	return nil
}

func (i *Integer) aencode(w writer) {
	w.WriteString(":")
	i.encode(w)
}

func (i *Integer) encode(w writer) {
	w.WriteString(strconv.FormatInt(int64(*i), 10))
}

func (i *Integer) valid() error {
	return nil
}

type Period struct {
	Start, End DateTime
	Duration   Duration
}

func (p *Period) decode(params map[string]string, data string) error {
	i := strings.IndexByte(data, '/')
	if i == -1 || len(data) == i+1 {
		return ErrInvalidPeriod
	}
	err := p.Start.decode(params, data[:i])
	if err != nil {
		return err
	}
	if data[i+1] == 'P' || data[i+1] == '+' {
		return p.Duration.decode(params, data[i+1:])
	}
	return p.End.decode(params, data[i+1:])
}

func (p *Period) aencode(w writer) {
	writeTimezone(w, p.Start.Time)
	p.encode(w)
}

func (p *Period) encode(w writer) {
	p.Start.encode(w)
	w.Write([]byte{'/'})
	if p.End.IsZero() {
		p.Duration.encode(w)
	} else {
		p.End.encode(w)
	}
}

func (p *Period) valid() error {
	if p.Start.IsZero() {
		return ErrInvalidPeriodStart
	}
	if p.End.IsZero() {
		if p.Duration.Negative {
			return ErrInvalidPeriodDuration
		}
	} else if !p.End.After(p.Start.Time) || p.Start.Location() != p.End.Location() {
		return ErrInvalidPeriodEnd
	}
	return nil
}

type Frequency uint8

const (
	Secondly Frequency = iota
	Minutely
	Hourly
	Daily
	Weekly
	Monthly
	Yearly
)

type WeekDay uint8

const (
	UnknownDay WeekDay = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Month uint8

const (
	UnknownMonth Month = iota
	January
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

type DayRecur struct {
	Day       WeekDay
	Occurence int8
}

type Recur struct {
	Frequency  Frequency
	Until      time.Time
	UntilTime  bool
	Count      uint64
	Interval   uint64
	BySecond   []uint8
	ByMinute   []uint8
	ByHour     []uint8
	ByDay      []DayRecur
	ByMonthDay []int8
	ByYearDay  []int16
	ByWeekNum  []int8
	ByMonth    []Month
	BySetPos   []int16
	WeekStart  WeekDay
}

func (r *Recur) decode(params map[string]string, data string) error {
	var freq bool
	for _, rule := range strings.Split(data, ";") {
		parts := strings.SplitN(rule, "=", 2)
		if len(parts) != 2 {
			return ErrInvalidRecur
		}
		switch parts[0] {
		case "FREQ":
			switch parts[1] {
			case "SECONDLY":
				r.Frequency = Secondly
			case "MINUTELY":
				r.Frequency = Minutely
			case "HOURLY":
				r.Frequency = Hourly
			case "DAILY":
				r.Frequency = Daily
			case "WEEKLY":
				r.Frequency = Weekly
			case "MONTHLY":
				r.Frequency = Monthly
			case "YEARLY":
				r.Frequency = Yearly
			default:
				return ErrInvalidRecur
			}
			freq = true
		case "UNTIL":
			if r.Count > 0 {
				return ErrInvalidRecur
			}
			if len(parts[1]) > 10 {
				var d DateTime
				if err := d.decode(params, parts[1]); err != nil {
					return ErrInvalidRecur
				}
				r.Until = d.Time
				r.UntilTime = true
			} else {
				var d Date
				if err := d.decode(params, parts[1]); err != nil {
					return ErrInvalidRecur
				}
				r.Until = d.Time
				r.UntilTime = false
			}
		case "COUNT":
			if !r.Until.IsZero() {
				return ErrInvalidRecur
			}
			n, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				return ErrInvalidRecur
			}
			r.Count = n
		case "INTERVAL":
			if r.Interval > 0 {
				return ErrInvalidRecur
			}
			n, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				return ErrInvalidRecur
			}
			r.Interval = n
		case "BYSECOND":
			if r.BySecond != nil {
				return ErrInvalidRecur
			}
			seconds := strings.Split(parts[1], ",")
			secondList := make([]uint8, len(seconds))
			for n, second := range seconds {
				i, err := strconv.ParseUint(second, 10, 8)
				if err != nil || i > 60 {
					return ErrInvalidRecur
				}
				secondList[n] = uint8(i)
			}
			r.BySecond = secondList
		case "BYMINUTE":
			if r.ByMinute != nil {
				return ErrInvalidRecur
			}
			minutes := strings.Split(parts[1], ",")
			minuteList := make([]uint8, len(minutes))
			for n, minute := range minutes {
				i, err := strconv.ParseUint(minute, 10, 8)
				if err != nil || i > 59 {
					return ErrInvalidRecur
				}
				minuteList[n] = uint8(i)
			}
			r.ByMinute = minuteList
		case "BYHOUR":
			if r.ByHour != nil {
				return ErrInvalidRecur
			}
			hours := strings.Split(parts[1], ",")
			hourList := make([]uint8, len(hours))
			for n, hour := range hours {
				i, err := strconv.ParseUint(hour, 10, 8)
				if err != nil || i > 23 {
					return ErrInvalidRecur
				}
				hourList[n] = uint8(i)
			}
			r.ByHour = hourList
		case "BYDAY":
			if r.ByDay != nil {
				return ErrInvalidRecur
			}
			days := strings.Split(parts[1], ",")
			dayList := make([]DayRecur, len(days))
			for n, day := range days {
				neg := false
				numCheck := true
				if len(day) < 2 {
					return ErrInvalidRecur
				}
				if day[0] == '+' {
					day = day[1:]
				} else if day[0] == '-' {
					neg = true
					numCheck = false
					day = day[1:]
					if len(day) < 2 {
						return ErrInvalidRecur
					}
				}
				var num int8
				if day[0] >= '0' && day[0] <= '9' {
					numCheck = true
					num = int8(day[0] - '0')
					day = day[1:]
					if day[0] >= '0' && day[0] <= '9' {
						num *= 10
						num += int8(day[0] - '0')
						day = day[1:]
					}
					if num == 0 || num > 53 {
						return ErrInvalidRecur
					}
					if neg {
						num = -num
					}
				}
				if !numCheck || len(day) != 2 {
					return ErrInvalidRecur
				}
				switch day {
				case "SU":
					dayList[n].Day = Sunday
				case "MO":
					dayList[n].Day = Monday
				case "TU":
					dayList[n].Day = Tuesday
				case "WE":
					dayList[n].Day = Wednesday
				case "TH":
					dayList[n].Day = Thursday
				case "FR":
					dayList[n].Day = Friday
				case "SA":
					dayList[n].Day = Saturday
				default:
					return ErrInvalidRecur
				}
				dayList[n].Occurence = num
			}
			r.ByDay = dayList
		case "BYMONTHDAY":
			if r.ByMonthDay != nil {
				return ErrInvalidRecur
			}
			monthDays := strings.Split(parts[1], ",")
			monthDayList := make([]int8, len(monthDays))
			for n, monthDay := range monthDays {
				i, err := strconv.ParseInt(monthDay, 10, 8)
				if err != nil || i == 0 || i > 31 || i < -31 {
					return ErrInvalidRecur
				}
				monthDayList[n] = int8(i)
			}
			r.ByMonthDay = monthDayList
		case "BYYEARDAY":
			if r.ByYearDay != nil {
				return ErrInvalidRecur
			}
			yearDays := strings.Split(parts[1], ",")
			yearDayList := make([]int16, len(yearDays))
			for n, yearDay := range yearDays {
				i, err := strconv.ParseInt(yearDay, 10, 16)
				if err != nil || i == 0 || i > 366 || i < -366 {
					return ErrInvalidRecur
				}
				yearDayList[n] = int16(i)
			}
			r.ByYearDay = yearDayList
		case "BYWEEKNO":
			if r.ByWeekNum != nil {
				return ErrInvalidRecur
			}
			weekNums := strings.Split(parts[1], ",")
			weekNumList := make([]int8, len(weekNums))
			for n, weekNum := range weekNums {
				i, err := strconv.ParseInt(weekNum, 10, 8)
				if err != nil || i == 0 || i > 53 || i < -53 {
					return ErrInvalidRecur
				}
				weekNumList[n] = int8(i)
			}
			r.ByWeekNum = weekNumList
		case "BYMONTH":
			if r.ByMonth != nil {
				return ErrInvalidRecur
			}
			months := strings.Split(parts[1], ",")
			monthList := make([]Month, len(months))
			for n, month := range months {
				i, err := strconv.ParseUint(month, 10, 8)
				if err != nil || i == 0 || i > 12 {
					return ErrInvalidRecur
				}
				monthList[n] = Month(i)
			}
			r.ByMonth = monthList
		case "BYSETPOS":
			if r.BySetPos != nil {
				return ErrInvalidRecur
			}
			setPoss := strings.Split(parts[1], ",")
			setPosList := make([]int16, len(setPoss))
			for n, setPos := range setPoss {
				i, err := strconv.ParseInt(setPos, 10, 16)
				if err != nil || i == 0 || i > 366 || i < -366 {
					return ErrInvalidRecur
				}
				setPosList[n] = int16(i)
			}
			r.BySetPos = setPosList
		case "WKST":
			if r.WeekStart != UnknownDay {
				return ErrInvalidRecur
			}
			switch parts[1] {
			case "SU":
				r.WeekStart = Sunday
			case "MO":
				r.WeekStart = Monday
			case "TU":
				r.WeekStart = Tuesday
			case "WE":
				r.WeekStart = Wednesday
			case "TH":
				r.WeekStart = Thursday
			case "FR":
				r.WeekStart = Friday
			case "SA":
				r.WeekStart = Saturday
			default:
				return ErrInvalidRecur
			}
		default:
			return ErrInvalidRecur
		}
	}
	if !freq {
		return ErrInvalidRecur
	}
	return nil
}

func (r *Recur) aencode(w writer) {
	writeTimezone(w, r.Until)
	w.WriteString(":")
	r.encode(w)
}

func (r *Recur) encode(w writer) {
	comma := []byte{','}
	switch r.Frequency {
	case Secondly:
		w.WriteString("FREQ=SECONDLY")
	case Minutely:
		w.WriteString("FREQ=MINUTELY")
	case Hourly:
		w.WriteString("FREQ=HOURLY")
	case Daily:
		w.WriteString("FREQ=DAILY")
	case Weekly:
		w.WriteString("FREQ=WEEKLY")
	case Monthly:
		w.WriteString("FREQ=MONTHLY")
	case Yearly:
		w.WriteString("FREQ=YEARLY")
	default:
		w.WriteString("FREQ=SECONDLY")
	}
	if r.Count != 0 {
		w.WriteString(";COUNT=")
		w.WriteString(strconv.FormatUint(r.Count, 10))
	} else if !r.Until.IsZero() {
		w.WriteString(";UNTIL=")
		if r.UntilTime {
			d := DateTime{r.Until}
			d.encode(w)
		} else {
			d := Date{r.Until}
			d.encode(w)
		}
	}
	if r.Interval != 0 {
		w.WriteString(";INTERVAL=")
		w.WriteString(strconv.FormatUint(r.Interval, 10))
	}
	if len(r.BySecond) > 0 {
		w.WriteString(";BYSECOND=")
		for n, second := range r.BySecond {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatUint(uint64(second), 10))
		}
	}
	if len(r.ByMinute) > 0 {
		w.WriteString(";BYMINUTE=")
		for n, minute := range r.ByMinute {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatUint(uint64(minute), 10))
		}
	}
	if len(r.ByHour) > 0 {
		w.WriteString(";BYHOUR=")
		for n, hour := range r.ByHour {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatUint(uint64(hour), 10))
		}
	}
	if len(r.ByDay) > 0 {
		w.WriteString(";BYDAY=")
		for n, day := range r.ByDay {
			if n > 0 {
				w.Write(comma)
			}
			if day.Occurence != 0 {
				w.WriteString(strconv.FormatInt(int64(day.Occurence), 10))
			}
			switch day.Day {
			case Sunday:
				w.Write([]byte{'S', 'U'})
			case Monday:
				w.Write([]byte{'M', 'O'})
			case Tuesday:
				w.Write([]byte{'T', 'U'})
			case Wednesday:
				w.Write([]byte{'W', 'E'})
			case Thursday:
				w.Write([]byte{'T', 'H'})
			case Friday:
				w.Write([]byte{'F', 'R'})
			case Saturday:
				w.Write([]byte{'S', 'A'})
			}
		}
	}
	if len(r.ByMonthDay) > 0 {
		w.WriteString(";BYMONTHDAY=")
		for n, month := range r.ByMonthDay {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatInt(int64(month), 10))
		}
	}
	if len(r.ByYearDay) > 0 {
		w.WriteString(";BYYEARDAY=")
		for n, year := range r.ByYearDay {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatInt(int64(year), 10))
		}
	}
	if len(r.ByWeekNum) > 0 {
		w.WriteString(";BYWEEKNO=")
		for n, week := range r.ByWeekNum {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatInt(int64(week), 10))
		}
	}
	if len(r.ByMonth) > 0 {
		w.WriteString(";BYMONTH=")
		for n, month := range r.ByMonth {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatUint(uint64(month), 10))
		}
	}
	if len(r.BySetPos) > 0 {
		w.WriteString(";BYSETPOS=")
		for n, setPos := range r.BySetPos {
			if n > 0 {
				w.Write(comma)
			}
			w.WriteString(strconv.FormatInt(int64(setPos), 10))
		}
	}
	if r.WeekStart != UnknownDay {
		w.WriteString(";WKST=")
		switch r.WeekStart {
		case Sunday:
			w.Write([]byte{'S', 'U'})
		case Monday:
			w.Write([]byte{'M', 'O'})
		case Tuesday:
			w.Write([]byte{'T', 'U'})
		case Wednesday:
			w.Write([]byte{'W', 'E'})
		case Thursday:
			w.Write([]byte{'T', 'H'})
		case Friday:
			w.Write([]byte{'F', 'R'})
		case Saturday:
			w.Write([]byte{'S', 'A'})
		}
	}
}

func (r *Recur) valid() error {
	switch r.Frequency {
	case Secondly, Minutely, Hourly, Daily, Weekly, Monthly, Yearly:
	default:
		return ErrInvalidRecurFrequency
	}
	for _, second := range r.BySecond {
		if second > 60 {
			return ErrInvalidRecurBySecond
		}
	}
	for _, minute := range r.ByMinute {
		if minute > 59 {
			return ErrInvalidRecurByMinute
		}
	}
	for _, hour := range r.ByHour {
		if hour > 23 {
			return ErrInvalidRecurByHour
		}
	}
	for _, day := range r.ByDay {
		switch day.Day {
		case Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday:
		default:
			return ErrInvalidRecurByDay
		}
	}
	for _, monthDay := range r.ByMonthDay {
		if monthDay == 0 || monthDay > 31 || monthDay < -31 {
			return ErrInvalidRecurByMonthDay
		}
	}
	for _, yearDay := range r.ByYearDay {
		if yearDay == 0 || yearDay > 366 || yearDay < -366 {
			return ErrInvalidRecurByYearDay
		}
	}
	for _, week := range r.ByWeekNum {
		if week == 0 || week > 53 || week < -53 {
			return ErrInvalidRecurByWeekNum
		}
	}
	for _, month := range r.ByMonth {
		if month == 0 || month > 12 {
			return ErrInvalidRecurByMonth
		}
	}
	for _, setPos := range r.BySetPos {
		if setPos == 0 || setPos > 366 || setPos < -366 {
			return ErrInvalidRecurBySetPos
		}
	}
	if r.WeekStart != UnknownDay {
		switch r.WeekStart {
		case Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday:
		default:
			return ErrInvalidRecurWeekStart
		}
	}
	return nil
}

type Text string

func (t *Text) decode(_ map[string]string, data string) error {
	st := parser.NewStringTokeniser(data)
	*t = Text(decodeText(st))
	if st.Peek() != -1 {
		return ErrInvalidText
	}
	return nil
}

func decodeText(t parser.Tokeniser) string {
	var d []byte
	var ru [4]byte
Loop:
	for {
		c := t.ExceptRun("\\,;:")
		d = append(d, t.Get()...)
		switch c {
		case '\\':
			t.Accept("\\")
			switch c := t.Peek(); c {
			case '\\':
				d = append(d, '\\')
			case ';':
				d = append(d, ';')
			case ',':
				d = append(d, ',')
			case 'N', 'n':
				d = append(d, '\n')
			default:
				d = append(d, '\\')
				l := utf8.EncodeRune(ru[:], c)
				d = append(d, ru[:l]...)
			}
			t.Except("")
			t.Get()
		default:
			break Loop
		}
	}
	return string(d)
}

func (t *Text) aencode(w writer) {
	w.WriteString(":")
	t.encode(w)
}

func (t *Text) encode(w writer) {
	d := make([]byte, 0, len(*t)+256)
	var ru [4]byte
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
		default:
			l := utf8.EncodeRune(ru[:], c)
			d = append(d, ru[:l]...)
		}
	}
	w.Write(d)
}

func (t *Text) valid() error {
	return nil
}

type MText []Text

func (t *MText) decode(_ map[string]string, data string) error {
	st := parser.NewStringTokeniser(data)
	for {
		*t = append(*t, Text(decodeText(st)))
		if st.Peek() == -1 {
			break
		}
	}
	return nil
}

func (t *MText) aencode(w writer) {
	w.WriteString(":")
	t.encode(w)
}

func (t *MText) encode(w writer) {
	for n, tx := range *t {
		if n > 0 {
			w.WriteString(",")
			tx.encode(w)
		}
	}
}

func (t *MText) valid() error {
	return nil
}

type Time struct {
	time.Time
}

func (t *Time) decode(params map[string]string, data string) error {
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

func (t *Time) aencode(w writer) {
	writeTimezone(w, t.Time)
	w.WriteString(":")
	t.encode(w)
}

func (t *Time) encode(w writer) {
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

func (t *Time) valid() error {
	if t.IsZero() {
		return ErrInvalidTime
	}
	return nil
}

type URI struct {
	url.URL
}

func (u *URI) decode(_ map[string]string, data string) error {
	cu, err := url.Parse(data)
	if err != nil {
		return err
	}
	u.URL = *cu
	return nil
}

func (u *URI) aencode(w writer) {
	w.WriteString(":")
	u.encode(w)
}

func (u *URI) encode(w writer) {
	w.WriteString(u.URL.String())
}

func (u *URI) valid() error {
	return nil
}

type UTCOffset int

func (u *UTCOffset) decode(_ map[string]string, data string) error {
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
func (u *UTCOffset) aencode(w writer) {
	w.WriteString(":")
	u.encode(w)
}

func (u *UTCOffset) encode(w writer) {
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

func (u *UTCOffset) valid() error {
	return nil
}

func writeTimezone(w writer, t time.Time) {
	switch l := t.Location(); l {
	case time.UTC, time.Local:
	default:
		w.WriteString(";")
		w.WriteString(l.String())
	}
}

// Errors
var (
	ErrInvalidEncoding        = errors.New("invalid Binary encoding")
	ErrInvalidPeriod          = errors.New("invalid Period")
	ErrInvalidDuration        = errors.New("invalid Duration")
	ErrInvalidText            = errors.New("invalid encoded text")
	ErrInvalidBoolean         = errors.New("invalid Boolean")
	ErrInvalidOffset          = errors.New("invalid UTC Offset")
	ErrInvalidRecur           = errors.New("invalid Recur")
	ErrInvalidTime            = errors.New("invalid time")
	ErrInvalidFloat           = errors.New("invalid float")
	ErrInvalidTFloat          = errors.New("invalid number of floats")
	ErrInvalidPeriodStart     = errors.New("invalid start of Period")
	ErrInvalidPeriodDuration  = errors.New("invalid Period duration")
	ErrInvalidPeriodEnd       = errors.New("invalid end of Period")
	ErrInvalidRecurFrequency  = errors.New("invalid Recur frequency")
	ErrInvalidRecurBySecond   = errors.New("invalid Recur BySecond")
	ErrInvalidRecurByMinute   = errors.New("invalid Recur ByMinute")
	ErrInvalidRecurByHour     = errors.New("invalid Recur ByHour")
	ErrInvalidRecurByDay      = errors.New("invalid Recur ByDay")
	ErrInvalidRecurByMonthDay = errors.New("invalid Recur ByMonthDay")
	ErrInvalidRecurByYearDay  = errors.New("invalid Recur ByYearDay")
	ErrInvalidRecurByWeekNum  = errors.New("invalid Recur ByWeekNum")
	ErrInvalidRecurByMonth    = errors.New("invalid Recur ByMonth")
	ErrInvalidRecurBySetPos   = errors.New("invalid Recur BySetPos")
	ErrInvalidRecurWeekStart  = errors.New("invalid Recur WeekStart")
)
