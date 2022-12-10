package rackAggregate

import "errors"

var (
	ErrRackNameNotSpecified        = errors.New("rack name not specified")
	ErrDatacenterIDNotProvided     = errors.New("datacenterId not provided")
	ErrDeviceIDNotProvided         = errors.New("deviceId not provided")
	ErrDeviceFormFactorNotProvided = errors.New("device form factor not provided")
)
