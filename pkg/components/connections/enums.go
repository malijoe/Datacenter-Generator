package connections

import "github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"

type Designation string

const (
	UnknownDesignation   = Designation(datacenter.UnknownDesignation)
	PrimaryDesignation   = Designation(datacenter.PrimaryDesignation)
	SecondaryDesignation = Designation(datacenter.SecondaryDesignation)
	MatchDesignation     = "="
	OppositeDesignation  = "!"
)
