package service

import (
	"context"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

type DiscountService interface {
	Create(ctx context.Context, request web.DiscountCreateRequest) web.DiscountResponse
	Update(ctx context.Context, request web.DiscountUpdateRequest) web.DiscountResponse
	Delete(ctx context.Context, discountId int)
	FindById(ctx context.Context, discountId int) web.DiscountResponse
	FindAll(ctx context.Context) []web.DiscountResponse
}
