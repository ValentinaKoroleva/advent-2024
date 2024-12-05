package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func fixLineAccordingToRules(line string, rules []string) int {
	cmp := func(a, b string) int {
		for _, s := range rules {
			if s := strings.Split(s, "|"); s[0] == a && s[1] == b {
				return -1
			}
		}
		return 0
	}
	numbers := regexp.MustCompile("[0-9]+").FindAllString(line, -1)
	slices.SortFunc(numbers, cmp)
	var middleNumber, _ = strconv.Atoi(numbers[len(numbers)/2])
	return middleNumber
}
