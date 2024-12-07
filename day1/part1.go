package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		a, err := strconv.Atoi(numbers[0])
		b, err2 := strconv.Atoi(numbers[1])
		if err != nil || err2 != nil {
			panic(err)
		}
		left = append(left, a)
		right = append(right, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += abs(left[i] - right[i])
	}
	fmt.Println(sum)
}
