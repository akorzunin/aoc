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

func main() {
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
		for _, v := range line {
			if unicode.IsDigit(v) {
				sub_arr = append(sub_arr, string(v))
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
