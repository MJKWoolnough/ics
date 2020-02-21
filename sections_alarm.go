package ics

import (
	"errors"
	"io"

	"vimagination.zapto.org/parser"
)

type psuedoTokeniser []parser.Phrase

func (p *psuedoTokeniser) GetPhrase() (parser.Phrase, error) {
	if len(*p) == 0 {
		return parser.Phrase{
			Type: parser.PhraseError,
			Data: nil,
		}, io.ErrUnexpectedEOF
	}
	ph := (*p)[0]
	*p = (*p)[1:]
	return ph, nil
}

// AlarmType is an interface this is fulfilled by AlarmAudio, AlarmDisplay and
// AlarmEmail
type AlarmType interface {
	section
	Type() string
}

// Alarm is the encompassing type for the three alarm types
type Alarm struct {
	AlarmType
}

func (a *Alarm) decode(t tokeniser) error {
	var pt psuedoTokeniser
Loop:
	for {
		ph, err := t.GetPhrase()
		if err != nil {
			return err
		}
		pt = append(pt, ph)
		switch ph.Data[0].Data {
		case "BEGIN":
			return ErrInvalidStructure
		case "ACTION":
			if a.AlarmType != nil {
				return ErrInvalidStructure
			}
			switch ph.Data[len(ph.Data)-1].Data {
			case "AUDIO":
				a.AlarmType = new(AlarmAudio)
			case "DISPLAY":
				a.AlarmType = new(AlarmDisplay)
			case "EMAIL":
				a.AlarmType = new(AlarmEmail)
			case "URI":
				a.AlarmType = new(AlarmURI)
			case "NONE":
				a.AlarmType = new(AlarmNone)
			}
		case "END":
			break Loop
		}
	}
	if a.AlarmType == nil {
		return ErrMissingAlarmAction
	}
	return a.AlarmType.decode(&pt)
}

func (a *Alarm) encode(w writer) {
	w.WriteString("BEGIN:VALARM\r\n")
	switch a.AlarmType.(type) {
	case *AlarmAudio:
		w.WriteString("ACTION:AUDIO\r\n")
	case *AlarmDisplay:
		w.WriteString("ACTION:DISPLAY\r\n")
	case *AlarmEmail:
		w.WriteString("ACTION:EMAIL\r\n")
	case *AlarmURI:
		w.WriteString("ACTION:URI\r\n")
	case *AlarmNone:
		w.WriteString("ACTION:NONE\r\n")
	}
	a.AlarmType.encode(w)
	w.WriteString("END:VALARM\r\n")
}

func (a *Alarm) valid() error {
	switch a.AlarmType.(type) {
	case *AlarmAudio, *AlarmDisplay, *AlarmEmail, *AlarmURI, *AlarmNone:
		return a.AlarmType.valid()
	}
	return ErrInvalidAlarm
}

// Type returns the type of the alarm "AUDIO"
func (AlarmAudio) Type() string {
	return "AUDIO"
}

// Type returns the type of the alarm "DISPLAY"
func (AlarmDisplay) Type() string {
	return "DISPLAY"
}

// Type returns the type of the alarm "EMAIL"
func (AlarmEmail) Type() string {
	return "EMAIL"
}

// Type returns the type of the alarm "URI"
func (AlarmURI) Type() string {
	return "URI"
}

// Type returns the type of the alarm "NONE"
func (AlarmNone) Type() string {
	return "NONE"
}

// Errors
var (
	ErrInvalidStructure   = errors.New("invalid structure")
	ErrMissingAlarmAction = errors.New("missing alarm action")
	ErrInvalidAlarm       = errors.New("invalid alarm type")
)
