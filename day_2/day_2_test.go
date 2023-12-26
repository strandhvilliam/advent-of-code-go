package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"example.txt",
			8,
		},
	}

	for _, test := range tests {
		result, _ := part1(test.input)
		if test.expected != result {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"example.txt",
			2286,
		},
	}

	for _, test := range tests {
		result, _ := part2(test.input)
		if test.expected != result {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestIsPossibleSet(t *testing.T) {
	tests := []struct {
		input    set
		expected bool
	}{
		{
			set{1, 1, 1},
			true,
		},
		{
			set{13, 1, 1},
			false,
		},
		{
			set{1, 14, 1},
			false,
		},
		{
			set{1, 1, 15},
			false,
		},
	}
	for _, test := range tests {
		result := isPossibleSet(test.input)
		if test.expected != result {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestParseGames(t *testing.T) {
	tests := []struct {
		input    []string
		expected []game
	}{
		{
			[]string{"Game 1: 1 red,1 blue,1 green;2 red,2 blue,2 green"},
			[]game{
				{
					sets: []set{
						{1, 1, 1},
					},
					id: 1,
				},
			},
		},
	}

	for _, test := range tests {
		result := parseGames(test.input)
		if len(test.expected) != len(result) {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestParseSets(t *testing.T) {
	tests := []struct {
		input    string
		expected []set
	}{
		{
			"1 red,1 blue,1 green;2 red,2 blue,2 green",
			[]set{
				{1, 1, 1},
				{2, 2, 2},
			},
		},
	}
	for _, test := range tests {
		result := parseSets(test.input)
		if len(test.expected) != len(result) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
		for i, v := range test.expected {
			if v != result[i] {
				t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			}
		}
	}
}

func TestParseIntoGameStruct(t *testing.T) {
	tests := []struct {
		input    string
		expected game
	}{
		{
			"Game 1: 1 red,1 blue,1 green;2 red,2 blue,2 green",
			game{
				[]set{
					{1, 1, 1},
					{2, 2, 2},
				},
				1,
			},
		},
	}
	for _, test := range tests {
		result := parseIntoGameStruct(test.input)
		if test.expected.id != result.id {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
		if len(test.expected.sets) != len(result.sets) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
		for i, v := range test.expected.sets {
			if v != result.sets[i] {
				t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			}
		}
	}
}

func TestParseColor(t *testing.T) {
	tests := []struct {
		input    string
		color    string
		expected int
	}{
		{
			"1 red,1 blue,1 green",
			"red",
			1,
		},
		{
			"1 red,1 blue,1 green",
			"blue",
			1,
		},
		{
			"1 red,1 blue,1 green",
			"green",
			1,
		},
		{
			"1 red,1 blue,1 green",
			"yellow",
			0,
		},
	}
	for _, test := range tests {
		result := countColor(test.input, test.color)
		if test.expected != result {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestSumGameIds(t *testing.T) {
	tests := []struct {
		input    []game
		expected int
	}{
		{
			[]game{
				{[]set{{1, 1, 1}}, 1},
				{[]set{{2, 2, 2}}, 2},
			},
			3,
		},
	}
	for _, test := range tests {
		result := sumGameIds(test.input)
		if test.expected != result {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestGetMinCubes(t *testing.T) {
	tests := []struct {
		input    []set
		expected set
	}{
		{
			[]set{{1, 1, 1}, {2, 2, 2}},
			set{2, 2, 2},
		},
	}

	for _, test := range tests {
		result := getMinCubes(test.input)
		if test.expected != result {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestPowerOfMins(t *testing.T) {
	tests := []struct {
		input    set
		expected int
	}{
		{
			set{2, 2, 2},
			8,
		},
	}
	for _, test := range tests {
		result := powerOfMins(test.input)
		if test.expected != result {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
