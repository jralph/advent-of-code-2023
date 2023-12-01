package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func extractCalibrationValue(line string) int {
	// Find the first and last digits in the line
	var first, last int

	for _, char := range line {
		if char >= '0' && char <= '9' {
			if first == 0 {
				first = int(char - '0')
			}
			last = int(char - '0')
		}
	}

	calibrationValue := first*10 + last
	return calibrationValue
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go input.txt")
		return
	}

	filename := os.Args[1]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	lines := strings.Split(string(content), "\n")

	totalSum := 0
	for _, line := range lines {
		if line != "" {
			calibrationValue := extractCalibrationValue(line)
			totalSum += calibrationValue
		}
	}

	fmt.Println("Total Sum of Calibration Values:", totalSum)
}
