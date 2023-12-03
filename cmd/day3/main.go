package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"strconv"
	"unicode"
)

type GearCoords struct {
	Row int
	Col int
}

func isPartNum(lines []string, row, col int) bool {
	for dr := -1; dr < 2; dr++ {
		for dc := -1; dc < 2; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			nr := row + dr
			nc := col + dc
			if (nr >= 0 && nr < len(lines)) && (nc >= 0 && nc < len(lines[0])) {
				ch := lines[nr][nc]
				if !unicode.IsDigit(rune(ch)) && ch != '.' {
					return true
				}
			}
		}
	}
	return false
}

func findAdjGear(lines []string, row, col int) (GearCoords, bool) {
	for dr := -1; dr < 2; dr++ {
		for dc := -1; dc < 2; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			nr := row + dr
			nc := col + dc
			if (nr >= 0 && nr < len(lines)) && (nc >= 0 && nc < len(lines[0])) {
				ch := lines[nr][nc]
				if ch == '*' {
					return GearCoords{Row: nr, Col: nc}, true
				}

			}
		}
	}
	return GearCoords{}, false
}

func solvePart1(lines []string) int {
	ans := 0
	for row, line := range lines {
		currNum := 0
		currNumIsPartNum := false
		for col, ch := range line {
			if unicode.IsDigit(ch) {
				n, _ := strconv.Atoi(string(ch))
				currNum = (currNum * 10) + n
				if isPartNum(lines, row, col) {
					currNumIsPartNum = true
				}
			} else {
				if currNumIsPartNum {
					ans += currNum
				}
				currNum = 0
				currNumIsPartNum = false
			}
		}
		// account for if the end of a digit is the end of a line
		if currNumIsPartNum {
			ans += currNum
		}
	}
	return ans
}

func solvePart2(lines []string) int {
	ans := 0
	// map of coordinates of a gear to its adjacent part numbers
	gearMap := map[GearCoords][]int{}
	for row, line := range lines {
		currNum := 0
		// use a map as a psuedo-set since go doesn't have a built-in set
		adjGears := map[GearCoords]bool{}
		for col, ch := range line {
			if unicode.IsDigit(ch) {
				adjGearCoords, hasAdjGear := findAdjGear(lines, row, col)
				if hasAdjGear {
					// store adjacent gear coords in the "set"
					adjGears[adjGearCoords] = true
				}
				n, _ := strconv.Atoi(string(ch))
				currNum = (currNum * 10) + n
			} else {
				// append the current part number to all its adjacent gears
				for adjGear := range adjGears {
					gearMap[adjGear] = append(gearMap[adjGear], currNum)
				}
				adjGears = map[GearCoords]bool{}
				currNum = 0
			}
		}
		// account for if the end of a digit is the end of a line
		for adjGear := range adjGears {
			gearMap[adjGear] = append(gearMap[adjGear], currNum)
		}
	}
	for _, v := range gearMap {
		if len(v) == 2 {
			ans += v[0] * v[1]
		}
	}
	return ans
}

func main() {
	lines := aoc.GetInputLinesForDay(3, false)
	p1 := solvePart1(lines)
	p2 := solvePart2(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
