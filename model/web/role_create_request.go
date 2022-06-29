package web

type RoleCreateRequest struct {
	RoleName        string `validate:"required,min=1,max=100" json:"name"`
	RoleDescription string `validate:"required,min=1,max=255" json:"description"`
}
