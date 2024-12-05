package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseRules(rules []string) map[int][]int {
	dependencies := map[int][]int{}
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		dependencies[before] = append(dependencies[before], after)
	}
	return dependencies
}

func isValidOrder(pages []int, rules map[int][]int) bool {
	pagePositions := map[int]int{}
	for i, page := range pages {
		pagePositions[page] = i
	}

	for before, afters := range rules {
		if _, exists := pagePositions[before]; !exists {
			continue
		}

		for _, after := range afters {
			if _, exists := pagePositions[after]; !exists {
				continue
			}

			if pagePositions[before] > pagePositions[after] {
				return false
			}
		}
	}
	return true
}

func sortPages(pages []int, rules map[int][]int) []int {
	sorted := make([]int, len(pages))
	copy(sorted, pages)

	sort.Slice(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]

		// Check direct dependencies
		for _, after := range rules[a] {
			if after == b {
				return true
			}
		}
		for _, after := range rules[b] {
			if after == a {
				return false
			}
		}

		// Check transitive dependencies
		visited := map[int]bool{}
		var dfs func(int, int) bool
		dfs = func(current, target int) bool {
			if current == target {
				return true
			}
			visited[current] = true
			for _, next := range rules[current] {
				if !visited[next] {
					if dfs(next, target) {
						return true
					}
				}
			}
			return false
		}

		if dfs(a, b) {
			return true
		}
		if dfs(b, a) {
			return false
		}

		return a < b
	})

	return sorted
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := []string{}
	updates := [][]int{}
	parsingRules := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			rules = append(rules, line)
		} else {
			update := []int{}
			for _, numStr := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(numStr)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	dependencies := parseRules(rules)
	sum := 0

	for _, update := range updates {
		if !isValidOrder(update, dependencies) {
			sorted := sortPages(update, dependencies)
			middleIndex := len(sorted) / 2
			sum += sorted[middleIndex]
		}
	}

	fmt.Println(sum)
}
