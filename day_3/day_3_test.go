package main

import (
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	test := struct {
		input    string
		expected int
	}{
		input:    "example.txt",
		expected: 4361,
	}

	result := part1(test.input)
	if test.expected != result {
		t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
	}
}

func TestPart2(t *testing.T) {
	test := struct {
		input    string
		expected int
	}{
		input:    "example.txt",
		expected: 467835,
	}

	result := part2(test.input)
	if test.expected != result {
		t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
	}
}

func TestSumParts(t *testing.T) {
	tests := []struct {
		inputSymbols []symbol
		inputParts   []part
		expected     int
	}{
		{
			inputSymbols: []symbol{
				{
					adjacent: []point{{1, 2}, {3, 5}, {4, 2}},
					position: point{2, 2},
				},
			},
			inputParts: []part{
				{
					positions: []point{{1, 2}, {3, 5}},
					value:     10,
				},
			},
			expected: 10,
		},
	}

	for _, test := range tests {
		result := sumParts(test.inputSymbols, test.inputParts)
		if result != test.expected {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputSymbols, test.inputParts, test.expected, result)
		}
	}
}

func TestParseRatios(t *testing.T) {
	tests := []struct {
		inputParts  []part
		expected    []part
		inputSymbol symbol
	}{
		{
			inputSymbol: symbol{
				adjacent: []point{{1, 2}, {3, 5}, {4, 2}},
				position: point{2, 2},
			},
			inputParts: []part{
				{
					positions: []point{{1, 2}, {3, 5}},
					value:     10,
				},
				{
					positions: []point{{3, 2}, {3, 5}},
					value:     10,
				},
				{
					positions: []point{{2, 2}},
					value:     5,
				},
			},
			expected: []part{
				{
					positions: []point{{1, 2}, {3, 5}},
					value:     10,
				},
				{
					positions: []point{{3, 2}, {3, 5}},
					value:     10,
				},
			},
		},
	}

	for _, test := range tests {
		result := parseRatios(test.inputSymbol, test.inputParts)
		if len(result) != len(test.expected) {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputSymbol, test.inputParts, test.expected, result)
			continue
		}
		for i, part := range result {
			if !reflect.DeepEqual(part, test.expected[i]) {
				t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputSymbol, test.inputParts, test.expected, result)
				break
			}
		}
	}
}

func TestSumRatios(t *testing.T) {
	tests := []struct {
		inputParts   []part
		inputSymbols []symbol
		expected     int
	}{
		{
			inputSymbols: []symbol{
				{
					adjacent: []point{{1, 2}, {3, 5}, {4, 2}},
					position: point{2, 2},
				},
			},
			inputParts: []part{
				{
					positions: []point{{1, 2}, {3, 5}},
					value:     10,
				},
				{
					positions: []point{{3, 2}, {3, 5}},
					value:     10,
				},
				{
					positions: []point{{2, 2}},
					value:     5,
				},
			},
			expected: 100,
		},
	}
	for _, test := range tests {
		result := sumRatios(test.inputSymbols, test.inputParts)
		if result != test.expected {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputSymbols, test.inputParts, test.expected, result)
			break
		}
	}
}

func TestContainsPosition(t *testing.T) {
	tests := []struct {
		adjacent  []point
		positions []point
		expected  bool
	}{
		{
			adjacent:  []point{{1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}, {0, 0}, {0, 1}, {0, 2}},
			positions: []point{{2, 2}, {2, 3}, {2, 4}},
			expected:  true,
		},
		{
			adjacent:  []point{{1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}, {0, 0}, {0, 1}, {0, 2}},
			positions: []point{{3, 2}, {3, 3}, {3, 4}},
			expected:  false,
		},
	}

	for _, test := range tests {
		result := containsPosition(test.adjacent, test.positions)
		if result != test.expected {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.adjacent, test.positions, test.expected, result)
			break
		}
	}
}

func TestGetAdjacentPos(t *testing.T) {
	tests := []struct {
		expected []point
		lines    []string
		x        int
		y        int
	}{
		{
			expected: []point{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
			lines:    []string{"467..114..", "...*......", "..35..633."},
			x:        1,
			y:        1,
		},
	}

	for _, test := range tests {

		result := getAdjacentPos(test.y, test.x, test.lines)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%v', '%v' and '%v', expected %v, but got %v", test.x, test.y, test.lines, test.expected, result)
			break
		}
	}
}

func TestGetLastIndexOfPart(t *testing.T) {
	tests := []struct {
		line     string
		expected int
		startIdx int
	}{
		{
			line:     "467..114..",
			expected: 7,
			startIdx: 5,
		},
	}

	for _, test := range tests {
		result := getLastIdxOfPart(test.line, test.startIdx)
		if result != test.expected {
			t.Errorf("For input '%s' and '%v', expected %v, but got %v", test.line, test.startIdx, test.expected, result)
			break
		}
	}
}

func TestGetSymbol(t *testing.T) {
	tests := []struct {
		lines    []string
		expected symbol
		y        int
		x        int
	}{
		{
			lines: []string{"467..114..", "...*......", "..35..633."},
			y:     1,
			x:     3,
			expected: symbol{
				adjacent: []point{{0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 4}, {2, 2}, {2, 3}, {2, 4}},
				position: point{1, 3},
			},
		},
	}

	for _, test := range tests {
		result := getSymbol(test.lines, test.y, test.x)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%v', '%v' and '%v', expected %v, but got %v", test.lines, test.y, test.x, test.expected, result)
			break
		}
	}
}

func TestGetPart(t *testing.T) {
	tests := []struct {
		line          string
		expectedPart  part
		y             int
		x             int
		expectedIndex int
	}{
		{
			line: "467..114..",
			y:    0,
			x:    0,
			expectedPart: part{
				positions: []point{{0, 0}, {0, 1}, {0, 2}},
				value:     467,
			},
			expectedIndex: 2,
		},
	}

	for _, test := range tests {
		resPart, resIdx := getPart(test.line, test.y, test.x)
		if !reflect.DeepEqual(resPart, test.expectedPart) || resIdx != test.expectedIndex {
			t.Errorf("For input '%s', '%v' and '%v', expected '%v' and '%v', but got '%v' and '%v'", test.line, test.y, test.x, test.expectedPart, test.expectedIndex, resPart, resIdx)
		}
	}
}

func TestReadInput(t *testing.T) {
	test := struct {
		input    string
		expected []string
	}{
		input: "example.txt",
		expected: []string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		},
	}

	result, _ := readInput(test.input)
	if !reflect.DeepEqual(result, test.expected) {
		t.Errorf("For input '%s', expected '%v', but got '%v'", test.input, test.expected, result)
	}
}
