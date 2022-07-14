package service

import (
	"context"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

type BookingService interface {
	Create(ctx context.Context, request web.BookingCreateRequest) web.BookingResponse
	Update(ctx context.Context, request web.BookingUpdateRequest) web.BookingResponse
	Delete(ctx context.Context, bookingId int)
	FindById(ctx context.Context, bookingId int) web.BookingResponse
	FindAll(ctx context.Context) []web.BookingResponse
}
