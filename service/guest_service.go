package service

import (
	"context"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

type GuestService interface {
	Create(ctx context.Context, request web.GuestCreateRequest) web.GuestResponse
	Update(ctx context.Context, request web.GuestUpdateRequest) web.GuestResponse
	Delete(ctx context.Context, guestId int)
	FindById(ctx context.Context, guestId int) web.GuestResponse
	FindAll(ctx context.Context) []web.GuestResponse
}
