package main

// Функция для поиска слова в филворде
func findCross(grid [][]string, word string) [][]int {

	var results [][]int

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == string(word[1]) {
				existsCondition := x-1 >= 0 && y-1 >= 0 && x+1 < len(grid) && y+1 < len(grid[0])
				if !existsCondition {
					continue
				}
				crossCondition := (grid[x-1][y-1] == string(word[0]) && grid[x+1][y+1] == string(word[2]) ||
					grid[x-1][y-1] == string(word[2]) && grid[x+1][y+1] == string(word[0])) &&
					(grid[x-1][y+1] == string(word[0]) && grid[x+1][y-1] == string(word[2]) ||
						grid[x-1][y+1] == string(word[2]) && grid[x+1][y-1] == string(word[0]))
				if crossCondition {
					results = append(results, []int{x, y, 0, 0})
				}
			}
		}
	}

	return results
}
