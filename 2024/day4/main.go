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
	// itrate slices 3x3
	for i := 0; i < len(arr_2d)-2; i++ {
		for j := 0; j < len(arr_2d[i])-2; j++ {
			a := [][]string{
				arr_2d[i][j : j+3],
				arr_2d[i+1][j : j+3],
				arr_2d[i+2][j : j+3],
			}
			sum += findX_mas(a)
		}
	}

	fmt.Println(sum)
}

func findXmas(arr []string) int {
	word := "MAS"
	result := 0
	s := strings.Join(arr, "")
	splits := strings.Split(s, word)
	result += len(splits) - 1
	slices.Reverse(arr)
	rs := strings.Join(arr, "")
	splits = strings.Split(rs, word)
	result += len(splits) - 1
	return result
}

func findX_mas(arr [][]string) int {
	// accept window 3x3
	if len(arr) != 3 || len(arr[0]) != 3 {
		panic("invalid window size")
	}
	arr_left := diagonalize(arr, true)
	is_mas_left := 0
	for _, line := range arr_left {
		is_mas_left += findXmas(line)
	}
	arr_right := diagonalize(arr, false)
	is_mas_right := 0
	for _, line := range arr_right {
		is_mas_right += findXmas(line)
	}
	if (is_mas_left + is_mas_right) == 2 {
		return 1
	}
	return 0
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
