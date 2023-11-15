package distance

import (
	"github.com/fuzzmatch/go/utils"
)

func Hamming(s []rune, t []rune) int {

	if len(s) != len(t) {
		panic("rune arrays must be the same length when using the Hamming Distance")
	}

	distance := 0
	for i := range s {
		if s[i] != t[i] {
			distance++
		}
	}
	return distance
}

func Levenshtein(s []rune, t []rune) int {
	// extract lengths
	m := len(s)
	n := len(t)

	// Initialise worker vectors
	v0 := make([]int, n+1)
	v1 := make([]int, n+1)

	for i := 0; i <= n; i++ {
		v0[i] = i
	}

	for i := 0; i < m; i++ {
		v1[0] = i + 1
		for j := 0; j < n; j++ {
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1
			substitutionCost := 0
			if s[i] == t[j] {
				substitutionCost = v0[j]
			} else {
				substitutionCost = v0[j] + 1
			}
			v1[j+1] = min(deletionCost, insertionCost, substitutionCost)
		}
		v0, v1 = utils.SwapIntArrays(v0, v1)
	}
	return v0[n]
}

func OptimalStringAlignment(s []rune, t []rune) int {
	matrix := utils.Make2DArray(len(s)+1, len(t)+1)

	for i := 0; i <= len(s); i++ {
		matrix[i][0] = i
	}

	for j := 0; j <= len(t); j++ {
		matrix[0][j] = j
	}

	var count int
	for i := 1; i <= len(s); i++ {
		matrix[i][0] = i
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				count = 0
			} else {
				count = 1
			}

			matrix[i][j] = min(min(matrix[i-1][j]+1, matrix[i][j-1]+1), matrix[i-1][j-1]+count) // insertion, deletion, substitution
			if i > 1 && j > 1 && s[i-1] == t[j-2] && s[i-2] == t[j-1] {
				matrix[i][j] = min(matrix[i][j], matrix[i-2][j-2]+1) // translation
			}
		}
	}
	return matrix[len(s)][len(t)]
}
