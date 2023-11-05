package distance

import (
	"github.com/fuzzmatch/go/utils"
)

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