package arrayr

import (
	"fmt"

	"github.com/eenti-utils/typr"
)

func convert[V, R any](ch typr.Resolver[V, R], a ...V) (r []R) {
	for _, e := range a {
		ne := ch(e)
		r = append(r, ne)
	}
	return
}

func doFilter[V any](ok typr.Qualifier[V], a ...V) (r []V) {
	for _, e := range a {
		if ok(e) {
			r = append(r, e)
		}
	}
	return
}

func doFirst[V any](ok typr.Qualifier[V], a ...V) (r V, f bool) {
	if len(a) == 0 {
		return
	}
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

func doFrequency[V comparable](a ...V) (r []typr.Pair[V, int]) {
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

func doLast[V any](ok typr.Qualifier[V], a ...V) (r V, f bool) {
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

func doNth[V any](i int, a ...V) (r V, e error) {
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
