package ics

import (
	"reflect"
	"strings"
	"testing"

	readerParser "github.com/MJKWoolnough/parser"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input  string
		tokens []readerParser.Token
	}{
		{"HELLO:World\r\n",
			[]readerParser.Token{
				{tokenName, "HELLO"},
				{tokenValue, "World"},
			},
		},
		{"BEEP;PN=PV;QN=\"QV\",RV:Value\r\n  Keeps \r\n Going\r\nTestQuoted;ONE=A^nB^^C^';TWO=\"A^nB^^C^'\":\r\n",
			[]readerParser.Token{
				{tokenName, "BEEP"},
				{tokenParamName, "PN"},
				{tokenParamValue, "PV"},
				{tokenParamName, "QN"},
				{tokenParamQValue, "QV"},
				{tokenParamValue, "RV"},
				{tokenValue, "Value Keeps Going"},
				{tokenName, "TESTQUOTED"},
				{tokenParamName, "ONE"},
				{tokenParamValue, "A\nB^C\""},
				{tokenParamName, "TWO"},
				{tokenParamQValue, "A\nB^C\""},
				{tokenValue, ""},
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
