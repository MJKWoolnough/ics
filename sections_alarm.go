package ics

import (
	"errors"
	"io"

	"github.com/MJKWoolnough/parser"
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

type AlarmType interface {
	section
	Type() string
}

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
func (a *Alarm) valid() bool {
	if a.AlarmType == nil {
		return false
	}
	return a.AlarmType.valid()
}

func (AlarmAudio) Type() string {
	return "AUDIO"
}

func (AlarmDisplay) Type() string {
	return "DISPLAY"
}

func (AlarmEmail) Type() string {
	return "EMAIL"
}

// Errors
var (
	ErrInvalidStructure   = errors.New("invalid structure")
	ErrMissingAlarmAction = errors.New("missing alarm action")
)
