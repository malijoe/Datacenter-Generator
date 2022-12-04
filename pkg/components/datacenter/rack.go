package datacenter

import (
	"errors"
	"fmt"
	"strings"
)

// Rack represents a rack.
type Rack struct {
	// the name of the rack
	Name string
	// the number of total RUs the rack has. (default 45)
	Size int

	// the datacenter this rack belongs to.
	Datacenter *Datacenter
	// the devices that are racked here.
	Devices []*Device
}

const defaultRackSize = 45

func NewRack(name string, size ...int) *Rack {
	rackSize := defaultRackSize
	if len(size) > 0 {
		rackSize = size[0]
	}
	return &Rack{
		Name: strings.ToLower(name),
		Size: rackSize,

		Devices: make([]*Device, rackSize),
	}
}

// CanFitDevice returns true if the Rack a valid range of RU(s) of size f.
// a range is valid if it is: composed of sequential RU(s), all RU(s) in the range are not occupied.
func (r *Rack) CanFitDevice(formFactor int) bool {
	return r.canFitDevice(formFactor, r.Size, false)
}

// CanFitDeviceAt is like CanFitDevice but returns true only if there is a valid range of RU(s) beginning at the provided elevation(el).
func (r *Rack) CanFitDeviceAt(formFactor int, el int) bool {
	return r.canFitDevice(formFactor, el, true)
}

// canFitDevice searches for a consecutive range of open RU(s) that is of size formFactor, starting at startingElevation
// and returns true on the first valid range found. the search satisfies the strict criteria if the first valid range found
// starts at startingElevation.
func (r *Rack) canFitDevice(formFactor int, startingElevation int, strict bool) bool {
	var (
		// tracks the number of consecutive open RU(s) in the currently tracked range.
		openRUCount = 0
		// the last index of the valid strict range
		strictLimit = startingElevation - (formFactor - 1)
	)
	for i := startingElevation; i >= 0; i-- {
		if r.Devices[i] == nil {
			openRUCount++
		} else if openRUCount > 0 {
			// if the RU at i is not empty and there are open RU(s) being tracked
			// reset the counter.
			openRUCount = 0
		}

		switch {
		case openRUCount >= formFactor:
			return true
		case i >= strictLimit && strict:
			// the search has failed the strict criteria.
			return false
		}
	}
	return false
}

// RackDevice attempts to put the passed device into the Rack and returns an error if there is no space
// available for the device.
func (r *Rack) RackDevice(device Device) error {
	return r.rackDevice(device)
}

// RackDeviceAt is like RackDevice but returns an error if the device cannot be racked at the elevation passed.
func (r *Rack) RackDeviceAt(device Device, el int) error {
	// rackDevice accepts the desired index the device is to be inserted. elevation = index+1
	return r.rackDevice(device, el-1)
}

var ErrUnableToFitDevice = errors.New("unable to fit device")

// rackDevice racks a device into the Rack and returns an error if no there is no space available for the device.
// if a value is passed for 'at', the device will attempt to be racked at the passed index and returns an error
// if it cannot fit there.
func (r *Rack) rackDevice(device Device, at ...int) error {
	// the beginning and end of the RU range to insert the device at.
	start := -1

	if strict := len(at) > 0 && at[0] > 0; strict {

		if !r.CanFitDeviceAt(device.Model.FormFactor, at[0]) {
			return fmt.Errorf("%w: cannot fit a device of size %d at elevation %d", ErrUnableToFitDevice, device.Model.FormFactor, at[0])
		}
		start = at[0]

	} else {

		for i := len(r.Devices) - 1; i >= 0; i-- {
			d := r.Devices[i]
			if d == nil {
				if r.CanFitDeviceAt(device.Model.FormFactor, i) {
					start = i
					break
				}
				i -= device.Model.FormFactor
			}
			i -= d.Model.FormFactor
		}

	}

	if start < 0 {
		return fmt.Errorf("%w: could not find a valid range of RU(s) to fit a device of size %d", ErrUnableToFitDevice, device.Model.FormFactor)
	}

	device.Elevation = start + 1
	end := start - device.Model.FormFactor
	for i := start; i > end; i-- {
		r.Devices[i] = &device
	}
	return nil
}
