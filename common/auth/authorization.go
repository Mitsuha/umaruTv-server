package auth

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"umarutv/common/models"
	"umarutv/database"
)

const _COST = 10

var UnavailableToken = errors.New("token is not available")

func AttemptLogin(user *models.User, password string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("incorrect username or password")
	}

	return GenerateToken(user), nil
}

func GenerateToken(user *models.User) string {
	var buffer = bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buffer)
	_ = encoder.Encode(&user)
	encoded, _ := ioutil.ReadAll(buffer)
	database.Redis.Set(context.Background(), user.Name, encoded, 0)
	return user.Name
}

func Authenticate(token string) (*models.User, error) {
	if token == "" {
		return nil, UnavailableToken
	}
	cmd := database.Redis.Get(context.Background(), token)
	encoded, err := cmd.Bytes()
	if err == redis.Nil {
		return nil, UnavailableToken
	}
	var buffer = bytes.NewBuffer(encoded)
	var decoder = gob.NewDecoder(buffer)
	var user models.User
	err = decoder.Decode(&user)
	if err != nil {
		return nil, UnavailableToken
	}
	err = user.RefreshFromDB()

	if err != nil || user.IsBeingBanned() {
		return nil, UnavailableToken
	}

	return &user, err
}

func EncryptionPassword(password string) string {
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), _COST)

	return string(encrypted)
}
