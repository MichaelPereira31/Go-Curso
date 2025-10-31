package database

import (
	internal "github.com/go-curso-michaelpereira31/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *internal.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) FindByEmail(email string) (*internal.User, error) {
	var user internal.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}