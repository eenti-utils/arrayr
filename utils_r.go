package arrayr

import (
	"reflect"

	"github.com/eenti-utils/typr"
)

func getFrequency[V any](e V, fs ...*frequency[V]) (r *frequency[V], b bool) {
	for _, f := range fs {
		if isFrequency(e, f) {
			r = f
			b = true
			return
		}
	}
	return
}

func isFrequency[V any](e V, f *frequency[V]) (r bool) {
	if f != nil {
		r = reflect.DeepEqual(f.X, e)
	}
	return
}

func pairX[V any](p typr.Pair[V, int]) (r V) {
	r = p.X
	return
}

func doFrequencyR[V any](a ...V) (r []typr.Pair[V, int]) {
	var f []*frequency[V]
	for _, el := range a {
		if p, exists := getFrequency(el, f...); exists {
			p.Y += 1
		} else {
			f = append(f, &frequency[V]{X: el, Y: 1})
		}
	}
	if len(f) > 0 {
		cast := func(v *frequency[V]) (r typr.Pair[V, int]) { return typr.Pair[V, int](*v).Value() }
		r = ChangeTo(cast, f...)
	}
	return
}
