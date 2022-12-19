package schema

// GetRoleResp get role  response
type GetRoleResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
