package arrayr

import "github.com/eenti-utils/typr"

// returns an array of pairs based on the specified elements, where
//  - the X property of each pair represents the originally specified element
//  - the Y property of each pair represents the number of times the element occurred, in the original specification
func Frequency[V comparable](a ...V) (r []typr.Pair[V, int]) {
	if len(a) == 0 {
		return
	}
	m := make(map[V]int)
	var a1 []V
	for _, el := range a {
		if c, ex := m[el]; ex {
			m[el] = c + 1
		} else {
			m[el] = 0
			a1 = append(a1, el)
		}
	}
	for _, el := range a1 {
		if c, ex := m[el]; ex {
			r = append(r, typr.Pair[V, int]{X: el, Y: c + 1})
		}
	}
	return
}
