class Solution {
public:
    int lengthOfLongestSubstring(string s)
    {
        int start=0,end=0,maxi=INT_MIN;
        set<char> st;
        while(end<s.length())
        {
            if(st.find(s[end])==st.end())
            {
                st.insert(s[end]);
                maxi=max(maxi,end-start+1);
            }
            else
            {
                while(s[start]!=s[end])
                {
                    st.erase(s[start]);
                    start++;
                }
                start++;
               maxi=max(maxi,end-start+1);
            }
            end++;
        }
        if(maxi==INT_MIN)
            return 0;
        return maxi;
        
    }
};