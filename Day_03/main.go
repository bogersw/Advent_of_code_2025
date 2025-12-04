package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readInput reads the contents of the specified file with
// battery banks in a string slice and returns it.
func readInput(fileName string) []string {
	banks := make([]string, 0, 100)
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("could not open file `%s` -> %s", fileName, err))
	}
	defer file.Close()
	// Process line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		banks = append(banks, line)
	}
	// Check if errors occuured during processing
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
	}
	return banks
}

// ############################################################################
// PART ONE
// ############################################################################

// processBank finds the TWO batteries in the specified bank that together
// produce the maximum joltage for that bank. Batteries can't be rearranged.
// The return value is the number formed by the digits on the two selected
// batteries.
func processBank(bank string) int {
	digits := strings.Split(bank, "")
	maximumJoltage := 0
	for i := range digits {
		for j := i + 1; j < len(digits); j++ {
			joltage, err := strconv.Atoi(digits[i] + digits[j])
			if err != nil {
				panic(fmt.Sprintf("could not process bank `%s` -> %s", bank, err))
			}
			if joltage > maximumJoltage {
				maximumJoltage = joltage
			}
		}
	}
	return maximumJoltage
}

// ############################################################################
// PART TWO
// ############################################################################

// processBank2 parses the specified bank string to repeatedly find the largest
// possible digit in a substring until twelve digits are found.
func processBank2(bank string) int {
    digits := strings.Split(bank, "")
    result := make([]string, 0, 12)
    currentIndex := 0
    for remaining := 12; remaining > 0; remaining-- {
        endIndex := len(digits) - remaining
        maxIndex := currentIndex
        maxDigit := digits[currentIndex]
        for i := currentIndex; i <= endIndex; i++ {
            if digits[i] > maxDigit {
                maxDigit = digits[i]
                maxIndex = i
                if maxDigit == "9" {
                    break
                }
            }
        }
        result = append(result, maxDigit)
        currentIndex = maxIndex + 1
    }
	// Return the maximum joltage as an integer
	maximumJoltage, err := strconv.Atoi(strings.Join(result, ""))
    if err != nil {
        panic(fmt.Sprintf("could not parse result `%s` -> %s", strings.Join(result, ""), err))
    }
    return maximumJoltage
}

func main() {
	// Read the input file with battery banks
	banks := readInput("banks.txt")

	// ############################################################################
	// PART ONE
	// ############################################################################
	totalJoltage := 0
	for _, bank := range banks {
		joltage := processBank(bank)
		totalJoltage += joltage
	}
	fmt.Printf("The total output joltage is: %d\n", totalJoltage)

	// ############################################################################
	// PART TWO
	// ############################################################################
	totalJoltage = 0
	for _, bank := range banks {
		joltage := processBank2(bank)
		totalJoltage += joltage
	}
	fmt.Printf("The total output joltage is: %d\n", totalJoltage)
}
