package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"example.txt",
			142,
		},
	}
	for _, test := range tests {
		result, _ := part1(test.input)
		if result != test.expected {
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
			"example2.txt",
			281,
		},
	}
	for _, test := range tests {
		result, _ := part2(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestParseCalibrationValue(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{"1", "2", "3", "5"},
			15,
		},
		{
			[]string{"1"},
			11,
		},
		{
			[]string{"4", "5"},
			45,
		},
	}

	for _, test := range tests {
		result := parseCalibrationValue(test.input)

		firstIn := test.input[0]
		lastIn := test.input[len(test.input)-1]
		correct, _ := strconv.Atoi(firstIn + lastIn)

		if correct != result {
			t.Errorf("the result %v did not match the expected %v", result, correct)
		}

	}
}

func TestFilterDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"abc123def456", []string{"1", "2", "3", "4", "5", "6"}},
		{"12345", []string{"1", "2", "3", "4", "5"}},
		{"abc", []string{}},
	}

	for _, test := range tests {
		result := filterDigits(test.input)

		if len(result) != len(test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			continue
		}

		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("For input '%s', expeted %v, but got %v", test.input, test.expected, result)
				break
			}
		}
	}
}

func TestFilterWordAndDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"one1abctwo", []string{"1", "1", "2"}},
		{"12three3", []string{"1", "2", "3", "3"}},
		{"abc123", []string{"1", "2", "3"}},
	}

	for _, test := range tests {
		result := filterWordAndDigits(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
				break
			}
		}
	}
}

func TestSumIntSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 10},
	}

	for _, test := range tests {
		result := sumIntSlice(test.input)
		if result != test.expected {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestReadInput(t *testing.T) {
	tests := []struct {
		err      error
		input    string
		expected []string
	}{
		{
			nil,
			"example.txt",
			[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
		},
		{
			fmt.Errorf("path '%s' was not found", "examle.txt"),
			"examle.txt",
			[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
		},
	}

	for _, test := range tests {
		result, err := readInput(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("For input '%s', expected '%s'", result[i], test.expected[i])
			}
		}
		if test.err != nil && err == nil {
			t.Errorf("For input '%s', expected error '%s'", test.input, test.err)
		}
	}
}
