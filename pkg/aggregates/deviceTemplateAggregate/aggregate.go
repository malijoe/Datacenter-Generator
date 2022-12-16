package deviceTemplateAggregate

import (
	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	"github.com/malijoe/DatacenterGenerator/pkg/components/hardware"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
)

const DeviceTemplateAggregateType events.AggregateType = "deviceTemplate"

type DeviceTemplateAggregate struct {
	*events.AggregateBase
	DeviceTemplate *datacenter.DeviceTemplate
}

func NewDeviceTemplateAggregateWithId(id string) *DeviceTemplateAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewDeviceTemplateAggregate()
	aggregate.SetId(id)
	return aggregate
}

func NewDeviceTemplateAggregate() *DeviceTemplateAggregate {
	aggregate := &DeviceTemplateAggregate{
		DeviceTemplate: datacenter.NewDeviceTemplate(),
	}
	base := events.NewAggregateBase(aggregate.When)
	base.SetType(DeviceTemplateAggregateType)
	aggregate.AggregateBase = base
	return aggregate
}

func (a *DeviceTemplateAggregate) When(event events.Event) error {
	switch event.GetEventType() {
	case eventsv1.DeviceTemplateCreated:
		return a.onCreate(event)
	default:
		return events.ErrInvalidEventType
	}
}

func (a *DeviceTemplateAggregate) onCreate(event events.Event) error {
	var data eventsv1.DeviceTemplateCreatedEvent
	if err := event.GetJsonData(&data); err != nil {
		return err
	}

	a.DeviceTemplate.ID = GetDeviceTemplateAggregateId(a.GetId())
	a.DeviceTemplate.Variant = data.Variant
	a.DeviceTemplate.Categories = data.Categories
	a.DeviceTemplate.HostnameTemplate = data.HostnameTemplate
	a.DeviceTemplate.Alias = data.Alias
	a.DeviceTemplate.Function = data.Function

	a.DeviceTemplate.Model = hardware.HardwareModel{ID: data.ModelId}

	return nil
}
