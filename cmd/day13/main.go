package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"log"
	"strings"
)

func numRowDiffs(grid [][]byte, rowIdx1, rowIdx2 int) int {
	diff := 0
	for colIdx := 0; colIdx < len(grid[0]); colIdx++ {
		if grid[rowIdx1][colIdx] != grid[rowIdx2][colIdx] {
			diff++
		}
	}
	return diff
}

func numColDiffs(grid [][]byte, colIdx1, colIdx2 int) int {
	diff := 0
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		if grid[rowIdx][colIdx1] != grid[rowIdx][colIdx2] {
			diff++
		}
	}
	return diff
}

func numSymmetricCols(grid [][]byte, isPart2 bool) int {
	C := len(grid[0])
	for colIdx := 0; colIdx < C; colIdx++ {
		fixedSmudge := false
		i := 0
		for {
			if colIdx-i < 0 || colIdx+i+1 >= C {
				break
			}

			colDiffs := numColDiffs(grid, colIdx-i, colIdx+1+i)

			if isPart2 && colDiffs == 1 && !fixedSmudge {
				fixedSmudge = true
				i++
				continue
			}

			if colDiffs != 0 {
				break
			}

			i++
		}
		if i > 0 && (fixedSmudge || !isPart2) && (colIdx+1-i == 0 || colIdx+i == C-1) {
			return colIdx + 1
		}
	}
	return -1
}

func numSymmetricRows(grid [][]byte, isPart2 bool) int {
	R := len(grid)
	for rowIdx := 0; rowIdx < R; rowIdx++ {
		fixedSmudge := false
		i := 0
		for {
			if rowIdx-i < 0 || rowIdx+i+1 >= R {
				break
			}

			rowDiffs := numRowDiffs(grid, rowIdx-i, rowIdx+1+i)

			if isPart2 && rowDiffs == 1 && !fixedSmudge {
				fixedSmudge = true
				i++
				continue
			}

			if rowDiffs != 0 {
				break
			}
			i++
		}
		if i > 0 && (fixedSmudge || !isPart2) && (rowIdx+1-i == 0 || rowIdx+i == R-1) {
			return rowIdx + 1
		}
	}
	return -1
}

func solve(input string, isPart2 bool) int {
	patterns := strings.Split(input, "\n\n")
	ans := 0
	for _, pattern := range patterns {
		grid := aoc.Map(strings.Split(pattern, "\n"), func(line string) []byte { return []byte(line) })
		sCol := numSymmetricCols(grid, isPart2)
		sRow := numSymmetricRows(grid, isPart2)
		if sCol == sRow {
			log.Fatal("invalid pattern")
		}
		if sCol != -1 {
			ans += sCol
		}
		if sRow != -1 {
			ans += (100 * sRow)
		}
	}
	return ans
}

func main() {
	input := aoc.GetInputForDay(13, false)
	p1 := solve(input, false)
	p2 := solve(input, true)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
