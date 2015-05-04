package ics

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
)

const (
	ianaTokenChars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"
	invSafeChars    = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\",:;\x7f"
	invQSafeChars   = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\"\x7f"
	invValueChars   = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x7f"
	paramDelim      = ";"
	paramValueDelim = "="
	nameValueDelim  = ":"
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
)

type stateFn func() (token, stateFn)

type parser struct {
	br                 *bufio.Reader
	buf                bytes.Buffer
	state              stateFn
	pos, lineNo, colNo int
	lastWidth          int
	wasNewLine         bool
	err                error
}

func newParser(r io.Reader) *parser {
	p := &parser{
		br: bufio.NewReader(r),
	}
	p.state = p.parseName
	return p
}

func (p *parser) GetToken() (token, error) {
	var t token
	p.buf.Reset()
	t, p.state = p.state()
	if p.err == io.EOF {
		p.state = p.errorFn
		if t.typ == TokenError {
			p.err = io.ErrUnexpectedEOF
		} else {
			return t, nil
		}
	}
	return t, p.err
}

func (p *parser) next() rune {
	if p.err != nil {
		return -1
	}
	r, s, err := p.br.ReadRune()
	if err != nil {
		p.lastWidth = 0
		p.err = err
		return -1
	}
	p.buf.WriteRune(r)
	p.pos += s
	p.lastWidth = s
	if s > 0 {
		if r == '\n' {
			p.lineNo++
			p.wasNewLine = true
			p.colNo = 0
		} else {
			p.colNo++
			p.wasNewLine = false
		}
	}
	return r
}

func (p *parser) backup() {
	if p.lastWidth > 0 {
		p.pos -= p.lastWidth
		if p.wasNewLine {
			p.lineNo--
		}
		p.colNo--
		p.br.UnreadRune()
		p.buf.Truncate(p.buf.Len() - p.lastWidth)
		p.lastWidth = 0
	}
}

func (p *parser) accept(valid string) bool {
	if strings.ContainsRune(valid, p.next()) {
		return true
	}
	p.backup()
	return false
}

func (p *parser) acceptRun(valid string) {
	for {
		r := p.next()
		if r == -1 {
			return
		}
		if !strings.ContainsRune(valid, r) {
			p.backup()
			return
		}
	}
}

func (p *parser) except(invalid string) bool {
	r := p.next()
	if r == -1 {
		return false
	}
	if !strings.ContainsRune(invalid, r) {
		return true
	}
	p.backup()
	return false
}

func (p *parser) exceptRun(invalid string) {
	for {
		r := p.next()
		if r == -1 {
			return
		}
		if strings.ContainsRune(invalid, r) {
			p.backup()
			return
		}
	}
}

func (p *parser) parseName() (token, stateFn) {
	p.acceptRun(ianaTokenChars)
	t := token{
		TokenName,
		p.buf.String(),
	}
	if p.buf.Len() == 0 {
		p.err = ErrNoName
	} else if p.accept(paramDelim) {
		return t, p.parseParamName
	} else if p.accept(nameValueDelim) {
		return t, p.parseValue
	} else if p.err == nil {
		p.err = ErrInvalidChar
	}
	return p.errorFn()
}

func (p *parser) parseParamName() (token, stateFn) {
	p.acceptRun(ianaTokenChars)
	t := token{
		TokenParamName,
		p.buf.String(),
	}
	if p.buf.Len() == 0 {
		p.err = ErrNoParamName
	} else if p.accept(paramValueDelim) {
		return t, p.parseParamValue
	} else if p.err == nil {
		p.err = ErrInvalidChar
	}
	return p.errorFn()
}

func (p *parser) parseParamValue() (token, stateFn) {
	var t token
	if p.accept("\"") {
		p.exceptRun(invQSafeChars)
		if !p.accept("\"") {
			p.err = ErrInvalidChar
			return p.errorFn()
		}
		t.typ = TokenParamQValue
		t.data = p.buf.String()
	} else {
		p.exceptRun(invSafeChars)
		t.typ = TokenParamValue
		t.data = p.buf.String()
	}
	if p.accept(paramDelim) {
		return t, p.parseParamName
	} else if p.accept(nameValueDelim) {
		return t, p.parseValue
	} else if p.err == nil {
		p.err = ErrInvalidChar
	}
	return p.errorFn()
}

func (p *parser) parseValue() (token, stateFn) {
	var toRet []byte
	for {
		p.exceptRun(invValueChars)
		if !p.accept("\r") || !p.accept("\n") {
			if p.err == nil {
				p.err = ErrInvalidChar
			}
			return p.errorFn()
		}
		toAdd := p.buf.Bytes()
		toRet = append(toRet, toAdd[:len(toAdd)-2]...)
		if !p.accept(" ") {
			break
		}
		p.buf.Reset()
	}
	return token{
		TokenValue,
		string(toRet),
	}, p.parseName
}

func (p *parser) errorFn() (token, stateFn) {
	return token{
		TokenError,
		p.err.Error(),
	}, p.errorFn
}

// Errors

var (
	ErrInvalidChar = errors.New("invalid character")
	ErrNoName      = errors.New("zero length name")
	ErrNoParamName = errors.New("zero length param name")
)
