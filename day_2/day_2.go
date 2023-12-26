package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	exampleFile string = "example.txt"
	inputFile   string = "input.txt"
)

type set struct {
	reds   int
	greens int
	blues  int
}

type game struct {
	sets []set
	id   int
}

func main() {
	part1Result, err := part1(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	part2Result, err := part2(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----DAY 2------")
	fmt.Println("Part 1: ", part1Result)
	fmt.Println("Part 2: ", part2Result)
}

func part1(inputFile string) (int, error) {
	lines, err := readInput(inputFile)
	if err != nil {
		return -1, err
	}

	games := parseGames(lines)
	filtered := filterPossibleGames(games)
	result := sumGameIds(filtered)
	return result, nil
}

func part2(inputFile string) (int, error) {
	lines, err := readInput(inputFile)
	if err != nil {
		return -1, err
	}

	games := parseGames(lines)

	result := 0
	for _, g := range games {
		min := getMinCubes(g.sets)
		pw := powerOfMins(min)
		result += pw
	}
	return result, nil
}

func parseGames(lines []string) []game {
	var games []game

	for _, line := range lines {
		games = append(games, parseIntoGameStruct(line))
	}
	return games
}

func powerOfMins(s set) int {
	return s.reds * s.greens * s.blues
}

func getMinCubes(sets []set) set {
	mins := set{
		reds:   0,
		blues:  0,
		greens: 0,
	}
	for _, s := range sets {
		if s.reds > mins.reds {
			mins.reds = s.reds
		}
		if s.greens > mins.greens {
			mins.greens = s.greens
		}
		if s.blues > mins.blues {
			mins.blues = s.blues
		}
	}
	return mins
}

func sumGameIds(games []game) int {
	sum := 0
	for _, g := range games {
		sum += g.id
	}
	return sum
}

func filterPossibleGames(games []game) []game {
	result := make([]game, 0)

	for _, g := range games {
		valid := true
		for _, s := range g.sets {
			if !isPossibleSet(s) {
				valid = false
			}
		}
		if valid {
			result = append(result, g)
		}
	}

	return result
}

func isPossibleSet(s set) bool {
	return s.reds <= 12 && s.greens <= 13 && s.blues <= 14
}

func parseIntoGameStruct(ln string) game {
	split := strings.Split(ln, ":")
	id, _ := strconv.Atoi(strings.Replace(split[0], "Game ", "", 1))
	sets := parseSets(strings.TrimSpace(split[1]))
	return game{
		sets: sets,
		id:   id,
	}
}

func parseSets(data string) []set {
	split := strings.Split(data, ";")

	var sets []set

	for _, str := range split {
		set := set{
			reds:   countColor(str, "red"),
			blues:  countColor(str, "blue"),
			greens: countColor(str, "green"),
		}

		sets = append(sets, set)
	}
	return sets
}

func countColor(data string, color string) int {
	split := strings.Split(data, ",")

	for _, str := range split {
		idx := strings.Index(str, color)
		if idx == -1 {
			continue
		}
		count, _ := strconv.Atoi(strings.TrimSpace(str[:idx]))
		return count
	}
	return 0
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
