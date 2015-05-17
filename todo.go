package ics

const vTodo = "VTODO"

type Todo struct {
}

func (c *Calendar) decodeTodo(d Decoder) error {
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case dateStamp:
		case uid:
		case class:
		case completed:
		case created:
		case description:
		case dateTimeStart:
		case geo:
		case lastModified:
		case location:
		case organizer:
		case percent:
		case priority:
		case recurrenceID:
		case sequence:
		case status:
		case summary:
		case url:
		case recurrenceRule:
		case due:
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
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vTodo {
				return ErrInvalidEnd
			}
			return nil
		}
	}
	return nil
}
