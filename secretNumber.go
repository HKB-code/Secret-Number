package main

import (
	"fmt"
	"math/rand"
	"time"
)

var bestScore int

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	bestScore = 0
	for {
		playGame()
		if !playAgain() {
			break
		}
	}
	fmt.Println("Thanks for playing! Your best score was:", bestScore)
}

func playGame() {
	welcome()
	difficulty := chooseDifficulty()
	guess := rand.Intn(difficulty) + 1 // Generate a number based on difficulty level
	var userInput int
	attempts := 0

	for {
		userInput = getInfo()
		attempts++
		if userInput < 1 || userInput > difficulty {
			fmt.Printf("Please enter a number between 1 and %d.\n", difficulty)
			continue
		}
		if guess == userInput {
			fmt.Printf("Your guess is right! It took you %d attempts.\n", attempts)
			if bestScore == 0 || attempts < bestScore {
				bestScore = attempts
			}
			break
		} else if guess < userInput {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Too low! Try again.")
		}
	}
}

func welcome() {
	fmt.Println("Welcome to the Number Guessing Game:")
	fmt.Println("Please guess the number from 1 to 10")
}

func chooseDifficulty() int {
	var level int
	fmt.Println("Choose difficulty level:")
	fmt.Println("1. Easy (1-10)")
	fmt.Println("2. Medium (1-50)")
	fmt.Println("3. Hard (1-100)")
	for {
		fmt.Print("Enter your choice (1-3): ")
		_, err := fmt.Scan(&level)
		if err == nil && level >= 1 && level <= 3 {
			break
		}
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
	}
	switch level {
	case 1:
		return 10
	case 2:
		return 50
	case 3:
		return 100
	default:
		return 10
	}
}
func getInfo() int {
	var value int
	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&value)
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Please enter an integer.")
	}
	return value
}

func playAgain() bool {
	var response string
	fmt.Print("Do you want to play again? (y/n): ")
	fmt.Scan(&response)
	return response == "y" || response == "Y"
}
