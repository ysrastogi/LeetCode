class Solution {
public:
    int maxWidthOfVerticalArea(vector<vector<int>>& points) {
        vector<int>x;
        int max = 0;
        int n= points.size();
        for(int i=0; i<n; i++){
            x.push_back(points[i][0]);
        }
        sort(x.begin(), x.end());
        for(int i= 0; i<n-1; i++){
            if((x[i+1]- x[i])> max){
                max = (x[i+1]-x[i]);
            }
        }
        
        return max;
        
    }
};