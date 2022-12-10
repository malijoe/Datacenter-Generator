package datacenterAggregate

import (
	"context"
	"errors"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

func GetDatacenterAggregateId(eventAggregateId string) string {
	return strings.ReplaceAll(eventAggregateId, string(DatacenterAggregateType)+"-", "")
}

func LoadDatacenterAggregate(ctx context.Context, store events.AggregateStore, aggregateId string) (*DatacenterAggregate, error) {
	datacenter := NewDatacenterAggregateWithId(aggregateId)

	err := store.Exists(ctx, datacenter.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err = store.Load(ctx, datacenter); err != nil {
		return nil, err
	}

	return datacenter, nil
}
