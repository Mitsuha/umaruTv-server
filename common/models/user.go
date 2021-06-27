package models

import (
	"time"
	"umarutv/database"
)

type User struct {
	ID        int `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Status    int `json:"status"`
	BanedAt   *time.Time `json:"baned_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const UserTableName = "users"

func (u *User) IsBeingBanned() bool {
	if u.BanedAt == nil {
		return false
	}
	return time.Now().Before(*u.BanedAt)
}

func (u *User) RefreshFromDB() error {
	return database.DB.Where("id", u.ID).Find(u).Error
}
