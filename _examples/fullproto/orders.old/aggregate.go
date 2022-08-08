package orders

import (
	"context"
	"fmt"
	"time"

	"github.com/vancelongwill/eventsource"
	"github.com/vancelongwill/eventsource/_examples/fullproto/pb"
)

type Order struct {
	ID              string
	Version         int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CancelledAt     time.Time
	CancelledReason string
	State           string
	FirstName       string
	Lastname        string
}

// On implements Aggregate interface
// This is where you describe how the event changes your aggregate.
// If your event has no change to the aggregate, then no changes will be made
func (item *Order) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *OrderCreated:
		item.CreatedAt = event.EventAt()
		item.State = StateCreated

	case *OrderShipped:
		item.State = StateShipped

	case *OrderCancelled:
		item.State = StateCancelled
		item.CancelledAt = event.EventAt()
		item.CancelledReason = v.CancelReason

	case *OrderNameChanged:
		item.FirstName = v.FirstName
		item.Lastname = v.LastName

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	item.Version = event.EventVersion()
	item.ID = event.AggregateID()
	item.UpdatedAt = event.EventAt()

	return nil
}

//Apply implements the CommandHandler interface
func (item *Order) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	// Does this application change state, and is the state change allowed?
	if sc, ok := command.(StateChanger); ok {
		if err := TransitionAllowed(item.State, sc.StateName()); err != nil {
			return nil, err
		}
	}
	nextVersion := int64(item.Version + 1)
	switch v := command.(type) {

	case *CreateOrder:
		e := pb.NewOrderCreated(v.AggregateID(), nextVersion)
		return []eventsource.Event{e}, nil

	case *ShipOrder:
		e := pb.NewOrderShipped(item.ID, nextVersion)
		return []eventsource.Event{e}, nil

	case *CancelOrder:
		e := pb.NewOrderCancelled(item.ID, nextVersion, v.CancelReason)
		return []eventsource.Event{e}, nil

	case *ChangeNameOrder:
		orderShipped := pb.NewOrderNameChanged(item.ID, nextVersion, v.FirstName, v.LastName)
		return []eventsource.Event{orderShipped}, nil

	default:
		return nil, fmt.Errorf("unhandled command, %T", v)
	}
}
