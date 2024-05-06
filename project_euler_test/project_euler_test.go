package project_euler_test

import (
	"project_euler/project_euler"
	"testing"
)

var message string = "Expected %d, current %d"

// TestMultiplesOf3Or5
func TestMultiplesOf3Or5(t *testing.T) {
	if result := project_euler.MultiplesOf3Or5(10); result != 23 {
		t.Errorf(message, 23, result)
	}

	if result := project_euler.MultiplesOf3Or5(100); result != 2318 {
		t.Errorf(message, 2318, result)
	}
}

// EvenFibonacciNumbersWithFor
func TestEvenFibonacciNumbersWithFor(t *testing.T) {
	if result := project_euler.EvenFibonacciNumbersWithFor(1); result != 0 {
		t.Errorf(message, 0, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithFor(2); result != 2 {
		t.Errorf(message, 2, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithFor(3); result != 2 {
		t.Errorf(message, 2, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithFor(8); result != 10 {
		t.Errorf(message, 10, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithFor(34); result != 44 {
		t.Errorf(message, 44, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithFor(59); result != 44 {
		t.Errorf(message, 44, result)
	}
}

// EvenFibonacciNumbersWithRecursion
func TestEvenFibonacciNumbersWithRecursion(t *testing.T) {

	if result := project_euler.EvenFibonacciNumbersWithRecursion(0, 1, 1, new(int64)); result != 0 {
		t.Errorf(message, 0, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithRecursion(1, 2, 2, new(int64)); result != 2 {
		t.Errorf(message, 2, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithRecursion(1, 2, 3, new(int64)); result != 2 {
		t.Errorf(message, 2, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithRecursion(1, 2, 8, new(int64)); result != 10 {
		t.Errorf(message, 10, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithRecursion(1, 2, 34, new(int64)); result != 44 {
		t.Errorf(message, 44, result)
	}
	if result := project_euler.EvenFibonacciNumbersWithRecursion(1, 2, 59, new(int64)); result != 44 {
		t.Errorf(message, 44, result)
	}
}
