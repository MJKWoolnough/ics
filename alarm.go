package ics

const vAlarm = "VALARM"

type Alarm struct {
}

func (c *Calendar) decodeAlarm(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		switch p := p.(type) {
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
