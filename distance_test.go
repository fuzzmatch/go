package hotfuzz

import (
	"reflect"
	"testing"
)

func TestSwapIntArrays(t *testing.T) {
	arr1 := []int{1}
	arr2 := []int{2}
	arr1_swapped, arr2_swapped := swapIntArrays(arr1, arr2)

	// Compare arrays using reflect.DeepEqual
	if !reflect.DeepEqual(arr1_swapped, []int{2}) || !reflect.DeepEqual(arr2_swapped, []int{1}) {
		t.Error("Expected arr1_swapped = [2] and arr2_swapped = [1], but got", arr1_swapped, arr2_swapped)
	}
}

func TestLevenshteinDistance(t *testing.T) {
	type testCase struct {
		query    string
		choice   string
		distance int
	}

	var tests = []testCase{
		{"test", "tests", 1}, // 1 deletion
		{"test", "tes", 1},   // 1 insertion
		{"best", "test", 1},  // 1 substitution
		{"Test", "test", 1},  // 1 Capital substitution
		{"test", "tets", 2},  // transposition counts as 2 edits
		{"test", "test", 0},  // should be equal
	}

	for _, test := range tests {
		query_array := []rune(test.query)
		choice_array := []rune(test.choice)
		distance := LevenstheinDistance(query_array, choice_array)
		if distance != test.distance {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected LevenshteinDistance = ", test.distance,
				", got LevenshteinDistance = ", distance,
			)
		}
	}
}
