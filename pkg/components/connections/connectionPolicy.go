package connections

import (
	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	"github.com/malijoe/DatacenterGenerator/pkg/components/hardware"
)

// ConnectionPolicy defines when and how to make connections between 2 or more device categories.
type ConnectionPolicy struct {
	// the name of the connection policy.
	Name string
	// the device category that connections will originate from
	Origin string
	// defines the boundary to search for devices
	Boundary BoundarySpecification
	// defines the requirements for valid terminal devices
	TerminalSpecification TerminalSpecification
	// defines the connections to be made by this policy.
	Connections ConnectionSpecification
	// defines the optional ACI indicators for this policy's connections.
	ACI ACISpecification
}

type TerminalSpecification struct {
	// the device categories (in order of priority) that are valid terminal device categories.
	Priority []string
	// the device categories that should be excluded from valid terminal devices.
	Exclude []string
	// the number of terminal devices to connect
	Quantity int
	// indicates that valid terminal devices must match the function value. ignored if function value is unspecified.
	Function datacenter.Function
}

type ACISpecification struct {
	// a template for an aci indicator to be applied to the origin end of the connection
	OriginIndicator string
	// a template for an aci indicator to be applied to the terminal end of the connection
	TerminalIndicator string
}

type BoundarySpecification struct {
	// indicates that valid terminal devices should be in the rack/not in the rack of the origin device
	// (true/false respectively). if no value is provided the rack is ignored.
	MatchRack *bool
	// indicates that valid terminal devices should be in/not in the same cluster as the origin device
	// (true/false respectively). if no value is provided the cluster is ignored.
	MatchCluster *bool
}

type ConnectionSpecification struct {
	// the medium to be used when making connections
	Medium hardware.Medium
	// an alternate medium to use for specified device categories.
	FallbackMedium struct {
		Categories []string
		Medium     hardware.Medium
	}
	// the number of ports to connect
	PortQuantity int
	// the ports that should be used for origin devices
	OriginPorts []PortSpecification
	// the ports that should be used for terminal devices
	TerminalPorts []PortSpecification
}

type PortSpecification struct {
	// indicates the designation of the device that the port ranges belong too.
	Designation Designation
	// the port ranges for connections to be made.
	PortRanges []PortRange
}
