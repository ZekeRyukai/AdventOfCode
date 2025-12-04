package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func D2P1() {
	sum := 0

	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string

		for rs := range strings.SplitSeq(line, ",") {
			start, end := strings.Split(rs, "-")[0], strings.Split(rs, "-")[1]

			s, err1 := strconv.Atoi(start)
			if err1 != nil {
				log.Fatalf("Error converting start to integer: %v", err1)
			}
			e, err2 := strconv.Atoi(end)
			if err2 != nil {
				log.Fatalf("Error converting start to integer: %v", err2)
			}

			for s <= e {
				// Convert s to string slice
				check := strconv.Itoa(s)

				l := len(check)
				if l%2 == 0 {
					mid := l / 2
					firstHalf := check[:mid]
					secondHalf := check[mid:]
					if firstHalf == secondHalf {
						sum += s
						println(firstHalf, secondHalf)
					}
				}
				s += 1
			}
		}
	}
	println(sum)
}

func D2P2() {
	sum := 0

	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string

		for rs := range strings.SplitSeq(line, ",") {
			start, end := strings.Split(rs, "-")[0], strings.Split(rs, "-")[1]

			s, err1 := strconv.Atoi(start)
			if err1 != nil {
				log.Fatalf("Error converting start to integer: %v", err1)
			}
			e, err2 := strconv.Atoi(end)
			if err2 != nil {
				log.Fatalf("Error converting start to integer: %v", err2)
			}

			for s <= e {
				// Convert s to string slice
				check := strconv.Itoa(s)
				l := len(check)
				mid := l / 2
				firstHalf := check[:mid]
				secondHalf := check[mid:]
				if firstHalf == secondHalf {
					sum += s
					println(s)
				} else {
					for i := 1; i < mid+1; i++ {
						if l%i == 0 {
							iCheck := check[:i]
							diff := l / i
							repeated := strings.Repeat(iCheck, diff)
							if repeated == check {
								sum += s
								println(s)
							}
						}

					}

				}
				s += 1
			}
		}
		println(sum)
	}
}
