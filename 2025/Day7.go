package main

import (
	"slices"
	"strings"
)

func D7P1() {
	fileLines := fileGrab()
	startIndx := -1
	count := 0
	var indecies []int
	for _, line := range fileLines {
		temp := strings.Split(line, "")
		if startIndx == -1 {
			startIndx = slices.Index(temp, "S")
			indecies = append(indecies, startIndx)
		} else {
			var tempIndecies []int
			for _, i := range indecies {
				if temp[i] == "^" {
					count += 1
					tempIndecies = append(tempIndecies, i-1)
					tempIndecies = append(tempIndecies, i+1)
				} else {
					tempIndecies = append(tempIndecies, i)
				}
			}
			indecies = slices.Compact(tempIndecies)
		}
	}
	println(count)
}

func D7P2() {
	fileLines := fileGrab()
	startIndx := -1
	sum := 0
	var indecies, counts []int
	for _, line := range fileLines {
		temp := strings.Split(line, "")
		if startIndx == -1 {
			for range temp {
				counts = append(counts, 0)
			}
			startIndx = slices.Index(temp, "S")
			indecies = append(indecies, startIndx)
			counts[startIndx] = 1
		} else {
			var tempIndecies []int
			for _, i := range indecies {
				if temp[i] == "^" {
					tempIndecies = append(tempIndecies, i-1)
					tempIndecies = append(tempIndecies, i+1)
					counts[i-1] += counts[i]
					counts[i+1] += counts[i]
					counts[i] = 0
				} else {
					tempIndecies = append(tempIndecies, i)
				}
			}
			indecies = slices.Compact(tempIndecies)
		}
	}
	for _, i := range counts {
		sum += i
	}
	println(sum)
}
