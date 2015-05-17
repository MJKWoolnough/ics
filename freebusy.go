package ics

const vFreeBusy = "VFREEBUSY"

type FreeBusy struct {
}

func (c *Calendar) decodeFreeBusy(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		switch p := p.(type) {
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vFreeBusy {
				return ErrInvalidEnd
			}
			return nil
		}
	}
	return nil
}
