package pb

import "time"

func (m *OrderCreated) AggregateID() string { return m.OrderId }
func (m *OrderCreated) EventVersion() int   { return int(m.Version) }
func (m *OrderCreated) EventAt() time.Time  { return m.At.AsTime() }

func (m *OrderShipped) AggregateID() string { return m.OrderId }
func (m *OrderShipped) EventVersion() int   { return int(m.Version) }
func (m *OrderShipped) EventAt() time.Time  { return m.At.AsTime() }

func (m *OrderCancelled) AggregateID() string { return m.OrderId }
func (m *OrderCancelled) EventVersion() int   { return int(m.Version) }
func (m *OrderCancelled) EventAt() time.Time  { return m.At.AsTime() }
