package main

import (
	"errors"
	"fmt"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type stack struct {
	values []string
}

func (s *stack) push(crate string) string {
	s.values = append(s.values, crate)
	return crate
}

func (s *stack) pop() string {
	stackLength := len(s.values)

	if stackLength == 0 {
		panic(errors.New("stack is empty"))
	}

	value := s.values[stackLength-1]
	s.values = s.values[:stackLength-1]

	return value
}

func move(count int, from int, to int) {
	temp := []string{}

	for i := 0; i < count; i += 1 {
		letter := stacks[from-1].pop()
		temp = append(temp, letter)
	}

	for i := len(temp) - 1; i >= 0; i -= 1 {
		stacks[to-1].push(temp[i])
	}
}

var stacks = []stack{
	{values: []string{"F", "C", "P", "G", "Q", "R"}},
	{values: []string{"W", "T", "C", "P"}},
	{values: []string{"B", "H", "P", "M", "C"}},
	{values: []string{"L", "T", "Q", "S", "M", "P", "R"}},
	{values: []string{"P", "H", "J", "Z", "V", "G", "N"}},
	{values: []string{"D", "P", "J"}},
	{values: []string{"L", "G", "P", "Z", "F", "J", "T", "R"}},
	{values: []string{"N", "L", "H", "C", "F", "P", "T", "J"}},
	{values: []string{"G", "V", "Z", "Q", "H", "T", "C", "W"}},
}

// var stacks = []stack{
// 	{values: []string{"Z", "N"}},
// 	{values: []string{"M", "C", "D"}},
// 	{values: []string{"P"}},
// }

func main() {
	// fmt.Println(stacks)
	// move(1, 2, 1)
	// fmt.Println(stacks)
	// move(3, 1, 3)
	// fmt.Println(stacks)
	// move(2, 2, 1)
	// fmt.Println(stacks)
	// move(1, 1, 2)
	// fmt.Println(stacks)

	move(2, 2, 8)
	move(2, 1, 6)
	move(8, 7, 1)
	move(7, 5, 4)
	move(1, 6, 4)
	move(1, 6, 3)
	move(6, 3, 5)
	move(9, 8, 1)
	move(3, 6, 7)
	move(14, 4, 1)
	move(6, 1, 7)
	move(16, 1, 9)
	move(6, 1, 4)
	move(1, 8, 6)
	move(4, 1, 5)
	move(11, 9, 7)
	move(2, 1, 8)
	move(1, 6, 7)
	move(1, 8, 7)
	move(1, 8, 3)
	move(7, 4, 3)
	move(14, 7, 6)
	move(8, 6, 9)
	move(19, 9, 2)
	move(1, 1, 2)
	move(2, 9, 7)
	move(9, 7, 8)
	move(2, 2, 8)
	move(16, 2, 9)
	move(4, 8, 2)
	move(1, 7, 9)
	move(3, 9, 6)
	move(3, 3, 6)
	move(11, 9, 2)
	move(7, 5, 3)
	move(2, 5, 9)
	move(6, 6, 4)
	move(1, 6, 4)
	move(4, 6, 8)
	move(5, 9, 1)
	move(4, 1, 7)
	move(3, 2, 6)
	move(3, 4, 1)
	move(1, 4, 1)
	move(2, 1, 3)
	move(4, 3, 7)
	move(1, 5, 2)
	move(3, 1, 6)
	move(15, 2, 5)
	move(3, 6, 3)
	move(13, 3, 8)
	move(2, 4, 2)
	move(9, 5, 4)
	move(2, 2, 5)
	move(5, 7, 5)
	move(10, 8, 6)
	move(1, 2, 5)
	move(10, 4, 6)
	move(4, 8, 6)
	move(3, 7, 1)
	move(3, 1, 9)
	move(1, 2, 1)
	move(8, 5, 2)
	move(3, 6, 9)
	move(6, 8, 5)
	move(6, 9, 2)
	move(1, 1, 9)
	move(10, 2, 1)
	move(4, 8, 5)
	move(10, 5, 9)
	move(11, 9, 7)
	move(5, 7, 9)
	move(1, 9, 2)
	move(3, 2, 9)
	move(2, 2, 8)
	move(4, 9, 5)
	move(4, 1, 9)
	move(5, 5, 2)
	move(5, 1, 4)
	move(21, 6, 9)
	move(3, 2, 9)
	move(2, 8, 1)
	move(25, 9, 6)
	move(4, 5, 7)
	move(1, 4, 6)
	move(6, 6, 4)
	move(3, 4, 6)
	move(7, 7, 3)
	move(4, 9, 1)
	move(3, 7, 8)
	move(2, 9, 8)
	move(2, 2, 8)
	move(4, 1, 3)
	move(9, 6, 2)
	move(13, 6, 4)
	move(13, 4, 5)
	move(1, 5, 8)
	move(2, 2, 3)
	move(6, 5, 3)
	move(19, 3, 6)
	move(1, 4, 9)
	move(2, 8, 1)
	move(5, 2, 3)
	move(5, 1, 9)
	move(7, 5, 4)
	move(1, 8, 3)
	move(1, 2, 6)
	move(8, 6, 3)
	move(1, 9, 8)
	move(11, 4, 2)
	move(1, 4, 6)
	move(1, 2, 8)
	move(5, 3, 4)
	move(4, 9, 6)
	move(1, 6, 8)
	move(9, 3, 1)
	move(7, 2, 9)
	move(1, 2, 6)
	move(3, 1, 8)
	move(2, 2, 3)
	move(3, 9, 7)
	move(3, 4, 7)
	move(2, 4, 3)
	move(2, 3, 5)
	move(8, 6, 4)
	move(6, 8, 6)
	move(2, 9, 4)
	move(5, 8, 6)
	move(3, 7, 5)
	move(1, 5, 8)
	move(1, 8, 2)
	move(1, 5, 1)
	move(11, 4, 9)
	move(2, 6, 3)
	move(2, 2, 4)
	move(6, 1, 2)
	move(6, 2, 1)
	move(3, 7, 3)
	move(2, 4, 7)
	move(4, 6, 5)
	move(7, 3, 7)
	move(5, 9, 6)
	move(22, 6, 8)
	move(2, 6, 5)
	move(2, 8, 4)
	move(14, 8, 7)
	move(11, 7, 4)
	move(3, 8, 1)
	move(9, 7, 8)
	move(10, 1, 4)
	move(1, 7, 4)
	move(4, 8, 7)
	move(6, 4, 9)
	move(7, 4, 1)
	move(3, 4, 8)
	move(1, 5, 8)
	move(8, 5, 3)
	move(4, 3, 9)
	move(7, 8, 9)
	move(3, 8, 3)
	move(2, 8, 2)
	move(7, 9, 1)
	move(2, 2, 8)
	move(8, 9, 1)
	move(8, 1, 7)
	move(7, 1, 5)
	move(7, 7, 1)
	move(11, 9, 8)
	move(9, 8, 5)
	move(2, 8, 5)
	move(3, 1, 8)
	move(2, 3, 7)
	move(6, 4, 1)
	move(6, 1, 6)
	move(5, 7, 1)
	move(2, 4, 6)
	move(1, 3, 5)
	move(4, 7, 4)
	move(2, 8, 7)
	move(10, 5, 6)
	move(9, 6, 1)
	move(8, 1, 6)
	move(1, 7, 2)
	move(9, 6, 4)
	move(2, 4, 3)
	move(3, 8, 1)
	move(1, 2, 4)
	move(4, 4, 1)
	move(7, 4, 3)
	move(3, 3, 2)
	move(1, 7, 6)
	move(9, 6, 7)
	move(6, 7, 4)
	move(2, 7, 2)
	move(6, 4, 7)
	move(2, 2, 9)
	move(1, 2, 4)
	move(1, 7, 4)
	move(4, 7, 6)
	move(4, 5, 4)
	move(1, 2, 5)
	move(1, 7, 5)
	move(1, 2, 6)
	move(6, 4, 3)
	move(9, 3, 9)
	move(4, 6, 2)
	move(7, 3, 8)
	move(22, 1, 7)
	move(1, 1, 7)
	move(2, 8, 3)
	move(4, 5, 6)
	move(2, 3, 2)
	move(6, 2, 8)
	move(3, 8, 6)
	move(1, 4, 8)
	move(1, 1, 8)
	move(8, 6, 7)
	move(7, 8, 9)
	move(22, 7, 4)
	move(3, 5, 6)
	move(1, 8, 1)
	move(2, 8, 2)
	move(3, 6, 4)
	move(1, 1, 3)
	move(15, 9, 1)
	move(5, 1, 5)
	move(3, 7, 6)
	move(5, 5, 6)
	move(4, 4, 3)
	move(6, 6, 9)
	move(7, 7, 6)
	move(5, 6, 7)
	move(4, 1, 9)
	move(3, 7, 4)
	move(2, 9, 7)
	move(5, 3, 5)
	move(3, 6, 3)
	move(5, 4, 6)
	move(10, 9, 5)
	move(1, 2, 9)
	move(1, 3, 5)
	move(1, 2, 9)
	move(3, 1, 6)
	move(2, 9, 2)
	move(7, 6, 5)
	move(15, 4, 9)
	move(2, 4, 5)
	move(1, 3, 4)
	move(9, 9, 1)
	move(1, 9, 2)
	move(2, 9, 4)
	move(11, 5, 4)
	move(1, 9, 3)
	move(1, 6, 8)
	move(4, 7, 8)
	move(4, 8, 9)
	move(15, 4, 7)
	move(1, 6, 7)
	move(1, 3, 7)
	move(6, 9, 6)
	move(1, 3, 7)
	move(1, 2, 1)
	move(1, 9, 5)
	move(3, 6, 1)
	move(11, 1, 4)
	move(6, 5, 1)
	move(2, 2, 5)
	move(1, 5, 7)
	move(2, 6, 1)
	move(7, 5, 7)
	move(3, 5, 6)
	move(4, 6, 1)
	move(11, 4, 3)
	move(1, 8, 5)
	move(23, 7, 6)
	move(18, 6, 9)
	move(1, 5, 9)
	move(1, 4, 2)
	move(3, 3, 7)
	move(3, 3, 8)
	move(17, 1, 8)
	move(5, 6, 5)
	move(2, 7, 1)
	move(20, 8, 2)
	move(4, 7, 2)
	move(3, 9, 5)
	move(7, 9, 7)
	move(6, 9, 2)
	move(1, 1, 8)
	move(3, 9, 4)
	move(7, 5, 2)
	move(6, 7, 1)
	move(1, 1, 8)
	move(3, 2, 6)
	move(1, 7, 6)
	move(2, 8, 9)
	move(35, 2, 4)
	move(3, 3, 2)
	move(1, 5, 7)
	move(2, 3, 9)
	move(3, 1, 6)
	move(2, 2, 1)
	move(32, 4, 7)
	move(3, 4, 8)
	move(3, 9, 5)
	move(1, 1, 2)
	move(21, 7, 5)
	move(2, 2, 1)
	move(3, 1, 2)
	move(15, 5, 1)
	move(3, 6, 7)
	move(3, 4, 6)
	move(3, 8, 5)
	move(1, 9, 3)
	move(8, 7, 2)
	move(6, 5, 2)
	move(9, 1, 6)
	move(4, 7, 1)
	move(2, 5, 4)
	move(2, 4, 3)
	move(3, 5, 4)
	move(17, 2, 7)
	move(3, 3, 5)
	move(2, 4, 8)
	move(1, 4, 3)
	move(5, 7, 9)
	move(1, 3, 6)
	move(4, 1, 7)
	move(4, 6, 7)
	move(2, 5, 2)
	move(1, 1, 3)
	move(10, 6, 4)
	move(1, 3, 7)
	move(20, 7, 8)
	move(8, 4, 8)
	move(1, 2, 8)
	move(4, 9, 1)
	move(3, 7, 4)
	move(2, 4, 9)
	move(2, 6, 3)
	move(1, 2, 8)
	move(1, 7, 6)
	move(1, 9, 5)
	move(3, 5, 9)
	move(4, 9, 2)
	move(1, 4, 5)
	move(1, 5, 3)
	move(3, 2, 4)
	move(1, 9, 7)
	move(1, 2, 1)
	move(1, 7, 1)
	move(11, 1, 2)
	move(3, 1, 7)
	move(25, 8, 5)
	move(1, 6, 3)
	move(1, 6, 2)
	move(7, 8, 2)
	move(9, 2, 8)
	move(2, 4, 7)
	move(2, 5, 7)
	move(2, 5, 2)
	move(5, 5, 1)
	move(7, 5, 1)
	move(2, 4, 9)
	move(3, 5, 6)
	move(1, 1, 8)
	move(1, 5, 6)
	move(1, 4, 7)
	move(1, 9, 2)
	move(3, 5, 2)
	move(2, 6, 9)
	move(3, 9, 8)
	move(1, 5, 4)
	move(3, 3, 9)
	move(10, 1, 5)
	move(4, 2, 8)
	move(2, 6, 1)
	move(3, 9, 7)
	move(1, 1, 9)
	move(1, 4, 3)
	move(1, 9, 2)
	move(9, 8, 2)
	move(2, 3, 7)
	move(2, 7, 6)
	move(3, 5, 6)
	move(4, 8, 6)
	move(4, 8, 3)
	move(4, 3, 2)
	move(4, 6, 8)
	move(1, 7, 9)
	move(2, 1, 8)
	move(2, 8, 3)
	move(1, 9, 2)
	move(13, 2, 4)
	move(6, 5, 7)
	move(2, 5, 7)
	move(10, 2, 4)
	move(11, 7, 8)
	move(1, 6, 4)
	move(4, 6, 7)
	move(24, 4, 9)
	move(11, 7, 4)
	move(1, 3, 8)
	move(1, 3, 5)
	move(4, 4, 2)
	move(5, 4, 2)
	move(9, 2, 5)
	move(4, 9, 5)
	move(1, 5, 1)
	move(2, 5, 7)
	move(2, 2, 5)
	move(1, 1, 7)
	move(2, 2, 3)
	move(18, 9, 6)
	move(9, 8, 1)
	move(2, 9, 5)
	move(5, 1, 8)
	move(2, 8, 7)
	move(4, 8, 4)
	move(5, 8, 7)
	move(10, 5, 1)
	move(10, 7, 4)
	move(4, 5, 8)
	move(14, 1, 9)
	move(6, 9, 8)
	move(1, 5, 1)
	move(12, 6, 9)
	move(4, 6, 8)
	move(11, 8, 5)
	move(1, 6, 1)
	move(19, 9, 7)
	move(2, 3, 5)
	move(13, 7, 5)
	move(3, 7, 1)
	move(4, 8, 9)
	move(2, 7, 6)
	move(7, 4, 8)
	move(5, 8, 1)
	move(1, 1, 3)
	move(1, 7, 2)
	move(6, 1, 6)
	move(1, 2, 5)
	move(1, 8, 1)
	move(1, 8, 2)
	move(2, 4, 8)
	move(5, 6, 1)
	move(2, 4, 7)
	move(2, 9, 6)
	move(1, 6, 5)
	move(4, 6, 2)
	move(1, 9, 5)
	move(2, 4, 5)
	move(4, 2, 4)
	move(2, 8, 3)
	move(3, 3, 2)
	move(4, 1, 2)
	move(2, 4, 7)
	move(4, 2, 3)
	move(4, 1, 2)
	move(13, 5, 1)
	move(1, 6, 2)
	move(1, 1, 8)
	move(15, 5, 2)
	move(4, 3, 1)
	move(5, 4, 3)
	move(1, 3, 6)
	move(1, 8, 7)
	move(1, 9, 8)
	move(1, 7, 8)
	move(3, 3, 2)
	move(1, 8, 2)
	move(1, 3, 7)
	move(13, 1, 4)
	move(3, 5, 3)
	move(1, 1, 2)
	move(1, 8, 5)
	move(5, 7, 2)
	move(1, 6, 5)
	move(2, 3, 4)
	move(10, 2, 5)
	move(1, 9, 5)
	move(3, 1, 8)
	move(3, 8, 3)
	move(11, 4, 5)
	move(12, 2, 8)
	move(4, 4, 7)
	move(10, 8, 5)
	move(2, 8, 1)
	move(1, 7, 3)
	move(1, 7, 9)
	move(5, 3, 7)
	move(1, 9, 4)
	move(7, 7, 6)
	move(13, 5, 8)
	move(6, 6, 7)
	move(5, 7, 4)
	move(1, 6, 4)
	move(2, 4, 9)
	move(1, 7, 9)
	move(3, 4, 3)
	move(1, 3, 6)
	move(4, 5, 7)

	for _, stack := range stacks {
		fmt.Printf("%s", stack.values[len(stack.values)-1])
	}
}