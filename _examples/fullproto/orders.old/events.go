package orders

import "github.com/vancelongwill/eventsource/_examples/fullproto/pb"

// OrderCreated event used a marker of order created
type OrderCreated struct {
	*pb.OrderCreated
}

// OrderShipped event used a marker of order shipped
type OrderShipped struct {
	*pb.OrderShipped
}

//OrderCancelled event used a marker of order created
type OrderCancelled struct {
	*pb.OrderCancelled
}

//OrderNameChanged event used a marker of order created
type OrderNameChanged struct {
	*pb.OrderNameChanged
}
