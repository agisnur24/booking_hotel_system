package web

type DiscountResponse struct {
	Id           int    `json:"id"`
	EmployeeId   int    `json:"employee_id"`
	HotelId      int    `json:"hotel_id"`
	RoomId       int    `json:"room_id"`
	Rate         string `json:"rate"`
	Status       string `json:"status"`
	RequestDate  string `json:"request_date"`
	EmployeeName string `json:"employee_name"`
	HotelName    string `json:"hotel_name"`
	RoomName     string `json:"room_name"`
}
