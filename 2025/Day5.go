package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func fileGrab() []string {
	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	f := []string{}
	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, line)
	}
	return f
}

func D5P1() {
	fileLines := fileGrab()
	var ranges, numStrings []string
	var nums, lower, upper, valid []int
	var match bool

	for i, line := range fileLines {
		// Process each line as needed
		if line == "" {
			ranges = fileLines[:i]
			numStrings = fileLines[i+1:]
			for _, r := range ranges {
				_ = r
				a, b := strings.Split(r, "-")[0], strings.Split(r, "-")[1]
				x, _ := strconv.Atoi(a)
				y, _ := strconv.Atoi(b)
				lower = append(lower, x)
				upper = append(upper, y)
			}
			for _, n := range numStrings {
				_ = n
				i, _ := strconv.Atoi(n)
				nums = append(nums, i)
			}
			break
		}
	}
	println("Lower:")
	for _, l := range lower {
		_ = l
		println(l)
	}
	println("----")
	println("Upper:")
	for _, u := range upper {
		_ = u
		println(u)
	}
	println("----")
	println("Numbers:")
	for _, n := range nums {
		_ = n
		println(n)
	}

	for i := range ranges {
		for _, n := range nums {
			if n >= lower[i] && n <= upper[i] {
				if len(valid) == 0 {
					valid = append(valid, n)
					//println("Valid number found:", n)
				} else {
					for _, v := range valid {
						if n == v {
							match = true
						} else {
							match = false
						}
						if match {
							goto Next
						}
					}
				Next:
					if !match {
						valid = append(valid, n)
						println("Valid number found:", n)
					}
				}
			}
		}
	}
	println("Total valid numbers found:", len(valid))
}

func D5P2() {
	fileLines := fileGrab()
	var ranges []string
	var lower, upper []int
	onoff := true
	sum := 0

	for i, line := range fileLines {
		// Process each line as needed
		if line == "" {
			ranges = fileLines[:i]
			for _, r := range ranges {
				_ = r
				a, b := strings.Split(r, "-")[0], strings.Split(r, "-")[1]
				x, _ := strconv.Atoi(a)
				y, _ := strconv.Atoi(b)
				lower = append(lower, x)
				upper = append(upper, y)
			}
		}
	}
	for i := range lower {
		println("Range", i, ":", lower[i], "-", upper[i])
	}

	for onoff {
		onoff = false
		for u := range upper {
			for l := range lower {
				if upper[u] >= lower[l] && upper[u] < upper[l] { //if upper is between lower and upper of another range, replace upper of matching range and lower of
					if lower[u] > lower[l] {
						lower[u] = lower[l]
					} else {
						lower[l] = lower[u]
					}
					upper[u] = upper[l]
					onoff = true
				}
				if lower[u] > lower[l] && lower[u] <= upper[l] { //if lower is between lower and upper of another range, replace lower of matching range and upper of
					if upper[u] < upper[l] {
						upper[u] = upper[l]
					} else {
						upper[l] = upper[u]
					}
					lower[u] = lower[l]
					onoff = true
				}
			}
		}
	}

	sort.Ints(lower)
	sort.Ints(upper)
	j := 0
	for i := 1; i < len(lower); i++ {
		if lower[j] == lower[i] {
			continue
		}
		if upper[j] == upper[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		lower[j] = lower[i]
		upper[j] = upper[i]
	}
	lower = lower[:j+1]
	upper = upper[:j+1]
	fmt.Println(lower) // [1 2 3 4]
	fmt.Println(upper) // [1 2 3 4]
	for i := range lower {
		sum += upper[i] - lower[i] + 1
		println("New Ranges:", lower[i], "and", upper[i])
	}
	println("Sum of all ranges:", sum)
}
