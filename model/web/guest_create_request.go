package web

type GuestCreateRequest struct {
	Name         string `validate:"required,min=1,max=100"json:"name"`
	Address      string `validate:"required,min=1,max=100"json:"address"`
	Phone_Number string `validate:"required,min=1,max=100"json:"phone_number"`
	Email        string `validate:"required,min=1,max=100"json:"email"`
}
