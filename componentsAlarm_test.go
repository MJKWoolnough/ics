package ics

import (
	"testing"
	"time"
)

func TestActionComponent(t *testing.T) {
	componentTester(t,
		"ACTION;Ignored=Param:AUDIO\r\nACTION:DISPLAY\r\nACTION:EMAIL\r\nACTION:HAMMER\r\n",
		[]component{
			actionAudio,
			actionDisplay,
			actionEmail,
			actionUnknown,
		},
	)
}

func TestRepeatComponent(t *testing.T) {
	componentTester(t,
		"REPEAT;Ignored=Param:0\r\nREPEAT:453\r\n",
		[]component{
			repeat(0),
			repeat(453),
		},
	)
}

func TestTriggerComponent(t *testing.T) {
	componentTester(t,
		"TRIGGER:P1W\r\nTRIGGER;VALUE=DATE-TIME:20150515T200530Z\r\nTRIGGER;RELATED=END:PT15M10S\r\n",
		[]component{
			trigger{Related: atrStart, Duration: time.Hour * 24 * 7},
			trigger{DateTime: time.Date(2015, 05, 15, 20, 05, 30, 0, time.UTC), Related: -1},
			trigger{Related: atrEnd, Duration: time.Minute*15 + time.Second*10},
		},
	)
}
