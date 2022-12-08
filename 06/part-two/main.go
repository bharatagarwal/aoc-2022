package main

import "os"

func allUnique(fourteenChar string) bool {
	record := make(map[rune]int)

	for _, char := range fourteenChar {
		record[char] += 1
	}

	for _, count := range record {
		if count > 1 {
			return false
		}
	}

	return true
}

func findLocation(buffer string) int {
	for i := 0; i < len(buffer)-13; i += 1 {
		if allUnique(buffer[i : i+14]) {
			return i + 1 + 13
		}
	}

	return -1
}

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	stringInput := string(input)

	println(findLocation("mjqjpqmgbljsphdztnvjfqwrcgsmlb")) // 30 characters
	println(findLocation("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	println(findLocation("nppdvjthqldpwncqszvftbrmjlhg"))
	println(findLocation("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	println(findLocation("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	println(findLocation(stringInput))
}
