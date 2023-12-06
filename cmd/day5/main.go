package main

import (
	"advent_of_code_2024/pkg/aoc"
	"bufio"
	"fmt"
	"math"
	"strings"
)

type Range struct {
	SourceStart      int
	DestinationStart int
	RangeLength      int
}

type ResourceMap struct {
	Ranges []Range
}

func skipTrailingLines(scn *bufio.Scanner) (string, bool) {
	for scn.Scan() {
		line := scn.Text()
		if strings.TrimSpace(line) != "" {
			return line, true
		}
	}
	return "", false
}

func getSeeds(scn *bufio.Scanner) []int {
	line, _ := skipTrailingLines(scn)
	seedsString := strings.Split(line, ":")[1]
	seeds := aoc.MapStrToInt(strings.Fields(seedsString))
	return seeds
}

func getResourceMap(scn *bufio.Scanner) (*ResourceMap, bool) {
	ranges := []Range{}
	_, hasMore := skipTrailingLines(scn)
	if !hasMore {
		return nil, false
	}
	for scn.Scan() {
		line := scn.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		nums := aoc.MapStrToInt(strings.Fields(line))
		ranges = append(ranges, Range{nums[1], nums[0], nums[2]})
	}
	return &ResourceMap{ranges}, true
}

func getDestinationNumber(sourceNumber int, resourceMap *ResourceMap) int {
	for _, resourceRange := range resourceMap.Ranges {
		if sourceNumber >= resourceRange.SourceStart && sourceNumber < resourceRange.SourceStart+resourceRange.RangeLength {
			requiredDestinationRange := sourceNumber - resourceRange.SourceStart
			if requiredDestinationRange <= resourceRange.RangeLength {
				return resourceRange.DestinationStart + requiredDestinationRange
			}
		}
	}
	return sourceNumber
}

func getDestinationRange(sourceNumber, sourceRange int, resourceMap *ResourceMap) (int, int) {
	fmt.Printf("sourceNumber: %d, sourceRange: %d\n", sourceNumber, sourceRange)
	for _, resourceRange := range resourceMap.Ranges {
		if (sourceNumber > resourceRange.SourceStart+resourceRange.RangeLength) || (sourceNumber+sourceRange < resourceRange.SourceStart) {
			continue
		}
		if sourceNumber >= resourceRange.SourceStart && sourceNumber+sourceRange <= resourceRange.SourceStart+resourceRange.RangeLength {
			fmt.Println("HERE MAN")
			return resourceRange.DestinationStart, sourceRange + (sourceNumber - resourceRange.SourceStart)
		}
		/*
										7 8 9
						3 4 5 6 7 8 9 10 11 12 13 14 15 16 17
			0 1 2 3 4 5												 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29
		*/

		if sourceNumber < resourceRange.SourceStart {
			return resourceRange.DestinationStart, (sourceNumber + sourceRange) - resourceRange.SourceStart
		}
		if sourceNumber+sourceRange > resourceRange.SourceStart+resourceRange.RangeLength {
			return resourceRange.DestinationStart, (resourceRange.SourceStart + resourceRange.RangeLength) - sourceNumber
		}
	}
	fmt.Println("HERE???")
	return sourceNumber, sourceRange
}

func solvePart1(scn *bufio.Scanner) int {
	resourceNumbers := getSeeds(scn)
	for resourceMap, hasMore := getResourceMap(scn); hasMore; resourceMap, hasMore = getResourceMap(scn) {
		for resourceNumberIdx, resourceNumber := range resourceNumbers {
			destinationNumber := getDestinationNumber(resourceNumber, resourceMap)
			resourceNumbers[resourceNumberIdx] = destinationNumber
		}
	}
	minResource := resourceNumbers[0]
	for _, resourceNumber := range resourceNumbers {
		if resourceNumber < minResource {
			minResource = resourceNumber
		}
	}
	return minResource
}

func getResourceMaps(scn *bufio.Scanner) []*ResourceMap {
	resourceMaps := []*ResourceMap{}
	for resourceMap, hasMore := getResourceMap(scn); hasMore; resourceMap, hasMore = getResourceMap(scn) {
		resourceMaps = append(resourceMaps, resourceMap)
	}
	return resourceMaps
}

func solvePart2BruteForce(scn *bufio.Scanner) int {
	resourceNumbersRanges := getSeeds(scn)
	resourceMaps := getResourceMaps(scn)

	minResource := math.MaxInt
	for i := 0; i < len(resourceNumbersRanges); i += 2 {
		sourceNumber := resourceNumbersRanges[i]
		sourceRange := resourceNumbersRanges[i+1]

		for resourceNumber := sourceNumber; resourceNumber < sourceNumber+sourceRange; resourceNumber++ {
			destinationNumber := resourceNumber
			for _, resourceMap := range resourceMaps {
				destinationNumber = getDestinationNumber(destinationNumber, resourceMap)
			}
			if destinationNumber < minResource {
				minResource = destinationNumber
			}
		}
	}

	return minResource
}

// tried using a more efficient approach but it didn't work
func solvePart2(scn *bufio.Scanner) int {
	resourceNumbersRanges := getSeeds(scn)
	for resourceMap, hasMore := getResourceMap(scn); hasMore; resourceMap, hasMore = getResourceMap(scn) {
		for i := 0; i < len(resourceNumbersRanges); i += 2 {
			sourceNumber := resourceNumbersRanges[i]
			sourceRange := resourceNumbersRanges[i+1]
			destinationStart, destinationRange := getDestinationRange(sourceNumber, sourceRange-1, resourceMap)
			resourceNumbersRanges[i] = destinationStart
			resourceNumbersRanges[i+1] = destinationStart + destinationRange
		}
	}
	return 0
}

func main() {
	// scn1 := aoc.GetInputScannerForDay(5, false)
	scn2 := aoc.GetInputScannerForDay(5, false)
	// p1 := solvePart1(scn1)
	p2 := solvePart2BruteForce(scn2)
	// fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
