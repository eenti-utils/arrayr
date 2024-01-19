package arrayr

import (
	"sync"

	"github.com/eenti-utils/typr"
)

// apply the user-defined Ranger function to elements of
// the specified array (or slice) and behave according to 
// the Range Options, if specified
func Range[V any](doRange typr.Ranger[V], a []V, o ...RangeOpts[V]) {
	if doRange == nil || len(a) == 0 {
		return
	}

	opts := newRangeOpts[V]()
	opts.assumeValues(o...)

	if runRange := genRanger(opts); runRange != nil {
		runRange(doRange, a)
	}
}

func genRanger[V any](o RangeOpts[V]) directionalRanger[V] {
	if o.Step > 0 {
		switch o.Concurrently {
		case true:
			return genRangeFC(o)
		default:
			return genRangeF(o)
		}
	}

	if o.Step < 0 {
		switch o.Concurrently {
		case true:
			return genRangeRC(o)
		default:
			return genRangeR(o)
		}
	}

	return nil
}

func genRangeF[V any](o RangeOpts[V]) (r directionalRanger[V]) {
	if o.FilterElements == nil && o.ValidateElements == nil {
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			rangeOp := typr.Continue
			for i := 0; i < l && rangeOp == typr.Continue; i += o.Step {
				rangeOp = rf(i, a[i])
			}
		}
		return
	}

	if o.FilterElements != nil && o.ValidateElements != nil {
		elemOk := o.FilterElements
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			rangeOp := typr.Continue
			for i := 0; i < l && rangeOp == typr.Continue; i += o.Step {
				elem := a[i]
				if err := checkElem(elem); err == nil && elemOk(elem) {
					rangeOp = rf(i, elem)
				} else {
					if o.HaltOnInvalidElements {
						rangeOp = typr.Break
					}
				}
			}
		}
		return
	}

	if o.FilterElements != nil {
		elemOk := o.FilterElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			rangeOp := typr.Continue
			for i := 0; i < l && rangeOp == typr.Continue; i += o.Step {
				if elem := a[i]; elemOk(elem) {
					rangeOp = rf(i, elem)
				}
			}
		}
		return
	} else {
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			rangeOp := typr.Continue
			for i := 0; i < l && rangeOp == typr.Continue; i += o.Step {
				if elem := a[i]; checkElem(elem) == nil {
					rangeOp = rf(i, elem)
				} else if o.HaltOnInvalidElements {
					rangeOp = typr.Break
				}
			}
		}
		return
	}
}

func genRangeR[V any](o RangeOpts[V]) (r directionalRanger[V]) {
	if o.FilterElements == nil && o.ValidateElements == nil {
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1
			rangeOp := typr.Continue
			for i := li; i > -1 && rangeOp == typr.Continue; i += o.Step {
				rangeOp = rf(i, a[i])
			}
		}
		return
	}

	if o.FilterElements != nil && o.ValidateElements != nil {
		elemOk := o.FilterElements
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1
			rangeOp := typr.Continue
			for i := li; i > -1 && rangeOp == typr.Continue; i += o.Step {
				elem := a[i]
				if err := checkElem(elem); err == nil && elemOk(elem) {
					rangeOp = rf(i, elem)
				} else {
					if o.HaltOnInvalidElements {
						rangeOp = typr.Break
					}
				}
			}
		}
		return
	}

	if o.FilterElements != nil {
		elemOk := o.FilterElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1
			rangeOp := typr.Continue
			for i := li; i > -1 && rangeOp == typr.Continue; i += o.Step {
				if elem := a[i]; elemOk(elem) {
					rangeOp = rf(i, elem)
				}
			}
		}
		return
	} else {
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1
			rangeOp := typr.Continue
			for i := li; i > -1 && rangeOp == typr.Continue; i += o.Step {
				if elem := a[i]; checkElem(elem) == nil {
					rangeOp = rf(i, elem)
				} else if o.HaltOnInvalidElements {
					rangeOp = typr.Break
				}
			}
		}
		return
	}
}

func genRangeFC[V any](o RangeOpts[V]) (r directionalRanger[V]) {
	if o.FilterElements == nil && o.ValidateElements == nil {
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)

			var w sync.WaitGroup
			for i := 0; i < l; i += o.Step {
				w.Add(1)

				go func(idx int, elmt V) {
					defer w.Done()
					rf(idx, elmt)
				}(i, a[i])
			}
			w.Wait()
		}
		return
	}

	if o.FilterElements != nil && o.ValidateElements != nil {
		elemOk := o.FilterElements
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)

			var w sync.WaitGroup
			for i := 0; i < l; i += o.Step {
				elem := a[i]
				if err := checkElem(elem); err == nil && elemOk(elem) {
					w.Add(1)
					go func(idx int, elmt V) {
						defer w.Done()
						rf(idx, elmt)
					}(i, elem)
				}
			}
			w.Wait()
		}
		return
	}

	if o.FilterElements != nil {
		elemOk := o.FilterElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)

			var w sync.WaitGroup
			for i := 0; i < l; i += o.Step {
				if elem := a[i]; elemOk(elem) {
					w.Add(1)
					go func(idx int, elmt V) {
						defer w.Done()
						rf(idx, elmt)
					}(i, elem)
				}
			}
			w.Wait()
		}
		return
	} else {
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)

			var w sync.WaitGroup
			for i := 0; i < l; i += o.Step {
				if elem := a[i]; checkElem(elem) == nil {
					w.Add(1)
					go func(idx int, elmt V) {
						defer w.Done()
						rf(idx, elmt)
					}(i, elem)
				}
			}
			w.Wait()
		}
		return
	}
}

func genRangeRC[V any](o RangeOpts[V]) (r directionalRanger[V]) {
	if o.FilterElements == nil && o.ValidateElements == nil {
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1

			var w sync.WaitGroup
			for i := li; i > -1; i += o.Step {
				w.Add(1)
				go func(idx int, elmt V) {
					defer w.Done()
					rf(idx, elmt)
				}(i, a[i])
			}
			w.Wait()
		}
		return
	}
	if o.FilterElements != nil && o.ValidateElements != nil {
		elemOk := o.FilterElements
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1

			var w sync.WaitGroup
			for i := li; i > -1; i += o.Step {
				elem := a[i]
				if err := checkElem(elem); err == nil && elemOk(elem) {
					w.Add(1)
					go func(idx int, elmt V) {
						defer w.Done()
						rf(idx, elmt)
					}(i, elem)
				}
			}
			w.Wait()
		}
		return
	}
	if o.FilterElements != nil {
		elemOk := o.FilterElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1

			var w sync.WaitGroup
			for i := li; i > -1; i += o.Step {
				if elem := a[i]; elemOk(elem) {
					w.Add(1)
					go func(idx int, elmt V) {
						defer w.Done()
						rf(idx, elmt)
					}(i, a[i])
				}
			}
			w.Wait()
		}
		return
	} else {
		checkElem := o.ValidateElements
		r = func(rf typr.Ranger[V], a []V) {
			l := len(a)
			li := l - 1

			var w sync.WaitGroup
			for i := li; i > -1; i += o.Step {
				if elem := a[i]; checkElem(elem) == nil {
					w.Add(1)
					go func(idx int, elmt V) {
						defer w.Done()
						rf(idx, elmt)
					}(i, elem)
				}
			}
			w.Wait()
		}
		return
	}
}
