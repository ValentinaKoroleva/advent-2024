package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
	// Define regex
	re := regexp.MustCompile(`\d+`)
	// Read input
	fileContent, _ := io.ReadAll(file)
	parts := strings.Split(string(fileContent), "\n\n")
	rules := strings.Split(parts[0], "\n")
	updateLines := strings.Split(parts[1], "\n")
	result1 := 0
	result2 := 0
	for _, line := range updateLines {
		numbers := strings.Split(line, ",")
		middleNumber1, _ := strconv.Atoi(numbers[len(numbers)/2])
		isLineValid := true
		for _, rule := range rules {
			ruleNumbers := re.FindAllString(rule, -1)
			isValid := checkRule(ruleNumbers[0], ruleNumbers[1], line)
			if !isValid {
				isLineValid = false
				middleNumber1 = 0
				break
			}
		}
		if !isLineValid {
			middleNumber2 := fixLineAccordingToRules(line, rules)
			// fmt.Println(middleNumber2)
			result2 += middleNumber2
		}
		result1 += middleNumber1
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

// 143
