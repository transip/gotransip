package util

import (
	"strconv"
)

func Atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// Contains is a case insenstive match, finding needle in a haystack
func Contains(haystack []int, needle int) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}
