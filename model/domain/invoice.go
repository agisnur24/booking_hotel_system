package domain

type Invoice struct {
	Id                       int
	Invoice_Number           int
	Employee_Id              int
	Meeting_Room_Pricings_Id int
	Discount_Id              int
	PIC                      string
	Employee_Name            string
	Price                    int
	Price_Type               int
	Discount_Rate            string
	Discount_Status          string
}
