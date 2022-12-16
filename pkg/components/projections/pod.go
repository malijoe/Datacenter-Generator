package projections

import "github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"

type PodProjection struct {
	BaseProjection `bson:",inline"`

	ID           string `json:"id,omitempty" bson:"id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Function     string `json:"function,omitempty" bson:"function,omitempty"`
	Instance     int    `json:"instance,omitempty" bson:"instance,omitempty"`
	DatacenterId string `json:"datacenterId,omitempty" bson:"datacenterId,omitempty"`
}

func projectionFromPod(p *datacenter.Pod, base BaseProjection) *PodProjection {
	var datacenterId string
	if p.Datacenter != nil {
		datacenterId = p.Datacenter.ID
	}

	return &PodProjection{
		BaseProjection: base,
		ID:             p.ID,
		Name:           p.Name,
		Function:       string(p.Function),
		Instance:       p.Instance,
		DatacenterId:   datacenterId,
	}
}

func NewCreatedPodProjection(p *datacenter.Pod) *PodProjection {
	return projectionFromPod(p, NewCreatedProjection())
}

func NewUpdatedPodProejction(p *datacenter.Pod) *PodProjection {
	return projectionFromPod(p, NewUpdatedProjection())
}

func NewDeletedPodProjection(p *datacenter.Pod) *PodProjection {
	return projectionFromPod(p, NewDeletedProjection())
}
