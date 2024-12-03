package main

import (
	"regexp"
	"strconv"
)

// Received a text that contains multiplication program 'mul(1,2)' and return result of multiplication
func multiplication(text string) int {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	match := re.FindStringSubmatch(text)
	if len(match) == 3 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		return a * b
	}
	return 0
}
