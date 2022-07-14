package web

type InvoiceResponse struct {
	/*Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_Number string `json:"phone_number"`
	Email        string `json:"email"`*/

	Id                       int    `json:"id"`
	Invoice_Number           int    `json:"invoice_number"`
	Employee_Id              int    `json:"employee_id"`
	Meeting_Room_Pricings_Id int    `json:"meeting_room_pricings_id"`
	Discount_Id              int    `json:"discount_id"`
	PIC                      string `json:"pic"`
	Employee_Name            string `json:"employee_name"`
	Price                    int    `json:"price"`
	Price_Type               int    `json:"price_type"`
	Discount_Rate            string `json:"discount_rate"`
	Discount_Status          string `json:"discount_status"`
}
