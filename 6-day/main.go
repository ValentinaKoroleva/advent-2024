package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type cursor struct {
	direction string
	x, y      int
}

func main() {
	// Define flags
	isTest1 := flag.Bool("test1", false, "a bool")
	isTest2 := flag.Bool("test2", false, "a bool")
	flag.Parse()
	// Read file
	filename := "input.txt"
	if *isTest1 {
		filename = "example1.txt"
	}
	if *isTest2 {
		filename = "example2.txt"
	}
	file, _ := os.Open(filename)
	defer file.Close()
	// Read input
	scanner := bufio.NewScanner(file)
	var matrix [][]string
	var matrix_copy [][]string
	var startCursor cursor
	var nCols, nRows int
	re := regexp.MustCompile(`v|>|\^|<`)
	for scanner.Scan() {
		line := scanner.Text()
		// Calculate numbers of columns
		if nCols == 0 {
			nCols = len(line)
		}
		// Find start -- cursor position
		if re.MatchString(line) {
			startCursor.x = nRows
			startCursor.y = strings.Index(line, re.FindString(line))
			startCursor.direction = re.FindString(line)
		}

		matrix = append(matrix, strings.Split(line, ""))
		matrix_copy = append(matrix_copy, strings.Split(line, ""))
		// Calculate numbers of rows
		nRows++
	}
	result1, exit := findExit(matrix, startCursor, nCols, nRows)
	fmt.Println(result1)
	fmt.Println(exit)
	result2 := findLoop(matrix_copy, startCursor, nCols, nRows, startCursor, exit)
	fmt.Println(result2)
}

// 143
