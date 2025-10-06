package main

import "fmt"

func main() {
	//Create an array of 10 integers.

	arr := [10]int{1, 2, 3, 5, 7, 8, 1, 2, 4, 19}
	fmt.Println(arr)

	//Fill it with user input values.//From class 12
	var arr1 [5]int
	fmt.Println("Enter 5 integers:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr1[i]) //here this raeds the input from the user and stops when the user hits enter
	}

	fmt.Println("Your values using scan:", arr1)

	var arr2 [5]int
	fmt.Println("Enter 5 integers:")
	for i := 0; i < 5; i++ {
		fmt.Scanln(&arr2[i]) //here this reads the input from the user one line at a time and stops when the user hits enter
	}
	fmt.Println("numbers(using scanln):", arr2)

	//Write functions to:
	//Find the average of the numbers

	arr3 := [5]int{10, 20, 30, 40, 50}
	sum := 0                     //I initialized the sum variable to store the sum of array elements
	for _, value := range arr3 { // syntax for range loop
		sum += value
	}
	fmt.Println("Sum of array elements:", sum)
	average := float64(sum) / float64(len(arr3)) //here we are using only float64 instead 34 is becx this is to convert the integer to float in precise values.
	{
		fmt.Println("Average of the elements :", average)
	}

	//Count how many numbers are greater than the average

	arr4 := [5]int{10, 20, 30, 40, 50} // This stores the elements in arr4
	count := 0
	for _, value := range arr4 { //keeping index as blank
		if value > int(average) { //here we are converting the average i,e avg is 30
			count++ //this increments the count value by 1 for every value which is greater than avg

		}
	}
	fmt.Println("Count of the elements greater than average:", count)
}
