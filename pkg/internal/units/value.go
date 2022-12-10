package units

import (
	"fmt"
	"strings"
)

// Value represents a whole number value with a unit of measurement.
type Value struct {
	mag  int
	unit Unit
}

func NewValue(mag int, u Unit) Value {
	return Value{
		mag:  mag,
		unit: u,
	}
}

// ParseValue attempts to parse the passed string into a Value, returning the parsed Value and an error
// if the string could not be parsed.
func ParseValue(s string) (Value, error) {
	var v Value
	if s == "" {
		return v, nil
	}
	n, u := ExtractInt(s)
	u = strings.ReplaceAll(u, " ", "")

	unit, err := FindUnit(u)
	if err != nil {
		return v, err
	}

	v.mag = n
	v.unit = unit
	return v, nil
}

func (v Value) String() string {
	return fmt.Sprintf("%v%s", v.mag, v.unit.Symbol)
}

func (v Value) MustConvert(u Unit) (Value, error) {
	if u.baseUnit.Name != v.unit.baseUnit.Name {
		return v, fmt.Errorf("unsupported conversion: %s to %s", u.Name, v.unit.Name)
	}

	m := v.mag * (u.baseRatio / v.unit.baseRatio)
	val := Value{
		mag:  m,
		unit: u,
	}

	return val, nil
}
