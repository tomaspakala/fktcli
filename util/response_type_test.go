package util

import (
	"testing"

	flag "github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestToType(t *testing.T) {
	tests := []struct {
		input        string
		responseType ResponseType
		want         string
	}{
		{
			input:        "tEsT",
			responseType: NormalCase,
			want:         "tEsT",
		},
		{
			input:        "tEsT",
			responseType: LowerCase,
			want:         "test",
		},
		{
			input:        "tEsT",
			responseType: UpperCase,
			want:         "TEST",
		},
	}

	for _, tt := range tests {
		result := tt.responseType.ToType(tt.input)

		assert.Equal(t, tt.want, result, "Input string should be updated accordingly by %s response type", tt.responseType)
	}
}

func TestGetValueByResponseType(t *testing.T) {
	flagSet := flag.FlagSet{}
	flagSet.String(FlagName, string(LowerCase), "")

	result := GetValueByResponseType(&flagSet, "tEsT")

	assert.Equal(t, "test", result, "Flags should be used to get value according to response type")
}
