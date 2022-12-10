package rackAggregate

import (
	"context"
	"strings"

	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
)

const defaultRackSize = 45

func (a *RackAggregate) CreateRack(ctx context.Context, name string, size int, datacenterId string) error {
	if name == "" {
		return ErrRackNameNotSpecified
	}
	name = strings.ToLower(name)

	if size == 0 {
		size = defaultRackSize
	}

	if datacenterId == "" {
		return ErrDatacenterIDNotProvided
	}

	event, err := eventsv1.NewRackCreatedEvent(a, name, size, datacenterId)
	if err != nil {
		return err
	}

	return a.Apply(event)
}

func (a *RackAggregate) AddDevice(ctx context.Context, deviceId string, elevation int, formFactor int) error {
	if deviceId == "" {
		return ErrDeviceIDNotProvided
	}

	if formFactor == 0 {
		return ErrDeviceFormFactorNotProvided
	}

	event, err := eventsv1.NewDeviceRackedEvent(a, deviceId, elevation, formFactor)
	if err != nil {
		return err
	}

	return a.Apply(event)
}
