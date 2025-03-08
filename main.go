package main

import "fmt"

func main() {
	sumup := sumup(1, 2, 3)
	fmt.Println(sumup)
}

// the parameters are variadic (dynamic)
func sumup(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
