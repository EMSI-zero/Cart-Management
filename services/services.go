package services

import (
	"cart-manager/registry"
	"cart-manager/services/user"
	"fmt"
)

func InitServices(r registry.ServiceRegistry) error {
	if r == nil {
		return fmt.Errorf("no registry found")
	}

	user.InitUserService(r)

	return nil
}
