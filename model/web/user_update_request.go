package web

type UserUpdateRequest struct {
	Id       int    `validate:"required"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Email    string `validate:"required,min=10,max=100" json:"email"`
	Password string `validate:"required,min=5,max=100" json:"password"`
}
