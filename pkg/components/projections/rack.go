package projections

import "github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"

type RackProjection struct {
	BaseProjection `bson:",inline"`

	ID           string `json:"id,omitempty" bson:"id,omitempty"`
	Size         int    `json:"size,omitempty" bson:"size,omitempty"`
	DatacenterId string `json:"datacenterId,omitempty" bson:"datacenterId,omitempty"`
}

func projectionFromRack(r *datacenter.Rack, base BaseProjection) *RackProjection {
	var datacenterId string
	if r.Datacenter != nil {
		datacenterId = r.Datacenter.ID
	}
	return &RackProjection{
		BaseProjection: base,
		ID:             r.ID,
		Size:           r.Size,
		DatacenterId:   datacenterId,
	}
}

func NewCreatedRackProjection(r *datacenter.Rack) *RackProjection {
	return projectionFromRack(r, NewCreatedProjection())
}

func NewUpdatedRackProjection(r *datacenter.Rack) *RackProjection {
	return projectionFromRack(r, NewUpdatedProjection())
}

func NewDeletedRackProjection(r *datacenter.Rack) *RackProjection {
	return projectionFromRack(r, NewDeletedProjection())
}
