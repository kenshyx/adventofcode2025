package day7

func CountPaths(grid [][]string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	dp := make([][]int, rows)
	for r := range dp {
		dp[r] = make([]int, cols)
	}

	// Find StartChar
	var sr, sc int
	found := false
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == StartChar {
				sr, sc = r, c
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return 0
	}

	// Start with 1 way at S
	dp[sr][sc] = 1

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dp[r][c] == 0 {
				continue
			}

			cell := grid[r][c]

			switch cell {

			case StartChar, BeamChar:
				// down
				nr := r + 1
				if nr < rows && grid[nr][c] != SpaceChar {
					dp[nr][c] += dp[r][c]
				}

			case SplitterChar:
				nr := r + 1
				if nr < rows {
					// down-left
					if c-1 >= 0 && grid[nr][c-1] != SpaceChar {
						dp[nr][c-1] += dp[r][c]
					}
					// down-right
					if c+1 < cols && grid[nr][c+1] != SpaceChar {
						dp[nr][c+1] += dp[r][c]
					}
				}
			}
		}
	}

	// Last row contains the number of paths
	total := 0
	for c := 0; c < cols; c++ {
		if grid[rows-1][c] == BeamChar {
			total += dp[rows-1][c]
		}
	}

	return total
}

func SolutionPart2() int {
	return CountPaths(matrix)
}
