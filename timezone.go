package ics

const (
	vTimezone = "VTIMEZONE"
	standard  = "STANDARD"
	daylight  = "DAYLIGHT"
)

type Timezone struct {
}

func (c *Calendar) decodeTimezone(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case timezoneID:
		case lastModified:
		case timezoneURL:
		case begin:
			switch p {
			case standard:
			case daylight:
			default:
				if err = d.readUnknownComponent(string(p)); err != nil {
					return err
				}
			}
		case end:
			if p != vTimezone {
				return ErrInvalidEnd
			}
			return nil
		}
	}
	return nil
}

func (t *Timezone) decodeStandard(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateTimeStart:
		case timezoneOffsetTo:
		case timezoneOffsetFrom:
		case recurrenceRule:
		case comment:
		case recurrenceDate:
		case timezoneName:
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != standard {
				return ErrInvalidEnd
			}
			return nil
		}
	}
}

func (t *Timezone) decodeDaylight(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateTimeStart:
		case timezoneOffsetTo:
		case timezoneOffsetFrom:
		case recurrenceRule:
		case comment:
		case recurrenceDate:
		case timezoneName:
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != daylight {
				return ErrInvalidEnd
			}
			return nil
		}
	}
}
