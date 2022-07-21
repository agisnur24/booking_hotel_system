package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToInvoiceResponse(invoice domain.Invoice) web.InvoiceResponse {
	return web.InvoiceResponse{
		Id:                   invoice.Id,
		Number:               invoice.Number,
		EmpeloyeeId:          invoice.EmpeloyeeId,
		MeetingRoomPricingId: invoice.MeetingRoomPricingId,
		DiscountId:           invoice.DiscountId,
		Pic:                  invoice.Pic,
		EmpeloyeeName:        invoice.EmpeloyeeName,
		Price:                invoice.Price,
		PriceType:            invoice.PriceType,
		DiscountRate:         invoice.DiscountRate,
		DiscountStatus:       invoice.DiscountStatus,
	}
}

func ToInvoiceResponses(invoices []domain.Invoice) []web.InvoiceResponse {
	var invoiceResponses []web.InvoiceResponse
	for _, invoice := range invoices {
		invoiceResponses = append(invoiceResponses, ToInvoiceResponse(invoice))
	}
	return invoiceResponses
}
