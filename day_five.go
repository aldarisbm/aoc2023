package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func dayFivePartOne() {
	data := loadData("five")
	seeds := strings.Split(strings.Split(strings.Split(data, "\n")[0], ": ")[1], " ")
	seedToSoil := getMapping(seeds, data, "seed-to-soil map:\n")
	fmt.Println(seedToSoil)
	//soilToFertilizer := getMapping(seeds, data, "soil-to-fertilizer map:\n")
	//fertilizerToWater := getMapping(seeds, data, "fertilizer-to-water map:\n")
	//waterToLight := getMapping(seeds, data, "water-to-light map:\n")
	//lightToTemperature := getMapping(seeds, data, "light-to-temperature map:\n")
	//temperatureToHumidity := getMapping(seeds, data, "temperature-to-humidity map:\n")
	//humidityToLocation :=  getMapping(seeds, data, "humidity-to-location map:\n")
}

func getMatrix(d, sep string) [][]int {
	matrix := make([][]int, 0)
	s := strings.Split(strings.Split(d, fmt.Sprintf("%s", sep))[1], "\n\n")[0]
	for _, line := range strings.Split(s, "\n") {
		lineLen := len(strings.Split(line, " "))
		row := make([]int, lineLen)
		for i, v := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(v)
			row[i] = n
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func getMapping(seeds []string, data string, sep string) *sync.Map {
	s := make([]int, len(seeds))
	for i, v := range seeds {
		n, _ := strconv.Atoi(v)
		s[i] = n
	}
	m := getMatrix(data, sep)

	mapping := &sync.Map{}

	for _, seed := range s {
		for _, r := range m {
			destStart := r[0]
			rangeLen := r[2]
			if seed-destStart > 0 && seed-destStart < rangeLen {
				mapping.Store(seed, seed+destStart)
			}
		}
	}

	mapping.Range(func(key, value any) bool {
		fmt.Printf("k %d, v %d", key, value)
		return true
	})
	//50 98 2
	// dest start 50
	// source start 98
	// range length 2
	// 98 99 == 50 51
	// 97 - 98 = -1
	// num - range start > 0 < rang length = good

	//52 50 48
	// seed
	//The first line has a destination range start of 50, a source range start of 98, and a range length of 2.
	//This line means that the source range starts at 98 and contains two values: 98 and 99.
	//The destination range is the same length, but it starts at 50, so its two values are 50 and 51.
	//With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.
	for _, row := range m {
		destinationStart := row[0]
		sourceStart := row[1]
		rangeLen := row[2]
		for i := 0; i < rangeLen; i++ {
			if intInSlice(sourceStart, s) {
				mapping.Store(sourceStart+1, destinationStart+1)
			}
		}
	}
	return mapping
}
