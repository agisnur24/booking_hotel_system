package web

type GuestUpdateRequest struct {
	Id           int    `validate:"required"`
	Name         string `validate:"required,max=100,min=1" json:"name"`
	Address      string `validate:"required,max=100,min=1" json:"address"`
	Phone_Number string `validate:"required,min=1,max=100"json:"phone_number"`
	Email        string `validate:"required,min=1,max=100"json:"email"`
}
