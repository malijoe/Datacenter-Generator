package datacenter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/malijoe/DatacenterGenerator/pkg/components/hardware"
	"github.com/malijoe/DatacenterGenerator/pkg/internal/templar"
)

type DeviceTemplate struct {
	// the unique identifier for the device template.
	ID string

	// the variant of hardware model that is this device template.
	Variant string
	// the categories to be associated with devices created using this template.
	Categories []string
	// a template for generating hostnames for devices created using this template.
	HostnameTemplate string
	// the function to apply to devices created using this template
	Function Function

	// an alias used to reference this device template.
	Alias string

	// the hardware model that this template creates devices with.
	Model hardware.HardwareModel
}

func NewDeviceTemplate() *DeviceTemplate {
	return &DeviceTemplate{}
}

func (t *DeviceTemplate) TemplateHostname(vars HostnameTemplateVars) (string, error) {
	if err := vars.canProcessTemplate(t.HostnameTemplate); err != nil {
		return "", err
	}

	hostname, err := templar.TemplateString(t.HostnameTemplate, vars)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hostname), nil
}

type HostnameTemplateVars struct {
	site        string
	function    Function
	podInstance int
	rackName    string
	designation Designation
	elevation   int
	instance    int
}

func NewHostnameTemplateVars(site string, function Function, podInstance int, rackName string, designation Designation, elevation int, instance int) HostnameTemplateVars {
	return HostnameTemplateVars{
		site:        site,
		function:    function,
		podInstance: podInstance,
		rackName:    rackName,
		designation: designation,
		elevation:   elevation,
		instance:    instance,
	}
}

func (htv HostnameTemplateVars) Site() string {
	return htv.site
}

func (htv HostnameTemplateVars) Type() string {
	return string(htv.function)
}

func (htv HostnameTemplateVars) Pod() int {
	return htv.podInstance
}

func (htv HostnameTemplateVars) Rack() string {
	return htv.rackName
}

func (htv HostnameTemplateVars) AB() string {
	return htv.designation.Alpha()
}

func (htv HostnameTemplateVars) Elevation() string {
	return fmt.Sprintf("%02d", htv.elevation)
}

func (htv HostnameTemplateVars) Number() string {
	return fmt.Sprintf("%02d", htv.instance)
}

const (
	siteTemplateVar        = "Site"
	functionTemplateVar    = "Type"
	podTemplateVar         = "Pod"
	rackTemplateVar        = "Rack"
	designationTemplateVar = "AB"
	elevationTemplateVar   = "Elevation"
	instanceTemplateVar    = "Number"
)

var ErrMissingHostnameTemplateVarValue = errors.New("missing value for hostname template variable")

func (htv HostnameTemplateVars) canProcessTemplate(template string) error {
	if !templateHasVars(template) {
		return nil
	}

	values := map[string]bool{
		siteTemplateVar:        htv.site != "",
		functionTemplateVar:    htv.function != UnknownFunction,
		podTemplateVar:         htv.podInstance >= 0,
		rackTemplateVar:        htv.rackName != "",
		designationTemplateVar: htv.designation != UnknownDesignation,
		elevationTemplateVar:   htv.elevation != 0,
		instanceTemplateVar:    htv.instance != 0,
	}

	params := map[string]bool{
		siteTemplateVar:        templateHasVar(template, siteTemplateVar),
		functionTemplateVar:    templateHasVar(template, functionTemplateVar),
		podTemplateVar:         templateHasVar(template, podTemplateVar),
		rackTemplateVar:        templateHasVar(template, rackTemplateVar),
		designationTemplateVar: templateHasVar(template, designationTemplateVar),
		elevationTemplateVar:   templateHasVar(template, elevationTemplateVar),
		instanceTemplateVar:    templateHasVar(template, instanceTemplateVar),
	}

	var templateVarErrs error
	for key, hasVar := range params {
		if hasVar && !values[key] {
			templateVarErrs = multierror.Append(templateVarErrs, fmt.Errorf("%w {%s}", ErrMissingHostnameTemplateVarValue, key))
		}
	}
	return templateVarErrs
}

func templateHasVar(template, key string) bool {
	if !templateHasVars(template) {
		return false
	}
	return strings.Contains(template, key)
}

func templateHasVars(template string) bool {
	return strings.Contains(template, "{")
}
