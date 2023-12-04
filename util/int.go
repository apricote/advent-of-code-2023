package util

import (
	"strconv"
)

func ParseInts(intStrings []string) ([]int, error) {
	ints := make([]int, len(intStrings))
	var err error

	for i, intString := range intStrings {
		ints[i], err = strconv.Atoi(intString)
		if err != nil {
			return nil, err
		}
	}

	return ints, nil
}

func Pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
