package ics

import (
	"errors"
	"io"

	"vimagination.zapto.org/parser"
)

const (
	tokenName parser.TokenType = iota
	tokenParamName
	tokenParamQuotedValue
	tokenParamValue
	tokenValue
)

const phraseContentLine parser.PhraseType = iota

const (
	ianaToken    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"
	nonsafeChars = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x7f\n\";:,"
)

type tokeniser interface {
	GetPhrase() (parser.Phrase, error)
}

func newTokeniser(r io.Reader) *parser.Parser {
	p := parser.New(parser.NewReaderTokeniser(r))
	p.TokeniserState(parseName)
	p.PhraserState(phrase)
	return &p
}

func parseName(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	next := t.AcceptRun(ianaToken)
	name := t.Get()
	if len(name) == 0 {
		if next == -1 {
			return t.Done()
		}
		t.Err = ErrInvalidContentLine
		return t.Error()
	}
	switch next {
	case ';':
		return parser.Token{
			Type: tokenName,
			Data: name,
		}, parseParamName
	case ':':
		return parser.Token{
			Type: tokenName,
			Data: name,
		}, parseValue
	default:
		t.Err = ErrInvalidContentLineName
		return t.Error()
	}
}

func parseParamName(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept(";")
	t.Get()
	if t.AcceptRun(ianaToken) != '=' {
		t.Err = ErrInvalidContentLineParamName
		return t.Error()
	}
	return parser.Token{
		Type: tokenParamName,
		Data: t.Get(),
	}, parseParamValue
}

func parseParamValue(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept("=,")
	t.Get()
	var tk parser.Token
	if t.Accept("\"") {
		t.Get()
		if t.ExceptRun(nonsafeChars[:33]) != '"' {
			t.Err = ErrInvalidContentLineQuotedParamValue
			return t.Error()
		}
		tk.Type = tokenParamQuotedValue
		tk.Data = t.Get()
		t.Accept("\"")
	} else {
		t.ExceptRun(nonsafeChars)
		tk.Type = tokenParamValue
		tk.Data = t.Get()
	}
	switch t.Peek() {
	case ',':
		return tk, parseParamValue
	case ';':
		return tk, parseParamName
	case ':':
		return tk, parseValue
	}
	t.Err = ErrInvalidContentLineParamValue
	return t.Error()
}

func parseValue(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept(":")
	t.Get()
	switch t.ExceptRun(nonsafeChars[:32]) {
	case '\r':
		data := t.Get()
		t.Accept("\r")
		if t.Accept("\n") {
			t.Get()
			return parser.Token{
				Type: tokenValue,
				Data: data,
			}, parseName
		}
	case -1:
		return parser.Token{
			Type: tokenValue,
			Data: t.Get(),
		}, (*parser.Tokeniser).Done
	}
	t.Err = ErrInvalidContentLineValue
	return t.Error()
}

func phrase(p *parser.Parser) (parser.Phrase, parser.PhraseFunc) {
	if !p.Accept(tokenName) {
		if p.Accept(parser.TokenDone) {
			return p.Done()
		}
		return p.Error()
	}
	for p.Accept(tokenParamName) {
		if !p.Accept(tokenParamValue, tokenParamQuotedValue) {
			return p.Error()
		}
		p.AcceptRun(tokenParamValue, tokenParamQuotedValue)
	}
	if !p.Accept(tokenValue) {
		return p.Error()
	}
	return parser.Phrase{
		Type: phraseContentLine,
		Data: p.Get(),
	}, phrase
}

// Errors
var (
	ErrInvalidContentLine                 = errors.New("invalid content line")
	ErrInvalidContentLineName             = errors.New("invalid content line name")
	ErrInvalidContentLineParamName        = errors.New("invalid content line param name")
	ErrInvalidContentLineQuotedParamValue = errors.New("invalid content line quoted param value")
	ErrInvalidContentLineParamValue       = errors.New("invalid content line param value")
	ErrInvalidContentLineValue            = errors.New("invalid content line value")
)
