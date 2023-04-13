package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber_Success(t *testing.T) {
	err := IsNumber("1122")

	assert.Nil(t, err)
}

func TestIsNumber_Failure(t *testing.T) {
	err := IsNumber("ghfdyr65")

	assert.NotNil(t, err)
}
