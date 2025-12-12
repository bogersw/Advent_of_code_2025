package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
)

type coordinate struct {
    row    int
    column int
}

// readInput reads the contents of the specified file into a slice of integer
// slices and returns it.
func readInput(fileName string) [][]int {
    coordinates := make([][]int, 0, 250)
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
    // Process line by line. Split each line in column, row coordinates.
    // Note that we reverse the order: column, row -> row, column because
    // the latter feels more natural to use.
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), ",")
        row, err := strconv.Atoi(strings.TrimSpace(line[1]))
        if err != nil {
            panic(fmt.Sprintf("could not parse line `%s` -> %s", line, err))
        }
        column, err := strconv.Atoi(strings.TrimSpace(line[0]))
        if err != nil {
            panic(fmt.Sprintf("could not parse line `%s` -> %s", line, err))
        }
        coordinates = append(coordinates, []int{row, column})
    }
    // Check if errors occurred during processing
    if err := scanner.Err(); err != nil {
        panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
    }
    return coordinates
}

// ############################################################################
// PART ONE
// ############################################################################

// determineLargestArea returns the area of the largest rectangle that can be
// made with two red tiles as opposite corners.
func determineLargestArea(coordinates [][]int) int {
    // Constants for readability
    const ROW = 0
    const COLUMN = 1
    // Sort coordinates (ascending row, ascending column)
    sort.Slice(coordinates, func(i, j int) bool {
        if coordinates[i][ROW] == coordinates[j][ROW] {
            return coordinates[i][COLUMN] < coordinates[j][COLUMN]
        }
        return coordinates[i][ROW] < coordinates[j][ROW]
    })
    // Consider all possible combinations of coordinates and calculate the area
    // for each combination. Return the largest area.
    maxArea := 0
    for i := 0; i < len(coordinates)-1; i++ {
        for j := i + 1; j < len(coordinates); j++ {
            rowWidth := coordinates[j][ROW] - coordinates[i][ROW] + 1
            columnHeight := coordinates[j][COLUMN] - coordinates[i][COLUMN] + 1
            area := rowWidth * columnHeight
            if area > maxArea {
                maxArea = area
            }
        }
    }
    return maxArea
}

// ############################################################################
// PART TWO
// ############################################################################

// determineLargestArea2 returns the area of the largest rectangle that can be
// made with two red tiles as opposite corners, but any other tiles it includes
// must now be red or green.
func determineLargestArea2(coordinates [][]int) int {
    // Time will tell if this part gets solved (although time already told me,
    // to be honest).
    return math.MaxInt
}

func main() {

    // ############################################################################
    // PART ONE
    // ############################################################################

    coordinates := readInput("example.txt")
    maxArea := determineLargestArea(coordinates)
    fmt.Printf("Largest area is (part one): %d\n", maxArea)

    // ############################################################################
    // PART TWO
    // ############################################################################

    coordinates = readInput("coordinates.txt")
    maxArea = determineLargestArea2(coordinates)
    fmt.Printf("Largest area is (part two): %d\n", maxArea)
}
