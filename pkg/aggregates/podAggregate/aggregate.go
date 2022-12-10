package podAggregate

import (
	"fmt"

	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

const PodAggregateType events.AggregateType = "pod"

type PodAggregate struct {
	*events.AggregateBase
	Pod *datacenter.Pod
}

func NewPodAggregateWithId(id string) *PodAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewPodAggregate()
	aggregate.SetId(id)
	return aggregate
}

func NewPodAggregate() *PodAggregate {
	aggregate := &PodAggregate{
		Pod: datacenter.NewPod(),
	}
	base := events.NewAggregateBase(aggregate.When)
	base.SetType(PodAggregateType)
	aggregate.AggregateBase = base
	return aggregate
}

func (a *PodAggregate) When(event events.Event) error {
	switch event.GetEventType() {
	default:
		return events.ErrInvalidEventType
	}
}

func (a *PodAggregate) onCreate(event events.Event) error {
	var data eventsv1.PodCreatedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	a.Pod.ID = GetPodAggregateId(a.GetId())
	a.Pod.Function = data.Function
	a.Pod.Instance = data.Instance
	a.Pod.Datacenter = datacenter.NewDatacenter()
	a.Pod.Datacenter.ID = data.DatacenterId
	a.Pod.Name = fmt.Sprintf("%s%d", data.Function, data.Instance)

	return nil
}
