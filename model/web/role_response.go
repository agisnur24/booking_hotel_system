package web

type RoleResponse struct {
	Id              int    `json:"id"`
	RoleName        string `json:"name"`
	RoleDescription string `json:"description"`
}
