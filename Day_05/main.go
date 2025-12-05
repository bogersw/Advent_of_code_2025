package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

// readInput reads the contents of the specified file
// into two string slices and returns them.
func readInput(fileName string) ([]string, []string) {
    freshIngredients := make([]string, 0, 100)
    availableIngredients := make([]string, 0, 100)
    // Open the file
    file, err := os.Open(fileName)
    if err != nil {
        panic(fmt.Sprintf("could not open file `%s` -> %s", fileName, err))
    }
    defer func(file *os.File) {
        err := file.Close()
        if err != nil {
            panic(fmt.Sprintf("could not close file `%s` -> %s", fileName, err))
        }
    }(file)
    // Process line by line. An empty line indicates the switch from
    // fresh ingredient ranges to available ingredient ID's.
    readAvailableIngredients := false
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if readAvailableIngredients {
            availableIngredients = append(availableIngredients, line)
        } else {
            readAvailableIngredients = line == ""
            if !readAvailableIngredients {
                freshIngredients = append(freshIngredients, line)
            }
        }
    }
    // Check if errors occurred during processing
    if err := scanner.Err(); err != nil {
        panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
    }
    return freshIngredients, availableIngredients
}

// processIngredientDatabase converts the ranges in text format to a
// slice of integer ranges (start / end).
func processIngredientDatabase(freshIngredients []string) [][]int {
    ranges := make([][]int, 0, len(freshIngredients))
    for i := range freshIngredients {
        start, err := strconv.Atoi(strings.Split(freshIngredients[i], "-")[0])
        if err != nil {
            panic(fmt.Sprintf("could not convert start `%s` -> %s", freshIngredients[i], err))
        }
        end, err := strconv.Atoi(strings.Split(freshIngredients[i], "-")[1])
        if err != nil {
            panic(fmt.Sprintf("could not convert end `%s` -> %s", freshIngredients[i], err))
        }
        ranges = append(ranges, []int{start, end})
    }
    return ranges
}

// ############################################################################
// PART ONE
// ############################################################################

// countFreshIngredients counts the number of fresh ingredients among the
// available ingredients and returns this number as an integer.
func countFreshIngredients(freshIngredients []string, availableIngredients []string) int {
    // Convert text ranges to integer ranges
    freshIngredientRanges := processIngredientDatabase(freshIngredients)
    // Check if available ID is in one of the fresh ID ranges
    freshCount := 0
    for i := range availableIngredients {
        id, _ := strconv.Atoi(availableIngredients[i])
        for _, idRange := range freshIngredientRanges {
            if id >= idRange[0] && id <= idRange[1] {
                freshCount++
                break
            }
        }
    }
    return freshCount
}

// ############################################################################
// PART TWO
// ############################################################################

// countUnique counts the number of unique ID's in the database with fresh
// ingredients. Ranges of ingredient ID's in the database can overlap.
func countUnique(freshIngredients []string) int {
    // Convert text ranges to integer ranges
    freshIngredientRanges := processIngredientDatabase(freshIngredients)
    // Sort by the first element of each inner slice (start of range).
    sort.Slice(freshIngredientRanges, func(i, j int) bool {
        a, b := freshIngredientRanges[i], freshIngredientRanges[j]
        return a[0] < b[0]
    })
    // Determine the number of unique ranges by checking start and end of each
    // range and comparing these with the unique ranges already found.
    uniqueRanges := make([][]int, 0, 100)
    for _, idRange := range freshIngredientRanges {
        processed := false
        if len(uniqueRanges) == 0 {
            uniqueRanges = append(uniqueRanges, []int{idRange[0], idRange[1]})
        } else {
            for _, uniqueRange := range uniqueRanges {
                if idRange[0] >= uniqueRange[0] && idRange[1] <= uniqueRange[1] {
                    // Ingredient range falls within existing unique range => no action
                    processed = true
                    break
                }
                if idRange[0] >= uniqueRange[0] && idRange[0] <= uniqueRange[1] {
                    // Ingredient range extends an existing unique range at the end
                    uniqueRange[1] = idRange[1]
                    processed = true
                    break
                }
                if idRange[0] <= uniqueRange[0] && idRange[1] >= uniqueRange[0] && idRange[1] <= uniqueRange[1] {
                    // Ingredient range extends an existing unique range at the start
                    uniqueRange[0] = idRange[0]
                    processed = true
                    break
                }
                if idRange[0] <= uniqueRange[0] && idRange[1] >= uniqueRange[1] {
                    // Ingredient range extends an existing unique range at the start and
                    // at the end
                    uniqueRange[0] = idRange[0]
                    uniqueRange[1] = idRange[1]
                    processed = true
                    break
                }
            }
            if !processed {
                // New, unique range
                uniqueRanges = append(uniqueRanges, []int{idRange[0], idRange[1]})
            }
        }
    }
    // Count unique ID's
    unique := 0
    for _, uniqueRange := range uniqueRanges {
        unique += uniqueRange[1] - uniqueRange[0] + 1
    }
    return unique
}

func main() {

    freshIngredients, availableIngredients := readInput("database.txt")

    // ############################################################################
    // PART ONE
    // ############################################################################
    freshCount := countFreshIngredients(freshIngredients, availableIngredients)
    fmt.Printf("The number of available ingredients that is fresh: %d\n", freshCount)

    // ############################################################################
    // PART TWO
    // ############################################################################
    uniqueCount := countUnique(freshIngredients)
    fmt.Printf("The number of unique, fresh ingredients is: %d\n", uniqueCount)
}
