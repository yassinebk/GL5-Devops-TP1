package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yassinebk/devops1/math"
)

var history []*math.MathOperation

// DisplayHistory prints all previous operations.
func DisplayHistory() {
	if len(history) == 0 {
		fmt.Println("No operations in history.")
		return
	}
	fmt.Println("History of Operations:")
	for i, op := range history {
		op.DisplayOperation(i)
	}
}

// ClearHistory removes all operations from history.
func ClearHistory() {
	history = []*math.MathOperation{}
	fmt.Println("History cleared.")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Perform a new calculation")
		fmt.Println("2. View history")
		fmt.Println("3. Clear history")
		fmt.Println("4. Exit")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			handleCalculation(reader)
		case "2":
			DisplayHistory()
		case "3":
			ClearHistory()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please enter a number between 1 and 4.")
		}
	}
}

// handleCalculation processes user input and performs a new calculation.
func handleCalculation(reader *bufio.Reader) {
	fmt.Print("Enter calculation (e.g., 5 + 3): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Invalid input. Please use the format: operand1 operation operand2")
		return
	}

	operand1, err1 := strconv.Atoi(parts[0])
	operation := parts[1][0]
	operand2, err2 := strconv.Atoi(parts[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid operands. Please enter valid integers.")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	mathOp := math.New(operand1, operand2, operation)
	result := mathOp.DoOperation()
	fmt.Printf("Result: %d %c %d = %d\n", operand1, operation, operand2, result)

	history = append(history, mathOp)
}
