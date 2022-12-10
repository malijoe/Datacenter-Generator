package rackAggregate

import (
	"fmt"

	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	"github.com/malijoe/DatacenterGenerator/pkg/components/hardware"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

const RackAggregateType events.AggregateType = "rack"

type RackAggregate struct {
	*events.AggregateBase
	Rack *datacenter.Rack
}

func NewRackAggregateWithId(id string) *RackAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewRackAggregate()
	aggregate.SetId(id)
	return aggregate
}

func NewRackAggregate() *RackAggregate {
	aggregate := &RackAggregate{
		Rack: datacenter.NewRack(),
	}

	base := events.NewAggregateBase(aggregate.When)
	base.SetType(RackAggregateType)
	aggregate.AggregateBase = base
	return aggregate
}

func (a *RackAggregate) When(event events.Event) error {
	switch event.GetEventType() {
	case eventsv1.RackCreated:
		return a.onCreate(event)
	case eventsv1.DeviceRacked:
		return a.onDeviceAdd(event)
	default:
		return events.ErrInvalidEventType
	}
}

func (a *RackAggregate) onCreate(event events.Event) error {
	var data eventsv1.RackCreatedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	a.Rack.ID = GetRackAggregateId(event.GetAggregateId())
	a.Rack.Name = data.Name
	a.Rack.Size = data.Size
	a.Rack.Datacenter = datacenter.NewDatacenter()
	a.Rack.Datacenter.ID = data.DatacenterId

	return nil
}

func (a *RackAggregate) onDeviceAdd(event events.Event) error {
	var data eventsv1.DeviceRackedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	device := datacenter.NewDevice()
	device.ID = data.DeviceId
	device.Model = hardware.HardwareModel{FormFactor: data.FormFactor}

	// if elevation is not specified
	if data.Elevation == 0 {
		// rack device at the next available elevation
		if err := a.Rack.RackDevice(device); err != nil {
			return fmt.Errorf("RackDevice: error {%w}", err)
		}
	} else {
		// rack device at the specified elevation
		if err := a.Rack.RackDeviceAt(device, data.Elevation); err != nil {
			return fmt.Errorf("RackDeviceAt: error {%w}", err)
		}
	}

	return nil
}
