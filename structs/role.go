package structs

type AddRoleRequest struct {
	Name string `json:"name" binding:"required"`
}
