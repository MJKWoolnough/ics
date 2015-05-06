package ics

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
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
	TokenError tokenType = iota
	TokenName
	TokenParamName
	TokenParamValue
	TokenParamQValue
	TokenValue
	TokenDone
)

type stateFn func() (token, stateFn)

type lexer struct {
	br        *bufio.Reader
	buf       bytes.Buffer
	state     stateFn
	lastWidth int
	err       error
}

func newLexer(r io.Reader) *lexer {
	l := &lexer{
		br: bufio.NewReader(r),
	}
	l.state = l.lexName
	return l
}

func (l *lexer) GetToken() (token, error) {
	if l.err == io.EOF {
		return token{TokenDone, ""}, l.err
	}
	var t token
	l.buf.Reset()
	t, l.state = l.state()
	if l.err == io.EOF {
		if t.typ == TokenError {
			l.err = io.ErrUnexpectedEOF
		} else {
			return t, nil
		}
	}
	return t, l.err
}

func (l *lexer) ClearError() {
	if l.err == io.EOF || l.err == io.ErrUnexpectedEOF {
		return
	}
	l.err = nil
	l.state = l.clearLine
}

func (l *lexer) next() rune {
	if l.err != nil {
		return -1
	}
	r, s, err := l.br.ReadRune()
	if err != nil {
		l.lastWidth = 0
		l.err = err
		return -1
	}
	l.buf.WriteRune(r)
	l.lastWidth = s
	return r
}

func (l *lexer) backup() {
	if l.lastWidth > 0 {
		l.br.UnreadRune()
		l.buf.Truncate(l.buf.Len() - l.lastWidth)
		l.lastWidth = 0
	}
}

func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for {
		r := l.next()
		if r == -1 {
			return
		}
		if !strings.ContainsRune(valid, r) {
			l.backup()
			return
		}
	}
}

func (l *lexer) exceptRun(invalid string) {
	for {
		r := l.next()
		if r == -1 {
			return
		}
		if strings.ContainsRune(invalid, r) {
			l.backup()
			return
		}
	}
}

func (l *lexer) lexName() (token, stateFn) {
	l.acceptRun(ianaTokenChars)
	t := token{
		TokenName,
		l.buf.String(),
	}
	if l.buf.Len() == 0 {
		if l.err == io.EOF {
			return token{TokenDone, ""}, nil
		}
		l.err = ErrNoName
	} else if l.accept(paramDelim) {
		return t, l.lexParamName
	} else if l.accept(nameValueDelim) {
		return t, l.lexValue
	} else if l.err == nil {
		l.err = ErrInvalidChar
	}
	return l.errorFn()
}

func (l *lexer) lexParamName() (token, stateFn) {
	l.acceptRun(ianaTokenChars)
	t := token{
		TokenParamName,
		l.buf.String(),
	}
	if l.buf.Len() == 0 {
		l.err = ErrNoParamName
	} else if l.accept(paramValueDelim) {
		return t, l.lexParamValue
	} else if l.err == nil {
		l.err = ErrInvalidChar
	}
	return l.errorFn()
}

func unescape6868(p []byte) []byte {
	u := p[:0]
	for i := 0; i < len(p); i++ {
		if p[i] == '^' && i+1 < len(p) {
			i++
			switch p[i] {
			case 'n':
				u = append(u, '\n') //crlf on windows?
			case '^':
				u = append(u, '^')
			case '\'':
				u = append(u, '"')
			default:
				u = append(u, '^', p[i])
			}
		} else {
			u = append(u, p[i])
		}
	}
	return u
}

func (l *lexer) lexParamValue() (token, stateFn) {
	var t token
	if l.accept(dquote) {
		l.exceptRun(invQSafeChars)
		if !l.accept(dquote) {
			l.err = ErrInvalidChar
			return l.errorFn()
		}
		t.typ = TokenParamQValue
		t.data = string(unescape6868(l.buf.Bytes()[1 : l.buf.Len()-1]))
	} else {
		l.exceptRun(invSafeChars)
		t.typ = TokenParamValue
		t.data = string(unescape6868(l.buf.Bytes()))
	}
	if l.accept(paramMultipleValueDelim) {
		return t, l.lexParamValue
	} else if l.accept(paramDelim) {
		return t, l.lexParamName
	} else if l.accept(nameValueDelim) {
		return t, l.lexValue
	} else if l.err == nil {
		l.err = ErrInvalidChar
	}
	return l.errorFn()
}

func (l *lexer) lexValue() (token, stateFn) {
	var toRet []byte
	for {
		l.exceptRun(invValueChars)
		if !l.accept(crlf[:1]) || !l.accept(crlf[1:]) {
			if l.err == nil {
				l.err = ErrInvalidChar
			}
			return l.errorFn()
		}
		toAdd := l.buf.Bytes()
		toRet = append(toRet, toAdd[:len(toAdd)-2]...)
		if !l.accept(" ") {
			break
		}
		l.buf.Reset()
	}
	return token{
		TokenValue,
		string(toRet),
	}, l.lexName
}

func (l *lexer) errorFn() (token, stateFn) {
	return token{
		TokenError,
		l.err.Error(),
	}, l.errorFn
}

func (l *lexer) clearLine() (token, stateFn) {
	for {
		l.exceptRun(crlf[:1])
		if l.err != nil {
			return l.errorFn()
		} else if l.accept(crlf[:1]) && l.accept(crlf[1:]) {
			return l.lexName()
		}
	}
}

// Errors

var (
	ErrInvalidChar = errors.New("invalid character")
	ErrNoName      = errors.New("zero length name")
	ErrNoParamName = errors.New("zero length param name")
)
