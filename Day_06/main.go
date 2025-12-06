package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// splitLine returns the items on a line as a string slice. The items
// are determined with the columns starting indices in `columnStarts`.
func splitLine(line string, columnStarts []int) []string {
    width := len(line)
    result := make([]string, 0, len(columnStarts))
    for i := 0; i < len(columnStarts); i++ {
        start := columnStarts[i]
        var end int
        if i+1 < len(columnStarts) {
            end = columnStarts[i+1] - 1
        } else {
            end = width
        }
        result = append(result, line[start:end])
    }
    return result
}

// determineColumnStarts returns indices where a column starts based on the
// line with the operators: these are always placed at the start of a column.
func determineColumnStarts(fileName string) []int {
    starts := make([]int, 0, 50)
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
    // Process line by line. The last line contains the operators.
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line[0] == '*' || line[0] == '+' {
            // Line with the operators
            for i := 0; i < len(line); i++ {
                if line[i] != ' ' {
                    starts = append(starts, i)
                }
            }
        }
    }
    // Check if errors occurred during processing
    if err := scanner.Err(); err != nil {
        panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
    }
    return starts
}

// readInput reads the contents of the specified file
// into two string slices and returns them.
func readInput(fileName string, columnStarts []int) [][]string {
    worksheet := make([][]string, 0, 50)
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
    // Process line by line. The last line contains the operators.
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        worksheet = append(worksheet, splitLine(line, columnStarts))
    }
    // Check if errors occurred during processing
    if err := scanner.Err(); err != nil {
        panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
    }
    return worksheet
}

// processProblem returns the result of a single problem. Inputs are a slice
// with numbers and the operator that has to be applied.
func processProblem(numbers []string, operator string) int {
    var result int
    // Process the rest of the numbers
    for i := 0; i < len(numbers); i++ {
        num, err := strconv.Atoi(numbers[i])
        if err != nil {
            panic(fmt.Sprintf("could not parse `%s` -> %s", numbers[i], err))
        }
        if i == 0 {
            result = num
            continue
        }
        switch operator {
        case "+":
            result += num
        case "*":
            result *= num
        case "-":
            result -= num
        default:
            panic(fmt.Sprintf("unknown operator `%s` -> %s", operator, numbers[i]))
        }
    }
    return result
}

// ############################################################################
// PART ONE
// ############################################################################

// processProblems solves the problems in the worksheet based on the description
// in part one. Numbers for each problem are arranged vertically.
func processProblems(worksheet [][]string) int {
    numberCount := len(worksheet[0])
    sum := 0
    for column := 0; column < numberCount; column++ {
        numbers := make([]string, 0, numberCount)
        for row := 0; row < len(worksheet)-1; row++ {
            numbers = append(numbers, strings.TrimSpace(worksheet[row][column]))
        }
        operator := strings.TrimSpace(worksheet[len(worksheet)-1][column])
        sum += processProblem(numbers, operator)
    }
    return sum
}

// ############################################################################
// PART TWO
// ############################################################################

// processProblems2 solves the problems in the worksheet based on the description
// in part two. In this part spaces in columns matter for alignment (the numbers
// we need are written right-to-left in columns with the most significant digit
// at the top and the least significant digit at the bottom).
func processProblems2(worksheet [][]string) int {
    numberCount := len(worksheet[0])
    sum := 0
    for column := 0; column < numberCount; column++ {
        numbers := make([]string, 0, numberCount)
        numberWidth := len(worksheet[0][column])
        for width := numberWidth - 1; width >= 0; width-- {
            number := ""
            for row := 0; row < len(worksheet)-1; row++ {
                number += string(worksheet[row][column][width])
            }
            numbers = append(numbers, strings.TrimSpace(number))
        }
        operator := strings.TrimSpace(worksheet[len(worksheet)-1][column])
        sum += processProblem(numbers, operator)
    }
    return sum
}

func main() {

    // Determine the starting indices of the columns in the file. We use
    // the line with operators for this (this is the last line in the file).
    // When we know the indices where the columns start, we can read in the file.
    columnsStarts := determineColumnStarts("worksheet.txt")
    worksheet := readInput("worksheet.txt", columnsStarts)

    // ############################################################################
    // PART ONE
    // ############################################################################

    sum := processProblems(worksheet)
    fmt.Printf("The total sum of all answers for part one is: %d\n", sum)

    // ############################################################################
    // PART TWO
    // ############################################################################

    sum = processProblems2(worksheet)
    fmt.Printf("The total sum of all answers for part 2 is: %d\n", sum)
}
