class Solution {
public:
    void sum(vector<vector<int>>&ans, vector<int>&ds, vector<int>&arr, int target, int ind){
        if(target ==0){
            ans.push_back(ds);
            return;
        }
        for(int i = ind; i<arr.size(); i++){
            if(i>ind && arr[i] == arr[i-1]) continue;
            if(target< arr[i])break;
            ds.push_back(arr[i]);
            sum(ans,ds,arr, target-arr[i], i+1);
            ds.pop_back();
        }
    }
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        sort(candidates.begin(), candidates.end());
        vector<vector<int>> ans;
        vector<int>ds;
        sum(ans, ds, candidates, target, 0);
        return ans;

    }
};