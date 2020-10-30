package islandPerimeter

func islandPerimeter(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				if i == 0 {
					count++
				} else if grid[i-1][j] == 0 {
					count++
				}
				if j == 0 {
					count++
				} else if grid[i][j-1] == 0 {
					count++
				}
			}
		}
	}
	return count * 2
}
