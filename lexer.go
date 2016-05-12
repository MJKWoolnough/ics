package ics

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode/utf8"

	readerParser "github.com/MJKWoolnough/parser"
)

const (
	ianaTokenChars          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"
	invSafeChars            = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\",:;\x7f"
	invQSafeChars           = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\"\x7f"
	invValueChars           = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x7f"
	paramDelim              = ";"
	paramValueDelim         = "="
	paramMultipleValueDelim = ","
	nameValueDelim          = ":"
	crlf                    = "\r\n"
	dQuote                  = "\""
)

const (
	tokenName readerParser.TokenType = iota + 1
	tokenParamName
	tokenParamValue
	tokenParamQValue
	tokenValue
)

type lexer struct {
	readerParser.Parser
}

func newLexer(r io.Reader) *readerParser.Parser {
	l := readerParser.New(readerParser.NewReaderTokeniser(&unfolder{r: r}))
	l.TokeniserState(lexName)
	return &l
}

func lexName(l *readerParser.Tokeniser) (readerParser.Token, readerParser.TokenFunc) {
	l.AcceptRun(ianaTokenChars)
	t := readerParser.Token{
		tokenName,
		strings.ToUpper(l.Lexeme()),
	}
	if len(t.Data) == 0 {
		if l.Err == io.EOF {
			return l.Done()
		}
		l.Err = ErrNoName
	} else if l.Accept(paramDelim) {
		return t, lexParamName
	} else if l.Accept(nameValueDelim) {
		return t, lexValue
	} else if l.Err == nil {
		l.Err = ErrInvalidChar
	}
	return l.Error()
}

func lexParamName(l *readerParser.Tokeniser) (readerParser.Token, readerParser.TokenFunc) {
	l.Lexeme()
	l.AcceptRun(ianaTokenChars)
	t := readerParser.Token{
		tokenParamName,
		strings.ToUpper(l.Lexeme()),
	}
	if len(t.Data) == 0 {
		l.Err = ErrNoParamName
	} else if !utf8.ValidString(t.Data) {
		l.Err = ErrNotUTF8
	} else if l.Accept(paramValueDelim) {
		return t, lexParamValue
	} else if l.Err == nil {
		l.Err = ErrInvalidChar
	}
	return l.Error()
}

func lexParamValue(l *readerParser.Tokeniser) (readerParser.Token, readerParser.TokenFunc) {
	l.Lexeme()
	var t readerParser.Token
	if l.Accept(dQuote) {
		l.ExceptRun(invQSafeChars)
		if !l.Accept(dQuote) {
			l.Err = ErrInvalidChar
			return l.Error()
		}
		t.Type = tokenParamQValue
		t.Data = l.Lexeme()
		t.Data = string(unescape6868(t.Data[1 : len(t.Data)-1]))
	} else {
		l.ExceptRun(invSafeChars)
		t.Type = tokenParamValue
		t.Data = string(bytes.ToUpper(unescape6868(l.Lexeme())))
	}
	if !utf8.ValidString(t.Data) {
		l.Err = ErrNotUTF8
	} else if l.Accept(paramMultipleValueDelim) {
		return t, lexParamValue
	} else if l.Accept(paramDelim) {
		return t, lexParamName
	} else if l.Accept(nameValueDelim) {

		return t, lexValue
	} else if l.Err == nil {
		l.Err = ErrInvalidChar
	}
	return l.Error()
}

func lexValue(l *readerParser.Tokeniser) (readerParser.Token, readerParser.TokenFunc) {
	l.Lexeme()
	l.ExceptRun(invValueChars)
	if !l.Accept(crlf[:1]) || !l.Accept(crlf[1:]) {
		if l.Err == nil {
			l.Err = ErrInvalidChar
		}
		return l.Error()
	}
	toRet := l.Lexeme()
	toRet = toRet[:len(toRet)-2]
	if !utf8.ValidString(toRet) {
		l.Err = ErrNotUTF8
		return l.Error()
	}
	return readerParser.Token{
		tokenValue,
		toRet,
	}, lexName
}

func clearLine(l *readerParser.Tokeniser) (readerParser.Token, readerParser.TokenFunc) {
	for {
		l.ExceptRun(crlf[:1])
		if l.Err != nil {
			return l.Error()
		} else if l.Accept(crlf[:1]) && l.Accept(crlf[1:]) {
			return lexName(l)
		}
	}
}

// Errors

var (
	ErrInvalidChar = errors.New("invalid character")
	ErrNoName      = errors.New("zero length name")
	ErrNoParamName = errors.New("zero length param name")
	ErrNotUTF8     = errors.New("invalid utf8 string")
)
