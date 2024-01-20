package arrayr

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test := Filter(pickEvensOnly, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	assrtEqual(t, 5, len(test))
	assrtEqual(t, []int{2, 4, 6, 8, 10}, test)

	test2 := Filter(pickEvensOnly, 1, 3, 5, 7, 9, 11)
	assrtEqual(t, 0, len(test2))
	assrtNotNil(t, test2)
}

func TestFilter2(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test := Filter(pickEvensOnly, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)

	assrtEqual(t, 5, len(test))
	assrtEqual(t, []int{2, 4, 6, 8, 10}, test)

	test2 := Filter(pickEvensOnly, 1, 3, 5, 7, 9, 11)
	assrtEqual(t, 0, len(test2))
	assrtNotNil(t, test2)
}

func TestFirst(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1, ok1 := First(pickEvensOnly, 1, 3, 5, 7, 9, 11, 12, 13)
	assrtEqual(t, 12, test1)
	assrtTrue(t, ok1)

	test2, ok2 := First(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13)
	assrtEqual(t, 0, test2)
	assrtFalse(t, ok2)
}

func TestFirst2(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1, ok1 := First(pickEvensOnly, []int{1, 3, 5, 7, 9, 11, 12, 13}...)
	assrtEqual(t, 12, test1)
	assrtTrue(t, ok1)

	test2, ok2 := First(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13)
	assrtEqual(t, 0, test2)
	assrtFalse(t, ok2)
}

func TestFirstOr(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	defaultValue := -1

	test1 := FirstOr(pickEvensOnly, defaultValue, 1, 3, 5, 7, 9, 11, 12, 13)
	assrtEqual(t, 12, test1)

	test2 := FirstOr(pickEvensOnly, defaultValue, 1, 3, 5, 7, 9, 11, 13)
	assrtEqual(t, defaultValue, test2)
}

func TestFirstOr2(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	defaultValue := -1

	test1 := FirstOr(pickEvensOnly, defaultValue, []int{1, 3, 5, 7, 9, 11, 12, 13}...)
	assrtEqual(t, 12, test1)

	test2 := FirstOr(pickEvensOnly, defaultValue, 1, 3, 5, 7, 9, 11, 13)
	assrtEqual(t, defaultValue, test2)
}

func TestFirstOrZero(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1 := FirstOrZero(pickEvensOnly, 1, 3, 5, 7, 9, 11, 12, 13)
	assrtEqual(t, 12, test1)

	test2 := FirstOrZero(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13)
	assrtEqual(t, 0, test2)
}

func TestFirstOrZero2(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1 := FirstOrZero(pickEvensOnly, []int{1, 3, 5, 7, 9, 11, 12, 13}...)
	assrtEqual(t, 12, test1)

	test2 := FirstOrZero(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13)
	assrtEqual(t, 0, test2)
}

func TestLast(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1, ok1 := Last(pickEvensOnly, 1, 3, 5, 7, 9, 11, 12, 13, 15, 16, 17)
	assrtEqual(t, 16, test1)
	assrtTrue(t, ok1)

	test2, ok2 := Last(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13, 15, 17)
	assrtEqual(t, 0, test2)
	assrtFalse(t, ok2)
}

func TestLast2(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1, ok1 := Last(pickEvensOnly, []int{1, 3, 5, 7, 9, 11, 12, 13, 15, 16, 17}...)
	assrtEqual(t, 16, test1)
	assrtTrue(t, ok1)

	test2, ok2 := Last(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13, 15, 17)
	assrtEqual(t, 0, test2)
	assrtFalse(t, ok2)
}

func TestLastOr(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	defaultValue := -1

	test1 := LastOr(pickEvensOnly, defaultValue, 1, 3, 5, 7, 9, 11, 12, 13, 15, 16, 17)
	assrtEqual(t, 16, test1)

	test2 := LastOr(pickEvensOnly, defaultValue, 1, 3, 5, 7, 9, 11, 13, 15, 17)
	assrtEqual(t, defaultValue, test2)
}

func TestLastOr2(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	defaultValue := -1

	test1 := LastOr(pickEvensOnly, defaultValue, []int{1, 3, 5, 7, 9, 11, 12, 13, 15, 16, 17}...)
	assrtEqual(t, 16, test1)

	test2 := LastOr(pickEvensOnly, defaultValue, 1, 3, 5, 7, 9, 11, 13, 15, 17)
	assrtEqual(t, defaultValue, test2)
}

func TestLastOrZero(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1 := LastOrZero(pickEvensOnly, 1, 3, 5, 7, 9, 11, 12, 13, 15, 16, 17)
	assrtEqual(t, 16, test1)

	test2 := LastOrZero(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13, 15, 17)
	assrtEqual(t, 0, test2)
}

func TestLastOrZero2(t *testing.T) {
	pickEvensOnly := func(e int) bool {
		return e%2 == 0
	}

	test1 := LastOrZero(pickEvensOnly, []int{1, 3, 5, 7, 9, 11, 12, 13, 15, 16, 17}...)
	assrtEqual(t, 16, test1)

	test2 := LastOrZero(pickEvensOnly, 1, 3, 5, 7, 9, 11, 13, 15, 17)
	assrtEqual(t, 0, test2)
}

func TestNth(t *testing.T) {

	test1, err1 := Nth(2, "a", "b", "c", "d")
	assrtEqual(t, "c", test1)
	assrtNil(t, err1)

	test2, err2 := Nth(20, "a", "b", "c", "d")
	assrtEqual(t, "", test2)
	assrtNotNil(t, err2)

	test3, err3 := Nth(-20, "a", "b", "c", "d")
	assrtEqual(t, "", test3)
	assrtNotNil(t, err3)

	test4, err4 := Nth(-2, "a", "b", "c", "d")
	assrtEqual(t, "c", test4)
	assrtNil(t, err4)

}

func TestNth2(t *testing.T) {

	test1, err1 := Nth(2, []string{"a", "b", "c", "d"}...)
	assrtEqual(t, "c", test1)
	assrtNil(t, err1)

	test2, err2 := Nth(20, "a", "b", "c", "d")
	assrtEqual(t, "", test2)
	assrtNotNil(t, err2)
}

func TestNthOr(t *testing.T) {

	defaultValue := "foo"

	test1 := NthOr(1, defaultValue, "a", "b", "c", "d")
	assrtEqual(t, "b", test1)

	test2 := NthOr(20, defaultValue, "a", "b", "c", "d")
	assrtEqual(t, defaultValue, test2)
}

func TestNthOr2(t *testing.T) {

	defaultValue := "foo"

	test1 := NthOr(1, defaultValue, []string{"a", "b", "c", "d"}...)
	assrtEqual(t, "b", test1)

	test2 := NthOr(20, defaultValue, "a", "b", "c", "d")
	assrtEqual(t, defaultValue, test2)
}

func TestNthOrZero(t *testing.T) {

	test1 := NthOrZero(0, "a", "b", "c", "d")
	assrtEqual(t, "a", test1)

	test2 := NthOrZero(-1, "a", "b", "c", "d")
	assrtEqual(t, "d", test2)
}

func TestNthOrZero2(t *testing.T) {

	test1 := NthOrZero(0, []string{"a", "b", "c", "d"}...)
	assrtEqual(t, "a", test1)

	test2 := NthOrZero(-2, "a", "b", "c", "d")
	assrtEqual(t, "c", test2)

	test3 := NthOrZero(-5, "a", "b", "c", "d")
	assrtEqual(t, "", test3)
}

func TestRepeated(t *testing.T) {
	test1 := Repeated(1, 2, 3, 3, 4, 5)

	assrtEqual(t, []int{3}, test1)

	test2 := Repeated(1, 2, 3, 4, 5)
	assrtEqual(t, 0, len(test2))
}

func TestRepeated2(t *testing.T) {
	test1 := Repeated([]int{1, 2, 3, 3, 4, 5}...)

	assrtEqual(t, []int{3}, test1)

	test2 := Repeated(1, 2, 3, 4, 5)
	assrtEqual(t, 0, len(test2))
}

func TestUnique(t *testing.T) {
	test1 := Unique(1, 2, 2, 2, 2, 2, 2, 3)

	assrtEqual(t, []int{1, 2, 3}, test1)
}

func TestUnique2(t *testing.T) {
	test1 := Unique([]int{1, 2, 2, 2, 2, 2, 2, 3}...)

	assrtEqual(t, []int{1, 2, 3}, test1)
}

func TestUnrepeated(t *testing.T) {
	test1 := Unrepeated(1, 2, 3, 3, 4, 5)

	assrtEqual(t, []int{1, 2, 4, 5}, test1)

	test2 := Unrepeated(1, 2, 3, 4, 5)
	assrtEqual(t, 5, len(test2))
}

func TestUnrepeated2(t *testing.T) {
	test1 := Unrepeated([]int{1, 2, 3, 3, 4, 5}...)

	assrtEqual(t, []int{1, 2, 4, 5}, test1)

	test2 := Unrepeated(1, 2, 3, 4, 5)
	assrtEqual(t, 5, len(test2))
}

func TestValid(t *testing.T) {
	checkValid := func(e string) (r error) {
		if len(e)%3 != 0 {
			r = fmt.Errorf("invalid element [ %s ]", e)
		}
		return
	}
	test := Valid(checkValid, "The", "Quick", "Brown", "Fox", "Jumped", "Over", "The", "Lazy", "Dog")

	assrtEqual(t, []string{"The", "Fox", "Jumped", "The", "Dog"}, test)
}

func TestValid2(t *testing.T) {
	checkValid := func(e string) (r error) {
		if len(e)%3 != 0 {
			r = fmt.Errorf("invalid element [ %s ]", e)
		}
		return
	}
	test := Valid(checkValid, []string{"The", "Quick", "Brown", "Fox", "Jumped", "Over", "The", "Lazy", "Dog"}...)

	assrtEqual(t, []string{"The", "Fox", "Jumped", "The", "Dog"}, test)
}
