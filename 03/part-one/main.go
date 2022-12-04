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

	for _, line := range lines {
		lineLength := len(line)
		firstHalf := line[:(lineLength / 2)]
		secondHalf := line[(lineLength / 2):]

		firstHalfPriorities := translateString(firstHalf)
		secondHalfPriorities := translateString(secondHalf)

		sort.Ints(firstHalfPriorities)
		sort.Ints(secondHalfPriorities)
		commonItems = append(commonItems, common(firstHalfPriorities, secondHalfPriorities))
	}

	for _, item := range commonItems {
		sum += item
	}

	fmt.Println(sum)
}

func common(first []int, second []int) int {
	var firstPointer, secondPointer int

	for {
		if first[firstPointer] == second[secondPointer] {
			fmt.Println(first[firstPointer])
			return first[firstPointer]
		}

		if first[firstPointer] > second[secondPointer] {
			secondPointer += 1
		}

		if first[firstPointer] < second[secondPointer] {
			firstPointer += 1
		}

		if firstPointer == len(first) || secondPointer == len(second) {
			break
		}
	}

	return -1
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
