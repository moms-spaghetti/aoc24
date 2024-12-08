package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	documentCount = 1000
	separator     = "   "
)

func main() {
	file, err := os.Open("../pt1/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		numbersA = make([]int, documentCount)
		numbersB = make([]int, documentCount)
		i        int
		result   int
	)

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), separator)
		if len(numbers) != 2 {
			panic("numbers missing 2 count")
		}

		firstNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		numbersA[i] = firstNumber

		secondNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		numbersB[i] = secondNumber

		i++
	}

	for _, numberA := range numbersA {
		var countOfAInB int
		for _, numberB := range numbersB {
			if numberA == numberB {
				countOfAInB++
			}
		}

		result += numberA * countOfAInB
	}

	fmt.Println(result)
}
