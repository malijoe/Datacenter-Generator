package deviceTemplateAggregate

import "errors"

var (
	ErrModelIdNotProvided       = errors.New("modelId not provided")
	ErrInvalidFunctionSpecified = errors.New("invalid function specified")
)
