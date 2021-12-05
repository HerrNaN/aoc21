package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct { x,y int }
type Vent struct { from, to Point }
type Vents []Vent

func getSolutionPart1(vs Vents) int {
	covered := make(map[Point]int)

	for _, v := range vs {
		traversed := covers(v.from, v.to, false)
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

func covers(from, to Point, wantDiagonals bool) []Point {
	dX := to.x - from.x
	dY := to.y - from.y

	var points []Point
	if dX == 0  {
		if dY > 0 {
			for i := 0; i <= dY; i++ {
				points = append(points, Point{from.x, from.y+i})
			}
		} else if dY < 0 {
			for i := 0; i >= dY; i-- {
				points = append(points, Point{from.x, from.y+i})
			}
		}
	} else if dY == 0 {
		if dX > 0 {
			for i := 0; i <= dX; i++ {
				points = append(points, Point{from.x+i, from.y})
			}
		} else if dX < 0 {
			for i := 0; i >= dX; i-- {
				points = append(points, Point{from.x+i, from.y})
			}
		}
	} else if wantDiagonals {
		for i := 0; i <= int(math.Abs(float64(dX))); i++ {
			points = append(points, Point{from.x+(i*sign(dX)), from.y+(i*sign(dY))})
		}
	}

	return points
}

func sign(i int) int {
	if i < 0 {
		return -1
	}

	return 1
}

func getSolutionPart2(vs Vents) int {
	covered := make(map[Point]int)

	for _, v := range vs {
		traversed := covers(v.from, v.to, true)
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
		pointStrings := strings.Split(line," -> ")
		from, to := parsePoint(pointStrings[0]), parsePoint(pointStrings[1])
		vents = append(vents, Vent{from, to})
	}

	return vents
}

func parsePoint(input string) Point {
	split := strings.Split(input, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return Point{x,y}
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
