package june

import (
	"bufio" // Importing bufio for buffered I/O operations
	"fmt"
	"math/rand" // Importing math/rand for random number generation
	"os"        // Importing os for standard input/output operations
	"strconv"
	"strings"
	"time" // Importing time for seeding the random number generator and handling time-related functions
)

// NumberGuessingGame implements a simple number guessing game where the user has to guess a randomly generated number between 1 and 100.
func NumberGuessingGame() {
	rand.Seed(time.Now().UnixNano())    // Seed the random number generator with the current time
	target := rand.Intn(100) + 1        // Generate a random number between 1 and 100
	reader := bufio.NewReader(os.Stdin) // Create a new buffered reader to read input from standard input

	fmt.Println("Guess the number between 1 and 100!")

	// Loop until the user guesses the correct number
	for {
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You guessed the number!")
			break
		}
	}
}

func Week3Result() string{
	fmt.Println("Welcome to the Number Guessing Game!")
	NumberGuessingGame()
	fmt.Println("Thanks for playing!")
	return "Game finished"
}
