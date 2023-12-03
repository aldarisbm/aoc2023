package main

import (
	"fmt"
	"os"
)

func loadData(day string) string {
	fileToOpen := fmt.Sprintf("data/%s", day)
	f, err := os.ReadFile(fileToOpen)
	if err != nil {
		panic(err)
	}
	return string(f)
}
