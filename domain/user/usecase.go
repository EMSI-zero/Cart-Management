package user

import "fmt"

type UserService interface {
	mustImplementBaseService()
	Login(*LoginRequest) (*LoginResponse, error)
	Register(*RegisterRequest) (*RegisterResponse, error)
	GetUserByID(int64) (*UserModel, error)
}

type UnimplementedUserService struct{}

func (us *UnimplementedUserService) mustImplementBaseService() {}

func (us *UnimplementedUserService) Login(*LoginRequest) (*LoginResponse, error) {
	return nil, fmt.Errorf("service not implemented")
}

func (us *UnimplementedUserService) Register(*RegisterRequest) (*RegisterResponse, error) {
	return nil, fmt.Errorf("service not implemented")

}
func (us *UnimplementedUserService) GetUserByID(int64) (*UserModel, error) {
	return nil, fmt.Errorf("service not implemented")
}
