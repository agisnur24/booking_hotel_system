package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToBookingResponse(booking domain.Booking) web.BookingResponse {
	return web.BookingResponse{
		Id:              booking.Id,
		MeetingRoomId:   booking.MeetingRoomId,
		InvoiceId:       booking.InvoiceId,
		Status:          booking.Status,
		MeetingRoomName: booking.MeetingRoomName,
		InvoiceNumber:   booking.InvoiceNumber,
		PicName:         booking.PicName,
	}
}

func ToBookingResponses(bookings []domain.Booking) []web.BookingResponse {
	var bookingResponses []web.BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}
	return bookingResponses
}
