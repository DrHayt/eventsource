// Code generated by eventsource-protobuf. DO NOT EDIT.
// source: commands.proto

package orders

// AggregateID implements the eventsource.Command interface for CreateOrder
func (c *CreateOrder) AggregateID() string {
	return c.Id
}

// AggregateID implements the eventsource.Command interface for CancelOrder
func (c *CancelOrder) AggregateID() string {
	return c.Id
}

// AggregateID implements the eventsource.Command interface for AuditOrder
func (c *AuditOrder) AggregateID() string {
	return c.Id
}

// AggregateID implements the eventsource.Command interface for ChangeName
func (c *ChangeName) AggregateID() string {
	return c.Id
}

// AggregateID implements the eventsource.Command interface for Ship
func (c *Ship) AggregateID() string {
	return c.Id
}
