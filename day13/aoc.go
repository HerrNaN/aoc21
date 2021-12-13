package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	FoldTypeVertical   = "x"
	FoldTypeHorizontal = "y"
)

type Fold struct {
	foldType string
	place    int
}
type Point struct{ x, y int }
type Paper map[Point]bool
type Manual struct {
	paper Paper
	folds []Fold
}

func getSolutionPart1(m Manual) int {
	m.doFold(m.folds[0])
	return len(m.paper)
}

func (m *Manual) foldAll() {
	for _, fold := range m.folds {
		m.doFold(fold)
	}
}

func (m *Manual) doFold(f Fold) {
	for p := range m.paper {
		switch f.foldType {
		case FoldTypeHorizontal:
			if p.y <= f.place {
				continue
			}

			delete(m.paper, p)
			diff := int(math.Abs(float64(f.place - p.y)))
			m.paper[Point{x: p.x, y: p.y - 2*diff}] = true
		case FoldTypeVertical:
			if p.x <= f.place {
				continue
			}

			delete(m.paper, p)
			diff := int(math.Abs(float64(f.place - p.x)))
			m.paper[Point{x: p.x - 2*diff, y: p.y}] = true
		}
	}
}

func (m *Manual) print() {
	var maxX, maxY int
	for p := range m.paper {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
	 	}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if m.paper[Point{x,y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func getSolutionPart2(m Manual) int {
	m.foldAll()
	m.print()
	return 0
}

func parseInput(input string) Manual {
	input = strings.TrimSpace(input)
	splits := strings.Split(input, "\n\n")
	return Manual{
		paper: parsePaper(splits[0]),
		folds: parseFolds(splits[1]),
	}
}

func parsePaper(input string) Paper {
	input = strings.TrimSpace(input)
	m := make(map[Point]bool)
	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, ",")
		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])
		m[Point{x, y}] = true
	}
	return m
}

func parseFolds(input string) []Fold {
	input = strings.TrimSpace(input)
	var folds []Fold
	for _, line := range strings.Split(input, "\n") {
		folds = append(folds, parseFold(line))
	}
	return folds
}

func parseFold(input string) Fold {
	input = strings.TrimSpace(input)
	splits := strings.Split(strings.Split(input, " ")[2], "=")
	place, _ := strconv.Atoi(splits[1])
	return Fold{
		foldType: splits[0],
		place:    place,
	}
}

func readInput() Manual {
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
