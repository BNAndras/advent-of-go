package main

import (
	"advent-of-go/utils/files"
	"sort"
	"strings"
)

type counter map[string]int
type rules map[string]string

func (c counter) Counts() []int {
	var counts []int
	for _, count := range c {
		counts = append(counts, count)
	}
	return counts
}

func main() {
	input := files.ReadFile(14, "\n")
	println(solvePartOne(input))
	println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	polymer := input[0]

	elementCounter := countCharsIn(polymer)
	pairCounter := countPairsIn(polymer)

	rules := getRuleMap(input[2:])

	for i := 0; i < 10; i++ {
		newPairCounter := make(counter)
		for pair, count := range pairCounter {
			inserted := rules[pair]
			elementCounter[inserted] += count

			pairBefore := string(pair[0]) + inserted
			newPairCounter[pairBefore] += count

			pairAfter := inserted + string(pair[1])
			newPairCounter[pairAfter] += count
			// fmt.Println(newPairCounter)
		}

		pairCounter = newPairCounter
	}

	counts := elementCounter.Counts()
	sort.Ints(counts)

	max := counts[len(counts)-1]
	min := counts[0]
	return max - min
}

func solvePartTwo(input []string) int {
	polymer := input[0]

	elementCounter := countCharsIn(polymer)
	pairCounter := countPairsIn(polymer)

	rules := getRuleMap(input[2:])

	for i := 0; i < 40; i++ {
		newPairCounter := make(counter)
		for pair, count := range pairCounter {
			inserted := rules[pair]
			elementCounter[inserted] += count

			pairBefore := string(pair[0]) + inserted
			newPairCounter[pairBefore] += count

			pairAfter := inserted + string(pair[1])
			newPairCounter[pairAfter] += count
		}

		pairCounter = newPairCounter
	}

	counts := elementCounter.Counts()
	sort.Ints(counts)

	max := counts[len(counts)-1]
	min := counts[0]
	return max - min
}

func getRuleMap(rules []string) rules {
	ruleMap := make(map[string]string)
	for _, line := range rules {
		sections := strings.Split(line, " -> ")

		pair := sections[0]
		insertion := sections[1]
		ruleMap[pair] = insertion
	}

	return ruleMap
}

func countPairsIn(s string) counter {
	counts := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		counts[pair]++
	}

	return counts
}

func countCharsIn(s string) counter {
	counts := make(map[string]int)
	for _, char := range s {
		counts[string(char)]++
	}

	return counts
}
