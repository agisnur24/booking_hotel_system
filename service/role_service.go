package service

import (
	"context"

	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

type RoleService interface {
	Create(ctx context.Context, request web.RoleCreateRequest) web.RoleResponse
	FindByRoleName(ctx context.Context, roleName string) web.RoleResponse
	FindAll(ctx context.Context) []web.RoleResponse
}
