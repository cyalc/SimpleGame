package game

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func generateUniqueRandomNumbers(n int) []int {
	integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n = n % 9
	s := make([]int, n)

	for i := 0; i < n; i++ {
		index := rand.Intn(len(integers))
		s[i] = integers[index]
		integers = append(integers[:index], integers[index+1:]...)
	}
	return s
}

func fillNumberToArray(b int) []int {
	length := getLengthOfInteger(b) //2
	s := make([]int, length)

	cofactor := 1
	for i := 0; i < length-1; i++ {
		cofactor *= 10 // 10 100
	}

	for i := 0; i < length; i++ {
		s[i] = b / cofactor       // 4321 / 1000 = 4
		b = b - (s[i] * cofactor) //
		cofactor = cofactor / 10
	}

	return s
}

func getLengthOfInteger(n int) int {
	i := n
	length := 0
	for i >= 1 {
		i = i / 10
		length++
	}
	return length
}

func getGamePoints(sourceNumbers []int, inputNumbers []int) (int, int) {
	pointPlusOne := 0
	pointMinusOne := 0

	for i := 0; i < len(sourceNumbers); i++ {
		if sourceNumbers[i] == inputNumbers[i] {
			pointPlusOne++
		}
		for j := 0; j < len(inputNumbers); j++ {
			if i != j {
				if sourceNumbers[i] == inputNumbers[j] {
					pointMinusOne--
				}
			}
		}
	}
	return pointPlusOne, pointMinusOne
}
