package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
)

type Coord struct {
	Row int
	Col int
}

// parseGrid build an adjacency list from the grid provided
// edges are constructed based on the "rules" given for each type of pipe
func parseGrid(grid []string) (map[Coord][]Coord, Coord) {
	pipeDirs := map[byte][]Coord{
		'|': {{1, 0}, {-1, 0}},
		'-': {{0, -1}, {0, 1}},
		'L': {{-1, 0}, {0, 1}},
		'J': {{0, -1}, {-1, 0}},
		'7': {{0, -1}, {1, 0}},
		'F': {{1, 0}, {0, 1}},
		'S': {{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
	}

	adjList := map[Coord][]Coord{}
	startingPos := Coord{}

	for rowIdx, line := range grid {
		for colIdx := 0; colIdx < len(line); colIdx++ {
			col := line[colIdx]
			if col == '.' {
				continue
			}
			if col == 'S' {
				startingPos.Row = rowIdx
				startingPos.Col = colIdx
			}
			currCoord := Coord{Row: rowIdx, Col: colIdx}
			for _, dir := range pipeDirs[col] {
				nr := rowIdx + dir.Row
				nc := colIdx + dir.Col
				if (nr >= 0 && nr < len(grid)) && (nc >= 0 && nc < len(grid[0])) {
					if grid[nr][nc] != '.' {
						adjList[currCoord] = append(adjList[currCoord], Coord{Row: nr, Col: nc})
					}
				}
			}
		}
	}
	return adjList, startingPos
}

// solvePart1 do a bfs starting from the 'S' position, marking the number of steps from 'S' for each node.
// in the end, return the number of nodes visited in the loop divided by 2
func solvePart1(lines []string) int {
	adjList, startingPos := parseGrid(lines)
	dists := map[Coord]int{}
	dists[startingPos] = 0

	q := []Coord{startingPos}
	visitedSet := map[Coord]bool{}

	for len(q) > 0 {
		currNode := q[0]
		q = q[1:]
		for _, nei := range adjList[currNode] {
			newDist := dists[currNode] + 1
			_, distSet := dists[nei]
			if !distSet || newDist < dists[nei] {
				dists[nei] = newDist
			}
			_, isVisited := visitedSet[nei]
			if !isVisited {
				visitedSet[nei] = true
				q = append(q, nei)
			}
		}
	}

	return len(visitedSet) / 2
}

func main() {
	lines := aoc.GetInputLinesForDay(10, false)
	p1 := solvePart1(lines)
	fmt.Println("Part 1:", p1)
}
