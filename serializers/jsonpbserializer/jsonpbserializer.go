package jsonpbserializer

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"

	"github.com/golang/protobuf/jsonpb"
	"github.com/vancelongwill/eventsource"
	"github.com/vancelongwill/eventsource/pbevent"
	"google.golang.org/protobuf/proto"
)

type JSONPBSerializer struct {
	eventTypes map[string]reflect.Type
	mu         sync.Mutex
}

// Bind registers the specified events with the serializer; may be called more than once
func (p *JSONPBSerializer) Bind(events ...eventsource.Event) {
	p.mu.Lock()
	for _, event := range events {
		if _, ok := event.(proto.Message); ok {
			eventType, reflectType := eventsource.EventType(event)
			p.eventTypes[eventType] = reflectType
		}
	}
	p.mu.Unlock()

}

// MarshalEvent converts an Event to a Record
func (p *JSONPBSerializer) MarshalEvent(v eventsource.Event) (eventsource.Record, error) {
	pv, ok := v.(proto.Message)
	if !ok {
		return eventsource.Record{}, fmt.Errorf("event must implement the ProtoEventer interface: %w", eventsource.ErrInvalidEncoding)
	}
	data, err := proto.Marshal(pv)
	if err != nil {
		return eventsource.Record{}, err
	}

	s, _ := eventsource.EventType(v)

	data, err = proto.Marshal(&pbevent.Event{
		EventTypename: s,
		EventData:     data,
	})

	if err != nil {
		return eventsource.Record{}, fmt.Errorf("unable to encode event: %w", eventsource.ErrInvalidEncoding)
	}

	return eventsource.Record{
		Version: v.EventVersion(),
		Data:    data,
	}, nil
}

// UnmarshalEvent converts an Event backed into a Record
func (p *JSONPBSerializer) UnmarshalEvent(record eventsource.Record) (eventsource.Event, error) {
	var evt pbevent.Event
	err := jsonpb.Unmarshal(bytes.NewReader(record.Data), &evt)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal event: %w", eventsource.ErrInvalidEncoding)
	}

	t, ok := p.eventTypes[evt.EventTypename]
	if !ok {
		return nil, fmt.Errorf("unbound event type, %v: %w", evt.EventTypename, eventsource.ErrUnboundEventType)
	}

	n := reflect.New(t)
	c := n.Interface().(proto.Message)
	v := c.ProtoReflect().New().Interface()
	err = jsonpb.Unmarshal(bytes.NewReader(evt.EventData), v)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal event data into %#v: %w", v, eventsource.ErrInvalidEncoding)
	}
	n.Elem().Field(0).Set(reflect.ValueOf(v))

	return n.Interface().(eventsource.Event), nil
}

// NewJSONPBSerializer constructs a new JSONPBSerializer and populates it with the specified events.
// Bind may be subsequently called to add more events.
// Events must also fulfill the proto.Message interface.
func NewJSONPBSerializer(events ...eventsource.Event) *JSONPBSerializer {
	serializer := &JSONPBSerializer{
		eventTypes: map[string]reflect.Type{},
	}
	serializer.Bind(events...)
	return serializer
}
