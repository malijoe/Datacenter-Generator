package datacenter

import (
	"fmt"
	"strings"
)

type Function string

const (
	UnknownFunction Function = "unspecified"
	ComputeFunction Function = "compute"
	EdgeFunction    Function = "edge"
	ServiceFunction Function = "service"
	StorageFunction Function = "storage"
)

// ParseFunction parses the passed string into a Function.
// UnknownFunction is returned if the input doesn't match any valid Function values.
func ParseFunction(s string) Function {
	switch strings.ToLower(s) {
	case "compute", "cpu":
		return ComputeFunction
	case "edge":
		return EdgeFunction
	case "service", "svc":
		return ServiceFunction
	case "storage", "strg":
		return StorageFunction
	}

	// unrecognized input
	return UnknownFunction
}

// Abv returns the abreviated version of the Function.
func (f Function) Abv() string {
	switch f {
	case ComputeFunction:
		return "cpu"
	case EdgeFunction:
		return "edge"
	case ServiceFunction:
		return "svc"
	case StorageFunction:
		return "strg"
	}

	return "n/a"
}

type Designation string

const (
	UnknownDesignation   Designation = "unspecified"
	PrimaryDesignation   Designation = "primary"
	SecondaryDesignation Designation = "secondary"
)

// ParseDesignation parses the passed string into a Designation.
// UnknownDesignation is returned if the input doesn't match any valid Designation values.
func ParseDesignation(s string) Designation {
	switch strings.ToLower(s) {
	case "primary", "a":
		return PrimaryDesignation
	case "secondary", "b":
		return SecondaryDesignation
	}

	// unrecognized input
	return UnknownDesignation
}

// Alpha returns the alphabetic representation of the Designation.
// PrimaryDesignation -> 'a'
// SecondaryDesignation -> 'b'
// UnknownDesignation -> 'o'
func (d Designation) Alpha() string {
	switch d {
	case PrimaryDesignation:
		return "a"
	case SecondaryDesignation:
		return "b"
	}

	return "o"
}

// Pod is a value object that represents a logical grouping of devices.
type Pod struct {
	// the name is found by concatenating the string representation of the pod's
	// function and instance number (i.e. 'service0').
	Name string
	// the instance number of the pod. represents the number of pods with a shared
	// function in the same datacenter.
	Instance int
	// the function of this pod. pods cannot have a function value of UnknownFunction.
	Function Function
}

func NewPod(instance int, function Function) Pod {
	return Pod{
		Name:     fmt.Sprintf("%s%d", function, instance),
		Instance: instance,
		Function: function,
	}
}

// IsZero returns true if the pod is the zero value.
func (p Pod) IsZero() bool {
	return p.Function == UnknownFunction
}

// Less returns true if: p (the calling pod) < b (the passed pod).
func (p Pod) Less(b Pod) bool {
	if p.Function == b.Function {
		return p.Instance < b.Instance
	}
	return false
}
