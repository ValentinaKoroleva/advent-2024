package main

import (
	"bufio"
	"flag"
	"fmt"
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
		filename = "example.txt"
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
	pattern := "\\s+"
	re := regexp.MustCompile(pattern)
	nOfReports := 0
	nOfReports2 := 0
	for scanner.Scan() { // internally, it advances token based on sperator
		// Read line of a file and create an array of int of this line of file
		report := re.Split(scanner.Text(), -1)
		// transform array of text to array of int
		reportInt := make([]int, len(report))
		for i, v := range report {
			reportInt[i], _ = strconv.Atoi(v)
		}
		isReportSafe := checkArray(reportInt)
		if isReportSafe {
			nOfReports++
		}
		isReportSafe2 := checkArray2(reportInt)
		if isReportSafe2 {
			nOfReports2++
		}
		// fmt.Println(reportInt)
		// fmt.Println(isReportSafe2)
	}
	fmt.Println(nOfReports)
	fmt.Println(nOfReports2)
}
