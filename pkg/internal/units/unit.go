package units

import "fmt"

var unitRegistry = make(map[string]Unit)

// Unit represents a unit of measurement
type Unit struct {
	// the name of the unit (e.g. bits)
	Name string
	// the symbol of the unit (e.g. b, Mb, Gb, ...)
	Symbol string

	// the base unit (e.g. the base unit for Gigabits would be bits)
	baseUnit *Unit
	// the ratio of base unit to the unit as a whole number (e.g. the ratio of bits to Megabits is 1000 b/ 1 Mb = 1000)
	baseRatio int
}

func NewBaseUnit(name, symbol string) Unit {
	if u, ok := unitRegistry[name]; ok {
		return u
	}

	u := Unit{
		Name:      name,
		Symbol:    symbol,
		baseUnit:  nil,
		baseRatio: 1,
	}
	unitRegistry[name] = u
	return u
}

func (u Unit) IsBaseUnit() bool {
	return u.baseUnit == nil
}

func FindUnit(u string) (Unit, error) {
	unit, ok := unitRegistry[u]
	if ok {
		return unit, nil
	}

	for _, unit = range unitRegistry {
		if unit.Symbol == u {
			return unit, nil
		}
	}

	return Unit{}, fmt.Errorf("no unit found with name or symbol '%s'", u)
}
