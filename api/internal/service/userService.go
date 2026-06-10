package service

import (
	"errors"
	"flowmoney/api/internal/models"
	"flowmoney/api/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(u models.User) (models.User, error)
	Login(email string, password string, secret string) (string, error)
	GetUserById(id int) (models.User, error)
	UpdateBalance(id int, balance float64) (models.User, error)
}

type UserServiceImpl struct {
	Ur repository.UserRepo
}

func NewUserService(ur repository.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{Ur: ur}
}

type Claims struct {
	UserId int
	Email  string
	jwt.RegisteredClaims
}

func (u *UserServiceImpl) CreateUser(user models.User) (models.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashed)

	res, err := u.Ur.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return res, err
}

func (u *UserServiceImpl) Login(email string, password string, secret string) (string, error) {
	res, err := u.Ur.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	comparePass := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if comparePass != nil {
		return "", errors.New("invalid password")
	}

	var c Claims
	c.Email = email
	c.UserId = res.Id
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *UserServiceImpl) GetUserById(id int) (models.User, error) {
	res, err := u.Ur.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (u *UserServiceImpl) UpdateBalance(id int, balance float64) (models.User, error) {
	res, err := u.Ur.UpdateBalance(id, balance)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}
