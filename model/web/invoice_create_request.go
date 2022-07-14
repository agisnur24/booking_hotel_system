package web

type InvoiceCreateRequest struct {
	/*Name         string `validate:"required,min=1,max=100"json:"name"`
	Address      string `validate:"required,min=1,max=100"json:"address"`
	Phone_Number string `validate:"required,min=1,max=100"json:"phone_number"`
	Email        string `validate:"required,min=1,max=100"json:"email"`*/

	Invoice_Number           int    `validate:"required"json:"Invoice_Number"`
	Employee_Id              int    `validate:"required"json:"Employee_Id"`
	Meeting_Room_Pricings_Id int    `validate:"required"json:"Meeting_Room_Pricings_Id"`
	Discount_Id              int    `validate:"required"json:"Discount_Id"`
	PIC                      string `validate:"required,min=1,max=100"json:"PIC"`
}
