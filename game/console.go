package game

import "fmt"

//StartGameFromConsole starts the game from console
func StartGameFromConsole() {
	computerInteger := generateUniqueRandomNumbers(4)

	inputInteger := getInput()

	for i := 0; ; {
		if inputInteger != -1 {
			inputIntegerArray := fillNumberToArray(inputInteger)
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
