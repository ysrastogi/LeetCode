class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int n = nums.size();
        int prod = 1;
        int maxi = INT_MIN; 
        for(int i = 0; i<n; i++){
            maxi = max(prod *= nums[i], maxi);
            if(nums[i] == 0){
                prod =1;
            }
        }
        prod =1;
        for(int i= n-1; i>= 0; i--){
            maxi = max(prod*= nums[i], maxi);
            if(nums[i] == 0){
                prod =1;
            }
        }
        return maxi;
    }
};