package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(numbers []int) int {
	s := 0

	for _, n := range numbers {
		s += n
	}

	return s
}

func twice(numbers []int) int {
	dict := map[int]bool{0: true}
	sum := 0

	for {
		for _, n := range numbers {
			sum += n

			if _, ok := dict[sum]; ok {
				return sum
			}

			dict[sum] = true
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	fmt.Println("Result 1: ", sum(numbers))
	fmt.Println("Result 2: ", twice(numbers))
}
