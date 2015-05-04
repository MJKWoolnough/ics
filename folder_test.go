package ics

import (
	"bytes"
	"testing"
)

func TestFolder(t *testing.T) {
	var buf bytes.Buffer
	f := newFolder(&buf)
	tests := []struct {
		input, output []byte
	}{
		{[]byte("1234"), []byte("1234\r\n")},
		{[]byte("1234567890123456789012345678901234567890123456789012345678901234567890"), []byte("1234567890123456789012345678901234567890123456789012345678901234567890\r\n")},
		{[]byte("123456789012345678901234567890123456789012345678901234567890123456789012345"), []byte("123456789012345678901234567890123456789012345678901234567890123456789012345\r\n")},
		{[]byte("1234567890123456789012345678901234567890123456789012345678901234567890123456"), []byte("123456789012345678901234567890123456789012345678901234567890123456789012345\r\n 6\r\n")},
		{[]byte("123456789012345678901234567890123456789012345678901234567890123456789012345123456789012345678901234567890123456789012345678901234567890123456789012345"), []byte("123456789012345678901234567890123456789012345678901234567890123456789012345\r\n 12345678901234567890123456789012345678901234567890123456789012345678901234\r\n 5\r\n")},
		{[]byte("123456789012345678901234567890123456789012345678901234567890123456789012345123456789012345678901234567890123456789012345678901234567890123456789012345123456789012345678901234567890123456789012345678901234567890123456789012345"), []byte("123456789012345678901234567890123456789012345678901234567890123456789012345\r\n 12345678901234567890123456789012345678901234567890123456789012345678901234\r\n 51234567890123456789012345678901234567890123456789012345678901234567890123\r\n 45\r\n")},
		{[]byte("123456789012345678901234567890123456789012345678901234567890123456789012345£"), []byte("123456789012345678901234567890123456789012345678901234567890123456789012345\r\n £\r\n")},
		{[]byte("12345678901234567890123456789012345678901234567890123456789012345678901234£"), []byte("12345678901234567890123456789012345678901234567890123456789012345678901234\r\n £\r\n")},
	}

	for n, test := range tests {
		buf.Reset()
		err := f.WriteLine(test.input)
		if err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		}
		if bytes.Compare(buf.Bytes(), test.output) != 0 {
			t.Errorf("test %d: expecting: -\n%s\ngot :-\n%s", n+1, test.output, buf.Bytes())
		}
	}
}