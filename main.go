package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	sumupNumbers := sumup(0, numbers...) // ... unpacks the slice
	fmt.Println(sumupNumbers)

	// sumup := sumup(1, 2, 3)
	// fmt.Println(sumup)
}

// the parameters are variadic (dynamic)
func sumup(startingValue int, numbers ...int) int { // ... catch all arguments
	sum := startingValue
	for _, number := range numbers {
		sum += number
	}
	return sum
}
