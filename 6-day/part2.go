package main

// В итоге смогла вывести только 3 штуки из 6 для тестового
import "fmt"

var i = 0
var obstacles []cursor

func findLoop(matrix [][]string, start cursor, nCols int, nRows int, initialStart cursor, stop cursor) int {
	if start == stop {
		return len(obstacles)
	}

	nextCursor := goForward(start)
	oldDirection := matrix[nextCursor.x][nextCursor.y]
	if matrix[nextCursor.x][nextCursor.y] != "#" {
		matrix[nextCursor.x][nextCursor.y] = nextCursor.direction
		matrix[start.x][start.y] = start.direction
	} else {
		nextCursor = turnRight(start)
		matrix[nextCursor.x][nextCursor.y] = nextCursor.direction
	}
	if turnRight(nextCursor).direction == oldDirection && !contains(obstacles, goForward(nextCursor)) {
		// found loop
		obstacles = append(obstacles, goForward(nextCursor))
		fmt.Println(obstacles)
	}

	// if nextCursor.x == 8 && nextCursor.y == 1 {
	// for _, row := range matrix {
	// 	fmt.Println(row)
	// }
	// fmt.Println(!contains(obstacles, goForward(nextCursor)))
	// fmt.Println(nextCursor)
	// fmt.Println(oldDirection)
	// fmt.Println(stop)
	// }

	// 	fmt.Println(oldDirection)
	// 	fmt.Println(matrix[nextCursor.x][nextCursor.y])
	// }

	i++
	return findLoop(matrix, nextCursor, nCols, nRows, initialStart, stop)
}

// Функция для проверки наличия элемента в массиве
func contains(arr []cursor, elem cursor) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
func turnRight2(cursor string) string {
	switch cursor {
	case ">":
		cursor = "v"
	case "v":
		cursor = "<"
	case "<":
		cursor = "^"
	case "^":
		cursor = ">"
	}
	return cursor
}
