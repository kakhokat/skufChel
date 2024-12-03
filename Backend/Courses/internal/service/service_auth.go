package service

import (
	"CoursesBack/internal/models"
	"CoursesBack/internal/store"
	"CoursesBack/internal/utils"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/segmentio/kafka-go"
)

type AuthService struct {
	salt  string
	repo  store.Auth
	kafka *kafka.Conn
}

func NewAuthService(repo store.Auth, salt string, kafka *kafka.Conn) *AuthService {
	return &AuthService{repo: repo, salt: salt, kafka: kafka}
}

func (a *AuthService) SignUp(req models.SignUpRequest) (bool, error) {
	req.Password = utils.HashPass(req.Password)
	randomInt := utils.Random(1000, 9999)

	innerJson, _ := json.Marshal(map[string]interface{}{
		"mail":     req.Email,
		"checkInt": randomInt,
	})

	resultJson, _ := json.Marshal(map[string]interface{}{
		"messageType": "mail",
		"message":     string(innerJson),
	})

	a.kafka.Write(resultJson)

	return a.repo.SignUp(req, randomInt)
}

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func (a *AuthService) SignIn(req models.SignUpRequest) (bool, string, error) {
	req.Password = utils.HashPass(req.Password)
	id, isConfirmed, err := a.repo.SignIn(req.Email, req.Password)

	if err != nil {
		return false, "", err
	}

	if !isConfirmed {
		return isConfirmed, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	}})

	signed, err := token.SignedString([]byte(a.salt))

	return isConfirmed, signed, err
}

func (a *AuthService) CheckKey(key, mail string) (bool, error) {
	return a.repo.CheckKey(key, mail)
}

func (a *AuthService) CheckCreator(id int) (bool, error) {
	return a.repo.CheckCreator(id)
}

func (a *AuthService) GetUserById(id int) (models.User, error) {
	return a.repo.GetUserById(id)
}

func (a *AuthService) SetCreator(id int) (bool, error) {
	return a.repo.SetCreator(id)
}

func (a *AuthService) GetIsConfirmed(id int) (bool, error) {
	return a.repo.GetIsConfirmed(id)
}
