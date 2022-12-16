package projections

import (
	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/units"
)

type DatacenterProjection struct {
	BaseProjection `bson:",inline"`

	ID        string                 `json:"id,omitempty" bson:"id,omitempty"`
	Site      string                 `json:"site,omitempty" bson:"site,omitempty"`
	Building  string                 `json:"building,omitempty" bson:"building,omitempty"`
	Room      string                 `json:"room,omitempty" bson:"room,omitempty"`
	Providers map[string]units.Value `json:"providers,omitempty" bson:"providers,omitempty"`
}

func projectionFromDatacenter(d *datacenter.Datacenter, base BaseProjection) *DatacenterProjection {
	return &DatacenterProjection{
		BaseProjection: base,
		ID:             d.ID,
		Site:           d.Site,
		Building:       d.Building,
		Room:           d.Room,
		Providers:      d.Providers,
	}
}

func NewCreatedDatacenterProjection(d *datacenter.Datacenter) *DatacenterProjection {
	return projectionFromDatacenter(d, NewCreatedProjection())
}

func NewUpdatedDatacenterProjection(d *datacenter.Datacenter) *DatacenterProjection {
	return projectionFromDatacenter(d, NewUpdatedProjection())
}

func NewDeletedDatacenterProjection(d *datacenter.Datacenter) *DatacenterProjection {
	return projectionFromDatacenter(d, NewDeletedProjection())
}
