package ics

import (
	"encoding/base64"
	"errors"
	"strings"
	"time"
)

func valueBinary(i string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(i)
}

func valueBoolean(i string) (bool, error) {
	switch strings.ToUpper(i) {
	case "FALSE":
		return false, nil
	case "TRUE":
		return true, nil
	}
	return false, ErrInvalidBoolean
}

func valueCalendarAddress(i string) (string, error) {
	return i, nil // add format checking?
}

func valueDate(i string) (time.Time, error) {
	return time.Parse("20060102", i)
}

func valueDateTime(i string) (time.Time, error) {
	return time.Now(), nil
}

// Errors

var (
	ErrInvalidBoolean = errors.New("invalid boolean value")
)
