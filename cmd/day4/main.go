package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func Map[T any, U any](l []T, f func(t T) U) []U {
	ret := make([]U, len(l))
	for i, a := range l {
		ret[i] = f(a)
	}
	return ret
}

func StrToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type Card struct {
	WinningNumbers []int
	NumbersOnHand  []int
}

func parseLineIntoCard(line string) Card {
	re := regexp.MustCompile(`\s+`)
	tokens := regexp.MustCompile(`\s*:\s*`).Split(line, -1)
	numbers := regexp.MustCompile(`\s*\|\s*`).Split(tokens[1], -1)
	winningNumbers := Map(re.Split(numbers[0], -1), StrToInt)
	numbersOnHand := Map(re.Split(numbers[1], -1), StrToInt)
	slices.Sort(numbersOnHand)
	return Card{WinningNumbers: winningNumbers, NumbersOnHand: numbersOnHand}
}

func (this *Card) CalculateTotalPoints() int {
	numWinningNumbers := 0
	for _, winningNumber := range this.WinningNumbers {
		_, found := slices.BinarySearch(this.NumbersOnHand, winningNumber)
		if found {
			numWinningNumbers++
		}
	}
	if numWinningNumbers == 0 {
		return 0
	}
	return (1 << (numWinningNumbers - 1))
}

func (this *Card) GetNumWinningNumbers() int {
	numWinningNumbers := 0
	for _, winningNumber := range this.WinningNumbers {
		_, found := slices.BinarySearch(this.NumbersOnHand, winningNumber)
		if found {
			numWinningNumbers++
		}
	}
	return numWinningNumbers
}

func solvePart1(lines []string) int {
	ans := 0
	for _, line := range lines {
		card := parseLineIntoCard(line)
		ans += card.CalculateTotalPoints()
	}
	return ans
}

func solvePart2(lines []string) int {
	cardNumMap := make([]int, len(lines))
	for idx := range lines {
		cardNumMap[idx] = 1
	}

	for idx, line := range lines {
		card := parseLineIntoCard(line)
		numWinningNumbers := card.GetNumWinningNumbers()
		lastCard := min(idx+numWinningNumbers, len(lines)-1)
		for id := idx + 1; id <= lastCard; id++ {
			cardNumMap[id] += cardNumMap[idx]
		}
	}
	ans := 0
	for _, numCopies := range cardNumMap {
		ans += numCopies
	}
	return ans
}

func main() {
	lines := aoc.GetInputLinesForDay(4, false)
	p1 := solvePart1(lines)
	p2 := solvePart2(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
