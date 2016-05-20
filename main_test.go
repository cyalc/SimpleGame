package main

import (
	"math/rand"
	"testing"
	"time"
)

var fillNumberToArrayTests = []struct {
	n        int   // input
	expected []int // expected result
}{
	{1234, []int{1, 2, 3, 4}},
	{2345, []int{2, 3, 4, 5}},
	{0, []int{}},
	{000, []int{}},
	{99, []int{9, 9}},
	{34875438, []int{3, 4, 8, 7, 5, 4, 3, 8}},
}

var lengthOfIntegerTable = []struct {
	input  int
	output int
}{
	{0, 0},
	{1, 1},
	{12, 2},
	{453, 3},
	{234234459004, 12},
}

var getGamePointsTable = []struct {
	sourceNumbers   []int
	inputNumber     []int
	outputPlusOnes  int
	outputMinusOnes int
}{
	{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 4, 0},
	{[]int{5, 2, 3, 4}, []int{3, 6, 8, 2}, 0, -2},
	{[]int{5, 3, 6, 1}, []int{1, 3, 6, 2}, 2, -1},
	{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}, 0, -4},
	{[]int{1, 2, 3, 4}, []int{5, 6, 7, 9}, 0, 0},
	{[]int{1, 2, 3, 4}, []int{4, 2, 3, 1}, 2, -2},
	{[]int{1, 2, 3, 4}, []int{1, 2, 4, 9}, 2, -1},
}

func TestGenerateUniqueNumbers(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 12; i++ {
		actual := GenerateUniqueRandomNumbers(i)

		if !isSliceUnique(actual) {
			t.Errorf("GenerateUniqueNumbers (%d) failed, slice is %s", i, actual)
		}
	}
}

func TestFillNumberToArray(t *testing.T) {
	for _, input := range fillNumberToArrayTests {
		actual := FillNumberToArray(input.n)
		if !isEqual(actual, input.expected) {
			t.Errorf("input: %d output : %d  expected : %d", input.n, actual, input.expected)
		}
	}
}

func TestGetLengthOfInteger(t *testing.T) {
	for _, test := range lengthOfIntegerTable {
		actual := getLengthOfInteger(test.input)

		if actual != test.output {
			t.Errorf("input: %d output: %d , we got %d", test.input, test.output, actual)
		}
	}
}

func TestGetGamePoints(t *testing.T) {
	for _, test := range getGamePointsTable {
		pointPlusOne, pointMinusOne := getGamePoints(test.sourceNumbers, test.inputNumber)
		if pointPlusOne != test.outputPlusOnes {
			t.Errorf("PointPlusOne Failed,  input: %d and %d  output: %d expected :%d", test.inputNumber, test.sourceNumbers, pointPlusOne, test.outputPlusOnes)
		}

		if pointMinusOne != test.outputMinusOnes {
			t.Errorf("PointMinusOne Failed,  input: %d and %d  output: %d expected :%d", test.inputNumber, test.sourceNumbers, pointMinusOne, test.outputMinusOnes)
		}
	}
}

func isSliceUnique(s []int) bool {
	for i, var1 := range s {
		for j, var2 := range s {
			if i != j {
				if var1 == var2 {
					return false
				}
			}
		}
	}
	return true
}

func isEqual(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
