package utils

import "user-management-golang/model"

func ToResponse(user model.User) model.UserResponse {
	var response model.UserResponse
	response.ID = user.ID
	response.Name = user.Name
	response.Age = user.Age
	response.Username = user.Username
	return response
}

func ToEntity(req model.RegisterRequest) model.User {
	return model.User{
		Name:     req.Name,
		Age:      req.Age,
		Username: req.Username,
		Password: req.Password}
}
