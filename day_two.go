package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTwoPartOne() {
	data := loadData("two")
	splitData := strings.Split(data, "\n")
	sum := 0
	for idx, line := range splitData {
		gameNumber := idx + 1
		if isGamePossible(line) {
			sum += gameNumber
		}
	}
	fmt.Printf("Day Two Part One Result: %d\n", sum)
}

func dayTwoPartTwo() {
	data := loadData("two")
	splitData := strings.Split(data, "\n")
	sum := 0
	for _, line := range splitData {
		or, og, ob := 0, 0, 0
		fullGame := strings.Split(line, ": ")[1]
		roundsInGame := strings.Split(fullGame, "; ")
		for _, game := range roundsInGame {
			gamePull := strings.Split(game, ", ")
			r, g, b := getCubeCount(gamePull)
			if r > or {
				or = r
			}
			if g > og {
				og = g
			}
			if b > ob {
				ob = b
			}
		}
		sum += or * og * ob
	}
	fmt.Printf("Day Two Part Two Result: %d\n", sum)

}

func isGamePossible(line string) bool {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	fullGame := strings.Split(line, ": ")[1]
	roundsInGame := strings.Split(fullGame, "; ")
	for _, game := range roundsInGame {
		gamePull := strings.Split(game, ", ")
		r, g, b := getCubeCount(gamePull)
		if r > maxRed || g > maxGreen || b > maxBlue {
			return false
		}
	}
	return true
}

func getCubeCount(pull []string) (int, int, int) {
	r, g, b := 0, 0, 0
	for _, cube := range pull {
		splitCubeRoll := strings.Split(cube, " ")
		color := splitCubeRoll[1]
		count := splitCubeRoll[0]
		switch color {
		case "red":
			n, _ := strconv.ParseInt(count, 10, 64)
			r = int(n)
		case "green":
			n, _ := strconv.ParseInt(count, 10, 64)
			g = int(n)
		case "blue":
			n, _ := strconv.ParseInt(count, 10, 64)
			b = int(n)
		}

	}
	return r, g, b
}
