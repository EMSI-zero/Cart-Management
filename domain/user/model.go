package user

import "time"

type UserModel struct {
	ID           int64
	UserName     string
	UserPassword string
	Salt         string
	CreatedAt    time.Time
}

// Data Transfer Objects (DTOs)
type LoginRequest struct {
	UserName string
	Password string
}

type LoginResponse struct {
	Success bool
}

type RegisterRequest struct {
}

type RegisterResponse struct {
}
