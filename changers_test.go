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

func TestChangeTo(t *testing.T) {
	intToFloat := func(i int) (r float64) {
		r = float64(i)
		return
	}

	test := ChangeTo(intToFloat, 1, 2, 3, 4, 5)

	assrtEqual(t, []float64{1.0, 2.0, 3.0, 4.0, 5.0}, test)
}

func TestChangeTo2(t *testing.T) {
	intToFloat := func(i int) (r float64) {
		r = float64(i)
		return
	}

	test := ChangeTo(intToFloat, []int{1, 2, 3, 4, 5}...)

	assrtEqual(t, []float64{1.0, 2.0, 3.0, 4.0, 5.0}, test)
}

func TestReverse(t *testing.T) {
	reverse := Reverse("apple","banana","coconut")
	
	assrtEqual(t,3,len(reverse))
	assrtEqual(t,[]string{"coconut","banana","apple"}, reverse)
}

func TestReverse2(t *testing.T) {
	testArr := []int{0,1,22,2,-5,3,4}
	reverse := Reverse(testArr...)

	assrtEqual(t,7,len(testArr))
	assrtEqual(t,7,len(reverse))
	assrtEqual(t,[]int{4, 3, -5, 2, 22, 1, 0}, reverse)
}