package ics

import (
	"reflect"
	"testing"
	"time"
)

func TestEscape(t *testing.T) {
	const want = "Hello\\, World!\\n\\\\/\\\\/\\;"
	if got := string(escape("Hello, World!\n\\/\\/;")); got != want {
		t.Errorf("expecting %q, got %q", want, got)
	}
}

func TestUnescape(t *testing.T) {
	const want = "Hello, World!\n\\/\\/;"
	if got := string(unescape("Hello\\, World!\\n\\\\/\\\\/\\;")); got != want {
		t.Errorf("expecting %q, got %q", want, got)
	}
}

func TestEscape6868(t *testing.T) {
	const want = "^'I'm ^^^^Happy^^^^^'^n^'Me too^'"
	if got := string(escape6868("\"I'm ^^Happy^^\"\n\"Me too\"")); got != want {
		t.Errorf("expecting %q, got %q", want, got)
	}
}

func TestUnescape6868(t *testing.T) {
	const want = "\"I'm ^^Happy^^\"\n\"Me too\""
	if got := string(unescape6868("^'I'm ^^^^Happy^^^^^'^n^'Me too^'")); got != want {
		t.Errorf("expecting %q, got %q", want, got)
	}
}

func TestTextSplit(t *testing.T) {
	tests := []struct {
		input  string
		delim  byte
		output []string
	}{
		{"Hello", ' ', []string{"Hello"}},
		{"Hello, World", ' ', []string{"Hello,", "World"}},
		{"Hello, World", 'o', []string{"Hell", ", W", "rld"}},
	}

	for n, test := range tests {
		got := textSplit(test.input, test.delim)
		if !reflect.DeepEqual(got, test.output) {
			t.Errorf("test %d: expecting %v, got %v", n+1, test.output, got)
		}
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input  string
		output time.Duration
	}{
		{"P2D", time.Hour * 24 * 2},
		{"P3W", time.Hour * 24 * 7 * 3},
		{"PT4H", time.Hour * 4},
		{"PT5M", time.Minute * 5},
		{"PT6S", time.Second * 6},
		{"P3DT6S", time.Hour*24*3 + time.Second*6},
		{"+P3DT6S", time.Hour*24*3 + time.Second*6},
		{"-P3DT6S", -time.Hour*24*3 - time.Second*6},
		{"P1DT2H3M4S", time.Hour*24 + time.Hour*2 + time.Minute*3 + time.Second*4},
	}

	for n, test := range tests {
		got, err := parseDuration(test.input)
		if err != nil {
			t.Errorf("test %d: received unexpected error: %q", n+1, err)
		} else if got != test.output {
			t.Errorf("test %d: expecting %v, got %v", n+1, test.output, got)
		}
	}
}

func TestParseOffset(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"+1200", 12 * 3600},
		{"-1200", -12 * 3600},
		{"+0230", 2*3600 + 30*60},
		{"+023045", 2*3600 + 30*60 + 45},
	}

	for n, test := range tests {
		got, err := parseOffset(test.input)
		if err != nil {
			t.Errorf("test %d: received unexpected error: %q", n+1, err)
		} else if got != test.output {
			t.Errorf("test %d: expecting %d, got %d", n+1, test.output, got)
		}
	}
}
