package ics

const vJournal = "VJournal"

type Journal struct {
}

func (c *Calendar) decodeJournal(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		switch p := p.(type) {
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vJournal {
				return ErrInvalidEnd
			}
			return nil
		}
	}
	return nil
}
