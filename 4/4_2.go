package main

import "github.com/apricote/advent-of-code-2023/util"

func SolveCurrentDayWithTwist(input string) int {
	cards := util.Must(ParseInput(input))

	result := 0

	for i, card := range cards {
		result += card.Count
		newCards := WinningNumbers(card)

		for j := i + 1; j < i+1+newCards; j++ {
			cards[j].Count += card.Count
		}
	}

	return result
}
