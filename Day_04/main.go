package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readInput reads the contents of the specified file with
// in a string slice and returns it.
func readInput(fileName string) []string {
	grid := make([]string, 0, 100)
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
		grid = append(grid, line)
	}
	// Check if errors occuured during processing
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
	}
	return grid
}

// ############################################################################
// PART ONE + PART 2
// ############################################################################

// extendGrid takes the input grid and puts a border of "." around it: in this
// way we can safely check the eight adjacent positions of points without the
// risk of out of bounds errors.
func extendGrid(grid []string) [][]string {
    rows := len(grid) + 2
    cols := len(grid[0]) + 2
    extendedGrid := make([][]string, rows)
    border := strings.Split(strings.Repeat(".", cols), "")
    extendedGrid[0] = border
    for i := range grid {
        extendedGrid[i+1] = strings.Split("."+grid[i]+".", "")
    }
    extendedGrid[rows-1] = border
    return extendedGrid
}

// findAccessibleRolls parses the specified grid and finds the rolls of paper
// that are accessible to a forklift. A paper roll (@) is accessible when there
// are fewer than four rolls of paper in the eight adjacent positions.
// If removeAccessibleRolls is true, then accessible rolls are removed from
// the grid by setting them to "." (the specified grid is modified in this case).
func findAccessibleRolls(grid []string, removeAccessibleRolls bool) int {
    extendedGrid := extendGrid(grid)
    countedRolls := 0
	accessibleRollPositions := make([][]int, 0, 100)
    for row := 1; row < len(extendedGrid)-1; row++ {
        for column := 1; column < len(extendedGrid[row])-1; column++ {
            mark := extendedGrid[row][column]
            if mark == "." {
                continue
            }
			// Check eight adjacent points
            adjacentRolls := 0
            for i := -1; i <= 1; i++ {
                for j := -1; j <= 1; j++ {
                    if i == 0 && j == 0 {
                        continue
                    }
                    if extendedGrid[row+i][column+j] == "@" {
                        adjacentRolls++
                    }
                }
            }
            if adjacentRolls < 4 {
				accessibleRollPositions = append(accessibleRollPositions, []int{row-1, column-1})
                countedRolls++
            }
        }
    }
	if removeAccessibleRolls {
		for _, position := range accessibleRollPositions {
			line := strings.Split(grid[position[0]], "")
			line[position[1]] = "."
			grid[position[0]] = strings.Join(line, "")
		}
	}
    return countedRolls
}

func main() {
	// Read the input file with the grid
	grid := readInput("grid.txt")

	// ############################################################################
	// PART ONE
	// ############################################################################
	fmt.Printf("The number of accessible rolls is: %d\n", findAccessibleRolls(grid, false))

	// ############################################################################
	// PART TWO
	// ############################################################################

	totalRollCount := 0
	for {
		// Keep counting until no accessible rolls are left
		rollCount := findAccessibleRolls(grid, true)
		if rollCount == 0 {
			break
		}
		totalRollCount += rollCount
	}
	fmt.Printf("The number of accessible rolls is: %d\n", totalRollCount)
}
