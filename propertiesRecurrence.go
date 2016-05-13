package ics

import (
	"strconv"
	"strings"
	"time"

	strparse "github.com/MJKWoolnough/parser"
)

type exceptionDate struct {
	dateTime
}

func (p *parser) readExceptionDateProperty() (property, error) {
	as, err := p.readAttributes(tzidparam, valuetypeparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var (
		e exceptionDate
		l *time.Location
	)
	if tzid, ok := as[tzidparam]; ok {
		l, err = time.LoadLocation(tzid.String())
		if err != nil {
			return nil, err
		}
	}
	if val, ok := as[valuetypeparam]; ok && val.(value) == valueDate {
		e.justDate = true
		e.dateTime, err = parseDate(v)
	} else {
		e.dateTime, err = parseDateTime(v, l)
	}
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e exceptionDate) Validate() bool {
	return true
}

func (e exceptionDate) Data() propertyData {
	params := make(map[string]attribute)
	if e.justDate {
		params[valuetypeparam] = valueDate
	}
	if e.Location() != time.UTC {
		params[tzidparam] = timezoneID(e.Location().String())
	}
	return propertyData{
		Name:   exdatep,
		Params: params,
		Value:  e.String(),
	}
}

type recurrenceDate struct {
	JustDate bool
	Periods  []period
}

func (p *parser) readRecurrenceDateProperty() (property, error) {
	as, err := p.readAttributes(tzidparam, valuetypeparam)
	if err != nil {
		return nil, err
	}
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var (
		r recurrenceDate
		l *time.Location
	)
	if tzid, ok := as[tzidparam]; ok {
		l, err = time.LoadLocation(tzid.String())
		if err != nil {
			return nil, err
		}
	}
	if val, ok := as[valuetypeparam]; ok && val.(value) == valuePeriod {
		r.Periods, err = parsePeriods(v, l)
		if err != nil {
			return nil, err
		}
	} else if ok && val.(value) == valueDate {
		r.JustDate = true
		parts := textSplit(v, ',')
		r.Periods = make([]period, 0, len(parts))
		for _, tm := range parts {
			t, err := parseDate(tm)
			if err != nil {
				return nil, err
			}
			r.Periods = append(r.Periods, period{Start: t})
		}
	} else {
		parts := textSplit(v, ',')
		r.Periods = make([]period, 0, len(parts))
		for _, tm := range parts {
			t, err := parseDateTime(tm, l)
			if err != nil {
				return nil, err
			}
			r.Periods = append(r.Periods, period{Start: t})
		}
	}
	return r, nil
}

func (r recurrenceDate) Validate() bool {
	return true
}

func (r recurrenceDate) Data() propertyData {
	params := make(map[string]attribute)
	val := make([]byte, len(r.Periods)*32)
	if r.Periods[0].End.IsZero() {
		if r.JustDate {
			params[valuetypeparam] = valueDate
		}
		for n, p := range r.Periods {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, p.Start.String()...)
		}
	} else {
		params[valuetypeparam] = valuePeriod
		for n, p := range r.Periods {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, p.Bytes()...)
		}
	}
	return propertyData{
		Name:   rdatep,
		Params: params,
		Value:  string(val),
	}
}

const (
	freqSecondly frequency = iota + 1
	freqMinutely
	freqHourly
	freqDaily
	freqWeekly
	freqMonthly
	freqYearly
)

type frequency int

func (f frequency) String() string {
	switch f {
	case freqSecondly:
		return "SECONDLY"
	case freqMinutely:
		return "MINETELY"
	case freqHourly:
		return "HOURLY"
	case freqDaily:
		return "DAILY"
	case freqWeekly:
		return "WEEKLY"
	case freqMonthly:
		return "MONTHLY"
	case freqYearly:
		return "YEARLY"
	default:
		return "UNKNOWN"
	}
}

type recurrenceRule struct {
	Frequency                                                                      frequency
	Until                                                                          dateTime
	Count, Interval                                                                int
	BySecond, ByMinute, ByHour, ByMonthDay, ByYearDay, ByWeekNo, ByMonth, BySetPos []int
	ByDay                                                                          [][2]int
	WeekStart                                                                      int
}

func (p *parser) readRecurrenceRuleProperty() (property, error) {
	v, err := p.readValue()
	if err != nil {
		return nil, err
	}
	var (
		r                                                                                                             recurrenceRule
		freqSet, untilCountSet, intervalSet, bsSet, bmSet, bhSet, bdSet, bmdSet, bydSet, bwnSet, bmoSet, bstSet, wSet bool
	)
	for _, rule := range strings.Split(v, ";") {
		parts := strings.SplitN(rule, "=", 2)
		if len(parts) != 2 {
			return nil, ErrUnsupportedValue
		}
		if !freqSet && parts[0] != "FREQ" {
			return nil, ErrUnsupportedValue
		}
		switch parts[0] {
		case "FREQ":
			if freqSet {
				return nil, ErrInvalidAttributeCombination
			}
			freqSet = true
			switch parts[1] {
			case "SECONDLY":
				r.Frequency = freqSecondly
			case "MINUTELY":
				r.Frequency = freqMinutely
			case "HOURLY":
				r.Frequency = freqHourly
			case "DAILY":
				r.Frequency = freqDaily
			case "WEEKLY":
				r.Frequency = freqWeekly
			case "MONTHLY":
				r.Frequency = freqMonthly
			case "YEARLY":
				r.Frequency = freqYearly
			default:
				return nil, ErrUnsupportedValue
			}
		case "UNTIL":
			if untilCountSet {
				return nil, ErrInvalidAttributeCombination
			}
			untilCountSet = true
			if strings.IndexByte(parts[1], 'T') >= 0 {
				r.Until, err = parseDateTime(parts[1], nil)
			} else {
				r.Until, err = parseDate(parts[1])
			}
			if err != nil {
				return nil, err
			}
		case "COUNT":
			if untilCountSet {
				return nil, ErrInvalidAttributeCombination
			}
			untilCountSet = true
			r.Count, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
		case "INTERVAL":
			if intervalSet {
				return nil, ErrInvalidAttributeCombination
			}
			intervalSet = true
			r.Interval, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
			if r.Interval <= 0 {
				return nil, ErrUnsupportedValue
			}
		case "BYSECOND":
			if bsSet {
				return nil, ErrInvalidAttributeCombination
			}
			bsSet = true
			for _, sec := range strings.Split(parts[1], ",") {
				s, err := strconv.Atoi(sec)
				if err != nil {
					return nil, err
				}
				if s < 0 || s > 60 {
					return nil, ErrUnsupportedValue
				}
				r.BySecond = append(r.BySecond, s)
			}
		case "BYMINUTE":
			if bmSet {
				return nil, ErrInvalidAttributeCombination
			}
			bmSet = true
			for _, min := range strings.Split(parts[1], ",") {
				m, err := strconv.Atoi(min)
				if err != nil {
					return nil, err
				}
				if m < 0 || m > 59 {
					return nil, ErrUnsupportedValue
				}
				r.ByMinute = append(r.ByMinute, m)
			}
		case "BYHOUR":
			if bhSet {
				return nil, ErrInvalidAttributeCombination
			}
			bhSet = true
			for _, hour := range strings.Split(parts[1], ",") {
				h, err := strconv.Atoi(hour)
				if err != nil {
					return nil, err
				}
				if h < 0 || h > 23 {
					return nil, ErrUnsupportedValue
				}
				r.ByHour = append(r.ByHour, h)
			}
		case "BYDAY":
			if bdSet {
				return nil, ErrInvalidAttributeCombination
			}
			bdSet = true
			for _, day := range strings.Split(parts[1], ",") {
				p := strparse.NewStringTokeniser(day)
				var (
					neg, num bool
					n, w     int
				)
				if p.Accept("-") {
					num = true
					neg = true
				} else if p.Accept("+") {
					num = true
				}
				pos := len(p.Get())
				p.AcceptRun("0123456789")
				if p.Len() == 0 {
					if num {
						return nil, ErrUnsupportedValue
					}
				} else {
					numStr := p.Get()
					pos += len(numStr)
					n, _ = strconv.Atoi(numStr)
					if neg {
						n = -n
					}
					if n < -53 || n > 53 || n == 0 {
						return nil, ErrUnsupportedValue
					}
				}
				switch parts[1][pos:] {
				case "SU":
				case "MO":
				case "TU":
				case "WE":
				case "TH":
				case "FR":
				case "SA":
				default:
					return nil, ErrUnsupportedValue
				}
				r.ByDay = append(r.ByDay, [2]int{n, w})
			}
		case "BYMONTHDAY":
			if bmdSet {
				return nil, ErrInvalidAttributeCombination
			}
			bmdSet = true
			for _, monthday := range strings.Split(parts[1], ",") {
				md, err := strconv.Atoi(monthday)
				if err != nil {
					return nil, err
				}
				if md < -31 || md > 31 || md == 0 {
					return nil, ErrUnsupportedValue
				}
				r.ByMonthDay = append(r.ByMonthDay, md)
			}
		case "BYYEARDAY":
			if bydSet {
				return nil, ErrInvalidAttributeCombination
			}
			bydSet = true
			for _, yearday := range strings.Split(parts[1], ",") {
				yd, err := strconv.Atoi(yearday)
				if err != nil {
					return nil, err
				}
				if yd < -366 || yd > 366 || yd == 0 {
					return nil, ErrUnsupportedValue
				}
				r.ByYearDay = append(r.ByYearDay, yd)
			}
		case "BYWEEKNO":
			if bwnSet {
				return nil, ErrInvalidAttributeCombination
			}
			bwnSet = true
			for _, week := range strings.Split(parts[1], ",") {
				w, err := strconv.Atoi(week)
				if err != nil {
					return nil, err
				}
				if w < -53 || w > 53 || w == 0 {
					return nil, ErrUnsupportedValue
				}
				r.ByWeekNo = append(r.ByWeekNo, w)
			}
		case "BYMONTH":
			if bmoSet {
				return nil, ErrInvalidAttributeCombination
			}
			bmoSet = true
			for _, month := range strings.Split(parts[1], ",") {
				m, err := strconv.Atoi(month)
				if err != nil {
					return nil, err
				}
				if m < 1 || m > 12 {
					return nil, ErrUnsupportedValue
				}
				r.ByMonth = append(r.ByMonth, m)
			}
		case "BYSETPOS":
			if bstSet {
				return nil, ErrInvalidAttributeCombination
			}
			bstSet = true
			for _, setpos := range strings.Split(parts[1], ",") {
				sp, err := strconv.Atoi(setpos)
				if err != nil {
					return nil, err
				}
				if sp < -366 || sp > 366 || sp == 0 {
					return nil, ErrUnsupportedValue
				}
				r.BySetPos = append(r.BySetPos, sp)
			}
		case "WKST":
			if wSet {
				return nil, ErrInvalidAttributeCombination
			}
			wSet = true
			r.WeekStart, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}
		default:
		}
	}
	if !intervalSet {
		r.Interval = 1
	}
	if !freqSet {
		return nil, ErrUnsupportedValue
	}
	return r, nil
}

func (r recurrenceRule) Validate() bool {
	return true
}

func (r recurrenceRule) Data() propertyData {
	val := make([]byte, 0, 1024)
	val = append(val, 'F', 'R', 'E', 'Q', '=')
	val = append(val, r.Frequency.String()...)
	if r.Count > 0 {
		val = append(val, ';', 'C', 'O', 'U', 'N', 'T', '=')
		val = append(val, strconv.Itoa(int(r.Count))...)
	} else if !r.Until.IsZero() {
		val = append(val, ';', 'U', 'N', 'T', 'I', 'L', '=')
		val = append(val, r.Until.String()...)
	}
	if r.Interval > 0 {
		val = append(val, ';', 'I', 'N', 'T', 'E', 'R', 'V', 'A', 'L', '=')
		val = append(val, strconv.Itoa(int(r.Interval))...)
	}

	if r.WeekStart > 0 {
		val = append(val, ';', 'W', 'K', 'S', 'T', '=')
		val = append(val, strconv.Itoa(int(r.WeekStart))...)
	}
	if len(r.BySecond) > 0 {
		val = append(val, ';', 'B', 'Y', 'S', 'E', 'C', 'O', 'N', 'D', '=')
		for n, s := range r.BySecond {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(s))...)
		}
	}
	if len(r.ByMinute) > 0 {
		val = append(val, ';', 'B', 'Y', 'M', 'I', 'N', 'U', 'T', 'E', '=')
		for n, m := range r.ByMinute {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(m))...)
		}
	}
	if len(r.ByHour) > 0 {
		val = append(val, ';', 'B', 'Y', 'H', 'O', 'U', 'R', '=')
		for n, h := range r.ByHour {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(h))...)
		}
	}
	if len(r.ByDay) > 0 {
		val = append(val, ';', 'B', 'Y', 'D', 'A', 'Y', '=')
		for n, d := range r.ByDay {
			if n > 0 {
				val = append(val, ',')
			}
			if d[0] != 0 {
				val = append(val, strconv.Itoa(int(d[0]))...)
			}
			switch d[1] {
			case 0:
				val = append(val, 'S', 'U')
			case 1:
				val = append(val, 'M', 'O')
			case 2:
				val = append(val, 'T', 'U')
			case 3:
				val = append(val, 'W', 'E')
			case 4:
				val = append(val, 'T', 'H')
			case 5:
				val = append(val, 'F', 'R')
			case 6:
				val = append(val, 'S', 'A')
			}
		}
	}
	if len(r.ByMonthDay) > 0 {
		val = append(val, ';', 'B', 'Y', 'M', 'O', 'N', 'T', 'H', 'D', 'A', 'Y', '=')
		for n, m := range r.ByMonthDay {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(m))...)
		}
	}
	if len(r.ByYearDay) > 0 {
		val = append(val, ';', 'B', 'Y', 'Y', 'E', 'A', 'R', 'D', 'A', 'Y', '=')
		for n, y := range r.BySecond {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(y))...)
		}
	}
	if len(r.ByWeekNo) > 0 {
		val = append(val, ';', 'B', 'Y', 'W', 'E', 'E', 'K', 'N', 'O', '=')
		for n, w := range r.ByWeekNo {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(w))...)
		}
	}
	if len(r.ByMonth) > 0 {
		val = append(val, ';', 'B', 'Y', 'M', 'O', 'N', 'T', 'H', '=')
		for n, m := range r.ByMonth {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(m))...)
		}
	}
	if len(r.BySetPos) > 0 {
		val = append(val, ';', 'B', 'Y', 'S', 'E', 'T', 'P', 'O', 'S', '=')
		for n, s := range r.BySetPos {
			if n > 0 {
				val = append(val, ',')
			}
			val = append(val, strconv.Itoa(int(s))...)
		}
	}
	return propertyData{
		Name:  rrulep,
		Value: string(val),
	}
}
