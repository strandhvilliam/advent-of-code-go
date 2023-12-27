package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	inputFile   string = "input.txt"
	exampleFile string = "example.txt"
)

func main() {
	part1Result := part1(exampleFile)
	part2Result := part2(exampleFile)

	fmt.Println("-----DAY 4------")
	fmt.Println("Part 1: ", part1Result)
	fmt.Println("Part 2: ", part2Result)
}

func part1(file string) int {
	lines, err := readInput(file)
	if err != nil {
		fmt.Println(err)
	}

	total := 0
	for _, line := range lines {
		wins, drawn := parseCard(line)
		total += calcPoints(wins, drawn)
	}
	return total
}

func part2(file string) int {
	lines, err := readInput(file)
	if err != nil {
		fmt.Println(err)
	}

	cards := make(map[int]int)

	for i, line := range lines {
		cards[i]++
		for k := 0; k < cards[i]; k++ {
			updateCardCounts(cards, i, line)
		}
	}

	result := sumCards(cards)
	return result
}

func calcPoints(wins []string, drawn []string) int {
	points := 0
	for _, card := range wins {
		if slices.Contains(drawn, card) {
			points *= 2
			if points == 0 {
				points = 1
			}
		}
	}
	return points
}

func updateCardCounts(cards map[int]int, idx int, line string) {
	wins, drawn := parseCard(line)
	count := intersections(wins, drawn)
	for j := idx; j < count+idx; j++ {
		cards[j+1]++
	}
}

func sumCards(cards map[int]int) int {
	sum := 0
	for _, v := range cards {
		sum += v
	}
	return sum
}

func intersections(wins []string, drawn []string) int {
	count := 0
	for _, w := range wins {
		if slices.Contains(drawn, w) {
			count++
		}
	}
	return count
}

func parseCard(line string) ([]string, []string) {
	fmt.Println("LINE: ", line)
	fields := strings.Fields(line)[2:]
	separator := slices.Index(fields, "|")
	wins := fields[:separator]
	drawn := fields[separator+1:]
	fmt.Println("WINS: ", wins)
	fmt.Println("DRAWN: ", drawn)
	fmt.Println()
	return wins, drawn
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
