package v1

import (
	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/events"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/units"
)

const (
	DatacenterCreated     = "V1_DATACENTER_CREATED"
	PodCreated            = "V1_POD_CREATED"
	DatacenterPodAdded    = "V1_DATACENTER_POD_ADDED"
	RackCreated           = "V1_RACK_CREATED"
	DatacenterRackAdded   = "V1_DATACENTER_RACK_ADDED"
	DeviceCreated         = "V1_DEVICE_CREATED"
	DeviceRacked          = "V1_DEVICE_RACKED"
	DeviceTemplateCreated = "V1_DEVICE_TEMPLATE_CREATED"
)

type DatacenterCreatedEvent struct {
	Site      string                 `json:"site"`
	Building  string                 `json:"building"`
	Room      string                 `json:"room"`
	Providers map[string]units.Value `json:"providers"`
}

func NewDatacenterCreatedEvent(aggregate events.Aggregate, site, building, room string, providers map[string]units.Value) (events.Event, error) {
	data := DatacenterCreatedEvent{
		Site:      site,
		Building:  building,
		Room:      room,
		Providers: providers,
	}
	event := events.NewBaseEvent(aggregate, DatacenterCreated)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil
}

type PodCreatedEvent struct {
	Function     datacenter.Function `json:"function"`
	Instance     int
	DatacenterId string `json:"datacenterId"`
}

func NewPodCreatedEvent(aggregate events.Aggregate, function datacenter.Function, instance int, datacenterId string) (events.Event, error) {
	data := PodCreatedEvent{
		Function:     function,
		Instance:     instance,
		DatacenterId: datacenterId,
	}
	event := events.NewBaseEvent(aggregate, PodCreated)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil
}

type DatacenterPodAddedEvent struct {
	PodId string `json:"podId"`
}

func NewDatacenterPodAddedEvent(aggregate events.Aggregate, podId string) (events.Event, error) {
	data := DatacenterPodAddedEvent{
		PodId: podId,
	}
	event := events.NewBaseEvent(aggregate, DatacenterPodAdded)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil
}

type RackCreatedEvent struct {
	Name         string `json:"name"`
	Size         int    `json:"size"`
	DatacenterId string `json:"datacenterId"`
}

func NewRackCreatedEvent(aggregate events.Aggregate, name string, size int, datacenterId string) (events.Event, error) {
	data := RackCreatedEvent{
		Name:         name,
		Size:         size,
		DatacenterId: datacenterId,
	}
	event := events.NewBaseEvent(aggregate, RackCreated)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil
}

type DatacenterRackAddedEvent struct {
	RackId string `json:"rackId"`
}

func NewDatacenterRackAddedEvent(aggregate events.Aggregate, rackId string) (events.Event, error) {
	data := DatacenterRackAddedEvent{
		RackId: rackId,
	}
	event := events.NewBaseEvent(aggregate, DatacenterRackAdded)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil
}

type DeviceCreatedEvent struct {
	Hostname    string                 `json:"hostname"`
	Elevation   int                    `json:"elevation"`
	Designation datacenter.Designation `json:"designation"`
	Cluster     int                    `json:"cluster"`
	Instance    int                    `json:"instance"`
	ModelId     string                 `json:"modelId"`
	Categories  []string               `json:"categories"`
	PodId       string                 `json:"podId"`
	RackId      string                 `json:"rackId"`
}

func NewDeviceCreatedEvent(aggregate events.Aggregate, hostname string, elevation int, designation datacenter.Designation, cluster, instance int, modelId string, categories []string, podId, rackId string) (events.Event, error) {
	data := DeviceCreatedEvent{
		Hostname:    hostname,
		Elevation:   elevation,
		Designation: designation,
		Cluster:     cluster,
		Instance:    instance,
		ModelId:     modelId,
		Categories:  categories,
		PodId:       podId,
		RackId:      rackId,
	}
	event := events.NewBaseEvent(aggregate, DeviceCreated)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil
}

type DeviceRackedEvent struct {
	DeviceId   string `json:"deviceId"`
	Elevation  int    `json:"elevation"`
	FormFactor int    `json:"formFactor"`
}

func NewDeviceRackedEvent(aggregate events.Aggregate, deviceId string, elevation int, formFactor int) (events.Event, error) {
	data := DeviceRackedEvent{
		DeviceId:   deviceId,
		Elevation:  elevation,
		FormFactor: formFactor,
	}
	event := events.NewBaseEvent(aggregate, DeviceRacked)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil

}

type DeviceTemplateCreatedEvent struct {
	ModelId          string              `json:"modelId"`
	Variant          string              `json:"variant"`
	Categories       []string            `json:"categories"`
	HostnameTemplate string              `json:"hostnameTemplate"`
	Alias            string              `json:"alias"`
	Function         datacenter.Function `json:"function"`
}

func NewDeviceTemplateCreatedEvent(aggregate events.Aggregate, modelId, variant string, categories []string, hostnameTemplate, alias string, function datacenter.Function) (events.Event, error) {
	data := DeviceTemplateCreatedEvent{
		ModelId:          modelId,
		Variant:          variant,
		Categories:       categories,
		HostnameTemplate: hostnameTemplate,
		Alias:            alias,
		Function:         function,
	}
	event := events.NewBaseEvent(aggregate, DeviceTemplateCreated)
	if err := event.SetJsonData(&data); err != nil {
		return events.Event{}, err
	}
	return event, nil
}
