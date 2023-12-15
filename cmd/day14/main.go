package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"log"

	"github.com/thoas/go-funk"
)

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func moveRockInDir(grid [][]byte, rowIdx, colIdx, dr, dc int) {
	for {
		nr := rowIdx + dr
		nc := colIdx + dc
		if !((nr >= 0 && nr < len(grid)) && (nc >= 0 && nc < len(grid[0]))) {
			return
		}
		if grid[nr][nc] == '.' {
			grid[rowIdx][colIdx] = '.'
			grid[nr][nc] = 'O'
		} else {
			return
		}
		rowIdx = nr
		colIdx = nc
	}
}

func tiltGridNorth(grid [][]byte) {
	for rowIdx, row := range grid {
		for colIdx, col := range row {
			if col == 'O' {
				moveRockInDir(grid, rowIdx, colIdx, -1, 0)
			}
		}
	}
}

func tiltGridWest(grid [][]byte) {
	for rowIdx, row := range grid {
		for colIdx, col := range row {
			if col == 'O' {
				moveRockInDir(grid, rowIdx, colIdx, 0, -1)
			}
		}
	}
}

func tiltGridSouth(grid [][]byte) {
	for rowIdx := len(grid) - 1; rowIdx >= 0; rowIdx-- {
		row := grid[rowIdx]
		for colIdx, col := range row {
			if col == 'O' {
				moveRockInDir(grid, rowIdx, colIdx, 1, 0)
			}
		}
	}
}

func tiltGridEast(grid [][]byte) {
	for rowIdx, row := range grid {
		for colIdx := len(row) - 1; colIdx >= 0; colIdx-- {
			col := row[colIdx]
			if col == 'O' {
				moveRockInDir(grid, rowIdx, colIdx, 0, 1)
			}
		}
	}
}

func spinGrid(grid [][]byte) {
	tiltGridNorth(grid)
	tiltGridWest(grid)
	tiltGridSouth(grid)
	tiltGridEast(grid)
}

func calculateLoad(grid [][]byte) int {
	ans := 0
	R := len(grid)
	for i, row := range grid {
		numRocks := 0
		for _, col := range row {
			if col == 'O' {
				numRocks++
			}
		}
		ans += (R - i) * numRocks
	}
	return ans
}

func solvePart1(lines []string) int {
	grid := aoc.Map(lines, func(line string) []byte { return []byte(line) })
	tiltGridNorth(grid)
	return calculateLoad(grid)
}

func solvePart2(lines []string) int {
	n := int(1e9)
	grid := aoc.Map(lines, func(line string) []byte { return []byte(line) })
	seen := map[string]int{}
	for i := 0; i <= n; i++ {
		spinGrid(grid)
		// flatten 2D byte array and stringify so that it can be hashed and used as a map key
		serialized := string(funk.Flatten(grid).([]byte))
		repeatedIdx, isRepeated := seen[serialized]
		if isRepeated {
			cycleLen := i - repeatedIdx
			// https://www.reddit.com/r/adventofcode/comments/18i0xtn/comment/kddj1rn/?utm_source=share&utm_medium=web2x&context=3
			// numSpinsNeeded := (n - preamble) % cycle
			numSpinsNeeded := (n - (repeatedIdx + 1)) % cycleLen
			for j := 0; j < numSpinsNeeded; j++ {
				spinGrid(grid)
			}
			return calculateLoad(grid)
		}
		seen[serialized] = i
	}
	log.Fatal("shouldn't reach here")
	return -1
}

func main() {
	lines := aoc.GetInputLinesForDay(14, false)
	p1 := solvePart1(lines)
	fmt.Println("Part 1:", p1)
	p2 := solvePart2(lines)
	fmt.Println("Part 2:", p2)
}
