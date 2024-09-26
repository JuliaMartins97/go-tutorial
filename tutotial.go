package main

import "fmt" // used to output and collect user input

// entry point
func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Print("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You can play!")
	} else {
		fmt.Println("You can`t play :( ")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is the best city in Brazil? ")
	var answer string
	var answer2 string
	var answer3 string
	fmt.Scan(&answer, &answer2, &answer3)

	if answer+" "+answer2+" "+answer3 == "Rio de Janeiro" || answer+" "+answer2+" "+answer3 == "rio de janeiro" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many provinces does Canada have? \n")
	var cores int
	fmt.Scan(&cores)

	if cores == 10 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v. ", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}
