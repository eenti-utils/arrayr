package arrayr

import "github.com/eenti-utils/typr"

// returns an array comprised of any elements that appeared more than once in the original specification
//  - if no elements appeared more than once, then the resulting array has a length of zero
//  - the resulting array, if not empty, will have unique elements
func Repeated[V comparable](a ...V) (r []V) {
	if len(a) == 0 {
		return
	}
	qualify := func(e typr.Pair[V, int]) bool { return e.Y > 1 }
	r = convert(pairX[V], doFilter(qualify, doFrequency(a...)...)...)
	return
}

// returns an array comprised of the elements specified, where each element appears exactly once,
// in the resulting array
func Unique[V comparable](a ...V) (r []V) {
	if len(a) == 0 {
		return
	}
	r = convert(pairX[V], doFrequency(a...)...)
	return
}

// returns an array comprised of any elements that only appeared once in the original specification
//  - if no elements appeared only once, then the resulting array has a length of zero
//  - the resulting array, if not empty, will have unique elements
func Unrepeated[V comparable](a ...V) (r []V) {
	if len(a) == 0 {
		return
	}
	qualify := func(e typr.Pair[V, int]) bool { return e.Y == 1 }
	r = convert(pairX[V], doFilter(qualify, doFrequency(a...)...)...)
	return
}
