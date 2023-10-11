package fuzzy

import "strings"

type distanceFunction func([]rune,[]rune)(int)

func simpleRatio(query string, choice string, scorer distanceFunction,caseSensitive bool) int {
	if !caseSensitive {
		query = strings.ToLower(query)
		choice = strings.ToLower(choice)
	}

	// convert s and t to array
	s := []rune(query)
	t := []rune(choice)

	return (len(s)+len(t)-scorer(s,t))/(len(s)+len(t))
}
