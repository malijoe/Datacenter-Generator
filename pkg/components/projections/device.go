package projections

import "github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"

type DeviceProjection struct {
	BaseProjection `bson:",inline"`

	ID          string   `json:"id,omitempty" bson:"id,omitempty"`
	Hostname    string   `json:"hostname,omitempty" bson:"hostname,omitempty"`
	Elevation   int      `json:"elevation,omitempty" bson:"elevation,omitmpty"`
	Designation string   `json:"designation,omitempty" bson:"designation,omitempty"`
	Cluster     int      `json:"cluster,omitmepty" bson:"cluster,omitempty"`
	Instance    int      `json:"instance,omitempty" bson:"instance,omitempty"`
	Categories  []string `json:"categories,omitempty" bson:"categories,omitempty"`

	PodId        string `json:"podId,omitempty" bson:"podId,omitempty"`
	RackId       string `json:"rackId,omitempty" bson:"rackId,omitempty"`
	DatacenterId string `json:"datacenterId,omitempty" bson:"datacenterId,omitempty"`
}

func projectionFromDevice(d *datacenter.Device, base BaseProjection) *DeviceProjection {
	var podId string
	if d.Pod != nil {
		podId = d.Pod.ID
	}

	var rackId string
	if d.Rack != nil {
		rackId = d.Rack.ID
	}

	var datacenterId string
	if d.Datacenter != nil {
		datacenterId = d.Datacenter.ID
	}

	return &DeviceProjection{
		ID:           d.ID,
		Hostname:     d.Hostname,
		Elevation:    d.Elevation,
		Designation:  string(d.Designation),
		Cluster:      d.Cluster,
		Instance:     d.Instance,
		Categories:   d.Categories,
		PodId:        podId,
		RackId:       rackId,
		DatacenterId: datacenterId,
	}
}

func NewCreatedDeviceProjection(d *datacenter.Device) *DeviceProjection {
	return projectionFromDevice(d, NewCreatedProjection())
}

func NewUpdatedDeviceProjection(d *datacenter.Device) *DeviceProjection {
	return projectionFromDevice(d, NewUpdatedProjection())
}

func NewDeletedDeviceProjection(d *datacenter.Device) *DeviceProjection {
	return projectionFromDevice(d, NewDeletedProjection())
}
