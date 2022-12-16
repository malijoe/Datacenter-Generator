package deviceTemplateAggregate

import (
	"context"
	"fmt"
	"strings"

	"github.com/malijoe/DatacenterGenerator/pkg/components/datacenter"
	eventsv1 "github.com/malijoe/DatacenterGenerator/pkg/events/v1"
)

const defaultVariant = "default"

var defaultCategories = []string{"x"}

func (a *DeviceTemplateAggregate) CreateDeviceTemplate(ctx context.Context, modelId, variant string, categories []string, hostnameTemplate, alias, function string) error {
	if modelId == "" {
		return ErrModelIdNotProvided
	}

	if variant == "" {
		variant = defaultVariant
	} else {
		variant = strings.ToLower(variant)
	}

	if len(categories) == 0 {
		if variant == defaultVariant {
			categories = defaultCategories
		} else {
			categories = []string{variant}
		}
	} else {
		for i := range categories {
			categories[i] = strings.ToLower(categories[i])
		}
	}

	alias = strings.ToLower(alias)

	parsedFunction := datacenter.UnknownFunction
	if function != "" {
		parsedFunction = datacenter.ParseFunction(function)
		// if parsedFunction is still UnknownFunction, the passed 'function' value was invalid
		if parsedFunction == datacenter.UnknownFunction {
			return fmt.Errorf("%w {%s}", ErrInvalidFunctionSpecified, function)
		}
	}

	event, err := eventsv1.NewDeviceTemplateCreatedEvent(a, modelId, variant, categories, hostnameTemplate, alias, parsedFunction)
	if err != nil {
		return err
	}

	return a.Apply(event)
}
