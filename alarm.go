package ics

const vAlarm = "VALARM"

type Alarm struct {
}

func (c *Calendar) decodeAlarm(d Decoder) error {
	properties := make([]property, 0, 32)
	var actionType action
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return err
		}
		switch p := p.(type) {
		case action:
			if actionType != actionUnknown {
				return ErrMultipleUnique
			}
			actionType = p
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return err
			}
		case end:
			if p != vAlarm {
				return ErrInvalidEnd
			}
			switch actionType {
			case actionUnknown:
				return ErrRequiredMissing
			case actionAudio:
			case actionDisplay:
			case actionEmail:
			}
			return nil
		default:
			properties = append(properties, p)
		}
	}
	return nil
}

func (c *Calendar) createAudioAlarm(ps []property) error {
	for _, p := range ps {
		switch p := p.(type) {
		case trigger:
		case duration:
		case repeat:
		case attach:
		}
	}
	return nil
}

func (c *Calendar) createDisplayAlarm(ps []property) error {
	for _, p := range ps {
		switch p := p.(type) {
		case description:
		case trigger:
		case duration:
		case repeat:
		}
	}
	return nil
}

func (c *Calendar) createEmailAlarm(ps []property) error {
	for _, p := range ps {
		switch p := p.(type) {
		case description:
		case trigger:
		case summary:
		case attendee:
		case duration:
		case repeat:
		}
	}
	return nil
}
