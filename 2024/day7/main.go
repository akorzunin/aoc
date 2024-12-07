package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalCalibration := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(parts[0])
		numbers := strings.Split(parts[1], " ")

		if canProduceTestValue(numbers, testValue) {
			totalCalibration += testValue
		}
	}

	fmt.Println("Total calibration result:", totalCalibration)
}

func canProduceTestValue(numbers []string, testValue int) bool {
	nums := make([]int, len(numbers))
	for i, num := range numbers {
		nums[i], _ = strconv.Atoi(num)
	}

	return tryOperators(nums, 0, nums[0], testValue)
}

func tryOperators(nums []int, index int, currentValue int, testValue int) bool {
	if index == len(nums)-1 {
		return currentValue == testValue
	}

	// Try addition
	if tryOperators(nums, index+1, currentValue+nums[index+1], testValue) {
		return true
	}

	// Try multiplication
	if tryOperators(nums, index+1, currentValue*nums[index+1], testValue) {
		return true
	}

	// Try concatenation
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, nums[index+1]))
	if tryOperators(nums, index+1, concatenated, testValue) {
		return true
	}

	return false
}
