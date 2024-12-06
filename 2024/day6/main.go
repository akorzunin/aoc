package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Point struct {
	x, y int
}

type Guard struct {
	pos    Point
	dir    Point
	facing string
}

type State struct {
	pos    Point
	facing string
}

func main() {
	// file, err := os.Open("test.in.txt")
	file, err := os.Open("in.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	var guard Guard
	directions := []rune("^>v<")

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []rune(line)
		for x, ch := range row {
			if slices.Contains(directions, ch) {
				guard = Guard{Point{x, y}, dirFromChar(ch), string(ch)}
				row[x] = '.'
			}
		}
		grid = append(grid, row)
	}

	loopCount := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '.' && (Point{x, y} != guard.pos) {
				if causesLoop(grid, guard, Point{x, y}) {
					loopCount++
				}
			}
		}
	}
	fmt.Println(loopCount)
}

func causesLoop(grid [][]rune, initialGuard Guard, obstacle Point) bool {
	newGrid := make([][]rune, len(grid))
	for i := range grid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	newGrid[obstacle.y][obstacle.x] = '#'

	guard := initialGuard
	visited := make(map[State]bool)

	for {
		if !isInGrid(guard.pos, newGrid) {
			return false
		}

		state := State{guard.pos, guard.facing}
		if visited[state] {
			return true
		}
		visited[state] = true

		next := Point{guard.pos.x + guard.dir.x, guard.pos.y + guard.dir.y}
		if !isInGrid(next, newGrid) {
			return false
		} else if newGrid[next.y][next.x] == '#' {
			guard.turnRight()
		} else {
			guard.move()
		}
	}
}

func isInGrid(p Point, grid [][]rune) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[0])
}

func (g *Guard) turnRight() {
	switch g.facing {
	case "^":
		g.dir = Point{1, 0}
		g.facing = ">"
	case ">":
		g.dir = Point{0, 1}
		g.facing = "v"
	case "v":
		g.dir = Point{-1, 0}
		g.facing = "<"
	case "<":
		g.dir = Point{0, -1}
		g.facing = "^"
	}
}

func (g *Guard) move() {
	g.pos.x += g.dir.x
	g.pos.y += g.dir.y
}

func dirFromChar(ch rune) Point {
	switch ch {
	case '^':
		return Point{0, -1}
	case 'v':
		return Point{0, 1}
	case '<':
		return Point{-1, 0}
	case '>':
		return Point{1, 0}
	}
	return Point{0, 0}
}
