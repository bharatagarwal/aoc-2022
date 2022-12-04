package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	byteLines, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(byteLines), "\n")
	lines = lines[:len(lines)-1]

	var commonItems []int
	var sum int

	// divide lines into groups of 3
	batches := divideIntoBatches(lines)

	// translate lines to priorities
	priorityBatches := translate(batches)

	for _, batch := range priorityBatches {
		for _, priorities := range batch {
			fmt.Println(priorities)
		}
		println()
	}

	commonItems = findCommonInBatches(priorityBatches) // find common item in each trio
	fmt.Println(commonItems)

	for _, item := range commonItems {
		sum += item
	}

	fmt.Println(sum)
}

func findCommonInBatches(batches [][][]int) []int {
	var common []int

	for _, batch := range batches {
		firstRecord := make(map[int]bool)
		secondRecord := make(map[int]bool)
		thirdRecord := make(map[int]bool)

		for i := 0; i < len(batch[0]); i++ {
			firstRecord[batch[0][i]] = true
		}

		for i := 0; i < len(batch[1]); i++ {
			secondRecord[batch[1][i]] = true
		}

		for i := 0; i < len(batch[2]); i++ {
			thirdRecord[batch[2][i]] = true
		}

		var foundCommonInBatch bool

		for i := 0; i < len(batch[0]); i++ {
			_, presentInSecond := secondRecord[batch[0][i]]
			_, presentInThird := thirdRecord[batch[0][i]]

			if presentInSecond && presentInThird && foundCommonInBatch == false {
				foundCommonInBatch = true
				common = append(common, batch[0][i])
			}
		}
	}

	return common
}

func translate(batches [][]string) [][][]int {
	var priorityBatches [][][]int

	for _, batch := range batches {
		var priorities [][]int

		for _, line := range batch {
			priorityForLines := translateString(line)
			sort.Ints(priorityForLines)
			priorities = append(priorities, priorityForLines)
		}

		priorityBatches = append(priorityBatches, priorities)

	}

	return priorityBatches
}

func common(first []int, second []int, third []int) int {
	var firstPointer, secondPointer, thirdPointer int
	fmt.Println(firstPointer, secondPointer, thirdPointer)

	return -1
}

func divideIntoBatches(lines []string) [][]string {
	var batches [][]string

	for i := 0; i < len(lines); i += 3 {
		batches = append(batches, lines[i:i+3])
	}

	return batches
}

func translateString(compartment string) []int {
	var priorities []int

	for i := range compartment {
		asciiValue := int(compartment[i])
		priorities = append(priorities, translateInt(asciiValue))
	}

	return priorities

}

func translateInt(value int) int {
	var priority int

	if value >= 97 && value <= 122 {
		priority = value - 96
	} else if value >= 65 && value <= 90 {
		priority = value - 38
	} else if value == 0 {
		return -1
	}

	return priority
}
