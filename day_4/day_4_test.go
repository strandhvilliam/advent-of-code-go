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
		expected: 13,
	}

	result := part1(test.input)

	if result != test.expected {
		t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
	}
}

func TestPart2(t *testing.T) {
	test := struct {
		input    string
		expected int
	}{
		input:    "example.txt",
		expected: 30,
	}

	result := part2(test.input)

	if result != test.expected {
		t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
	}
}

func TestCalcPoints(t *testing.T) {
	tests := []struct {
		inputWins  []string
		inputDrawn []string
		expected   int
	}{
		{
			inputWins:  []string{"41", "48", "83", "86", "17"},
			inputDrawn: []string{"83", "86", "6", "31", "17", "9", "48", "53"},
			expected:   8,
		},
		{
			inputWins:  []string{"13", "32", "20", "16", "61"},
			inputDrawn: []string{"61", "30", "68", "82", "17", "32", "24", "19"},
			expected:   2,
		},
	}

	for _, test := range tests {
		result := calcPoints(test.inputWins, test.inputDrawn)
		if result != test.expected {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputWins, test.inputDrawn, test.expected, result)
		}
	}
}

// CARDS:  map[0:1 1:2 2:4 3:3 4:1]
// INDEX:  2
// LINE:  Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// RESULT:  map[0:1 1:2 2:4 3:4 4:2]

func TestUpdateCardCount(t *testing.T) {
	tests := []struct {
		inputCards map[int]int
		expected   map[int]int
		inputLine  string
		inputIndex int
	}{
		{
			inputCards: map[int]int{
				0: 1,
				1: 2,
				2: 4,
				3: 3,
				4: 1,
			},
			inputIndex: 2,
			inputLine:  "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expected: map[int]int{
				0: 1,
				1: 2,
				2: 4,
				3: 4,
				4: 2,
			},
		},
	}

	for _, test := range tests {

		cardsCopy := make(map[int]int)
		for k, v := range test.inputCards {
			cardsCopy[k] = v
		}

		updateCardCounts(cardsCopy, test.inputIndex, test.inputLine)

		if !reflect.DeepEqual(cardsCopy, test.expected) {
			t.Errorf("For input '%v', '%v' and '%v', expected %v, but got %v", test.inputCards, test.inputIndex, test.inputLine, test.expected, cardsCopy)
		}
	}
}

func TestSumCards(t *testing.T) {
	tests := []struct {
		input    map[int]int
		expected int
	}{
		{
			input: map[int]int{
				0: 1,
				1: 2,
				2: 4,
				3: 8,
				4: 14,
				5: 1,
			},
			expected: 30,
		},
	}

	for _, test := range tests {

		result := sumCards(test.input)

		if result != test.expected {
			t.Errorf("For input '%v', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestIntersections(t *testing.T) {
	tests := []struct {
		inputWins  []string
		inputDrawn []string
		expected   int
	}{
		{
			inputWins:  []string{"13", "32", "20", "16", "61"},
			inputDrawn: []string{"61", "30", "68", "82", "17", "32", "24", "19"},
			expected:   2,
		},
	}

	for _, test := range tests {
		result := intersections(test.inputWins, test.inputDrawn)
		if test.expected != result {
			t.Errorf("For input '%v' and '%v', expected %v, but got %v", test.inputWins, test.inputDrawn, test.expected, result)
		}
	}
}

func TestParseCard(t *testing.T) {
	tests := []struct {
		input         string
		expectedWins  []string
		expectedDrawn []string
	}{
		{
			input:         "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			expectedWins:  []string{"31", "18", "13", "56", "72"},
			expectedDrawn: []string{"74", "77", "10", "23", "35", "67", "36", "11"},
		},
	}

	for _, test := range tests {
		resWins, resDrawn := parseCard(test.input)
		if !reflect.DeepEqual(test.expectedWins, resWins) || !reflect.DeepEqual(test.expectedDrawn, resDrawn) {
			t.Errorf("For input '%v', expected %v and %v, but got %v and %v", test.input, test.expectedWins, test.expectedDrawn, resWins, resDrawn)
		}
	}
}
