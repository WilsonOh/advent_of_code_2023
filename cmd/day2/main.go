package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"strconv"
	"strings"
)

type CubeSet struct {
	NumRed   int
	NumGreen int
	NumBlue  int
}

func (this *CubeSet) String() string {
	return fmt.Sprintf("CubeSet{ NumRed: %d, NumGreen: %d, NumBlue: %d }", this.NumRed, this.NumGreen, this.NumBlue)
}

type Game struct {
	id       int
	CubeSets []*CubeSet
}

func (this *Game) String() string {
	return fmt.Sprintf("Game{ id: %d, CubeSets: %v }", this.id, this.CubeSets)
}

func parseStringToGame(s string) *Game {
	game := new(Game)
	tokens := strings.Split(s, ": ")
	idString := strings.Split(tokens[0], " ")[1]
	id, _ := strconv.Atoi(idString)
	game.id = id
	cubeSetsString := strings.Split(tokens[1], "; ")
	cubeSets := make([]*CubeSet, len(cubeSetsString))
	for idx, cubeSetString := range cubeSetsString {
		cubeSet := new(CubeSet)
		cubes := strings.Split(cubeSetString, ", ")
		for _, cube := range cubes {
			cubeTokens := strings.Split(cube, " ")
			numCubesString := cubeTokens[0]
			numCubes, _ := strconv.Atoi(numCubesString)
			color := cubeTokens[1]
			switch color {
			case "red":
				cubeSet.NumRed = numCubes
			case "blue":
				cubeSet.NumBlue = numCubes
			case "green":
				cubeSet.NumGreen = numCubes
			}
		}
		cubeSets[idx] = cubeSet
	}
	game.CubeSets = cubeSets
	return game
}

func isValidCubeSet(cubeSet *CubeSet, numRedAllowed, numGreenAllowed, numBlueAllowed int) bool {
	return cubeSet.NumRed <= numRedAllowed && cubeSet.NumGreen <= numGreenAllowed && cubeSet.NumBlue <= numBlueAllowed
}

func solvePart1(lines []string) int {
	ans := 0
	for _, line := range lines {
		game := parseStringToGame(line)
		isValidGame := true
		for _, cubeSet := range game.CubeSets {
			if !isValidCubeSet(cubeSet, 12, 13, 14) {
				isValidGame = false
				break
			}
		}
		if isValidGame {
			ans += game.id
		}
	}
	return ans
}

func solvePart2(lines []string) int {
	ans := 0

	for _, line := range lines {
		game := parseStringToGame(line)
		var maxRed, maxGreen, maxBlue int
		for _, cubeSet := range game.CubeSets {
			maxRed = max(maxRed, cubeSet.NumRed)
			maxGreen = max(maxGreen, cubeSet.NumGreen)
			maxBlue = max(maxBlue, cubeSet.NumBlue)
		}
		ans += maxRed * maxBlue * maxGreen
	}
	return ans

}

func main() {
	lines := aoc.GetInputLinesForDay(2, false)
	p1 := solvePart1(lines)
	p2 := solvePart2(lines)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
}
