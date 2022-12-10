package datacenter

// Pod is a value object that represents a logical grouping of devices.
type Pod struct {
	// the unique identifier for the pod.
	ID string

	// the name is found by concatenating the string representation of the pod's
	// function and instance number (i.e. 'service0').
	Name string
	// the instance number of the pod. represents the number of pods with a shared
	// function in the same datacenterAggregate.
	Instance int
	// the function of this pod. pods cannot have a function value of UnknownFunction.
	Function Function

	// the datacenter the pod belongs to
	Datacenter *Datacenter
}

func NewPod() *Pod {
	return &Pod{}
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
