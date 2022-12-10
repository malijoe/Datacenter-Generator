package rackAggregate

import (
	"context"
	"errors"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

func GetRackAggregateId(eventAggregateId string) string {
	return strings.ReplaceAll(eventAggregateId, string(RackAggregateType)+"-", "")
}

func LoadRackAggregate(ctx context.Context, store events.AggregateStore, aggregateId string) (*RackAggregate, error) {
	rack := NewRackAggregateWithId(aggregateId)

	err := store.Exists(ctx, rack.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err = store.Load(ctx, rack); err != nil {
		return nil, err
	}

	return rack, nil
}
