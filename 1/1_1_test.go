package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2023/util"
)

func TestSolveCurrentDay(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 142},
		{input: util.GetInput(), want: 55123},
	}

	for _, tc := range tests {
		got := SolveCurrentDay(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
