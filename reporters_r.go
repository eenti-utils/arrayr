package arrayr

import "github.com/eenti-utils/typr"

// ⚠ uses reflection
//
// returns an array of pairs based on the specified elements, where
//  - the X property of each pair represents the originally specified element
//  - the Y property of each pair represents the number of times the element occurred, in the original specification
func FrequencyR[V any](a ...V) (r []typr.Pair[V,int]) {
	if len(a) == 0 {
		return
	}

	return doFrequencyR(a...)
}
