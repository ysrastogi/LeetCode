func dfs(i, j, idx int, word string, board [][]byte, m, n int) bool {
	if idx == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[idx] {
		return false
	}
	temp := board[i][j]
	board[i][j] = '#'

	found := dfs(i+1, j, idx+1, word, board, m, n) ||
		dfs(i-1, j, idx+1, word, board, m, n) ||
		dfs(i, j+1, idx+1, word, board, m, n) ||
		dfs(i, j-1, idx+1, word, board, m, n)

	board[i][j] = temp

	return found
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0, word, board, m, n) {
				return true
			}
		}
	}
	return false
}
