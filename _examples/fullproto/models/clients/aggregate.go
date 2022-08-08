package clients

//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative ./events.proto
//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative ./commands.proto
//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative ./aggregate.proto
//go:generate protoc --eventsource_out=. events.proto
//go:generate protoc --commands_out=. commands.proto

import (
	"context"
	"fmt"
	"time"

	"github.com/vancelongwill/eventsource"
)

// On implements Aggregate interface
// This is where you describe how the event changes your aggregate.
// If your event has no change to the aggregate, then no changes will be made
func (item *Client) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *Created:
		item.CreatedAt = event.EventAt().Unix()

	case *Activated:
		item.Active = true

	case *DeActivated:
		item.Active = false

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	// Make this automatic by adding it to the interface.
	item.SetBaseFields(event)

	return nil
}

func (item *Client) SetBaseFields(event eventsource.Event) {
	item.Version = int32(event.EventVersion())
	item.Id = event.AggregateID()
	item.UpdatedAt = event.EventAt().Unix()
}

//Apply implements the CommandHandler interface
func (item *Client) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	// Does this application change state, and is the state change allowed?
	//if sc, ok := command.(StateChanger); ok {
	//	if err := TransitionAllowed(item.State, sc.StateName()); err != nil {
	//		return nil, err
	//	}
	//}
	nextVersion := item.Version + 1
	switch v := command.(type) {

	case *Create:
		e := &Created{
			Id:      v.AggregateID(),
			Version: nextVersion,
			At:      time.Now().Unix(),
			By:      v.By,
		}
		return []eventsource.Event{e}, nil

	case *Activate:
		e := &Activated{
			Id:      v.AggregateID(),
			Version: nextVersion,
			At:      time.Now().Unix(),
			By:      v.By,
		}
		return []eventsource.Event{e}, nil

	case *DeActivate:
		e := &DeActivated{
			Id:      v.AggregateID(),
			Version: nextVersion,
			At:      time.Now().Unix(),
			By:      v.By,
		}
		return []eventsource.Event{e}, nil

	case *ChangeName:
		e := &NameChanged{
			Id:      v.AggregateID(),
			Version: nextVersion,
			At:      time.Now().Unix(),
			Name:    v.Name,
			By:      v.By,
		}
		return []eventsource.Event{e}, nil

	default:
		return nil, fmt.Errorf("unhandled command, %T", v)
	}
}
