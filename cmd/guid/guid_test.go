package guid

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetGuid(t *testing.T) {
	result := getGuid()

	_, err := uuid.Parse(result)
	assert.NoError(t, err, fmt.Sprintf("Should return valid guid; Invalid value=%s", result))
}
