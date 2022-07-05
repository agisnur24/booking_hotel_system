package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToHotelResponse(hotel domain.Hotel) web.HotelResponse {
	return web.HotelResponse{
		Id:      hotel.Id,
		Name:    hotel.Name,
		Address: hotel.Address,
		City:    hotel.City,
		ZipCode: hotel.ZipCode,
		Rate:    hotel.Rate,
	}
}

func ToHotelResponses(hotels []domain.Hotel) []web.HotelResponse {
	var hotelResponses []web.HotelResponse
	for _, hotel := range hotels {
		hotelResponses = append(hotelResponses, ToHotelResponse(hotel))
	}
	return hotelResponses
}
