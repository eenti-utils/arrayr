package arrayr

import (
	"fmt"

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

	for _, e := range a {
		if ok(e) {
			r = append(r, e)
		}
	}
	return
}

// returns the first element, where the specified Qualifier function returns bool true, and
//  - bool true, if an element was qualified
//  - bool false, otherwise
func First[V any](ok typr.Qualifier[V], a ...V) (r V, f bool) {
	if ok == nil {
		if c := len(a); c > 0 {
			r = a[0]
			f = true
		}
		return
	}
	for _, e := range a {
		if ok(e) {
			r = e
			f = true
			return
		}
	}
	return
}

// returns the first element, where the specified Qualifier function returns bool true, or
// the specified default value, if no element was qualified
func FirstOr[V any](ok typr.Qualifier[V], defaultVal V, a ...V) (r V) {
	r = defaultVal
	if fv, exists := First(ok, a...); exists {
		r = fv
	}
	return
}

// returns the first element, where the specified Qualifier function returns bool true, or
// the zero value of the given type, if no element was qualified
func FirstOrZero[V any](ok typr.Qualifier[V], a ...V) (r V) {
	if fv, exists := First(ok, a...); exists {
		r = fv
	}
	return
}

// returns the last element, where the specified Qualifier function returns bool true, and
//  - bool true, if an element was qualified
//  - bool false, otherwise
func Last[V any](ok typr.Qualifier[V], a ...V) (r V, f bool) {
	l := len(a)
	if l == 0 {
		return
	}
	li := l - 1
	if ok == nil {
		r = a[li]
		f = true
		return
	}
	for i := li; i > -1; i-- {
		if e := a[i]; ok(e) {
			r = e
			f = true
			return
		}
	}
	return
}

// returns the last element, where the specified Qualifier function returns bool true, or
// the specified default value, if no element was qualified
func LastOr[V any](ok typr.Qualifier[V], defaultVal V, a ...V) (r V) {
	r = defaultVal
	if fv, exists := Last(ok, a...); exists {
		r = fv
	}
	return
}

// returns the last element, where the specified Qualifier function returns bool true, or
// the zero value of the given type, if no element was qualified
func LastOrZero[V any](ok typr.Qualifier[V], a ...V) (r V) {
	if fv, exists := Last(ok, a...); exists {
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
	l := len(a)
	if i > -1 && i < l {
		r = a[i]
	} else if i < 0 && l+i > -1 {
		r = a[l+i]
	} else {
		e = fmt.Errorf("index out of range [ %d ]", i)
	}
	return
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
	if nth, err := Nth(i, a...); err == nil {
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
	if nth, err := Nth(i, a...); err == nil {
		r = nth
	}
	return
}

// returns an array comprised of any elements that appeared more than once in the original specification
//  - if no elements appeared more than once, then the resulting array has a length of zero
//  - the resulting array, if not empty, will have unique elements
func Repeated[V comparable](a ...V) (r []V) {
	m := make(map[V]bool)
	for _, el := range a {
		if _, ex := m[el]; ! ex {
			m[el]=true
			r = append(r, el)
		}
	}
	return
}

// returns an array comprised of the elements specified, where each element appears exactly once,
// in the resulting array
func Unique[V comparable](a ...V) (r []V) {
	m := make(map[V]int)
	for _, el := range a {
		if c, ex := m[el]; ex {
			m[el] = c + 1
		} else {
			m[el] = 0
			r = append(r, el)
		}
	}
	return
}

// returns an array comprised of any elements that only appeared once in the original specification
//  - if no elements appeared only once, then the resulting array has a length of zero
//  - the resulting array, if not empty, will have unique elements
func Unrepeated[V comparable](a ...V) (r []V) {
	m := make(map[V]int)
	for _, el := range a {
		if c, ex := m[el]; ex {
			m[el] = c + 1
		} else {
			m[el] = 0
		}
	}
	for _, el := range a {
		if m[el] == 0 {
			r = append(r, el)
		}
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
