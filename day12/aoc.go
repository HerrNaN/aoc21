package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	CaveNameStart = "start"
	CaveNameEnd   = "end"
)

type Cave string
type CaveMap map[Cave][]Cave

func getSolutionPart1(m CaveMap) int {
	return m.numUniquePaths(true)
}

func (m CaveMap) numUniquePaths(disallowVisitTwice bool) int {
	return len(m.uniquePathsFrom(CaveNameStart, []Cave{}, disallowVisitTwice))
}

func (m CaveMap) uniquePathsFrom(cave Cave, visited []Cave, visitedSmallCaveTwice bool) [][]Cave {
	if cave == CaveNameEnd {
		return [][]Cave{{CaveNameEnd}}
	}

	if cave.isSmall() && contains(visited, cave) {
		if visitedSmallCaveTwice || cave == CaveNameStart {
			return nil
		}

		visitedSmallCaveTwice = true
	}

	paths := make([][]Cave, 0, len(m[cave]))
	for _, nextCave := range m[cave] {
		paths = append(paths, m.uniquePathsFrom(nextCave, append(visited, cave), visitedSmallCaveTwice)...)
	}

	for i := range paths {
		paths[i] = append(paths[i], cave)
	}

	return paths
}

func (c Cave) isSmall() bool {
	return string(c) == strings.ToLower(string(c))
}

func contains(ss []Cave, s Cave) bool {
	for i := range ss {
		if ss[i] == s {
			return true
		}
	}
	return false
}

func getSolutionPart2(m CaveMap) int {
	return m.numUniquePaths(false)
}

func parseInput(input string) CaveMap {
	input = strings.TrimSpace(input)
	m := make(map[Cave][]Cave)
	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, "-")
		m[Cave(splits[0])] = append(m[Cave(splits[0])], Cave(splits[1]))
		m[Cave(splits[1])] = append(m[Cave(splits[1])], Cave(splits[0]))
	}
	return m
}

func readInput() CaveMap {
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
