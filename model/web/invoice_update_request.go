package web

type InvoiceUpdateRequest struct {
	Id                       int    `validate:"required"`
	Invoice_Number           int    `validate:"required"`
	Employee_Id              int    `validate:"required"`
	Meeting_Room_Pricings_Id int    `validate:"required"`
	Discount_Id              int    `validate:"required"`
	PIC                      string `validate:"required,min=1,max=100"json:"PIC"`

	/*Name         string `validate:"required,max=100,min=1" json:"name"`
	Address      string `validate:"required,max=100,min=1" json:"address"`
	Phone_Number string `validate:"required,min=1,max=100"json:"phone_number"`
	Email        string `validate:"required,min=1,max=100"json:"email"`*/
}
