package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/oleiade/reflections"
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

type GameId = int

func parseGameId(data string) GameId {
	var digits []string
	for _, v := range data {
		if unicode.IsDigit(v) {
			digits = append(digits, string(v))
		}
	}
	gameId, _ := strconv.Atoi(strings.Join(digits, ""))
	return GameId(gameId)
}

type GameResult = struct {
	gameId     GameId
	redRolls   []Balls
	greenRolls []Balls
	blueRolls  []Balls
}
type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

type Balls = struct {
	color Color
	count int
}

func parseGameRolls(data string) [][]Balls {
	var rolls [][]Balls
	for _, game := range strings.Split(data, "; ") {
		var ballSlice []Balls
		for _, balls := range strings.Split(game, ", ") {
			f := strings.SplitN(balls, " ", 2)
			count, color := f[0], f[1]
			int_count, _ := strconv.Atoi(count)
			bl := Balls{color: Color(color), count: int_count}
			ballSlice = append(ballSlice, bl)
		}

		rolls = append(rolls, ballSlice)
	}
	return rolls
}
func parseGame(line string) *GameResult {
	parts := strings.SplitN(line, ": ", 2)
	gameId := parseGameId(parts[0])
	gameResult := &GameResult{gameId: gameId}
	gameRolls := parseGameRolls(parts[1])
	for _, roll := range gameRolls {
		for _, balls := range roll {
			switch balls.color {
			case Red:
				gameResult.redRolls = append(gameResult.redRolls, balls)
			case Green:
				gameResult.greenRolls = append(gameResult.greenRolls, balls)
			case Blue:
				gameResult.blueRolls = append(gameResult.blueRolls, balls)
			}
		}
	}
	return gameResult
}

type GameRule = struct {
	RedNum   int
	GreenNum int
	BlueNum  int
}
type Pair[T, U any] struct {
	First  T
	Second U
}

func Zip[T, U any](ts []T, us []U) []Pair[T, U] {
	if len(ts) != len(us) {
		panic("slices have different length")
	}
	pairs := make([]Pair[T, U], len(ts))
	for i := 0; i < len(ts); i++ {
		pairs[i] = Pair[T, U]{ts[i], us[i]}
	}
	return pairs
}
func getColorCount(gameRule GameRule, color Color) int {
	switch color {
	case Red:
		return gameRule.RedNum
	case Green:
		return gameRule.GreenNum
	case Blue:
		return gameRule.BlueNum
	}
	panic("color not found")
}
func checkGameRule(gameResult *GameResult, gameRule GameRule) bool {
	sum := 0
	rolls := [][]Balls{gameResult.redRolls, gameResult.greenRolls, gameResult.blueRolls}
	colors := []Color{Red, Green, Blue}
	for _, pair := range Zip(rolls, colors) {
		colorRoll, color := pair.First, pair.Second
		for _, roll := range colorRoll {
			if roll.count > sum {
				sum = roll.count
			}
			if sum > getColorCount(gameRule, color) {
				fmt.Printf(
					"Game %v not valid. Reason: %d %s balls.\n",
					gameResult.gameId,
					roll.count,
					color,
				)
				return false
			}
		}
	}
	return true
}
func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
func getFieldName(color Color) string {
	switch color {
	case Red:
		return "RedNum"
	case Green:
		return "GreenNum"
	case Blue:
		return "BlueNum"
	}
	panic("color not found")
}
func getPower(gameRule GameRule) int {
	sum := 1
	for _, fieldName := range [...]string{"RedNum", "GreenNum", "BlueNum"} {
		value, _ := reflections.GetField(gameRule, fieldName)
		sum *= value.(int)
	}
	return sum
}
func geGamePower(gameResult *GameResult) int {
	rolls := [][]Balls{gameResult.redRolls, gameResult.greenRolls, gameResult.blueRolls}
	colors := []Color{Red, Green, Blue}
	sum := GameRule{}
	for _, pair := range Zip(rolls, colors) {
		colorRoll, color := pair.First, pair.Second
		for _, roll := range colorRoll {
			fieldName := getFieldName(color)
			f, err := reflections.GetField(sum, fieldName)
			if err != nil {
				log.Fatal(err)
			}
			if roll.count > f.(int) {
				_ = reflections.SetField(&sum, fieldName, roll.count)
			}
		}
	}
	pow := getPower(sum)
	fmt.Printf("GameId %d, %v, pow: %d\n", gameResult.gameId, sum, pow)
	return pow

}
func main() {
	s := Read("./in.txt")
	defer s.Close()
	var validGameIds []GameId
	for s.Scan() {
		line := s.Text()
		gameData := parseGame(line)
		power := geGamePower(gameData)
		validGameIds = append(validGameIds, power)
	}
	res := sum(validGameIds)
	fmt.Printf("%+v\n", res)
	if err := os.WriteFile("out.txt", []byte(fmt.Sprintf("%d", res)), 0644); err != nil {
		panic(err)
	}
}
