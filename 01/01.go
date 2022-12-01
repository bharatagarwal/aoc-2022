package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	// returns a string of the input
	// in the form of a byte array

	if err != nil {
		log.Fatal(err)
	}

	var calories []int

	stringInput := string(input[:])
	elfCalories := strings.Split(stringInput, "\n\n")

	for _, word := range elfCalories {
		individualCalories := strings.Split(word, "\n")
		var total int
		for _, calories := range individualCalories {
			integerCalorie, err := strconv.Atoi(calories)

			if err != nil {
				break
			}

			total += integerCalorie
		}

		calories = append(calories, total)
	}

	sortedCalories := make([]int, len(calories))
	copy(sortedCalories, calories)
	sort.Ints(sortedCalories)

	noOfElves := len(sortedCalories)

	fmt.Printf("Maximum number of calories: %d\n", sortedCalories[noOfElves-1])

	topThree := sortedCalories[noOfElves-1] +
		sortedCalories[noOfElves-2] +
		sortedCalories[noOfElves-3]

	fmt.Printf("Top three calories: %d\n", topThree)
}
