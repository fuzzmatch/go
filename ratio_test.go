package hotfuzz

import "testing"

func TestCaseSensitiveSimpleRatioWithLevenshteinDistance(t *testing.T) {
	type testCase struct {
		query  string
		choice string
		score  float32
	}

	var tests = []testCase{
		{"test", "tests", 8.0 / 9.0}, // 1 deletion
		{"test", "tes", 6.0 / 7.0},   // 1 insertion
		{"best", "test", 7.0 / 8.0},  // 1 substitution
		{"Test", "test", 7.0 / 8.0},  // 1 Capital substitution
		{"test", "tets", 6.0 / 8.0},  // transposition counts as 2 edits
		{"test", "test", 1.0},        // should be equal
	}

	for _, test := range tests {
		score := simpleRatio(test.query, test.choice, LevenstheinDistance, true)
		if score != test.score {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected CaseSensitiveSimpleRatio = ", test.score,
				", got CaseSensitiveSimpleRatio = ", score,
			)
		}
	}
}

func TestCaseInsensitiveSimpleRatioWithLevenshteinDistance(t *testing.T) {
	type testCase struct {
		query  string
		choice string
		score  float32
	}

	var tests = []testCase{
		{"test", "tests", 8.0 / 9.0}, // 1 deletion
		{"test", "tes", 6.0 / 7.0},   // 1 insertion
		{"best", "test", 7.0 / 8.0},  // 1 substitution
		{"Test", "test", 1.0},        // 1 Capital substitution should not be counted
		{"test", "tets", 6.0 / 8.0},  // transposition counts as 2 edits
		{"test", "test", 1.0},        // should be equal
	}

	for _, test := range tests {
		score := simpleRatio(test.query, test.choice, LevenstheinDistance, false)
		if score != test.score {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected CaseSensitiveSimpleRatio = ", test.score,
				", got CaseSensitiveSimpleRatio = ", score,
			)
		}
	}
}

func TestCaseSensitivePartialRatioWithLevenshteinDistance(t *testing.T) {
	type testCase struct {
		query  string
		choice string
		score  float32
	}

	var tests = []testCase{
		{"test", "tests", 1.0},      // 1 deletion
		{"test", "tes", 1.0},        // 1 insertion
		{"best", "test", 7.0 / 8.0}, // 1 substitution
		{"Test", "test", 7.0 / 8.0}, // 1 Capital substitution
		{"test", "tets", 6.0 / 8.0}, // transposition counts as 2 edits
		{"test", "test", 1.0},       // should be equal
	}

	for _, test := range tests {
		score := partialRatio(test.query, test.choice, LevenstheinDistance, true)
		if score != test.score {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected CaseSensitiveSimpleRatio = ", test.score,
				", got CaseSensitiveSimpleRatio = ", score,
			)
		}
	}
}

func TestCaseInsensitivePartialRatioWithLevenshteinDistance(t *testing.T) {
	type testCase struct {
		query  string
		choice string
		score  float32
	}

	var tests = []testCase{
		{"test", "tests", 1.0},      // 1 deletion
		{"test", "tes", 1.0},        // 1 insertion
		{"best", "test", 7.0 / 8.0}, // 1 substitution
		{"Test", "test", 1.0},       // 1 Capital substitution
		{"test", "tets", 6.0 / 8.0}, // transposition counts as 2 edits
		{"test", "test", 1.0},       // should be equal
	}

	for _, test := range tests {
		score := partialRatio(test.query, test.choice, LevenstheinDistance, false)
		if score != test.score {
			t.Error(
				"For ", test.query,
				" and ", test.choice,
				" expected CaseSensitiveSimpleRatio = ", test.score,
				", got CaseSensitiveSimpleRatio = ", score,
			)
		}
	}
}