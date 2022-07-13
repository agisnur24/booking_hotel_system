package web

type DiscountCreateRequest struct {
	EmployeeId  int    `validate:"required"json:"employee_id"`
	HotelId     int    `validate:"required"json:"hotel_id"`
	RoomId      int    `validate:"required"json:"room_id"`
	Rate        int    `validate:"required"json:"rate"`
	Status      string `validate:"required,min=1,max=100"json:"status"`
	RequestDate string `validate:"required,min=1,max=100"json:"request_date"`
}
