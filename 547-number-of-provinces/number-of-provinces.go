func dfs(node int, isConnected [][]int, visited []bool) {
    visited[node] = true

    for j := 0; j < len(isConnected); j++ {
        if isConnected[node][j] == 1 && !visited[j] {
            dfs(j, isConnected, visited)
        }
    }
}

func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    visited := make([]bool, n)
    provinces := 0

    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i, isConnected, visited)
            provinces++
        }
    }

    return provinces
}