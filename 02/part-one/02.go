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
	input, err := os.ReadFile("input.txt")

	handle(err)

	var game []round
	var total int

	gameRecord := strings.Split(string(input), "\n")
	gameRecord = gameRecord[:len(gameRecord)-1]

	for _, move := range gameRecord {
		moves := strings.Split(move, " ")
		currentRound := round{}

		currentRound.opponentsMove, err = translateMove(moves[0])
		handle(err)

		currentRound.yourMove, err = translateMove(moves[1])
		handle(err)

		score := computeScore(currentRound.yourMove, currentRound.opponentsMove)
		currentRound.score = score

		total += score
		game = append(game, currentRound)
	}

	fmt.Println(total)

}

func translateMove(move string) (string, error) {
	switch move {
	case "A", "X":
		return "rock", nil
	case "B", "Y":
		return "paper", nil
	case "C", "Z":
		return "scissors", nil
	default:
		return "", errors.New("incorrect data")
	}
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
		default:
			outcome = ""
		}
	case "paper":
		switch opponentsMove {
		case "rock":
			outcome = "win"
		case "paper":
			outcome = "draw"
		case "scissors":
			outcome = "loss"
		default:
			outcome = ""
		}
	case "scissors":
		switch opponentsMove {
		case "rock":
			outcome = "loss"
		case "paper":
			outcome = "win"
		case "scissors":
			outcome = "draw"
		default:
			outcome = ""
		}
	default:
		outcome = ""
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
