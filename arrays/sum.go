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
