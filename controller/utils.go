package controller

import "user-management-golang/model"

type loginRequest struct {
	Name     string
	Password string
}
type userResponse struct {
	ID   uint
	Name string
	Age  int
}

func toResponse(user model.User) userResponse {
	var response userResponse
	response.ID = user.ID
	response.Name = user.Name
	response.Age = user.Age
	return response
}
