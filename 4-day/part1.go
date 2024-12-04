package main

// Функция для поиска слова в филворде
func findWords(grid [][]string, word string) [][]int {
	directions := [][]int{
		{0, 1},   // горизонтально
		{1, 0},   // вертикально
		{1, 1},   // диагонально вправо-вниз
		{1, -1},  // диагонально вправо-вверх
		{-1, 1},  // диагонально влево-вниз
		{-1, -1}, // диагонально влево-вверх
		{0, -1},  // горизонтально влево
		{-1, 0},  // вертикально вверх
	}

	var results [][]int

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if search(grid, word, x, y, dx, dy) {
					results = append(results, []int{x, y, dx, dy})
				}
			}
		}
	}

	return results
}

// Вспомогательная функция для проверки наличия слова в заданном направлении
func search(grid [][]string, word string, x, y, dx, dy int) bool {
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		nx, ny := x+i*dx, y+i*dy
		if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) || grid[nx][ny] != string(word[i]) {
			return false
		}
	}
	return true
}
