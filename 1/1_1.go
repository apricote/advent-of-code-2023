package main

import (
	"github.com/apricote/advent-of-code-2023/util"
	"regexp"
	"strconv"
)

var (
	GetDigits = regexp.MustCompile(`\d`)
)

func SolveCurrentDay(input string) int {
	result := 0
	for _, line := range util.SplitLines(input) {
		calibration, err := ProcessLine(line)
		if err != nil {
			panic(err)
		}
		result += calibration
	}

	return result
}

func ProcessLine(line string) (int, error) {
	digitsAsStrings := GetDigits.FindAllString(line, -1)

	first, err := strconv.Atoi(digitsAsStrings[0])
	if err != nil {
		return 0, err
	}
	last, err := strconv.Atoi(digitsAsStrings[len(digitsAsStrings)-1])
	if err != nil {
		return 0, err
	}

	return first*10 + last, nil
}
