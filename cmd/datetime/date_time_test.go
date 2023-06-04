package datetime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDateTime(t *testing.T) {
	location, _ := time.LoadLocation("America/New_York")
	timeNow = time.Date(2023, time.June, 4, 16, 5, 45, 123456789, location)
	expected := "2023-06-04T20:05:45.123456789Z"

	result := getDateTime()

	assert.Equal(t, expected, result, fmt.Sprintf("DateTime should be in RFC3339 and in UTC format with milliseconds"))
}
