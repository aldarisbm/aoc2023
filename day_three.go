package main

import (
	"fmt"
	"strconv"
	"strings"
)

const nums = "1234567890"

func dayThreePartOne() {
	data := loadData("three")
	sum := getSumDayOne(data)

	fmt.Printf("Day Three Part One Result: %d\n", sum)
}

func dayThreePartTwo() {
	data := loadData("three")
	sum := getSumDayTwo(data)
	fmt.Printf("Day Three Part Two Result: %d\n", sum)
}

type MatrixIndex struct {
	Data [][]int
}

func NewMatrixSymbolIndex(height, width int, data string) *MatrixIndex {
	mi := MatrixIndex{}
	mi.Data = make([][]int, height)
	for i := 0; i < height; i++ {
		mi.Data[i] = make([]int, width)
	}

	dot := "."
	for i, l := range strings.Split(data, "\n") {
		for j, char := range l {
			if !strings.Contains(nums, string(char)) && string(char) != dot {
				mi.Data[i][j] = 1
			}
		}
	}

	return &mi
}

func getSumDayOne(data string) int {
	smi := NewMatrixSymbolIndex(140, 140, data)
	sum := 0

	for i, l := range strings.Split(data, "\n") {
		for j := 0; j < len(l); {
			if strings.Contains(nums, string(l[j])) {
				num := processNumber(l[j:])
				lenOfNum := len(strconv.Itoa(num))
				if isNumSurroundedBySymbol(i, j, num, smi) {
					sum += num
				}
				j = j + lenOfNum // starting j at the end of the number
			} else {
				j++
			}
		}
	}
	return sum
}

func getSumDayTwo(data string) int {

	intID := make(map[int]int)
	id := 0

	sum := 0
	m := make([][]map[string]int, 140)
	for i := 0; i < 140; i++ {
		m[i] = make([]map[string]int, 140)
	}

	mn := make([][]string, 140)
	for i := 0; i < 140; i++ {
		mn[i] = make([]string, 140)
	}
	for idx, l := range strings.Split(data, "\n") {
		for j, c := range l {
			mn[idx][j] = string(c)
		}
	}

	for idx, l := range strings.Split(data, "\n") {
		for j := 0; j < len(l); {
			if strings.Contains(nums, string(l[j])) {
				num := processNumber(l[j:])
				lenOfNum := len(strconv.Itoa(num))
				for x := j; x < j+lenOfNum; x++ {
					if m[idx][x] == nil {
						m[idx][x] = make(map[string]int)
					}
					m[idx][x][string(l[x])] = id
				}
				intID[id] = num
				id++
				j += lenOfNum
			} else {
				if m[idx][j] == nil {
					m[idx][j] = make(map[string]int)
				}
				m[idx][j] = map[string]int{string(l[j]): -1}
				j++
			}
		}
	}

	for i, l := range strings.Split(data, "\n") {
		for j := 0; j < len(l); j++ {
			if string(l[j]) == "*" {
				gearRatio := getSurroundingNumbers(i, j, m, mn, intID)
				if gearRatio != -1 {
					sum += gearRatio
				}
			}
		}
	}
	return sum
}

func processNumber(s string) int {
	for idx, c := range s {
		if !strings.Contains(nums, string(c)) {
			parsed, _ := strconv.ParseInt(s[:idx], 10, 64)
			n := int(parsed)
			return n
		}
		if idx == len(s)-1 {
			parsed, _ := strconv.ParseInt(s, 10, 64)
			n := int(parsed)
			return n
		}
	}
	return -1
}

func isNumSurroundedBySymbol(i, j, num int, symbolMatrix *MatrixIndex) bool {
	lenOfNum := len(strconv.Itoa(num))
	for currIdx := j; currIdx < j+lenOfNum; currIdx++ {
		if isCharSurroundedBySymbol(i, currIdx, symbolMatrix) {
			return true
		}
	}
	return false
}

func getSurroundingNumbers(i, j int, m [][]map[string]int, mn [][]string, ids map[int]int) int {
	found := make(map[int]bool)
	var numbers []int

	// check left
	if j != 0 {
		s := mn[i][j-1]
		_, ok := found[m[i][j-1][s]]
		if m[i][j-1][s] != -1 && !ok {
			found[m[i][j-1][s]] = true
			numbers = append(numbers, ids[m[i][j-1][s]])
		}
	}
	// check right
	if j != 139 {
		s := mn[i][j+1]
		_, ok := found[m[i][j+1][s]]
		if m[i][j+1][s] != -1 && !ok {
			found[m[i][j+1][s]] = true
			numbers = append(numbers, ids[m[i][j+1][s]])
		}
	}
	// check above
	if i != 0 {
		s := mn[i-1][j]
		_, ok := found[m[i-1][j][s]]
		if m[i-1][j][s] != -1 && !ok {
			found[m[i-1][j][s]] = true
			numbers = append(numbers, ids[m[i-1][j][s]])
		}
	}
	// check below
	if i != 139 {
		s := mn[i+1][j]
		_, ok := found[m[i+1][j][s]]
		if m[i+1][j][s] != -1 && !ok {
			found[m[i+1][j][s]] = true
			numbers = append(numbers, ids[m[i+1][j][s]])
		}
	}
	// check top left
	if i != 0 && j != 0 {
		s := mn[i-1][j-1]
		_, ok := found[m[i-1][j-1][s]]
		if m[i-1][j-1][s] != -1 && !ok {
			found[m[i-1][j-1][s]] = true
			numbers = append(numbers, ids[m[i-1][j-1][s]])
		}
	}
	// check top right
	if i != 0 && j != 139 {
		s := mn[i-1][j+1]
		_, ok := found[m[i-1][j+1][s]]
		if m[i-1][j+1][s] != -1 && !ok {
			found[m[i-1][j+1][s]] = true
			numbers = append(numbers, ids[m[i-1][j+1][s]])
		}
	}
	// check bottom left
	if i != 139 && j != 0 {
		s := mn[i+1][j-1]
		_, ok := found[m[i+1][j-1][s]]
		if m[i+1][j-1][s] != -1 && !ok {
			found[m[i+1][j-1][s]] = true
			numbers = append(numbers, ids[m[i+1][j-1][s]])
		}
	}
	// check bottom right
	if i != 139 && j != 139 {
		s := mn[i+1][j+1]
		_, ok := found[m[i+1][j+1][s]]
		if m[i+1][j+1][s] != -1 && !ok {
			found[m[i+1][j+1][s]] = true
			numbers = append(numbers, ids[m[i+1][j+1][s]])
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}
	return -1
}

func isCharSurroundedBySymbol(i, j int, symbolMatrix *MatrixIndex) bool {
	// check left
	if j != 0 {
		if symbolMatrix.Data[i][j-1] == 1 {
			return true
		}
	}
	// check right
	if j != 139 {
		if symbolMatrix.Data[i][j+1] == 1 {
			return true
		}
	}
	//check above
	if i != 0 {
		if symbolMatrix.Data[i-1][j] == 1 {
			return true
		}
	}
	// check below
	if i != 139 {
		if symbolMatrix.Data[i+1][j] == 1 {
			return true
		}
	}
	// check top left
	if i != 0 && j != 0 {
		if symbolMatrix.Data[i-1][j-1] == 1 {
			return true
		}
	}
	// check top right
	if i != 0 && j != 139 {
		if symbolMatrix.Data[i-1][j+1] == 1 {
			return true
		}
	}
	// check bottom left
	if i != 139 && j != 0 {
		if symbolMatrix.Data[i+1][j-1] == 1 {
			return true
		}
	}
	// check bottom right
	if i != 139 && j != 139 {
		if symbolMatrix.Data[i+1][j+1] == 1 {
			return true
		}
	}

	return false
}
