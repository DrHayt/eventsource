package orders

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vancelongwill/eventsource"
	"github.com/vancelongwill/eventsource/_examples/fullproto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//Order is an example of state generated from left fold of events
type Order struct {
	ID        string
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
	State     string
}

const (
	EmptyState     = ""
	CreatedState   = "created"
	ShippedState   = "shipped"
	CancelledState = "cancelled"
)

var (
	ErrTransitionNotAllowed = errors.New("transition not allowed")
	allowedTransitions      = map[string]map[string]struct{}{
		EmptyState: {
			CreatedState: struct{}{},
		},
		CreatedState: {
			ShippedState:   struct{}{},
			CancelledState: struct{}{},
		},
		ShippedState:   {},
		CancelledState: {},
	}
)

func TransitionAllowed(from, to string) error {
	if stateRecord, ok := allowedTransitions[from]; ok {
		// Ok, we found the from state in the map, is the value of the to correct?
		if _, ok := stateRecord[to]; ok {
			return nil
		}
	}
	return fmt.Errorf("transition from %s to %s: %w", from, to, ErrTransitionNotAllowed)
}

//OrderCreated event used a marker of order created
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

//CreateOrder command
type CreateOrder struct {
	eventsource.CommandModel
}

func (s CreateOrder) StateName() string {
	return CreatedState
}

//ShipOrder command
type ShipOrder struct {
	eventsource.CommandModel
}

func (s ShipOrder) StateName() string {
	return ShippedState
}

//CancelOrder command
type CancelOrder struct {
	eventsource.CommandModel
}

func (s CancelOrder) StateName() string {
	return CancelledState
}

//On implements Aggregate interface
func (item *Order) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *OrderCreated:
		item.State = CreatedState

	case *OrderShipped:
		item.State = ShippedState

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
	switch v := command.(type) {
	case *CreateOrder:
		if err := TransitionAllowed(item.State, v.StateName()); err != nil {
			return nil, err
		}
		orderCreated := &OrderCreated{
			OrderCreated: &pb.OrderCreated{
				OrderId: command.AggregateID(), Version: int32(item.Version + 1), At: timestamppb.New(time.Now()),
			},
		}
		return []eventsource.Event{orderCreated}, nil
	case *ShipOrder:
		if err := TransitionAllowed(item.State, v.StateName()); err != nil {
			return nil, err
		}
		orderShipped := &OrderShipped{
			OrderShipped: &pb.OrderShipped{
				OrderId: command.AggregateID(), Version: int32(item.Version + 1), At: timestamppb.New(time.Now()),
			},
		}
		return []eventsource.Event{orderShipped}, nil

	default:
		return nil, fmt.Errorf("unhandled command, %v", v)
	}
}
