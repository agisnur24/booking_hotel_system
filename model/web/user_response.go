package web

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}
