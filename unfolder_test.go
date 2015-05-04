package ics

import (
	"bytes"
	"testing"
)

func TestUnfolder(t *testing.T) {
	tests := []struct {
		input   []byte
		outputs [][]byte
	}{
		{[]byte("1234\r\n"), [][]byte{[]byte("1234")}},
		{[]byte("1234\r\n 5678\r\n"), [][]byte{[]byte("12345678")}},
		{[]byte("1234\r\n5678\r\n"), [][]byte{[]byte("1234"), []byte("5678")}},
	}
	for n, test := range tests {
		u, _ := newUnfolder(bytes.NewReader(test.input))
		for o, output := range test.outputs {
			got, err := u.ReadLine()
			if err != nil {
				t.Errorf("test %d-%d: unexpected error: %s", n+1, o+1, err)
			}
			if bytes.Compare(got, output) != 0 {
				t.Errorf("test %d-%d: expecting: -\n%s\ngot :-\n%s", n+1, o+1, output, got)
			}
		}
	}
}
