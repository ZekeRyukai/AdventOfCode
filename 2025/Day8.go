package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func D8P1() {
	fileLines := fileGrab()
	var pos [][]float64
	var boxes, circuits [][]int
	var order, lens, t []int
	var prox, cloneProx []float64
	product := 1

	for _, line := range fileLines {
		box := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(box[0], 64)
		y, _ := strconv.ParseFloat(box[1], 64)
		z, _ := strconv.ParseFloat(box[2], 64)

		pos = append(pos, []float64{x, y, z})
	}

	for i := range pos {
		n := i + 1
		for n < len(pos) {
			pq1 := math.Pow((pos[i][0] - pos[n][0]), 2)
			pq2 := math.Pow((pos[i][1] - pos[n][1]), 2)
			pq3 := math.Pow((pos[i][2] - pos[n][2]), 2)
			sum := pq1 + pq2 + pq3
			distance := math.Sqrt(sum)
			prox = append(prox, distance)
			boxes = append(boxes, []int{i + 1, n + 1})
			n++
		}
	}

	cloneProx = slices.Clone(prox)
	slices.Sort(cloneProx)

	for n := range cloneProx {
		order = append(order, slices.Index(prox, cloneProx[n]))
	}

	for n := range 1000 {
		b1 := boxes[order[n]][0]
		b2 := boxes[order[n]][1]

		if n == 0 {
			circuits = append(circuits, []int{b1, b2})
		} else {
			i1, i2 := -1, -1
			for x := range circuits {
				if slices.Contains(circuits[x], b1) {
					// note index
					i1 = x
				}
				if slices.Contains(circuits[x], b2) {
					// note index
					i2 = x
				}
			}
			println("Here  i1 = ", i1, " i2 = ", i2, "----- Box 1: ", b1, " Box 2: ", b2)
			//checking indexes
			if i1 == -1 && i2 == -1 {
				// new circuits slice
				circuits = append(circuits, []int{b1, b2})
			} else if i1 > -1 && i2 > -1 {
				// if match: nothing
				if i1 == i2 {
					continue
				} else if i1 < i2 { // if i1 lower
					// concat i2 TO i1, delete i2
					circuits[i1] = slices.Concat(circuits[i1], circuits[i2])
					slices.Delete(circuits, i2, i2+1)
				} else { // i1 > i2  // if i2 lower
					// concat i1 TO i2, delete i1
					circuits[i2] = slices.Concat(circuits[i2], circuits[i1])
					slices.Delete(circuits, i1, i1+1)
				}
			} else if i1 == -1 {
				// add  box 1 to box 2 circuit
				circuits[i2] = append(circuits[i2], b1)
			} else { // i2 == -1
				// add box 2 to box 1 circuit
				circuits[i1] = append(circuits[i1], b2)
			}
		}
	}

	for x := range circuits {
		for y := range circuits[x] {
			println("Circuit[] = ", x, "Circuits[x][] = ", y, "Value = ", circuits[x][y])
		}
		lens = append(lens, len(circuits[x]))
	}
	for range 3 {
		max := slices.Max(lens)
		maxIndex := slices.Index(lens, max)
		slices.Delete(lens, maxIndex, maxIndex+1)
		t = append(t, max)
	}
	for x := range t {
		println(t[x])
		product *= t[x]
	}
	println(product)

	//distance (p, q) = sqrt( (p1-q1)^2 + (p2-q2)^2 + (p3-q3)^2 )
	/*
		for each box, look at each circuit

		if box 1 is found, note index
		if box 2 is found, note index

		if box 1 not found, box 2 found - add box 1 to box 2 circuit
		if box 2 not found, box 1 found - add box 2 to box 1 circuit

		if neither found, new circuit

		If both found:

		if indexes match, do nothing
		if indexes don't match:
			note lower index:
			concat higher index w/ lower index
			delete higher index slice

	*/

}

func D8P2() {
	fileLines := fileGrab()
	var pos [][]float64
	var boxes, circuits, intPos [][]int
	var order []int
	var prox, cloneProx []float64
	product := 1
	var box1, box2 int

	for _, line := range fileLines {
		box := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(box[0], 64)
		y, _ := strconv.ParseFloat(box[1], 64)
		z, _ := strconv.ParseFloat(box[2], 64)

		ix, _ := strconv.Atoi(box[0])
		iy, _ := strconv.Atoi(box[1])
		iz, _ := strconv.Atoi(box[2])

		pos = append(pos, []float64{x, y, z})
		intPos = append(intPos, []int{ix, iy, iz})
	}

	for i := range pos {
		n := i + 1
		for n < len(pos) {
			pq1 := math.Pow((pos[i][0] - pos[n][0]), 2)
			pq2 := math.Pow((pos[i][1] - pos[n][1]), 2)
			pq3 := math.Pow((pos[i][2] - pos[n][2]), 2)
			sum := pq1 + pq2 + pq3
			distance := math.Sqrt(sum)
			prox = append(prox, distance)
			boxes = append(boxes, []int{i, n})
			n++
		}
	}

	cloneProx = slices.Clone(prox)
	slices.Sort(cloneProx)

	for n := range cloneProx {
		order = append(order, slices.Index(prox, cloneProx[n]))
	}

	for n := range cloneProx {
		b1 := boxes[order[n]][0]
		b2 := boxes[order[n]][1]

		if n == 0 {
			circuits = append(circuits, []int{b1, b2})
		} else {
			i1, i2 := -1, -1
			for x := range circuits {
				if slices.Contains(circuits[x], b1) {
					// note index
					i1 = x
				}
				if slices.Contains(circuits[x], b2) {
					// note index
					i2 = x
				}
			}
			println("Here  i1 = ", i1, " i2 = ", i2, "----- Box 1: ", b1, " Box 2: ", b2)
			//checking indexes
			if i1 == -1 && i2 == -1 {
				// new circuits slice
				circuits = append(circuits, []int{b1, b2})
			} else if i1 > -1 && i2 > -1 {
				// if match: nothing
				if i1 == i2 {
					continue
				} else if i1 < i2 { // if i1 lower
					// concat i2 TO i1, delete i2
					circuits[i1] = slices.Concat(circuits[i1], circuits[i2])
					slices.Delete(circuits, i2, i2+1)
				} else { // i1 > i2  // if i2 lower
					// concat i1 TO i2, delete i1
					circuits[i2] = slices.Concat(circuits[i2], circuits[i1])
					slices.Delete(circuits, i1, i1+1)
				}
			} else if i1 == -1 {
				// add  box 1 to box 2 circuit
				circuits[i2] = append(circuits[i2], b1)
			} else { // i2 == -1
				// add box 2 to box 1 circuit
				circuits[i1] = append(circuits[i1], b2)
			}
			if len(circuits[0]) == len(fileLines) {
				box1 = b1
				box2 = b2
				break
			}
		}
	}
	product = intPos[box1][0] * intPos[box2][0]
	println(product, intPos[box1][0], intPos[box2][0])
}
