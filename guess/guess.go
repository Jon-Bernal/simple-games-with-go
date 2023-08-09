package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	low := 1
	high := 100
	tries := 0
	prevGuess := 0
	
	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan()
	
	for {
		guess := (low + high) / 2
		if guess == prevGuess {
			fmt.Println("You're CHEATTING!!! I QUIT.")
			break
		}
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that :")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()
		response := scanner.Text()
		tries++


		prevGuess = guess
		if response == "a" {
			high = guess -1
			} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won in", tries, "tries!")
			break
		} else if response == "exit" {
			break
		} else {
			fmt.Println("Invalid response, try again")
		}
	}
}