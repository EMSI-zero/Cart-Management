package registry

import "cart-manager/domain/user"

type ServiceRegistry interface {
	mustImplementBaseRegistry()
	RegisterUserService(user.UserService)
	GetUserService() user.UserService
}

type serviceRegistry struct {
	services ServiceMap
}

func (sr *serviceRegistry) mustImplementBaseRegistry() {}

type ServiceMap struct {
	UserServices user.UserService
	// ProfileService ProfileService
}

func NewServiceRegistry() *serviceRegistry {
	//create an empty service registry
	sr := new(serviceRegistry)
	return sr
}

func (sr *serviceRegistry) RegisterUserService(service user.UserService) {
	sr.services.UserServices = service
}

func (sr *serviceRegistry) GetUserService() user.UserService {
	return sr.services.UserServices
}
