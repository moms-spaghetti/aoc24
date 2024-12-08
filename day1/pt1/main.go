package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	documentCount = 1000
	separator     = "   "
)

func main() {
	file, err := os.Open("data.txt")
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

	sort.SliceStable(numbersA, func(i, j int) bool {
		return numbersA[i] < numbersA[j]
	})

	sort.SliceStable(numbersB, func(i, j int) bool {
		return numbersB[i] < numbersB[j]
	})

	for i := 0; i < len(numbersA); i++ {
		diff := numbersA[i] - numbersB[i]
		if diff < 0 {
			diff = numbersB[i] - numbersA[i]
		}
		result += diff
	}

	fmt.Println(result)
}
