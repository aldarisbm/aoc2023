package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func dayFivePartOne() {
	data := loadData("five")
	seedsString := strings.Split(strings.Split(strings.Split(data, "\n")[0], ": ")[1], " ")
	seeds := make([]int, len(seedsString))
	for i, v := range seedsString {
		n, _ := strconv.Atoi(v)
		seeds[i] = n
	}

	seedToSoil := getMapping(seeds, data, "seed-to-soil map:\n")
	soilToFertilizer := getMapping(getValues(seedToSoil), data, "soil-to-fertilizer map:\n")
	fertilizerToWater := getMapping(getValues(soilToFertilizer), data, "fertilizer-to-water map:\n")
	waterToLight := getMapping(getValues(fertilizerToWater), data, "water-to-light map:\n")
	lightToTemperature := getMapping(getValues(waterToLight), data, "light-to-temperature map:\n")
	temperatureToHumidity := getMapping(getValues(lightToTemperature), data, "temperature-to-humidity map:\n")
	humidityToLocation := getMapping(getValues(temperatureToHumidity), data, "humidity-to-location map:\n")

	minRes := slices.Min(getValues(humidityToLocation))
	fmt.Printf("Day Five Part One Result: %d\n", minRes)
}

func dayFivePartTwo() {
	_ = loadData("five")
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

func getMapping(seeds []int, data string, sep string) map[int]int {
	mappings := make(map[int]int)
	m := getMatrix(data, sep)

	for _, seed := range seeds {
		for _, row := range m {
			destStart := row[0]
			sourceStart := row[1]
			length := row[2]
			if seed >= sourceStart && seed < sourceStart+length {
				mappings[seed] = destStart + (seed - sourceStart)
			}
		}
	}

	for _, seed := range seeds {
		if _, ok := mappings[seed]; !ok {
			mappings[seed] = seed
		}
	}
	return mappings
}

func getValues(m map[int]int) []int {
	values := make([]int, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}
