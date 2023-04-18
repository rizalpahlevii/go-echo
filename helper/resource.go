package helper

import (
	"go-echo/model"
	"go-echo/resource"
)

func ToUserResponse(user model.User) resource.UserResource {
	return resource.UserResource{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
func ToUserResponses(users []model.User) []resource.UserResource {
	var userResponses []resource.UserResource
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
