package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToDiscountResponse(discount domain.Discount) web.DiscountResponse {
	return web.DiscountResponse{
		Id:           discount.Id,
		EmployeeId:   discount.EmployeeId,
		HotelId:      discount.HotelId,
		RoomId:       discount.RoomId,
		Rate:         discount.Rate,
		Status:       discount.Status,
		RequestDate:  discount.RequestDate,
		EmployeeName: discount.EmployeeName,
		HotelName:    discount.HotelName,
		RoomName:     discount.RoomName,
	}
}

func ToDiscountResponses(discounts []domain.Discount) []web.DiscountResponse {

	var discountresponses []web.DiscountResponse
	for _, discount := range discounts {
		discountresponses = append(discountresponses, ToDiscountResponse(discount))
	}
	return discountresponses
}
