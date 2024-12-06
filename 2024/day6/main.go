package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []rune(line)
		for x, ch := range row {
			if ch == '^' || ch == 'v' || ch == '<' || ch == '>' {
				guard = Guard{Point{x, y}, dirFromChar(ch), string(ch)}
				row[x] = '.'
			}
		}
		grid = append(grid, row)
	}

	visited := make(map[Point]bool)
	states := make(map[State]bool)
	visited[guard.pos] = true

	// prindGeridWithGuard(grid, guard, visited)
	for {
		if !isInGrid(guard.pos, grid) {
			break
		}

		state := State{guard.pos, guard.facing}
		if states[state] {
			break // We've been here before with the same direction, so we're in a loop
		}
		states[state] = true

		next := Point{guard.pos.x + guard.dir.x, guard.pos.y + guard.dir.y}
		if !isInGrid(next, grid) {
			break
		} else if grid[next.y][next.x] == '#' {
			guard.turnRight()
		} else {
			guard.move()
			visited[guard.pos] = true
		}
		// prindGeridWithGuard(grid, guard, visited)
	}

	fmt.Println(len(visited))
}

func isVisited(p Point, visited map[Point]bool) bool {
	_, exists := visited[p]
	return exists
}

func prindGeridWithGuard(grid [][]rune, guard Guard, visited map[Point]bool) {
	for y, line := range grid {
		for x, ch := range line {
			if guard.pos.x == x && guard.pos.y == y {
				fmt.Printf("%s", guard.facing)
			} else if isVisited(Point{x, y}, visited) {
				fmt.Printf("X")
			} else {
				fmt.Printf("%c", ch)
			}
		}
		fmt.Println()
	}
	fmt.Println()
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
