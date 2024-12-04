package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	arr_2d := [][]string{}
	for scanner.Scan() {
		line := scanner.Bytes()
		arr := make([]string, len(line))
		for num, char := range line {
			arr[num] = string(char)
		}
		arr_2d = append(arr_2d, arr)
	}
	sum := 0
	for _, line := range arr_2d {
		sum += findXmas(line)
	}
	arr_2d_transposed := transpose(arr_2d)
	for _, line := range arr_2d_transposed {
		sum += findXmas(line)
	}
	arr_2d_diagonalized_left := diagonalize(arr_2d, true)
	for _, line := range arr_2d_diagonalized_left {
		sum += findXmas(line)
	}
	arr_2d_diagonalized_right := diagonalize(arr_2d, false)
	for _, line := range arr_2d_diagonalized_right {
		sum += findXmas(line)
	}
	fmt.Println(sum)
}

func findXmas(arr []string) int {
	result := 0
	s := strings.Join(arr, "")
	splits := strings.Split(s, "XMAS")
	result += len(splits) - 1
	slices.Reverse(arr)
	rs := strings.Join(arr, "")
	splits = strings.Split(rs, "XMAS")
	result += len(splits) - 1
	return result
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func diagonalize(slice [][]string, direction bool) [][]string {
	rows := len(slice)
	cols := len(slice[0])

	result := make([][]string, rows+cols-1)

	for i := range result {
		result[i] = make([]string, 0)

		for j := 0; j < rows; j++ {
			k := i - j

			if k >= 0 && k < cols {
				if !direction {
					result[i] = append(result[i], slice[j][k])
				} else {
					result[i] = append(result[i], slice[j][cols-k-1])
				}
			}
		}
	}
	return result
}
