package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func largestSubsequence(s string, length int) string {
	var stack []rune
	
	for i, char := range s {
		// Pop smaller elements from the stack if we have enough characters 
		// remaining in the original string to still reach our target length
		for len(stack) > 0 && stack[len(stack)-1] < char && (len(stack)-1+(len(s)-i)) >= length {
			stack = stack[:len(stack)-1]
		}
		
		// Push the current character if the stack hasn't reached the target length
		if len(stack) < length {
			stack = append(stack, char)
		}
	}
	
	// Convert the slice of runes back into a string
	var sb strings.Builder
	for _, char := range stack {
		sb.WriteRune(char)
	}
	return sb.String()
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input file")
	}
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		largestVal := runes[:12]
		i := 1
		for i < len(runes) {
			currentVal := int(runes[i] - '0')
			// appendToLast := false
			for j := 0; j < len(largestVal); j++ {
				digit := int(largestVal[j] - '0')

				if currentVal > digit {
					remaining := len(largestVal) - j
					if remaining <= len(runes)-i {
						largestVal = append(largestVal[:j], runes[i:i+remaining]...)
						break
					}
				}
			}
			i++
		}
		largestInt, err := strconv.Atoi(string(largestVal))
		if err != nil {
			fmt.Println("error converting char to string")
		}
		//fmt.Printf("largest int at line: %d\n", largestInt)
		result += largestInt
	}
	return result
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input file")
	}
	scanner := bufio.NewScanner(file)

	resultSlice := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		largestVal := []rune{runes[0], 0}
		for i := 1; i < len(runes); i++ {
			// this will return actual digit
			firstDigit := int(largestVal[0] - '0')
			secondDigit := int(largestVal[1] - '0')
			currentVal := int(runes[i] - '0')

			if currentVal > firstDigit && i+1 < len(runes) {
				largestVal[0] = runes[i]
				largestVal[1] = '0'
			} else if currentVal > secondDigit {
				largestVal[1] = runes[i]
			}
		}
		largestInt, err := strconv.Atoi(string(largestVal))
		if err != nil {
			fmt.Println("error converting char to string")
		}
		resultSlice += largestInt
	}
	fmt.Println(resultSlice)
}

func main() {
	fmt.Println(part2())
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input file")
	}
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		val, err := strconv.Atoi(largestSubsequence(text,12))
		if err != nil {
			fmt.Println("error converting char to string")
		}
		result += val 
	}
	fmt.Println(result)
}
