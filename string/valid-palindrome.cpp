class Solution {
public:
    bool isPalindrome(string s) {
        int n = s.length();
        string news;
        for(int i= 0; i<n; i++){
            if((s[i]>=97 && s[i]<=122)|| (s[i]>= 48 && s[i]<=57)){
                news.push_back(s[i]);
            }
            else if(s[i]>=65&& s[i]<=90){
                news.push_back(s[i]+32);
            }
            else{
                continue;
            }
        }
        for(int i= 0; i<news.length()/2; i++){
            if(news[i] != news[news.length()-i-1]){
                return false;
            }
        }
        return true;
    }
};