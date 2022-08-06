package pb

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Force compliance with the command and event interfaces.
func (m *OrderCreated) AggregateID() string { return m.OrderId }
func (m *OrderCreated) EventVersion() int   { return int(m.Version) }
func (m *OrderCreated) EventAt() time.Time  { return m.At.AsTime() }

func (m *OrderShipped) AggregateID() string { return m.OrderId }
func (m *OrderShipped) EventVersion() int   { return int(m.Version) }
func (m *OrderShipped) EventAt() time.Time  { return m.At.AsTime() }

func (m *OrderCancelled) AggregateID() string { return m.OrderId }
func (m *OrderCancelled) EventVersion() int   { return int(m.Version) }
func (m *OrderCancelled) EventAt() time.Time  { return m.At.AsTime() }

func (m *OrderNameChanged) AggregateID() string { return m.OrderId }
func (m *OrderNameChanged) EventVersion() int   { return int(m.Version) }
func (m *OrderNameChanged) EventAt() time.Time  { return m.At.AsTime() }

func NewOrderCreated(id string, version int64) *OrderCreated {
	return &OrderCreated{
		OrderId: id,
		At:      timestamppb.New(time.Now()),
		Version: version,
	}
}

func NewOrderShipped(id string, version int64) *OrderShipped {
	return &OrderShipped{
		OrderId: id,
		At:      timestamppb.New(time.Now()),
		Version: version,
	}
}

func NewOrderCancelled(id string, version int64, cancelReason string) *OrderCancelled {
	return &OrderCancelled{
		OrderId:      id,
		At:           timestamppb.New(time.Now()),
		Version:      version,
		CancelReason: cancelReason,
	}
}

func NewOrderNameChanged(id string, version int64, firstName, lastName string) *OrderNameChanged {
	return &OrderNameChanged{
		OrderId:   id,
		At:        timestamppb.New(time.Now()),
		Version:   version,
		FirstName: firstName,
		LastName:  lastName,
	}
}
