package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err, "they should be equal")
	assert.Equal(t, 10.0, tax, "they should be equal")

	tax, err = CalculateTax(0)
	assert.Error(t, err)
	assert.Equal(t, 0.0, tax, "they should be equal")
}
