package arrayr

import "testing"

func TestFrom(t *testing.T) {
	test := From("a", "b", "c")

	assrtEqual(t, []string{"a", "b", "c"}, test)
}

func TestFrom_Empty(t *testing.T) {
	test := From[any]()

	assrtEqual(t, 0, len(test))
	assrtNotNil(t, test)
}

func TestFrom_Any(t *testing.T) {
	test := From[any](1, "a", t)
	assrtEqual(t, []any{1, "a", t}, test)
}

func TestFromMap(t *testing.T) {
	test := FromMap(map[string]string{
		"do": "a deer, a female deer",
		"re": "a drop of golden sun",
		"mi": "a name I call myself",
	})

	assrtEqual(t, 3, len(test))

	for _, pair := range test {
		switch pair.X {
		case "do":
		assrtEqual(t, "a deer, a female deer", pair.Y)

		case "re":
		assrtEqual(t, "a drop of golden sun", pair.Y)

		case "mi":
		assrtEqual(t, "a name I call myself", pair.Y)
		
		default:
			t.Fatalf(`unexpected values -- pair{X: "%s", Y: "%s"}`,pair.X,pair.Y)
		}
	}
}
