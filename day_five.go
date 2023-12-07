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
	soilToFertilizer := getMapping(seedToSoil, data, "soil-to-fertilizer map:\n")
	fertilizerToWater := getMapping(soilToFertilizer, data, "fertilizer-to-water map:\n")
	waterToLight := getMapping(fertilizerToWater, data, "water-to-light map:\n")
	lightToTemperature := getMapping(waterToLight, data, "light-to-temperature map:\n")
	temperatureToHumidity := getMapping(lightToTemperature, data, "temperature-to-humidity map:\n")
	humidityToLocation := getMapping(temperatureToHumidity, data, "humidity-to-location map:\n")

	minRes := slices.Min(humidityToLocation)
	fmt.Printf("Day Five Part One Result: %d\n", minRes)
}

func dayFivePartTwo() {
	data := loadData("five")
	seedsString := strings.Split(strings.Split(strings.Split(data, "\n")[0], ": ")[1], " ")
	seeds := make([]int, len(seedsString))
	for i, v := range seedsString {
		n, _ := strconv.Atoi(v)
		seeds[i] = n
	}
	realSeeds := getRealSeeds(seeds)

	seedToSoil := getMapping(realSeeds, data, "seed-to-soil map:\n")
	realSeeds = nil
	soilToFertilizer := getMapping(seedToSoil, data, "soil-to-fertilizer map:\n")
	seedToSoil = nil
	fertilizerToWater := getMapping(soilToFertilizer, data, "fertilizer-to-water map:\n")
	soilToFertilizer = nil
	waterToLight := getMapping(fertilizerToWater, data, "water-to-light map:\n")
	fertilizerToWater = nil
	lightToTemperature := getMapping(waterToLight, data, "light-to-temperature map:\n")
	waterToLight = nil
	temperatureToHumidity := getMapping(lightToTemperature, data, "temperature-to-humidity map:\n")
	lightToTemperature = nil
	humidityToLocation := getMapping(temperatureToHumidity, data, "humidity-to-location map:\n")
	temperatureToHumidity = nil

	minRes := slices.Min(humidityToLocation)
	fmt.Printf("Day Five Part Two Result: %d\n", minRes)
}

func getRealSeeds(seeds []int) []int {
	realSeeds := make([]int, 0)
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := seeds[i] + seeds[i+1]
		for j := start; j < end; j++ {
			realSeeds = append(realSeeds, j)
		}
	}
	return realSeeds
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

func getMapping(seeds []int, data string, sep string) []int {
	m := getMatrix(data, sep)
	listOfMappedSeeds := make([]int, len(seeds))

	for i, seed := range seeds {
		listOfMappedSeeds[i] = -1
		for _, row := range m {
			destStart := row[0]
			sourceStart := row[1]
			length := row[2]
			if seed >= sourceStart && seed < sourceStart+length {
				listOfMappedSeeds[i] = destStart + (seed - sourceStart)
				break
			}
		}
		if listOfMappedSeeds[i] == -1 {
			listOfMappedSeeds[i] = seed
		}
	}

	return listOfMappedSeeds
}
