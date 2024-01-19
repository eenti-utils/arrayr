package arrayr

import "testing"

func TestDedupe(t *testing.T) {
	test1 := Dedupe(1, 2, 3, 3, 4, 5)

	assrtEqual(t, []int{1, 2, 3, 4, 5}, test1)

	test2 := Dedupe(1, 2, 3, 4, 5)
	assrtEqual(t, 5, len(test2))

}

func TestDedupe2(t *testing.T) {
	test1 := Dedupe([]int{1, 2, 3, 3, 4, 5}...)

	assrtEqual(t, []int{1, 2, 3, 4, 5}, test1)

	test2 := Dedupe(1, 2, 3, 4, 5)
	assrtEqual(t, 5, len(test2))

}

func TestModify(t *testing.T) {
	intToFloat := func(i int) (r float64) {
		r = float64(i)
		return
	}

	test := Modify(intToFloat, 1, 2, 3, 4, 5)

	assrtEqual(t, []float64{1.0, 2.0, 3.0, 4.0, 5.0}, test)
}

func TestModify2(t *testing.T) {
	intToFloat := func(i int) (r float64) {
		r = float64(i)
		return
	}

	test := Modify(intToFloat, []int{1, 2, 3, 4, 5}...)

	assrtEqual(t, []float64{1.0, 2.0, 3.0, 4.0, 5.0}, test)
}
