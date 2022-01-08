package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input Input = nil

func TestAOC_parseInput(t *testing.T) {
	// 	inputString := `target area: x=20..30, y=-10..-5
	// `
	// 	actualParsedInput := parseInput(inputString)
	// 	assert.Equal(t, input, actualParsedInput)
}

func TestPoint_Rotations(t *testing.T) {
	tests := []struct {
		name   string
		point Point
		want   []Point
	}{
		{
			name:  "",
			point: Point{},
			want:  []Point{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.point.Rotations())
		})
	}
}

func TestAOC_getSolutionPart1(t *testing.T) {
	// assert.Equal(t, 45, getSolutionPart1(parseInput("target area: x=20..30, y=-10..-5")))
}

func TestAOC_getSolutionPart2(t *testing.T) {
	// assert.Equal(t, 112, getSolutionPart2(parseInput("target area: x=20..30, y=-10..-5")))
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

