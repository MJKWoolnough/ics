package ics

import (
	"reflect"
	"strings"
	"testing"
)

func TestParserParts(t *testing.T) {
	p := newParser(strings.NewReader("NAME;Param=Pvalue1,\"Pvalue2\";Oparam=\"qValue\":VAL\r\n UE\r\nNAME;Ignored=THIS;And=\"this\":But not this\r\nParamTest;GetThisParam=VAL;ANDThisParam=VAL;ButNotThisOne=NO:\r\n"))
	n, err := p.readName()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if n != "NAME" {
		t.Errorf("expecting name %q, got %q", "NAME", n)
	}
	atts, err := p.readAttributes("PARAM", "OPARAM")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	exp := map[string]attribute{
		"PARAM": unknownParam{
			{tokenParamValue, "PVALUE1"},
			{tokenParamQValue, "Pvalue2"},
		},
		"OPARAM": unknownParam{
			{tokenParamQValue, "qValue"},
		},
	}
	if !reflect.DeepEqual(atts, exp) {
		t.Errorf("expecting attributes %v, got %v", exp, atts)
		return
	}
	val, err := p.readValue()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if val != "VALUE" {
		t.Errorf("expecting value %q, got %q", "VALUE", val)
		return
	}
	n, err = p.readName()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if n != "NAME" {
		t.Errorf("expecting name %q, got %q", "NAME", n)
	}
	val, err = p.readValue()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if val != "But not this" {
		t.Errorf("expecting value %q, got %q", "But not this", val)
		return
	}
	n, err = p.readName()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if n != "PARAMTEST" {
		t.Errorf("expecting name %q, got %q", "NAME", n)
	}
	atts, err = p.readAttributes("GETTHISPARAM", "ANDTHISPARAM")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	exp = map[string]attribute{
		"GETTHISPARAM": unknownParam{
			{tokenParamValue, "VAL"},
		},
		"ANDTHISPARAM": unknownParam{
			{tokenParamValue, "VAL"},
		},
	}
	if !reflect.DeepEqual(atts, exp) {
		t.Errorf("expecting attributes %v, got %v", exp, atts)
		return
	}
	val, err = p.readValue()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	if val != "" {
		t.Errorf("expecting value %q, got %q", "", val)
		return
	}
}
