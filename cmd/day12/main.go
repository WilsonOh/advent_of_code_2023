package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"slices"
	"strings"

	"github.com/thoas/go-funk"
)

type RowState struct {
	row string
	idx int
	ch  byte
}

func isValidRow(row string, blocks []int) bool {
	cBlocks := strings.FieldsFunc(row, func(r rune) bool { return r == '.' })
	blockSizes := aoc.Map(cBlocks, func(b string) int { return len(b) })
	return slices.Equal(blockSizes, blocks)
}

func getNumValidRows(row string, blocks []int, idx int, cache map[RowState]int) int {
	if idx == len(row) {
		if isValidRow(row, blocks) {
			return 1
		} else {
			return 0
		}
	}
	if val, ok := cache[RowState{row, idx, row[idx]}]; ok {
		return val
	}
	ret := 0
	if row[idx] == '?' {
		ret += getNumValidRows(row[:idx]+"."+row[idx+1:], blocks, idx+1, cache)
		ret += getNumValidRows(row[:idx]+"#"+row[idx+1:], blocks, idx+1, cache)
	} else {
		ret += getNumValidRows(row, blocks, idx+1, cache)
	}
	cache[RowState{row, idx, row[idx]}] = ret
	return ret
}

func unFoldRow(row string, blocks []int) (string, []int) {
	newRowList := []string{}
	newBlocksList := [][]int{}
	for i := 0; i < 5; i++ {
		newRowList = append(newRowList, row)
		newBlocksList = append(newBlocksList, blocks)
	}
	newRow := strings.Join(newRowList, "?")
	newBlocks := funk.Flatten(newBlocksList)
	return newRow, newBlocks.([]int)
}

func solvePart1(lines []string) int {
	ans := 0
	for _, line := range lines {
		cache := map[RowState]int{}
		tokens := strings.Split(line, " ")
		row := tokens[0]
		blocks := aoc.MapStrToInt(strings.Split(tokens[1], ","))
		ans += getNumValidRows(row, blocks, 0, cache)
	}
	return ans
}

func solvePart2(lines []string) int {
	ans := 0
	for _, line := range lines {
		cache := map[RowState]int{}
		tokens := strings.Split(line, " ")
		row := tokens[0]
		blocks := aoc.MapStrToInt(strings.Split(tokens[1], ","))
		newRow, newBlocks := unFoldRow(row, blocks)
		ans += getNumValidRows(newRow, newBlocks, 0, cache)
	}
	return ans
}

func main() {
	lines := aoc.GetInputLinesForDay(12, false)
	p1 := solvePart1(lines)
	fmt.Println("Part 1:", p1)
	// p2 := solvePart2(lines)
	// fmt.Println("Part 2:", p2)
}
