package arrayr

import "github.com/eenti-utils/typr"

type directionalRanger[V any] func(rf typr.Ranger[V], a []V)

// directives and flags passed to the Range function 
type RangeOpts[V any] struct {
	// a non-zero int value, where
	//  - positive values imply ranging in order, from index 0 to n
	//  - negative values imply ranging in order, from index n to 0
	// the value 1 (default) means handle each element, starting at index 0
	//
	// the value 2 means handle every second element, starting at index 0
	Step int
	// a function that returns bool true, when the given element
	// should be handled by the user-defined Ranger function, and bool false, otherwise
	//
	// filtered elements are not passed on to the user-defined Ranger function
	FilterElements typr.Qualifier[V]
	// a function that returns nil, unless the given element is not considered valid, otherwise
	// a non-nil error value is returned
	//
	// elements deemed to be invalid are not passed on to the user-defined Ranger function
	ValidateElements typr.Validator[V]
	// when set to true, stops ranging when the first "invalid" element is encountered
	//  - this option is only effective when a viable ValidateElements function has been defined
	HaltOnInvalidElements bool
	// ⚠ Use With Caution! ⚠
	//
	// when set to true:
	//  - the arrayr.Range(...) function (as usual) blocks, until all elements are finished processing
	//  - all eligible elements are submitted to the user-defined Ranger function and processed concurrently
	//  - the return value of the user-defined Ranger function (eg. typr.Break) is ignored
	//  - elements do not necessarily finish processing in order of submission
	//  - ValidateElements function behaves as a FilterElements function, and HaltOnInvalidElements
	//    is ignored, if set
	// ⚠ User-defined Ranger functions should generally take care to operate in a concurrent-safe
	// manner, when using this option ⚠
	Concurrently bool
}

func newRangeOpts[V any]() RangeOpts[V] {
	return RangeOpts[V]{Step: 1}
}

func (o *RangeOpts[V]) assumeValues(uOpts ...RangeOpts[V]) {
	if len(uOpts) == 0 {
		return
	}
	userOpts := uOpts[0]
	if userOpts.Step != 0 {
		o.Step = userOpts.Step
	}
	o.FilterElements = userOpts.FilterElements
	o.ValidateElements = userOpts.ValidateElements
	o.Concurrently = userOpts.Concurrently
}

