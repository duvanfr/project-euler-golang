package project_euler

// EvenFibonacciNumbers
/**
 * Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
 * 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 * By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
 */
func EvenFibonacciNumbersWithFor(limit int64) int64 {

	sum := int64(0)
	firstNumber, secondNumber := int64(0), int64(1)

	for secondNumber <= limit {
		if secondNumber%2 == 0 {
			sum += secondNumber
		}
		firstNumber, secondNumber = secondNumber, firstNumber+secondNumber
	}
	return sum
}

func EvenFibonacciNumbersWithRecursion(firstNumber, secondNumber, limit int64, counter *int64) int64 {
	if firstNumber <= limit {
		if firstNumber%2 == 0 {
			*counter += firstNumber
		}
		EvenFibonacciNumbersWithRecursion(secondNumber, firstNumber+secondNumber, limit, counter)
	}
	return *counter
}
