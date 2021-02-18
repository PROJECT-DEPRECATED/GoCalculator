package calculating

import "fmt"

func DivideRemainder(number1 int, number2 int, memory int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 % number2

	return number1 % number2
}

func MDivideRemainder(number int, memory int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	return memory % number
}
