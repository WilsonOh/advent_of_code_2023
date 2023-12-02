package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func solve(lines []string, partTwo bool) int {
	ans := 0
	letterDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, line := range lines {
		var digits []string
		for idx, letter := range line {
			if unicode.IsDigit(letter) {
				digits = append(digits, string(letter))
			}
			if !partTwo {
				continue
			}
			for digit, letterDigit := range letterDigits {
				if strings.HasPrefix(line[idx:], letterDigit) {
					digits = append(digits, strconv.Itoa(digit+1))
				}
			}
		}
		combinedDigit := digits[0] + digits[len(digits)-1]
		num, err := strconv.Atoi(combinedDigit)
		if err != nil {
			log.Fatalf("failed to convert string %s to an int", combinedDigit)
		}
		ans += num
	}
	return ans
}

func main() {
	lines := aoc.GetInputLinesForDay(1)
	p1 := solve(lines, false)
	p2 := solve(lines, true)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
}
