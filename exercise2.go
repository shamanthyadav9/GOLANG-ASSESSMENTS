package main

import "fmt"

func main() {
	numbers := []int{10, 23, 45, 12, 34, 56, 78}

	sum := 0
	max := numbers[0]
	min := numbers[0]

	for _, num := range numbers {

		fmt.Println(num)

		sum += num

		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
