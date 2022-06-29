package service

import (
	"context"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userEmail string)
	FindByEmail(ctx context.Context, userEmail string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
