package datacenterAggregate

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/units"
)

func (a *DatacenterAggregate) CreateDatacenter(ctx context.Context, site string, building string, room string, providers map[string]string) error {
	if site == "" {
		return ErrSiteNotSpecified
	}
	site = strings.ToLower(site)
	building = strings.ToLower(building)
	room = strings.ToLower(room)

	parsedProviders := make(map[string]units.Value)
	if len(providers) > 0 {
		var parseProviderErr error
		for provider, speed := range providers {
			value, err := units.ParseValue(speed)
			if err != nil {
				parseProviderErr = multierror.Append(parseProviderErr, fmt.Errorf("%w provider: {%s}, speed: {%s}", ErrInvalidProviderTransferSpeed, provider, speed))
				continue
			}
			parsedProviders[strings.ToLower(provider)] = value
		}

		if parseProviderErr != nil {
			return parseProviderErr
		}
	}

	event, err := eventsv1.NewDatacenterCreatedEvent(a, site, building, room, parsedProviders)
	if err != nil {
		return err
	}
	return a.Apply(event)
}

func (a *DatacenterAggregate) AddPod(ctx context.Context, podId string) error {
	if podId == "" {
		return ErrPodIDNotProvided
	}

	event, err := eventsv1.NewDatacenterPodAddedEvent(a, podId)
	if err != nil {
		return err
	}

	return a.Apply(event)
}

func (a *DatacenterAggregate) AddRack(ctx context.Context, rackId string) error {
	if rackId == "" {
		return ErrRackIDNotProvided
	}

	event, err := eventsv1.NewDatacenterRackAddedEvent(a, rackId)
	if err != nil {
		return err
	}

	return a.Apply(event)
}
