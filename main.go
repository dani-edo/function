package main

import "fmt"

func main() {
	sumup := sumup([]int{1, 2, 3})
	fmt.Println(sumup)
}

func sumup(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
