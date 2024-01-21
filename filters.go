package arrayr

import (
	"github.com/eenti-utils/typr"
)

// applies the specified Qualifier function to each of the specified elements
//
// each element, where the Qualifier function returns bool true is added to the array
// that is returned
func Filter[V any](ok typr.Qualifier[V], a ...V) (r []V) {
	if ok == nil {
		r = a
		return
	}

	return doFilter(ok, a...)
}

// returns the first element, where the specified Qualifier function returns bool true, and
//  - bool true, if an element was qualified
//  - bool false, otherwise
func First[V any](ok typr.Qualifier[V], a ...V) (r V, f bool) {
	return doFirst(ok, a...)
}

// returns the first element, where the specified Qualifier function returns bool true, or
// the specified default value, if no element was qualified
func FirstOr[V any](ok typr.Qualifier[V], defaultVal V, a ...V) (r V) {
	r = defaultVal
	if fv, exists := doFirst(ok, a...); exists {
		r = fv
	}
	return
}

// returns the first element, where the specified Qualifier function returns bool true, or
// the zero value of the given type, if no element was qualified
func FirstOrZero[V any](ok typr.Qualifier[V], a ...V) (r V) {
	if fv, exists := doFirst(ok, a...); exists {
		r = fv
	}
	return
}

// returns the last element, where the specified Qualifier function returns bool true, and
//  - bool true, if an element was qualified
//  - bool false, otherwise
func Last[V any](ok typr.Qualifier[V], a ...V) (r V, f bool) {
	return doLast(ok, a...)
}

// returns the last element, where the specified Qualifier function returns bool true, or
// the specified default value, if no element was qualified
func LastOr[V any](ok typr.Qualifier[V], defaultVal V, a ...V) (r V) {
	r = defaultVal

	if len(a) == 0 {
		return
	}

	if fv, exists := doLast(ok, a...); exists {
		r = fv
	}
	return
}

// returns the last element, where the specified Qualifier function returns bool true, or
// the zero value of the given type, if no element was qualified
func LastOrZero[V any](ok typr.Qualifier[V], a ...V) (r V) {
	if len(a) == 0 {
		return
	}

	if fv, exists := doLast(ok, a...); exists {
		r = fv
	}
	return
}

// returns the element at the specified index, and a nil error, if the index is valid
//  - valid non-negative index values range from 0 to n-1, where n is the number of elements in the array
//  - valid negative index value range from -n to -1, where n is the number of elements in the array
// a negative index references an offset from the end of the array, such that
//  - negative 1 (-1) references the last element
//  - negative 2 (-1) references the next to last element
// and so forth
func Nth[V any](i int, a ...V) (r V, e error) {
	return doNth(i, a...)
}

// returns the element at the specified index, if the index is valid, or
// the specified default value
//
// valid index values range from -n to n-1, where n is the number of elements in the array
//
// zero references the first element of the array
//
// a negative index references an offset from the end of the array, such that
//  - negative 1 (-1) references the last element
//  - negative 2 (-2) references the next to last element
// and so forth
func NthOr[V any](i int, defaultVal V, a ...V) (r V) {
	r = defaultVal
	if nth, err := doNth(i, a...); err == nil {
		r = nth
	}
	return
}

// returns the element at the specified index, if the index is valid, or
// the zero value of the given type
//
// valid index values range from -n to n-1, where n is the number of elements in the array
//
// zero references the first element of the array
//
// a negative index references an offset from the end of the array, such that
//  - negative 1 (-1) references the last element
//  - negative 2 (-2) references the next to last element
// and so forth
func NthOrZero[V any](i int, a ...V) (r V) {
	if nth, err := doNth(i, a...); err == nil {
		r = nth
	}
	return
}

// returns an array comprised of all of the specified elements considered to be valid, by the specified Validator function
//  - if no elements were considered valid, then the resulting array has a length of zero
func Valid[V any](validate typr.Validator[V], a ...V) (r []V) {
	if validate == nil {
		if len(a) > 0 {
			r = a
		}
		return
	}
	for _, e := range a {
		if err := validate(e); err == nil {
			r = append(r, e)
		}
	}
	return
}
