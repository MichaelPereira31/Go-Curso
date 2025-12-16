package database

import (
	"context"
	"fmt"

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

func (u *UserDB) FindAll(page, limit int, sort string) ([]internal.User, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	var users []internal.User
	var err error
	fmt.Println(page, limit, sort)
	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Order("id " + sort).Find(&users).Error
	} else {
		err = u.DB.Order("id" + sort).Find(&users).Error
	}
	return users, err
}

func (u *UserDB) FindByEmail(email string) (*internal.User, error) {
	var user internal.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) Delete(id string) error {
	ctx := context.Background()
	return u.DB.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&internal.User{}).
		Error
}
