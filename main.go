package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Algorithm :
1 - ) generate a unique n length integer by GenerateUniqueRandomNumbers
2 - ) Wait for a input from user for n length integer
3 - ) Get input and fill it to array
4 - ) compare it and return the point
5 - ) Do this until Point +n


**/
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	computerInteger := GenerateUniqueRandomNumbers(4)

	fmt.Println("Welcome to game")
	
	
	inputInteger := getInput()

	for i := 0; ; {
		if inputInteger != -1 {
			inputIntegerArray := FillNumberToArray(inputInteger)
			plusOne, minusOne := getGamePoints(computerInteger, inputIntegerArray)
			if plusOne == 4 {
				fmt.Println("Congrulations, you won ", i, ". attempt")
				break
			} else {
				fmt.Println("+1 : ", plusOne, "and  -1 : ", minusOne)
				inputInteger = getInput()
			}
		} else {
			fmt.Println("düzgün gir")
			inputInteger = getInput()
		}
	}
}

//GenerateUniqueRandomNumbers Create length of n random unique number
func GenerateUniqueRandomNumbers(n int) []int {
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

//FillNumberToArray convert n digit integer to array
func FillNumberToArray(b int) []int {
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

func getInput() int {

	fmt.Println("Please enter a 4-digit number...")


	var i int
	_, err := fmt.Scan(&i)

	if err != nil {
		fmt.Println("An error has occurred : ", err)
		return -1
	}

	if getLengthOfInteger(i) != 4 {
		fmt.Println("The integer you entered is not 4 digit")
		return -1
	}

	return i
}
