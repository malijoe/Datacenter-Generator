package ranges

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/exp/slices"
)

const (
	modifierSeparator = ":"
	rangeWildcard     = "*"
	boundSeparator    = "-"
	sequenceSeparator = ","
)

// Range describes a range of numbers (int)
type Range interface {
	// InRange returns true if the passed number is within the range.
	InRange(i int) bool
	fmt.Stringer
}

var ErrMalformedRange = errors.New("malformed range")

// ParseRange attempts to parse the passed string into a Range implementation,
// and returns the Range implementation parsed, and an error if one occurred.
// convention:
// `1-5` consecutive range, min: 1, max: 5 (both upper and lower boundaries are inclusive)
// '1,2,4,5' non-consecutive range
// '*' signifies any number is valid
// '10-*' denotes a continuous range with no upper boundary (inclusive)
// '*-20' denotes a continuous range with no lower boundary (inclusive).
//  *note for all boundaries valid numbers are defined as: any number x, where x >= 0.
// modifiers: 'even', 'odd'
// 'even:1-5' only even numbers from this consecutive range
// 'odd:1,2,4,5' only odd numbers from this non-consecutive range
// 'even:*' signifies any even number is valid
func ParseRange(s string) (Range, error) {
	s = strings.ToLower(s)
	original := strings.Clone(s)
	// parse and strip modifier from the string if one exists.
	mod := UnknownModifier
	if i := strings.Index(s, modifierSeparator); i >= 0 {
		mStr := s[:i]
		mod = ParseModifier(mStr)
		s = s[i+1:]
	}

	var r Range
	switch {
	case strings.Contains(s, boundSeparator):
		var max, min *int

		parts := strings.Split(s, boundSeparator)
		if len(parts) != 2 {
			return nil, fmt.Errorf("%w {%s}: too many range boundaries parsed", ErrMalformedRange, original)
		}

		for part, setter := range map[string]func(int){parts[0]: func(i int) { min = &i }, parts[1]: func(i int) { max = &i }} {
			if part != rangeWildcard {
				num, err := strconv.Atoi(part)
				if err != nil {
					return nil, fmt.Errorf("%w {%s}: error: {%v}", ErrMalformedRange, original, err)
				}
				setter(num)
			}
		}
		if min != nil && max != nil && *min > *max {
			min, max = max, min
		}

		r = NewBoundedRange(min, max)
	case strings.Contains(s, sequenceSeparator):
		parts := strings.Split(s, ",")
		values := make([]int, len(parts))

		var strconvErrs error
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				strconvErrs = multierror.Append(strconvErrs, err)
				continue
			}
			values[i] = num
		}
		if strconvErrs != nil {
			return nil, fmt.Errorf("%w {%s}: error {%v}", ErrMalformedRange, original, strconvErrs)
		}
		r = NewSpecificRange(values)
	case s == rangeWildcard:
		r = NewUnboundedRange()
	}
	if mod != UnknownModifier {
		r = NewModifiedRange(mod, r)
	}
	return r, nil
}

// unboundedRange represents an unbounded range of sequential numbers.
type unboundedRange struct{}

func NewUnboundedRange() *unboundedRange {
	return &unboundedRange{}
}

func (*unboundedRange) InRange(_ int) bool {
	return true
}

func (*unboundedRange) String() string {
	return "*"
}

// boundedRange represents a range of sequential numbers with at least an upper or lower limit.
type boundedRange struct {
	min, max                   int
	maxUnbounded, minUnbounded bool
}

// NewBoundedRange returns a Range implementation with bounded limit(s). To leave a limit unbounded
// pass a nil value for either min/max (lower/upper limit respectively).
func NewBoundedRange(min, max *int) *boundedRange {
	var ma, mi int
	if min != nil {
		mi = *min
	}
	if max != nil {
		ma = *max
	}
	return &boundedRange{
		min:          mi,
		max:          ma,
		maxUnbounded: max == nil,
		minUnbounded: min == nil,
	}
}

func (r *boundedRange) InRange(x int) bool {
	maxFilter := defaultFilter
	if !r.maxUnbounded {
		maxFilter = func(i int) bool {
			return i <= r.max
		}
	}

	minFilter := defaultFilter
	if !r.minUnbounded {
		minFilter = func(i int) bool {
			return r.min <= i
		}
	}

	return compositeNumFilter(maxFilter, minFilter)(x)
}

func (r *boundedRange) String() string {
	maxPiece := strconv.Itoa(r.max)
	if r.maxUnbounded {
		maxPiece = "*"
	}

	minPiece := strconv.Itoa(r.min)
	if r.minUnbounded {
		minPiece = "*"
	}
	return strings.Join([]string{minPiece, maxPiece}, "-")
}

// specificRange represents a range of specific integer values that are not necessarily sequential.
type specificRange struct {
	values []int
}

// NewSpecificRange returns a Range implementation that consists of only the values passed.
func NewSpecificRange(values []int) *specificRange {
	return &specificRange{values: values}
}

func (r *specificRange) InRange(x int) bool {
	return slices.Contains(r.values, x)
}

func (r *specificRange) String() string {
	pieces := make([]string, len(r.values))
	for i := range r.values {
		pieces[i] = strconv.Itoa(r.values[i])
	}
	return strings.Join(pieces, ",")
}

// modifiedRange represents a range with an applied modifier.
type modifiedRange struct {
	modifier Modifier
	r        Range
}

// NewModifiedRange returns a Range implementation that is modified using the passed Modifier.
func NewModifiedRange(modifier Modifier, r Range) *modifiedRange {
	return &modifiedRange{modifier: modifier, r: r}
}

func (r *modifiedRange) InRange(i int) bool {
	return r.modifier.modify(r.r.InRange)(i)
}

func (r *modifiedRange) String() string {
	base := r.r.String()
	if r.modifier != UnknownModifier {
		return strings.Join([]string{string(r.modifier), base}, ":")
	}
	return base
}
