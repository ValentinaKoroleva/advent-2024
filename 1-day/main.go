package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read input file
	isTest := flag.Bool("test", false, "a bool")
	flag.Parse()
	filename := "input.txt"
	if *isTest {
		filename = "example1.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// Parse input file
	// Fill the array
	scanner := bufio.NewScanner(file)
	var arrA, arrB []int
	pattern := "\\s+"
	re := regexp.MustCompile(pattern)
	for scanner.Scan() { // internally, it advances token based on sperator
		line := re.Split(scanner.Text(), -1)
		a, err := strconv.Atoi(line[0])
		if err != nil {
			// ... handle error
			panic(err)
		}
		b, err := strconv.Atoi(line[1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		arrA = append(arrA, a)
		arrB = append(arrB, b)
	}
	// Find total d
	findTotalDistance(arrA, arrB)
	// Find similarity
	findSimilarity(arrA, arrB)
}
