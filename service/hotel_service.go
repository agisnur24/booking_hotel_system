package service

import (
	"context"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

type HotelService interface {
	Create(ctx context.Context, request web.HotelCreateRequest) web.HotelResponse
	Update(ctx context.Context, request web.HotelUpdateRequest) web.HotelResponse
	Delete(ctx context.Context, hotelId int)
	FindById(ctx context.Context, hotelId int) web.HotelResponse
	FindAll(ctx context.Context) []web.HotelResponse
}
