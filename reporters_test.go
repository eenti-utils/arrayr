package arrayr

import (
	"testing"

	"github.com/eenti-utils/typr"
)

func TestOccurred(t *testing.T) {
	test := Frequency("a", "b", "c", "a")

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

func TestOccurred2(t *testing.T) {
	test := Frequency([]string{"a", "b", "c", "a"}...)

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

func TestOccurred_Empty(t *testing.T) {
	test := Frequency[int]()
	assrtEqual(t, 0, len(test))
	assrtNotNil(t, test)
}

func TestOccurred_Empty2(t *testing.T) {
	test := Frequency([]int{}...)
	assrtEqual(t, 0, len(test))
	assrtNotNil(t, test)
}
