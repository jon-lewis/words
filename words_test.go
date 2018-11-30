package words

import "testing"

func TestStringIsCapitalized(t *testing.T) {
	scenarios := [][]string{
		[]string{"", ""},
		[]string{"1", "1"},
		[]string{"/", "/"},
		[]string{"[", "["},
		[]string{"&", "&"},
		[]string{"pickle", "Pickle"},
		[]string{"two words", "Two words"},
		[]string{"This is capitalized", "This is capitalized"},
	}

	for i := 0; i < len(scenarios); i++ {
		if Capitalize(scenarios[i][0]) != scenarios[i][1] {
			t.Error("Expected", scenarios[i][1], "Actual", Capitalize(scenarios[i][0]))
		}
	}
}
