package main

import "fmt"

func main() {
	s := []int{12, 34, 56, 78, 90, 11}
	val := 78
	found := false

	for i, v := range s {
		if v == val {
			fmt.Printf("Value %d found at index %d\n", val, i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("NotFound")
	}
} // <- this is the closing brace for main
