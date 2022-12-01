package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type elf struct {
	calories []int
	total    int
}

func main() {
	input, err := os.ReadFile("input.txt")
	// returns a string of the input
	// in the form of a byte array

	if err != nil {
		log.Fatal(err)
	}

	var elves []elf

	stringInput := string(input[:])
	elfCalories := strings.Split(stringInput, "\n\n")

	for _, word := range elfCalories {
		currentElf := elf{}
		individualCalories := strings.Split(word, "\n")

		for _, calories := range individualCalories {
			integerCalorie, err := strconv.Atoi(calories)

			if err != nil {
				break
			}

			currentElf.calories = append(currentElf.calories, integerCalorie)
			currentElf.total += integerCalorie
		}

		elves = append(elves, currentElf)
	}

	var elfTotals []int

	for _, currentElf := range elves {
		elfTotals = append(elfTotals, currentElf.total)
	}

	sortedTotals := make([]int, len(elfTotals))
	copy(sortedTotals, elfTotals)
	sort.Ints(sortedTotals)
	fmt.Printf("Maximum number of calories: %d\n", sortedTotals[len(sortedTotals)-1])
	topThree := sortedTotals[len(sortedTotals)-1] +
		sortedTotals[len(sortedTotals)-2] +
		sortedTotals[len(sortedTotals)-3]

	fmt.Printf("Top three calories: %d\n", topThree)
}
