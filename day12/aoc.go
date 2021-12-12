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

type CaveMap map[string][]string

func getSolutionPart1(m CaveMap) int {
	return m.numUniquePaths(true)
}

func (m CaveMap) numUniquePaths(disallowVisitTwice bool) int {
	return len(m.uniquePathsFrom(CaveNameStart, make(map[string]bool), disallowVisitTwice))
}

func (m CaveMap) uniquePathsFrom(cave string, visited map[string]bool, visitedSmallCaveTwice bool) [][]string {
	if cave == CaveNameEnd {
		return [][]string{{CaveNameEnd}}
	}

	if isSmall(cave) && visited[cave] {
		if cave == CaveNameStart || visitedSmallCaveTwice {
			return nil
		}

		visitedSmallCaveTwice = true
	}

	visited[cave] = true


	var paths [][]string
	for _, nextCave := range m[cave] {
		paths = append(paths, m.uniquePathsFrom(nextCave, copyMap(visited), visitedSmallCaveTwice)...)
	}

	for i := range paths {
		paths[i] = append(paths[i], cave)
	}

	return paths
}

func isSmall(s string) bool {
	return s == strings.ToLower(s)
}

func copyMap(src map[string]bool) map[string]bool {
	dst := make(map[string]bool)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func getSolutionPart2(m CaveMap) int {
	return m.numUniquePaths(false)
}

func parseInput(input string) CaveMap {
	input = strings.TrimSpace(input)
	m := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, "-")
		m[splits[0]] = append(m[splits[0]], splits[1])
		m[splits[1]] = append(m[splits[1]], splits[0])
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
