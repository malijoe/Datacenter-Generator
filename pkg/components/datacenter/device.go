package datacenter

import "github.com/malijoe/DatacenterGenerator/pkg/components/hardware"

// Device represents a device.
type Device struct {
	// the unique identifier for the device.
	ID string

	// the hostname of the device.
	Hostname string
	// the elevation of the device. corresponds to the first RU that this device occupies.
	Elevation int
	// the designation given to this device. (primary/secondary/unspecified)
	Designation Designation
	// the cluster number of the device. (a value of 0 is unclustered)
	Cluster int
	// the instance number for this device. represents the number of instances
	// with the same config as this device.
	Instance int
	// the hardware model that is this device.
	Model hardware.HardwareModel

	// the categories this device falls under.
	Categories []string

	// the pod this device belongs to
	Pod *Pod
	// the rack this device is located in.
	Rack *Rack
	// the datacenterAggregate this device is located in.
	Datacenter *Datacenter
}

func NewDevice() *Device {
	return &Device{}
}
