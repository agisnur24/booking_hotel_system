package web

type RoleUpdateRequest struct {
	Id              int    `validate:"required"`
	RoleName        string `validate:"required,min=1,max=100" json:"role_name"`
	RoleDescription string `validate:"required,min=1,max=100" json:"role_description"`
}
