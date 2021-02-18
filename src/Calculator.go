package main

import (
	"calculating"
	"fmt"
	"time"
)

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
			fmt.Println("Result:", calculating.Add(number1, number2, memory))

		case "subtract":
			fmt.Println("Result:", calculating.Subtract(number1, number2, memory))

		case "multiply":
			fmt.Println("Result:", calculating.Multiply(number1, number2, memory))

		case "divide":
			fmt.Println("Result:", calculating.Divide(number1, number2, memory))

		case "remainder":
			fmt.Println("Result:", calculating.DivideRemainder(number1, number2, memory))

		case "bit_left":
			fmt.Println("Result:", calculating.LeftBitMove(number1, number2, memory))

		case "bit_right":
			fmt.Println("Result:", calculating.RightBitMove(number1, number2, memory))

		case "m_add":
			fmt.Println("Result:", calculating.MAdd(number1, memory))

		case "m_subtract":
			fmt.Println("Result:", calculating.MSubtract(number1, memory))

		case "m_multiply":
			fmt.Println("Result:", calculating.MMultiply(number1, memory))

		case "m_divide":
			fmt.Println("Result:", calculating.MDivide(number1, memory))

		case "m_remainder":
			fmt.Println("Result:", calculating.MDivideRemainder(number1, memory))

		case "m_bit_left":
			fmt.Println("Result:", calculating.MLeftBitMove(number1, memory))

		case "m_bit_right":
			fmt.Println("Result:", calculating.MRightBitMove(number1, memory))

		case "reset_memory":
			fmt.Println("Debug memory value:", memory)
			time.Sleep(1000)
			fmt.Println("Memory initializing...")
			InitializingMemory()

		case "help":
			fmt.Println("Available Default Command = [ add, subtract, multiply, divide, remainder, bit_left, bit_right ]")
			fmt.Println("Available Memory Command = [ reset_memory, m_add, m_subtract, m_multiply, m_divide, m_remainder, m_bit_left, m_bit_right ]")
			fmt.Println("Available Util Command = [ help, exit ]")

		case "exit":
			running = false

		default:
			fmt.Println("Calculator> Error. if you want to command list, you have type \"help\" command.")
		}
	}
}

func InitializingMemory() {
	memory = 0
}
