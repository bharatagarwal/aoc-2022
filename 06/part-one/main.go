package main

import "os"

func allUnique(fourChar string) bool {
	record := make(map[rune]int)

	for _, char := range fourChar {
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
	for i := 0; i < len(buffer)-3; i += 1 {
		if allUnique(buffer[i : i+4]) {
			return i + 3 + 1
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

	buffer := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	noOfCharacters := findLocation(buffer)
	println(noOfCharacters)
	println(findLocation("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	println(findLocation("nppdvjthqldpwncqszvftbrmjlhg"))
	println(findLocation("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	println(findLocation("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	println(findLocation(stringInput))
}
