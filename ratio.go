// contains all ratios from https://chairnerd.seatgeek.com/fuzzywuzzy-fuzzy-string-matching-in-python/

package hotfuzz

import "strings"

type distanceFunction func([]rune, []rune) int

func swapRuneArrays(arr1 []rune, arr2 []rune) ([]rune, []rune) {
	return arr2, arr1
}

func simpleRatio(query string, choice string, scorer distanceFunction, caseSensitive bool) float32 {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// convert s and t to array
	s := []rune(query)
	t := []rune(choice)

	return (float32(len(s) + len(t) - scorer(s, t))) / float32(len(s)+len(t))
}

func partialRatio(query string, choice string, scorer distanceFunction, caseSensitive bool) float32 {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// convert s and t to array
	s := []rune(query)
	t := []rune(choice)

	// enforce len(s) < len(t)
	if len(s) > len(t) {
		s, t = swapRuneArrays(s, t)
	}

	m := len(s)
	n := len(t)
	lenDif := n - m
	minDistance := 2 * m

	for i := 0; i <= lenDif; i++ {
		sub := t[i : i+m]
		distance := scorer(s, sub)
		if distance < minDistance {
			minDistance = distance
		}
	}
	return float32(2*m-minDistance) / float32(2*m)
}
