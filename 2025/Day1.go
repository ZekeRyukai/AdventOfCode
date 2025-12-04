package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func D1P1() { // Open the file
	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	currentNumber := 50
	count := 0
	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		//fmt.Println(line)      // Process the line (e.g., print it)

		direction := strings.TrimRight(line, "0123456789")
		turns, err := strconv.Atoi(strings.TrimLeft(line, "LR"))

		if err != nil {
			log.Fatalf("Error converting turns to integer: %v", err)
		}

		fmt.Println(direction, turns)

		switch direction {
		case "R":
			currentNumber += turns
		case "L":
			currentNumber -= turns
		default:
			log.Fatalf("Invalid direction: %s", direction)
		}

		if currentNumber%100 == 0 {
			count += 1
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %v", err)
	}
	fmt.Println(count)
}

func D1P2() {
	// Implementation for Day 1 Part 2 goes here
	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	currentNumber := 50
	count := 0
	currentState := "Positive"
	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		//fmt.Println(line)      // Process the line (e.g., print it)

		direction := strings.TrimRight(line, "0123456789")
		turns, err := strconv.Atoi(strings.TrimLeft(line, "LR"))

		if err != nil {
			log.Fatalf("Error converting turns to integer: %v", err)
		}

		fmt.Println(direction, turns)

		switch direction {
		case "R":
			currentNumber += turns
		case "L":
			currentNumber -= turns
		default:
			log.Fatalf("Invalid direction: %s", direction)
		}

		if currentNumber == 0 {
			count += 1
			currentState = "Zero"
		} else if currentNumber > 0 {
			if currentState == "Negative" {
				count += 1
			}
			currentState = "Positive"
		} else if currentNumber < 0 {
			if currentState == "Positive" {
				count += 1
			}
			currentState = "Negative"
		}

		count += int(math.Abs(float64(currentNumber / 100)))
		currentNumber = currentNumber - (100 * (currentNumber / 100))
		if currentNumber == 0 {
			currentState = "Zero"
		}
		fmt.Println(currentNumber, count)
	}

}
