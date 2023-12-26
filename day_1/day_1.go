package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	exampleFile  string = "example.txt"
	example2File string = "example2.txt"
	inputFile    string = "input.txt"
)

func main() {
	part1Result, err := part1(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	part2Result, err := part2(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-----DAY 1------")
	fmt.Println("Part 1: ", part1Result)
	fmt.Println("Part 2: ", part2Result)
}

func part1(inputFile string) (int, error) {
	input, err := readInput(inputFile)
	if err != nil {
		return 0, err
	}

	var resultValues []int
	for _, word := range input {
		digits := filterDigits(word)
		calValue := parseCalibrationValue(digits)
		resultValues = append(resultValues, calValue)
	}
	return sumIntSlice(resultValues), nil
}

func part2(inputFile string) (int, error) {
	input, err := readInput(inputFile)
	if err != nil {
		return 0, err
	}

	var resultValues []int
	for _, line := range input {
		digits := filterWordAndDigits(line)
		calValues := parseCalibrationValue(digits)
		resultValues = append(resultValues, calValues)
	}

	return sumIntSlice(resultValues), nil
}

func sumIntSlice(intSlice []int) int {
	var sum int
	for _, num := range intSlice {
		sum += num
	}
	return sum
}

func parseCalibrationValue(values []string) int {
	if len(values) <= 0 {
		return 0
	}

	firstVal := values[0]
	lastVal := values[len(values)-1]
	calValue := firstVal + lastVal
	res, _ := strconv.Atoi(calValue)
	return res
}

func filterDigits(str string) []string {
	var digits []string
	for _, c := range str {
		if !unicode.IsDigit(c) {
			continue
		}
		digits = append(digits, string(c))
	}
	return digits
}

func readInput(path string) ([]string, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("path '%s' was not found", path)
	}
	defer input.Close()
	sc := bufio.NewScanner(input)

	lines := make([]string, 0)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func filterWordAndDigits(line string) []string {
	spelledDigits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var result []string

	for i, letter := range line {
		if unicode.IsDigit(letter) {
			result = append(result, string(letter))
			continue
		}

		for k, v := range spelledDigits {
			wordLen := len(k)
			if wordLen > len(line[i:]) {
				continue
			}
			if k == line[i:i+wordLen] {
				result = append(result, fmt.Sprint(v))
			}
		}

	}

	return result
}
