package ranges

type numFilter func(int) bool

var defaultFilter numFilter = func(_ int) bool {
	return true
}

func compositeNumFilter(filters ...numFilter) numFilter {
	return func(i int) bool {
		for _, filter := range filters {
			if !filter(i) {
				return false
			}
		}
		return true
	}
}
