package podAggregate

import (
	"context"

	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
)

func (a *PodAggregate) CreatePod(ctx context.Context, function string, dc *datacenter.Datacenter) error {
	if function == "" {
		return ErrFunctionNotSpecified
	}

	parsedFunction := datacenter.ParseFunction(function)
	if parsedFunction == datacenter.UnknownFunction {
		return ErrInvalidFunctionSpecified
	}
	// get the number of pod instances with the same function
	n := dc.NumPodInstances(parsedFunction)
	// add 1 to reflect the addition of this pod
	n += 1

	event, err := eventsv1.NewPodCreatedEvent(a, parsedFunction, n, dc.ID)
	if err != nil {
		return err
	}

	return a.Apply(event)
}
