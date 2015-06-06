package ics

import "github.com/MJKWoolnough/bitmask"

const vAlarm = "VALARM"

type Alarm interface {
	writeAlarmData(*Encoder)
}

func (c *Calendar) decodeAlarm(d Decoder) (Alarm, error) {
	properties := make([]property, 0, 32)
	var actionType action
	for {
		p, err := d.p.GetProperty()
		if err != nil {
			return nil, err
		}
		switch p := p.(type) {
		case action:
			if actionType != actionUnknown {
				return nil, ErrMultipleUnique
			}
			actionType = p
		case begin:
			if err = d.readUnknownComponent(string(p)); err != nil {
				return nil, err
			}
		case end:
			if p != vAlarm {
				return nil, ErrInvalidEnd
			}
			var a Alarm
			switch actionType {
			case actionAudio:
				aa := AudioAlarm{}
				err = aa.decode(properties)
				a = aa
			case actionDisplay:
				ad := DisplayAlarm{}
				err = ad.decode(properties)
				a = ad
			case actionEmail:
				ae := EmailAlarm{}
				err = ae.decode(properties)
				a = ae
			default:
				return nil, ErrRequiredMissing
			}
			if err != nil {
				return nil, err
			}
			return a, nil
		default:
			properties = append(properties, p)
		}
	}
}

type AudioAlarm struct {
	Trigger    trigger
	Duration   duration
	Repeat     repeat
	Attachment attach
}

func (a *AudioAlarm) decode(ps []property) error {
	bm := bitmask.New(4)
	for _, p := range ps {
		switch p := p.(type) {
		case trigger:
			if !bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			a.Trigger = p
		case duration:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			a.Duration = p
		case repeat:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			a.Repeat = p
		case attach:
			if !bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			a.Attachment = p
		}
	}
	if !bm.Get(0) || bm.Get(1) != bm.Get(2) {
		return ErrRequiredMissing
	}
	return nil
}

func (a AudioAlarm) writeAlarmData(e *Encoder) {
	e.writeProperty(actionAudio)
	e.writeProperty(a.Trigger)
	if a.Repeat > 0 {
		e.writeProperty(a.Duration)
		e.writeProperty(a.Repeat)
	}
	if len(a.Attachment.Bytes) != 0 {
		e.writeProperty(a.Attachment)
	}
}

type DisplayAlarm struct {
	Description description
	Trigger     trigger
	Duration    duration
	Repeat      repeat
}

func (d *DisplayAlarm) decode(ps []property) error {
	bm := bitmask.New(4)
	for _, p := range ps {
		switch p := p.(type) {
		case description:
			if !bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			d.Description = p
		case trigger:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			d.Trigger = p
		case duration:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			d.Duration = p
		case repeat:
			if !bm.SetIfNot(3, true) {
				return ErrMultipleUnique
			}
			d.Repeat = p
		}
	}
	if !bm.Get(0) || !bm.Get(1) || bm.Get(2) != bm.Get(3) {
		return ErrRequiredMissing
	}
	return nil
}

func (d DisplayAlarm) writeAlarmData(e *Encoder) {
	e.writeProperty(actionDisplay)
	e.writeProperty(d.Description)
	e.writeProperty(d.Trigger)
	if d.Repeat > 0 {
		e.writeProperty(d.Duration)
		e.writeProperty(d.Repeat)
	}
}

type EmailAlarm struct {
	Description description
	Trigger     trigger
	Summary     summary
	Attendee    attendee
	Duration    duration
	Repeat      repeat
	Attachments []attach
}

func (e *EmailAlarm) decode(ps []property) error {
	bm := bitmask.New(6)
	for _, p := range ps {
		switch p := p.(type) {
		case description:
			if !bm.SetIfNot(0, true) {
				return ErrMultipleUnique
			}
			e.Description = p
		case trigger:
			if !bm.SetIfNot(1, true) {
				return ErrMultipleUnique
			}
			e.Trigger = p
		case summary:
			if !bm.SetIfNot(2, true) {
				return ErrMultipleUnique
			}
			e.Summary = p
		case attendee:
			bm.Set(3, true)
			e.Attendee = p
		case duration:
			if !bm.SetIfNot(4, true) {
				return ErrMultipleUnique
			}
			e.Duration = p
		case repeat:
			if !bm.SetIfNot(5, true) {
				return ErrMultipleUnique
			}
			e.Repeat = p
		case attach:
			e.Attachments = append(e.Attachments, p)
		}
	}
	if !bm.Get(0) || !bm.Get(1) || !bm.Get(2) || !bm.Get(3) || bm.Get(4) != bm.Get(5) || !bm.Get(6) {
		return ErrRequiredMissing
	}
	return nil
}

func (ea EmailAlarm) writeAlarmData(e *Encoder) {
	e.writeProperty(actionEmail)
	e.writeProperty(ea.Description)
	e.writeProperty(ea.Trigger)
	e.writeProperty(ea.Summary)
	e.writeProperty(ea.Attendee)
	if ea.Repeat > 0 {
		e.writeProperty(ea.Duration)
		e.writeProperty(ea.Repeat)
	}
	for _, a := range ea.Attachments {
		e.writeProperty(a)
	}
}
