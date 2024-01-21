package arrayr

import (
	"testing"
)

func TestRepeatedR(t *testing.T) {
	test1 := RepeatedR(1, 2, 3, 3, 4, 5)

	assrtEqual(t, []int{3}, test1)

	test2 := RepeatedR(1, 2, 3, 4, 5)
	assrtEqual(t, 0, len(test2))
}

func TestRepeatedR2(t *testing.T) {
	test1 := RepeatedR([]int{1, 2, 3, 3, 4, 5}...)

	assrtEqual(t, []int{3}, test1)

	test2 := RepeatedR(1, 2, 3, 4, 5)
	assrtEqual(t, 0, len(test2))
}

func TestRepeatedR3(t *testing.T) {
	nt := new(testing.T)
	test1 := RepeatedR(t, t, t, nt)

	assrtEqual(t, []*testing.T{t}, test1)

	test2 := RepeatedR(t, t, nt, nt)
	assrtEqual(t, 2, len(test2))
}

func TestUniqueR(t *testing.T) {
	test1 := UniqueR(1, 2, 2, 2, 2, 2, 2, 3)

	assrtEqual(t, []int{1, 2, 3}, test1)
}

func TestUniqueR2(t *testing.T) {
	test1 := UniqueR([]int{1, 2, 2, 2, 2, 2, 2, 3}...)

	assrtEqual(t, []int{1, 2, 3}, test1)
}

func TestUniqueR3(t *testing.T) {
	nt := new(testing.T)
	nt3 := new(testing.T)
	nt3.Fail() // to distinguish values from nt ...

	test1 := UniqueR(t, nt, nt, nt, nt, nt, nt, nt3)

	assrtEqual(t, []*testing.T{t, nt, nt3}, test1)
}

func TestUnrepeatedR(t *testing.T) {
	test1 := UnrepeatedR(1, 2, 3, 3, 4, 5)

	assrtEqual(t, []int{1, 2, 4, 5}, test1)

	test2 := UnrepeatedR(1, 2, 3, 4, 5)
	assrtEqual(t, 5, len(test2))
}

func TestUnrepeatedR2(t *testing.T) {
	test1 := UnrepeatedR([]int{1, 2, 3, 3, 4, 5}...)

	assrtEqual(t, []int{1, 2, 4, 5}, test1)

	test2 := UnrepeatedR(1, 2, 3, 4, 5)
	assrtEqual(t, 5, len(test2))
}

func TestUnrepeatedR3(t *testing.T) {
	nt := new(testing.T)
	nt3 := new(testing.T)
	nt3.Fail() // to distinguish values from nt ...

	test1 := UnrepeatedR(t, nt, nt3, nt3, nil)

	assrtEqual(t, []*testing.T{t, nt, nil}, test1)

	test2 := UnrepeatedR(t, nt, nt3, nil)
	assrtEqual(t, 4, len(test2))
}
