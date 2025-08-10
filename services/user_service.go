package services

import (
	"errors"
	"strings"
	"user-management-golang/model"
	"user-management-golang/utils"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser(name string) []model.UserResponse
	GetUserDetail(id string) (*model.UserResponse, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) GetUser(name string) []model.UserResponse {
	name = strings.ToLower(name)
	var users []model.User

	query := s.db
	if name != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+name+"%")
	}
	query.Find(&users)

	var responses []model.UserResponse
	for _, user := range users {
		response := utils.ToResponse(user)
		responses = append(responses, response)
	}
	return responses
}

func (s *userService) GetUserDetail(id string) (*model.UserResponse, error) {
	var user model.User
	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	response := utils.ToResponse(user)
	return &response, nil
}
