package events

import "context"

// AggregateStore is responsible for loading and saving aggregates.
type AggregateStore interface {
	// LoadAggregate loads the most recent version of an aggregate to provided into params aggregate with a type and id.
	LoadAggregate(ctx context.Context, aggregate Aggregate) error

	// SaveAggregate saves the uncommitted events for an aggregate.
	SaveAggregate(ctx context.Context, aggregate Aggregate) error

	DeleteAggregate(ctx context.Context, aggregateType string, id string) error
	GetAggregateTypeEntries(aggregateType string) (entries []*Entry, err error)
}

// EventStore is an interface for an event sourcing event store.
type EventStore interface {
	// SaveEvents appends all events in the event stream to the store.
	SaveEvents(ctx context.Context, events []Event) error

	// LoadEvents loads all events for the aggregate id from the store.
	LoadEvents(ctx context.Context, streamId string) ([]Event, error)
}

type AggregateStoreMiddleware func(store AggregateStore) AggregateStore

func AggregateStoreWithMiddleware(store AggregateStore, middleware ...AggregateStoreMiddleware) AggregateStore {
	s := store
	for _, m := range middleware {
		s = m(s)
	}
	return s
}
