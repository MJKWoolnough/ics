package ics

const vAlarm = "VALARM"

type Alarm struct {
}

func (c *Calendar) decodeAlarm(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case action:
		case trigger:
		case description:
		case summary:
		case attendee:
		case duration:
		case repeat:
		case attach:
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vAlarm {
				return ErrInvalidEnd
			}
			return nil
		}
	}
	return nil
}
