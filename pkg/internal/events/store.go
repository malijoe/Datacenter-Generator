package events

import "context"

// AggregateStore is responsible for loading and saving aggregates.
type AggregateStore interface {
	// Load loads the most recent version of an aggregate to provided into params aggregate with a type and id.
	Load(ctx context.Context, aggregate Aggregate) error

	// Save saves the uncommitted events for an aggregate.
	Save(ctx context.Context, aggregate Aggregate) error

	// Exists check aggregate exists by id.
	Exists(ctx context.Context, streamId string) error
}

// EventStore is an interface for an event sourcing event store.
type EventStore interface {
	// SaveEvents appends all events in the event stream to the store.
	SaveEvents(ctx context.Context, events []Event) error

	// LoadEvents loads all events for the aggregate id from the store.
	LoadEvents(ctx context.Context, streamId string) ([]Event, error)
}
