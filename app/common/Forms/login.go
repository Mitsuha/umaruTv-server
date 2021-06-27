package Forms

import (
	"errors"
)

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (l *LoginForm) Check() error {
	if l.Username == "" || l.Password == "" {
		return errors.New("bad request, username and password is required")
	}

	return nil
}
