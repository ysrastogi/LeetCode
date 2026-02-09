class Solution {
public:
    void dfs(vector<vector<int>>&grid, vector<vector<int>>&vis, int r, int c){
        int m= grid.size();
        int n = grid[0].size();
        int drow[4] = {1,0,-1, 0};
        int dcol[4] = {0,-1, 0, 1};
        for(int i=0; i<4; i++){
            int row = r+drow[i];
            int col = c+dcol[i];
            if(row>=0 && row<m && col>=0 && col<n){
                if(grid[row][col] && !vis[row][col]){
                    vis[row][col] =1;
                    dfs(grid, vis, row, col);
                }
            }
        }

    }
    int numEnclaves(vector<vector<int>>& grid) {
        int m= grid.size();
        int n= grid[0].size();
        vector<vector<int>>vis(m, vector<int>(n,0));
        for(int i=0; i<n; i++){
            if(grid[0][i] == 1){
                vis[0][i] =1;
                dfs(grid, vis, 0, i);
            }
            if(grid[m-1][i] == 1){
                vis[m-1][i] = 1;
                dfs(grid, vis, m-1, i);
            }
        }
        for(int i=0; i<m; i++){
            if(grid[i][0] ==1){
                vis[i][0] = 1;
                dfs(grid, vis, i, 0);
            }
            if(grid[i][n-1]){
                vis[i][n-1] =1;
                dfs(grid, vis, i, n-1);
            }
        }
        int count =0;
        for(int i=0; i<m; i++){
            for(int j=0; j<n; j++){
             if(grid[i][j] && !vis[i][j]) count++;   
            }
        }
        return count;
    }
};