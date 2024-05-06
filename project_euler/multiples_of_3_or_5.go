package project_euler

// MultiplesOf3Or5
/*
*	If we list all the natural numbers below 10 that are multiples of 3 or 5,
*   we get 3,5,6 and 9. The sum of these multiples is 23.
*   Find the sum of all the multiples of 3 or 5 below 10.
 */
func MultiplesOf3Or5(limit int64) int64 {
	var sum int64
	for i := int64(1); i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
