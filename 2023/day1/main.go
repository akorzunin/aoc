package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

type DigitArr = [10]string

func getIndex(arr DigitArr, index string) int {
	predicate := func(i int) bool { return arr[i] == index }
	for i := 0; i < len(arr); i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func main() {
	digits := DigitArr{"_", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	file, err := os.Open("./in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	const maxCapacity = 10_000 * 1024 // 20GB == 20_000*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	var arr []int
	for scanner.Scan() {
		line := scanner.Text()
		var sub_arr []string
		for i, v := range line {
			cutted_line := line[i:]
			if unicode.IsDigit(v) {
				sub_arr = append(sub_arr, string(v))
			}
			for _, digit := range digits {
				if strings.HasPrefix(cutted_line, digit) {
					parsed_digit := getIndex(digits, digit)
					sub_arr = append(sub_arr, strconv.Itoa(parsed_digit))
				}

			}
		}
		if len(sub_arr) == 0 {
			continue
		}
		var cal_value = []string{sub_arr[0], sub_arr[len(sub_arr)-1]}
		cal_int, _ := strconv.Atoi(strings.Join(cal_value, ""))
		arr = append(arr, cal_int)
	}
	sum_val := sum(arr)
	fmt.Println(sum_val)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("out.txt", []byte(fmt.Sprintf("%d", sum_val)), 0644)
	if err != nil {
		panic(err)
	}
}
