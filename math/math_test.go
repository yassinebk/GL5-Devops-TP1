package math

import (
	"testing"
)

func TestNewValidOperations(t *testing.T) {
	tests := []struct {
		operand1  int
		operand2  int
		operation byte
	}{
		{5, 3, '+'},
		{10, 4, '-'},
		{6, 2, '/'},
		{7, 3, '*'},
	}

	for _, tt := range tests {
		mathOp := New(tt.operand1, tt.operand2, tt.operation)
		if mathOp == nil {
			t.Errorf("Expected a valid MathOperation, but got nil for operation %c", tt.operation)
		}
	}
}

func TestNewInvalidOperationPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid operation but did not get one")
		}
	}()
	New(5, 3, '%') // Invalid operation
}

func TestDivideByZeroPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for division by zero but did not get one")
		}
	}()
	New(5, 0, '/')
}

func TestDoOperation(t *testing.T) {
	tests := []struct {
		operand1  int
		operand2  int
		operation byte
		expected  int
	}{
		{5, 3, '+', 8},
		{10, 4, '-', 6},
		{6, 2, '/', 3},
		{7, 3, '*', 21},
	}

	for _, tt := range tests {
		mathOp := New(tt.operand1, tt.operand2, tt.operation)
		result := mathOp.DoOperation()
		if result != tt.expected {
			t.Errorf(
				"Expected %d %c %d = %d, but got %d",
				tt.operand1,
				tt.operation,
				tt.operand2,
				tt.expected,
				result,
			)
		}
	}
}

func TestGetResult(t *testing.T) {
	mathOp := New(5, 3, '+')
	mathOp.DoOperation()
	result := mathOp.GetResult()
	if result != 8 {
		t.Errorf("Expected GetResult() to return 8, but got %d", result)
	}
}

func TestOperationWithoutCallingDoOperation(t *testing.T) {
	mathOp := New(5, 3, '+')
	result := mathOp.GetResult() // Should be 0 since DoOperation() was not called
	if result != 0 {
		t.Errorf("Expected initial result to be 0, but got %d", result)
	}
}

func TestMultipleOperations(t *testing.T) {
	mathOp := New(8, 2, '/')
	mathOp.DoOperation()
	if result := mathOp.GetResult(); result != 4 {
		t.Errorf("Expected 8 / 2 = 4, but got %d", result)
	}

	mathOp.operand1 = 4
	mathOp.operand2 = 3
	mathOp.operation = '*'
	mathOp.DoOperation()
	if result := mathOp.GetResult(); result != 12 {
		t.Errorf("Expected 4 * 3 = 12, but got %d", result)
	}
}
