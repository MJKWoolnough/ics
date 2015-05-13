package ics

import (
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
	dquote                  = "\""
)

type token struct {
	typ  tokenType
	data string
}

type tokenType uint8

const (
	tokenError tokenType = iota
	tokenName
	tokenParamName
	tokenParamValue
	tokenParamQValue
	tokenValue
	tokenDone
)

type stateFn func() (token, stateFn)

type lexer struct {
	p     readerParser.Parser
	state stateFn
	err   error
}

func newLexer(r io.Reader) *lexer {
	var l lexer
	l.p = readerParser.NewReaderParser(&unfolder{r: r})
	l.state = l.lexName
	return &l
}

func (l *lexer) GetToken() (token, error) {
	if l.err == io.EOF {
		return token{tokenDone, ""}, l.err
	}
	var t token
	t, l.state = l.state()
	l.p.Get()
	if l.err == io.EOF {
		if t.typ == tokenError {
			l.err = io.ErrUnexpectedEOF
		} else {
			return t, nil
		}
	}
	return t, l.err
}

func (l *lexer) lexName() (token, stateFn) {
	l.p.AcceptRun(ianaTokenChars)
	t := token{
		tokenName,
		strings.ToUpper(l.p.Get()),
	}
	if len(t.data) == 0 {
		if l.err == io.EOF {
			return token{tokenDone, ""}, nil
		}
		l.err = ErrNoName
	} else if l.p.Accept(paramDelim) {
		return t, l.lexParamName
	} else if l.p.Accept(nameValueDelim) {
		return t, l.lexValue
	} else if l.err == nil {
		l.err = ErrInvalidChar
	}
	return l.errorFn()
}

func (l *lexer) lexParamName() (token, stateFn) {
	l.p.AcceptRun(ianaTokenChars)
	t := token{
		tokenParamName,
		strings.ToUpper(l.p.Get()),
	}
	if len(t.data) == 0 {
		l.err = ErrNoParamName
	} else if !utf8.ValidString(t.data) {
		l.err = ErrNotUTF8
	} else if l.p.Accept(paramValueDelim) {
		return t, l.lexParamValue
	} else if l.err == nil {
		l.err = ErrInvalidChar
	}
	return l.errorFn()
}

func (l *lexer) lexParamValue() (token, stateFn) {
	var t token
	if l.p.Accept(dquote) {
		l.p.ExceptRun(invQSafeChars)
		if !l.p.Accept(dquote) {
			l.err = ErrInvalidChar
			return l.errorFn()
		}
		t.typ = tokenParamQValue
		t.data = l.p.Get()
		t.data = unescape6868(t.data[1 : len(t.data)-1])
	} else {
		l.p.ExceptRun(invSafeChars)
		t.typ = tokenParamValue
		t.data = strings.ToUpper(unescape6868(l.p.Get()))
	}
	if !utf8.ValidString(t.data) {
		l.err = ErrNotUTF8
	} else if l.p.Accept(paramMultipleValueDelim) {
		return t, l.lexParamValue
	} else if l.p.Accept(paramDelim) {
		return t, l.lexParamName
	} else if l.p.Accept(nameValueDelim) {
		return t, l.lexValue
	} else if l.err == nil {
		l.err = ErrInvalidChar
	}
	return l.errorFn()
}

func (l *lexer) lexValue() (token, stateFn) {
	l.p.ExceptRun(invValueChars)
	if !l.p.Accept(crlf[:1]) || !l.p.Accept(crlf[1:]) {
		if l.err == nil {
			l.err = ErrInvalidChar
		}
		return l.errorFn()
	}
	toRet := l.p.Get()
	toRet = toRet[:len(toRet)-2]
	if !utf8.ValidString(toRet) {
		l.err = ErrNotUTF8
		return l.errorFn()
	}
	return token{
		tokenValue,
		toRet,
	}, l.lexName
}

func (l *lexer) errorFn() (token, stateFn) {
	return token{
		tokenError,
		l.err.Error(),
	}, l.errorFn
}

func (l *lexer) clearLine() (token, stateFn) {
	for {
		l.p.ExceptRun(crlf[:1])
		if l.err != nil {
			return l.errorFn()
		} else if l.p.Accept(crlf[:1]) && l.p.Accept(crlf[1:]) {
			return l.lexName()
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
