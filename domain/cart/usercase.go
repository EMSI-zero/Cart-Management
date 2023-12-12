package cart

import "context"

type CartService interface {
	Carts(context.Context) ([]*Cart, error)
	Cart(context.Context) (*Cart, error)
	UpdateCart(context.Context, *Cart) (*Cart, error)
	CreateCart(context.Context, *Cart) (*Cart, error)
	DeleteCart(context.Context, *CartToDelete) error
}
