package main

import (
	"advent_of_code_2024/pkg/aoc"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var orderMapPart1 map[byte]int = map[byte]int{'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0}

var orderMapPart2 map[byte]int = map[byte]int{'A': 12, 'K': 11, 'Q': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1, 'J': 0}

func getCardCounts(cards [5]byte) map[byte]int {
	return funk.Reduce(cards, func(acc map[byte]int, curr byte) map[byte]int {
		_, ok := acc[curr]
		if !ok {
			acc[curr] = 0
		}
		acc[curr]++
		return acc
	}, map[byte]int{}).(map[byte]int)
}

func isFiveOfAKind(counts map[byte]int) bool {
	for _, count := range counts {
		if count == 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(counts map[byte]int) bool {
	for _, count := range counts {
		if count == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(counts map[byte]int) bool {
	hasThree := false
	hasTwo := false
	for _, count := range counts {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func isThreeOfAKind(counts map[byte]int) bool {
	hasThree := false
	hasTwo := false
	for _, count := range counts {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && !hasTwo
}

func isTwoPair(counts map[byte]int) bool {
	numPairs := 0
	for _, count := range counts {
		if count == 2 {
			numPairs++
		}
	}
	return numPairs == 2
}

func isOnePair(counts map[byte]int) bool {
	numOnes := 0
	numPairs := 0
	for _, count := range counts {
		if count == 2 {
			numPairs++
		}
		if count == 1 {
			numOnes++
		}
	}
	return numOnes == 3 && numPairs == 1
}

func isHighCard(counts map[byte]int) bool {
	for _, count := range counts {
		if count != 1 {
			return false
		}
	}
	return true
}

func getHandType(hand [5]byte, counts map[byte]int) (int, error) {
	m := map[int]func(counts map[byte]int) bool{
		FIVE_OF_A_KIND:  isFiveOfAKind,
		FOUR_OF_A_KIND:  isFourOfAKind,
		FULL_HOUSE:      isFullHouse,
		THREE_OF_A_KIND: isThreeOfAKind,
		TWO_PAIR:        isTwoPair,
		ONE_PAIR:        isOnePair,
		HIGH_CARD:       isHighCard,
	}
	for handType, validator := range m {
		if validator(counts) {
			return handType, nil
		}
	}
	return -1, errors.New("Unknown hand type")
}

func parseHands(lines []string) map[[5]byte]int {
	hands := map[[5]byte]int{}
	for _, line := range lines {
		tokens := strings.Fields(line)
		hand := [5]byte{}
		copy(hand[:], tokens[0])
		bidString := tokens[1]
		bid, _ := strconv.Atoi(bidString)
		hands[hand] = bid
	}
	return hands
}

func mostCommonCard(counts map[byte]int, hand [5]byte) byte {
	maxCount := 0
	var maxCard byte
	for _, card := range hand {
		if card == 'J' {
			continue
		}
		if counts[card] > maxCount {
			maxCount = counts[card]
			maxCard = card
		}
	}
	return maxCard
}

func replaceJokers(counts map[byte]int, hand [5]byte) {
	mostCommon := mostCommonCard(counts, hand)
	if mostCommon == 'J' {
		return
	}
	numJokers, ok := counts['J']
	if !ok {
		return
	}

	counts[mostCommon] += numJokers
	delete(counts, 'J')
}

func sortHandsPart1(hands [][5]byte) {
	slices.SortFunc(hands, func(a, b [5]byte) int {
		aCounts := getCardCounts(a)
		aType, _ := getHandType(a, aCounts)
		bCounts := getCardCounts(b)
		bType, _ := getHandType(b, bCounts)
		if aType != bType {
			return aType - bType
		}
		for i := 0; i < 5; i++ {
			aRank := orderMapPart1[a[i]]
			bRank := orderMapPart1[b[i]]
			if aRank != bRank {
				return aRank - bRank
			}
		}
		return 0
	})
}

func sortHandsPart2(hands [][5]byte) {
	slices.SortFunc(hands, func(a, b [5]byte) int {
		aCounts := getCardCounts(a)
		replaceJokers(aCounts, a)
		aType, _ := getHandType(a, aCounts)
		bCounts := getCardCounts(b)
		replaceJokers(bCounts, b)
		bType, _ := getHandType(b, bCounts)
		if aType != bType {
			return aType - bType
		}
		for i := 0; i < 5; i++ {
			aRank := orderMapPart2[a[i]]
			bRank := orderMapPart2[b[i]]
			if aRank != bRank {
				return aRank - bRank
			}
		}
		return 0
	})
}

func solvePart1(lines []string) int {
	handBidMap := parseHands(lines)
	hands := funk.Keys(handBidMap).([][5]byte)
	sortHandsPart1(hands)
	ans := 0
	for rank, hand := range hands {
		ans += (rank + 1) * handBidMap[hand]
	}
	return ans
}

func solvePart2(lines []string) int {
	handBidMap := parseHands(lines)
	hands := funk.Keys(handBidMap).([][5]byte)
	sortHandsPart2(hands)
	ans := 0
	for rank, hand := range hands {
		ans += (rank + 1) * handBidMap[hand]
	}
	return ans
}

func main() {
	lines := aoc.GetInputLinesForDay(7, false)
	p1 := solvePart1(lines)
	p2 := solvePart2(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
