package arrayr

import (
	"reflect"

	"github.com/eenti-utils/typr"
)

func getPair[X, Y any](e X, ps ...*typr.Pair[X, Y]) (r *typr.Pair[X, Y], b bool) {
	for _, p := range ps {
		if isPairX(e, p) {
			r = p
			b = true
			return
		}
	}
	return
}

func isPairX[X, Y any](e X, p *typr.Pair[X, Y]) (r bool) {
	if p != nil {
		r = reflect.DeepEqual(p.X, e)
	}
	return
}

func convertFreqToX[V any](e typr.Pair[V, int]) (r V) {
	return e.X
}
