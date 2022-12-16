package projections

import "github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"

type DeviceTemplateProjection struct {
	BaseProjection `bson:",inline"`

	ID               string   `json:"id,omitempty" bson:"id,omitempty"`
	Variant          string   `json:"variant,omitempty" bson:"variant,omitempty"`
	Function         string   `json:"function,omitempty" bson:"function,omitempty"`
	Categories       []string `json:"categories,omitempty" bson:"categories,omitempty"`
	HostnameTemplate string   `json:"hostnameTemplate,omitempty" bson:"hostnameTemplate,omitempty"`
	Alias            string   `json:"alias,omitempty" bson:"alias,omitempty"`

	ModelId string `json:"modelId,omitempty" bson:"modelId,omitempty"`
}

func projectionFromDeviceTemplate(dt *datacenter.DeviceTemplate, base BaseProjection) *DeviceTemplateProjection {
	return &DeviceTemplateProjection{
		BaseProjection:   base,
		ID:               dt.ID,
		Variant:          dt.Variant,
		Function:         string(dt.Function),
		Categories:       dt.Categories,
		HostnameTemplate: dt.HostnameTemplate,
		Alias:            dt.Alias,
		ModelId:          dt.Model.ID,
	}
}

func NewCreatedDeviceTemplateProjection(dt *datacenter.DeviceTemplate) *DeviceTemplateProjection {
	return projectionFromDeviceTemplate(dt, NewCreatedProjection())
}

func NewUpdatedDeviceTemplateProjection(dt *datacenter.DeviceTemplate) *DeviceTemplateProjection {
	return projectionFromDeviceTemplate(dt, NewUpdatedProjection())
}

func NewDeletedDeviceTemplateProjection(dt *datacenter.DeviceTemplate) *DeviceTemplateProjection {
	return projectionFromDeviceTemplate(dt, NewDeletedProjection())
}
