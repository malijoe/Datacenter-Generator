package ranges

import "strings"

type Modifier string

const (
	UnknownModifier Modifier = "unknown"
	EvenModifier    Modifier = "even"
	OddModifier     Modifier = "odd"
)

func ParseModifier(s string) Modifier {
	switch strings.ToLower(s) {
	case "even":
		return EvenModifier
	case "odd":
		return OddModifier
	}

	return UnknownModifier
}

func (m Modifier) modify(filter numFilter) numFilter {
	modifier := func(i int) bool {
		mod := i % 2
		switch m {
		case EvenModifier:
			return mod == 0
		case OddModifier:
			return mod != 0
		}
		// default case defer to the original filter
		return true
	}
	return func(i int) bool {
		if filter(i) {
			return modifier(i)
		}
		return false
	}
}
