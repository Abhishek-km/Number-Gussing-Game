package main

import (
	"fmt"
	"math/rand"
	"errors"
)

func startGame(noOfChances int, randomNumber int) (bool,error) {
	for i := 1; i <= noOfChances; i++ {

		// take the input of the guessed number
		var gussedNumber int
		fmt.Println("")
		fmt.Print("\nEnter your guess:")
		_, err := fmt.Scan(&gussedNumber)
		if err != nil {
			fmt.Println("Error reading the input", err)
			return false, errors.New("Invalid Input")
		}

		if gussedNumber > randomNumber {
			fmt.Printf("Incorrect! The number is less than %v", gussedNumber)
		} else if gussedNumber < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %v", gussedNumber)
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.", i)
			fmt.Println()
			return true,nil
		}
	}
	return false,nil
}

func main() {

	// Print the welcome message of the Game
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	randomNumber := rand.Intn(100) + 1

	// Print the levels of the game
	fmt.Println("")
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	// Take the input of the user regarding the diffuculity level
	var difficultyLevel int
	fmt.Print("\nEnter your choice:")
	_, err := fmt.Scan(&difficultyLevel)
	if err != nil {
		fmt.Println("Error reading the input", err)
		return
	}

	if difficultyLevel < 1 || difficultyLevel > 3 {
		fmt.Println("Wrong choice, Restart the game.")
	}

	// Make variable to store the difficultiy level string value and no. of chances and assign value
	var difficultyLevelString string
	var noOfChances int

	switch difficultyLevel {
	case 1:
		difficultyLevelString = "Easy"
		noOfChances = 10
	case 2:
		difficultyLevelString = "Medium"
		noOfChances = 5
	case 3:
		difficultyLevelString = "Hard"
		noOfChances = 3
	}

	fmt.Printf("\nGreat! You have selected the %s difficulty level.", difficultyLevelString)
	fmt.Printf("\nLet's start the game!")

	for {
		result, error := startGame(noOfChances, randomNumber)
		if result {
			break
		} else if(error == nil) {
			fmt.Printf("\nDo you want to still continue playing?(y/n) ")
			var answer string
			_, err := fmt.Scan(&answer)
			if err != nil {
				fmt.Printf("\nError reading the input", err)
				break
			}
			if answer == "y" {
				continue
			} else if answer == "n" {
				break
			} else {
				fmt.Printf("\nError reading the input,")
				break
			}
		} else{
			break
		}
	}
}
