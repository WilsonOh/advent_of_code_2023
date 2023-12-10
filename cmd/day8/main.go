package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

type Node struct {
	Left  string
	Right string
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func solvePart1(input string) int {
	tokens := strings.Split(input, "\n\n")
	instructions := tokens[0]
	nodesString := strings.Split(tokens[1], "\n")
	graph := map[string]Node{}

	for _, nodeString := range nodesString {
		re := regexp.MustCompile("[A-Z0-9]{3}")
		matches := re.FindAllString(nodeString, -1)
		graph[matches[0]] = Node{Left: matches[1], Right: matches[2]}
	}

	currInstructionIdx := 0
	numInstructions := len(instructions)

	numSteps := 0

	currNode := "AAA"

	for currNode != "ZZZ" {
		numSteps++

		switch string(instructions[currInstructionIdx]) {
		case "L":
			currNode = graph[currNode].Left
		case "R":
			currNode = graph[currNode].Right
		}
		currInstructionIdx = (currInstructionIdx + 1) % numInstructions
	}
	return numSteps

}

func solvePart2(input string) int {
	tokens := strings.Split(input, "\n\n")
	instructions := tokens[0]
	nodesString := strings.Split(tokens[1], "\n")
	graph := map[string]Node{}

	currNodes := []string{}
	for _, nodeString := range nodesString {
		re := regexp.MustCompile("[A-Z0-9]{3}")
		matches := re.FindAllString(nodeString, -1)
		graph[matches[0]] = Node{Left: matches[1], Right: matches[2]}
		if matches[0][2] == 'A' {
			currNodes = append(currNodes, matches[0])
		}
	}

	currInstructionIdx := 0
	numInstructions := len(instructions)

	numSteps := 0

	timings := map[string]int{}

	for {
		numSteps++

		for idx, node := range currNodes {
			switch string(instructions[currInstructionIdx]) {
			case "L":
				currNodes[idx] = graph[node].Left
			case "R":
				currNodes[idx] = graph[node].Right
			}
			if node[2] == 'Z' {
				timings[node] = numSteps - 1
			}
			if len(timings) == len(currNodes) {
				timingsVals := funk.Values(timings).([]int)
				return LCM(timingsVals[0], timingsVals[1], timingsVals[2:]...)
			}
		}
		currInstructionIdx = (currInstructionIdx + 1) % numInstructions
	}

}

func main() {
	input := aoc.GetInputForDay(8, false)
	p1 := solvePart1(input)
	p2 := solvePart2(input)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
