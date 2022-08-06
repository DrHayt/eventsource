package orders

import (
	"errors"
	"fmt"

	"github.com/vancelongwill/eventsource"
)

// StateChanger is the interface used to determine if a command causes a change in state.
type StateChanger interface {
	StateName() string
}

// These are the states we are concerned with.
const (
	StateEmpty     = ""
	StateCreated   = "created"
	StateShipped   = "shipped"
	StateCancelled = "cancelled"
)

var (
	ErrTransitionNotAllowed = errors.New("transition not allowed")
	allowedTransitions      = map[string]map[string]struct{}{
		StateEmpty: {
			StateCreated: struct{}{},
		},
		StateCreated: {
			StateShipped:   struct{}{},
			StateCancelled: struct{}{},
		},
		StateShipped:   {},
		StateCancelled: {},
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

//CreateOrder command
type CreateOrder struct {
	eventsource.CommandModel
}

func (s CreateOrder) StateName() string {
	return StateCreated
}

//ShipOrder command
type ShipOrder struct {
	eventsource.CommandModel
}

func (s ShipOrder) StateName() string {
	return StateShipped
}

//CancelOrder command
type CancelOrder struct {
	eventsource.CommandModel
	CancelReason string
}

func (s CancelOrder) StateName() string {
	return StateCancelled
}

// ChangeNameOrder command
type ChangeNameOrder struct {
	eventsource.CommandModel
	FirstName string
	LastName  string
}
