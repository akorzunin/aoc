package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	scanner.Split(bufio.ScanWords)
	sum := 0
	enable := true
	for scanner.Scan() {
		word := scanner.Text()
		pattern := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)

		matches := pattern.FindAllString(word, -1)

		fmt.Println("Matches:")
		for _, match := range matches {
			if match == "do()" {
				enable = true
				continue
			} else if match == "don't()" {
				enable = false
				continue
			}
			if !enable {
				continue
			}
			sum += getMultiplier(match)
			fmt.Println(match)
		}
	}
	fmt.Println(sum)
}

func getMultiplier(match string) int {
	_, m, _ := strings.Cut(match, "mul(")
	m, _, _ = strings.Cut(m, ")")
	m = strings.TrimSpace(m)
	nums := strings.Split(m, ",")
	num1, num2 := nums[0], nums[1]
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	return a * b
}
