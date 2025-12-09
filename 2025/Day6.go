package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func D6P1() {
	fileLines := fileGrab()
	tall := len(fileLines) - 1
	re1 := regexp.MustCompile(`\b\s+\b`)
	re2 := regexp.MustCompile(` `)
	var noSpaces, iterated, operators []string
	total := 0

	for _, line := range fileLines {
		line = re1.ReplaceAllString(line, ",")
		line = re2.ReplaceAllString(line, "")
		noSpaces = append(noSpaces, line)
	}
	for _, l := range noSpaces {
		temp := strings.SplitSeq(l, ",")
		for i := range temp {
			iterated = append(iterated, i)
		}
	}
	lastitem := iterated[len(iterated)-1]
	iterated = slices.Delete(iterated, (len(iterated) - 1), (len(iterated)))
	operators = strings.Split(lastitem, "")
	operations := len(operators)
	for i, o := range operators {
		var products, sums []int
		count := 0
		product := 1
		sum := 0
		switch o {
		case "*":
			//multiply
			for count < tall {
				num, _ := strconv.Atoi(iterated[i+(operations*count)])
				products = append(products, num)
				count++
			}
			for _, p := range products {
				product *= p
			}
			total += product

		case "+":
			//add
			for count < tall {
				num, _ := strconv.Atoi(iterated[i+(operations*count)])
				sums = append(sums, num)
				count++
			}
			for _, s := range sums {
				sum += s
			}
			total += sum
		}
	}
	println(total)
}

/*

if operator - add / multiply
while < tall 0, 1, 2...

count = 0, 1, 2...
count ++

iterated[i + (len(operators)*count)]
iterated[i=0, + 0], count = 1 : 0 + 4x1 = 4)

*/

func D6P2() {
	fileLines := fileGrab()
	//tall := len(fileLines) - 1
	var rotated, operations []string
	var prods, sums []int
	lastitem := fileLines[len(fileLines)-1]
	fileLines = slices.Delete(fileLines, (len(fileLines) - 1), (len(fileLines)))
	var total int

	for x, line := range fileLines {
		temp := strings.Split(line, "")
		for i, t := range temp {
			if x == 0 {
				rotated = temp
			} else {
				rotated[i] += t
			}
		}
	}

	for i, r := range rotated {
		rotated[i] = strings.TrimSpace(r)
	}

	//getProducts
	prod := 1
	for x, r := range rotated {
		temp, _ := strconv.Atoi(r)
		if temp == 0 {
			prods = append(prods, prod)
			prod = 1
		} else if x == len(rotated)-1 {
			prod *= temp
			prods = append(prods, prod)
		} else {
			prod *= temp
		}
	}

	//getSums
	sum := 0
	for x, r := range rotated {
		temp, _ := strconv.Atoi(r)
		if temp == 0 {
			sums = append(sums, sum)
			sum = 0
		} else if x == len(rotated)-1 {
			sum += temp
			sums = append(sums, sum)
		} else {
			sum += temp
		}
	}

	operations = strings.Fields(lastitem)
	for i, o := range operations {
		switch o {
		case "+":
			total += sums[i]
		case "*":
			total += prods[i]
		}
	}
	println(total)
}
