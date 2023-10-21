// contains all ratios from https://chairnerd.seatgeek.com/fuzzywuzzy-fuzzy-string-matching-in-python/

package fuzzmatch

import (
	"strings"
)

type distanceFunction func([]rune, []rune) int

func simpleRatio(query string, choice string, metric distanceFunction, caseSensitive bool) float32 {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// convert s and t to array
	s := []rune(query)
	t := []rune(choice)

	totalLength := float32(len(s) + len(t))
	distance := float32(metric(s, t))
	return (totalLength - distance) / totalLength
}

func partialRatio(query string, choice string, metric distanceFunction, caseSensitive bool) float32 {
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
		distance := metric(s, sub)
		if distance < minDistance {
			minDistance = distance
		}
	}
	return float32(2*m-minDistance) / float32(2*m)
}

func tokenSortRatio(query string, choice string, delimeter string, metric distanceFunction, caseSensitive bool) float32 {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// Tokenize query and choice
	query_tokens_slice := strings.Split(query, delimeter)
	query_tokens_slice = caseInsensitiveSort(query_tokens_slice)

	choice_tokens_slice := strings.Split(choice, delimeter)
	choice_tokens_slice = caseInsensitiveSort(choice_tokens_slice)

	// recombine into a single string
	query_tokens := strings.Join(query_tokens_slice, " ")
	choice_tokens := strings.Join(choice_tokens_slice, " ")

	// convert s and t to array
	s := []rune(query_tokens)
	t := []rune(choice_tokens)

	totalLength := float32(len(s) + len(t))
	distance := float32(metric(s, t))
	return (totalLength - distance) / totalLength
}

// Token Set
