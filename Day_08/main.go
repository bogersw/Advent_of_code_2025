package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "slices"
    "sort"
    "strconv"
    "strings"
)

type pair struct {
    junctionBoxes []int
    distance      float64
}

type circuit struct {
    junctionBoxes []int
}

// readInput reads the contents of the specified file
// into a string slice and returns it.
func readInput(fileName string) [][]int {
    positions := make([][]int, 0, 250)
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
    // Process line by line. Split each line in x, y, z coordinates.
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        xyzString := strings.Split(line, ",")
        xyzInt := make([]int, len(xyzString))
        for i, coordinate := range xyzString {
            if xyzInt[i], err = strconv.Atoi(strings.TrimSpace(coordinate)); err != nil {
                panic(fmt.Sprintf("could not parse line `%s` -> %s", line, err))
            }
        }
        positions = append(positions, xyzInt)
    }
    // Check if errors occurred during processing
    if err := scanner.Err(); err != nil {
        panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
    }
    return positions
}

// ############################################################################
// PART ONE
// ############################################################################

// calculateDistance calculates the straight-line distance between two points
// in three-dimensional space. It returns this distance as a float64.
func calculateDistance(position1 []int, position2 []int) float64 {
    dX := float64(position2[0] - position1[0])
    dY := float64(position2[1] - position1[1])
    dZ := float64(position2[2] - position1[2])
    return math.Sqrt(dX*dX + dY*dY + dZ*dZ)
}

// sortByDistance sorts a slice of pair structs in ascending order based on the
// value of the distance field.
func sortByDistance(p []pair) {
    sort.Slice(p, func(i, j int) bool {
        if p[i].distance == p[j].distance {
            return p[i].junctionBoxes[0] < p[j].junctionBoxes[0]
        }
        return p[i].distance < p[j].distance
    })
}

// findCircuitIndex return the index of a circuit that contains the
// specified junctionBox or returns -1 if the junctionBox is not part
// of any circuit yet.
func findCircuitIndex(circuits []circuit, junctionBox int) int {
    for j := 0; j < len(circuits); j++ {
        if slices.Index(circuits[j].junctionBoxes, junctionBox) != -1 {
            return j
        }
    }
    return -1
}

// connectJunctionBoxes processes 1000 pairs of junction boxes and returns the
// product of the sizes of the three largest circuits.
func connectJunctionBoxes(positions [][]int) int {
    // Calculate distances between all possible pairs (note that the distance
    // between 1 and 2 is the same as the distance between 2 and 1, so count
    // only 1).
    pairs := make([]pair, 0, len(positions)*2)
    for i := 0; i < len(positions)-1; i++ {
        for j := i + 1; j < len(positions); j++ {
            distance := calculateDistance(positions[i], positions[j])
            pairs = append(pairs, pair{[]int{i, j}, distance})
        }
    }
    // Sort ascending by distance
    sortByDistance(pairs)
    // We now have all possible pairs sorted by their distance: we can now
    // process the pairs one by one and assign them to circuits (if any).
    // For part 1, we only check 1000 pairs (=len(positions)).
    circuits := make([]circuit, 0, len(positions))
    for i := 0; i < len(positions); i++ {
        // Junction boxes in the pair we're considering
        junctionBox1 := pairs[i].junctionBoxes[0]
        junctionBox2 := pairs[i].junctionBoxes[1]
        // Indexes of circuits the junction boxes belong to (-1 if not)
        idx1 := findCircuitIndex(circuits, junctionBox1)
        idx2 := findCircuitIndex(circuits, junctionBox2)
        // Process options for processing the junction boxes
        if idx1 != -1 && idx2 != -1 && idx1 != idx2 {
            // Both junction boxes are part of a circuit => join circuits
            circuits[idx1].junctionBoxes = append(circuits[idx1].junctionBoxes, circuits[idx2].junctionBoxes...)
            circuits = slices.Delete(circuits, idx2, idx2+1)
        }
        if idx1 != -1 && idx2 == -1 {
            // JunctionBox1 is part of a circuit => add JunctionBox2
            circuits[idx1].junctionBoxes = append(circuits[idx1].junctionBoxes, junctionBox2)
        }
        if idx1 == -1 && idx2 != -1 {
            // JunctionBox2 is part of a circuit => add JunctionBox1
            circuits[idx2].junctionBoxes = append(circuits[idx2].junctionBoxes, junctionBox1)
        }
        if idx1 == -1 && idx2 == -1 {
            // Both junction boxes are not part of a circuit => new circuit
            circuits = append(circuits, circuit{junctionBoxes: []int{junctionBox1, junctionBox2}})
        }
    }
    // Find circuit lengths and sort them in descending order
    sort.Slice(circuits, func(i, j int) bool {
        return len(circuits[i].junctionBoxes) >= len(circuits[j].junctionBoxes)
    })
    // Multiply size of first 3 circuits (or less)
    size := 1
    for i := range circuits {
        if i > 2 {
            break
        }
        size *= len(circuits[i].junctionBoxes)
    }
    return size
}

// ############################################################################
// PART TWO
// ############################################################################

// connectJunctionBoxes2 processes all pairs of junction boxes and returns the
// product of the x-coordinates of the last two junction boxes that are connected.
func connectJunctionBoxes2(positions [][]int) int {
    // Calculate distances between all possible pairs (note that the distance
    // between 1 and 2 is the same as the distance between 2 and 1, so count
    // only 1).
    pairs := make([]pair, 0, len(positions)*2)
    for i := 0; i < len(positions)-1; i++ {
        for j := i + 1; j < len(positions); j++ {
            distance := calculateDistance(positions[i], positions[j])
            pairs = append(pairs, pair{[]int{i, j}, distance})
        }
    }
    // Sort ascending by distance
    sortByDistance(pairs)
    // We now have all possible pairs sorted by their distance: we can now
    // process the pairs one by one and assign them to circuits (if any).
    circuits := make([]circuit, 0, len(positions))
    lastTwo := make([]int, 2, 2) // Keep track of last two junction boxes
    for i := 0; i < len(pairs); i++ {
        // Junction boxes in the pair we're considering
        junctionBox1 := pairs[i].junctionBoxes[0]
        junctionBox2 := pairs[i].junctionBoxes[1]
        // Indexes of circuits the junction boxes belong to (-1 if not)
        idx1 := findCircuitIndex(circuits, junctionBox1)
        idx2 := findCircuitIndex(circuits, junctionBox2)
        // Process options for processing the junction boxes
        if idx1 != -1 && idx2 != -1 && idx1 != idx2 {
            // Both junction boxes are part of a circuit => join circuits
            circuits[idx1].junctionBoxes = append(circuits[idx1].junctionBoxes, circuits[idx2].junctionBoxes...)
            circuits = slices.Delete(circuits, idx2, idx2+1)
            lastTwo = []int{junctionBox1, junctionBox2}
        }
        if idx1 != -1 && idx2 == -1 {
            // JunctionBox1 is part of a circuit => add JunctionBox2
            circuits[idx1].junctionBoxes = append(circuits[idx1].junctionBoxes, junctionBox2)
            lastTwo = []int{junctionBox1, junctionBox2}
        }
        if idx1 == -1 && idx2 != -1 {
            // JunctionBox2 is part of a circuit => add JunctionBox1
            circuits[idx2].junctionBoxes = append(circuits[idx2].junctionBoxes, junctionBox1)
            lastTwo = []int{junctionBox1, junctionBox2}
        }
        if idx1 == -1 && idx2 == -1 {
            // Both junction boxes are not part of a circuit => new circuit
            circuits = append(circuits, circuit{junctionBoxes: []int{junctionBox1, junctionBox2}})
            lastTwo = []int{junctionBox1, junctionBox2}
        }
    }
    // Get the x-coordinates of the last two junction boxes that were connected
    // and return the product.
    x1 := positions[lastTwo[0]][0]
    x2 := positions[lastTwo[1]][0]
    return x1 * x2
}

func main() {

    positions := readInput("positions.txt")

    // ############################################################################
    // PART ONE
    // ############################################################################
    sizeOfFirstThree := connectJunctionBoxes(positions)
    fmt.Printf("Product of sizes of 3 largest circuits: %d\n", sizeOfFirstThree)

    // ############################################################################
    // PART TWO
    // ############################################################################

    product := connectJunctionBoxes2(positions)
    fmt.Printf("Product of last 2 x-coordinates: %d\n", product)
}
