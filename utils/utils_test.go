package utils

import (
	"reflect"
	"testing"
)

func TestSwapIntArrays(t *testing.T) {
	arr1 := []int{1}
	arr2 := []int{2}
	arr1_swapped, arr2_swapped := SwapIntArrays(arr1, arr2)

	// Compare arrays using reflect.DeepEqual
	if !reflect.DeepEqual(arr1_swapped, []int{2}) || !reflect.DeepEqual(arr2_swapped, []int{1}) {
		t.Error("Expected arr1_swapped = [2] and arr2_swapped = [1], but got", arr1_swapped, arr2_swapped)
	}
}

func TestSwapRuneArrays(t *testing.T) {
	arr1 := []rune{'a'}
	arr2 := []rune{'b'}
	arr1_swapped, arr2_swapped := SwapRuneArrays(arr1, arr2)

	// Compare arrays using reflect.DeepEqual
	if !reflect.DeepEqual(arr1_swapped, []rune{'b'}) || !reflect.DeepEqual(arr2_swapped, []rune{'a'}) {
		t.Error("Expected arr1_swapped = ['b'] and arr2_swapped = ['a'], but got", arr1_swapped, arr2_swapped)
	}
}

func TestCaseInsensitiveSort(t *testing.T) {
	input_arr := []string{"Apple", "Apple", "Cat", "banana", "cat", "apple", "Banana"}
	expected_sorted := []string{"Apple", "Apple", "apple", "Banana", "banana", "Cat", "cat"}
	sorted := CaseInsensitiveSort(input_arr)

	// Compare arrays
	if !reflect.DeepEqual(expected_sorted, sorted) {
		t.Error(
			"Expected ", expected_sorted,
			" got ", sorted,
		)
	}
}
