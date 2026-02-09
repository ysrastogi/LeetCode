class Solution {
public:
    bool isSafe(int row, int col, vector<string> board, int n){
        int duprow = row;
        int dupcol = col;
        while(col>= 0 && row>= 0){
            if(board[row][col] == 'Q') return false;
            row--;
            col--;
        }
        row = duprow;
        col = dupcol;
        while(col>= 0){
            if(board[row][col] == 'Q') return false;
            col--;
        }
        row = duprow;
        col = dupcol;
        while(row< n && col >= 0){
            if(board[row][col] == 'Q') return false;
            row++;
            col--;
        }
        return true;
    }

    void fillBoard(vector<vector<string>>&ans, vector<string>&board, int col, int n){
        if(col == n){
            ans.push_back(board);
            return;
        }
        for(int row= 0; row<n; row++){
            if(isSafe(row, col, board, n)){
                board[row][col] = 'Q';
                fillBoard(ans, board, col+1, n);
                board[row][col] = '.';

            }
        }
    }
    vector<vector<string>> solveNQueens(int n) {
        vector<vector<string>>ans;
        vector<string> board(n);
        string s(n, '.');
        for(int i=0; i<n; i++){
            board[i] = s;
        }
        fillBoard(ans, board, 0, n);
        return ans;

        
    }
};