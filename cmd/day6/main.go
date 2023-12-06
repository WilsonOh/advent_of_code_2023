package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"strings"
)

type Race struct {
	AllowedTime    int
	RecordDistance int
}

func parseRaces(lines []string) []Race {
	times := aoc.MapStrToInt(strings.Fields(lines[0])[1:])
	dists := aoc.MapStrToInt(strings.Fields(lines[1])[1:])
	races := []Race{}
	for i := 0; i < len(times); i++ {
		races = append(races, Race{AllowedTime: times[i], RecordDistance: dists[i]})
	}
	return races
}

func parseRace(lines []string) Race {
	times := strings.Fields(lines[0])[1:]
	dists := strings.Fields(lines[1])[1:]
	time := strings.Join(times, "")
	dist := strings.Join(dists, "")
	return Race{AllowedTime: aoc.StrToInt(time), RecordDistance: aoc.StrToInt(dist)}
}

func solve(races []Race) int {
	ans := 1
	for _, race := range races {
		numWaysToBeat := 0
		for i := 1; i < race.AllowedTime; i++ {
			distTravelled := i * (race.AllowedTime - i)
			if distTravelled > race.RecordDistance {
				numWaysToBeat++
			}
		}
		ans *= numWaysToBeat
	}
	return ans

}

func main() {
	lines := aoc.GetInputLinesForDay(6, false)
	p1 := solve(parseRaces(lines))
	p2 := solve([]Race{parseRace(lines)})
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
