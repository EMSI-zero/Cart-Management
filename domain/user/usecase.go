package user

type UserService interface {
	Login(*LoginRequest) *LoginResponse
	Register(*RegisterRequest) *RegisterResponse
	GetUserByID(int64) *UserModel
}
