package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
	// output:
	// best student in math: aziz with 100
	// best student in science: jamie with 88
	// best student in english: jamie with 90
	// best student in average: aziz with 84

}

func getInfo(students []Student) {
	var bestMathStudent, bestScienceStudent, bestEnglishStudent *Student
	var highestMathScore, highestScienceScore, highestEnglishScore int
	var bestAvgStudent *Student
	var highestAvgScore float64

	for i := 0; i < len(students); i++ {
		student := &students[i]
		if student.MathScore > highestMathScore {
			highestMathScore = student.MathScore
			bestMathStudent = student
		}
		if student.ScienceScore > highestScienceScore {
			highestScienceScore = student.ScienceScore
			bestScienceStudent = student
		}
		if student.EnglishScore > highestEnglishScore {
			highestEnglishScore = student.EnglishScore
			bestEnglishStudent = student
		}
		avgScore := float64(student.MathScore+student.ScienceScore+student.EnglishScore) / 3
		if avgScore > highestAvgScore {
			highestAvgScore = avgScore
			bestAvgStudent = student
		}
	}

	fmt.Printf("best student in math: %s with %d\n", bestMathStudent.Name, highestMathScore)
	fmt.Printf("best student in science: %s with %d\n", bestScienceStudent.Name, highestScienceScore)
	fmt.Printf("best student in english: %s with %d\n", bestEnglishStudent.Name, highestEnglishScore)
	fmt.Printf("best student in average: %s with %.0f\n", bestAvgStudent.Name, highestAvgScore)
}
