package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generateComputerChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	randomIndex := rand.Intn(len(choices))
	return choices[randomIndex]
}

func determineWinner(playerChoice, computerChoice string) {
	switch playerChoice {
	case computerChoice:
		fmt.Println("A stalemate! Try again.") // Fun "tie" message
	case "rock":
		if computerChoice == "scissors" {
			fmt.Println("You crush the scissors! Victory!")
		} else {
			fmt.Println("Alas, your rock is smothered by paper. Defeat!")
		}
	case "paper":
		if computerChoice == "rock" {
			fmt.Println("Your paper smothers the rock! A glorious win!")
		} else {
			fmt.Println("The computer's scissors shred your paper. Try again!")
		}
	case "scissors":
		if computerChoice == "paper" {
			fmt.Println("Snip snip! Your scissors shred the paper. You win!")
		} else {
			fmt.Println("Your scissors are no match for the mighty rock. Defeat!")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose rock, paper, or scissors:")
		playerChoice, _ := reader.ReadString('\n')
		playerChoice = strings.ToLower(strings.TrimSpace(playerChoice))

		computerChoice := generateComputerChoice()

		fmt.Println("Computer chose:", computerChoice) // Display computer's choice
		determineWinner(playerChoice, computerChoice)

		fmt.Println("Play again? (yes/no)")
		playAgain, _ := reader.ReadString('\n')
		playAgain = strings.ToLower(strings.TrimSpace(playAgain))

		if playAgain != "yes" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
