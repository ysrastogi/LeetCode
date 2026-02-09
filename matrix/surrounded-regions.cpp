class Solution {
public:
    void dfs(vector<vector<char>>&board, vector<vector<int>>&vis, int r, int c){
        int m = board.size();
        int n = board[0].size();
        int drow[4] = {1, 0, -1, 0};
        int dcol[4] = {0, -1, 0, 1};
        for(int i=0; i<4; i++){
            int row = r+drow[i];
            int col = c+dcol[i];
            if(row>=0 && row<m && col>=0 && col<n){
                if(board[row][col] == 'O' && !vis[row][col]){
                    vis[row][col]=1;
                    dfs(board, vis, row, col);
                }
            }
        }

    }
    void solve(vector<vector<char>>& board) {
        int m = board.size();
        int n= board[0].size();
        vector<vector<int>>vis(m,(vector<int>(n,0)));
        for(int i=0; i<n; i++){
            if(board[0][i]=='O'){
                vis[0][i] =1;
                dfs(board, vis, 0, i);
            }
            if(board[m-1][i] == 'O'){
                vis[m-1][i] =1;
                dfs(board, vis, m-1, i);
            }
        }
        for(int i=0; i<m; i++){
            if(board[i][0] == 'O'){
                vis[i][0] =1;
                dfs(board, vis, i, 0);
            }
            if(board[i][n-1] == 'O'){
                vis[i][n-1]=1;
                dfs(board, vis, i, n-1);
            }
        }
        for(int i=0; i<m; i++){
            for(int j=0; j<n; j++){
                if(vis[i][j]== 0) board[i][j]= 'X';
            }
        }
    }
};