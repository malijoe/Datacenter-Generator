package podAggregate

import (
	"context"
	"errors"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

func GetPodAggregateId(eventAggregateId string) string {
	return strings.ReplaceAll(eventAggregateId, string(PodAggregateType)+"-", "")
}

func LoadPodAggregate(ctx context.Context, store events.AggregateStore, aggregateId string) (*PodAggregate, error) {
	pod := NewPodAggregateWithId(aggregateId)

	err := store.Exists(ctx, pod.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err = store.Load(ctx, pod); err != nil {
		return nil, err
	}

	return pod, nil
}
