package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

const (
	exampleFile string = "example.txt"
	inputFile   string = "input.txt"
)

type symbol struct {
	adjacent []point
	ratio    []part
	position point
}

type part struct {
	positions []point
	value     int
}

type point struct {
	y int
	x int
}

func main() {
	part1Result := part1(inputFile)
	part2Result := part2(inputFile)

	fmt.Println("-----DAY 3------")
	fmt.Println("Part 1: ", part1Result)
	fmt.Println("Part 2: ", part2Result)
}

func part1(file string) int {
	lines, err := readInput(file)
	if err != nil {
		fmt.Println(err)
	}

	symbols, values := parsePartsAndSymbols(lines, isSymbol)

	sum := 0
	for _, v := range values {
		for _, s := range symbols {
			if containsPosition(s.adjacent, v.positions) {
				sum += v.value
				continue
			}
		}
	}

	return sum
}

func part2(filePath string) int {
	lines, err := readInput(filePath)
	if err != nil {
		fmt.Println(err)
	}

	symbols, parts := parsePartsAndSymbols(lines, isGear)
	res := sumRatios(symbols, parts)
	return res
}

func sumRatios(symbols []symbol, parts []part) int {
	sum := 0
	for _, g := range symbols {
		ratios := parseRatios(g, parts)
		if len(ratios) > 1 {
			mult := 1
			for _, r := range ratios {
				mult *= r.value
			}
			sum += mult
		}
	}
	return sum
}

func parseRatios(g symbol, p []part) []part {
	var ratios []part

	for _, val := range p {
		if containsPosition(g.adjacent, val.positions) {
			ratios = append(ratios, val)
		}
	}
	return ratios
}

func isSymbol(b byte) bool {
	return string(b) != "." && !unicode.IsDigit(rune(b))
}

func isGear(b byte) bool {
	return string(b) == "*"
}

func parsePartsAndSymbols(lines []string, checkFn func(byte) bool) ([]symbol, []part) {
	symbols := make([]symbol, 0)
	parts := make([]part, 0)
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if unicode.IsDigit(rune(line[x])) {
				part, lastIdx := getPart(line, y, x)
				parts = append(parts, part)
				x = lastIdx
			}
			if checkFn(lines[y][x]) {
				s := getSymbol(lines, y, x)
				symbols = append(symbols, s)
			}
		}
	}
	return symbols, parts
}

func getSymbol(lines []string, y int, x int) symbol {
	s := symbol{
		adjacent: getAdjacentPos(y, x, lines),
		position: point{y, x},
		ratio:    make([]part, 0),
	}
	return s
}

func getPart(line string, y int, x int) (part, int) {
	lIndex := getLastIdxOfPart(line, x)
	val, _ := strconv.Atoi(line[x : lIndex+1])
	positions := make([]point, 0)
	for i := x; i <= lIndex; i++ {
		positions = append(positions, point{y, i})
	}
	p := part{
		positions: positions,
		value:     val,
	}
	return p, lIndex
}

func containsPosition(adjacent []point, positions []point) bool {
	for _, a := range adjacent {
		for _, p := range positions {
			if p.y == a.y && p.x == a.x {
				return true
			}
		}
	}
	return false
}

func getAdjacentPos(y int, x int, lines []string) []point {
	aPos := []point{
		{y - 1, x - 1},
		{y - 1, x},
		{y - 1, x + 1},
		{y, x - 1},
		{y, x + 1},
		{y + 1, x - 1},
		{y + 1, x},
		{y + 1, x + 1},
	}
	for i, pos := range aPos {
		if pos.y > len(lines) || pos.y < 0 || pos.x > len(lines[y])-1 || pos.x < 0 {
			aPos = slices.Delete(aPos, i, 1)
		}
	}
	return aPos
}

func getLastIdxOfPart(ln string, startIdx int) int {
	for i := startIdx; i < len(ln); i++ {
		c := rune(ln[i])
		if !unicode.IsDigit(c) {
			return i - 1
		}
	}
	return len(ln) - 1
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
