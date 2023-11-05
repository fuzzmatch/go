// contains all ratios from https://chairnerd.seatgeek.com/fuzzywuzzy-fuzzy-string-matching-in-python/

package ratio

import (
	"strings"

	set "github.com/fuzzmatch/set"

	utils "github.com/fuzzmatch/go/utils"
)

type distanceFunction func([]rune, []rune) int

func SimpleRatio(query string, choice string, metric distanceFunction, caseSensitive bool) float32 {
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

func PartialRatio(query string, choice string, metric distanceFunction, caseSensitive bool) float32 {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// convert s and t to array
	s := []rune(query)
	t := []rune(choice)

	// enforce len(s) < len(t)
	if len(s) > len(t) {
		s, t = utils.SwapRuneArrays(s, t)
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

func TokenSortRatio(query string, choice string, delimiter string, metric distanceFunction, caseSensitive bool) float32 {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// Tokenize query and choice
	query_tokens_slice := strings.Split(query, delimiter)
	query_tokens_slice = utils.CaseInsensitiveSort(query_tokens_slice)

	choice_tokens_slice := strings.Split(choice, delimiter)
	choice_tokens_slice = utils.CaseInsensitiveSort(choice_tokens_slice)

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

func TokenSetRatio(query string, choice string, delimiter string, metric distanceFunction, caseSensitive bool) float32 {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// Tokenize query and choice
	query_tokens_slice := strings.Split(query, delimiter)
	query_tokens_set := set.FromSlice(query_tokens_slice)

	choice_tokens_slice := strings.Split(choice, delimiter)
	choice_tokens_set := set.FromSlice(choice_tokens_slice)

	// Get Intersection
	intersection_tokens_set := query_tokens_set.Intersection(choice_tokens_set)

	// Get Slices again
	query_tokens_slice = utils.CaseInsensitiveSort(query_tokens_set.Elements())
	choice_tokens_slice = utils.CaseInsensitiveSort(choice_tokens_set.Elements())
	intersection_tokens_slice := utils.CaseInsensitiveSort(intersection_tokens_set.Elements())

	// recombine into a single string
	query_tokens_string := strings.Join(query_tokens_slice, " ")
	choice_tokens_string := strings.Join(choice_tokens_slice, " ")
	intersection_tokens_string := strings.Join(intersection_tokens_slice, " ")

	// convert s and t to rune array
	query_tokens_runes := []rune(query_tokens_string)
	choice_tokens_runes := []rune(choice_tokens_string)
	intersection_tokens_runes := []rune(intersection_tokens_string)

	// Compare Different Pairs
	distance := float32(metric(intersection_tokens_runes, query_tokens_runes))
	totalLength := float32(len(intersection_tokens_runes) + len(query_tokens_runes))
	r0 := (totalLength - distance) / totalLength

	distance = float32(metric(intersection_tokens_runes, choice_tokens_runes))
	totalLength = float32(len(intersection_tokens_runes) + len(choice_tokens_runes))
	r1 := (totalLength - distance) / totalLength

	distance = float32(metric(query_tokens_runes, choice_tokens_runes))
	totalLength = float32(len(query_tokens_runes) + len(choice_tokens_runes))
	r2 := (totalLength - distance) / totalLength

	return max(r0, r1, r2)
}
