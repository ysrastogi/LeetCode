class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int m = matrix.size();
        int n = matrix[0].size();
        int top = 0;
        int bottom = m-1;
        while(top<= bottom){
            int row = (top +bottom)/2;
            if(target<matrix[row][0]){
                bottom = row-1;
            }
            else if(target>matrix[row][n-1]){
                top = row+1;
            }
            else{
                break;
            }
        }
        if(top > bottom){
            return false;
        }
        int row = (top+bottom)/2;
        int left = 0;
        int right = n-1;
        while(left<= right){
            int mid = (left+right)/2;
            if(target<matrix[row][mid]){
                right = mid -1;
            }
            else if(target>matrix[row][mid]){
                left = mid+1;
            }
            else{
                return true;
            }
        }
        return false;
        
    }
};