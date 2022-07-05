package web

type GuestResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_Number string `json:"phone_number"`
	Email        string `json:"email"`
}
