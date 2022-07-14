package web

type InvoiceUpdateRequest struct {
	Id                       int    `validate:"required"`
	Invoice_Number           int    `validate:"required"json:"invoice_number"`
	Employee_Id              int    `validate:"required"json:"employee_id"`
	Meeting_Room_Pricings_Id int    `validate:"required"json:"meeting_room_pricings_id"`
	Discount_Id              int    `validate:"required"json:"discount_id"`
	PIC                      string `validate:"required,min=1,max=100"json:"pic"`

	/*Name         string `validate:"required,max=100,min=1" json:"name"`
	Address      string `validate:"required,max=100,min=1" json:"address"`
	Phone_Number string `validate:"required,min=1,max=100"json:"phone_number"`
	Email        string `validate:"required,min=1,max=100"json:"email"`*/
}
