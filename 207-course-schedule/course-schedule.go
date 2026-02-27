func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    for _, p := range prerequisites{
        course := p[0]
        prereq := p[1]
        graph[prereq] = append(graph[prereq], course)
    }

    state := make([]int, numCourses)

    var dfs func(int)bool
    dfs = func(node int) bool {
        if state[node] == 1{
            return false
        }
        if state[node] == 2{
            return true
        }
        state[node]=1
        for _, neighbour := range graph[node]{
            if !dfs(neighbour){
                return false
            }
        }
        state[node]=2
        return true
    }

    for i:=0; i<numCourses; i++{
        if state[i]==0{
            if !dfs(i){
                return false
            }
        }
    }
    return true
    


    
}