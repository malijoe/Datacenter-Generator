package datacenter

import "github.com/malijoe/DatacenterGenerator/pkg/components/hardware"

type DeviceTemplate struct {
	// the variant of hardware model that is this device template.
	Variant string
	// the categories to be associated with devices created using this template.
	Categories []string
	// a template for generating hostnames for devices created using this template.
	HostnameTemplate string

	// an alias used to reference this device template.
	Alias string

	// the hardware model that this template creates devices with.
	Model hardware.HardwareModel
}
