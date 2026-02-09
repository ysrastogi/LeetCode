class Solution {
public:
    string kthDistinct(vector<string>& arr, int k) {
       unordered_map<string,int> mp;
        
        for(int i=0; i<arr.size(); i++){
            mp[arr[i]]++;
        }
        
        for(int i=0; i<arr.size(); i++){
            if(mp.find(arr[i])!=mp.end()){
                if(mp[arr[i]]==1){
                    k-=1;
                    if(k==0){
                        return arr[i];
                    }
                    else 
                        continue;
                }
            }
        }
        return "";
        
    }
};