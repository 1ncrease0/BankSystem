package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	salt       = "abobasolenaya12345"
	tokenTTL   = 12 * time.Hour
	signingKey = "aboba"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

type AuthService struct {
	authRepo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{authRepo: repo}
}

// Клиентские методы

func (s *AuthService) EnterpriseSpecialist(id int) (models.EnterpriseSpecialist, error) {
	return s.authRepo.EnterpriseSpecialistById(id)
}

func (s *AuthService) CreateClient(client models.Client) (int, error) {
	client.Password = generatePasswordHash(client.Password)
	return s.authRepo.CreateClient(client)
}

func (s *AuthService) Client(clientId int) (models.Client, error) {
	return s.authRepo.ClientById(clientId)
}

func (s *AuthService) GenerateClientToken(username, password string) (string, error) {
	client, err := s.authRepo.Client(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	userId := client.Id
	role := "client"

	now := time.Now()
	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "FinanceSystem",
			Subject:   fmt.Sprintf("%d", userId),
		},
		UserId: userId,
		Role:   role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

// Методы для банковского сотрудника

func (s *AuthService) CreateBankEmployee(employee models.BankEmployee) (int, error) {
	employee.Password = generatePasswordHash(employee.Password)
	return s.authRepo.CreateBankEmployee(employee)
}

func (s *AuthService) GenerateEmployeeToken(username, password string) (string, error) {
	employee, err := s.authRepo.BankEmployee(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	userId := employee.Id
	role := employee.Role

	now := time.Now()
	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "FinanceSystem",
			Subject:   fmt.Sprintf("%d", userId),
		},
		UserId: userId,
		Role:   role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

// Методы для EnterpriseSpecialist

func (s *AuthService) CreateEnterpriseSpecialist(es models.EnterpriseSpecialist) (int, error) {
	es.Password = generatePasswordHash(es.Password)
	return s.authRepo.CreateEnterpriseSpecialist(es)
}

func (s *AuthService) GenerateEnterpriseSpecialistToken(username, password string) (string, error) {
	// Получаем EnterpriseSpecialist по логину и хешированному паролю
	es, err := s.authRepo.EnterpriseSpecialist(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	userId := es.Id
	role := "enterprise_specialist" // можно изменить в зависимости от требований

	now := time.Now()
	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "FinanceSystem",
			Subject:   fmt.Sprintf("%d", userId),
		},
		UserId: userId,
		Role:   role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
