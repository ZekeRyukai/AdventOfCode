package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func D3P1() {
	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		bank := scanner.Text()
		nums := strings.SplitN(bank, "", -1)
		first := 0
		second := 0
		index := -1
		for i := range nums {
			n, _ := strconv.Atoi(nums[i])
			if n > first && i > index && i < len(nums)-1 {
				first = n
				index = i
				second, _ = strconv.Atoi(nums[i+1])
			}
			if n > second && i != index {
				second = n
			}
		}
		jolts := (first * 10) + second
		println(jolts)
		sum += jolts
	}
	println(sum)
}

func D3P2() {
	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		bank := scanner.Text()
		nums := strings.SplitN(bank, "", -1)
		high := 0
		highIndex := -1
		intNums := []int{}
		twelve := 0

		for i := len(nums) - 1; i >= 0; i-- {
			n, _ := strconv.Atoi(nums[i])
			intNums = append(intNums, n)
		}
		start := 11
		end := len(intNums)

		for j := start; start > -1 && j < end; j++ {
			//println(j, start, end)
			if intNums[j] >= high {
				high = intNums[j]
				highIndex = j
			}
			if j == end-1 {
				end = highIndex
				twelve = (twelve * 10) + high
				start += -1
				j = start - 1
				high = 0
				highIndex = -1
			}
		}
		println(twelve)
		sum += twelve
	}
	println(sum)
}
