package ics

const vEvent = "VEVENT"

type Event struct {
}

func (c *Calendar) decodeEvent(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateStamp:
		case uid:
		case dateTimeStart:
		case class:
		case created:
		case description:
		case geo:
		case lastModified:
		case location:
		case organizer:
		case priority:
		case sequence:
		case status:
		case summary:
		case timeTransparency:
		case url:
		case recurrenceID:
		case recurrenceRule:
		case dateTimeEnd:
		case duration:
		case attach:
		case attendee:
		case categories:
		case comment:
		case contact:
		case exceptionDate:
		case requestStatus:
		case related:
		case resources:
		case recurrenceDate:
		case begin:
			switch p {
			case vAlarm:
			default:
				if err = d.readUnknownComponent(string(p)); err != nil {
					return err
				}
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
