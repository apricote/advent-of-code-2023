package main

import (
	"fmt"
	"github.com/apricote/advent-of-code-2023/util"
	"regexp"
	"strconv"
	"strings"
)

var (
	NextNumber = regexp.MustCompile(`\d+`)
)

type Game struct {
	ID      int
	Guesses [][3]int
}

const (
	Red = iota
	Green
	Blue
)

func SolveCurrentDay(input string) int {
	games, err := ParseInput(input)
	if err != nil {
		panic(err)
	}

	result := 0

	totalCubes := [3]int{12, 13, 14}

	for _, game := range games {
		if ValidateGame(game, totalCubes) {
			result += game.ID
		}
	}

	return result
}

func ParseInput(input string) ([]Game, error) {
	var games []Game
	var err error = nil

	for _, line := range util.SplitLines(input) {
		game := Game{}

		game.ID, err = GetNextInt(line)
		if err != nil {
			return nil, err
		}

		line, _ = strings.CutPrefix(line, fmt.Sprintf("Game %d: ", game.ID))

		for _, guess := range strings.Split(line, "; ") {
			cubes := [3]int{0, 0, 0}
			for _, cube := range strings.Split(guess, ", ") {
				var color int
				switch {
				case strings.HasSuffix(cube, "red"):
					color = Red
				case strings.HasSuffix(cube, "green"):
					color = Green
				case strings.HasSuffix(cube, "blue"):
					color = Blue
				}

				cubes[color], err = GetNextInt(cube)
				if err != nil {
					return nil, err
				}
			}

			game.Guesses = append(game.Guesses, cubes)
		}

		games = append(games, game)
	}

	return games, nil
}

func GetNextInt(input string) (int, error) {
	intString := NextNumber.FindString(input)
	return strconv.Atoi(intString)
}

func ValidateGame(game Game, total [3]int) bool {
	for _, guess := range game.Guesses {
		for i := 0; i < 3; i++ {
			if guess[i] > total[i] {
				return false
			}
		}
	}

	return true
}
