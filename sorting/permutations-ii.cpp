class Solution {
public:

    void unique_permutations(int ind, vector<vector<int>>&ans, vector<int>nums){
        if(ind == nums.size()- 1){
            ans.push_back(nums);
            return;
        }
        for(int i = ind; i<nums.size(); i++){
            if(!((ind != i) &&(nums[ind]== nums[i]))){
                swap(nums[i], nums[ind]);
                unique_permutations(ind+1, ans, nums);
                }
            
        }
    }
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>> ans;
        sort(nums.begin(), nums.end());
        unique_permutations(0, ans, nums);
        return ans;

    }
};