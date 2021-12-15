package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var age uint
	var answer string
	var answer2 string
	score := 0

	fmt.Println("Welcome to the quiz game !!")

	fmt.Print("Enter your name: ")

	fmt.Scan(&name)

	fmt.Printf("Welcome to the game %v ! \n", name)

	fmt.Print("Enter your age: ")

	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Fantastic you can play !!")
	} else {
		fmt.Println("OOPS you can't play !!")
		return
	}

	fmt.Print("Which one is electric, tesla or mercedes? ")

	fmt.Scan(&answer)
	if strings.ToLower(answer) == "tesla" {
		fmt.Println("Correct!!")
		score += 1
	} else {
		fmt.Println("Inorrect!!")
	}

	fmt.Print("Which one is latest, Windows 10 or Windows 11? ")
	fmt.Scan(&answer, &answer2)
	if strings.ToLower(answer)+" "+strings.ToLower(answer2) == "windows 11" {
		fmt.Println("Correct!!")
		score++
	} else {
		fmt.Println("Inorrect!!")
	}

	fmt.Printf("You scored %v out of 2 !!\n", score)
	percent := (float64(score) / float64(2)) * 100

	fmt.Printf("You scored: %v%% !!\n", percent)

}
