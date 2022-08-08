package orders

import (
	"context"
	"fmt"
	"time"

	"github.com/vancelongwill/eventsource"
)

// On implements Aggregate interface
// This is where you describe how the event changes your aggregate.
// If your event has no change to the aggregate, then no changes will be made
func (item *Aggregate) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *OrderCreated:
		item.CreatedAt = event.EventAt().Unix()
		item.State = OrderState_CREATED

	case *OrderShipped:
		item.State = OrderState_SHIPPED

	case *OrderCancelled:
		item.State = OrderState_CANCELLED
		item.CancelledAt = event.EventAt().Unix()
		item.CancelReason = v.CancelReason

	case *OrderNameChanged:
		item.FirstName = v.FirstName
		item.LastName = v.LastName

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	item.Version = int32(event.EventVersion())
	item.Id = event.AggregateID()
	item.UpdatedAt = event.EventAt().Unix()

	return nil
}

//Apply implements the CommandHandler interface
func (item *Aggregate) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	// Does this application change state, and is the state change allowed?
	//if sc, ok := command.(StateChanger); ok {
	//	if err := TransitionAllowed(item.State, sc.StateName()); err != nil {
	//		return nil, err
	//	}
	//}
	nextVersion := item.Version + 1
	switch v := command.(type) {

	case *CreateOrder:
		e := &OrderCreated{
			Id:      v.AggregateID(),
			Version: nextVersion,
			At:      time.Now().Unix(),
			By:      v.By,
		}
		return []eventsource.Event{e}, nil

	case *Ship:
		e := &OrderShipped{
			Id:          v.AggregateID(),
			Version:     nextVersion,
			At:          time.Now().Unix(),
			Destination: v.Destination,
		}
		return []eventsource.Event{e}, nil

	case *CancelOrder:
		e := &OrderCancelled{
			Id:           v.AggregateID(),
			Version:      nextVersion,
			At:           time.Now().Unix(),
			CancelReason: v.Reason,
		}
		return []eventsource.Event{e}, nil

	case *ChangeName:
		e := &OrderNameChanged{
			Id:        v.AggregateID(),
			Version:   nextVersion,
			At:        time.Now().Unix(),
			FirstName: v.FirstName,
			LastName:  v.LastName,
		}
		return []eventsource.Event{e}, nil

	default:
		return nil, fmt.Errorf("unhandled command, %T", v)
	}
}
