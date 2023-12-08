package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Kind int

const (
	highCard Kind = iota + 1
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func daySevenPartOne() {
	data := loadData("seven")
	games := getGames(data)
	sort.Sort(ByGame(games))
	sum := 0
	for i, g := range games {
		rank := i + 1
		sum += g.Bid * rank
	}

	fmt.Printf("Day Seven Part One: %d\n", sum)
}

type Game struct {
	Bid     int
	Hand    string
	Kind    Kind
	HandInt []int
}

type ByGame []Game

func (g ByGame) Len() int {
	return len(g)
}
func (g ByGame) Less(i, j int) bool {
	if g[i].Kind == g[j].Kind {
		for idx := 0; idx < len(g[i].HandInt); idx++ {
			if g[i].HandInt[idx] != g[j].HandInt[idx] {
				return g[i].HandInt[idx] < g[j].HandInt[idx]
			}
		}
	} else {
		return g[i].Kind < g[j].Kind
	}
	return false
}

func (g ByGame) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func handToIntSlice(hand string) []int {
	m := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	var res []int
	for _, v := range hand {
		res = append(res, m[string(v)])
	}
	return res
}

func getGames(s string) []Game {
	lines := strings.Split(s, "\n")
	var games []Game
	for _, line := range lines {
		g := strings.Split(line, " ")
		hand := g[0]
		bid, _ := strconv.Atoi(g[1])
		kind := getKind(hand)
		games = append(games, Game{
			Hand:    hand,
			Bid:     bid,
			Kind:    kind,
			HandInt: handToIntSlice(hand),
		})
	}

	return games
}

func getKind(hand string) Kind {
	if isFiveOfAKind(hand) {
		return fiveOfAKind
	} else if isFourOfAKind(hand) {
		return fourOfAKind
	} else if isFullHouse(hand) {
		return fullHouse
	} else if isThreeOfAKind(hand) {
		return threeOfAKind
	} else if isTwoPair(hand) {
		return twoPair
	} else if isOnePair(hand) {
		return onePair
	} else {
		return highCard
	}
}

func isFiveOfAKind(hand string) bool {
	if strings.Count(hand, string(hand[0])) == 5 {
		return true
	}
	return false
}

func isFourOfAKind(hand string) bool {
	if strings.Count(hand, string(hand[0])) == 4 || strings.Count(hand, string(hand[1])) == 4 {
		return true
	}
	return false
}

func isFullHouse(hand string) bool {
	a := string(hand[0])
	b := ""
	for i := 1; i < len(hand); i++ {
		if string(hand[i]) != a {
			b = string(hand[i])
			break
		}
	}
	if strings.Count(hand, a) == 3 && strings.Count(hand, b) == 2 {
		return true
	}
	if strings.Count(hand, b) == 3 && strings.Count(hand, a) == 2 {
		return true
	}
	return false
}

func isThreeOfAKind(hand string) bool {
	for _, c := range hand {
		if strings.Count(hand, string(c)) == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(hand string) bool {
	first := ""
	count := 0
	for _, c := range hand {
		if strings.Count(hand, string(c)) == 2 && string(c) != first {
			first = string(c)
			count++
			if count == 2 {
				return true
			}
		}
	}
	return false
}

func isOnePair(hand string) bool {
	for _, c := range hand {
		if strings.Count(hand, string(c)) == 2 {
			return true
		}
	}
	return false
}
