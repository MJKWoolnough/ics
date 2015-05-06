package ics

import (
	"reflect"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
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
		{"BEEP;PN=PV;QN=\"QV\",RV:Value\r\n  Keeps \r\n Going\r\nTestQuoted;ONE=A^nB^^C^';TWO=\"A^nB^^C^'\":\r\n",
			[]token{
				{TokenName, "BEEP"},
				{TokenParamName, "PN"},
				{TokenParamValue, "PV"},
				{TokenParamName, "QN"},
				{TokenParamQValue, "QV"},
				{TokenParamValue, "RV"},
				{TokenValue, "Value Keeps Going"},
				{TokenName, "TESTQUOTED"},
				{TokenParamName, "ONE"},
				{TokenParamValue, "A\nB^C\""},
				{TokenParamName, "TWO"},
				{TokenParamQValue, "A\nB^C\""},
				{TokenValue, ""},
			},
		},
	}
	for n, test := range tests {
		l := newLexer(strings.NewReader(test.input))
		for o, token := range test.tokens {
			got, err := l.GetToken()
			if err != nil {
				t.Errorf("test %d-%d: unexpected error: %s", n+1, o+1, err)
			} else if !reflect.DeepEqual(got, token) {
				t.Errorf("test %d-%d: expecting :-\n%v\ngot :-\n%v", n+1, o+1, token, got)
			}
		}
	}
}