package main

func Sum(numbers []int) int {
	totalValue := 0
	for _, number := range numbers {
		totalValue += number
	}
	return totalValue
}

// Go can let you write variadic functions that can take a variable number of arguments.
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
