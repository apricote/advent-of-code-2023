package main

func SolveCurrentDayWithTwist(input string) int {
	games, err := ParseInput(input)
	if err != nil {
		panic(err)
	}

	result := 0

	for _, game := range games {
		result += GamePower(game)
	}

	return result
}

func GamePower(game Game) int {
	minCubes := [3]int{0, 0, 0}

	for _, guess := range game.Guesses {
		for color, amount := range guess {
			if amount > minCubes[color] {
				minCubes[color] = amount
			}
		}
	}

	return minCubes[0] * minCubes[1] * minCubes[2]
}
