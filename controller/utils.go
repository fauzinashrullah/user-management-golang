package controller

import "user-management-golang/model"

type loginRequest struct {
	Username string
	Password string
}
type userResponse struct {
	ID       uint
	Name     string
	Age      int
	Username string
}

func toResponse(user model.User) userResponse {
	var response userResponse
	response.ID = user.ID
	response.Name = user.Name
	response.Age = user.Age
	response.Username = user.Username
	return response
}
