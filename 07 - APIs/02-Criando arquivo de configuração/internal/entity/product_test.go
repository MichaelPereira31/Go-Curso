package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.0)
	assert.Nil(t, err)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.NotNil(t, product.ID)
}
