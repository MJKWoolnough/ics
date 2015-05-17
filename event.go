package ics

const vEvent = "VEVENT"

type Event struct {
}

func (c *Calendar) decodeEvent(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		switch p := p.(type) {
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vEvent {
				return ErrInvalidEnd
			}
			return nil
		}
	}
	return nil
}
