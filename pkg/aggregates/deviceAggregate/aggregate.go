package deviceAggregate

import (
	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	"github.com/malijoe/DatacenterGenerator/pkg/components/hardware"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

const DeviceAggregateType events.AggregateType = "device"

type DeviceAggregate struct {
	*events.AggregateBase
	Device *datacenter.Device
}

func NewDeviceAggregateWithId(id string) *DeviceAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewDeviceAggregate()
	aggregate.SetId(id)
	return aggregate
}

func NewDeviceAggregate() *DeviceAggregate {
	aggregate := &DeviceAggregate{
		Device: datacenter.NewDevice(),
	}
	base := events.NewAggregateBase(aggregate.When)
	base.SetType(DeviceAggregateType)
	aggregate.AggregateBase = base
	return aggregate
}

func (a *DeviceAggregate) When(event events.Event) error {
	switch event.GetEventType() {
	case eventsv1.DeviceCreated:
		return a.onCreate(event)
	default:
		return events.ErrInvalidEventType
	}
}

func (a *DeviceAggregate) onCreate(event events.Event) error {
	var data eventsv1.DeviceCreatedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	a.Device.ID = GetDeviceAggregateId(event.GetAggregateId())
	a.Device.Hostname = data.Hostname
	a.Device.Elevation = data.Elevation
	a.Device.Designation = data.Designation
	a.Device.Cluster = data.Cluster
	a.Device.Instance = data.Instance
	a.Device.Categories = data.Categories

	if data.PodId != "" {
		a.Device.Pod = datacenter.NewPod()
		a.Device.Pod.ID = data.PodId
	}

	a.Device.Rack = datacenter.NewRack()
	a.Device.Rack.ID = data.RackId

	a.Device.Model = hardware.HardwareModel{ID: data.ModelId}

	return nil
}
