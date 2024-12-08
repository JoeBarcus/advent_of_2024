package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coords struct {
	x     int
	y     int
	value string
}

func main() {
	data := readFile("input.txt")
	coords := buildMap(data)
	totals := findXmas(coords)
	fmt.Println(totals)
}

func findXmas(coords []Coords) int {
	totals := 0
	for i := range coords {
		if coords[i].value == "X" {
			totals += checkForward(coords[i], coords)
			totals += checkBackward(coords[i], coords)
			totals += checkUp(coords[i], coords)
			totals += checkDown(coords[i], coords)
			totals += checkDownRight(coords[i], coords)
			totals += checkDownLeft(coords[i], coords)
			totals += checkUpLeft(coords[i], coords)
			totals += checkUpRight(coords[i], coords)
		}
	}

	return totals
}

func checkForward(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x+1, coord.y)
	aCoord := filterCoords(coords, coord.x+2, coord.y)
	sCoord := filterCoords(coords, coord.x+3, coord.y)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func checkBackward(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x-1, coord.y)
	aCoord := filterCoords(coords, coord.x-2, coord.y)
	sCoord := filterCoords(coords, coord.x-3, coord.y)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func checkUp(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x, coord.y+1)
	aCoord := filterCoords(coords, coord.x, coord.y+2)
	sCoord := filterCoords(coords, coord.x, coord.y+3)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func checkDown(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x, coord.y-1)
	aCoord := filterCoords(coords, coord.x, coord.y-2)
	sCoord := filterCoords(coords, coord.x, coord.y-3)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func checkDownRight(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x+1, coord.y-1)
	aCoord := filterCoords(coords, coord.x+2, coord.y-2)
	sCoord := filterCoords(coords, coord.x+3, coord.y-3)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func checkDownLeft(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x-1, coord.y-1)
	aCoord := filterCoords(coords, coord.x-2, coord.y-2)
	sCoord := filterCoords(coords, coord.x-3, coord.y-3)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func checkUpLeft(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x-1, coord.y+1)
	aCoord := filterCoords(coords, coord.x-2, coord.y+2)
	sCoord := filterCoords(coords, coord.x-3, coord.y+3)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func checkUpRight(coord Coords, coords []Coords) int {
	mCoord := filterCoords(coords, coord.x+1, coord.y+1)
	aCoord := filterCoords(coords, coord.x+2, coord.y+2)
	sCoord := filterCoords(coords, coord.x+3, coord.y+3)

	if mCoord == "M" && aCoord == "A" && sCoord == "S" {
		return 1
	}

	return 0
}

func filterCoords(coords []Coords, x int, y int) string {
	for _, coord := range coords {
		if coord.x == x && coord.y == y {
			return coord.value
		}
	}

	return "none"
}

func buildMap(data []string) []Coords {
	var xCoords []Coords
	for i := range data {
		for j := range data[i] {
			xLoc := Coords{
				x:     j,
				y:     i,
				value: string(data[i][j]),
			}
			xCoords = append(xCoords, xLoc)
		}
	}

	return xCoords
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
