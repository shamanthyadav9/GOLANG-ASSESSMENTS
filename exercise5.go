package main

import "fmt"

func main() { //creating a map of students and their scores
	students := map[string]int{
		"a": 85,
		"b": 78,
		"c": 92,
		"d": 60,
	}
	//print all students and their scores
	fmt.Println("Students and their scores:")
	for name, students := range students {
		fmt.Printf("student %s has score %d\n", name, students)

	}
	//find the student with highest score

	highestname := ""
	highscore := -1 // score(starts with something very low)
	for name, score := range students {
		if score > highscore {
			highscore = score
			highestname = name
		}
	}
	fmt.Printf("Student with highest score: %s (%d)\n", highestname, highscore)

	// 3. Add a new student "E" with score 88
	students["E"] = 88
	fmt.Println("Added E:", students["E"])

	// 4. Check if "C" exists using the 'ok' idiom
	scoreC, ok := students["C"]
	if ok {
		fmt.Printf("Student C exists with score %d\n", scoreC)
	} else {
		fmt.Println("Student C does not exist")
	}
}
