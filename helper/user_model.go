package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {

	return web.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleId:   user.RoleId,
		RoleName: user.RoleName,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
