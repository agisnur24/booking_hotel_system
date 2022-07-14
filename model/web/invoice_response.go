package web

type InvoiceResponse struct {
	/*Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_Number string `json:"phone_number"`
	Email        string `json:"email"`*/

	Id                       int    `json:"Id"`
	Invoice_Number           int    `json:"Invoice_Number"`
	Employee_Id              int    `json:"Employee_Id"`
	Meeting_Room_Pricings_Id int    `json:"Meeting_Room_Pricings_Id"`
	Discount_Id              int    `json:"Discount_Id"`
	PIC                      string `json:"PIC"`
	Employee_Name            string `json:"Employee_Name"`
	Price                    int    `json:"Price"`
	Price_Type               int    `json:"Price_Type"`
	Discount_Rate            string `json:"Discount_Rate"`
	Discount_Status          string `json:"Discount_Status"`
}
