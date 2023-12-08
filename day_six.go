package main

import (
	"fmt"
	"strings"
)

func daySixPartOne() {
	data := loadData("six")
	races := getRaces(data)

	var waysToWin []int
	for _, race := range races {
		numOfWaysToWin := race.processRace()
		waysToWin = append(waysToWin, numOfWaysToWin)
	}

	result := 1
	for _, w := range waysToWin {
		result *= w
	}
	fmt.Printf("Day Six Part One: %d\n", result)
}

func daySixPartTwo() {
	data := loadData("six_two")
	races := getRaces(data)

	var waysToWin []int
	for _, race := range races {
		numOfWaysToWin := race.processRace()
		waysToWin = append(waysToWin, numOfWaysToWin)
	}

	result := 1
	for _, w := range waysToWin {
		result *= w
	}
	fmt.Printf("Day Six Part Two: %d\n", result)
}

type Race struct {
	Record   int
	Distance int
}

func getRaces(s string) []*Race {
	firstLine := strings.Split(s, "\n")[0]
	timeStrings := strings.TrimSpace(strings.Split(firstLine, ":")[1])
	times := processNumbers(timeStrings)

	secondLine := strings.Split(s, "\n")[1]
	distanceStrings := strings.TrimSpace(strings.Split(secondLine, ":")[1])
	distances := processNumbers(distanceStrings)

	var races []*Race

	for i := 0; i < len(times); i++ {
		r := &Race{
			Record:   times[i],
			Distance: distances[i],
		}
		races = append(races, r)
	}

	return races
}

func (r *Race) processRace() int {
	wins := 0

	for hold := 1; hold < r.Record; hold++ {
		timeToCover := r.Record - hold
		distanceTraveled := hold * timeToCover

		if distanceTraveled > r.Distance {
			wins++
		}

	}

	return wins
}
