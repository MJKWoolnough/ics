package ics

const vFreeBusy = "VFREEBUSY"

type FreeBusy struct {
}

func (c *Calendar) decodeFreeBusy(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateStamp:
		case uid:
		case contact:
		case dateTimeStart:
		case dateTimeEnd:
		case organizer:
		case url:
		case attendee:
		case comment:
		case freebusy:
		case requestStatus:
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
