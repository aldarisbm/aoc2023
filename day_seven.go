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
	games := getGames(data, false)
	sort.Sort(ByGame(games))
	sum := 0
	for i, g := range games {
		rank := i + 1
		sum += g.Bid * rank
	}

	// QC
	if sum != 249638405 {
		panic("wrong sum")
	}
	fmt.Printf("Day Seven Part One: %d\n", sum)
}

func daySevenPartTwo() {
	data := loadData("seven")
	games := getGames(data, true)
	sort.Sort(ByGame(games))

	fmt.Printf("Day Seven Part Two: %d\n", 0)
}

type Game struct {
	Bid     int
	Kind    Kind
	HandInt []int
}

func NewGame(bid int, hand string, specialJoker bool) Game {
	ogHand := hand
	if specialJoker {
		hand = processHand(hand)
	}
	kind := getKind(hand)

	return Game{
		Bid:     bid,
		Kind:    kind,
		HandInt: handToIntSlice(ogHand, specialJoker),
	}
}

func processHand(hand string) string {
	k := getKind(hand)
	if k == fiveOfAKind {
		return hand
	}
	if k == fourOfAKind && strings.Count(hand, "J") == 1 {
		for _, v := range hand {
			if string(v) != "J" {
				return strings.Replace(hand, "J", string(v), 1)
			}
		}
	}
	if k == threeOfAKind && strings.Count(hand, "J") >= 1 {
		for _, v := range hand {
			if string(v) != "J" && strings.Count(hand, string(v)) == 3 {
				return strings.Replace(hand, "J", string(v), strings.Count(hand, "J"))
			}
		}
	}
	if k == twoPair && strings.Count(hand, "J") >= 1 {
		for _, v := range hand {
			if strings.Count(hand, "J") == 2 {
				for _, v := range hand {
					if string(v) != "J" && strings.Count(hand, string(v)) == 2 {
						return strings.Replace(hand, "J", string(v), strings.Count(hand, "J"))
					}
				}
			}
			if string(v) != "J" && strings.Count(hand, string(v)) == 2 {
				strongest := getStrongest(hand)
				return strings.Replace(hand, "J", strongest, -1)
			}
		}
	}
	if k == onePair && strings.Count(hand, "J") >= 1 {
		if strings.Count(hand, "J") == 2 {
			fmt.Println()
			for _, v := range hand {
				if string(v) != "J" {
					return strings.Replace(hand, "J", string(v), 1)
				}
			}
		}
		moreOf := ""
		for _, c := range hand {
			if strings.Count(hand, string(c)) == 2 {
				moreOf = string(c)
				fmt.Printf("more of %s\n", moreOf)
				break
			}
		}
		return hand
	}
	return hand
}

func getStrongest(s string) string {
	m := getMap(false)
	first := ""
	second := ""

	for _, v := range s {
		if string(v) != "J" && strings.Count(s, string(v)) == 2 {
			if first == "" && string(v) != first {
				first = string(v)
			}
			if first != "" && string(v) != first {
				second = string(v)
				break
			}
		}
	}

	if m[first] > m[second] {
		return first
	} else {
		return second
	}
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

func getMap(specialJoker bool) map[string]int {
	m := make(map[string]int)
	s := "23456789TJQKA"
	if specialJoker {
		s = "J23456789TQK"
	}
	for i := 0; i < len(s); i++ {
		m[string(s[i])] = i + 2
	}
	return m
}

func handToIntSlice(hand string, specialJoker bool) []int {
	m := getMap(specialJoker)

	var res []int
	for _, v := range hand {
		res = append(res, m[string(v)])
	}
	return res
}

func getGames(s string, specialJoker bool) []Game {
	lines := strings.Split(s, "\n")
	var games []Game
	for _, line := range lines {
		g := strings.Split(line, " ")
		hand := g[0]
		bid, _ := strconv.Atoi(g[1])
		game := NewGame(bid, hand, specialJoker)
		games = append(games, game)
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
