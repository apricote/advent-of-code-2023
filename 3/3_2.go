package main

import "github.com/apricote/advent-of-code-2023/util"

func SolveCurrentDayWithTwist(input string) int {
	schematic := ParseInput(input)

	var gearMap = make(map[[2]int][]int)

	for i := range schematic {
		currentNumber := 0
		var neighbourGears [][2]int
		for j := range schematic[i] {
			currentRune := schematic[i][j]

			if IsDigit(currentRune) {
				schematic[i][j] = Empty

				currentNumber = currentNumber*10 + int(currentRune-DigitMin)

				gears := GetNeighbourGears(schematic, i, j)
				neighbourGears = append(neighbourGears, gears...)
			}

			if j == len(schematic[i])-1 || !IsDigit(currentRune) {
				for _, gear := range util.Unique(neighbourGears) {
					gearMap[gear] = append(gearMap[gear], currentNumber)
				}

				currentNumber = 0
				neighbourGears = [][2]int{}
			}

		}
	}

	result := 0
	for _, partNumbers := range gearMap {
		if len(partNumbers) != 2 {
			continue
		}
		result += partNumbers[0] * partNumbers[1]
	}
	return result
}

func GetNeighbourGears(schematic [][]rune, i, j int) [][2]int {
	var gears [][2]int

	for _, neighbour := range Neighbours {
		neighbourI := i + neighbour[0]
		neighbourJ := j + neighbour[1]
		if neighbourI < 0 || neighbourI >= len(schematic) || neighbourJ < 0 || neighbourJ >= len(schematic[i]) {
			continue
		}

		value := schematic[neighbourI][neighbourJ]
		if value == Gear {
			gears = append(gears, [2]int{neighbourI, neighbourJ})
		}
	}

	return gears
}
