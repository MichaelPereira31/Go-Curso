package database

import (
	"fmt"
	"testing"

	"github.com/go-curso-michaelpereira31/internal/entity"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i:=1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*10)
		assert.NoError(t, err)
		db.Create(product)
	}
}

