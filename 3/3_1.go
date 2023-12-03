package main

import (
	"strings"
)

const (
	Empty    = '.'
	DigitMin = '0'
	DigitMax = '9'
	Gear     = '*'
)

var (
	Neighbours = [8][2]int{
		{-1, -1}, // top left
		{-1, 0},  // top
		{-1, 1},  // top right
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // bottom left
		{1, 0},   // bottom
		{1, 1},   // bottom right
	}
)

func SolveCurrentDay(input string) int {
	schematic := ParseInput(input)

	var partNumbers []int

	for i := range schematic {
		currentNumber := 0
		hasNeighbour := false
		for j := range schematic[i] {
			currentRune := schematic[i][j]

			if IsDigit(currentRune) {
				schematic[i][j] = Empty

				currentNumber = currentNumber*10 + int(currentRune-DigitMin)

				if HasNeighbourSymbol(schematic, i, j) {
					hasNeighbour = true
				}
			}

			if j == len(schematic[i])-1 || !IsDigit(currentRune) {
				if hasNeighbour {
					partNumbers = append(partNumbers, currentNumber)
				}
				currentNumber = 0
				hasNeighbour = false
			}

		}
	}

	result := 0
	for _, partNumber := range partNumbers {
		result += partNumber
	}
	return result
}

func ParseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	schematic := make([][]rune, 0, len(lines))

	for i, line := range strings.Split(input, "\n") {
		schematic = append(schematic, make([]rune, len(line)))
		for j, char := range line {
			schematic[i][j] = char
		}
	}
	return schematic
}

func HasNeighbourSymbol(schematic [][]rune, i, j int) bool {
	for _, neighbour := range Neighbours {
		neighbourI := i + neighbour[0]
		neighbourJ := j + neighbour[1]
		if neighbourI < 0 || neighbourI >= len(schematic) || neighbourJ < 0 || neighbourJ >= len(schematic[i]) {
			continue
		}

		value := schematic[neighbourI][neighbourJ]
		if IsSymbol(value) {
			return true
		}
	}

	return false
}

func IsDigit(char rune) bool {
	return char >= DigitMin && char <= DigitMax
}

func IsSymbol(char rune) bool {
	return char != Empty && !IsDigit(char)
}
