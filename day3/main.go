package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func Read(filepath string) *FileScanner {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	const maxCapacity = 10_000 * 1024 // 20GB == 20_000*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	return &FileScanner{file, scanner}
}

type Point = struct {
	X int
	Y int
}
type SerialNumber = struct {
	Location Point
	IsValid  bool
	Length   int
	Content  []int
}

func main() {
	s := Read("./in.txt")
	defer s.Close()
	var potentialSerialNumbers []SerialNumber
	for i := 0; s.Scan(); i++ {
		line := s.Text()
		serialNumber := &SerialNumber{}
		for j, char := range line {
			if unicode.IsDigit(char) {

				if (j != 0) && unicode.IsDigit(rune(line[j-1])) {
					// add to existing SerialNumber
					serialDigit, _ := strconv.Atoi(string(char))
					serialNumber.Content = append(serialNumber.Content, serialDigit)
					serialNumber.Length++
					continue
				}
				serialNumber := &SerialNumber{Location: Point{X: j, Y: i}}

			}

		}
	}
	a := [][]string{
		{".11", ".12", ".13", ".14", ".15"},
		{".21", "*1", "*2", "*3", ".25"},
		{".31", ".32", ".33", ".34", ".35"},
	}
	sample := a
	for i := 0; i < len(sample); i++ {
		for j := 0; j < len(sample[i]); j++ {
			fmt.Println(sample[i][j])
		}
	}
}
