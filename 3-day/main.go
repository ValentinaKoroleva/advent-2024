package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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
	// Чтение всего содержимого файла
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Преобразование содержимого файла в строку
	line := string(fileContent)

	multiplyPattern := "mul\\(\\d+\\,\\d+\\)"

	re := regexp.MustCompile(multiplyPattern)
	sum := 0
	linesByDo := strings.Split(line, "do()")
	fmt.Println(len(linesByDo))
	for i := 0; i < len(linesByDo); i++ {
		linesByDont := strings.Split(linesByDo[i], "don't()")
		multiplyLines := re.FindAllString(linesByDont[0], -1)
		for i := 0; i < len(multiplyLines); i++ {
			sum += multiplication(multiplyLines[i])
		}
	}
	fmt.Println(sum)
}

// 89798695
