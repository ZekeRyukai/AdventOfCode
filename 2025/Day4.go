package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func D4P1() {
	file, err := os.Open("Input.txt") // Replace "your_file.txt" with the actual file path
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	posLength := 0
	sum := 0
	tfs := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		posLength := len(line)

		s := strings.Split(line, "")
		for c := range s{
			if s[c] == "."
			{
				tfs = append(tfs, 0)
			}
			else{
				tfs = append(tfs, 1)
			}
		}
	}
	
	nwse := posLength + 1
	nesw := posLength - 1

	min := posLength
	max := len(tfs) - posLength
	
	for i := range tfs{
		switch tfs[i]{
		case 0:
			sum += 0
		case 1:
			var n int
			var s int
			var w int
			var e int
			var nw int
			var ne int
			var sw int
			var se int
			currentRolls := 0

			if i < min{
				nw, n, ne = 0, 0, 0
			}
			if i > max{
				sw, s, se = 0, 0, 0
			}
			if i % posLength = 0{
				nw, w, sw = 0, 0, 0
			}
			if i % nesw = 0{
				ne, e, se = 0, 0, 0
			}
			border := []int{nw, n, ne, e, se, s, sw, w}
			for x := range border{
				switch x {
				case 0:
					if border[x] != 0{
						currentRolls += tfs[i-nwse]
					}
				case 1:
					if border[x] != 0{
						currentRolls += tfs[i-posLength]
					}
				case 2:
					if border[x] != 0{
						currentRolls += tfs[i-nesw]
					}
				case 3:
					if border[x] != 0{
						currentRolls += tfs[i+1]
					}
				case 4:
					if border[x] != 0{
						currentRolls += tfs[i+nwse]
					}
				case 5:
					if border[x] != 0{
						currentRolls += tfs[i+posLength]
					}
				case 6:
					if border[x] != 0{
						currentRolls += tfs[i+nesw]
					}
				case 7:
					if border[x] != 0{
						currentRolls += tfs[i-1]
					}
				}
			}
			if currentRolls < 4{
				sum += 1
			}

		}
	}
}
