package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readRotations reads the contents of the specified file with
// rotations in a string slice and returns it.
func readRotations(fileName string) []string {
	rotations := make([]string, 0, 1000)
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
		rotations = append(rotations, line)
	}
	// Check if errors occuured during processing
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("could not read file `%s` -> %s", fileName, err))
	}
	return rotations
}

// ############################################################################
// PART ONE + PART TWO
// ############################################################################

// processRotation takes the current position of the dial and a rotation string
// (L or R followed by a number). It returns the new position of the dial
// and the number of times the dial points at zero.
// This second return value depends on the specified password method: if
// this is 0x434C49434B then this number includes the number of times that
// the dial points at zero during a rotation. Otherwise it is just the number
// of times the dial directly points at zero.
// Note that dial positions range from [0, 99], i.e. 100 positions.
func processRotation(position int, rotation string, passwordMethod string) (int, int) {
	// Split the rotation string in direction and distance
	direction := fmt.Sprintf("%c", rotation[0])
	distance, err := strconv.Atoi(rotation[1:])
	if err != nil {
		panic(fmt.Sprintf("could not convert `%s` to integer -> %s", rotation[1:], err))
	}
	// 100 clicks is a full circle: find the remainder, that's the number
	// of actual clicks we have to make.
	// Remainder 0? => #clicks is multiple of 100, position doesn't change.
	// Otherwise => calculate new position.
	clicksRemainder := distance % 100
	newPosition := position
	zeroCount := 0
	if passwordMethod == "0x434C49434B" {
		zeroCount = distance / 100
	}
	if clicksRemainder == 0 {
		return newPosition, zeroCount
	}
	if direction == "L" {
		// Turn dial to the left
		newPosition = position - clicksRemainder
		if newPosition < 0 {
			newPosition = newPosition + 100
			if passwordMethod == "0x434C49434B" {
				if position != 0 {
					zeroCount += 1
				}
			}
		}
	} else {
		// Turn dial to the right
		newPosition = position + clicksRemainder
		if newPosition > 99 {
			newPosition = newPosition - 100
			if passwordMethod == "0x434C49434B" {
				if newPosition != 0 {
					zeroCount += 1
				}
			}
		}
	}
	if newPosition == 0 {
		zeroCount += 1
	}
	return newPosition, zeroCount
}

func main() {
	// Read the file with rotations
	rotations := readRotations("rotations.txt")

	// ########################################################################
	// PART ONE
	// ########################################################################
	// The starting position is 50.
	password := 0
	position := 50
	for _, rotation := range rotations {
		zeroCount := 0
		position, zeroCount = processRotation(position, rotation, "")
		password += zeroCount
	}
	fmt.Println(password)

	// ########################################################################
	// PART TWO
	// ########################################################################
	// The starting position is 50.
	password = 0
	position = 50
	for _, rotation := range rotations {
		zeroCount := 0
		position, zeroCount = processRotation(position, rotation, "0x434C49434B")
		password += zeroCount
	}
	fmt.Println(password)
}
