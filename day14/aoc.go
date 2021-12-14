package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type Template string
type Rules map[string]rune
type Formula struct {
	template Template
	rules Rules
}

type Reaction struct {
	step int
	pair string
}

func getSolutionPart1(f Formula) int {
	freqs := f.freqAfterSteps(10)
	// fmt.Println(sum(freqs))
	min, max := min(freqs), max(freqs)
	return max - min
}

func (f Formula) freqAfterSteps(n int) map[rune]int {
	m := make(map[rune]int)
	for _, r := range f.template {
		m[r]++
	}
	cache := make(map[Reaction]map[rune]int)
	for i := 0; i < len(f.template)-1; i++ {
		unionWith(m, f.rules.countReaction(Reaction{n, string(f.template)[i:i+2]}, cache), add)
	}
	return m
}

func max(m map[rune]int) int {
	var n int
	for _, v := range m {
		if v > n {
			n = v
		}
	}
	return n
}

func min(m map[rune]int) int {
	n := math.MaxInt
	for _, v := range m {
		if v < n {
			n = v
		}
	}
	return n
}

func sum(m map[rune]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}


// Unions into m1
func unionWith(m1,m2 map[rune]int, f func(int,int) int) {
	if m1 == nil {
		m1 = m2
	} else if m2 != nil {
		for k := range m2 {
			m1[k] = f(m1[k],m2[k])
		}
	}
}

func add(a,b int) int {
	return a+b
}

func (rules Rules) countReaction(r Reaction, cache map[Reaction]map[rune]int) map[rune]int {
	// fmt.Println(r.pair)
	if r.step == 0 {
		return nil
	}

	if m, ok := cache[r]; ok {
		return m
	}

	m := map[rune]int{
		rules[r.pair]: 1,
	}
	ar := Reaction{r.step-1, string(rune(r.pair[0]))+string(rules[r.pair])}
	a := rules.countReaction(ar, cache)
	unionWith(m, a, add)

	br := Reaction{r.step-1, string(rules[r.pair])+string(rune(r.pair[1]))}
	b := rules.countReaction(br, cache)
	unionWith(m, b, add)

	cache[r] = m

	return m
}

func getSolutionPart2(f Formula) int {
	freqs := f.freqAfterSteps(40)
	// fmt.Println(sum(freqs))
	min, max := min(freqs), max(freqs)
	return max - min
}

func parseInput(input string) Formula {
	input = strings.TrimSpace(input)
	splits := strings.Split(input, "\n\n")
	return Formula{
		template: Template(splits[0]),
		rules:    parseRules(splits[1]),
	}
}

func parseRules(input string) Rules {
	input = strings.TrimSpace(input)
	rules := make(Rules)
	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, " -> ")
		rules[splits[0]] = rune(splits[1][0])
	}
	return rules
}

func readInput() Formula {
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
