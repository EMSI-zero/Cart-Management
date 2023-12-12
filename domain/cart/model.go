package cart

import "time"

type CartState int64

const (
	Pending CartState = iota
	Completed
)

type Cart struct {
	ID        int64
	Data      []byte
	State     CartState
	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
}

//DTO

type CartToDelete struct {
	ID        int64
	UpdatedAt time.Time
}
