package zendesk

import (
	"fmt"
	"strconv"
)

var (
	NotImplementedErr = fmt.Errorf("feature not yet implemented")
)

func toStringArray(ints []int) []string {
	values := make([]string, len(ints))
	for index, i := range ints {
		values[index] = strconv.Itoa(i)
	}

	return values
}
