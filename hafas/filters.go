package hafas

import "github.com/mboufous/berlin-departure-board/util"

type LineFilter map[string]bool

func NewLineFilter(values ...string) LineFilter {
	f := LineFilter{}
	for _, value := range values {
		f[util.Normalize(value)] = true
	}
	return f
}

func (f LineFilter) Filter(value string) bool {
	if len(f) == 0 {
		return true // Allow all if filter is empty
	}
	_, exists := f[util.Normalize(value)]
	return exists
}
