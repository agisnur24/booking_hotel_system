package helper

import (
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
)

func ToInvoiceResponse(invoice domain.Invoice) web.InvoiceResponse {
	return web.InvoiceResponse{
		Id:                       invoice.Id,
		Invoice_Number:           invoice.Invoice_Number,
		Employee_Id:              invoice.Employee_Id,
		Meeting_Room_Pricings_Id: invoice.Meeting_Room_Pricings_Id,
		Discount_Id:              invoice.Discount_Id,
		PIC:                      invoice.PIC,
		Employee_Name:            invoice.Employee_Name,
		Price:                    invoice.Price,
		Price_Type:               invoice.Price_Type,
		Discount_Rate:            invoice.Discount_Rate,
		Discount_Status:          invoice.Discount_Status,
	}
}

func ToInvoiceResponses(invoices []domain.Invoice) []web.InvoiceResponse {

	var invoiceResponses []web.InvoiceResponse
	for _, invoice := range invoices {
		invoiceResponses = append(invoiceResponses, ToInvoiceResponse(invoice))
	}
	return invoiceResponses
}
