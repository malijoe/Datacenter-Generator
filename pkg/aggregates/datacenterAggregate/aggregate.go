package datacenterAggregate

import (
	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

const DatacenterAggregateType events.AggregateType = "datacenter"

type DatacenterAggregate struct {
	*events.AggregateBase
	Datacenter *datacenter.Datacenter
}

func NewDatacenterAggregateWithId(id string) *DatacenterAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewDatacenterAggregate()
	aggregate.SetId(id)
	return aggregate
}

func NewDatacenterAggregate() *DatacenterAggregate {
	aggregate := &DatacenterAggregate{
		Datacenter: datacenter.NewDatacenter(),
	}
	base := events.NewAggregateBase(aggregate.When)
	base.SetType(DatacenterAggregateType)
	aggregate.AggregateBase = base
	return aggregate
}

func (a *DatacenterAggregate) When(event events.Event) error {
	switch event.GetEventType() {
	case eventsv1.DatacenterCreated:
		return a.onCreate(event)
	case eventsv1.DatacenterPodAdded:
		return a.onPodAdd(event)
	case eventsv1.DatacenterRackAdded:
		return a.onRackAdd(event)
	default:
		return events.ErrInvalidEventType
	}
}

func (a *DatacenterAggregate) onCreate(event events.Event) error {
	var data eventsv1.DatacenterCreatedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	a.Datacenter.ID = GetDatacenterAggregateId(event.GetAggregateId())
	a.Datacenter.Site = data.Site
	a.Datacenter.Building = data.Building
	a.Datacenter.Room = data.Room
	a.Datacenter.Providers = data.Providers

	return nil
}

func (a *DatacenterAggregate) onPodAdd(event events.Event) error {
	var data eventsv1.DatacenterPodAddedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	pod := datacenter.NewPod()
	pod.ID = data.PodId
	pod.Datacenter = a.Datacenter

	a.Datacenter.Pods = append(a.Datacenter.Pods, pod)

	return nil
}

func (a *DatacenterAggregate) onRackAdd(event events.Event) error {
	var data eventsv1.DatacenterRackAddedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	rack := datacenter.NewRack()
	rack.ID = data.RackId
	rack.Datacenter = a.Datacenter

	a.Datacenter.Racks = append(a.Datacenter.Racks, rack)

	return nil
}
