package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"strings"
)

func getNewSequence(nums []int) []int {
	newSeq := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		newSeq[i] = nums[i+1] - nums[i]
	}
	return newSeq
}
func isAllZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true

}

func solve(nums []int, isPart2 bool) int {
	if isAllZeroes(nums) {
		return 0
	}
	seq := getNewSequence(nums)
	if isPart2 {
		return nums[0] - solve(seq, isPart2)
	}
	return nums[len(nums)-1] + solve(seq, isPart2)
}

func main() {
	lines := aoc.GetInputLinesForDay(9, false)
	p1 := 0
	p2 := 0
	for _, line := range lines {
		nums := aoc.MapStrToInt(strings.Fields(line))
		p1 += solve(nums, false)
		p2 += solve(nums, true)
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
