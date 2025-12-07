package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
)

// readInput reads the contents of the specified file
// into a string slice and returns it.
func readInput(fileName string) []string {
    manifold := make([]string, 0, 250)
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
    // Process line by line.
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        manifold = append(manifold, line)
    }
    // Check if errors occurred during processing
    if err := scanner.Err(); err != nil {
        panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
    }
    return manifold
}

// ############################################################################
// PART ONE
// ############################################################################

// determineStartIndex identifies the index where the beam starts its journey
// through the manifold.
func determineStartIndex(manifold []string) int {
    // The first line in the manifold diagram contains an `S`: this gives us
    // the start index.
    startIndex := 0
    for i := range manifold[0] {
        if manifold[0][i] == 'S' {
            startIndex = i
            break
        }
    }
    return startIndex
}

// processTachyonBeam follows the path of the beam in the manifold from top
// to bottom, counts the number of times the beam is split and returns that
// number.
func processTachyonBeam(manifold []string, startIndex int) int {
    // Initialize split count
    splitCount := 0
    // Start on line 1
    beamIndices := make([]int, 0, 25)
    beamIndices = append(beamIndices, startIndex)
    // Process line 1 -> len - 1: trace the beam from top to bottom and
    // keep track of the beam indices on each line and when they encounter
    // a `^` on the next line.
    for i := 1; i < len(manifold)-1; i++ {
        newBeamIndices := make([]int, 0, 25)
        for _, beamIndex := range beamIndices {
            if manifold[i+1][beamIndex] == '^' {
                if !slices.Contains(newBeamIndices, beamIndex-1) {
                    newBeamIndices = append(newBeamIndices, beamIndex-1)
                }
                if !slices.Contains(newBeamIndices, beamIndex+1) {
                    newBeamIndices = append(newBeamIndices, beamIndex+1)
                }
                splitCount++
            } else {
                if !slices.Contains(newBeamIndices, beamIndex) {
                    newBeamIndices = append(newBeamIndices, beamIndex)
                }
            }
        }
        beamIndices = newBeamIndices
    }
    return splitCount
}

// ############################################################################
// PART TWO
// ############################################################################

// processTachyonBeam2 follows the path of a particle in the manifold from top
// to bottom, counts the number of different timelines this particle ends up
// on and returns this number.
func processTachyonBeam2(manifold []string, startIndex int) int {
    // Initialize timeline count. We basically count how many times each
    // index is involved in a distinct timeline. We start with 1.
    timeLineCount := make([]int, len(manifold))
    timeLineCount[startIndex] = 1
    // Process line 1 -> len - 1: trace the particle from top to bottom and
    // keep track of how many possibilities there are (each path is another
    // timeline), taking into account the timelines we have already counted.
    for i := 1; i < len(manifold)-1; i++ {
        newTimeLineCount := make([]int, len(manifold[0]))
        for j, count := range timeLineCount {
            if count == 0 {
                // No timelines yet
                continue
            }
            // Check next line in the manifold
            if manifold[i+1][j] == '^' {
                newTimeLineCount[j-1] += count
                newTimeLineCount[j+1] += count
            } else {
                newTimeLineCount[j] += count
            }
        }
        timeLineCount = newTimeLineCount
    }
    // Count the total number of timelines
    totalTimeLineCount := 0
    for i := range timeLineCount {
        totalTimeLineCount += timeLineCount[i]
    }
    return totalTimeLineCount
}

func main() {

    manifold := readInput("manifold.txt")
    startIndex := determineStartIndex(manifold)

    // ############################################################################
    // PART ONE
    // ############################################################################

    splitCount := processTachyonBeam(manifold, startIndex)
    fmt.Printf("The number of times the beam is split: %d\n", splitCount)

    // ############################################################################
    // PART TWO
    // ############################################################################

    timelineCount := processTachyonBeam2(manifold, startIndex)
    fmt.Printf("The number of different time lines is: %d\n", timelineCount)

}
