package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input file
	isTest1 := flag.Bool("test1", false, "a bool")
	isTest2 := flag.Bool("test2", false, "a bool")
	flag.Parse()
	filename := "input.txt"
	if *isTest1 {
		filename = "example1.txt"
	}
	if *isTest2 {
		filename = "example2.txt"
	}
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	arrayOfLetters := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arrayOfLetters = append(arrayOfLetters, strings.Split(line, ""))
	}
	word := "XMAS"
	results1 := findWords(arrayOfLetters, word)
	xword := "MAS"
	results2 := findCross(arrayOfLetters, xword)
	fmt.Println(len(results1))
	// fmt.Println(results2)
	fmt.Println(len(results2))
}
