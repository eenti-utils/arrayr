package arrayr

import "github.com/eenti-utils/typr"

// returns an array with unique elements,
// based on the elements specified
func Dedupe[V comparable](a ...V) (r []V) {
	m := make(map[V]bool)
	for _, el := range a {
		if _, ex := m[el]; !ex {
			r = append(r, el)
			m[el] = true
		}
	}
	return
}

// applies the specified change function to each element,
// and returns an array of changed elements
//
// Example
// sleet := ChangeTo[water,ice](freeze,[]water{drop01,drop02,drop03}...)
func ChangeTo[V, R any](change typr.Resolver[V, R], a ...V) (r []R) {
	if change != nil {
		for _, e := range a {
			ne := change(e)
			r = append(r, ne)
		}
	}
	return
}
