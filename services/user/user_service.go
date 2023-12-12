package user

import (
	"cart-manager/domain/user"
	"cart-manager/registry"
	"fmt"
)

func InitUserService(r registry.ServiceRegistry) error {
	if r == nil {
		return fmt.Errorf("no registry found")
	}

	r.RegisterUserService(&UserService{})

	return nil
}

type UserService struct {
	user.UnimplementedUserService
}
