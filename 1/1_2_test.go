package main

import (
	"fmt"
	"github.com/apricote/advent-of-code-2023/util"
	"testing"
)

func TestSolveCurrentDayWithTwist(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetInputFromPath("input_example2.txt"), want: 281},
		{input: util.GetInput(), want: 55260},
		{input: util.GetInputFromPath("input_rc.txt"), want: 92},
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
