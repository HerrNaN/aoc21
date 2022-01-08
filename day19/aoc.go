package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Scanner []Point
type Rotation Scanner
type Point struct {
	Vec3
}
type Vec3 [3]int

func (p Point) Rotations() []Point {
	// ps := []Point{p}
	// for _, rm := range rotations {
	// 	ps = append(ps, Point{rm.Mul3x1(p.Vec3)})
	// }
	// return ps
	panic("unimplemented")
}

func (s Scanner) Rotations() []Rotation {
	panic("unimplemented")
}

func getSolutionPart1(ss []Scanner) int {
	return len(mergeRelativeToFirst(ss))
}

func getSolutionPart2(ss []Scanner) int {
	panic("unimplemented")
}

func mergeRelativeToFirst(ss []Scanner) Scanner {
	s0 := ss[0]
	for range ss {
		s0, ss = s0.mergeWithAny(ss)
	}
	return s0
}

func (s Scanner) mergeWithAny(ss []Scanner) (Scanner, []Scanner) {
	for i, other := range ss {
		if sm, ok := s.mergeWith(other); ok {
			s = sm
			return s, append(ss[0:i], ss[i+1:]...)
		}
	}
	panic("couldn't merge with any")
}

func (s Scanner) mergeWith(other Scanner) (Scanner, bool) {
	for _, r := range other.Rotations() {
		if sm, ok := s.mergeWithRotation(r); ok {
			s = sm
			return s, true
		}
	}
	return s, false
}

func (s Scanner) mergeWithRotation(r Rotation) (Scanner, bool) {
	if s.overlapsWith(r) {
		return s.merge(r), true
	}

	return s, false
}

func (s Scanner) overlapsWith(r Rotation) bool {
    panic("unimplemented")
}

func (s Scanner) merge(r Rotation) Scanner {
    panic("unimplemented")
}

func parseInput(input string) []Scanner {
	panic("unimplemented")
}

func readInput() []Scanner {
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
