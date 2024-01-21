package arrayr

import "github.com/eenti-utils/typr"

type frequency[V any] typr.Pair[V,int]

func (f frequency[V]) Value() (r frequency[V]) {
	r = frequency[V]{X: f.X, Y: f.Y}
	return
}