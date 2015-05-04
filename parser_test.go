package ics

import (
	"reflect"
	"strings"
	"testing"
)

func TestParset(t *testing.T) {
	tests := []struct {
		input  string
		tokens []token
	}{
		{"HELLO:World\r\n",
			[]token{
				{TokenName, "HELLO"},
				{TokenValue, "World"},
			},
		},
		{"BEEP;PN=PV;QN=\"QV\":Value\r\n  Keeps \r\n Going\r\n",
			[]token{
				{TokenName, "BEEP"},
				{TokenParamName, "PN"},
				{TokenParamValue, "PV"},
				{TokenParamName, "QN"},
				{TokenParamQValue, "\"QV\""},
				{TokenValue, "Value Keeps Going"},
			},
		},
	}
	for n, test := range tests {
		p := newParser(strings.NewReader(test.input))
		for o, token := range test.tokens {
			got, err := p.GetToken()
			if err != nil {
				t.Errorf("test %d-%d: unexpected error: %s", n+1, o+1, err)
			} else if !reflect.DeepEqual(got, token) {
				t.Errorf("test %d-%d: expecting :-\n%v\ngot :-\n%v", n+1, o+1, token, got)
			}
		}
	}
}
