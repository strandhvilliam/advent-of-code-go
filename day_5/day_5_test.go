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
		expected: 35,
	}

	result := part1(test.input)

	if result != test.expected {
		t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
	}
}

func TestPart2(t *testing.T) {
	test := struct {
		input    string
		expected int
	}{
		input:    "example.txt",
		expected: 46,
	}

	result := part2(test.input)

	if result != test.expected {
		t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
	}
}

func TestParseRangeMapping(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]rangePair
	}{
		{
			input: []string{"seed-to-soil map:", "50 98 2", "52 50 48"},
			expected: [][]rangePair{
				{
					{src: 50, dest: 98, rng: 2},
					{src: 52, dest: 50, rng: 48},
				},
			},
		},
	}

	for _, test := range tests {
		result := parseRangeMapping(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestTranslateVal(t *testing.T) {
	tests := []struct {
		inputMapping [][]rangePair
		inputValue   int
		expected     int
	}{
		{
			inputMapping: [][]rangePair{
				{
					{src: 50, dest: 98, rng: 2},
					{src: 52, dest: 50, rng: 48},
				},
			},
			inputValue: 53,
			expected:   51,
		},
	}

	for _, test := range tests {
		result := translateVal(test.inputMapping, test.inputValue)

		if result != test.expected {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputMapping, test.inputValue, test.expected, result)
		}
	}
}

func TestIsInRange(t *testing.T) {
	tests := []struct {
		inputRange []seedRange
		inputVal   int
		expected   bool
	}{
		{
			inputRange: []seedRange{
				{79, 92},
				{55, 67},
			},
			inputVal: 80,
			expected: true,
		},
		{
			inputRange: []seedRange{
				{79, 92},
				{55, 67},
			},
			inputVal: 70,
			expected: false,
		},
	}

	for _, test := range tests {
		result := isInRange(test.inputRange, test.inputVal)

		if result != test.expected {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputRange, test.inputVal, test.expected, result)
		}
	}
}

func TestParseSeedRange(t *testing.T) {
	tests := []struct {
		input    string
		expected []seedRange
	}{
		{
			input: "seeds: 79 14 55 13",
			expected: []seedRange{
				{79, 92},
				{55, 67},
			},
		},
	}

	for _, test := range tests {
		result := parseSeedRanges(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestParseSeedLine(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			input:    "seeds: 79 14 55 13",
			expected: []int{79, 14, 55, 13},
		},
	}

	for _, test := range tests {
		result := parseSeedLine(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestGetValuesFromLine(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			input:    "49 53 8",
			expected: []int{49, 53, 8},
		},
	}

	for _, test := range tests {
		result := getValuesFromLine(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
