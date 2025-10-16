type UpdateUserRequest struct {
	Email     string `json:"email" binding:"omitempty,email"`
	FirstName string `json:"first_name" binding:"omitempty"`
	LastName  string `json:"last_name" binding:"omitempty"`
	status    string `json:"status" binding:"omitempty,oneof=active inactive"`
}