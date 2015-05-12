package ics

import (
	"strings"
	"testing"
)

func TestUnfolder(t *testing.T) {
	u := unfolder{r: strings.NewReader(
		"123" +
			"123\r\n 123" +
			"\r\n12" +
			"3\r\n123\r" + "\n 1" +
			"3\r\n123\r\n" + " 12" +
			"3\r\n123\r\n " + "123" +
			"123\r\n" +
			"123\r\n" +
			"123\r\n",
	)}
	var buf [64]byte

	// Simple Read Test
	n, err := u.Read(buf[:3])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 3 {
		t.Errorf("expecting %d bytes read, got %d", 3, n)
		return
	} else if string(buf[:3]) != "123" {
		t.Errorf("expecting %q, got %q", "123", buf[:3])
	}
	// Simple unfolding test
	n, err = u.Read(buf[:6])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 6 {
		t.Errorf("expecting %d bytes read, got %d", 6, n)
		return
	} else if string(buf[:6]) != "123123" {
		t.Errorf("expecting %q, got %q", "123123", buf[:10])
	}

	// Non-fold line-break test
	n, err = u.Read(buf[:4])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 4 {
		t.Errorf("expecting %d bytes read, got %d", 6, n)
		return
	} else if string(buf[:4]) != "\r\n12" {
		t.Errorf("expecting %q, got %q", "\r\n12", buf[:4])
	}

	// Non-fold line-break + unfold @ '\r' end of buffer
	n, err = u.Read(buf[:7])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 7 {
		t.Errorf("expecting %d bytes read, got %d", 7, n)
		return
	} else if string(buf[:7]) != "3\r\n1231" {
		t.Errorf("expecting %q, got %q", "3\r\n1231", buf[:7])
	}

	// Non-fold line-break + unfold @ '\n' end of buffer
	n, err = u.Read(buf[:8])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 8 {
		t.Errorf("expecting %d bytes read, got %d", 8, n)
		return
	} else if string(buf[:8]) != "3\r\n12312" {
		t.Errorf("expecting %q, got %q", "3\r\n12312", buf[:8])
	}

	// Non-fold line-break + unfold @ ' ' end of buffer
	n, err = u.Read(buf[:9])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 9 {
		t.Errorf("expecting %d bytes read, got %d", 9, n)
		return
	} else if string(buf[:9]) != "3\r\n123123" {
		t.Errorf("expecting %q, got %q", "3\r\n123123", buf[:9])
	}

	// Non-fold line-break, end of buffer
	n, err = u.Read(buf[:5])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 5 {
		t.Errorf("expecting %d bytes read, got %d", 5, n)
		return
	} else if string(buf[:5]) != "123\r\n" {
		t.Errorf("expecting %q, got %q", "123\r\n", buf[:5])
	}

	// Non-fold line-break, end of buffer, after a buffer hold
	n, err = u.Read(buf[:5])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	} else if n != 5 {
		t.Errorf("expecting %d bytes read, got %d", 5, n)
		return
	} else if string(buf[:5]) != "123\r\n" {
		t.Errorf("expecting %q, got %q", "123\r\n", buf[:5])
	}
}
