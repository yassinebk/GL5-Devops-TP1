package math

import (
	"fmt"
	"slices"
)

type MathOperation struct {
	operand1  int
	operand2  int
	operation byte
	result    int
}

var ALLOWED_OPERATIONS = []byte{'+', '-', '/', '*'}

func New(operand1 int, operand2 int, operation byte) *MathOperation {
	if !slices.Contains(ALLOWED_OPERATIONS, operation) {
		panic("Only the following operations are allowed '+','-','/','*'")
	}

	if operand2 == 0 && operation == '/' {
		panic("You can't divide by zero :(")
	}

	return &MathOperation{
		result:    0,
		operand1:  operand1,
		operand2:  operand2,
		operation: operation,
	}
}

func (m *MathOperation) DoOperation() int {
	switch m.operation {
	case '-':
		m.result = m.operand1 - m.operand2
	case '+':
		m.result = m.operand1 + m.operand2
	case '/':
		m.result = m.operand1 / m.operand2
	case '*':
		m.result = m.operand1 * m.operand2
	default:
		panic("The provided operation is not defined ( I don't know how you got here )")
	}

	return m.result
}

func (m *MathOperation) GetResult() int {
	return m.result
}

func (op *MathOperation) DisplayOperation(i int) {
	fmt.Printf("%d. %d %c %d = %d\n", i+1, op.operand1, op.operation, op.operand2, op.result)
}
