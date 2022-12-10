package deviceAggregate

import "errors"

var (
	ErrCantFitDeviceInRack         = errors.New("can't fit device in rack")
	ErrInvalidDesignationSpecified = errors.New("invalid designation specified")
	ErrFunctionConflict            = errors.New("function conflict")
)
