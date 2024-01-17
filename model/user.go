package model

type SignUpUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required,min=5,max=15"`
}

type LoginRequest struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required,min=5,max=15"`
}

type User struct {
	Email    string `db:"email"`
	Username string `db:"username"`
}