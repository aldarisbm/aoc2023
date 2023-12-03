package main

import (
	"fmt"
	"strconv"
	"strings"
)

const nums = "1234567890"

func dayThreePartOne() {
	data := loadData("three")
	sum := getSum(data)

	fmt.Printf("Day Three Part One Result: %d\n", sum)
}

func dayThreePartTwo() {
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

func getSum(data string) int {
	smi := NewMatrixSymbolIndex(140, 140, data)
	sum := 0

	for i, l := range strings.Split(data, "\n") {
		for j := 0; j < len(l); j++ {
			if strings.Contains(nums, string(l[j])) {
				num := processNumber(l[j:])
				lenOfNum := len(strconv.Itoa(num))
				if isNumSurroundedBySymbol(i, j, num, smi) {
					fmt.Printf("found number surrounded: %d, at: %d, %d\n", num, i, j)
					sum += num
				}
				j = j + lenOfNum // starting j at the end of the number
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
	if num == 712 {
		fmt.Println()
	}
	lenOfNum := len(strconv.Itoa(num))
	for currIdx := j; currIdx < j+lenOfNum; currIdx++ {
		if isCharSurroundedBySymbol(i, currIdx, symbolMatrix) {
			return true
		}
	}
	return false
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
