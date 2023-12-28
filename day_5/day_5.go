package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rangePair struct {
	src  int
	dest int
	rng  int
}

type seedRange struct {
	start int
	end   int
}

func main() {
	part1Result := part1("example.txt")
	part2Result := part2("example.txt")

	fmt.Println("-----DAY 5------")
	fmt.Println("Part 1: ", part1Result)
	fmt.Println("Part 2: ", part2Result)
}

func part1(file string) int {
	lines, err := readInput(file)
	if err != nil {
		fmt.Println(err)
	}

	seeds := parseSeedLine(lines[0])
	rngMap := parseRangeMapping(lines[2:])

	for loc := 0; ; loc++ {
		value := translateVal(rngMap, loc)
		if slices.Contains(seeds, value) {
			return loc
		}
	}
}

func part2(file string) int {
	lines, err := readInput(file)
	if err != nil {
		fmt.Println(err)
	}

	seeds := parseSeedRanges(lines[0])
	rngMap := parseRangeMapping(lines[2:])

	for loc := 0; ; loc++ {
		value := translateVal(rngMap, loc)
		if isInRange(seeds, value) {
			return loc
		}
	}
}

func parseRangeMapping(lines []string) [][]rangePair {
	mapping := make([][]rangePair, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if string(line[len(line)-1]) == ":" {
			mapping = append(mapping, []rangePair{})
			continue
		}
		values := getValuesFromLine(line)
		p := rangePair{values[0], values[1], values[2]}

		idx := len(mapping) - 1
		mapping[idx] = append(mapping[idx], p)
	}

	slices.Reverse(mapping)
	return mapping
}

func translateVal(mapping [][]rangePair, value int) int {
	for _, pairs := range mapping {
		for _, p := range pairs {
			if isMapped(value, p.src, p.rng) {
				n := calcNewVal(value, p.src, p.dest)
				value = n
				break
			}
		}
	}
	return value
}

func isInRange(seeds []seedRange, val int) bool {
	for _, rng := range seeds {
		if val >= rng.start && val <= rng.end {
			return true
		}
	}
	return false
}

func parseSeedRanges(line string) []seedRange {
	var ranges []seedRange
	ints := strings.Fields(strings.Split(line, ":")[1])
	size := 2

	for i := 0; i < len(ints); i += size {
		endIdx := i + size

		pair := ints[i:endIdx]

		start, _ := strconv.Atoi(pair[0])
		rng, _ := strconv.Atoi(pair[1])
		end := start + rng - 1
		r := seedRange{start: start, end: end}

		ranges = append(ranges, r)
	}
	return ranges
}

func parseSeedLine(line string) []int {
	seeds := strings.Fields(strings.Split(line, ":")[1])
	var parsed []int
	for _, s := range seeds {
		i, _ := strconv.Atoi(s)
		parsed = append(parsed, i)
	}
	return parsed
}

func getValuesFromLine(line string) []int {
	values := strings.Fields(line)
	dest, _ := strconv.Atoi(values[0])
	src, _ := strconv.Atoi(values[1])
	rng, _ := strconv.Atoi(values[2])
	return []int{dest, src, rng}
}

func isMapped(val int, start int, mRange int) bool {
	return start <= val && (start+mRange) >= val
}

func calcNewVal(val int, prevStart int, nextStart int) int {
	return nextStart + (val - prevStart)
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
