package units

import "math"

type magnitude struct {
	Symbol string
	Prefix string
	Power  int
}

func (m magnitude) makeUnit(base Unit) Unit {
	var pow int
	if !base.IsBaseUnit() {
		base = *base.baseUnit
		pow = base.baseRatio
	}

	name := m.Prefix + base.Name
	symbol := m.Symbol + base.Symbol

	pow = m.Power - pow
	ratio := 1 * int(math.Pow(10.0, float64(pow)))

	u := Unit{
		Name:      name,
		Symbol:    symbol,
		baseUnit:  &base,
		baseRatio: ratio,
	}
	unitRegistry[name] = u
	return u
}

var mags = map[int]magnitude{
	18:  {"E", "exa", 18},
	15:  {"P", "peta", 15},
	12:  {"T", "tera", 12},
	9:   {"G", "giga", 9},
	6:   {"M", "mega", 6},
	3:   {"K", "kilo", 3},
	2:   {"H", "hecto", 2},
	1:   {"Da", "deca", 1},
	-1:  {"d", "deci", -1},
	-2:  {"c", "centi", -2},
	-3:  {"m", "milli", -3},
	-6:  {"μ", "micro", -6},
	-9:  {"n", "nano", -9},
	-12: {"p", "pico", -12},
	-15: {"f", "femto", -15},
	-18: {"a", "atto", -18},
}

func Exa(b Unit) Unit   { return mags[18].makeUnit(b) }
func Peta(b Unit) Unit  { return mags[15].makeUnit(b) }
func Tera(b Unit) Unit  { return mags[12].makeUnit(b) }
func Giga(b Unit) Unit  { return mags[9].makeUnit(b) }
func Mega(b Unit) Unit  { return mags[6].makeUnit(b) }
func Kilo(b Unit) Unit  { return mags[3].makeUnit(b) }
func Hecto(b Unit) Unit { return mags[2].makeUnit(b) }
func Deca(b Unit) Unit  { return mags[1].makeUnit(b) }
func Deci(b Unit) Unit  { return mags[-1].makeUnit(b) }
func Centi(b Unit) Unit { return mags[-2].makeUnit(b) }
func Milli(b Unit) Unit { return mags[-3].makeUnit(b) }
func Micro(b Unit) Unit { return mags[-6].makeUnit(b) }
func Nano(b Unit) Unit  { return mags[-9].makeUnit(b) }
func Pico(b Unit) Unit  { return mags[-12].makeUnit(b) }
func Femto(b Unit) Unit { return mags[-15].makeUnit(b) }
func Atto(b Unit) Unit  { return mags[-18].makeUnit(b) }
