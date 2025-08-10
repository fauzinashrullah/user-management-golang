package services

import (
	"errors"
	"strconv"
	"user-management-golang/config"
	"user-management-golang/model"
	"user-management-golang/utils"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(username, password string) (*[]any, error)
	Register(user model.User) (*model.User, error)
	isUsernameExist(username string) bool
}

type authService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db: db}
}

func (s *authService) Login(username, password string) (*[]any, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if !config.VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, _ := config.GenerateJwt(strconv.FormatUint(uint64(user.ID), 10))
	tokenResponse := map[string]string{"Token": token}

	response := utils.ToResponse(user)

	var responses []any
	responses = append(responses, response)
	responses = append(responses, tokenResponse)
	return &responses, nil
}

func (s *authService) Register(user model.User) (*model.User, error) {
	if s.isUsernameExist(user.Username) {
		return nil, errors.New("username not available")
	}

	hash, err := config.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("failed generate password")
	}
	user.Password = hash
	s.db.Create(&user)

	return &user, nil
}

func (s *authService) isUsernameExist(username string) bool {
	var user model.User
	err := s.db.Where("username = ?", username).First(&user).Error
	return err == nil
}
