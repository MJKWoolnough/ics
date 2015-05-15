package ics

import (
	"reflect"
	"strings"
	"testing"
)

func componentTester(t *testing.T, testStr string, tests []component) {
	p := newParser(strings.NewReader(testStr))
	for n, test := range tests {
		c, err := p.GetComponent()
		if err != nil {
			t.Errorf("test %d: got unexpected error: %q", n+1, err)
		} else if !reflect.DeepEqual(c, test) {
			t.Errorf("test %d: got %v, expecting %v", n+1, test, c)
		}
	}
}

func TestBeginComponent(t *testing.T) {
	componentTester(t,
		"BEGIN:HELLO\r\nBEGIN;Ignored=Param:WORLD\r\n",
		[]component{
			begin("HELLO"),
			begin("WORLD"),
		},
	)
}

func TestEndComponent(t *testing.T) {
	componentTester(t,
		"END;Ignored=Param:WORLD\r\nEND:HELLO\r\n",
		[]component{
			end("WORLD"),
			end("HELLO"),
		},
	)
}

func TestRequestStatusComponent(t *testing.T) {
	componentTester(t,
		"REQUEST-STATUS:1.02;Almost Successful\r\nREQUEST-STATUS;LANGUAGE=\"EN\":3.94;Client died;Heart-Attack\r\n",
		[]component{
			requestStatus{StatusCode: 102, StatusDescription: "Almost Successful"},
			requestStatus{Language: "EN", StatusCode: 394, StatusDescription: "Client died", Extra: "Heart-Attack"},
		},
	)
}

func TestUnknownComponent(t *testing.T) {
	componentTester(t,
		"SOMECOMP;LANGUAGE=\"FRENCH\";COLA=CHERRY:SomeValue\r\n",
		[]component{
			unknown{Name: "SOMECOMP", Params: map[string]attribute{"LANGUAGE": language("FRENCH"), "COLA": unknownParam{token{tokenParamValue, "CHERRY"}}}, Value: "SomeValue"},
		},
	)
}
