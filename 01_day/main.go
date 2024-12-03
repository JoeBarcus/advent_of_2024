package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	data := readText("input.txt")

	list1, list2 := getLists(data)

	sort.Ints(list1)
	sort.Ints(list2)

	total := deltaLists(list1, list2)
	fmt.Println(total)
}

func deltaLists(list1 []int, list2 []int) int {
	var totals int
	for i := range list1 {
		sum := math.Abs(float64(list1[i]) - float64(list2[i]))
		totals += int(sum)
	}
	return totals
}

func getLists(data []string) ([]int, []int) {
	var list1 []int
	var list2 []int

	for i := range data {
		row := data[i]
		fmt.Println(row)
		num1, err1 := strconv.Atoi(string(row[0:5]))
		num2, err2 := strconv.Atoi(string(row[8:13]))
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing integers:", err1, err2)
			continue
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}

func readText(fileName string) []string {
	file, err := os.Open(fileName)
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
