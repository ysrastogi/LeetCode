class Solution {
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>>pascal;
        for(int i=1; i<= numRows; i++){
            vector<int>triangle;
            int res = 1;
            triangle.push_back(res);
            for(int j=1; j<i; j++){
                res = res*(i-j);
                res =res/j;
                triangle.push_back(res);
            }
            pascal.push_back(triangle);
        }
        return pascal;
        
    }
};