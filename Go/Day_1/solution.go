package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// ans: 6768
func main() {
	startPos := 50
	count := 0
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

	for scanner.Scan() {
		line := scanner.Text()
		operator := string(line[0])
		rotateVal, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		posBeforeRotate := startPos

		switch operator {
		case "L":
			startPos -= rotateVal
		case "R":
			startPos += rotateVal
		}

		if startPos <= 0 {
			passZero := startPos / 100
			startPos = startPos % 100
			count += int(math.Abs(float64(passZero)))
			if startPos < 0 {
				startPos += 100
				if posBeforeRotate != 0 {
					count++
				}
			}
			if startPos == 0 {
				count++
			}
		}

		if startPos >= 100 {
			passZero := startPos / 100
			count += passZero
			startPos = startPos % 100
		}
		fmt.Println(count)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
