package user

import (
	"cart-manager/domain/user"
	"cart-manager/registry"
)

func InitUserService(r registry.ServiceRegistry) {
	r.RegisterUserService(&UserService{})
}

type UserService struct {
	user.UnimplementedUserService
}
