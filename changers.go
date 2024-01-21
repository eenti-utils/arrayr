package arrayr

import "github.com/eenti-utils/typr"

// applies the specified change function to each element,
// and returns an array of changed elements
//
// Example
// sleet := ChangeTo[water,ice](freeze,[]water{drop01,drop02,drop03}...)
func ChangeTo[V, R any](change typr.Resolver[V, R], a ...V) (r []R) {
	if change != nil {
		r = convert(change, a...)
	}
	return
}

// returns an array of the specified elements in reverse order
func Reverse[V any](a ...V) (r []V) {
	if l := len(a); l > 1 {
		if l == 2 {
			r = []V{a[1], a[0]}
			return
		}
		li := l - 1
		for i := li; i > -1; i-- {
			r = append(r, a[i])
		}
	} else {
		r = a
	}
	return
}
