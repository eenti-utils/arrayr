package arrayr

import "github.com/eenti-utils/typr"

// ⚠ uses reflection
//
// returns an array comprised of any elements that appeared more than once in the original specification
//  - if no elements appeared more than once, then the resulting array has a length of zero
//  - the resulting array, if not empty, will have unique elements
func RepeatedR[V any](a ...V) (r []V) {
	if len(a) > 0 {
		qualify := func(p typr.Pair[V, int]) (r bool) { return p.Y > 1 }
		r = ChangeTo(convertFreqToX[V], Filter(qualify, FrequencyR(a...)...)...)
	}
	return
}

// ⚠ uses reflection
//
// returns an array comprised of the elements specified, where each element appears exactly once,
// in the resulting array
func UniqueR[V any](a ...V) (r []V) {
	if len(a) > 0 {
		r = ChangeTo(convertFreqToX[V], FrequencyR(a...)...)
	}
	return
}

// ⚠ uses reflection
//
// returns an array comprised of any elements that only appeared once in the original specification
//  - if no elements appeared only once, then the resulting array has a length of zero
//  - the resulting array, if not empty, will have unique elements
func UnrepeatedR[V any](a ...V) (r []V) {
	if len(a) > 0 {
		qualify := func(p typr.Pair[V, int]) (r bool) { return p.Y == 1 }
		r = ChangeTo(convertFreqToX[V], Filter(qualify, FrequencyR(a...)...)...)
	}
	return
}
