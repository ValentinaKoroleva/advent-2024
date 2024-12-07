package main

func goForward(cursor cursor) cursor {
	switch cursor.direction {
	case ">":
		cursor.y++
	case "v":
		cursor.x++
	case "<":
		cursor.y--
	case "^":
		cursor.x--
	}
	return cursor
}

// turnRight changes the direction to the right from the current orientation.
func turnRight(cursor cursor) cursor {
	switch cursor.direction {
	case ">":
		cursor.direction = "v"
	case "v":
		cursor.direction = "<"
	case "<":
		cursor.direction = "^"
	case "^":
		cursor.direction = ">"
	}
	return cursor
}

func findExit(matrix [][]string, start cursor, nCols int, nRows int) (int, cursor) {
	nextCursor := goForward(start)
	if nextCursor.x > nRows-1 || nextCursor.y > nCols-1 || nextCursor.x < 0 || nextCursor.y < 0 {
		res := 0
		for _, row := range matrix {
			for _, con := range row {
				if con == "X" {
					res++
				}
			}
		}
		return res + 1, start
	}
	if matrix[nextCursor.x][nextCursor.y] != "#" {
		matrix[nextCursor.x][nextCursor.y] = nextCursor.direction
		matrix[start.x][start.y] = "X"

	} else {
		nextCursor = turnRight(start)
		matrix[nextCursor.x][nextCursor.y] = nextCursor.direction

	}
	return findExit(matrix, nextCursor, nCols, nRows)
}
