package arrayr

import "github.com/eenti-utils/typr"

// returns an array based on the specified values,
// in the specified order
func From[V any](a ...V) (r []V) {
	if len(a) > 0 {
		return a
	}
	return
}

// returns an array of paired elements based on the specified map, where
//  - the X property of each array element represents a map key
//  - the Y property of each array element represents a map value
//
// order of the resulting elements is not guaranteed
func FromMap[K comparable,V any](m map[K]V) (r []typr.Pair[K,V]) {
	for k,v := range m {
		r = append(r, typr.Pair[K, V]{X: k, Y: v})
	}
	return
}