package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = Formula{
	template: "NNCB",
	rules: map[string]rune{
		"CH": 'B',
		"HH": 'N',
		"CB": 'H',
		"NH": 'C',
		"HB": 'C',
		"HC": 'B',
		"HN": 'C',
		"NN": 'C',
		"BH": 'H',
		"NC": 'B',
		"NB": 'B',
		"BN": 'B',
		"BB": 'N',
		"BC": 'B',
		"CC": 'N',
		"CN": 'C',
	},
}

func TestAOC_parseInput(t *testing.T) {
	inputString := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`
	actualParsedInput := parseInput(inputString)
	assert.Equal(t, input, actualParsedInput)
}

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := 1588
	actualSolution := getSolutionPart1(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 2188189693529
	actualSolution := getSolutionPart2(input)
	assert.Equal(t, expectedSolution, actualSolution)
}

func Benchmark_Parse(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		readInput()
	}
}

func Benchmark_Part1(b *testing.B) {
	b.ReportAllocs()

	input := readInput()
	for i := 0; i < b.N; i++ {
		getSolutionPart1(input)
	}
}

func Benchmark_Part2(b *testing.B) {
	b.ReportAllocs()

	input := readInput()
	for i := 0; i < b.N; i++ {
		getSolutionPart2(input)
	}
}
