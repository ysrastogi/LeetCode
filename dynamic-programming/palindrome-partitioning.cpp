class Solution {
public:

    bool isValidPalindrome(string s, int start, int end){
        while(start <= end){
            if(s[start] != s[end])
                return false;
            start++;
            end--;
        }
        return true;
    }
    void palindromePartioning(vector<vector<string>>&ans, vector<string>&ds, string s, int ind){
        if(ind == s.size()){
            ans.push_back(ds);
            return;
        }
        for(int i=ind; i< s.size(); ++i){
            if(isValidPalindrome(s, ind, i)){
                ds.push_back(s.substr(ind, i-ind+1));
                palindromePartioning(ans, ds,s, i+1);
                ds.pop_back();
            }
        }
    }
    vector<vector<string>> partition(string s) {
        vector<vector<string>>ans;
        vector<string> ds;
        palindromePartioning(ans, ds, s, 0);
        return ans;
        
    }
};