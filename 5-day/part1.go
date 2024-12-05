package main

import (
	"strings"
)

func checkRule(leftNumber string, rightNumber string, line string) bool {
	indexLeft := strings.Index(line, leftNumber)
	indexRight := strings.Index(line, rightNumber)
	isValid := true
	if indexLeft >= 0 && indexRight >= 0 && indexLeft >= indexRight {
		isValid = false
	}
	return isValid
}
