package ics

const vTimezone = "VTIMEZONE"

type Timezone struct {
}

func (c *Calendar) decodeTimezone(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		switch p := p.(type) {
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
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
