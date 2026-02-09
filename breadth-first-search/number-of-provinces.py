class Solution:
    def dfs(self, vis, node, isConnected):
        for neighbor in range(len(isConnected)):
            if isConnected[node][neighbor] and not vis[neighbor]:
                vis[neighbor] =True
                self.dfs(vis, neighbor, isConnected)

    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        n = len(isConnected)
        provinces =0
        vis = [False]*n
        for i in range(n):
            if not vis[i]:
                provinces +=1
                vis[i] = True
                self.dfs(vis, i, isConnected)

        return provinces

        