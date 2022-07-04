package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToRoleResponse(role domain.Role) web.RoleResponse {

	return web.RoleResponse{
		Id:       role.Id,
		RoleName: role.RoleName,
	}
}

func ToRoleResponses(roles []domain.Role) []web.RoleResponse {
	var roleResponses []web.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}
	return roleResponses
}
