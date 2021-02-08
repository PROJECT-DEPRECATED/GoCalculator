package main

import "fmt"

var memory int

func main() {
	var running = true

	for running {
		var number1, number2 int
		var calculatingType string

		fmt.Println("Calculator> Please type commands.")
		fmt.Scan(&calculatingType)

		switch calculatingType {
		case "add":
			fmt.Println("Result:", Add(number1, number2))

		case "subtract":
			fmt.Println("Result:", Subtract(number1, number2))

		case "multiply":
			fmt.Println("Result:", Multiply(number1, number2))

		case "divide":
			fmt.Println("Result:", Divide(number1, number2))

		case "remainder":
			fmt.Println("Result:", DivideRemainder(number1, number2))

		case "bit_left":
			fmt.Println("Result:", LeftBitMove(number1, number2))

		case "bit_right":
			fmt.Println("Result:", RightBitMove(number1, number2))

		case "m_add":
			fmt.Println("Result:", MAdd(number1))

		case "m_subtract":
			fmt.Println("Result:", MSubtract(number1))

		case "m_multiply":
			fmt.Println("Result:", MMultiply(number1))

		case "m_divide":
			fmt.Println("Result:", MDivide(number1))

		case "m_remainder":
			fmt.Println("Result:", MDivideRemainder(number1))

		case "m_bit_left":
			fmt.Println("Result:", MLeftBitMove(number1))

		case "m_bit_right":
			fmt.Println("Result:", MRightBitMove(number1))

		case "help":
			fmt.Println("Available Default Command = [ add, subtract, multiply, divide, remainder, bit_left, bit_right ]")
			fmt.Println("Available Memory Command = [ m_add, m_subtract, m_multiply, m_divide, m_remainder, m_bit_left, m_bit_right ]")
			fmt.Println("Available Util Command = [ help, exit ]")

		case "exit":
			running = false

		default:
			fmt.Println("Calculator> Error. if you want to command list, you have type \"help\" command.")
		}
	}
}

func MAdd(number int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	memory = memory + number

	return memory + number
}

func MSubtract(number int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	memory = memory - number

	return memory - number
}

func MMultiply(number int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	memory = memory * number

	return memory * number
}

func MDivide(number int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	memory = memory / number

	return memory / number
}

func MDivideRemainder(number int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	return memory % number
}

func MLeftBitMove(number int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	memory = memory << number

	return memory << number
}

func MRightBitMove(number int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number)

	memory = memory >> number

	return memory >> number
}

func Add(number1 int, number2 int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 + number2

	return number1 + number2
}

func Subtract(number1 int, number2 int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 - number2

	return number1 - number2
}

func Multiply(number1 int, number2 int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 * number2

	return number1 * number2
}

func Divide(number1 int, number2 int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 / number2

	return number1 / number2
}

func DivideRemainder(number1 int, number2 int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 % number2

	return number1 % number2
}

func LeftBitMove(number1 int, number2 int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 << number2

	return number1 << number2
}

func RightBitMove(number1 int, number2 int) int {
	fmt.Println("Calculator> Please type first number.")
	fmt.Scan(&number1)

	fmt.Println("Calculator> Please type second number.")
	fmt.Scan(&number2)

	memory = number1 >> number2

	return number1 >> number2
}

func InitializingMemory() {
	memory = 0
}