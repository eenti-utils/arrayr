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

// applies each element to the specified resolver function,
// and returns an array of resolved elements, in the desired type
//
// Example
// sleet := Modify[water,ice](freeze,[]water{drop01,drop02,drop03}...)
func Modify[V,R any](resolve typr.Resolver[V,R],a ...V) (r []R) {
	if resolve != nil {
		for _, e := range a {
			ne := resolve(e)
			r = append(r, ne)
		}
	}
	return
}
