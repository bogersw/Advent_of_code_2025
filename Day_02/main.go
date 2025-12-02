package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readInput reads the contents of the specified file with
// product ID ranges in a string slice and returns it.
// Each item in the string slice will be an ID range:
// xxxxx-yyyyy, with xxxxx the first ID and yyyyy the last ID.
func readInput(fileName string) []string {
	// Read the file (one line with comma-separated ID ranges)
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("could not open file `%s` -> %s", fileName, err))
	}
	line := strings.TrimSpace(string(data))
	// Read product ID ranges in a string slice
	idRanges := make([]string, 0, 50)
	if line != "" {
		idRanges = append(idRanges, strings.Split(line, ",")...)
		for i := range idRanges {
			idRanges[i] = strings.TrimSpace(idRanges[i])
		}
	}
	return idRanges
}

// SplitRange splits the specified string with a product ID range and returns
// the first ID and the last ID as integers
func splitRange(idRange string) (int, int) {
	endPoints := strings.Split(idRange, "-")
	// Get the first and last ID
	firstID, err := strconv.Atoi(endPoints[0])
	if err != nil {
		panic(fmt.Sprintf("could not process first ID `%s` -> %s", endPoints[0], err))
	}
	lastID, err := strconv.Atoi(endPoints[1])
	if err != nil {
		panic(fmt.Sprintf("could not process last ID `%s` -> %s", endPoints[1], err))
	}
	return firstID, lastID
}

// ############################################################################
// PART ONE
// ############################################################################

// processRange checks all the ID's in the specified range and returns
// those ID's that are made only of some sequence of digits repeated twice.
// The invalid ID's are returned as an integer slice.
func processRange(idRange string) []int {
	// Get the first ID and the last ID
	firstID, lastID := splitRange(idRange)
	// Check all numbers: numbers with an uneven number of digits
	// can't be invalid, for numbers with an even number of digits
	// we have to check if the first half equals the second half.
	// For this we convert the number to a string, split it and
	// compare the parts.
	result := make([]int, 0, 100)
	for i := firstID; i <= lastID; i++ {
		productID := strconv.Itoa(i)
		if len(productID) % 2 != 0 {
			continue
		}
		leftPart := productID[0:len(productID)/2]
		rightPart := productID[len(productID)/2:]
		if leftPart == rightPart {
			result = append(result, i)
		}
	}
	return result
}

// ############################################################################
// PART TWO
// ############################################################################

// splitNumber splits the specified number into parts of given size
func splitNumber(number string, size int) []string {
	// Check size and make sure that the length is divisible by size
	if size <= 0 || len(number)%size != 0 {
		return []string{}
	}
	// Split string in n parts and save parts in string slice
    n := len(number) / size
    parts := make([]string, n)
    for i := range n {
        parts[i] = number[i*size : (i+1)*size]
    }
    return parts
}

// IsInvalid checks if the the specified number is invalid by splitting it
// into equal parts and checking if all these parts are equal to each other
// and occur at least twice.
func isInvalid(number string) bool {
    n := len(number)
    for partSize := 1; partSize <= n/2; partSize++ {
        parts := splitNumber(number, partSize)
		// Count occurences of parts
        counts := make(map[string]int)
        for _, p := range parts {
            counts[p]++
            if counts[p] == len(parts) && counts[p] >= 2 {
                return true
            }
        }
    }
    return false
}

// processRange2 checks all the ID's in the specified range and returns
// those ID's that are invalid. Now invalid means that some sequence of
// digits is repeated AT LEAST twice.
// The invalid ID's are returned as an integer slice.
func processRange2(idRange string) []int {
	// Get the first ID and the last ID
	firstID, lastID := splitRange(idRange)
	// Check all numbers and identify those numbers that consist of
	// sequences of digits that ae repeated at least twice.
	result := make([]int, 0, 100)
	for i := firstID; i <= lastID; i++ {
		productID := strconv.Itoa(i)
		if isInvalid(productID) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// Read the file with product ID ranges
	idRanges := readInput("ranges.txt")

	// ########################################################################
	// PART ONE
	// ########################################################################
	sum := 0
	for _, idRange := range idRanges {
		numbers := processRange(idRange)
		for _, number := range numbers {
			sum += number
		}
	}
	fmt.Printf("Solution for part one: %d\n", sum)

	// ########################################################################
	// PART TWO
	// ########################################################################
	sum = 0
	for _, idRange := range idRanges {
		numbers := processRange2(idRange)
		for _, number := range numbers {
			sum += number
		}
	}
	fmt.Printf("Solution for part two: %d\n", sum)
}
