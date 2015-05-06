package ics

import (
	"errors"
	"regexp"
)

type attribute interface{}

func attributeFromTokens(pn token, pvs []token) (attribute, error) {
	switch pn.data {
	case "ALTREP":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		if pvs[0].typ != tokenParamQValue {
			return nil, ErrIncorrectParamValueType
		}
		return altRep(pvs[0].data), nil
	case "CN":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return commonName(pvs[0].data), nil
	case "CUTYPE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return calendarUserType(pvs[0].data), nil
	case "DELEGATED-FROM":
		vals := make(delegatedFrom, len(pvs))
		for n := range pvs {
			if pvs[n].typ != tokenParamQValue {
				return nil, ErrIncorrectParamValueType
			}
			vals[n] = pvs[n].data
		}
		return vals, nil
	case "DELEGATED-TO":
		vals := make(delegatedTo, len(pvs))
		for n := range pvs {
			if pvs[n].typ != tokenParamQValue {
				return nil, ErrIncorrectParamValueType
			}
			vals[n] = pvs[n].data
		}
		return vals, nil
	case "DIR":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		if pvs[0].typ != tokenParamQValue {
			return nil, ErrIncorrectParamValueType
		}
		return directoryEntryRef(pvs[0].data), nil
	case "ENCODING":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		switch pvs[0].data {
		case "8BIT":
			return encoding8Bit, nil
		case "BASE64":
			return encodingBase64, nil
		}
		return nil, ErrUnknownEncoding
	case "FMTYPE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		if !fmTypeRegex.MatchString(pvs[0].data) {
			return nil, ErrInvalidValue
		}
		return formatType(pvs[0].data), nil
	case "FBTYPE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return freeBusy(pvs[0].data), nil
	case "LANGUAGE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return language(pvs[0].data), nil //should really validate this
	case "MEMBER":
		vals := make(member, len(pvs))
		for n := range pvs {
			if pvs[n].typ != tokenParamQValue {
				return nil, ErrIncorrectParamValueType
			}
			vals[n] = pvs[n].data
		}
		return vals, nil
	case "PARTSTAT":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return partStat(pvs[0].data), nil
	case "RANGE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		switch pvs[0].data {
		case "THISANDFUTURE":
			return rangeThisAndFuture, nil
		case "THISANDPRIOR":
			return rangeThisAndPrior, nil
		}
		return nil, ErrUnknownRange
	case "RELATED":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		switch pvs[0].data {
		case "START":
			return relatedStart, nil
		case "END":
			return relatedEnd, nil
		}
		return nil, ErrUnknownRelated
	case "RELTYPE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return relType(pvs[0].data), nil
	case "ROLE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return role(pvs[0].data), nil
	case "RSVP":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		switch pvs[0].data {
		case "FALSE":
			return rsvp(false), nil
		case "TRUE":
			return rsvp(true), nil
		}
		return nil, ErrUnknownRSVP
	case "SENT-BY":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		if pvs[0].typ != tokenParamQValue {
			return nil, ErrIncorrectParamValueType
		}
		return sentBy(pvs[0].data), nil
	case "TZID":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return timezone(pvs[0].data), nil
	case "VALUE":
		if len(pvs) != 1 {
			return nil, ErrIncorrectNumValues
		}
		return value(pvs[0].data), nil
	default:
		values := make([]string, len(pvs))
		for i := 0; i < len(pvs); i++ {
			values[i] = pvs[i].data
		}
		return unknownAttribute{pn.data, values}, nil
	}
}

type unknownAttribute struct {
	Name   string
	Values []string
}

type altRep string

type commonName string

type calendarUserType string

type delegatedFrom []string

type delegatedTo []string

type directoryEntryRef string

const (
	encoding8Bit   encoding = false
	encodingBase64 encoding = true
)

type encoding bool

var fmTypeRegex *regexp.Regexp

func init() {
	fmTypeRegex = regexp.MustCompile("^[a-zA-Z0-9!#$&.+-^_]{1,127}/[a-zA-Z0-9!#$&.+-^_]{1,127}$")
}

type formatType string

type freeBusy string

type language string

type member []string

type partStat string

const (
	rangeThisAndFuture rangeAttr = false
	rangeThisAndPrior  rangeAttr = true
)

type rangeAttr bool

const (
	relatedStart related = false
	relatedEnd   related = false
)

type related bool

type relType string

type role string

type rsvp bool

type sentBy string

type timezone string

type value string

// Errors

var (
	ErrIncorrectNumValues      = errors.New("incorrect numbers of values for attribute")
	ErrIncorrectParamValueType = errors.New("incorrect param value type")
	ErrUnknownEncoding         = errors.New("unknown encoding type")
	ErrInvalidValue            = errors.New("invalid value")
	ErrUnknownRange            = errors.New("unknown range value")
	ErrUnknownRelated          = errors.New("unknown related value")
	ErrUnknownRSVP             = errors.New("unknown rsvp value")
)
