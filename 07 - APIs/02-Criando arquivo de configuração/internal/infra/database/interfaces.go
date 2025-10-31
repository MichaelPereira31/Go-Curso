package database

import internal "github.com/go-curso-michaelpereira31/internal/entity"

type UserInterface interface {
	Create(user *internal.User) error
	FindByEmail(email string) (*internal.User, error)
}
