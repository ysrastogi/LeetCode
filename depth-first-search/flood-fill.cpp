class Solution {
public:
    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int color) {
        int m = image.size();
        int n = image[0].size();
        vector<vector<int>>pixels = image;
        queue<pair<int, int>> q;
        q.push({sr, sc});
        pixels[sr][sc]=color;
        q.push({sr,sc});
        int drow[4] ={1, 0, -1, 0};
        int dcol[4] ={0, -1, 0, 1};
        while(!q.empty()){
            int r = q.front().first;
            int c = q.front().second;
            q.pop();
            for(int i= 0; i<4; i++){
                int row = r+drow[i];
                int col = c+dcol[i];
                if(row>=0 && row<m && col>=0 && col<n && image[row][col] == image[r][c] && pixels[row][col] != color){
                    q.push({row, col});
                    pixels[row][col] = color;
                }
            }
        }
        return pixels;
    }
};