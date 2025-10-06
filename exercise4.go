package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "orange", "grapes", "mango"}

	//append means adding a new slicew to the existing slice

	fruits = append(fruits, "kiwi") //appending kiwi to the slice

	idx := -1
	for i, fruit := range fruits {
		if fruit == "orange" {
			idx = i
			break
		}
	}
	if idx != -1 {
		fruits = append(fruits[:idx], fruits[idx+1:]...) //removing orange from the slice
	}
	fmt.Println(fruits)

}
