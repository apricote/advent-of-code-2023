#!/usr/bin/env bash
set -e -o pipefail

DAY=$1
mkdir $DAY
pushd $DAY

cat <<EOF > ${DAY}_1.go
package main

func SolveCurrentDay(input string) int {
	return 0
}
EOF

cat <<EOF > ${DAY}_1_test.go
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
		{input: util.GetExampleInput(), want: 0},
		//{input: util.GetInput(), want: 0},
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
EOF

cat <<EOF > ${DAY}_2.go
package main

func SolveCurrentDayWithTwist(input string) int {
	return 0
}
EOF

cat <<EOF > ${DAY}_2_test.go
package main

import (
	"fmt"
	"testing"

	//"github.com/apricote/advent-of-code-2023/util"
)

func TestSolveCurrentDayWithTwist(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		//{input: util.GetExampleInput(), want: 0},
		//{input: util.GetInput(), want: 0},
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
EOF

touch input_example.txt input.txt

git add ./
git commit -m "chore: bootstrap day ${DAY}"
