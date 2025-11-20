package dto

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32" example:"john_doe"`
	Email    string `json:"email" binding:"omitempty,email" example:"john@example.com"`
	Phone    string `json:"phone" binding:"omitempty,len=11" example:"13800138000"`
	Password string `json:"password" binding:"required,min=6,max=32" example:"password123"`
}

// UpdateUserRequest represents the request to update user information
type UpdateUserRequest struct {
	Email  string `json:"email" binding:"omitempty,email" example:"john@example.com"`
	Phone  string `json:"phone" binding:"omitempty,len=11" example:"13800138000"`
	Avatar string `json:"avatar" binding:"omitempty,url" example:"https://example.com/avatar.jpg"`
	Status int8   `json:"status" binding:"omitempty,oneof=0 1" example:"1"`
}

