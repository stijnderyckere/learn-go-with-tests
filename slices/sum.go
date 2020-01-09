package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		fmt.Printf("index '%d' has value '%v'\n", i, number)
	}
}

//Sum : returns the sum of an array
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

//SumAll : returns the sum for each passed slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

//SumAllTails : returns a slice containing every tail of every slice passed
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return
}
