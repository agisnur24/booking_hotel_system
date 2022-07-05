package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToGuestResponse(guest domain.Guest) web.GuestResponse {
	return web.GuestResponse{
		Id:           guest.Id,
		Name:         guest.Name,
		Address:      guest.Address,
		Phone_Number: guest.Phone_Number,
		Email:        guest.Email,
	}
}

func ToGuestResponses(guests []domain.Guest) []web.GuestResponse {

	var guestresponses []web.GuestResponse
	for _, guest := range guests {
		guestresponses = append(guestresponses, ToGuestResponse(guest))
	}
	return guestresponses
}
