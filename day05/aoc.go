package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct{ x, y int }
type Vent struct{ from, to Point }
type Vents []Vent

func getSolutionPart1(vs Vents) int {
	covered := make(map[Point]int)

	for _, v := range vs {
		if v.isDiagonal() {
			continue
		}

		traversed := covers(v.from, v.to)
		for _, p := range traversed {
			covered[p]++
		}
	}

	sum := 0
	for _, p := range covered {
		if p >= 2 {
			sum++
		}
	}

	return sum
}

func (v Vent) isDiagonal() bool {
	return v.to.x != v.from.x && v.to.y != v.from.y
}

func covers(from, to Point) []Point {
	dX := to.x - from.x
	dY := to.y - from.y

	nPoints := 1 + int(math.Max(math.Abs(float64(dX)), math.Abs(float64(dY))))
	points := make([]Point, nPoints)

	for i := 0; i < nPoints; i++ {
		points[i] = Point{
			x: from.x + (i * multiplier(dX)),
			y: from.y + (i * multiplier(dY)),
		}
	}

	return points
}

func multiplier(i int) int {
	if i == 0 {
		return 0
	}
	if i > 0 {
		return 1
	}
	return -1
}

func getSolutionPart2(vs Vents) int {
	covered := make(map[Point]int)

	for _, v := range vs {
		traversed := covers(v.from, v.to)
		for _, p := range traversed {
			covered[p] = covered[p] + 1
		}
	}

	sum := 0
	for _, p := range covered {
		if p >= 2 {
			sum++
		}
	}

	return sum
}

func parseInput(input string) Vents {
	input = strings.TrimSpace(input)
	split := strings.Split(input, "\n")

	var vents Vents
	for _, line := range split {
		pointStrings := strings.Split(line, " -> ")
		from, to := parsePoint(pointStrings[0]), parsePoint(pointStrings[1])
		vents = append(vents, Vent{from, to})
	}

	return vents
}

func parsePoint(input string) Point {
	split := strings.Split(input, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return Point{x, y}
}

func readInput() Vents {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	return parseInput(string(inputBytes))
}

func main() {
	input := readInput()

	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
