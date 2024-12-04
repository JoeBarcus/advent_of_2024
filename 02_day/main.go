package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readText("input.txt")

	totals := 0
	for i := range data {
		valid := checkRow(data[i])
		if valid {
			totals += 1
		}
	}
	fmt.Println(totals)
}

func checkRow(rowData string) bool {
	stringSlice := strings.Fields(rowData)

	integerSlice := convertToInts(stringSlice)

	var ascending bool
	if integerSlice[0] < integerSlice[1] {
		ascending = true
	} else if integerSlice[0] > integerSlice[1] {
		ascending = false
	} else {
		return false
	}

	if !ascending {
		reverse(integerSlice)
	}

	var previous int
	for i := range integerSlice {
		if i == 0 {
			previous = integerSlice[i]
			continue
		}
		if integerSlice[i] <= previous || integerSlice[i]-previous >= 4 {
			return false
		}
		previous = integerSlice[i]
	}
	return true
}

func reverse(slice []int) {
	left, right := 0, len(slice)-1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}

func convertToInts(stringSlice []string) []int {
	var nums []int

	for _, part := range stringSlice {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func readText(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
