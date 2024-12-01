package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func main() {
	// file, err := os.Open("test.in.txt")
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var left_arr []int
	var right_arr []int

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")
		v, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left_arr = append(left_arr, v)
		v, err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right_arr = append(right_arr, v)
	}
	slices.Sort(left_arr)
	slices.Sort(right_arr)
	var sum int
	for _, item := range left_arr {
		score := 0
		for _, right_item := range right_arr {
			if item == right_item {
				score++
			}
		}
		sim_score := score * item
		sum += sim_score
	}

	fmt.Println(sum)
}
