package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var signature = []byte("mySecre4atSignatur3")

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	NoHp      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email, password, noHp string) (User, error) {
	if name == "" {
		return User{}, errors.New("name cannot be empty")
	}
	if email == "" {
		return User{}, errors.New("email cannot be empty")
	}
	if password == "" {
		return User{}, errors.New("password cannot be empty")
	}
	if len(password) < 7 {
		return User{}, errors.New("password length must be 6 character or more")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		NoHp:     noHp,
	}, nil
}

func (u User) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signature)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func (u User) DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return signature, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}
	if !parsedToken.Valid {
		return data, errors.New("invalid token")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
