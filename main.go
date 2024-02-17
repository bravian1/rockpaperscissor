package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Declare the playerScore and computerScore variables as global
var playerScore int
var computerScore int

func generateComputerChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	randomIndex := rand.Intn(len(choices))
	return choices[randomIndex]
}

func determineWinner(playerChoice, computerChoice string) {
	// Your code for tieMessages, playerWinMessages, computerWinMessages
	tieMessages := []string{
		"It's a tie!",
		"Nobody wins this round.",
		"You both chose the same thing.",
	}

	playerWinMessages := []string{
		"You win! You have bested the computer with your skills.",
		"The computer is no match for you. You rock!",
		"You have defeated the computer. Congratulations!",
		"You are the champion of rock, paper, scissors. Well done!",
		"You have outwitted the computer. You win this round.",
	}

	computerWinMessages := []string{
		"You have been defeated by the superior intelligence of the computer. I win this game of rock, paper, scissors!",
		"You have no idea what you are up against. I have calculated every possible outcome. I win this game of rock, paper, scissors!",
		"You have met your match. I am the ultimate rock, paper, scissors player. You lose!",
		"You have made a fatal error. I have anticipated your every move. I win this game of rock, paper, scissors!",
		"You have been outsmarted by the computer. I am the master of rock, paper, scissors. You are finished!",
	}

	switch playerChoice {
	case "rock":
		if computerChoice == "rock" {
			fmt.Println(tieMessages[rand.Intn(len(tieMessages))])
		} else if computerChoice == "paper" {
			fmt.Println(computerWinMessages[rand.Intn(len(computerWinMessages))])
			computerScore++
		} else if computerChoice == "scissors" {
			fmt.Println(playerWinMessages[rand.Intn(len(playerWinMessages))])
			playerScore++
		}
	case "paper":
		if computerChoice == "rock" {
			fmt.Println(playerWinMessages[rand.Intn(len(playerWinMessages))])
			playerScore++
		} else if computerChoice == "paper" {
			fmt.Println(tieMessages[rand.Intn(len(tieMessages))])
		} else if computerChoice == "scissors" {
			fmt.Println(computerWinMessages[rand.Intn(len(computerWinMessages))])
			computerScore++
		}
	case "scissors":
		if computerChoice == "rock" {
			fmt.Println(computerWinMessages[rand.Intn(len(computerWinMessages))])
			computerScore++
		} else if computerChoice == "paper" {
			fmt.Println(playerWinMessages[rand.Intn(len(playerWinMessages))])
			playerScore++
		} else if computerChoice == "scissors" {
			fmt.Println(tieMessages[rand.Intn(len(tieMessages))])
		}
	default:
		fmt.Println("Invalid choice. Please choose rock, paper, or scissors.")
	}
}

func main() {
	// Seed the random number generator once at the beginning
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Rock Paper Scissors!")
	fmt.Println("How many sets (best-of-X matches) to determine the ultimate champion?")
	targetSetsInput, _ := reader.ReadString('\n')
	targetSetsInput = strings.TrimSpace(targetSetsInput)
	targetSets, err := strconv.Atoi(targetSetsInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	gameSets := 0
	for gameSets < targetSets {
		fmt.Println("How many wins for a set victory? (e.g., 3)")
		targetWins, _ := reader.ReadString('\n')
		targetWins = strings.TrimSpace(targetWins)
		targetWinsInt, err := strconv.Atoi(targetWins)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue // Skip to the next set if invalid input
		}

		// Reset the playerScore and computerScore variables for each set
		playerScore = 0
		computerScore = 0

		for playerScore < targetWinsInt && computerScore < targetWinsInt {
			fmt.Println("Current Score - Player:", playerScore, "Computer:", computerScore)
			fmt.Println("Choose rock, paper, or scissors:")
			playerChoice, _ := reader.ReadString('\n')
			playerChoice = strings.ToLower(strings.TrimSpace(playerChoice))

			computerChoice := generateComputerChoice()
			fmt.Println("Computer chose:", computerChoice)

			determineWinner(playerChoice, computerChoice)
		}

		if playerScore > computerScore {
			fmt.Println("You win this set!")
		} else {
			fmt.Println("The computer wins this set!")
		}

		gameSets++ // Increment after a set is completed
	}

	fmt.Println("All sets completed. Thanks for playing!") // Game over
}
