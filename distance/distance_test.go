package distance

import (
	"testing"
)

func TestHamming(t *testing.T) {
	type testCase struct {
		query    string
		choice   string
		distance int
	}

	var tests = []testCase{
		{"test", "test", 0}, // should be equal
		{"test", "tost", 1},
		{"Test", "test", 1}, // should be case sensitive
		{"test", "tets", 2}, // transposition should count as two edits
	}

	for _, test := range tests {
		query_array := []rune(test.query)
		choice_array := []rune(test.choice)
		distance := Hamming(query_array, choice_array)
		if distance != test.distance {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected Hamming = ", test.distance,
				", got Hamming = ", distance,
			)
		}
	}
}

func TestLevenshtein(t *testing.T) {
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
		distance := Levenshtein(query_array, choice_array)
		if distance != test.distance {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected Levenshtein = ", test.distance,
				", got Levenshtein = ", distance,
			)
		}
	}
}

func TestOptimalStringAlignment(t *testing.T) {
	type testCase struct {
		query    string
		choice   string
		distance int
	}

	var tests = []testCase{
		{"ca", "abc", 3},
		{"ca", "ac", 1},
	}

	for _, test := range tests {
		query_array := []rune(test.query)
		choice_array := []rune(test.choice)
		distance := OptimalStringAlignment(query_array, choice_array)
		// distance := edlib.OSADamerauLevenshteinDistance(test.query, test.choice)
		if distance != test.distance {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected OptimalStringAlignment = ", test.distance,
				", got OptimalStringAlignment = ", distance,
			)
		}
	}

}
