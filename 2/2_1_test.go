package main

import (
	"fmt"
	"testing"

	"github.com/apricote/advent-of-code-2023/util"
)

func TestSolveCurrentDay(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 8},
		{input: util.GetInput(), want: 2265},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := SolveCurrentDay(tc.input)

			if tc.want != got {
				t.Errorf("Expected %d but got %d", tc.want, got)
			}
		})
	}
}
