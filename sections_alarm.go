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

type SectionAlarm struct {
	AlarmType
}

func (a *SectionAlarm) decode(t tokeniser) error {
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
				a.AlarmType = new(SectionAlarmAudio)
			case "DISPLAY":
				a.AlarmType = new(SectionAlarmDisplay)
			case "EMAIL":
				a.AlarmType = new(SectionAlarmEmail)
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

func (a *SectionAlarm) encode(w writer) {
	w.WriteString("BEGIN:VALARM\r\n")
	switch a.AlarmType.(type) {
	case *SectionAlarmAudio:
		w.WriteString("ACTION:AUDIO\r\n")
	case *SectionAlarmDisplay:
		w.WriteString("ACTION:DISPLAY\r\n")
	case *SectionAlarmEmail:
		w.WriteString("ACTION:EMAIL\r\n")
	}
	a.encode(w)
	w.WriteString("END:VALARM\r\n")
}

func (a *SectionAlarm) valid() error {
	switch a.AlarmType.(type) {
	case *SectionAlarmAudio, *SectionAlarmDisplay, *SectionAlarmEmail:
		return a.AlarmType.valid()
	}
	return ErrInvalidAlarm
}

func (SectionAlarmAudio) Type() string {
	return "AUDIO"
}

func (SectionAlarmDisplay) Type() string {
	return "DISPLAY"
}

func (SectionAlarmEmail) Type() string {
	return "EMAIL"
}

// Errors
var (
	ErrInvalidStructure   = errors.New("invalid structure")
	ErrMissingAlarmAction = errors.New("missing alarm action")
	ErrInvalidAlarm       = errors.New("invalid alarm type")
)
