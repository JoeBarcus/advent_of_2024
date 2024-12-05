package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readText("input_test.txt")

	totals := 0
	for i := range data {
		valid := checkRow(data[i])
		if valid {
			totals += 1
		}
	}
	fmt.Println("Part 1:", totals)

	totals2 := 0
	for i := range data {
		valid := checkRow2(data[i])
		if valid {
			totals2 += 1
		}
	}
	fmt.Println("Part 2:", totals2)

}

func checkRow2(rowData string) bool {
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
	allFiveGood := true
	for i := range integerSlice {
		if i == 0 {
			previous = integerSlice[i]
			continue
		}
		if integerSlice[i] <= previous || integerSlice[i]-previous >= 4 {
			allFiveGood = false
		}
		previous = integerSlice[i]
	}

	if allFiveGood {
		return true
	}

	strikes := 0
	var integerArray [5]int
	copy(integerArray[:], integerSlice)
	for j := range integerArray {
		filteredSlice := append(integerArray[:j], integerArray[j+1:]...)
		copy(integerArray[:], integerSlice)

		var previous int
		for i := range filteredSlice {
			if i == 0 {
				previous = filteredSlice[i]
				continue
			}
			fmt.Println("Previous", previous)
			fmt.Println("Filtered", filteredSlice[i])
			if filteredSlice[i] <= previous || filteredSlice[i]-previous >= 4 {
				fmt.Println("Here")
				strikes++
			}
			previous = filteredSlice[i]
		}

		filteredSlice = nil
	}

	if strikes > 2 {
		return false
	}
	return true
}

func remove(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Index out of range")
		return slice
	}
	newSlice := append(slice[:i], slice[i+1:]...)
	return newSlice
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
