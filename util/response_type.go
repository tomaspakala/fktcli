package util

import (
	"errors"
	"strings"

	flag "github.com/spf13/pflag"
)

type ResponseType string

const (
	FlagName string = "responseType"

	NormalCase ResponseType = "normalCase"
	UpperCase  ResponseType = "upperCase"
	LowerCase  ResponseType = "lowerCase"
)

func (rt *ResponseType) String() string {
	return string(*rt)
}

func (rt *ResponseType) Set(v string) error {
	switch v {
	case string(UpperCase), string(LowerCase):
		*rt = ResponseType(v)
		return nil
	default:
		return errors.New(`Must be one of "normalCase", "upperCase" or "lowerCase"`)
	}
}

func (rt *ResponseType) Type() string {
	return "ResponseType"
}

func (rt *ResponseType) ToType(v string) string {
	switch *rt {
	case LowerCase:
		return strings.ToLower(v)
	case UpperCase:
		return strings.ToUpper(v)
	default:
		return v
	}
}

func GetValueByResponseType(flags *flag.FlagSet, value string) string {
	respTypeValueString := flags.Lookup(FlagName).Value.String()
	respType := ResponseType(respTypeValueString)
	return respType.ToType(value)
}
