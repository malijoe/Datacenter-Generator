package datacenterAggregate

import "errors"

var (
	ErrSiteNotSpecified             = errors.New("site not specified")
	ErrInvalidProviderTransferSpeed = errors.New("invalid provider transfer speed")
	ErrPodIDNotProvided             = errors.New("podID not provided")
	ErrRackIDNotProvided            = errors.New("rackID not provided")
)
