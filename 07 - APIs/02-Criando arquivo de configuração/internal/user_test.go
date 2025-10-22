package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Michael", "d9yHd@example.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Michael", user.Name)
	assert.Equal(t, "d9yHd@example.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Michael", "d9yHd@example.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, user.Password, "123456")
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.False(t, user.ValidatePassword(""))
}
