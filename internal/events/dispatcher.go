package events

import (
	"context"
	"sync"
)

type (
	EventHandler interface {
		HandleEvent(ctx context.Context, event Event) error
	}
	EventHandlerFunc func(ctx context.Context, event Event) error
	EventSubscriber  interface {
		Subscribe(typ EventType, handler EventHandler)
	}
	EventPublisher interface {
		Publish(ctx context.Context, events ...Event) error
	}
	EventDispatcher struct {
		handlers map[EventType][]EventHandler
		mu       sync.Mutex
	}
)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[EventType][]EventHandler),
	}
}

func (d *EventDispatcher) Subscribe(typ EventType, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.handlers[typ] = append(d.handlers[typ], handler)
}

func (d *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		for _, handler := range d.handlers[event.GetEventType()] {
			err := handler.HandleEvent(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f EventHandlerFunc) HandleEvent(ctx context.Context, event Event) error {
	return f(ctx, event)
}
