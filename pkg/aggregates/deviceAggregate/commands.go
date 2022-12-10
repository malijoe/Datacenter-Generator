package deviceAggregate

import (
	"context"
	"fmt"

	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
)

func (a *DeviceAggregate) CreateDevice(ctx context.Context, template *datacenter.DeviceTemplate, dc *datacenter.Datacenter, rack *datacenter.Rack, pod *datacenter.Pod, elevation int, cluster int, designation string) error {
	// get current number of relevant instances for device
	n := dc.NumDeviceInstances(template.Model.PID, template.Variant)
	// add 1 to account for adding this device
	n += 1

	if elevation != 0 {
		if !rack.CanFitDeviceAt(template.Model.FormFactor, elevation) {
			return fmt.Errorf("%w {%s}, elevation {%d}, formFactor: {%d}", ErrCantFitDeviceInRack, rack.Name, elevation, template.Model.FormFactor)
		}
	} else {
		el, canFit := rack.CanFitDevice(template.Model.FormFactor)
		if !canFit {
			return fmt.Errorf("%w {%s}, formFactor: {%d}", ErrCantFitDeviceInRack, rack.Name, template.Model.FormFactor)
		}
		elevation = el
	}

	parsedDesignation := datacenter.UnknownDesignation
	if designation != "" {
		parsedDesignation = datacenter.ParseDesignation(designation)
		if parsedDesignation == datacenter.UnknownDesignation {
			return fmt.Errorf("%w designation: {%s}", ErrInvalidDesignationSpecified, designation)
		}
	}

	var podId string
	if pod != nil {
		podId = pod.ID
		if template.Function != datacenter.UnknownFunction && pod.Function != template.Function {
			return fmt.Errorf("%w: pod function does not match function specified by template. pod: {%s}, template {%s}", ErrFunctionConflict, pod.Function, template.Function)
		}
	}

	hostname, err := template.TemplateHostname(datacenter.NewHostnameTemplateVars(dc.Site, pod.Function, pod.Instance, rack.Name, parsedDesignation, elevation, n))
	if err != nil {
		return fmt.Errorf("TemplateHostname: %w", err)
	}

	event, err := eventsv1.NewDeviceCreatedEvent(a, hostname, elevation, parsedDesignation, cluster, n, template.Model.ID, template.Categories, podId, rack.ID)
	if err != nil {
		return err
	}

	return a.Apply(event)
}
