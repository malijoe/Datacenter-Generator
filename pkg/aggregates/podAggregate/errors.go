package podAggregate

import "errors"

var (
	ErrFunctionNotSpecified     = errors.New("function not specified")
	ErrInvalidFunctionSpecified = errors.New("invalid function specified")
)
