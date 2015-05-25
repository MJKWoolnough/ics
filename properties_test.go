package ics

import (
	"reflect"
	"strings"
	"testing"
)

func propertyTester(t *testing.T, testStr string, tests []property) {
	p := newParser(strings.NewReader(testStr))
	for n, test := range tests {
		c, err := p.GetProperty()
		if err != nil {
			t.Errorf("test %d: got unexpected error: %q", n+1, err)
		} else if !reflect.DeepEqual(c, test) {
			t.Errorf("test %d: got %v, expecting %v", n+1, test, c)
		}
	}
}

func TestBeginproperty(t *testing.T) {
	propertyTester(t,
		"BEGIN:HELLO\r\nBEGIN;Ignored=Param:WORLD\r\n",
		[]property{
			begin("HELLO"),
			begin("WORLD"),
		},
	)
}

func TestEndproperty(t *testing.T) {
	propertyTester(t,
		"END;Ignored=Param:WORLD\r\nEND:HELLO\r\n",
		[]property{
			end("WORLD"),
			end("HELLO"),
		},
	)
}

func TestRequestStatusproperty(t *testing.T) {
	propertyTester(t,
		"REQUEST-STATUS:1.02;Almost Successful\r\nREQUEST-STATUS;LANGUAGE=\"EN\":3.94;Client died;Heart-Attack\r\n",
		[]property{
			requestStatus{StatusCode: 102, StatusDescription: "Almost Successful"},
			requestStatus{Language: "EN", StatusCode: 394, StatusDescription: "Client died", Extra: "Heart-Attack"},
		},
	)
}

func TestUnknownproperty(t *testing.T) {
	propertyTester(t,
		"SOMECOMP;LANGUAGE=\"FRENCH\";COLA=CHERRY:SomeValue\r\n",
		[]property{
			propertyData{Name: "SOMECOMP", Params: map[string]attribute{"LANGUAGE": language("FRENCH"), "COLA": unknownParam{token{tokenParamValue, "CHERRY"}}}, Value: "SomeValue"},
		},
	)
}
