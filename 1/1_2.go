package main

import (
	"github.com/apricote/advent-of-code-2023/util"
	"github.com/dlclark/regexp2"
	"strconv"
)

var (
	GetDigitsWithLetters = regexp2.MustCompile(`(?=(\d|one|two|three|four|five|six|seven|eight|nine))`, 0)
)

func SolveCurrentDayWithTwist(input string) int {
	result := 0
	for _, line := range util.SplitLines(input) {
		calibration, err := ProcessLineWithTwist(line)
		if err != nil {
			panic(err)
		}
		result += calibration
	}

	return result
}

func ProcessLineWithTwist(line string) (int, error) {
	digitsAsStrings := regexp2FindAllString(GetDigitsWithLetters, line)

	first, err := ParseDigit(digitsAsStrings[0])
	if err != nil {
		return 0, err
	}
	last, err := ParseDigit(digitsAsStrings[len(digitsAsStrings)-1])
	if err != nil {
		return 0, err
	}

	return first*10 + last, nil
}

func ParseDigit(digitString string) (int, error) {
	switch digitString {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		return strconv.Atoi(digitString)
	}
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.GroupByNumber(1).String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
