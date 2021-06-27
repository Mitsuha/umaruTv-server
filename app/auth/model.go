package auth

import (
	"umarutv/common/models"
)

type ResponseAfterLogin struct {
	ID     int
	Name   string
	Email  string
	Avatar string
	Token  string
}

func (r *ResponseAfterLogin) LoadFullUser(user *models.User, token string) *ResponseAfterLogin {
	r.Name = user.Name
	r.Avatar = user.Avatar
	r.Token = token
	return r
}
