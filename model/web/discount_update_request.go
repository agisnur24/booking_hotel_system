package web

type DiscountUpdateRequest struct {
	Id          int    `validate:"required"`
	EmployeeId  int    `validate:"required" json:"employeeid"`
	HotelId     int    `validate:"required" json:"hotelid"`
	RoomId      int    `validate:"required" json:"roomid"`
	Rate        string `validate:"required,min=1,max=10" json:"rate"`
	Status      string `validate:"required,min=1,max=100"json:"status"`
	RequestDate string `validate:"required,min=1,max=100"json:"requestdate"`
}
