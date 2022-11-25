package gen

import (
	"fmt"
	"sort"
	"strings"
)

func array(elems ...string) []string {
	return append([]string{}, elems...)
}

func appendFn(arr []string, elems ...string) []string {
	return append(arr, elems...)
}

func first(elems []string) (string, error) {
	if len(elems) == 0 {
		return "", fmt.Errorf("first: no elements in string array")
	}
	return elems[0], nil
}

func last(elems []string) (string, error) {
	if len(elems) == 0 {
		return "", fmt.Errorf("last: no elements in string array")
	}
	return elems[len(elems)-1], nil
}

func sortFn(array []string) []string {
	sorted := make([]string, len(array))
	copy(sorted, array)
	sort.Strings(sorted)
	return sorted
}

func split(sep string, str string) []string {
	return strings.Split(str, sep)
}
