package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func getSolutionPart1(input []int) int {
	sum := 0
	for i, value := range input {
		if IsPrime(value) {
			sum += i * value
		}
	}
	return sum
}

func IsPrime(n int) bool {
    for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return n > 1
}

func getSolutionPart2(input []int) int {
	sum := 0
	for i, value := range input {
		if !IsPrime(value) {
			if i%2 == 0 {
				sum += value
			} else {
				sum -= value
			}
		}
	}
	return sum
}

func parseInput(input string) ([]int, error) {
	var ints []int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("couldn't read input")
	}

	input, err := parseInput(string(inputBytes))
	if err != nil {
		panic("couldn't parse input")
	}

	fmt.Println("Go")
	part := os.Getenv("part")

	if part == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
