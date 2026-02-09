class Solution {
public:
    void combinationalSum3(int ind, int num, vector<vector<int>>&ans, vector<int>&ds, int k, int n){
        if(ind == k){
            if(n==0){
                ans.push_back(ds);
            }
            return;
        }
        for(int i = num+1; i<10; i++){
            if(i>n) break;
            ds.push_back(i);
            combinationalSum3(ind+1, i, ans, ds, k, n-i);
            ds.pop_back();
        }
        
    }
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<vector<int>>ans;
        vector<int>ds;
        combinationalSum3(0,0,ans, ds, k, n);
        return ans;
    }
};