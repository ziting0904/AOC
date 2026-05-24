package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getRepeatedInt(start, end string) []int {
	var result []int
	var startPos, endPos, length int
	var err error
	if len(start)%2 == 1 && len(end)%2 == 1 {
		return result
	}

	if len(start)%2 == 0 && len(end)%2 == 1 {
		startPos, err = strconv.Atoi(start)
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		length = len(start)
		endPos = int(math.Pow10(length))
	}

	if len(start)%2 == 1 && len(end)%2 == 0 {
		endPos, err = strconv.Atoi(end)
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		length = len(end)
		startPos = int(math.Pow10(length - 1))
	}

	if len(start)%2 == 0 && len(end)%2 == 0 {
		fmt.Println("Both start and end values have the same number of digits!")
		startPos, err = strconv.Atoi(start)
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		endPos, err = strconv.Atoi(end)
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		length = len(start)
	}

	//fmt.Printf("startPos: %d, endPos: %d\n, length: %d\n", startPos, endPos, length)
	for i := startPos; i < endPos; i++ {
		divisor := int(math.Pow10(length / 2))
		firstHalf := i / divisor
		secondHalf := i % divisor
		if firstHalf == secondHalf {
			fmt.Println("Found invalid Id: ", i)
			result = append(result, i)
		}

		for j := length/2 - 1; j > 0; j-- {
			if length%j == 0 {

			}

		}
	}

	return result
}

func regexMatcher(start, end string) []int {
	var err error
	var result []int
	startRange, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}
	endRange, err := strconv.Atoi(end)
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}

	for i := startRange; i < endRange; i++ {
		currentStr := strconv.Itoa(i)
		lengthStr := len(currentStr)
		for j := lengthStr / 2; j > 0; j-- {
			if len(currentStr)%j == 0 {
				matchesNo := len(currentStr) / j
				subStr := currentStr[:j]
				re := regexp.MustCompile(subStr)
				matches := re.FindAllString(currentStr, matchesNo)
				if len(matches) == matchesNo {
					result = append(result, i)
					break
				}
			}
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	strList := strings.Split(text, ",")

	var result []int
	resultSum := 0
	for _, str := range strList {
		//fmt.Printf("Current range: %s\n", str)
		values := strings.Split(str, "-")
		repeatedInts := regexMatcher(values[0], values[1])
		result = append(result, repeatedInts...)
	}
	//fmt.Println(strList[0])

	for _, value := range result {
		resultSum += value
	}
	fmt.Println(resultSum)
}
