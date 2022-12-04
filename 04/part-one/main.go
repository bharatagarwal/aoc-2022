package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines, err := os.ReadFile("input.txt")
	handle(err)

	pairs := strings.Split(string(lines), "\n")
	fmt.Println(pairs)

	var count int

	for _, pair := range pairs {
		firstRange := strings.Split(pair, ",")[0]
		secondRange := strings.Split(pair, ",")[1]

		firstStarting, firstEnding := rangeIntEndpoints(firstRange)
		secondStarting, secondEnding := rangeIntEndpoints(secondRange)

		firstRangeElements := storeRange(firstStarting, firstEnding)
		secondRangeElements := storeRange(secondStarting, secondEnding)

		if firstContainsSecond(firstRangeElements, secondRangeElements) ||
			firstContainsSecond(secondRangeElements, firstRangeElements) {
			count += 1
		}

	}

	fmt.Println(count)

}

func firstContainsSecond(first map[int]bool, second map[int]bool) bool {
	allCovered := true

	for key := range first {
		_, ok := second[key]
		if !ok {
			allCovered = false
		}
	}

	if allCovered {
		return true
	}

	return false
}

func rangeIntEndpoints(stringRange string) (int, int) {
	first := strings.Split(stringRange, "-")[0]
	second := strings.Split(stringRange, "-")[1]

	firstInt, err := strconv.Atoi(first)
	handle(err)
	secondInt, err := strconv.Atoi(second)
	handle(err)

	return firstInt, secondInt
}

func storeRange(first int, last int) map[int]bool {
	hash := make(map[int]bool)

	for i := first; i <= last; i++ {
		hash[i] = true
	}

	return hash
}
