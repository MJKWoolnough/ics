package ics

import (
	"reflect"
	"strings"
	"testing"

	"github.com/MJKWoolnough/parser"
)

func TestTokeniser(t *testing.T) {
	tests := []struct {
		Input  string
		Output parser.Phrase
		Error  error
	}{
		{
			Input: "Name:Value\r\n",
			Output: parser.Phrase{
				Type: phraseContentLine,
				Data: []parser.Token{
					{
						Type: tokenName,
						Data: "Name",
					},
					{
						Type: tokenValue,
						Data: "Value",
					},
				},
			},
		},
		{
			Input: ":Value\r\n",
			Error: ErrInvalidContentLine,
		},
		{
			Input: "Name+:Value\r\n",
			Error: ErrInvalidContentLineName,
		},
		{
			Input: "Name:!Value HERE09	zø\r\n",
			Output: parser.Phrase{
				Type: phraseContentLine,
				Data: []parser.Token{
					{
						Type: tokenName,
						Data: "Name",
					},
					{
						Type: tokenValue,
						Data: "!Value HERE09	zø",
					},
				},
			},
		},
		{
			Input: "Name:\nValue\r\n",
			Error: ErrInvalidContentLineValue,
		},
		{
			Input: "Name:Va\x00lue\r\n",
			Error: ErrInvalidContentLineValue,
		},
		{
			Input: "Name:Value\x7f\r\n",
			Error: ErrInvalidContentLineValue,
		},
		{
			Input: "Name:Value",
			Error: ErrInvalidContentLineValue,
		},
		{
			Input: "Name;param=paramValue:Value\r\n",
			Output: parser.Phrase{
				Type: phraseContentLine,
				Data: []parser.Token{
					{
						Type: tokenName,
						Data: "Name",
					},
					{
						Type: tokenParamName,
						Data: "param",
					},
					{
						Type: tokenParamValue,
						Data: "paramValue",
					},
					{
						Type: tokenValue,
						Data: "Value",
					},
				},
			},
		},
		{
			Input: "Name;param=\"paramValue\":Value\r\n",
			Output: parser.Phrase{
				Type: phraseContentLine,
				Data: []parser.Token{
					{
						Type: tokenName,
						Data: "Name",
					},
					{
						Type: tokenParamName,
						Data: "param",
					},
					{
						Type: tokenParamQuotedValue,
						Data: "paramValue",
					},
					{
						Type: tokenValue,
						Data: "Value",
					},
				},
			},
		},
		{
			Input: "Name;param=\":;,\":Value\r\n",
			Output: parser.Phrase{
				Type: phraseContentLine,
				Data: []parser.Token{
					{
						Type: tokenName,
						Data: "Name",
					},
					{
						Type: tokenParamName,
						Data: "param",
					},
					{
						Type: tokenParamQuotedValue,
						Data: ":;,",
					},
					{
						Type: tokenValue,
						Data: "Value",
					},
				},
			},
		},
		{
			Input: "Name;param=paramValue1,\"paramValue2\":Value\r\n",
			Output: parser.Phrase{
				Type: phraseContentLine,
				Data: []parser.Token{
					{
						Type: tokenName,
						Data: "Name",
					},
					{
						Type: tokenParamName,
						Data: "param",
					},
					{
						Type: tokenParamValue,
						Data: "paramValue1",
					},
					{
						Type: tokenParamQuotedValue,
						Data: "paramValue2",
					},
					{
						Type: tokenValue,
						Data: "Value",
					},
				},
			},
		},
		{
			Input: "Name;param1=\"ABC\";param2=DEF:Value\r\n",
			Output: parser.Phrase{
				Type: phraseContentLine,
				Data: []parser.Token{
					{
						Type: tokenName,
						Data: "Name",
					},
					{
						Type: tokenParamName,
						Data: "param1",
					},
					{
						Type: tokenParamQuotedValue,
						Data: "ABC",
					},
					{
						Type: tokenParamName,
						Data: "param2",
					},
					{
						Type: tokenParamValue,
						Data: "DEF",
					},
					{
						Type: tokenValue,
						Data: "Value",
					},
				},
			},
		},
	}
	for n, test := range tests {
		p, err := newTokeniser(strings.NewReader(test.Input)).GetPhrase()
		if !reflect.DeepEqual(err, test.Error) {
			t.Errorf("test %d: expecting error %q, got %q", n+1, test.Error, err)
		} else if test.Error == nil && !reflect.DeepEqual(p, test.Output) {
			t.Errorf("test %d: expecting %v, got %v", n+1, test.Output, p)
		}
	}
}
