package deviceAggregate

import (
	"context"
	"errors"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

func GetDeviceAggregateId(eventAggregateId string) string {
	return strings.ReplaceAll(eventAggregateId, string(DeviceAggregateType)+"-", "")
}

func LoadDeviceAggregate(ctx context.Context, store events.AggregateStore, aggregateId string) (*DeviceAggregate, error) {
	device := NewDeviceAggregateWithId(aggregateId)

	err := store.Exists(ctx, device.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err = store.Load(ctx, device); err != nil {
		return nil, err
	}

	return device, nil
}
