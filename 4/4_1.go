package main

import (
	"strings"

	"github.com/apricote/advent-of-code-2023/util"
)

func SolveCurrentDay(input string) int {
	cards := util.Must(ParseInput(input))

	result := 0
	for _, card := range cards {
		result += Points(card)
	}

	return result
}

type Scratchcard struct {
	Count          int
	WinningNumbers []int
	Numbers        []int
}

func ParseInput(input string) ([]Scratchcard, error) {
	var cards []Scratchcard
	var err error

	for _, line := range util.SplitLines(input) {
		var card Scratchcard
		card.Count = 1

		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		segments := strings.Split(line, ": ")
		segments = strings.Split(segments[1], " | ")

		card.WinningNumbers, err = util.ParseInts(strings.Fields(segments[0]))
		if err != nil {
			return nil, err
		}

		card.Numbers, err = util.ParseInts(strings.Fields(segments[1]))
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}

func Points(card Scratchcard) int {
	winningNumbers := WinningNumbers(card)
	if winningNumbers == 0 {
		return 0
	}

	return util.Pow(2, WinningNumbers(card)-1)
}

func WinningNumbers(card Scratchcard) int {
	unique := util.Unique(append(card.WinningNumbers, card.Numbers...))

	return len(card.WinningNumbers) + len(card.Numbers) - len(unique)

}
