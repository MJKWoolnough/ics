package ics

const vJournal = "VJournal"

type Journal struct {
}

func (c *Calendar) decodeJournal(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateStamp:
		case uid:
		case class:
		case created:
		case dateTimeStart:
		case lastModified:
		case organizer:
		case recurrenceID:
		case sequence:
		case status:
		case summary:
		case url:
		case recurrenceRule:
		case attach:
		case attendee:
		case categories:
		case comment:
		case contact:
		case description:
		case exceptionDate:
		case related:
		case recurrenceDate:
		case requestStatus:
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
