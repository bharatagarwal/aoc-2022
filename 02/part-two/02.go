package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type round struct {
	yourMove      string
	opponentsMove string
	score         int
}

func main() {
	result("rock", "scissors")
	input, err := os.ReadFile("input.txt") // byte array
	handle(err)

	var game []round
	var total int

	gameRecord := strings.Split(string(input), "\n")
	gameRecord = gameRecord[:len(gameRecord)-1]

	for c, move := range gameRecord {
		moves := strings.Split(move, " ")
		currentRound := round{}
		fmt.Printf("%d- %s\n", c, move)

		currentRound.opponentsMove, err = translateMove(moves[0])
		handle(err)

		result, err := translateResult(moves[1])
		fmt.Println(result)
		handle(err)

		currentRound.yourMove = findRelevantMove(result, currentRound.opponentsMove)
		fmt.Println(currentRound.yourMove)
		fmt.Println(currentRound.opponentsMove)
		score := computeScore(currentRound.yourMove, currentRound.opponentsMove)
		currentRound.score = score

		total += score
		game = append(game, currentRound)
	}

	fmt.Println(total)

}

func translateMove(move string) (string, error) {
	switch move {
	case "A":
		return "rock", nil
	case "B":
		return "paper", nil
	case "C":
		return "scissors", nil
	}

	return "", errors.New("incorrect data")
}

func translateResult(outcome string) (string, error) {
	switch outcome {
	case "X":
		return "loss", nil
	case "Y":
		return "draw", nil
	case "Z":
		return "win", nil
	}

	return "", errors.New("incorrect data")
}

func findRelevantMove(result string, opponentMove string) (string) {
	var outcome string

	switch result {
	case "win":
		switch opponentMove {
		case "rock":
			outcome = "paper"
		case "paper":
			outcome = "scissors"
		case "scissors":
			outcome = "rock"
		}
	case "draw":
		switch opponentMove {
		case "rock":
			outcome = "rock"
		case "paper":
			outcome = "paper"
		case "scissors":
			outcome = "scissors"
		}
	case "loss":
		switch opponentMove {
		case "rock":
			outcome = "scissors"
		case "paper":
			outcome = "rock"
		case "scissors":
			outcome = "paper"
		}
		
	}

	return outcome
}

func result(yourMove string, opponentsMove string) string {

	var outcome string
	switch yourMove {
	case "rock":
		switch opponentsMove {
		case "rock":
			outcome = "draw"
		case "paper":
			outcome = "loss"
		case "scissors":
			outcome = "win"
		}
	case "paper":
		switch opponentsMove {
		case "rock":
			outcome = "win"
		case "paper":
			outcome = "draw"
		case "scissors":
			outcome = "loss"
		}
	case "scissors":
		switch opponentsMove {
		case "rock":
			outcome = "loss"
		case "paper":
			outcome = "win"
		case "scissors":
			outcome = "draw"
		}
	}

	return outcome
}

func computeScore(yourMove string, opponentsMove string) int {
	var total int
	switch yourMove {
	case "rock":
		total += 1
	case "paper":
		total += 2
	case "scissors":
		total += 3
	}

	switch result(yourMove, opponentsMove) {
	case "win":
		total += 6
	case "draw":
		total += 3
	}

	return total
}
