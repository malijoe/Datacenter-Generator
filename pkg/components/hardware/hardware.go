package hardware

type HardwareModel struct {
	// the unique identifier for the hardware model
	ID string

	PID        string
	FormFactor int
	Weight     float32

	PortGroups []*PortGroup
}
