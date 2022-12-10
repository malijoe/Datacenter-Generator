package logger

import "fmt"

const (
	defaultJSONOutput  = false
	defaultOutputLevel = "info"
)

// Options defines the sets of options for logging.
type Options struct {
	// JSONFormatEnabled is the flag to enable JSON formatted logs.
	JSONFormatEnabled bool

	// OutputLevel is the level of logging.
	OutputLevel string
}

// SetOutputLevel sets the log output level.
func (o *Options) SetOutputLevel(lvl string) error {
	if toLogLevel(lvl) == UndefinedLevel {
		return fmt.Errorf("undefined Log Output Level: %s", lvl)
	}
	o.OutputLevel = lvl
	return nil
}

func (o *Options) AttachCmdFlags(
	stringVar func(p *string, name string, value string, usage string),
	boolVar func(p *bool, name string, value bool, usage string),
) {
	if stringVar != nil {
		stringVar(
			&o.OutputLevel,
			"log-level",
			defaultOutputLevel,
			"Options are debug, info, warn, error, or fatal (default info)",
		)
	}
	if boolVar != nil {
		boolVar(
			&o.JSONFormatEnabled,
			"log-as-json",
			defaultJSONOutput,
			"print log as JSON (default false)",
		)
	}
}

func DefaultOptions() Options {
	return Options{
		JSONFormatEnabled: defaultJSONOutput,
		OutputLevel:       defaultOutputLevel,
	}
}

func ApplyOptionsToLoggers(options *Options) error {
	internalLoggers := getLoggers()

	// apply formatting options first
	for _, v := range internalLoggers {
		v.EnableJSONOutput(options.JSONFormatEnabled)
	}

	lvl := toLogLevel(options.OutputLevel)
	if lvl == UndefinedLevel {
		return fmt.Errorf("invalid value for --log-level: %s", options.OutputLevel)
	}

	for _, v := range internalLoggers {
		v.SetOutputLevel(lvl)
	}
	return nil
}
