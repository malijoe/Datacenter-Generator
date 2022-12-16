package deviceTemplateAggregate

import (
	"context"
	"errors"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

func GetDeviceTemplateAggregateId(eventAggregateId string) string {
	return strings.ReplaceAll(eventAggregateId, string(DeviceTemplateAggregateType)+"-", "")
}

func LoadDeviceTemplateAggregate(ctx context.Context, store events.AggregateStore, aggregateId string) (*DeviceTemplateAggregate, error) {
	deviceTemplate := NewDeviceTemplateAggregateWithId(aggregateId)

	err := store.Exists(ctx, deviceTemplate.GetId())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err = store.Load(ctx, deviceTemplate); err != nil {
		return nil, err
	}

	return deviceTemplate, nil
}
