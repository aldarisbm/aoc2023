package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOnePartOne() {
	data := loadData("one")
	splitData := strings.Split(data, "\n")
	var sum int64
	for _, line := range splitData {
		var digits []string
		for _, r := range line {
			_, err := strconv.ParseInt(string(r), 10, 64)
			if err != nil {
				continue
			} else {
				digits = append(digits, string(r))
			}
		}
		first := digits[0]
		last := digits[len(digits)-1]
		n, _ := strconv.ParseInt(fmt.Sprintf("%s%s", first, last), 10, 64)
		sum += n
	}
	fmt.Printf("Day One Part One Result: %d\n", sum)
}

func dayOnePartTwo() {
	data := loadData("one")
	splitData := strings.Split(data, "\n")
	var sum int64
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for _, line := range splitData {
		indices := make([]int, len(line))

		for idx, char := range line {
			n, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				continue
			}
			indices[idx] = int(n)
		}

		for k, v := range digits {
			for i := 0; i < strings.Count(line, k); i++ {
				replacement := strings.Repeat(" ", len(k))
				parsedLine := strings.Replace(line, k, replacement, i)
				idxOfK := strings.Index(parsedLine, k)
				indices[idxOfK] = v
			}
		}
		numString := ""
		for i := 0; i < len(indices); i++ {
			if indices[i] != 0 {
				numString += strconv.Itoa(indices[i])
				break
			}

		}
		for i := len(indices) - 1; i >= 0; i-- {
			if indices[i] != 0 {
				numString += strconv.Itoa(indices[i])
				break
			}
		}
		n, _ := strconv.ParseInt(numString, 10, 64)
		sum += n
	}
	fmt.Printf("Day One Part Two Result: %d\n", sum)
}
