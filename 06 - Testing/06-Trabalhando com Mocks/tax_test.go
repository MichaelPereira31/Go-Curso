package tax

import (
	"errors"
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

func TestCalculateTaxAndSave(t *testing.T) {
	repoMock := &TaxRepository{}
	repoMock.On("SaveTax", 10.0).Return(nil)

	repoMock.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	err := CalculateTaxAndSave(1000.0, repoMock)
	assert.Nil(t, err, "they should be equal")

	err = CalculateTaxAndSave(0, repoMock)
	assert.Error(t, err)

	repoMock.AssertExpectations(t)
}
