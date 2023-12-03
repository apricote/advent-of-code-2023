package main

import (
	"fmt"
	"github.com/apricote/advent-of-code-2023/util"
	"testing"
	//"github.com/apricote/advent-of-code-2023/util"
)

func TestSolveCurrentDayWithTwist(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 467835},
		{input: util.GetInput(), want: 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := SolveCurrentDayWithTwist(tc.input)

			if tc.want != got {
				t.Errorf("Expected %d but got %d", tc.want, got)
			}
		})
	}
}
