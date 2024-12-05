package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileData := readFile("input.txt")

	totals := 0
	for i := range fileData {
		multipliers := findPattern(fileData[i])
		totals += findTotals(multipliers)
	}
	fmt.Println(totals)

	totals2 := 0
	dont := false
	for j := range fileData {
		multipliers2 := findPattern2(fileData[j])
		totals2 += findTotals2(multipliers2, &dont)
	}
	fmt.Println(totals2)

}

func findTotals2(multipliers []string, dont *bool) int {
	re := regexp.MustCompile(`.*?(\d+),(\d+).*?`)

	productTotals := 0
	for j := range multipliers {
		if multipliers[j] == "don't()" {
			*dont = true
			continue
		}

		if multipliers[j] == "do()" {
			*dont = false
			continue
		}

		if !*dont {
			matches := re.FindStringSubmatch(multipliers[j])
			left, err := strconv.Atoi(matches[1])
			if err != nil {
				fmt.Println("error converting:", err)
			}
			right, err := strconv.Atoi(matches[2])
			if err != nil {
				fmt.Println("error converting:", err)
			}
			productTotals += left * right

		}
	}

	return productTotals
}

func findTotals(multipliers []string) int {
	re := regexp.MustCompile(`.*?(\d+),(\d+).*?`)

	productTotals := 0
	for j := range multipliers {
		matches := re.FindStringSubmatch(multipliers[j])
		left, err := strconv.Atoi(matches[1])
		if err != nil {
			fmt.Println("error converting:", err)
		}
		right, err := strconv.Atoi(matches[2])
		if err != nil {
			fmt.Println("error converting:", err)
		}
		productTotals += left * right
	}

	return productTotals
}

func findPattern2(fileRow string) []string {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(((\d+),)*(\d+)\)`)
	matches := re.FindAllString(fileRow, -1)
	return matches
}

func findPattern(fileRow string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(fileRow, -1)
	return matches
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
