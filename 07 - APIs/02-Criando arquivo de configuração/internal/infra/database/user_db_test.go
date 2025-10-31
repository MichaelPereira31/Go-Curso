package database

import (
	"testing"

	"github.com/go-curso-michaelpereira31/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	// Usar :memory: para banco em memória
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	user, err := entity.NewUser("Jonh", "j@j.com", "123456")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jonh", "j@j.com", "123456")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)


	userFound, _ := userDB.FindByEmail("j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
