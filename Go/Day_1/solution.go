package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	startPos := 50
	count := 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		operator := string(line[0])
		rotateVal, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		fmt.Printf("Position before: %v\n", startPos)

		switch operator {
		case "L":
			startPos -= rotateVal
		case "R":
			startPos += rotateVal
		}

		if startPos < 0 || startPos > 99 {
			startPos = int(math.Mod(float64(startPos), 100))
			if startPos < 0 {
				startPos += 100
			}
		}

		if startPos == 0 {
			count += 1
		}
		fmt.Printf("Position after: %v, rotateVal: %v\n", startPos, rotateVal)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(count)
}
