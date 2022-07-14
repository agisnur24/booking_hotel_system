package domain

type Discount struct {
	Id           int
	EmployeeId   int
	HotelId      int
	RoomId       int
	Rate         string
	Status       string
	RequestDate  string
	EmployeeName string
	HotelName    string
	RoomName     string
}
