class Solution {
public:
    int characterReplacement(string s, int k) {
        int start = 0, end = 0;
        int res = INT_MIN;
        int maxFreq=0;
        unordered_map<char, int> st;
        while(end<s.length()){
            st[s[end]]++;
             maxFreq=max(maxFreq,st[s[end]]);
            while((end-start+1)-maxFreq >k){
                st[s[start]]--;
                start++;
            }
            res = max(res, (end-start+1));
            end++;
            
        }
        return res;
    }
};