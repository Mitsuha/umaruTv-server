package auth

import (
	"errors"
	"umarutv/common/dao"
	"umarutv/common/models"
)

type RegisterForm struct {
	Username string
	Email    string
	Password string
}

func (r *RegisterForm) Check() error {
	if r.Username == "" || r.Email == "" || r.Password == "" {
		return errors.New("username, email and password is required")
	}
	if dao.UserByNameExists(r.Username) {
		return errors.New("username is already in use")
	}

	return nil
}

func (r *RegisterForm) GenerateAUser() *models.User {
	return &models.User{
		Name:     r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}
