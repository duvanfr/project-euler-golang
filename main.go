package main

import (
	"fmt"
	"project_euler/project_euler"
)

func main() {
	fmt.Printf("Exercises 1:\nThe sum of all the multiples of 3 or 5 below 100 is: %v\n", project_euler.MultiplesOf3Or5(1000))
	fmt.Printf("Exercises 2:\nThe sum of the even-valued terms using for cycle is: %v\n", project_euler.EvenFibonacciNumbersWithFor(4000000))
	fmt.Printf("Exercises 2:\nThe sum of the even-valued terms using recursion is: %v\n", project_euler.EvenFibonacciNumbersWithRecursion(1, 2, 4000000, new(int64)))
}
