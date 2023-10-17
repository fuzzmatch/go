package hotfuzz

import (
	"sort"
	"unicode"
	"unicode/utf8"
)

func swapRuneArrays(arr1 []rune, arr2 []rune) ([]rune, []rune) {
	return arr2, arr1
}

func swapIntArrays(arr1 []int, arr2 []int) ([]int, []int) {
	return arr2, arr1
}

func lessCaseInsensitive(s string, t string) bool {
	for {
		if len(t) == 0 {
			return false
		}
		if len(s) == 0 {
			return true
		}
		c, sizec := utf8.DecodeRuneInString(s)
		d, sized := utf8.DecodeRuneInString(t)

		lowerc := unicode.ToLower(c)
		lowerd := unicode.ToLower(d)

		if lowerc < lowerd {
			return true
		}
		if lowerc > lowerd {
			return false
		}
		if c < d {
			return true
		}
		if d < c {
			return false
		}

		s = s[sizec:]
		t = t[sized:]
	}
}

func caseInsensitiveSort(data []string) []string {
	sort.Slice(data, func(i, j int) bool { return lessCaseInsensitive(data[i], data[j]) })
	return data
}
