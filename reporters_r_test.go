package arrayr

import (
	"testing"

	"github.com/eenti-utils/typr"
)

func TestFrequencyR(t *testing.T) {
	test := FrequencyR("a", "b", "c", "a")

	assrtNotNil(t, test)

	assrtEqual(t, 3, len(test))

	expected := []typr.Pair[string, int]{
		{X: "a", Y: 2},
		{X: "b", Y: 1},
		{X: "c", Y: 1},
	}

	assrtEqual(t, expected, test)

	assrtEqual(t, "a", test[0].X)
	assrtEqual(t, 2, test[0].Y)

	assrtEqual(t, "b", test[1].X)
	assrtEqual(t, 1, test[1].Y)

	assrtEqual(t, "c", test[2].X)
	assrtEqual(t, 1, test[2].Y)

}

func TestFrequencyR2(t *testing.T) {
	test := FrequencyR([]string{"a", "b", "c", "a"}...)

	assrtNotNil(t, test)

	assrtEqual(t, 3, len(test))

	expected := []typr.Pair[string, int]{
		{X: "a", Y: 2},
		{X: "b", Y: 1},
		{X: "c", Y: 1},
	}

	assrtEqual(t, expected, test)

	assrtEqual(t, "a", test[0].X)
	assrtEqual(t, 2, test[0].Y)

	assrtEqual(t, "b", test[1].X)
	assrtEqual(t, 1, test[1].Y)

	assrtEqual(t, "c", test[2].X)
	assrtEqual(t, 1, test[2].Y)

}

func TestFrequencyR3(t *testing.T) {
	nt := new(testing.T)
	nt3 := new(testing.T)
	nt3.Fail() // to distinguish values from nt ...
	test := FrequencyR([]*testing.T{t, nt, nt3, t}...)

	assrtNotNil(t, test)

	assrtEqual(t, 3, len(test))

	expected := []typr.Pair[*testing.T, int]{
		{X: t, Y: 2},
		{X: nt, Y: 1},
		{X: nt3, Y: 1},
	}

	assrtEqual(t, expected, test)

	assrtEqual(t, t, test[0].X)
	assrtEqual(t, 2, test[0].Y)

	assrtEqual(t, nt, test[1].X)
	assrtEqual(t, 1, test[1].Y)

	assrtEqual(t, nt3, test[2].X)
	assrtEqual(t, 1, test[2].Y)

}

func TestFrequencyR_Empty(t *testing.T) {
	test := FrequencyR[int]()
	assrtEqual(t, 0, len(test))
	assrtNotNil(t, test)
}

func TestFrequencyR_Empty2(t *testing.T) {
	test := FrequencyR([]int{}...)
	assrtEqual(t, 0, len(test))
	assrtNotNil(t, test)
}

func TestFrequencyR_Empty3(t *testing.T) {
	test := FrequencyR([]*testing.T{}...)
	assrtEqual(t, 0, len(test))
	assrtNotNil(t, test)
}
