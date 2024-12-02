package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("test.in.txt")
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safe_levels := 0
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		parsed_levels := []int{}
		for _, level := range levels {
			i_level, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			parsed_levels = append(parsed_levels, i_level)
		}
		fmt.Println("[IN]: ", parsed_levels)
		isSafe := isFloorSafe(parsed_levels)
		fmt.Println("> ", isSafe, parsed_levels)
		fmt.Println()
		if isSafe {
			safe_levels++
		}
	}
	fmt.Println(safe_levels)
}

func checkAscending(levels []int, tolerateUnsafe bool) bool {
	prevLevel := levels[0]
	for num, level := range levels {
		if num == 0 {
			continue
		}
		if level <= prevLevel {
			fmt.Println("[FAIL]: ", prevLevel, " >= ", level)
			return tolerated(levels, num, tolerateUnsafe, checkAscending)
		}
		if (level - prevLevel) > 3 {
			fmt.Println("[FAIL]: ", level, " - ", prevLevel, "diff: ", prevLevel-level)
			return tolerated(levels, num, tolerateUnsafe, checkAscending)
		}
		prevLevel = level
	}
	if !tolerateUnsafe {
		fmt.Println("rescued", levels)
	}
	return true
}

func tolerated(levels []int, _ int, tolerateUnsafe bool, _ func([]int, bool) bool) bool {
	if tolerateUnsafe {
		for i := 0; i < len(levels); i++ {
			if checkDescending(removeIndex(levels, i), false) || checkAscending(removeIndex(levels, i), false) {
				return true
			}
		}
	}
	fmt.Println("non recoverable", levels)
	return false
}

func checkDescending(levels []int, tolerateUnsafe bool) bool {
	prevLevel := levels[0]
	for num, level := range levels {
		if num == 0 {
			continue
		}
		if level >= prevLevel {
			fmt.Println("[FAIL]: ", prevLevel, " <= ", level)
			return tolerated(levels, num, tolerateUnsafe, checkDescending)
		}
		if (prevLevel - level) > 3 {
			fmt.Println("[FAIL]: ", level, " - ", prevLevel, "diff: ", prevLevel-level)
			return tolerated(levels, num, tolerateUnsafe, checkDescending)
		}
		prevLevel = level
	}
	if !tolerateUnsafe {
		fmt.Println("rescued", levels)
	}
	return true
}

func removeIndex(levels []int, index int) []int {
	tmp := make([]int, len(levels))
	copy(tmp, levels)
	tmp = append(tmp[:index], tmp[index+1:]...)
	return tmp
}

func isFloorSafe(levels []int) bool {
	if levels[0] > levels[1] {
		return checkDescending(levels, true)
	} else {
		return checkAscending(levels, true)
	}
}
