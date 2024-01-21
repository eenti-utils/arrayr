package arrayr

import "github.com/eenti-utils/typr"

// ⚠ uses reflection
//
// returns an array of pairs based on the specified elements, where
//  - the X property of each pair represents the originally specified element
//  - the Y property of each pair represents the number of times the element occurred, in the original specification
func FrequencyR[V any](a ...V) (r []typr.Pair[V, int]) {
	if len(a) == 0 {
		return
	}

	var f []*typr.Pair[V, int]
	for _, el := range a {
		if p, exists := getPair[V, int](el, f...); exists {
			p.Y += 1
		} else {
			f = append(f, &typr.Pair[V, int]{X: el, Y: 1})
		}
	}
	if len(f) > 0 {
		cast := func(v *typr.Pair[V, int]) (r typr.Pair[V, int]) {
			return typr.Pair[V, int]{X: v.X, Y: v.Y}
		}
		r = ChangeTo[*typr.Pair[V, int], typr.Pair[V, int]](cast, f...)
	}
	return
}
