class Solution {
public:
    int numIdenticalPairs(vector<int>& nums) {
        map<int,int> m;
        for(int i=0; i<nums.size(); i++){
            m[nums[i]]++;          
        }
        map<int, int>::iterator it = m.begin();
        int out =0;
        while(it!= m.end()){
            if(it->second >1){
                int n = it->second;
                out = out+(n*(n-1)/2);
            }
            it++;
        }
        return out;
    }
};