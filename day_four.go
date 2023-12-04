package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	cardNumber     int
	winningNumbers []int
	currentNumbers []int
	matches        int
}

func dayFourPartOne() {
	data := loadData("four")
	sum := 0
	for i, l := range strings.Split(data, "\n") {
		s := strings.Split(l, fmt.Sprintf("%d: ", i+1))
		s = strings.Split(s[1], " | ")
		winningNumbers := processNumbers(s[0])
		myNumbers := processNumbers(s[1])
		sum += getPoints(winningNumbers, myNumbers)
	}
	fmt.Printf("Day Four Part One: %d\n", sum)
}

func dayFourPartTwo() {
	data := loadData("four")
	var allCards []Card
	for i, l := range strings.Split(data, "\n") {
		s := strings.Split(l, fmt.Sprintf("%d: ", i+1))
		s = strings.Split(s[1], " | ")
		winningNumbers := processNumbers(s[0])
		myNumbers := processNumbers(s[1])
		matches := getRealPoints(winningNumbers, myNumbers)
		allCards = append(allCards, Card{i + 1, winningNumbers, myNumbers, matches})
	}
	fmt.Printf("%+v\n", allCards)
}

func getRealPoints(w, m []int) int {
	points := 0
	for _, v := range m {
		if intInSlice(v, w) {
			points *= 2
		}
	}
	return points
}

func processNumbers(s string) []int {
	var nums []int
	digits := "0123456789"
	for i := 0; i < len(s); {
		if strings.Contains(digits, string(s[i])) {
			num := processNumber(s[i:])
			nums = append(nums, num)
			lenOfNum := len(strconv.Itoa(num))
			i += lenOfNum
		} else {
			i++
		}

	}
	return nums
}

func getPoints(w, m []int) int {
	points := 0
	for _, v := range m {
		if intInSlice(v, w) && points != 0 {
			points *= 2
		}
		if intInSlice(v, w) && points == 0 {
			points = 1
		}
	}
	return points
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
