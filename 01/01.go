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
	
	if err != nil {
		log.Fatal(err)
	}

	var calories []int

	stringInput := string(input[:])
	caloriesForEachElf := strings.Split(stringInput, "\n\n")

	for _, calorieSets := range caloriesForEachElf {
		individualCalories := strings.Split(calorieSets, "\n")

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

	sumOfTopThreeElves := sortedCalories[noOfElves-1] +
		sortedCalories[noOfElves-2] +
		sortedCalories[noOfElves-3]

	fmt.Printf("Top three calories: %d\n", sumOfTopThreeElves)
}
