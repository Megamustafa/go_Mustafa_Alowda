package main

import "fmt"

func main() {
	//Priority 1, Task 1
	var radius float32
	fmt.Print("Please enter the raidus of the tube: ")
	fmt.Scan(&radius)

	var height float32
	fmt.Print("Please enter the height of the tube: ")
	fmt.Scan(&height)

	var area float32 = 3.14 * radius * radius * height

	fmt.Println("The area of the tube is: ", area)

	//Priority 1, Task 2
	var score int
	fmt.Print("Please enter your score: ")
	fmt.Scan(&score)

	if score <= 100 && score >= 85 {
		fmt.Println("Grade A")
	} else if score < 85 && score >= 70 {
		fmt.Println("Grade B")
	} else if score < 70 && score >= 55 {
		fmt.Println("Grade C")
	} else if score <= 50 && score >= 40 {
		fmt.Println("Grade D")
	} else if score < 40 && score >= 0 {
		fmt.Println("Grade E")
	} else if score < 0 || score > 00 {
		fmt.Println("Invalid score")
	}

	//Priority 1, Task 3
	for i := 1; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i * i)
		} else if i%7 == 0 {
			fmt.Println(i * i * i)
		} else if i%4 == 0 && i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}

	//Priority 2, Task 1
	var nu int
	fmt.Print("please enter a number: ")
	fmt.Scan(&nu)
	for i := 1; i <= nu; i++ {
		if i%2 == 0 && nu%i == 0 {
			fmt.Print("I")
		} else if i%2 != 0 && nu%i == 0 {
			fmt.Print("O")
		}

	}

	//Priorty 2, Task 2
	var budget int
	var budgetscore int
	var duration int
	var durationscore int
	var diffculty int
	var diffcultyscore int
	var totalscore int

	fmt.Println("Please insert budget: ")
	fmt.Scan(&budget)

	if budget <= 50 && budget >= 0 {
		budgetscore = 4
	} else if budget <= 80 && budget >= 51 {
		budgetscore = 3
	} else if budget <= 100 && budget >= 81 {
		budgetscore = 2
	} else if budget > 100 {
		budgetscore = 1
	}

	fmt.Println("Please insert duration: ")
	fmt.Scan(&duration)

	if duration <= 20 && duration >= 0 {
		durationscore = 10
	} else if duration <= 30 && duration >= 21 {
		durationscore = 5
	} else if duration <= 50 && duration >= 31 {
		durationscore = 2
	} else if duration > 50 {
		durationscore = 1
	}

	fmt.Println("Please insert difficulty: ")
	fmt.Scan(&diffculty)

	if diffculty <= 3 && diffculty >= 0 {
		diffcultyscore = 10
	} else if diffculty <= 6 && diffculty >= 4 {
		diffcultyscore = 5
	} else if diffculty <= 10 && diffculty >= 8 {
		diffcultyscore = 1
	} else if diffculty > 10 {
		diffcultyscore = 0
	}

	totalscore = budgetscore + durationscore + diffcultyscore

	fmt.Println("total score is: ", totalscore)

	if totalscore <= 24 && totalscore >= 17 {
		fmt.Println("Output: High")
	} else if totalscore <= 16 && totalscore >= 10 {
		fmt.Println("Output: Medium")
	} else if totalscore <= 9 && totalscore >= 3 {
		fmt.Println("Output: Low")
	} else if totalscore <= 2 {
		fmt.Println("Output: Impossible")
	}
}
