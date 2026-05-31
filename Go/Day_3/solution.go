package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
		largestInt,err := strconv.Atoi(string(largestVal))
		if err != nil {
			fmt.Println("error converting char to string")
		}
		fmt.Printf("largest int val: %d\n", largestInt)
		resultSlice += largestInt
	}
	fmt.Println(resultSlice)
}
